package isureal

import (
	"net/http"
	"sort"
	"time"

	"github.com/labstack/echo/v5"
)

type getRelationshipsRequest struct {
	Kind            string     `param:"kind"`
	BeforeUpdatedAt *time.Time `query:"beforeUpdatedAt"`
	BeforeID        string     `query:"beforeId"`
}

type relationshipUserResponse struct {
	ID              string  `json:"id"`
	Username        string  `json:"username"`
	DisplayName     string  `json:"displayName"`
	ProfileImageURL *string `json:"profileImageUrl"`
	Bio             string  `json:"bio"`
	AccountType     string  `json:"accountType"`
}

type relationshipResponse struct {
	RelationID string                   `json:"relationId"`
	User       relationshipUserResponse `json:"user"`
	Status     string                   `json:"status"`
	UpdatedAt  time.Time                `json:"updatedAt"`
}

type getRelationshipsResponse struct {
	Items   []relationshipResponse `json:"items"`
	HasMore bool                   `json:"hasMore"`
}

type relationRecord struct {
	ID        string
	UpdatedAt time.Time
	OtherID   string
	user      User
}

func getRelationshipsHandler(h *HandlerContext) echo.HandlerFunc {
	return func(c *echo.Context) error {
		var input getRelationshipsRequest
		if err := c.Bind(&input); err != nil {
			return err
		}
		if (input.BeforeUpdatedAt == nil) != (input.BeforeID == "") {
			return NewAPIError(http.StatusBadRequest, "invalid_pagination")
		}
		switch input.Kind {
		case "friends", "following", "followers", "incoming-requests", "outgoing-requests":
		default:
			return NewAPIError(http.StatusBadRequest, "invalid_relationship")
		}
		actor := GetSession(c).UserID
		s := h.Store
		var records []relationRecord
		s.mu.RLock()
		if u := s.users[actor]; u != nil {
			switch input.Kind {
			case "friends":
				for other, f := range u.friends {
					records = append(records, relationRecord{ID: f.ID, UpdatedAt: f.UpdatedAt, OtherID: other})
				}
			case "following":
				for other, f := range u.following {
					records = append(records, relationRecord{ID: f.ID, UpdatedAt: f.UpdatedAt, OtherID: other})
				}
			case "followers":
				for other, f := range u.followers {
					records = append(records, relationRecord{ID: f.ID, UpdatedAt: f.UpdatedAt, OtherID: other})
				}
			case "incoming-requests":
				for other, r := range u.pendingIn {
					records = append(records, relationRecord{ID: r.ID, UpdatedAt: r.UpdatedAt, OtherID: other})
				}
			case "outgoing-requests":
				for other, r := range u.pendingOut {
					records = append(records, relationRecord{ID: r.ID, UpdatedAt: r.UpdatedAt, OtherID: other})
				}
			}
		}
		filtered := records[:0]
		for _, r := range records {
			if input.BeforeUpdatedAt != nil && !(r.UpdatedAt.Before(*input.BeforeUpdatedAt) || (r.UpdatedAt.Equal(*input.BeforeUpdatedAt) && r.ID < input.BeforeID)) {
				continue
			}
			other := s.activeUser(r.OtherID)
			if other == nil {
				continue
			}
			r.user = other.User
			filtered = append(filtered, r)
		}
		s.mu.RUnlock()
		sort.Slice(filtered, func(i, j int) bool {
			if !filtered[i].UpdatedAt.Equal(filtered[j].UpdatedAt) {
				return filtered[i].UpdatedAt.After(filtered[j].UpdatedAt)
			}
			return filtered[i].ID > filtered[j].ID
		})
		if len(filtered) > 51 {
			filtered = filtered[:51]
		}
		items := make([]relationshipResponse, 0, len(filtered))
		for _, record := range filtered {
			user := record.user
			if user.AccountType != "certified_public" && input.Kind != "friends" {
				user.Bio = ""
			}
			status := input.Kind
			if input.Kind == "incoming-requests" || input.Kind == "outgoing-requests" {
				status = "pending"
			}
			items = append(items, relationshipResponse{
				RelationID: record.ID, Status: status, UpdatedAt: record.UpdatedAt,
				User: newRelationshipUserResponse(user),
			})
		}
		hasMore := len(items) > 50
		if hasMore {
			items = items[:50]
		}
		return c.JSON(http.StatusOK, getRelationshipsResponse{Items: items, HasMore: hasMore})
	}
}
