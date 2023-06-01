package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

// 考生试题答案表
type AnswerItem struct {
	ID         int    `orm:"column(id)" form:"id" json:"id"`
	AnswerID   int    `orm:"column(answer_id)" form:"answer_id" json:"answer_id"`
	QuestionID int    `orm:"column(question_id)" form:"question_id" json:"question_id"`
	OptionIds  string `orm:"column(option_ids)" form:"option_ids" json:"option_ids"`
	Content    string `orm:"column(content)" form:"content" json:"content"`
	Check      int8   `orm:"column(check)" form:"check" json:"check"`
	Score      int    `orm:"column(score)" form:"score" json:"score"`
	Memo       string `orm:"column(memo)" form:"memo" json:"memo"`
}

// 状态
const (
	AnswerItemUnCheck = 0 // 未评分
	AnswerItemChecked = 1 // 已评分
)

// 查询列表参数
type ReadAnswerItemListParam struct {
	BaseQueryParam
	ClosePage bool `form:"close_page" json:"close_page"`
}

// 初始化
func init() {
	orm.RegisterModel(new(AnswerItem))
}

// 表名
func AnswerItemTBName() (name string) {
	return "answer_item"
}

// 自定义表名
func (m *AnswerItem) TableName() (name string) {
	return AnswerItemTBName()
}

// 多字段索引
func (m *AnswerItem) TableIndex() [][]string {
	return [][]string{}
}

// 多字段唯一键
func (m *AnswerItem) TableUnique() [][]string {
	return [][]string{}
}

// 自定义引擎
func (m *AnswerItem) TableEngine() string {
	return "INNODB"
}

// 查询单个对象
func ReadAnswerItemOne(id int) (m AnswerItem, err error) {
	o := orm.NewOrm()
	m.ID = id
	err = o.Read(&m)
	return
}

// 查询多个对象
func ReadAnswerItemList(param ReadAnswerItemListParam) (list []*AnswerItem, total int64, err error) {
	list = make([]*AnswerItem, 0)
	query := orm.NewOrm().QueryTable(AnswerItemTBName())

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
func ReadAnswerItemListRaw(param ReadAnswerItemListParam) (list []*AnswerItem, total int64, err error) {
	list = make([]*AnswerItem, 0)
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
	var fields = "T0.`id`, T0.`answer_id`, T0.`question_id`, T0.`option_ids`, T0.`content`, T0.`check`, T0.`score`, T0.`memo`"

	// 关联查询
	var relatedSql string

	// 连表查询
	var sql = fmt.Sprintf("SELECT %s FROM answer_item AS T0 %s %s %s %s", fields, relatedSql, whereSql, orderSql, pageSql)

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
	var countSql = fmt.Sprintf("SELECT count(*) AS count FROM answer_item AS T0 %s %s", relatedSql, whereSql)
	var count RawCount
	err = orm.NewOrm().Raw(countSql, args...).QueryRow(&count)
	if err != nil {
		return
	}
	total = count.Count
	return
}

// 插入单个对象
func InsertAnswerItemOne(m AnswerItem) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(&m)
	return
}

// 插入多个对象
func InsertAnswerItemMulti(list []AnswerItem) (num int64, err error) {
	o := orm.NewOrm()
	num, err = o.InsertMulti(100, list)
	return
}

// 更新单个对象
func UpdateAnswerItemOne(m AnswerItem, fields ...string) (num int64, err error) {
	o := orm.NewOrm()
	if len(fields) == 0 {
		fields = []string{"answer_id", "question_id", "option_ids", "content", "check", "score", "memo"}
	}
	num, err = o.Update(&m, fields...)
	return
}

// 更新多个对象
func UpdateAnswerItemMulti(ids []int, params map[string]interface{}) (num int64, err error) {
	o := orm.NewOrm()
	num, err = o.QueryTable(AnswerItemTBName()).Filter("id__in", ids).Update(params)
	return
}

// 删除单个对象
func DeleteAnswerItemOne(id int) (num int64, err error) {
	o := orm.NewOrm()
	num, err = o.Delete(&AnswerItem{ID: id})
	return
}

// 删除多个对象
func DeleteAnswerItemMulti(ids []int) (num int64, err error) {
	o := orm.NewOrm()
	num, err = o.QueryTable(AnswerItemTBName()).Filter("id__in", ids).Delete()
	return
}
