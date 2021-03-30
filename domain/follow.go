package domain

// Follow struct
type Follow struct {
	UserID string `json:"userid,omitempty"`
	Followers [] string `json:"followers,omitempty"`
}