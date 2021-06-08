package docs

type Bookmark struct {
	UserID string `json:"userid"`
	Bookmarks [] string `json:"bookmarks"`
}

type StatusResponse struct {
	Status string `json:"status"`
}