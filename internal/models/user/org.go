package user

import "go-to-cloud/internal/utils"

type Org struct {
	Key         uint           `json:"key"`
	Id          uint           `json:"id"`
	CreatedAt   utils.JsonTime `json:"-"`
	Name        string         `json:"name"` // 组织名称
	MemberCount uint           `json:"member_count"`
	Remark      string         `json:"remark"`
}
