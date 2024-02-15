package entities

import "time"

type User struct {
	ID        uint   `json:"id"`
	RoleID    uint   `json:"role_id" gorm:"column:role_id"`
	Role      *Role   `gorm:"foreignKey:RoleID;references:ID"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	NoTelp    string `json:"no_telp"`
	Photo     string `json:"photo"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (e *User) TableName() string {
	return "users"
}
