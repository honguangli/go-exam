package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

// 考试计划班级表
type PlanClassRelModel struct {
	ID      int `orm:"column(id)" form:"id" json:"id"`
	PlanID  int `orm:"column(plan_id)" form:"plan_id" json:"plan_id"`
	ClassID int `orm:"column(class_id)" form:"class_id" json:"class_id"`

	// 班级信息
	ClassName string `orm:"column(class_name)" form:"class_name" json:"class_name"`
}

// 查询列表参数
type ReadPlanClassRelModelListParam struct {
	BaseQueryParam
	PlanID    int  `json:"plan_id"`
	ClosePage bool `form:"close_page" json:"close_page"`
}

// 查询多个对象
func ReadPlanClassRelModelListRaw(param ReadPlanClassRelModelListParam) (list []*PlanClassRelModel, total int64, err error) {
	list = make([]*PlanClassRelModel, 0)
	var args = make([]interface{}, 0)
	var whereSql = "WHERE 1=1"

	if param.PlanID > 0 {
		whereSql += " AND T0.plan_id = ?"
		args = append(args, param.PlanID)
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
	var fields = "T0.`id`, T0.`plan_id`, T0.`class_id`"

	// 关联查询
	var relatedSql = "LEFT JOIN `class` AS T1 ON T1.id = T0.class_id"
	fields += ", T1.name AS class_name"

	// 连表查询
	var sql = fmt.Sprintf("SELECT %s FROM plan_class_rel AS T0 %s %s %s %s", fields, relatedSql, whereSql, orderSql, pageSql)

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
	var countSql = fmt.Sprintf("SELECT count(*) AS count FROM plan_class_rel AS T0 %s %s", relatedSql, whereSql)
	var count RawCount
	err = orm.NewOrm().Raw(countSql, args...).QueryRow(&count)
	if err != nil {
		return
	}
	total = count.Count
	return
}
