package models

import "go-to-cloud/internal/utils"

type GitModel struct {
	Address      string `json:"address"`
	Branch       string `json:"branch"`
	EncodedToken string `json:"token"`
}

func (m *GitModel) EncodeToken(token *string) {
	m.EncodedToken = string(utils.AesEny([]byte(*token)))
}

func (m *GitModel) DecodeToken() *string {
	token := string(utils.AesEny([]byte(m.EncodedToken)))

	return &token
}
