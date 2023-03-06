package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

// 考试成绩表
type Grade struct {
	ID              int    `orm:"column(id)" form:"id" json:"id"`
	PlanID          int    `orm:"column(plan_id)" form:"plan_id" json:"plan_id"`
	PaperID         int    `orm:"column(paper_id)" form:"paper_id" json:"paper_id"`
	UserID          int    `orm:"column(user_id)" form:"user_id" json:"user_id"`
	Score           int    `orm:"column(score)" form:"score" json:"score"`
	ObjectiveScore  int    `orm:"column(objective_score)" form:"objective_score" json:"objective_score"`
	SubjectiveScore int    `orm:"column(subjective_score)" form:"subjective_score" json:"subjective_score"`
	Status          int    `orm:"column(status)" form:"status" json:"status"`
	StartTime       int    `orm:"column(start_time)" form:"start_time" json:"start_time"`
	EndTime         int    `orm:"column(end_time)" form:"end_time" json:"end_time"`
	Duration        int    `orm:"column(duration)" form:"duration" json:"duration"`
	Memo            string `orm:"column(memo)" form:"memo" json:"memo"`
}

// 查询列表参数
type ReadGradeListParam struct {
	BaseQueryParam
	ClosePage bool `form:"close_page" json:"close_page"`
}

// 初始化
func init() {
	orm.RegisterModel(new(Grade))
}

// 表名
func GradeTBName() (name string) {
	return "grade"
}

// 自定义表名
func (m *Grade) TableName() (name string) {
	return GradeTBName()
}

// 多字段索引
func (m *Grade) TableIndex() [][]string {
	return [][]string{}
}

// 多字段唯一键
func (m *Grade) TableUnique() [][]string {
	return [][]string{}
}

// 自定义引擎
func (m *Grade) TableEngine() string {
	return "INNODB"
}

// 查询单个对象
func ReadGradeOne(id int) (m Grade, err error) {
	o := orm.NewOrm()
	m.ID = id
	err = o.Read(&m)
	return
}

// 查询多个对象
func ReadGradeList(param ReadGradeListParam) (list []*Grade, total int64, err error) {
	list = make([]*Grade, 0)
	query := orm.NewOrm().QueryTable(GradeTBName())

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
func ReadGradeListRaw(param ReadGradeListParam) (list []*Grade, total int64, err error) {
	list = make([]*Grade, 0)
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
	var fields = "T0.`id`, T0.`plan_id`, T0.`paper_id`, T0.`user_id`, T0.`score`, T0.`objective_score`, T0.`subjective_score`, T0.`status`, T0.`start_time`, T0.`end_time`, T0.`duration`, T0.`memo`"

	// 关联查询
	var relatedSql string

	// 连表查询
	var sql = fmt.Sprintf("SELECT %s FROM grade AS T0 %s %s %s %s", fields, relatedSql, whereSql, orderSql, pageSql)

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
	var countSql = fmt.Sprintf("SELECT count(*) AS count FROM grade AS T0 %s %s", relatedSql, whereSql)
	var count RawCount
	err = orm.NewOrm().Raw(countSql, args...).QueryRow(&count)
	if err != nil {
		return
	}
	total = count.Count
	return
}

// 插入单个对象
func InsertGradeOne(m Grade) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(&m)
	return
}

// 插入多个对象
func InsertGradeMulti(list []Grade) (num int64, err error) {
	o := orm.NewOrm()
	num, err = o.InsertMulti(100, list)
	return
}

// 更新单个对象
func UpdateGradeOne(m Grade, fields ...string) (num int64, err error) {
	o := orm.NewOrm()
	if len(fields) == 0 {
		fields = []string{"plan_id", "paper_id", "user_id", "score", "objective_score", "subjective_score", "status", "start_time", "end_time", "duration", "memo"}
	}
	num, err = o.Update(&m, fields...)
	return
}

// 更新多个对象
func UpdateGradeMulti(ids []int, params map[string]interface{}) (num int64, err error) {
	o := orm.NewOrm()
	num, err = o.QueryTable(GradeTBName()).Filter("id__in", ids).Update(params)
	return
}

// 删除单个对象
func DeleteGradeOne(id int) (num int64, err error) {
	o := orm.NewOrm()
	num, err = o.Delete(&Grade{ID: id})
	return
}

// 删除多个对象
func DeleteGradeMulti(ids []int) (num int64, err error) {
	o := orm.NewOrm()
	num, err = o.QueryTable(GradeTBName()).Filter("id__in", ids).Delete()
	return
}
