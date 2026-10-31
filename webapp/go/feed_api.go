package isureal

import (
	"net/http"
	"sort"
	"time"

	"github.com/labstack/echo/v5"
)

type getFeedRequest struct {
	Feed           string     `param:"feed"`
	BeforeScore    *int64     `query:"beforeScore"`
	BeforePostedAt *time.Time `query:"beforePostedAt"`
	BeforeID       string     `query:"beforeId"`
}

type feedItemResponse struct {
	createPostResponse
	Score *int64 `json:"score,omitempty"`
}

type getFeedResponse struct {
	Items   []feedItemResponse `json:"items"`
	HasMore bool               `json:"hasMore"`
}

type feedCandidate struct {
	post   Post
	author User
	row    *PostRow
}

func getFeedHandler(h *HandlerContext) echo.HandlerFunc {
	return func(c *echo.Context) error {
		var input getFeedRequest
		if err := c.Bind(&input); err != nil {
			return err
		}
		switch input.Feed {
		case "friends", "friends-of-friends", "public", "following":
		default:
			return NewAPIError(http.StatusBadRequest, "invalid_post")
		}
		if (input.BeforePostedAt == nil) != (input.BeforeID == "") {
			return NewAPIError(http.StatusBadRequest, "invalid_pagination")
		}
		ranked := input.Feed == "friends-of-friends"
		if ranked && (input.BeforeScore == nil) != (input.BeforeID == "") {
			return NewAPIError(http.StatusBadRequest, "invalid_pagination")
		}
		s := h.Store
		s.mu.RLock()
		viewer, ok := s.loadViewerContextLocked(GetSession(c).UserID)
		if !ok {
			s.mu.RUnlock()
			return c.JSON(http.StatusOK, getFeedResponse{Items: []feedItemResponse{}})
		}
		following := map[string]bool{}
		if input.Feed == "following" {
			for id := range s.users[viewer.ID].following {
				following[id] = true
			}
		}
		beforePostedAt, beforeID := input.BeforePostedAt, input.BeforeID
		if ranked {
			beforePostedAt, beforeID = nil, ""
		}
		var rows []*PostRow
		for _, w := range s.activeWindows {
			for _, p := range w.posts {
				if p.DeletedAt != nil || p.AuthorID == viewer.ID {
					continue
				}
				if beforePostedAt != nil && !(p.PostedAt.Before(*beforePostedAt) || (p.PostedAt.Equal(*beforePostedAt) && p.ID < beforeID)) {
					continue
				}
				rows = append(rows, p)
			}
		}
		sort.Slice(rows, func(i, j int) bool { return postLess(rows[j], rows[i]) })
		candidates := make([]feedCandidate, 0)
		for _, p := range rows {
			author := s.activeUser(p.AuthorID)
			if author == nil {
				continue
			}
			post := p.post()
			if viewer.canView(post) && matchesFeed(viewer, following, post, input.Feed) {
				candidates = append(candidates, feedCandidate{post: post, author: author.User, row: p})
				if !ranked && len(candidates) == 51 {
					break
				}
			}
		}
		var scores map[string]int64
		if ranked {
			scores = s.recommendationScoresLocked(viewer, candidates)
		}
		s.mu.RUnlock()
		if ranked {
			sort.SliceStable(candidates, func(i, j int) bool {
				return scores[candidates[i].post.ID] > scores[candidates[j].post.ID]
			})
		}
		items := make([]feedItemResponse, 0)
		for _, cand := range candidates {
			score := scores[cand.post.ID]
			if ranked && input.BeforeScore != nil && !rankedBefore(score, cand.post, *input.BeforeScore, *input.BeforePostedAt, input.BeforeID) {
				continue
			}
			item := feedItemResponse{createPostResponse: newPostResponse(cand.post, cand.author)}
			if ranked {
				item.Score = &score
			}
			items = append(items, item)
			if len(items) == 51 {
				break
			}
		}
		hasMore := len(items) > 50
		if hasMore {
			items = items[:50]
		}
		return c.JSON(http.StatusOK, getFeedResponse{Items: items, HasMore: hasMore})
	}
}

func rankedBefore(score int64, post Post, beforeScore int64, beforePostedAt time.Time, beforeID string) bool {
	if score != beforeScore {
		return score < beforeScore
	}
	if !post.PostedAt.Equal(beforePostedAt) {
		return post.PostedAt.Before(beforePostedAt)
	}
	return post.ID < beforeID
}

func (s *Store) recommendationScoresLocked(viewer viewerContext, candidates []feedCandidate) map[string]int64 {
	scores := make(map[string]int64, len(candidates))
	connections := map[string]int64{}
	for _, cand := range candidates {
		authorID := cand.post.AuthorID
		if _, ok := connections[authorID]; ok {
			continue
		}
		var sum int64
		if a := s.users[authorID]; a != nil {
			for f := range a.friends {
				sum += int64(viewer.MutualFriends[f])
			}
		}
		connections[authorID] = sum
	}
	for _, cand := range candidates {
		var engaged int64
		seen := map[string]bool{}
		for _, r := range cand.row.realmojis {
			if !r.current() || seen[r.ReactorID] || s.activeUser(r.ReactorID) == nil {
				continue
			}
			seen[r.ReactorID] = true
			if viewer.Friends[r.ReactorID] {
				engaged++
			}
		}
		for _, cm := range cand.row.comments {
			if cm.DeletedAt != nil || seen[cm.AuthorID] || s.activeUser(cm.AuthorID) == nil {
				continue
			}
			seen[cm.AuthorID] = true
			if viewer.Friends[cm.AuthorID] {
				engaged++
			}
		}
		scores[cand.post.ID] = 10*engaged + 3*int64(viewer.MutualFriends[cand.post.AuthorID]) + min(connections[cand.post.AuthorID], 30)
	}
	return scores
}

func matchesFeed(viewer viewerContext, following map[string]bool, post Post, feed string) bool {
	switch feed {
	case "public":
		return post.Visibility == "public"
	case "following":
		return following[post.AuthorID]
	case "friends":
		return viewer.Friends[post.AuthorID]
	case "friends-of-friends":
		return !viewer.Friends[post.AuthorID] && viewer.FriendsOfFriends[post.AuthorID]
	default:
		return false
	}
}
