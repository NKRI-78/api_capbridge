package entities

import "time"

type AdminListUser struct {
	Id        string    `json:"id"`
	Avatar    string    `json:"avatar"`
	Fullname  string    `json:"fullname"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Role      string    `json:"role"`
	Verified  bool      `json:"verified"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type AdminListUserResponse struct {
	Id        string    `json:"id"`
	Avatar    string    `json:"avatar"`
	Fullname  string    `json:"fullname"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Role      string    `json:"role"`
	Verified  bool      `json:"verified"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type AdminVerifyUser struct {
	UserId string `json:"user_id"`
}

type AdminVerifyProject struct {
	Id string `json:"id"`
}
