package entities

import "time"

type AdminListUser struct {
	Id        string    `json:"id"`
	Avatar    string    `json:"avatar"`
	Fullname  string    `json:"fullname"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
