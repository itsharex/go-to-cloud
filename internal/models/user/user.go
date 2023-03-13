package user

import "go-to-cloud/internal/utils"

type User struct {
	Key            uint           `json:"key"` // 适配前端默认关键词，与Id同值
	Id             uint           `json:"id"`
	CreatedAt      utils.JsonTime `json:"created_at"`
	Account        string         `json:"account"`        // 账号
	RealName       string         `json:"name"`           // 真实名称
	Pinyin         string         `json:"pinyin"`         // 拼音
	PinyinInit     string         `json:"pinyinInit"`     // 拼音首字母
	OriginPassword string         `json:"originPassword"` // 原始密码（只能由前端传至后端，后端会忽略这个字段）
	Email          string         `json:"email"`          // 邮箱
	Mobile         string         `json:"mobile"`         // 电话，可用于接收验证码、钉钉被艾特
}

func (m *User) TransPinyin() (full, short string) {
	if len(m.RealName) > 0 {
		m.Pinyin, m.PinyinInit = utils.GetShortcut(m.RealName)
	} else if len(m.Account) > 0 {
		m.Pinyin, m.PinyinInit = utils.GetShortcut(m.RealName)
	}
	return m.Pinyin, m.PinyinInit
}