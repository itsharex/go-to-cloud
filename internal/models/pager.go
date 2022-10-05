package models

type Pager struct {
	PageIndex uint `form:"pageIndex"`
	PageSize  uint `form:"pageSize"`
}
