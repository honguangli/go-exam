package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

// 用户角色表
type UserRoleRel struct {
	ID     int `orm:"column(id)" form:"id" json:"id"`
	UserID int `orm:"column(user_id)" form:"user_id" json:"user_id"`
	RoleID int `orm:"column(role_id)" form:"role_id" json:"role_id"`
}

// 查询列表参数
type ReadUserRoleRelListParam struct {
	BaseQueryParam
	ClosePage bool `form:"close_page" json:"close_page"`
}

// 初始化
func init() {
	orm.RegisterModel(new(UserRoleRel))
}

// 表名
func UserRoleRelTBName() (name string) {
	return "user_role_rel"
}

// 自定义表名
func (m *UserRoleRel) TableName() (name string) {
	return UserRoleRelTBName()
}

// 多字段索引
func (m *UserRoleRel) TableIndex() [][]string {
	return [][]string{}
}

// 多字段唯一键
func (m *UserRoleRel) TableUnique() [][]string {
	return [][]string{
		{"user_id", "role_id"},
	}
}

// 自定义引擎
func (m *UserRoleRel) TableEngine() string {
	return "INNODB"
}

// 查询单个对象
func ReadUserRoleRelOne(id int) (m UserRoleRel, err error) {
	o := orm.NewOrm()
	m.ID = id
	err = o.Read(&m)
	return
}

// 查询多个对象
func ReadUserRoleRelList(param ReadUserRoleRelListParam) (list []*UserRoleRel, total int64, err error) {
	list = make([]*UserRoleRel, 0)
	query := orm.NewOrm().QueryTable(UserRoleRelTBName())

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
func ReadUserRoleRelListRaw(param ReadUserRoleRelListParam) (list []*UserRoleRel, total int64, err error) {
	list = make([]*UserRoleRel, 0)
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
	var fields = "T0.`id`, T0.`user_id`, T0.`role_id`"

	// 关联查询
	var relatedSql string

	// 连表查询
	var sql = fmt.Sprintf("SELECT %s FROM user_role_rel AS T0 %s %s %s %s", fields, relatedSql, whereSql, orderSql, pageSql)

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
	var countSql = fmt.Sprintf("SELECT count(*) AS count FROM user_role_rel AS T0 %s %s", relatedSql, whereSql)
	var count RawCount
	err = orm.NewOrm().Raw(countSql, args...).QueryRow(&count)
	if err != nil {
		return
	}
	total = count.Count
	return
}

// 插入单个对象
func InsertUserRoleRelOne(m UserRoleRel) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(&m)
	return
}

// 插入多个对象
func InsertUserRoleRelMulti(list []UserRoleRel) (num int64, err error) {
	o := orm.NewOrm()
	num, err = o.InsertMulti(100, list)
	return
}

// 更新单个对象
func UpdateUserRoleRelOne(m UserRoleRel, fields ...string) (num int64, err error) {
	o := orm.NewOrm()
	if len(fields) == 0 {
		fields = []string{"user_id", "role_id"}
	}
	num, err = o.Update(&m, fields...)
	return
}

// 更新多个对象
func UpdateUserRoleRelMulti(ids []int, params map[string]interface{}) (num int64, err error) {
	o := orm.NewOrm()
	num, err = o.QueryTable(UserRoleRelTBName()).Filter("id__in", ids).Update(params)
	return
}

// 删除单个对象
func DeleteUserRoleRelOne(id int) (num int64, err error) {
	o := orm.NewOrm()
	num, err = o.Delete(&UserRoleRel{ID: id})
	return
}

// 删除多个对象
func DeleteUserRoleRelMulti(ids []int) (num int64, err error) {
	o := orm.NewOrm()
	num, err = o.QueryTable(UserRoleRelTBName()).Filter("id__in", ids).Delete()
	return
}
