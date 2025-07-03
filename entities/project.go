package entities

import "time"

type ProjectListScan struct {
	Id                       string    `json:"id"`
	Title                    string    `json:"title"`
	Goal                     string    `json:"goal"`
	Capital                  string    `json:"capital"`
	Roi                      string    `json:"roi"`
	MinInvest                string    `json:"min_invest"`
	UnitPrice                string    `json:"unit_price"`
	UnitTotal                string    `json:"unit_total"`
	NumberOfUnit             string    `json:"number_of_unit"`
	Periode                  string    `json:"periode"`
	TypeOfBond               string    `json:"type_of_bond"`
	NominalValue             string    `json:"nominal_value"`
	TimePeriode              string    `json:"time_periode"`
	InterestRate             string    `json:"interest_rate"`
	UseOfFunds               string    `json:"use_of_funds"`
	InterestPaymentSchedule  string    `json:"interest_payment_schedule"`
	PrincipalPaymentSchedule string    `json:"principal_payment_schedule"`
	CollateralGuarantee      string    `json:"collateral_guarantee"`
	DescJob                  string    `json:"desc_job"`
	IsApbn                   bool      `json:"is_apbn"`
	IsApproved               bool      `json:"is_approved"`
	UserId                   string    `json:"user_id"`
	CreatedAt                time.Time `json:"created_at"`
	UpdatedAt                time.Time `json:"updated_at"`
}

type ProjectListResponse struct {
	Id                       string          `json:"id"`
	Title                    string          `json:"title"`
	Goal                     string          `json:"goal"`
	Medias                   []ProjectMedia  `json:"medias"`
	Location                 ProjectLocation `json:"location"`
	Doc                      ProjectDoc      `json:"doc"`
	Capital                  string          `json:"capital"`
	Roi                      string          `json:"roi"`
	MinInvest                string          `json:"min_invest"`
	UnitPrice                string          `json:"unit_price"`
	UnitTotal                string          `json:"unit_total"`
	NumberOfUnit             string          `json:"number_of_unit"`
	Periode                  string          `json:"periode"`
	TypeOfBond               string          `json:"type_of_bond"`
	NominalValue             string          `json:"nominal_value"`
	TimePeriode              string          `json:"time_periode"`
	InterestRate             string          `json:"interest_rate"`
	InterestPaymentSchedule  string          `json:"interest_payment_schedule"`
	PrincipalPaymentSchedule string          `json:"principal_payment_schedule"`
	UseOfFunds               string          `json:"use_of_funds"`
	CollateralGuarantee      string          `json:"collateral_guarantee"`
	DescJob                  string          `json:"desc_job"`
	IsApbn                   bool            `json:"is_apbn"`
	IsApproved               bool            `json:"is_approved"`
	Company                  Company         `json:"company"`
	CreatedAt                time.Time       `json:"created_at"`
	UpdatedAt                time.Time       `json:"updated_at"`
}

type Company struct {
	Name string `json:"name"`
}

type ProjectStore struct {
	Id           string          `json:"id"`
	Title        string          `json:"title"`
	Goal         string          `json:"goal"`
	Capital      string          `json:"capital"`
	Roi          string          `json:"roi"`
	MinInvest    string          `json:"min_invest"`
	UnitPrice    string          `json:"unit_price"`
	UnitTotal    string          `json:"unit_total"`
	NumberOfUnit string          `json:"number_of_unit"`
	Periode      string          `json:"periode"`
	Medias       []ProjectMedia  `json:"medias"`
	Location     ProjectLocation `json:"location"`
}

type ProjectMedia struct {
	Id   int    `json:"id"`
	Path string `json:"path"`
}

type ProjectStoreMedia struct {
	Id        string `json:"id"`
	Path      string `json:"path"`
	ProjectId string `json:"project_id"`
}

type ProjectLocation struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Url  string `json:"url"`
	Lat  string `json:"lat"`
	Lng  string `json:"lng"`
}

type ProjectStoreLocation struct {
	Id        int    `json:"id"`
	ProjectId string `json:"project_id"`
	Name      string `json:"name"`
	Url       string `json:"url"`
	Lat       string `json:"lat"`
	Lng       string `json:"lng"`
}

type ProjectDoc struct {
	Id   string `json:"id"`
	Path string `json:"path"`
}

type ProjectCompany struct {
	CompanyName string `json:"company_name"`
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
