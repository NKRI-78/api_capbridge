package entities

import "time"

type ProfileScan struct {
	Id        string    `json:"id"`
	Fullname  string    `json:"fullname"`
	Avatar    string    `json:"avatar"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type GetProfile struct {
	UserId string `json:"user_id"`
}

type UpdateProfile struct {
	UserId   string `json:"user_id"`
	Fullname string `json:"fullname"`
}
