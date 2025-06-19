package entities

import "time"

type PPOBTransactionListScan struct {
	AppId     int       `json:"app_id"`
	AppName   string    `json:"app_name"`
	Value     string    `json:"value"`
	Idpel     string    `json:"idpel"`
	Product   string    `json:"product"`
	CreatedAt time.Time `json:"created_at"`
}

type PPOBTransactionListResponse struct {
	App       App       `json:"app"`
	Value     string    `json:"value"`
	Idpel     string    `json:"idpel"`
	Product   string    `json:"product"`
	CreatedAt time.Time `json:"created_at"`
}

type App struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
