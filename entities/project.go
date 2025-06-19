package entities

import "time"

type ProjectListScan struct {
	Id           string    `json:"id"`
	Title        string    `json:"title"`
	Goal         string    `json:"goal"`
	Capital      string    `json:"capital"`
	Roi          string    `json:"roi"`
	MinInvest    string    `json:"min_invest"`
	UnitPrice    string    `json:"unit_price"`
	UnitTotal    string    `json:"unit_total"`
	NumberOfUnit string    `json:"number_of_unit"`
	Periode      string    `json:"periode"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type ProjectListResponse struct {
	Id           string          `json:"id"`
	Title        string          `json:"title"`
	Goal         string          `json:"goal"`
	Medias       []ProjectMedia  `json:"medias"`
	Location     ProjectLocation `json:"location"`
	Doc          ProjectDoc      `json:"doc"`
	Capital      string          `json:"capital"`
	Roi          string          `json:"roi"`
	MinInvest    string          `json:"min_invest"`
	UnitPrice    string          `json:"unit_price"`
	UnitTotal    string          `json:"unit_total"`
	NumberOfUnit string          `json:"number_of_unit"`
	Periode      string          `json:"periode"`
	CreatedAt    time.Time       `json:"created_at"`
	UpdatedAt    time.Time       `json:"updated_at"`
}

type ProjectMedia struct {
	Id   string `json:"id"`
	Path string `json:"path"`
}

type ProjectLocation struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Url  string `json:"url"`
	Lat  string `json:"lat"`
	Lng  string `json:"lng"`
}

type ProjectDoc struct {
	Id   string `json:"id"`
	Path string `json:"path"`
}

type ProjectStore struct {
	Id        string `json:"id"`
	Title     string `json:"title"`
	Goal      string `json:"goal"`
	Capital   string `json:"capital"`
	Roi       string `json:"roi"`
	MinInvest string `json:"min_invest"`
	UnitPrice string `json:"unit_price"`
	UnitTotal string `json:"unit_total"`
}

type ProjectUpdate struct {
	Id        string `json:"id"`
	Title     string `json:"title"`
	Goal      string `json:"goal"`
	Capital   string `json:"capital"`
	Roi       string `json:"roi"`
	MinInvest string `json:"min_invest"`
	UnitPrice string `json:"unit_price"`
	UnitTotal string `json:"unit_total"`
}

type ProjectDelete struct {
	Id string `json:"id"`
}
