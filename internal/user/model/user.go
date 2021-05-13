package model

type User struct {
	ID   ID      `json:"id"`
	Name *string `json:"Name,omitempty"`
}
