package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

// 权限表
type Permission struct {
	ID   int    `orm:"column(id)" form:"id" json:"id"`
	Name string `orm:"column(name)" form:"name" json:"name"`
	Type int    `orm:"column(type)" form:"type" json:"type"`
	Pid  int    `orm:"column(pid)" form:"pid" json:"pid"`
	Icon string `orm:"column(icon)" form:"icon" json:"icon"`
	Path string `orm:"column(path)" form:"path" json:"path"`
	Memo string `orm:"column(memo)" form:"memo" json:"memo"`
}

// 查询列表参数
type ReadPermissionListParam struct {
	BaseQueryParam
	ClosePage bool `form:"close_page" json:"close_page"`
}

// 初始化
func init() {
	orm.RegisterModel(new(Permission))
}

// 表名
func PermissionTBName() (name string) {
	return "permission"
}

// 自定义表名
func (m *Permission) TableName() (name string) {
	return PermissionTBName()
}

// 多字段索引
func (m *Permission) TableIndex() [][]string {
	return [][]string{}
}

// 多字段唯一键
func (m *Permission) TableUnique() [][]string {
	return [][]string{}
}

// 自定义引擎
func (m *Permission) TableEngine() string {
	return "INNODB"
}

// 查询单个对象
func ReadPermissionOne(id int) (m Permission, err error) {
	o := orm.NewOrm()
	m.ID = id
	err = o.Read(&m)
	return
}

// 查询多个对象
func ReadPermissionList(param ReadPermissionListParam) (list []*Permission, total int64, err error) {
	list = make([]*Permission, 0)
	query := orm.NewOrm().QueryTable(PermissionTBName())

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
func ReadPermissionListRaw(param ReadPermissionListParam) (list []*Permission, total int64, err error) {
	list = make([]*Permission, 0)
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
	var fields = "T0.`id`, T0.`name`, T0.`type`, T0.`pid`, T0.`icon`, T0.`path`, T0.`memo`"

	// 关联查询
	var relatedSql string

	// 连表查询
	var sql = fmt.Sprintf("SELECT %s FROM permission AS T0 %s %s %s %s", fields, relatedSql, whereSql, orderSql, pageSql)

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
	var countSql = fmt.Sprintf("SELECT count(*) AS count FROM permission AS T0 %s %s", relatedSql, whereSql)
	var count RawCount
	err = orm.NewOrm().Raw(countSql, args...).QueryRow(&count)
	if err != nil {
		return
	}
	total = count.Count
	return
}

// 插入单个对象
func InsertPermissionOne(m Permission) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(&m)
	return
}

// 插入多个对象
func InsertPermissionMulti(list []Permission) (num int64, err error) {
	o := orm.NewOrm()
	num, err = o.InsertMulti(100, list)
	return
}

// 更新单个对象
func UpdatePermissionOne(m Permission, fields ...string) (num int64, err error) {
	o := orm.NewOrm()
	if len(fields) == 0 {
		fields = []string{"name", "type", "pid", "icon", "path", "memo"}
	}
	num, err = o.Update(&m, fields...)
	return
}

// 更新多个对象
func UpdatePermissionMulti(ids []int, params map[string]interface{}) (num int64, err error) {
	o := orm.NewOrm()
	num, err = o.QueryTable(PermissionTBName()).Filter("id__in", ids).Update(params)
	return
}

// 删除单个对象
func DeletePermissionOne(id int) (num int64, err error) {
	o := orm.NewOrm()
	num, err = o.Delete(&Permission{ID: id})
	return
}

// 删除多个对象
func DeletePermissionMulti(ids []int) (num int64, err error) {
	o := orm.NewOrm()
	num, err = o.QueryTable(PermissionTBName()).Filter("id__in", ids).Delete()
	return
}
