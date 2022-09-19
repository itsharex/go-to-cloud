package models

type LoginModel struct {
	Account      string `json:"username"`
	OrigPassword string `json:"password"`
}
