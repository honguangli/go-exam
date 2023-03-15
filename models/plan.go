package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

// 考试计划表
type Plan struct {
	ID          int    `orm:"column(id)" form:"id" json:"id"`
	Name        string `orm:"column(name)" form:"name" json:"name"`
	PaperID     int    `orm:"column(paper_id)" form:"paper_id" json:"paper_id"`
	StartTime   int    `orm:"column(start_time)" form:"start_time" json:"start_time"`
	EndTime     int    `orm:"column(end_time)" form:"end_time" json:"end_time"`
	Duration    int    `orm:"column(duration)" form:"duration" json:"duration"`
	PublishTime int    `orm:"column(publish_time)" form:"publish_time" json:"publish_time"`
	Status      int    `orm:"column(status)" form:"status" json:"status"`
	QueryGrade  int8   `orm:"column(query_grade)" form:"query_grade" json:"query_grade"`
	CreateTime  int64  `orm:"column(create_time)" form:"create_time" json:"create_time"`
	UpdateTime  int64  `orm:"column(update_time)" form:"update_time" json:"update_time"`
	Memo        string `orm:"column(memo)" form:"memo" json:"memo"`
}

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

// 删除参数
type DeletePlanParam struct {
	ID   int   `json:"id"`
	List []int `json:"list"`
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
