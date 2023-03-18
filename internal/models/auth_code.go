package models

type AuthCode int

const (
	PodViewLog AuthCode = 400002
	PodShell   AuthCode = 400003
	PodDelete  AuthCode = 400004
)

type Kind string

type KindPair struct {
	Key     Kind   `json:"key"`
	ValueCN string `json:"valueCN"`
	ValueEN string `json:"valueEN"`
}

const (
	Root  Kind = "root"
	Dev   Kind = "dev"
	Ops   Kind = "ops"
	Guest Kind = "guest"
)
