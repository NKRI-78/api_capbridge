package entities

// ProfileScan hanya untuk hasil scan dari `profiles`
type ProfileScan struct {
	Id                string `json:"id"`
	Fullname          string `json:"fullname"`
	Avatar            string `json:"avatar"`
	Gender            string `json:"gender"`
	LastEdu           string `json:"last_edu"`
	StatusMarital     string `json:"status_marital"`
	Nik               string `json:"nik"`
	PlaceAndDatebirth string `json:"place_and_datebirth"`
	No                string `json:"no"`
	BankName          string `json:"bank_name"`
	BankOwner         string `json:"bank_owner"`
	BankBranch        string `json:"bank_branch"`
	CompanyName       string `json:"company_name"`
	CompanyAddress    string `json:"company_address"`
	MonthlyIncome     string `json:"monthly_income"`
	Position          string `json:"position"`
	CreatedAt         string `json:"created_at"`
	UpdatedAt         string `json:"updated_at"`
}

// Hasil scan dari `companies`
type ProfileEmiten struct {
	CompanyName          string `json:"company_name"`
	CompanyData          string `json:"company_data"`
	CompanyNib           string `json:"company_nib"`
	CompanyAddress       string `json:"company_address"`
	CompanyNpwp          string `json:"company_npwp"`
	CapitalStructure     string `json:"capital_structure"`
	FinancialStatements  string `json:"financial_statements"`
	CommissionerName     string `json:"commissioner_name"`
	CommissionerPosition string `json:"commissioner_position"`
	CommissionerKtp      string `json:"commissioner_ktp"`
	CommissionerNpwp     string `json:"commissioner_npwp"`
	DirectorName         string `json:"director_name"`
	DirectorPosition     string `json:"director_position"`
	DirectorKtp          string `json:"director_ktp"`
	DirectorNpwp         string `json:"director_npwp"`
	TotalEmployees       string `json:"total_employees"`
	DeedOfIncorporation  string `json:"deed_of_incorporation"`
	LatestAmendmentDeed  string `json:"latest_amendment_deed"`
	SkKemenkumham        string `json:"sk_kemenkumham"`
}

type ProfileInvestor struct {
	Gender        string      `json:"gender"`
	LastEdu       string      `json:"last_edu"`
	StatusMarital string      `json:"status_marital"`
	Ktp           string      `json:"ktp"`
	AddressKtp    string      `json:"address_ktp"`
	Bank          ProfileBank `json:"bank"`
	Job           ProfileJob  `json:"job"`
}

type ProfileBank struct {
	No     string `json:"no"`
	Name   string `json:"name"`
	Owner  string `json:"owner"`
	Branch string `json:"branch"`
}

type ProfileJob struct {
	CompanyName    string `json:"company_name"`
	CompanyAddress string `json:"company_address"`
	MonthlyIncome  string `json:"monthly_income"`
	Position       string `json:"position"`
}

// Struct response akhir
type ProfileResponse struct {
	Profile
	Investor ProfileInvestor `json:"investor"`
	Emiten   ProfileEmiten   `json:"emiten"`
}

type Profile struct {
	Id        string `json:"id"`
	Fullname  string `json:"fullname"`
	Avatar    string `json:"avatar"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type GetProfile struct {
	UserId string `json:"user_id"`
}

type UpdateProfile struct {
	UserId   string `json:"user_id"`
	Fullname string `json:"fullname"`
}
