package docs

type StatusResponse struct{
	Status string `json:"status"`
}

type Follow struct {
	UserID string `json:"userid"`
	Followers [] string `json:"followers"`
	Following [] string `json:"following"`
}