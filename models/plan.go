package models

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/orm"
	"time"
)

// 考试计划表
type Plan struct {
	ID          int    `orm:"column(id)" form:"id" json:"id"`
	Name        string `orm:"column(name)" form:"name" json:"name"`
	PaperID     int    `orm:"column(paper_id)" form:"paper_id" json:"paper_id"`
	StartTime   int64  `orm:"column(start_time)" form:"start_time" json:"start_time"`
	EndTime     int64  `orm:"column(end_time)" form:"end_time" json:"end_time"`
	Duration    int    `orm:"column(duration)" form:"duration" json:"duration"`
	PublishTime int64  `orm:"column(publish_time)" form:"publish_time" json:"publish_time"`
	Status      int    `orm:"column(status)" form:"status" json:"status"`
	QueryGrade  int8   `orm:"column(query_grade)" form:"query_grade" json:"query_grade"`
	CreateTime  int64  `orm:"column(create_time)" form:"create_time" json:"create_time"`
	UpdateTime  int64  `orm:"column(update_time)" form:"update_time" json:"update_time"`
	Memo        string `orm:"column(memo)" form:"memo" json:"memo"`
}

// 考试计划表 关联
type PlanRel struct {
	ID          int    `orm:"column(id)" form:"id" json:"id"`
	Name        string `orm:"column(name)" form:"name" json:"name"`
	PaperID     int    `orm:"column(paper_id)" form:"paper_id" json:"paper_id"`
	StartTime   int64  `orm:"column(start_time)" form:"start_time" json:"start_time"`
	EndTime     int64  `orm:"column(end_time)" form:"end_time" json:"end_time"`
	Duration    int    `orm:"column(duration)" form:"duration" json:"duration"`
	PublishTime int64  `orm:"column(publish_time)" form:"publish_time" json:"publish_time"`
	Status      int    `orm:"column(status)" form:"status" json:"status"`
	QueryGrade  int8   `orm:"column(query_grade)" form:"query_grade" json:"query_grade"`
	CreateTime  int64  `orm:"column(create_time)" form:"create_time" json:"create_time"`
	UpdateTime  int64  `orm:"column(update_time)" form:"update_time" json:"update_time"`
	Memo        string `orm:"column(memo)" form:"memo" json:"memo"`

	// 试卷信息
	PaperName string `orm:"column(paper_name)" form:"paper_name" json:"paper_name"`
}

// 状态
const (
	PlanDefault   = 0 // 待发布
	PlanPublished = 1 // 已发布
	PlanCanceled  = 2 // 已取消
	PlanEnded     = 3 // 已结束
)

// 成绩状态
const (
	PlanGradeDisable = 0 // 未出成绩不可查询
	PlanGradeEnable  = 1 // 可查询
)

// 异常
var ErrStudentsEmpty = errors.New("考生列表为空")

// 查询详情参数
type ReadPlanDetailParam struct {
	ID int `json:"id"`
}

// 查询列表参数
type ReadPlanListParam struct {
	BaseQueryParam
	Name      string `json:"name"`
	Status    int    `json:"status"`
	ClosePage bool   `form:"close_page" json:"close_page"`
}

// 发布参数
type PublishPlanParam struct {
	ID int `json:"id"`
}

// 删除参数
type DeletePlanParam struct {
	ID int `json:"id"`
}

// 初始化
func init() {
	orm.RegisterModel(new(Plan))
}

// 表名
func PlanTBName() (name string) {
	return "plan"
}

// 自定义表名
func (m *Plan) TableName() (name string) {
	return PlanTBName()
}

// 多字段索引
func (m *Plan) TableIndex() [][]string {
	return [][]string{}
}

// 多字段唯一键
func (m *Plan) TableUnique() [][]string {
	return [][]string{}
}

// 自定义引擎
func (m *Plan) TableEngine() string {
	return "INNODB"
}

// 查询单个对象
func ReadPlanOne(id int) (m Plan, err error) {
	o := orm.NewOrm()
	m.ID = id
	err = o.Read(&m)
	return
}

// 查询单个对象
func ReadPlanRelOne(id int) (m PlanRel, err error) {
	o := orm.NewOrm()
	m.ID = id
	err = o.Read(&m)
	return
}

// 查询多个对象
func ReadPlanList(param ReadPlanListParam) (list []*Plan, total int64, err error) {
	list = make([]*Plan, 0)
	query := orm.NewOrm().QueryTable(PlanTBName())

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
func ReadPlanListRaw(param ReadPlanListParam) (list []*Plan, total int64, err error) {
	list = make([]*Plan, 0)
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
	var fields = "T0.`id`, T0.`name`, T0.`paper_id`, T0.`start_time`, T0.`end_time`, T0.`duration`, T0.`publish_time`, T0.`status`, T0.`query_grade`, T0.`create_time`, T0.`update_time`, T0.`memo`"

	// 关联查询
	var relatedSql string

	// 连表查询
	var sql = fmt.Sprintf("SELECT %s FROM plan AS T0 %s %s %s %s", fields, relatedSql, whereSql, orderSql, pageSql)

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
	var countSql = fmt.Sprintf("SELECT count(*) AS count FROM plan AS T0 %s %s", relatedSql, whereSql)
	var count RawCount
	err = orm.NewOrm().Raw(countSql, args...).QueryRow(&count)
	if err != nil {
		return
	}
	total = count.Count
	return
}

// 查询多个对象
func ReadPlanRelListRaw(param ReadPlanListParam) (list []*PlanRel, total int64, err error) {
	list = make([]*PlanRel, 0)
	var args = make([]interface{}, 0)
	var whereSql = "WHERE 1=1"

	if len(param.Name) > 0 {
		whereSql += " AND T0.`name` LIKE ?"
		args = append(args, fmt.Sprintf("%%%s%%", param.Name))
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
	var fields = "T0.`id`, T0.`name`, T0.`paper_id`, T0.`start_time`, T0.`end_time`, T0.`duration`, T0.`publish_time`, T0.`status`, T0.`query_grade`, T0.`create_time`, T0.`update_time`, T0.`memo`"

	// 关联查询
	var relatedSql = "LEFT JOIN paper AS T1 ON T1.id = T0.paper_id"
	fields += ", T1.`name` AS paper_name"

	// 连表查询
	var sql = fmt.Sprintf("SELECT %s FROM plan AS T0 %s %s %s %s", fields, relatedSql, whereSql, orderSql, pageSql)

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
	var countSql = fmt.Sprintf("SELECT count(*) AS count FROM plan AS T0 %s %s", relatedSql, whereSql)
	var count RawCount
	err = orm.NewOrm().Raw(countSql, args...).QueryRow(&count)
	if err != nil {
		return
	}
	total = count.Count
	return
}

// 插入单个对象
func InsertPlanOne(m Plan) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(&m)
	return
}

// 插入多个对象
func InsertPlanMulti(list []Plan) (num int64, err error) {
	o := orm.NewOrm()
	num, err = o.InsertMulti(100, list)
	return
}

// 更新单个对象
func UpdatePlanOne(m Plan, fields ...string) (num int64, err error) {
	o := orm.NewOrm()
	if len(fields) == 0 {
		fields = []string{"name", "paper_id", "start_time", "end_time", "duration", "publish_time", "status", "query_grade", "create_time", "update_time", "memo"}
	}
	num, err = o.Update(&m, fields...)
	return
}

// 更新多个对象
func UpdatePlanMulti(ids []int, params map[string]interface{}) (num int64, err error) {
	o := orm.NewOrm()
	num, err = o.QueryTable(PlanTBName()).Filter("id__in", ids).Update(params)
	return
}

// 删除单个对象
func DeletePlanOne(id int) (num int64, err error) {
	o := orm.NewOrm()
	num, err = o.Delete(&Plan{ID: id})
	return
}

// 删除多个对象
func DeletePlanMulti(ids []int) (num int64, err error) {
	o := orm.NewOrm()
	num, err = o.QueryTable(PlanTBName()).Filter("id__in", ids).Delete()
	return
}

// 发布考试计划
func PublishPlan(param PublishPlanParam) (err error) {
	o := orm.NewOrm()

	// 查询考试计划
	var m = Plan{
		ID: param.ID,
	}
	err = o.Read(&m)
	if err != nil {
		return
	}

	if m.Status != PlanDefault {
		return errors.New("状态异常")
	}

	// 查询考生列表
	var list = make([]ClassUserRel, 0)
	_, err = o.Raw("SELECT DISTINCT T0.user_id FROM class_user_rel AS T0"+
		" LEFT JOIN plan_class_rel AS T1 ON T1.class_id = T0.class_id"+
		" WHERE T1.plan_id = ?"+
		" ORDER BY T0.user_id ASC", param.ID).QueryRows(&list)
	if err == orm.ErrNoRows {
		return ErrStudentsEmpty
	}
	if err != nil {
		return
	}
	if len(list) == 0 {
		return ErrStudentsEmpty
	}

	var grades = make([]Grade, len(list))
	for k, v := range list {
		grades[k] = Grade{
			PlanID:          m.ID,
			PaperID:         m.PaperID,
			UserID:          v.UserID,
			Score:           0,
			ObjectiveScore:  0,
			SubjectiveScore: 0,
			Status:          GradeDefault,
			StartTime:       0,
			EndTime:         0,
			Duration:        0,
			Memo:            "",
		}
	}

	var nowUnix = time.Now().Unix()

	// 开启事务
	err = o.Begin()
	if err != nil {
		return
	}

	// 更新考试计划状态
	m.PublishTime = nowUnix
	m.Status = PlanPublished
	m.UpdateTime = nowUnix
	_, err = o.Update(&m, "publish_time", "status", "update_time")
	if err != nil {
		o.Rollback()
		return
	}

	// 保存考生考试计划
	_, err = o.InsertMulti(100, grades)
	if err != nil {
		o.Rollback()
		return
	}

	err = o.Commit()
	return
}
