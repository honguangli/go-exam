package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

// 知识点表
type Knowledge struct {
	ID   int    `orm:"column(id)" form:"id" json:"id"`
	Name string `orm:"column(name)" form:"name" json:"name"`
	Desc string `orm:"column(desc)" form:"desc" json:"desc"`
}

// 查询详情参数
type ReadKnowledgeDetailParam struct {
	ID int `json:"id"`
}

// 查询列表参数
type ReadKnowledgeListParam struct {
	BaseQueryParam
	Name      string `json:"name"`
	ClosePage bool   `form:"close_page" json:"close_page"`
}

// 删除参数
type DeleteKnowledgeParam struct {
	ID   int   `json:"id"`
	List []int `json:"list"`
}

// 初始化
func init() {
	orm.RegisterModel(new(Knowledge))
}

// 表名
func KnowledgeTBName() (name string) {
	return "knowledge"
}

// 自定义表名
func (m *Knowledge) TableName() (name string) {
	return KnowledgeTBName()
}

// 多字段索引
func (m *Knowledge) TableIndex() [][]string {
	return [][]string{}
}

// 多字段唯一键
func (m *Knowledge) TableUnique() [][]string {
	return [][]string{}
}

// 自定义引擎
func (m *Knowledge) TableEngine() string {
	return "INNODB"
}

// 查询单个对象
func ReadKnowledgeOne(id int) (m Knowledge, err error) {
	o := orm.NewOrm()
	m.ID = id
	err = o.Read(&m)
	return
}

// 查询多个对象
func ReadKnowledgeList(param ReadKnowledgeListParam) (list []*Knowledge, total int64, err error) {
	list = make([]*Knowledge, 0)
	query := orm.NewOrm().QueryTable(KnowledgeTBName())

	if len(param.Name) > 0 {
		query = query.Filter("name__icontains", param.Name)
	}

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
func ReadKnowledgeListRaw(param ReadKnowledgeListParam) (list []*Knowledge, total int64, err error) {
	list = make([]*Knowledge, 0)
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
	var fields = "T0.`id`, T0.`name`, T0.`desc`"

	// 关联查询
	var relatedSql string

	// 连表查询
	var sql = fmt.Sprintf("SELECT %s FROM knowledge AS T0 %s %s %s %s", fields, relatedSql, whereSql, orderSql, pageSql)

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
	var countSql = fmt.Sprintf("SELECT count(*) AS count FROM knowledge AS T0 %s %s", relatedSql, whereSql)
	var count RawCount
	err = orm.NewOrm().Raw(countSql, args...).QueryRow(&count)
	if err != nil {
		return
	}
	total = count.Count
	return
}

// 插入单个对象
func InsertKnowledgeOne(m Knowledge) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(&m)
	return
}

// 插入多个对象
func InsertKnowledgeMulti(list []Knowledge) (num int64, err error) {
	o := orm.NewOrm()
	num, err = o.InsertMulti(100, list)
	return
}

// 更新单个对象
func UpdateKnowledgeOne(m Knowledge, fields ...string) (num int64, err error) {
	o := orm.NewOrm()
	if len(fields) == 0 {
		fields = []string{"name", "desc"}
	}
	num, err = o.Update(&m, fields...)
	return
}

// 更新多个对象
func UpdateKnowledgeMulti(ids []int, params map[string]interface{}) (num int64, err error) {
	o := orm.NewOrm()
	num, err = o.QueryTable(KnowledgeTBName()).Filter("id__in", ids).Update(params)
	return
}

// 删除单个对象
func DeleteKnowledgeOne(id int) (num int64, err error) {
	o := orm.NewOrm()
	num, err = o.Delete(&Knowledge{ID: id})
	return
}

// 删除多个对象
func DeleteKnowledgeMulti(ids []int) (num int64, err error) {
	o := orm.NewOrm()
	num, err = o.QueryTable(KnowledgeTBName()).Filter("id__in", ids).Delete()
	return
}
