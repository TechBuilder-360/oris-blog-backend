package domain

// UserInterest struct
type UserInterest struct {
	UserID string `json:"userid,omitempty"`
	Categories [] string `json:"categories,omitempty"`
}