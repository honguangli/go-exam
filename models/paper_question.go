package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

// 试卷试题表
type PaperQuestion struct {
	ID int `orm:"column(id)" form:"id" json:"id"`
}

// 查询列表参数
type ReadPaperQuestionListParam struct {
	BaseQueryParam
	ClosePage bool `form:"close_page" json:"close_page"`
}

// 初始化
func init() {
	orm.RegisterModel(new(PaperQuestion))
}

// 表名
func PaperQuestionTBName() (name string) {
	return "paper_question"
}

// 自定义表名
func (m *PaperQuestion) TableName() (name string) {
	return PaperQuestionTBName()
}

// 多字段索引
func (m *PaperQuestion) TableIndex() [][]string {
	return [][]string{}
}

// 多字段唯一键
func (m *PaperQuestion) TableUnique() [][]string {
	return [][]string{}
}

// 自定义引擎
func (m *PaperQuestion) TableEngine() string {
	return "INNODB"
}

// 查询单个对象
func ReadPaperQuestionOne(id int) (m PaperQuestion, err error) {
	o := orm.NewOrm()
	m.ID = id
	err = o.Read(&m)
	return
}

// 查询多个对象
func ReadPaperQuestionList(param ReadPaperQuestionListParam) (list []*PaperQuestion, total int64, err error) {
	list = make([]*PaperQuestion, 0)
	query := orm.NewOrm().QueryTable(PaperQuestionTBName())

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
func ReadPaperQuestionListRaw(param ReadPaperQuestionListParam) (list []*PaperQuestion, total int64, err error) {
	list = make([]*PaperQuestion, 0)
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
	var fields = "T0.`id`"

	// 关联查询
	var relatedSql string

	// 连表查询
	var sql = fmt.Sprintf("SELECT %s FROM paper_question AS T0 %s %s %s %s", fields, relatedSql, whereSql, orderSql, pageSql)

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
	var countSql = fmt.Sprintf("SELECT count(*) AS count FROM paper_question AS T0 %s %s", relatedSql, whereSql)
	var count RawCount
	err = orm.NewOrm().Raw(countSql, args...).QueryRow(&count)
	if err != nil {
		return
	}
	total = count.Count
	return
}

// 插入单个对象
func InsertPaperQuestionOne(m PaperQuestion) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(&m)
	return
}

// 插入多个对象
func InsertPaperQuestionMulti(list []PaperQuestion) (num int64, err error) {
	o := orm.NewOrm()
	num, err = o.InsertMulti(100, list)
	return
}

// 更新单个对象
func UpdatePaperQuestionOne(m PaperQuestion, fields ...string) (num int64, err error) {
	o := orm.NewOrm()
	if len(fields) == 0 {
		fields = []string{""}
	}
	num, err = o.Update(&m, fields...)
	return
}

// 更新多个对象
func UpdatePaperQuestionMulti(ids []int, params map[string]interface{}) (num int64, err error) {
	o := orm.NewOrm()
	num, err = o.QueryTable(PaperQuestionTBName()).Filter("id__in", ids).Update(params)
	return
}

// 删除单个对象
func DeletePaperQuestionOne(id int) (num int64, err error) {
	o := orm.NewOrm()
	num, err = o.Delete(&PaperQuestion{ID: id})
	return
}

// 删除多个对象
func DeletePaperQuestionMulti(ids []int) (num int64, err error) {
	o := orm.NewOrm()
	num, err = o.QueryTable(PaperQuestionTBName()).Filter("id__in", ids).Delete()
	return
}
