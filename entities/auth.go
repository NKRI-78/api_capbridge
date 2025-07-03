package entities

type Register struct {
	Email    string   `json:"email"`
	Phone    string   `json:"phone"`
	Fullname string   `json:"fullname"`
	Role     string   `json:"role"`
	Password string   `json:"password"`
	UserId   string   `json:"user_id"`
	Investor Investor `json:"investor"`
	Emiten   Emiten   `json:"emiten"`
}

type Investor struct {
	Gender        string `json:"gender"`
	LastEdu       string `json:"last_edu"`
	StatusMarital string `json:"status_marital"`
	Ktp           string `json:"ktp"`
	AddressKtp    string `json:"address_ktp"`
	Bank          Bank   `json:"bank"`
	Job           Job    `json:"job"`
}

type Bank struct {
	No     string `json:"no"`
	Name   string `json:"name"`
	Owner  string `json:"owner"`
	Branch string `json:"branch"`
}

type Job struct {
	CompanyName    string `json:"company_name"`
	CompanyAddress string `json:"company_address"`
	MonthlyIncome  string `json:"monthly_income"`
	Position       string `json:"position"`
}

type Emiten struct {
	CompanyName         string   `json:"company_name"`
	CompanyData         string   `json:"company_data"`
	CompanyNib          string   `json:"company_nib"`
	DeedOfIncorporation string   `json:"deed_of_incorporation"`
	LatestAmendmentDeed string   `json:"latest_amendment_deed"`
	SkKemenkumham       string   `json:"sk_kemenkumham"`
	CompanyAddress      string   `json:"company_address"`
	CompanyNpwp         string   `json:"company_npwp"`
	TotalEmployees      string   `json:"total_employees"`
	CapitalStructure    string   `json:"capital_structure"`
	FinancialStatements string   `json:"financial_statements"`
	CommisionerName     string   `json:"commisioner_name"`
	CommisionerPosition string   `json:"commisioner_position"`
	CommisionerKtp      string   `json:"commisioner_ktp"`
	CommisionerNpwp     string   `json:"commisioner_npwp"`
	DirectorName        string   `json:"director_name"`
	DirectorPosition    string   `json:"director_position"`
	DirectorKtp         string   `json:"director_ktp"`
	DirectorNpwp        string   `json:"director_npwp"`
	InfoBond            InfoBond `json:"info_bond"`
}

type InfoBond struct {
	Title                    string   `json:"title"`
	Img                      string   `json:"img"`
	Doc                      string   `json:"doc"`
	Location                 Location `json:"location"`
	TypeOfBond               string   `json:"type_of_bond"`
	NominalValue             string   `json:"nominal_value"`
	TimePeriode              string   `json:"time_periode"`
	InterestRate             string   `json:"interest_rate"`
	InterestPaymentSchedule  string   `json:"interest_payment_schedule"`
	PrincipalPaymentSchedule string   `json:"principal_payment_schedule"`
	UseOfFunds               string   `json:"use_of_funds"`
	CollateralGuarantee      string   `json:"collateral_guarantee"`
	DescJob                  string   `json:"desc_job"`
	IsApbn                   bool     `json:"is_apbn"`
}

type Location struct {
	Name string `json:"name"`
	Url  string `json:"url"`
	Lat  string `json:"lat"`
	Lng  string `json:"lng"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginScan struct {
	Id       string `json:"id"`
	Enabled  bool   `json:"enabled"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	Verify   bool   `json:"verify"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Id      string `json:"id"`
	Enabled bool   `json:"enabled"`
	Email   string `json:"email"`
	Role    string `json:"role"`
	Verify  bool   `json:"verify"`
	Token   string `json:"token"`
}

type RegisterResponse struct {
	Id      string `json:"id"`
	Enabled bool   `json:"enabled"`
	Email   string `json:"email"`
	Role    string `json:"role"`
	Verify  bool   `json:"verify"`
	Token   string `json:"token"`
}

type CheckRole struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
