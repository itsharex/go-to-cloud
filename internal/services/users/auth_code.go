package users

import (
	"go-to-cloud/internal/models"
)

// GetAuthCodes 获取用户权限点
func GetAuthCodes(userId uint) []models.AuthCode {

	return []models.AuthCode{
		models.PodViewLog,
		models.PodDelete,
	}
}
