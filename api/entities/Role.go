package entities

import "time"

type Role struct {
	ID        uint   `json:"id"`
	NameRole  string `json:"name_role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (e *Role) TableName() string {
	return "roles"
}