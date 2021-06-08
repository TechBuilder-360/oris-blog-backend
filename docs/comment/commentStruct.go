package docs

import "time"

type CreateCommentRequest struct {
	PostID   string `json:"postid"`
	AuthorID string `json:"authorid"`
	Text     string `json:"text"`
}

type StatusResponse struct {
	Status string `json:"status"`
}

type Comment struct {
	_id          string
	PostID       string
	Author       string
	Text         string
	DateCreated  time.Time
	Replies      []string //userID array
	RepliesCount int
	Likes        []string //userID array
	LikeCount    int
}