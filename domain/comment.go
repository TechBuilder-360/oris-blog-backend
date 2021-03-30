package domain

// Comment struct
type Comment struct {
	PostID string `json:"postid,omitempty"`
	CommentID string `json:"commentid,omitempty"`
	Author string `json:"author,omitempty"`
	Text string `json:"text,omitempty"`
	DateCreated string `json:"datecreated,omitempty"`
	Replies []Comment `json:"replies,omitempty"`
	RepliesCount int
	Likes [] string `json:"likes,omitempty"` //userID array
	LikeCount int
}