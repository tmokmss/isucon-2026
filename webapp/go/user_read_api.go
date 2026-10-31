package isureal

import (
	"net/http"
	"sort"
	"strings"

	"github.com/labstack/echo/v5"
	"golang.org/x/text/cases"
	"golang.org/x/text/unicode/norm"
)

type searchUsersRequest struct {
	Q              string `query:"q"`
	BeforeUsername string `query:"beforeUsername"`
	BeforeID       string `query:"beforeId"`
}

type searchUsersResponse struct {
	Items   []relationshipUserResponse `json:"items"`
	HasMore bool                       `json:"hasMore"`
}

func normalizeSearchText(value string) string {
	return norm.NFC.String(cases.Fold().String(norm.NFC.String(value)))
}

func searchUsersHandler(h *HandlerContext) echo.HandlerFunc {
	return func(c *echo.Context) error {
		var input searchUsersRequest
		if err := c.Bind(&input); err != nil {
			return err
		}
		if !validUserText(input.Q, 1, 100) {
			return NewAPIError(http.StatusBadRequest, "invalid_relationship")
		}
		if (input.BeforeUsername == "") != (input.BeforeID == "") {
			return NewAPIError(http.StatusBadRequest, "invalid_pagination")
		}
		viewer := GetSession(c).UserID
		query := normalizeSearchText(input.Q)
		s := h.Store
		s.mu.RLock()
		users := s.searchUsersLocked(viewer, query, input.BeforeUsername, input.BeforeID)
		s.mu.RUnlock()
		items := make([]relationshipUserResponse, 0, len(users))
		for _, user := range users {
			if user.AccountType != "certified_public" {
				user.Bio = ""
			}
			items = append(items, newRelationshipUserResponse(user))
		}
		hasMore := len(items) > 50
		if hasMore {
			items = items[:50]
		}
		return c.JSON(http.StatusOK, searchUsersResponse{Items: items, HasMore: hasMore})
	}
}

// 利用者名の順に、条件に合う利用者を 51 人まで返す。読みロックを持って呼ぶ
func (s *Store) searchUsersLocked(viewer, query, beforeUsername, beforeID string) []User {
	users := make([]User, 0, 51)
	list := s.usersSorted
	start := 0
	if beforeUsername != "" {
		start = sort.Search(len(list), func(i int) bool {
			u := list[i].u
			return u.Username > beforeUsername || (u.Username == beforeUsername && u.ID > beforeID)
		})
	}
	v := s.users[viewer]
	for i := start; i < len(list); i++ {
		e := &list[i]
		if !e.active || e.u == v {
			continue
		}
		if !strings.Contains(e.text[:e.nameLen], query) && !strings.Contains(e.text[e.nameLen:], query) {
			continue
		}
		u := e.u
		if v != nil && (v.blockOut[u.ID] != nil || v.blockIn[u.ID] != nil) {
			continue
		}
		users = append(users, u.User)
		if len(users) == 51 {
			break
		}
	}
	return users
}

type getUserRequest struct {
	UserID string `param:"userId"`
}

type getUserResponse relationshipUserResponse

func getUserHandler(h *HandlerContext) echo.HandlerFunc {
	return func(c *echo.Context) error {
		var input getUserRequest
		if err := c.Bind(&input); err != nil {
			return err
		}
		viewer := GetSession(c).UserID
		s := h.Store
		s.mu.RLock()
		u := s.activeUser(input.UserID)
		if u == nil {
			s.mu.RUnlock()
			return NewAPIError(http.StatusNotFound, "relationship_not_found")
		}
		user := u.User
		blocked := s.isBlocked(viewer, user.ID)
		isFriend := false
		if v := s.users[viewer]; v != nil {
			isFriend = v.friends[user.ID] != nil
		}
		s.mu.RUnlock()
		if blocked {
			return NewAPIError(http.StatusNotFound, "relationship_not_found")
		}
		if user.AccountType != "certified_public" && !isFriend {
			user.Bio = ""
		}
		return c.JSON(http.StatusOK, getUserResponse(newRelationshipUserResponse(user)))
	}
}

func newRelationshipUserResponse(user User) relationshipUserResponse {
	return relationshipUserResponse{ID: user.ID, Username: user.Username, DisplayName: user.DisplayName, ProfileImageURL: profileImageURL(user), Bio: user.Bio, AccountType: user.AccountType}
}
