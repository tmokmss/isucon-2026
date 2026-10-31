package isureal

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v5"
)

type createPostRequest struct {
	// 画像の本体は readPostImage が読む。この欄は、image がファイルでない値で来たときに Bind が元と同じく 400 を返すために残す
	Image        *multipart.FileHeader `form:"image"`
	Caption      string                `form:"caption"`
	Visibility   string                `form:"visibility"`
	LocationKind string                `form:"location_kind"`
	Latitude     *float64              `form:"latitude"`
	Longitude    *float64              `form:"longitude"`
	Locality     *string               `form:"locality"`
}

type postAuthorResponse struct {
	ID              string  `json:"id"`
	Username        string  `json:"username"`
	DisplayName     string  `json:"displayName"`
	ProfileImageURL *string `json:"profileImageUrl"`
	Bio             string  `json:"bio"`
	AccountType     string  `json:"accountType"`
}

type createPostResponse struct {
	ID              string             `json:"id"`
	Author          postAuthorResponse `json:"author"`
	PostingWindowID string             `json:"postingWindowId"`
	ImageURL        string             `json:"imageUrl"`
	Caption         string             `json:"caption"`
	Visibility      string             `json:"visibility"`
	Location        PostLocation       `json:"location"`
	PostedAt        time.Time          `json:"postedAt"`
	OnTime          bool               `json:"onTime"`
	LateSeconds     int64              `json:"lateSeconds"`
}

func createPostHandler(h *HandlerContext) echo.HandlerFunc {
	return func(c *echo.Context) error {
		imageData, hasImage, err := readPostImage(c)
		if err != nil {
			return err
		}
		var input createPostRequest
		if err := c.Bind(&input); err != nil {
			return err
		}
		if !hasImage || !validUserText(input.Caption, 0, 1000) {
			return NewAPIError(http.StatusBadRequest, "invalid_post")
		}
		switch input.Visibility {
		case "friends", "friends_of_friends", "public":
		default:
			return NewAPIError(http.StatusBadRequest, "invalid_post")
		}
		location := PostLocation{
			Kind: input.LocationKind, Latitude: input.Latitude,
			Longitude: input.Longitude, Locality: input.Locality,
		}
		if !location.valid() {
			return NewAPIError(http.StatusBadRequest, "invalid_post")
		}
		width, height, err := checkPostJPEG(c.Request(), imageData)
		if err != nil {
			return NewAPIError(http.StatusBadRequest, "invalid_post")
		}
		if width > 4096 || height > 4096 {
			return NewAPIError(http.StatusBadRequest, "invalid_post")
		}

		userID := GetSession(c).UserID
		s := h.Store
		unlock := lockUsers(userID)
		defer unlock()
		s.mu.RLock()
		u := s.activeUser(userID)
		if u == nil {
			s.mu.RUnlock()
			return NewAPIError(http.StatusNotFound, "account_not_found")
		}
		user := u.User
		window := s.activeWindows[user.Region]
		if window == nil {
			s.mu.RUnlock()
			return NewAPIError(http.StatusConflict, "post_slot_unavailable")
		}
		active, deleted := postCountsLocked(u, window.LocalDate)
		windowID, startsAt := window.ID, window.StartsAt
		s.mu.RUnlock()
		postedAt := now()
		if active != 0 || deleted >= 2 || postedAt.Before(startsAt) {
			return NewAPIError(http.StatusConflict, "post_slot_unavailable")
		}
		imageID := uuid.NewString()
		postID := uuid.NewString()
		if err := storeNewImage(imageID, imageData); err != nil {
			return err
		}
		p := &PostRow{
			ID: postID, AuthorID: userID, WindowID: windowID, ImageID: imageID, Caption: input.Caption,
			Visibility: input.Visibility, Location: location, PostedAt: postedAt, CreatedAt: postedAt,
		}
		imageCache.put(imageID, imageData, "image/jpeg")
		s.mu.Lock()
		rw := s.commitLocked(newOpPostAdd(p))
		post := p.post()
		s.mu.Unlock()
		rw.wait()
		return c.JSON(http.StatusCreated, newPostResponse(post, user))
	}
}

// multipart の本文を自分で読み、画像のパート（最初の image）を本文の長さで確保した 1 つの領域へ 1 回のコピーで取る。
// ParseMultipartForm は画像を伸ばしながら複写し、ReadAll でもう一度複写するので、一斉の投稿で asia1 の CPU と GC を食う。
// 画像以外のフィールドは ReadForm と同じ形で MultipartForm に置き、Bind に任せる（Bind は MultipartForm があれば本文を読まない）
func readPostImage(c *echo.Context) ([]byte, bool, error) {
	req := c.Request()
	if !strings.HasPrefix(req.Header.Get(echo.HeaderContentType), echo.MIMEMultipartForm) {
		return nil, false, nil
	}
	mr, err := req.MultipartReader()
	if err != nil {
		return nil, false, echo.ErrBadRequest.Wrap(err)
	}
	form := &multipart.Form{Value: map[string][]string{}, File: map[string][]*multipart.FileHeader{}}
	req.MultipartForm = form
	var image bytes.Buffer
	hasImage := false
	for {
		p, err := mr.NextPart()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, false, echo.ErrBadRequest.Wrap(err)
		}
		name := p.FormName()
		switch {
		case name == "":
		case p.FileName() == "":
			var b strings.Builder
			if _, err := io.Copy(&b, p); err != nil {
				return nil, false, echo.ErrBadRequest.Wrap(err)
			}
			form.Value[name] = append(form.Value[name], b.String())
		case name == "image" && !hasImage:
			hasImage = true
			image.Grow(max(int(req.ContentLength), 0) + bytes.MinRead)
			if _, err := image.ReadFrom(p); err != nil {
				return nil, false, echo.ErrBadRequest.Wrap(err)
			}
		}
	}
	return image.Bytes(), hasImage, nil
}

func newPostResponse(post Post, user User) createPostResponse {
	response := createPostResponse{
		ID: post.ID,
		Author: postAuthorResponse{
			ID: user.ID, Username: user.Username, DisplayName: user.DisplayName,
			ProfileImageURL: profileImageURL(user),
			Bio:             user.Bio, AccountType: user.AccountType,
		},
		PostingWindowID: post.PostingWindowID, ImageURL: signedImageURL(user, post.ImageID),
		Caption: post.Caption, Visibility: post.Visibility, Location: post.PostLocation,
		PostedAt: post.PostedAt, OnTime: !post.PostedAt.After(post.OnTimeDeadline),
	}
	if !response.OnTime {
		response.LateSeconds = int64(post.PostedAt.Sub(post.OnTimeDeadline) / time.Second)
	}
	return response
}

type getMyCurrentPostResponse createPostResponse

func getMyCurrentPostHandler(h *HandlerContext) echo.HandlerFunc {
	return func(c *echo.Context) error {
		userID := GetSession(c).UserID
		s := h.Store
		s.mu.RLock()
		u := s.users[userID]
		var p *PostRow
		if u != nil {
			p = s.currentPostLocked(u)
		}
		if p == nil {
			s.mu.RUnlock()
			return NewAPIError(http.StatusNotFound, "post_not_found")
		}
		post := p.post()
		active := u.active()
		user := u.User
		s.mu.RUnlock()
		if !active {
			return NewAPIError(http.StatusInternalServerError, "user_not_found")
		}
		return c.JSON(http.StatusOK, getMyCurrentPostResponse(newPostResponse(post, user)))
	}
}

type getPostRequest struct {
	PostID string `param:"postId"`
}

type getPostResponse createPostResponse

func getPostHandler(h *HandlerContext) echo.HandlerFunc {
	return func(c *echo.Context) error {
		var input getPostRequest
		if err := c.Bind(&input); err != nil {
			return err
		}
		s := h.Store
		s.mu.RLock()
		p, author := s.livePostWithAuthorLocked(input.PostID)
		if p == nil {
			s.mu.RUnlock()
			return NewAPIError(http.StatusNotFound, "post_not_found")
		}
		post := p.post()
		authorUser := author.User
		viewer, ok := s.loadViewerContextLocked(GetSession(c).UserID)
		s.mu.RUnlock()
		if !ok || !viewer.canView(post) {
			return NewAPIError(http.StatusNotFound, "post_not_found")
		}
		return c.JSON(http.StatusOK, getPostResponse(newPostResponse(post, authorUser)))
	}
}

type updatePostRequest struct {
	PostID     string  `param:"postId" json:"-"`
	Caption    *string `json:"caption"`
	Visibility *string `json:"visibility"`
}

type updatePostResponse createPostResponse

func updatePostHandler(h *HandlerContext) echo.HandlerFunc {
	return func(c *echo.Context) error {
		var input updatePostRequest
		if err := c.Bind(&input); err != nil {
			return err
		}
		if input.Caption == nil && input.Visibility == nil {
			return NewAPIError(http.StatusBadRequest, "invalid_post")
		}
		if input.Caption != nil && !validUserText(*input.Caption, 0, 1000) {
			return NewAPIError(http.StatusBadRequest, "invalid_post")
		}
		if input.Visibility != nil {
			switch *input.Visibility {
			case "friends", "friends_of_friends", "public":
			default:
				return NewAPIError(http.StatusBadRequest, "invalid_post")
			}
		}
		userID := GetSession(c).UserID
		s := h.Store
		unlock := lockUsers(userID)
		defer unlock()
		s.mu.RLock()
		p := s.livePostLocked(input.PostID)
		if p == nil {
			s.mu.RUnlock()
			return NewAPIError(http.StatusNotFound, "post_not_found")
		}
		authorID, ended := p.AuthorID, p.window.EndedAt != nil
		u := s.activeUser(userID)
		s.mu.RUnlock()
		if authorID != userID || ended {
			return NewAPIError(http.StatusForbidden, "post_not_editable")
		}
		if u == nil {
			return NewAPIError(http.StatusInternalServerError, "user_not_found")
		}
		s.mu.Lock()
		rw := s.commitLocked(&opPostUpdate{ID: p.ID, Caption: opt(input.Caption), Visibility: opt(input.Visibility)})
		post := p.post()
		user := u.User
		s.mu.Unlock()
		rw.wait()
		return c.JSON(http.StatusOK, updatePostResponse(newPostResponse(post, user)))
	}
}

type deletePostRequest struct {
	PostID string `param:"postId" json:"-"`
}

func deletePostHandler(h *HandlerContext) echo.HandlerFunc {
	return func(c *echo.Context) error {
		var input deletePostRequest
		if err := c.Bind(&input); err != nil {
			return err
		}
		userID := GetSession(c).UserID
		s := h.Store
		unlock := lockUsers(userID)
		defer unlock()
		s.mu.RLock()
		p := s.livePostLocked(input.PostID)
		if p == nil {
			s.mu.RUnlock()
			return NewAPIError(http.StatusNotFound, "post_not_found")
		}
		authorID := p.AuthorID
		deletions := 0
		if u := s.users[userID]; u != nil {
			_, deletions = postCountsLocked(u, p.window.LocalDate)
		}
		s.mu.RUnlock()
		if authorID != userID {
			return NewAPIError(http.StatusForbidden, "post_not_editable")
		}
		if deletions >= 2 {
			return NewAPIError(http.StatusConflict, "post_slot_unavailable")
		}
		deletedAt := now()
		s.mu.Lock()
		rw := s.commitLocked(&opPostDelete{ID: p.ID, At: deletedAt})
		s.mu.Unlock()
		rw.wait()
		imagesInvalidated(p.ImageID)
		return c.NoContent(http.StatusNoContent)
	}
}
