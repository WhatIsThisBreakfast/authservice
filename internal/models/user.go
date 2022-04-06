package models

type User struct {
	ID       int    `json:"id"`
	PublicID string `json:"public_id"`
	Payload  string
}
