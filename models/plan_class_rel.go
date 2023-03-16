package models

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/orm"
)

// 考试计划班级表
type PlanClassRel struct {
	ID      int `orm:"column(id)" form:"id" json:"id"`
	PlanID  int `orm:"column(plan_id)" form:"plan_id" json:"plan_id"`
	ClassID int `orm:"column(class_id)" form:"class_id" json:"class_id"`
}

// 查询详情参数
type ReadPlanClassRelDetailParam struct {
	ID int `json:"id"`
}

// 查询列表参数
type ReadPlanClassRelListParam struct {
	BaseQueryParam
	ClosePage bool `form:"close_page" json:"close_page"`
}

// 创建参数
type InsertPlanClassRelParam struct {
	PlanID    int   `json:"plan_id"`
	ClassList []int `json:"class_list"`
}

// 删除参数
type DeletePlanClassRelParam struct {
	ID   int   `json:"id"`
	List []int `json:"list"`
}

// 初始化
func init() {
	orm.RegisterModel(new(PlanClassRel))
}

// 表名
func PlanClassRelTBName() (name string) {
	return "plan_class_rel"
}

// 自定义表名
func (m *PlanClassRel) TableName() (name string) {
	return PlanClassRelTBName()
}

// 多字段索引
func (m *PlanClassRel) TableIndex() [][]string {
	return [][]string{}
}

// 多字段唯一键
func (m *PlanClassRel) TableUnique() [][]string {
	return [][]string{
		{"plan_id", "class_id"},
	}
}

// 自定义引擎
func (m *PlanClassRel) TableEngine() string {
	return "INNODB"
}

// 查询单个对象
func ReadPlanClassRelOne(id int) (m PlanClassRel, err error) {
	o := orm.NewOrm()
	m.ID = id
	err = o.Read(&m)
	return
}

// 查询多个对象
func ReadPlanClassRelList(param ReadPlanClassRelListParam) (list []*PlanClassRel, total int64, err error) {
	list = make([]*PlanClassRel, 0)
	query := orm.NewOrm().QueryTable(PlanClassRelTBName())

	sortOrder := "id"
	switch param.Sort {
	}
	if param.Order == "desc" {
		sortOrder = "-" + sortOrder
	}

	total, err = query.Count()
	if err != nil {
		return
	}
	_, err = query.OrderBy(sortOrder).Limit(param.Limit, param.Offset).All(&list)
	return
}

// 查询多个对象
func ReadPlanClassRelListRaw(param ReadPlanClassRelListParam) (list []*PlanClassRel, total int64, err error) {
	list = make([]*PlanClassRel, 0)
	var args = make([]interface{}, 0)
	var whereSql = "WHERE 1=1"

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
	var relatedSql string

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

// 插入单个对象
func InsertPlanClassRelOne(m PlanClassRel) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(&m)
	return
}

// 插入多个对象
func InsertPlanClassRelMulti(list []PlanClassRel) (num int64, err error) {
	o := orm.NewOrm()
	num, err = o.InsertMulti(100, list)
	return
}

// 更新单个对象
func UpdatePlanClassRelOne(m PlanClassRel, fields ...string) (num int64, err error) {
	o := orm.NewOrm()
	if len(fields) == 0 {
		fields = []string{"plan_id", "class_id"}
	}
	num, err = o.Update(&m, fields...)
	return
}

// 更新多个对象
func UpdatePlanClassRelMulti(ids []int, params map[string]interface{}) (num int64, err error) {
	o := orm.NewOrm()
	num, err = o.QueryTable(PlanClassRelTBName()).Filter("id__in", ids).Update(params)
	return
}

// 删除单个对象
func DeletePlanClassRelOne(id int) (num int64, err error) {
	o := orm.NewOrm()
	num, err = o.Delete(&PlanClassRel{ID: id})
	return
}

// 删除多个对象
func DeletePlanClassRelMulti(ids []int) (num int64, err error) {
	o := orm.NewOrm()
	num, err = o.QueryTable(PlanClassRelTBName()).Filter("id__in", ids).Delete()
	return
}

// 插入或更新多个对象
func InsertOrUpdatePlanClassRelMulti(param InsertPlanClassRelParam) (num int64, err error) {
	var m = make(map[int]struct{})
	for _, v := range param.ClassList {
		m[v] = struct{}{}
	}
	var list = make([]PlanClassRel, 0)
	for k := range m {
		list = append(list, PlanClassRel{PlanID: param.PlanID, ClassID: k})
	}
	if len(list) == 0 {
		return 0, errors.New("参数错误: 班级列表不能为空")
	}
	num = int64(len(list))

	o := orm.NewOrm()
	err = o.Begin()
	if err != nil {
		return
	}

	for _, v := range list {
		_, err = o.InsertOrUpdate(&v)
		if err != nil {
			o.Rollback()
			return
		}
	}

	err = o.Commit()
	return
}
