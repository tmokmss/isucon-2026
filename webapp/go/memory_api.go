package isureal

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v5"
)

type getMemoriesRequest struct {
	Month          string     `query:"month"`
	BeforePostedAt *time.Time `query:"beforePostedAt"`
	BeforeID       string     `query:"beforeId"`
}

type memoryPostResponse struct {
	createPostResponse
	RealMojis []replaceRealMojiResponse `json:"realmojis"`
	Comments  []createCommentResponse   `json:"comments"`
}

type getMemoriesResponse struct {
	Items   []memoryPostResponse `json:"items"`
	HasMore bool                 `json:"hasMore"`
}

type getMemoriesByDateRequest struct {
	Date           string     `param:"date"`
	BeforePostedAt *time.Time `query:"beforePostedAt"`
	BeforeID       string     `query:"beforeId"`
}

type getMemoriesByDateResponse getMemoriesResponse

func getMemoriesHandler(h *HandlerContext) echo.HandlerFunc {
	return func(c *echo.Context) error {
		var input getMemoriesRequest
		if err := c.Bind(&input); err != nil {
			return err
		}
		start, err := time.Parse("2006-01", input.Month)
		if err != nil {
			return NewAPIError(http.StatusBadRequest, "invalid_memory_period")
		}
		response, err := listMemories(h.Store, GetSession(c).UserID, start, start.AddDate(0, 1, 0), input.BeforePostedAt, input.BeforeID)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, response)
	}
}

func getMemoriesByDateHandler(h *HandlerContext) echo.HandlerFunc {
	return func(c *echo.Context) error {
		var input getMemoriesByDateRequest
		if err := c.Bind(&input); err != nil {
			return err
		}
		start, err := time.Parse(time.DateOnly, input.Date)
		if err != nil {
			return NewAPIError(http.StatusBadRequest, "invalid_memory_period")
		}
		response, err := listMemories(h.Store, GetSession(c).UserID, start, start.AddDate(0, 0, 1), input.BeforePostedAt, input.BeforeID)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, getMemoriesByDateResponse(response))
	}
}

type memoryEntry struct {
	post      Post
	author    User
	realmojis []realMojiWithReactor
	comments  []commentWithAuthor
}

func listMemories(s *Store, userID string, start, end time.Time, beforeAt *time.Time, beforeID string) (getMemoriesResponse, error) {
	var response getMemoriesResponse
	if (beforeAt == nil) != (beforeID == "") {
		return response, NewAPIError(http.StatusBadRequest, "invalid_pagination")
	}
	entries := make([]memoryEntry, 0)
	s.mu.RLock()
	u := s.activeUser(userID)
	if u != nil {
		for i := len(u.posts) - 1; i >= 0; i-- {
			p := u.posts[i]
			if p.DeletedAt != nil {
				continue
			}
			d := p.window.LocalDate
			if d.Before(start) || !d.Before(end) {
				continue
			}
			if beforeAt != nil && !(p.PostedAt.Before(*beforeAt) || (p.PostedAt.Equal(*beforeAt) && p.ID < beforeID)) {
				continue
			}
			entries = append(entries, memoryEntry{post: p.post(), author: u.User})
			if len(entries) == 51 {
				break
			}
		}
		if len(entries) > 50 {
			response.HasMore = true
			entries = entries[:50]
		}
		for i := range entries {
			p := s.posts[entries[i].post.ID]
			entries[i].realmojis = s.currentRealMojisLocked(p, nil, "", 50)
			entries[i].comments = s.commentsLocked(p, nil, "", 50)
		}
	}
	s.mu.RUnlock()
	items := make([]memoryPostResponse, 0, len(entries))
	for _, e := range entries {
		item := memoryPostResponse{
			createPostResponse: newPostResponse(e.post, e.author),
			RealMojis:          make([]replaceRealMojiResponse, 0, len(e.realmojis)),
			Comments:           make([]createCommentResponse, 0, len(e.comments)),
		}
		for _, r := range e.realmojis {
			item.RealMojis = append(item.RealMojis, newRealMojiResponse(r.reaction, r.reactor))
		}
		for _, cm := range e.comments {
			item.Comments = append(item.Comments, newCommentResponse(cm.comment, cm.author))
		}
		items = append(items, item)
	}
	response.Items = items
	return response, nil
}
