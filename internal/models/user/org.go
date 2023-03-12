package user

import "go-to-cloud/internal/utils"

type Org struct {
	Id          uint           `json:"id"`
	CreatedAt   utils.JsonTime `json:"created_at"`
	Name        string         `json:"name"` // 组织名称
	MemberCount uint           `json:"member_count"`
	Remark      string         `json:"remark"`
}
