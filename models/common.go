package models

import "go-exam/enums"

type JsonResult struct {
	Code enums.ResponseCode `json:"code"`
	Msg  string             `json:"msg"`
	Obj  interface{}        `json:"obj"`
}

type BaseQueryParam struct {
	Sort   string `form:"sort" json:"sort"`
	Order  string `form:"order" json:"order"`
	Offset int64  `form:"offset" json:"offset"`
	Limit  int    `form:"limit" json:"limit"`
}

type RawCount struct {
	Count int64 `json:"count"`
}
