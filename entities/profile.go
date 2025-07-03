package entities

// ProfileScan hanya untuk hasil scan dari `profiles`
type ProfileScan struct {
	Id            string `json:"id"`
	Fullname      string `json:"fullname"`
	Avatar        string `json:"avatar"`
	Gender        string `json:"gender"`
	LastEdu       string `json:"last_edu"`
	StatusMarital string `json:"status_marital"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
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

// Struct response akhir
type ProfileResponse struct {
	ProfileScan
	Emiten ProfileEmiten `json:"emiten"`
}

type GetProfile struct {
	UserId string `json:"user_id"`
}

type UpdateProfile struct {
	UserId   string `json:"user_id"`
	Fullname string `json:"fullname"`
}
