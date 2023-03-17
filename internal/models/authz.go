package models

const RootUserName = "root"

type AuthCode int

const (
	PodViewLog AuthCode = 400002
	PodShell   AuthCode = 400003
	PodDelete  AuthCode = 400004
)
