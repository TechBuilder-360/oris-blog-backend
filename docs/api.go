package docs

// CreatePostRequest represents body of CreatePost request.
type CreatePostRequest struct{
	Author string `json:"author,omitempty"`
	Title string `json:"title,omitempty"`
	Categories []string `json:"categories,omitempty"`
	Article string `json:"article,omitempty"`
	Status string `json:"status,omitempty"`
}

// CreatePostResponse represents body of CreatePost response.
type CreatePostResponse struct{
	PostID string `json:"postid,omitempty"`
}