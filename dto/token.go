package dto

type Token struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Level    string `json:"level"`
	Token    string `json:"auth_token"`
	IsAuth   bool   `json:"is_auth"`
}
