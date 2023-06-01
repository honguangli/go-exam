package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

// 试卷试题表
type PaperQuestion struct {
	ID           int     `orm:"column(id)" form:"id" json:"id"`
	PaperID      int     `orm:"column(paper_id)" form:"paper_id" json:"paper_id"`
	OriginID     int     `orm:"column(origin_id)" form:"origin_id" json:"origin_id"`
	SubjectID    int     `orm:"column(subject_id)" form:"subject_id" json:"subject_id"`
	Name         string  `orm:"column(name)" form:"name" json:"name"`
	Type         int     `orm:"column(type)" form:"type" json:"type"`
	Content      string  `orm:"column(content)" form:"content" json:"content"`
	Tips         string  `orm:"column(tips)" form:"tips" json:"tips"`
	Analysis     string  `orm:"column(analysis)" form:"analysis" json:"analysis"`
	Difficulty   float64 `orm:"column(difficulty)" form:"difficulty" json:"difficulty"`
	KnowledgeIds string  `orm:"column(knowledge_ids)" form:"knowledge_ids" json:"knowledge_ids"`
	Score        int     `orm:"column(score)" form:"score" json:"score"`
	UpdateTime   int64   `orm:"column(update_time)" form:"update_time" json:"update_time"`
	Memo         string  `orm:"column(memo)" form:"memo" json:"memo"`
}

// 查询列表参数
type ReadPaperQuestionListParam struct {
	BaseQueryParam
	PaperID   int  `json:"paper_id"`
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
	return [][]string{
		{"subject_id"},
		{"type"},
	}
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

	if param.PaperID > 0 {
		whereSql += " AND T0.`paper_id` = ?"
		args = append(args, param.PaperID)
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
	var fields = "T0.`id`, T0.`paper_id`, T0.`origin_id`, T0.`subject_id`, T0.`name`, T0.`type`, T0.`content`, T0.`tips`, T0.`analysis`, T0.`difficulty`, T0.`knowledge_ids`, T0.`score`, T0.`update_time`, T0.`memo`"

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
		fields = []string{"paper_id", "origin_id", "subject_id", "name", "type", "content", "tips", "analysis", "difficulty", "knowledge_ids", "score", "update_time", "memo"}
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
