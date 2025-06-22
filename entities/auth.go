package entities

type Register struct {
	UserId   string `json:"user_id"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Fullname string `json:"fullname"`
	Role     string `json:"role"`
	Password string `json:"password"`
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
