package model

// City the city model
type City struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	State       string `json:"state"`
	Description string `json:"description"`
}
