package model

type UserModel struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Guid     string `json:"guid"`
}
