package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

// 试题选项表
type QuestionOption struct {
	ID         int    `orm:"column(id)" form:"id" json:"id"`
	QuestionID int    `orm:"column(question_id)" form:"question_id" json:"question_id"`
	Tag        string `orm:"column(tag)" form:"tag" json:"tag"`
	Content    string `orm:"column(content)" form:"content" json:"content"`
	IsRight    int8   `orm:"column(is_right)" form:"is_right" json:"is_right"`
	Memo       string `orm:"column(memo)" form:"memo" json:"memo"`
}

// 是否正确答案
const (
	QUESTION_OPTION_IS_NOT_RIGHT = iota // 错误
	QUESTION_OPTION_IS_RIGHT            // 正确
)

// 查询列表参数
type ReadQuestionOptionListParam struct {
	BaseQueryParam
	QuestionID int  `json:"question_id"`
	ClosePage  bool `form:"close_page" json:"close_page"`
}

// 初始化
func init() {
	orm.RegisterModel(new(QuestionOption))
}

// 表名
func QuestionOptionTBName() (name string) {
	return "question_option"
}

// 自定义表名
func (m *QuestionOption) TableName() (name string) {
	return QuestionOptionTBName()
}

// 多字段索引
func (m *QuestionOption) TableIndex() [][]string {
	return [][]string{}
}

// 多字段唯一键
func (m *QuestionOption) TableUnique() [][]string {
	return [][]string{}
}

// 自定义引擎
func (m *QuestionOption) TableEngine() string {
	return "INNODB"
}

// 查询单个对象
func ReadQuestionOptionOne(id int) (m QuestionOption, err error) {
	o := orm.NewOrm()
	m.ID = id
	err = o.Read(&m)
	return
}

// 查询多个对象
func ReadQuestionOptionList(param ReadQuestionOptionListParam) (list []*QuestionOption, total int64, err error) {
	list = make([]*QuestionOption, 0)
	query := orm.NewOrm().QueryTable(QuestionOptionTBName())

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
func ReadQuestionOptionListRaw(param ReadQuestionOptionListParam) (list []*QuestionOption, total int64, err error) {
	list = make([]*QuestionOption, 0)
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
	var fields = "T0.`id`, T0.`question_id`, T0.`tag`, T0.`content`, T0.`is_right`, T0.`memo`"

	// 关联查询
	var relatedSql string

	// 连表查询
	var sql = fmt.Sprintf("SELECT %s FROM question_option AS T0 %s %s %s %s", fields, relatedSql, whereSql, orderSql, pageSql)

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
	var countSql = fmt.Sprintf("SELECT count(*) AS count FROM question_option AS T0 %s %s", relatedSql, whereSql)
	var count RawCount
	err = orm.NewOrm().Raw(countSql, args...).QueryRow(&count)
	if err != nil {
		return
	}
	total = count.Count
	return
}

// 插入单个对象
func InsertQuestionOptionOne(m QuestionOption) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(&m)
	return
}

// 插入多个对象
func InsertQuestionOptionMulti(list []QuestionOption) (num int64, err error) {
	o := orm.NewOrm()
	num, err = o.InsertMulti(100, list)
	return
}

// 更新单个对象
func UpdateQuestionOptionOne(m QuestionOption, fields ...string) (num int64, err error) {
	o := orm.NewOrm()
	if len(fields) == 0 {
		fields = []string{"question_id", "tag", "content", "is_right", "memo"}
	}
	num, err = o.Update(&m, fields...)
	return
}

// 更新多个对象
func UpdateQuestionOptionMulti(ids []int, params map[string]interface{}) (num int64, err error) {
	o := orm.NewOrm()
	num, err = o.QueryTable(QuestionOptionTBName()).Filter("id__in", ids).Update(params)
	return
}

// 删除单个对象
func DeleteQuestionOptionOne(id int) (num int64, err error) {
	o := orm.NewOrm()
	num, err = o.Delete(&QuestionOption{ID: id})
	return
}

// 删除多个对象
func DeleteQuestionOptionMulti(ids []int) (num int64, err error) {
	o := orm.NewOrm()
	num, err = o.QueryTable(QuestionOptionTBName()).Filter("id__in", ids).Delete()
	return
}
