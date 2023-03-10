package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

// 班级表
type Class struct {
	ID     int    `orm:"column(id)" form:"id" json:"id"`
	Name   string `orm:"column(name)" form:"name" json:"name"`
	Status int    `orm:"column(status)" form:"status" json:"status"`
	Desc   string `orm:"column(desc)" form:"desc" json:"desc"`
}

// 状态
const (
	CLASS_STATUS_IGNORE = -1 // 忽略状态
	CLASS_DISABLE       = 0  // 禁用
	CLASS_ENABLE        = 1  // 正常
)

// 查询详情参数
type ReadClassDetailParam struct {
	ID int `json:"id"`
}

// 查询列表参数
type ReadClassListParam struct {
	BaseQueryParam
	Name      string `json:"name"`
	Status    int    `json:"status"`
	ClosePage bool   `form:"close_page" json:"close_page"`
}

// 删除参数
type DeleteClassParam struct {
	ID   int   `json:"id"`
	List []int `json:"list"`
}

// 更新班级用户参数
type UpdateClassUserParam struct {
	ID       int   `json:"id"`
	UserList []int `json:"user_list"`
}

// 初始化
func init() {
	orm.RegisterModel(new(Class))
}

// 表名
func ClassTBName() (name string) {
	return "class"
}

// 自定义表名
func (m *Class) TableName() (name string) {
	return ClassTBName()
}

// 多字段索引
func (m *Class) TableIndex() [][]string {
	return [][]string{}
}

// 多字段唯一键
func (m *Class) TableUnique() [][]string {
	return [][]string{}
}

// 自定义引擎
func (m *Class) TableEngine() string {
	return "INNODB"
}

// 查询单个对象
func ReadClassOne(id int) (m Class, err error) {
	o := orm.NewOrm()
	m.ID = id
	err = o.Read(&m)
	return
}

// 查询多个对象
func ReadClassList(param ReadClassListParam) (list []*Class, total int64, err error) {
	list = make([]*Class, 0)
	query := orm.NewOrm().QueryTable(ClassTBName())

	if len(param.Name) > 0 {
		query = query.Filter("name_icontains", param.Name)
	}

	if param.Status != CLASS_STATUS_IGNORE {
		query = query.Filter("status", param.Status)
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
func ReadClassListRaw(param ReadClassListParam) (list []*Class, total int64, err error) {
	list = make([]*Class, 0)
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
	var fields = "T0.`id`, T0.`name`, T0.`status`, T0.`desc`"

	// 关联查询
	var relatedSql string

	// 连表查询
	var sql = fmt.Sprintf("SELECT %s FROM class AS T0 %s %s %s %s", fields, relatedSql, whereSql, orderSql, pageSql)

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
	var countSql = fmt.Sprintf("SELECT count(*) AS count FROM class AS T0 %s %s", relatedSql, whereSql)
	var count RawCount
	err = orm.NewOrm().Raw(countSql, args...).QueryRow(&count)
	if err != nil {
		return
	}
	total = count.Count
	return
}

// 插入单个对象
func InsertClassOne(m Class) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(&m)
	return
}

// 插入多个对象
func InsertClassMulti(list []Class) (num int64, err error) {
	o := orm.NewOrm()
	num, err = o.InsertMulti(100, list)
	return
}

// 更新单个对象
func UpdateClassOne(m Class, fields ...string) (num int64, err error) {
	o := orm.NewOrm()
	if len(fields) == 0 {
		fields = []string{"name", "status", "desc"}
	}
	num, err = o.Update(&m, fields...)
	return
}

// 更新多个对象
func UpdateClassMulti(ids []int, params map[string]interface{}) (num int64, err error) {
	o := orm.NewOrm()
	num, err = o.QueryTable(ClassTBName()).Filter("id__in", ids).Update(params)
	return
}

// 删除单个对象
func DeleteClassOne(id int) (num int64, err error) {
	o := orm.NewOrm()
	num, err = o.Delete(&Class{ID: id})
	return
}

// 删除多个对象
func DeleteClassMulti(ids []int) (num int64, err error) {
	o := orm.NewOrm()
	num, err = o.QueryTable(ClassTBName()).Filter("id__in", ids).Delete()
	return
}

// 更新班级用户
func UpdateClassUserMulti(param UpdateClassUserParam) (err error) {
	var list = make([]*ClassUserRel, len(param.UserList))
	for k, v := range param.UserList {
		list[k] = &ClassUserRel{
			ClassID: param.ID,
			UserID:  v,
		}
	}

	o := orm.NewOrm()
	err = o.Begin()
	if err != nil {
		return
	}

	// 删除旧关系
	_, err = o.Raw("DELETE FROM class_user_rel WHERE class_id = ?", param.ID).Exec()
	if err != nil {
		o.Rollback()
		return
	}

	// 添加新关系
	_, err = o.InsertMulti(100, list)
	if err != nil {
		o.Rollback()
		return
	}

	err = o.Commit()
	return
}
