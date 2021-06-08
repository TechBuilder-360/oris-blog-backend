package docs

import "time"

type GetPostResponse struct {
	_id            string
	AuthorID       string `json:"authorID"`
	Title          string `json:"title"`
	Summary        string
	CoverImage     string `json:"coverimage"`
	Slug           string
	URL            string
	Categories     []string `json:"categories"`
	Likes          []string // user Id array
	Like_count     int
	Comments       []string // comment Id array
	Comments_count int
	Article        string `json:"article"`
	DateCreated    time.Time
	DateUpdated    time.Time
	Status         string `json:"status"`
	ReadTime       string
}

// CreatePostRequest represents body of CreatePost request.
type CreatePostRequest struct{
	Author string `json:"author"`
	Title string `json:"title"`
	Categories []string `json:"categories"`
	Article string `json:"article"`
	Status string `json:"status"`
}

// CreatePostResponse represents body of CreatePost response.
type CreatePostResponse struct{
	PostID string `json:"postid"`
}

type StatusReponse struct{
	Status string `json:"status"`
}