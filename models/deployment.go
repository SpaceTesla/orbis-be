package models

type Deployment struct {
	ID      int    `json:"id"`
	Project string `json:"project"`
	Status  string `json:"status"`
}
