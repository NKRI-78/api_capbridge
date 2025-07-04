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

type AdminListProject struct {
	Id                       string `json:"id"`
	Title                    string `json:"title"`
	Goal                     string `json:"goal"`
	Capital                  string `json:"capital"`
	Roi                      string `json:"roi"`
	MinInvest                string `json:"min_invest"`
	UnitPrice                string `json:"unit_price"`
	UnitTotal                string `json:"unit_total"`
	NumberOfUnit             string `json:"number_of_init"`
	Periode                  string `json:"periode"`
	TypeOfBond               string `json:"type_of_bond"`
	NominalValue             string `json:"nominal_value"`
	TimePeriode              string `json:"time_periode"`
	InterestRate             string `json:"interest_rate"`
	InterestPaymentSchedule  string `json:"interest_payment_schedule"`
	PrincipalPaymentSchedule string `json:"principal_payment_schedule"`
	UseOfFunds               string `json:"use_of_funds"`
	CollateralGuarantee      string `json:"collateral_guarantee"`
	DescJob                  string `json:"desc_job"`
	IsApbn                   bool   `json:"is_apbn"`
	IsApproved               bool   `json:"is_approved"`
	UserId                   string `json:"user_id"`
	UserEmail                string `json:"user_email"`
	UserName                 string `json:"user_name"`
	UserPhone                string `json:"user_phone"`
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

type AdminListProjectResponse struct {
	Id                       string               `json:"id"`
	Title                    string               `json:"title"`
	Goal                     string               `json:"goal"`
	Capital                  string               `json:"capital"`
	Roi                      string               `json:"roi"`
	MinInvest                string               `json:"min_invest"`
	UnitPrice                string               `json:"unit_price"`
	UnitTotal                string               `json:"unit_total"`
	NumberOfUnit             string               `json:"number_of_init"`
	Periode                  string               `json:"periode"`
	TypeOfBond               string               `json:"type_of_bond"`
	NominalValue             string               `json:"nominal_value"`
	TimePeriode              string               `json:"time_periode"`
	InterestRate             string               `json:"interest_rate"`
	InterestPaymentSchedule  string               `json:"interest_payment_schedule"`
	PrincipalPaymentSchedule string               `json:"principal_payment_schedule"`
	UseOfFunds               string               `json:"use_of_funds"`
	CollateralGuarantee      string               `json:"collateral_guarantee"`
	DescJob                  string               `json:"desc_job"`
	IsApbn                   bool                 `json:"is_apbn"`
	IsApproved               bool                 `json:"is_approved"`
	Media                    []AdminListMedia       `json:"media"`
	Location                 AdminListLocation    `json:"location"`
	User                     AdminListProjectUser `json:"user"`
}

type AdminListLocation struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Url  string `json:"url"`
	Lat  string `json:"lat"`
	Lng  string `json:"lng"`
}

type AdminListMedia struct {
	Id   int    `json:"id"`
	Path string `json:"path"`
}

type AdminListProjectUser struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type AdminVerifyUser struct {
	UserId string `json:"user_id"`
}

type AdminVerifyProject struct {
	Id string `json:"id"`
}
