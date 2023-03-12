package user

import "go-to-cloud/internal/utils"

type User struct {
	Key       uint           `json:"key"` // 适配前端默认关键词，与Id同值
	Id        uint           `json:"id"`
	CreatedAt utils.JsonTime `json:"created_at"`
	Account   string         `json:"account"`  // 账号
	Name      string         `json:"name"`     // 真实名称
	Shortcut  string         `json:"shortcut"` // 名称快捷方式，默认拼音首字母
}
