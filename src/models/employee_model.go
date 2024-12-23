package models

type Employee struct {
	ID        int    `json:"id"`
	Position  string `json:"position"`
	ProfileID int    `json:"profile_id"`
}
