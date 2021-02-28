package models

// Comment struct
type Comment struct {
	ParentID string `json:"parentid,omitempty"`
	CommentID string `json:"commentid,omitempty"`
	Author string `json:"author,omitempty"`
	Text string `json:"text,omitempty"`
	DateCreated string `json:"datecreated,omitempty"`
	Replies []Comment `json:"replies,omitempty"`
	Likes int `json:"likes,omitempty"`
}