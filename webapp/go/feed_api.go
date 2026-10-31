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
		candidates := s.feedCandidatesLocked(viewer, following, input.Feed, ranked, beforePostedAt, beforeID)
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

// feedCandidatesLocked は、受付中の投稿期間の投稿を新しい順にたどり、閲覧できてフィードに合う投稿を返す。
// friends・following・friends-of-friends は matchesFeed に合う投稿者が決まっているので、投稿期間ごとの投稿者の索引からその投稿者の投稿だけを集める。
// public は投稿期間ごとの並んだ投稿を新しい側からまとめてたどり、51 件で止める
func (s *Store) feedCandidatesLocked(viewer viewerContext, following map[string]bool, feed string, ranked bool, beforePostedAt *time.Time, beforeID string) []feedCandidate {
	candidates := make([]feedCandidate, 0)
	// 自分以外の投稿は、今日投稿していなければ canView を通らない
	if !viewer.HasPostedToday {
		return candidates
	}
	older := func(p *PostRow) bool {
		return beforePostedAt == nil || p.PostedAt.Before(*beforePostedAt) || (p.PostedAt.Equal(*beforePostedAt) && p.ID < beforeID)
	}
	keep := func(p *PostRow) bool {
		return p.DeletedAt == nil && p.AuthorID != viewer.ID && older(p)
	}
	// visit は候補に足し、続けてたどるかを返す
	visit := func(p *PostRow) bool {
		author := s.activeUser(p.AuthorID)
		if author == nil {
			return true
		}
		post := p.post()
		if viewer.canView(post) && matchesFeed(viewer, following, post, feed) {
			candidates = append(candidates, feedCandidate{post: post, author: author.User, row: p})
			if !ranked && len(candidates) == 51 {
				return false
			}
		}
		return true
	}
	if feed == "public" {
		lists := make([][]*PostRow, 0, len(s.activeWindows))
		for _, w := range s.activeWindows {
			// w.posts は古い順なので、ページの境目より古い投稿は先頭からの範囲になる
			lists = append(lists, w.posts[:sort.Search(len(w.posts), func(i int) bool { return !older(w.posts[i]) })])
		}
		for {
			best := -1
			for i, l := range lists {
				if len(l) > 0 && (best < 0 || postLess(lists[best][len(lists[best])-1], l[len(l)-1])) {
					best = i
				}
			}
			if best < 0 {
				return candidates
			}
			l := lists[best]
			p := l[len(l)-1]
			lists[best] = l[:len(l)-1]
			if keep(p) && !visit(p) {
				return candidates
			}
		}
	}
	var authors map[string]bool
	switch feed {
	case "friends":
		authors = viewer.Friends
	case "following":
		authors = following
	case "friends-of-friends":
		authors = viewer.FriendsOfFriends
	}
	var rows []*PostRow
	for _, w := range s.activeWindows {
		for id := range authors {
			if feed == "friends-of-friends" && viewer.Friends[id] {
				continue
			}
			for _, p := range w.byAuthor[id] {
				if keep(p) {
					rows = append(rows, p)
				}
			}
		}
	}
	sort.Slice(rows, func(i, j int) bool { return postLess(rows[j], rows[i]) })
	for _, p := range rows {
		if !visit(p) {
			break
		}
	}
	return candidates
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
