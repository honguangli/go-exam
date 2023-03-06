package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

// 角色权限表
type RolePermissionRel struct {
	ID           int `orm:"column(id)" form:"id" json:"id"`
	RoleID       int `orm:"column(role_id)" form:"role_id" json:"role_id"`
	PermissionID int `orm:"column(permission_id)" form:"permission_id" json:"permission_id"`
}

// 查询列表参数
type ReadRolePermissionRelListParam struct {
	BaseQueryParam
	ClosePage bool `form:"close_page" json:"close_page"`
}

// 初始化
func init() {
	orm.RegisterModel(new(RolePermissionRel))
}

// 表名
func RolePermissionRelTBName() (name string) {
	return "role_permission_rel"
}

// 自定义表名
func (m *RolePermissionRel) TableName() (name string) {
	return RolePermissionRelTBName()
}

// 多字段索引
func (m *RolePermissionRel) TableIndex() [][]string {
	return [][]string{}
}

// 多字段唯一键
func (m *RolePermissionRel) TableUnique() [][]string {
	return [][]string{
		{"role_id", "permission_id"},
	}
}

// 自定义引擎
func (m *RolePermissionRel) TableEngine() string {
	return "INNODB"
}

// 查询单个对象
func ReadRolePermissionRelOne(id int) (m RolePermissionRel, err error) {
	o := orm.NewOrm()
	m.ID = id
	err = o.Read(&m)
	return
}

// 查询多个对象
func ReadRolePermissionRelList(param ReadRolePermissionRelListParam) (list []*RolePermissionRel, total int64, err error) {
	list = make([]*RolePermissionRel, 0)
	query := orm.NewOrm().QueryTable(RolePermissionRelTBName())

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
func ReadRolePermissionRelListRaw(param ReadRolePermissionRelListParam) (list []*RolePermissionRel, total int64, err error) {
	list = make([]*RolePermissionRel, 0)
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
	var fields = "T0.`id`, T0.`role_id`, T0.`permission_id`"

	// 关联查询
	var relatedSql string

	// 连表查询
	var sql = fmt.Sprintf("SELECT %s FROM role_permission_rel AS T0 %s %s %s %s", fields, relatedSql, whereSql, orderSql, pageSql)

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
	var countSql = fmt.Sprintf("SELECT count(*) AS count FROM role_permission_rel AS T0 %s %s", relatedSql, whereSql)
	var count RawCount
	err = orm.NewOrm().Raw(countSql, args...).QueryRow(&count)
	if err != nil {
		return
	}
	total = count.Count
	return
}

// 插入单个对象
func InsertRolePermissionRelOne(m RolePermissionRel) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(&m)
	return
}

// 插入多个对象
func InsertRolePermissionRelMulti(list []RolePermissionRel) (num int64, err error) {
	o := orm.NewOrm()
	num, err = o.InsertMulti(100, list)
	return
}

// 更新单个对象
func UpdateRolePermissionRelOne(m RolePermissionRel, fields ...string) (num int64, err error) {
	o := orm.NewOrm()
	if len(fields) == 0 {
		fields = []string{"role_id", "permission_id"}
	}
	num, err = o.Update(&m, fields...)
	return
}

// 更新多个对象
func UpdateRolePermissionRelMulti(ids []int, params map[string]interface{}) (num int64, err error) {
	o := orm.NewOrm()
	num, err = o.QueryTable(RolePermissionRelTBName()).Filter("id__in", ids).Update(params)
	return
}

// 删除单个对象
func DeleteRolePermissionRelOne(id int) (num int64, err error) {
	o := orm.NewOrm()
	num, err = o.Delete(&RolePermissionRel{ID: id})
	return
}

// 删除多个对象
func DeleteRolePermissionRelMulti(ids []int) (num int64, err error) {
	o := orm.NewOrm()
	num, err = o.QueryTable(RolePermissionRelTBName()).Filter("id__in", ids).Delete()
	return
}
