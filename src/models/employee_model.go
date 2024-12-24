package models

type Employee struct {
	ID        int      `json:"id" sql:"id"`
	Position  string   `json:"position"  sql:"position"`
	ProfileID int      `json:"profile_id"  sql:"profile_id"`
	Profile   *Profile `json:"profile" sql:""`
}
