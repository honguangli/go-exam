package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"time"
)

// 考生试卷答题表
type Answer struct {
	ID         int    `orm:"column(id)" form:"id" json:"id"`
	PlanID     int    `orm:"column(plan_id)" form:"plan_id" json:"plan_id"`
	PaperID    int    `orm:"column(paper_id)" form:"paper_id" json:"paper_id"`
	UserID     int    `orm:"column(user_id)" form:"user_id" json:"user_id"`
	GradeID    int    `orm:"column(grade_id)" form:"grade_id" json:"grade_id"`
	SubmitTime int64  `orm:"column(submit_time)" form:"submit_time" json:"submit_time"`
	Memo       string `orm:"column(memo)" form:"memo" json:"memo"`
}

// 查询列表参数
type ReadAnswerListParam struct {
	BaseQueryParam
	ClosePage bool `form:"close_page" json:"close_page"`
}

// 初始化
func init() {
	orm.RegisterModel(new(Answer))
}

// 表名
func AnswerTBName() (name string) {
	return "answer"
}

// 自定义表名
func (m *Answer) TableName() (name string) {
	return AnswerTBName()
}

// 多字段索引
func (m *Answer) TableIndex() [][]string {
	return [][]string{}
}

// 多字段唯一键
func (m *Answer) TableUnique() [][]string {
	return [][]string{}
}

// 自定义引擎
func (m *Answer) TableEngine() string {
	return "INNODB"
}

// 查询单个对象
func ReadAnswerOne(id int) (m Answer, err error) {
	o := orm.NewOrm()
	m.ID = id
	err = o.Read(&m)
	return
}

// 查询多个对象
func ReadAnswerList(param ReadAnswerListParam) (list []*Answer, total int64, err error) {
	list = make([]*Answer, 0)
	query := orm.NewOrm().QueryTable(AnswerTBName())

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
func ReadAnswerListRaw(param ReadAnswerListParam) (list []*Answer, total int64, err error) {
	list = make([]*Answer, 0)
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
	var fields = "T0.`id`, T0.`plan_id`, T0.`paper_id`, T0.`user_id`, T0.`grade_id`, T0.`submit_time`, T0.`memo`"

	// 关联查询
	var relatedSql string

	// 连表查询
	var sql = fmt.Sprintf("SELECT %s FROM answer AS T0 %s %s %s %s", fields, relatedSql, whereSql, orderSql, pageSql)

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
	var countSql = fmt.Sprintf("SELECT count(*) AS count FROM answer AS T0 %s %s", relatedSql, whereSql)
	var count RawCount
	err = orm.NewOrm().Raw(countSql, args...).QueryRow(&count)
	if err != nil {
		return
	}
	total = count.Count
	return
}

// 插入单个对象
func InsertAnswerOne(m Answer) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(&m)
	return
}

// 插入多个对象
func InsertAnswerMulti(list []Answer) (num int64, err error) {
	o := orm.NewOrm()
	num, err = o.InsertMulti(100, list)
	return
}

// 更新单个对象
func UpdateAnswerOne(m Answer, fields ...string) (num int64, err error) {
	o := orm.NewOrm()
	if len(fields) == 0 {
		fields = []string{"plan_id", "paper_id", "user_id", "grade_id", "submit_time", "memo"}
	}
	num, err = o.Update(&m, fields...)
	return
}

// 更新多个对象
func UpdateAnswerMulti(ids []int, params map[string]interface{}) (num int64, err error) {
	o := orm.NewOrm()
	num, err = o.QueryTable(AnswerTBName()).Filter("id__in", ids).Update(params)
	return
}

// 删除单个对象
func DeleteAnswerOne(id int) (num int64, err error) {
	o := orm.NewOrm()
	num, err = o.Delete(&Answer{ID: id})
	return
}

// 删除多个对象
func DeleteAnswerMulti(ids []int) (num int64, err error) {
	o := orm.NewOrm()
	num, err = o.QueryTable(AnswerTBName()).Filter("id__in", ids).Delete()
	return
}

// 保存答题卡
func SubmitAnswer(grade *GradeRel, param SubmitExamParam) (err error) {
	o := orm.NewOrm()
	err = o.Begin()
	if err != nil {
		return
	}

	var nowUnix = time.Now().Unix()

	// 更新考试状态
	_, err = o.Update(&Grade{ID: grade.ID, Status: GradeSubmit, EndTime: nowUnix, Duration: nowUnix - grade.StartTime}, "status", "end_time", "duration")
	if err != nil {
		o.Rollback()
		return
	}

	// 保存答题卡
	var m = Answer{
		PlanID:     grade.PlanID,
		PaperID:    grade.PaperID,
		UserID:     grade.UserID,
		GradeID:    grade.ID,
		SubmitTime: nowUnix,
	}
	_, err = o.Insert(&m)
	if err != nil {
		o.Rollback()
		return
	}

	// 保存答题项
	var items = make([]*AnswerItem, len(param.Answers))
	for k, v := range param.Answers {
		v.AnswerID = m.ID
		v.Check = AnswerItemUnCheck
		v.Score = 0
		items[k] = &v
	}
	_, err = o.InsertMulti(100, items)
	if err != nil {
		o.Rollback()
		return
	}

	err = o.Commit()
	return
}
