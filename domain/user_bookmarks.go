package domain

// UserBookmark struct
type UserBookmark struct {
	UserID string `json:"userid,omitempty"`
	Bookmarks [] string `json:"bookmarks,omitempty"`
}