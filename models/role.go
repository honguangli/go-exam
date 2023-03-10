package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"strings"
)

// 角色表
type Role struct {
	ID         int    `orm:"column(id)" form:"id" json:"id"`
	Name       string `orm:"column(name)" form:"name" json:"name"`
	Code       string `orm:"column(code)" form:"code" json:"code"`
	Seq        int    `orm:"column(seq)" form:"seq" json:"seq"`
	Status     int    `orm:"column(status)" form:"status" json:"status"`
	CreateTime int64  `orm:"column(create_time)" form:"create_time" json:"create_time"`
	UpdateTime int64  `orm:"column(update_time)" form:"update_time" json:"update_time"`
	Memo       string `orm:"column(memo)" form:"memo" json:"memo"`
}

// 状态
const (
	ROLE_STATUS_IGNORE = -1 // 忽略状态
	ROLE_DISABLE       = 0  // 禁用
	ROLE_ENABLE        = 1  // 正常
)

// 查询详情参数
type ReadRoleDetailParam struct {
	ID int `json:"id"`
}

// 查询列表参数
type ReadRoleListParam struct {
	BaseQueryParam
	Name      string `json:"name"`
	Code      string `json:"code"`
	Status    int    `json:"status"`
	ClosePage bool   `form:"close_page" json:"close_page"`
}

// 删除参数
type DeleteRoleParam struct {
	ID   int   `json:"id"`
	List []int `json:"list"`
}

// 更新角色权限参数
type UpdateRolePermissionParam struct {
	ID             int   `json:"id"`
	PermissionList []int `json:"permission_list"`
}

// 初始化
func init() {
	orm.RegisterModel(new(Role))
}

// 表名
func RoleTBName() (name string) {
	return "role"
}

// 自定义表名
func (m *Role) TableName() (name string) {
	return RoleTBName()
}

// 多字段索引
func (m *Role) TableIndex() [][]string {
	return [][]string{}
}

// 多字段唯一键
func (m *Role) TableUnique() [][]string {
	return [][]string{
		{"code"},
	}
}

// 自定义引擎
func (m *Role) TableEngine() string {
	return "INNODB"
}

// 查询单个对象
func ReadRoleOne(id int) (m Role, err error) {
	o := orm.NewOrm()
	m.ID = id
	err = o.Read(&m)
	return
}

// 查询多个对象
func ReadRoleList(param ReadRoleListParam) (list []*Role, total int64, err error) {
	list = make([]*Role, 0)
	query := orm.NewOrm().QueryTable(RoleTBName())

	if len(param.Name) > 0 {
		query = query.Filter("name__icontains", param.Name)
	}

	if len(param.Code) > 0 {
		query = query.Filter("code", param.Code)
	}

	if param.Status != ROLE_STATUS_IGNORE {
		query = query.Filter("status", param.Status)
	}

	sortOrder := "id"
	switch param.Sort {
	case "seq":
		sortOrder = "seq"
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
func ReadRoleListRaw(param ReadRoleListParam) (list []*Role, total int64, err error) {
	list = make([]*Role, 0)
	var args = make([]interface{}, 0)
	var whereSql = "WHERE 1=1"

	// 排序
	var orderSql = "ORDER BY "
	switch param.Sort {
	case "seq":
		orderSql += "T0.seq"
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
	var fields = "T0.`id`, T0.`name`, T0.`code`, T0.`seq`, T0.`status`, T0.`create_time`, T0.`update_time`, T0.`memo`"

	// 关联查询
	var relatedSql string

	// 连表查询
	var sql = fmt.Sprintf("SELECT %s FROM role AS T0 %s %s %s %s", fields, relatedSql, whereSql, orderSql, pageSql)

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
	var countSql = fmt.Sprintf("SELECT count(*) AS count FROM role AS T0 %s %s", relatedSql, whereSql)
	var count RawCount
	err = orm.NewOrm().Raw(countSql, args...).QueryRow(&count)
	if err != nil {
		return
	}
	total = count.Count
	return
}

// 插入单个对象
func InsertRoleOne(m Role) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(&m)
	return
}

// 插入多个对象
func InsertRoleMulti(list []Role) (num int64, err error) {
	o := orm.NewOrm()
	num, err = o.InsertMulti(100, list)
	return
}

// 更新单个对象
func UpdateRoleOne(m Role, fields ...string) (num int64, err error) {
	o := orm.NewOrm()
	if len(fields) == 0 {
		fields = []string{"name", "code", "seq", "status", "create_time", "update_time", "memo"}
	}
	num, err = o.Update(&m, fields...)
	return
}

// 更新多个对象
func UpdateRoleMulti(ids []int, params map[string]interface{}) (num int64, err error) {
	o := orm.NewOrm()
	num, err = o.QueryTable(RoleTBName()).Filter("id__in", ids).Update(params)
	return
}

// 删除单个对象
func DeleteRoleOne(id int) (num int64, err error) {
	o := orm.NewOrm()
	num, err = o.Delete(&Role{ID: id})
	return
}

// 删除多个对象
func DeleteRoleMulti(ids []int) (num int64, err error) {
	o := orm.NewOrm()
	num, err = o.QueryTable(RoleTBName()).Filter("id__in", ids).Delete()
	return
}

// 更新角色权限
func UpdateRolePermissionMulti(param UpdateRolePermissionParam) (err error) {
	var list = make([]*RolePermissionRel, len(param.PermissionList))
	for k, v := range param.PermissionList {
		list[k] = &RolePermissionRel{
			RoleID:       param.ID,
			PermissionID: v,
		}
	}

	o := orm.NewOrm()
	err = o.Begin()
	if err != nil {
		return
	}

	// 删除旧关系
	_, err = o.Raw("DELETE FROM role_permission_rel WHERE role_id = ?", param.ID).Exec()
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

// 删除多个对象
func DeleteRoleOneWithRel(id int) (num int64, err error) {
	o := orm.NewOrm()

	err = o.Begin()
	if err != nil {
		return
	}

	// 删除角色权限
	_, err = o.Raw("DELETE FROM role_permission_rel WHERE role_id = ?", id).Exec()
	if err != nil {
		o.Rollback()
		return
	}

	// 删除用户角色
	_, err = o.Raw("DELETE FROM user_role_rel WHERE role_id = ?", id).Exec()
	if err != nil {
		o.Rollback()
		return
	}

	// 删除角色
	_, err = o.Raw("DELETE FROM role WHERE id = ?", id).Exec()
	if err != nil {
		o.Rollback()
		return
	}

	err = o.Commit()
	return
}

// 删除多个对象
func DeleteRoleMultiWithRel(ids []int) (num int64, err error) {
	o := orm.NewOrm()

	err = o.Begin()
	if err != nil {
		return
	}

	// 删除角色权限
	_, err = o.Raw(fmt.Sprintf("DELETE FROM role_permission_rel WHERE role_id IN (?%s)", strings.Repeat(", ?", len(ids)-1)), ids).Exec()
	if err != nil {
		o.Rollback()
		return
	}

	// 删除用户角色
	_, err = o.Raw(fmt.Sprintf("DELETE FROM user_role_rel WHERE role_id IN (?%s)", strings.Repeat(", ?", len(ids)-1)), ids).Exec()
	if err != nil {
		o.Rollback()
		return
	}

	// 删除角色
	_, err = o.Raw(fmt.Sprintf("DELETE FROM role WHERE id IN (?%s)", strings.Repeat(", ?", len(ids)-1)), ids).Exec()
	if err != nil {
		o.Rollback()
		return
	}

	err = o.Commit()
	return
}
