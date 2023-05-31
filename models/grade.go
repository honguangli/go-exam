package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"strings"
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

// 考试成绩表 关联
type GradeRel struct {
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

	// 考试计划信息
	PlanName       string `orm:"column(plan_name)" form:"plan_name" json:"plan_name"`
	PlanStartTime  int64  `orm:"column(plan_start_time)" form:"plan_start_time" json:"plan_start_time"`
	PlanEndTime    int64  `orm:"column(plan_end_time)" form:"plan_end_time" json:"plan_end_time"`
	PlanDuration   int    `orm:"column(plan_duration)" form:"plan_duration" json:"plan_duration"`
	PlanStatus     int    `orm:"column(plan_status)" form:"plan_status" json:"plan_status"`
	PlanQueryGrade int8   `orm:"column(plan_query_grade)" form:"plan_query_grade" json:"plan_query_grade"`

	// 试卷信息
	PaperName         string `orm:"column(paper_name)" form:"paper_name" json:"paper_name"`
	PaperSubjectID    int    `orm:"column(paper_subject_id)" form:"paper_subject_id" json:"paper_subject_id"`
	PaperKnowledgeIds string `orm:"column(paper_knowledge_ids)" form:"paper_knowledge_ids" json:"paper_knowledge_ids"`
	PaperScore        int    `orm:"column(paper_score)" form:"paper_score" json:"paper_score"`
	PaperPassScore    int    `orm:"column(paper_pass_score)" form:"paper_pass_score" json:"paper_pass_score"`

	// 考生信息
	UserName     string `orm:"column(user_name)" form:"user_name" json:"user_name"`
	UserTrueName string `orm:"column(user_true_name)" form:"user_true_name" json:"user_true_name"`
	UserMobile   string `orm:"column(user_mobile)" form:"user_mobile" json:"user_mobile"`
	UserEmail    string `orm:"column(user_email)" form:"user_email" json:"user_email"`
	UserStatus   int    `orm:"column(user_status)" form:"user_status" json:"user_status"`
}

// 状态
const (
	GradeDefault  = 0 // 待参加考试
	GradeUnSubmit = 1 // 待交卷
	GradeSubmit   = 2 // 已交卷待评分
	GradeMarking  = 3 // 部分评分
	GradeMarked   = 4 // 评分完成
	GradeCancel   = 5 // 考试取消
)

// 查询列表参数
type ReadGradeListParam struct {
	BaseQueryParam
	ClosePage bool `form:"close_page" json:"close_page"`
}

// 查询列表参数
type ReadGradeRelListParam struct {
	BaseQueryParam
	PlanID         int    `json:"plan_id"`
	UserID         int    `json:"user_id"`
	StatusList     []int  `json:"status"`
	PlanName       string `json:"plan_name"`
	PlanStatusList []int  `json:"plan_status"`
	PlanQueryGrade int    `json:"plan_query_grade"`
	UserName       string `json:"user_name"`
	ClosePage      bool   `form:"close_page" json:"close_page"`
}

// 查询详情参数
type ReadGradeDetailParam struct {
	ID int `json:"id"`
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

// 查询多个对象
func ReadGradeRelListRaw(param ReadGradeRelListParam) (list []*GradeRel, total int64, err error) {
	list = make([]*GradeRel, 0)
	var args = make([]interface{}, 0)
	var whereSql = "WHERE 1=1"

	if param.PlanID > 0 {
		whereSql += " AND T0.`plan_id` = ?"
		args = append(args, param.PlanID)
	}

	if param.UserID > 0 {
		whereSql += " AND T0.`user_id` = ?"
		args = append(args, param.UserID)
	}

	if len(param.StatusList) > 0 {
		whereSql += fmt.Sprintf(" AND T0.`status` IN (?%s)", strings.Repeat(", ?", len(param.StatusList)-1))
		args = append(args, param.StatusList)
	}

	if len(param.PlanName) > 0 {
		whereSql += " AND T1.`name` LIKE ?"
		args = append(args, fmt.Sprintf("%%%s%%", param.PlanName))
	}

	if len(param.PlanStatusList) > 0 {
		whereSql += fmt.Sprintf(" AND T1.`status` IN (?%s)", strings.Repeat(", ?", len(param.PlanStatusList)-1))
		args = append(args, param.PlanStatusList)
	}

	if param.PlanQueryGrade >= 0 {
		whereSql += " AND T1.`query_grade` = ?"
		args = append(args, param.PlanQueryGrade)
	}

	if len(param.UserName) > 0 {
		whereSql += " AND T3.`name` = ?"
		args = append(args, param.UserName)
	}

	// 排序
	var orderSql = "ORDER BY "
	switch param.Sort {
	case "score":
		orderSql += "T0.score"
	case "objective_score":
		orderSql += "T0.objective_score"
	case "subjective_score":
		orderSql += "T0.subjective_score"
	case "start_time":
		orderSql += "T0.start_time"
	case "end_time":
		orderSql += "T0.end_time"
	case "duration":
		orderSql += "T0.duration"
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
	var relatedSql = "LEFT JOIN plan AS T1 ON T1.id = T0.plan_id" +
		" LEFT JOIN paper AS T2 ON T2.id = T0.paper_id" +
		" LEFT JOIN user AS T3 ON T3.id = T0.user_id"
	fields += ", T1.`name` AS plan_name, T1.`start_time` AS plan_start_time, T1.`end_time` AS plan_end_time, T1.`duration` AS plan_duration, T1.`status` AS plan_status, T1.`query_grade` AS plan_query_grade" +
		", T2.`name` AS paper_name, T2.`subject_id` AS paper_subject_id, T2.`knowledge_ids` AS paper_knowledge_ids, T2.`score` AS paper_score, T2.`pass_score` AS paper_pass_score" +
		", T3.`name` AS user_name, T3.`true_name` AS user_true_name, T3.`mobile` AS user_mobile, T3.`email` AS user_email, T3.`status` AS user_status"

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
