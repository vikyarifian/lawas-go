package dto

type Token struct {
	UserID        string         `json:"user_id"`
	Username      string         `json:"username"`
	Name          string         `json:"name"`
	Email         string         `json:"email"`
	Level         string         `json:"level"`
	Token         string         `json:"auth_token"`
	IsAuth        bool           `json:"is_auth"`
	Notifications []Notification `json:"notificaion"`
}

type Notification struct {
	Code  string `json:"code"`
	Name  string `json:"name"`
	Count int    `json:"count"`
}
