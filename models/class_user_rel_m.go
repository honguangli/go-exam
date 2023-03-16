package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

// 班级考生表
type ClassUserRelModel struct {
	ID      int `orm:"column(id)" form:"id" json:"id"`
	ClassID int `orm:"column(class_id)" form:"class_id" json:"class_id"`
	UserID  int `orm:"column(user_id)" form:"user_id" json:"user_id"`

	// 用户信息
	UserName     string `orm:"column(user_name)" form:"user_name" json:"user_name"`
	UserTrueName string `orm:"column(user_true_name)" form:"user_true_name" json:"user_true_name"`
}

// 查询列表参数
type ReadClassUserRelModelListParam struct {
	BaseQueryParam
	ClassID   int  `json:"class_id"`
	ClosePage bool `form:"close_page" json:"close_page"`
}

// 查询多个对象
func ReadClassUserRelModelListRaw(param ReadClassUserRelModelListParam) (list []*ClassUserRelModel, total int64, err error) {
	list = make([]*ClassUserRelModel, 0)
	var args = make([]interface{}, 0)
	var whereSql = "WHERE 1=1"

	if param.ClassID > 0 {
		whereSql += " AND T0.class_id = ?"
		args = append(args, param.ClassID)
	}

	// 排序
	var orderSql = "ORDER BY "
	switch param.Sort {
	default:
		orderSql += "T0.id"
	}
	if param.Order == "desc" {
		orderSql += " DESC"
	}

	// 分页
	var pageSql string
	if !param.ClosePage {
		pageSql = fmt.Sprintf("LIMIT %d OFFSET %d", param.Limit, param.Offset)
	}

	// 查询字段
	var fields = "T0.`id`, T0.`class_id`, T0.`user_id`"

	// 关联查询
	var relatedSql = "LEFT JOIN `user` AS T1 ON T1.id = T0.user_id"
	fields += ", T1.name AS user_name, T1.true_name AS user_true_name"

	// 连表查询
	var sql = fmt.Sprintf("SELECT %s FROM class_user_rel AS T0 %s %s %s %s", fields, relatedSql, whereSql, orderSql, pageSql)

	// 查询列表
	total, err = orm.NewOrm().Raw(sql, args...).QueryRows(&list)
	if err != nil {
		return
	}

	// 关闭分页时不查询count
	if param.ClosePage {
		return
	}

	// 查询总数
	var countSql = fmt.Sprintf("SELECT count(*) AS count FROM class_user_rel AS T0 %s %s", relatedSql, whereSql)
	var count RawCount
	err = orm.NewOrm().Raw(countSql, args...).QueryRow(&count)
	if err != nil {
		return
	}
	total = count.Count
	return
}
