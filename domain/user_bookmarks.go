package domain

// UserBookmark struct
type UserBookmark struct {
	UserID string `json:"userid"`
	Bookmarks [] string `json:"bookmarks"`
}