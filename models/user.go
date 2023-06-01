package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"strings"
)

// 用户表
type User struct {
	ID         int    `orm:"column(id)" form:"id" json:"id"`
	Name       string `orm:"column(name)" form:"name" json:"name"`
	Password   string `orm:"column(password)" form:"password" json:"password"`
	Type       int    `orm:"column(type)" form:"type" json:"type"`
	TrueName   string `orm:"column(true_name)" form:"true_name" json:"true_name"`
	Mobile     string `orm:"column(mobile)" form:"mobile" json:"mobile"`
	Email      string `orm:"column(email)" form:"email" json:"email"`
	Status     int    `orm:"column(status)" form:"status" json:"status"`
	CreateTime int64  `orm:"column(create_time)" form:"create_time" json:"create_time"`
	UpdateTime int64  `orm:"column(update_time)" form:"update_time" json:"update_time"`
	Memo       string `orm:"column(memo)" form:"memo" json:"memo"`
}

// 类型
const (
	UserTypeIgnore = -1 // 忽略类型
	UserAdmin      = 1  // 管理员
	UserTeacher    = 2  // 教师
	UserStudent    = 3  // 学生
)

// 状态
const (
	UserStatusIgnore = -1 // 忽略状态
	UserDisable      = 0  // 禁用
	UserEnable       = 1  // 正常
)

// 默认密码
const DEFAULT_PASSWORD = "123456"

// 查询详情参数
type ReadUserDetailParam struct {
	ID int `json:"id"`
}

// 查询列表参数
type ReadUserListParam struct {
	BaseQueryParam
	Name      string `json:"name"`
	Type      int    `json:"type"`
	Status    int    `json:"status"`
	ClosePage bool   `form:"close_page" json:"close_page"`
}

// 删除参数
type DeleteUserParam struct {
	ID   int   `json:"id"`
	List []int `json:"list"`
}

// 更新用户角色参数
type UpdateUserRoleParam struct {
	ID       int   `json:"id"`
	RoleList []int `json:"role_list"`
}

// 初始化
func init() {
	orm.RegisterModel(new(User))
}

// 表名
func UserTBName() (name string) {
	return "user"
}

// 自定义表名
func (m *User) TableName() (name string) {
	return UserTBName()
}

// 多字段索引
func (m *User) TableIndex() [][]string {
	return [][]string{}
}

// 多字段唯一键
func (m *User) TableUnique() [][]string {
	return [][]string{
		{"name"},
	}
}

// 自定义引擎
func (m *User) TableEngine() string {
	return "INNODB"
}

// 查询单个对象
func ReadUserOne(id int) (m User, err error) {
	o := orm.NewOrm()
	m.ID = id
	err = o.Read(&m)
	return
}

// 查询单个对象
func ReadUserOneByName(name string) (m User, err error) {
	err = orm.NewOrm().QueryTable(UserTBName()).Filter("name", name).One(&m)
	return
}

// 查询多个对象
func ReadUserList(param ReadUserListParam) (list []*User, total int64, err error) {
	list = make([]*User, 0)
	query := orm.NewOrm().QueryTable(UserTBName())

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
func ReadUserListRaw(param ReadUserListParam) (list []*User, total int64, err error) {
	list = make([]*User, 0)
	var args = make([]interface{}, 0)
	var whereSql = "WHERE 1=1"

	if len(param.Name) > 0 {
		whereSql += " AND T0.`name` LIKE ?"
		args = append(args, fmt.Sprintf("%%%s%%", param.Name))
	}

	if param.Type != UserTypeIgnore {
		whereSql += " AND T0.`type` = ?"
		args = append(args, param.Type)
	}

	if param.Status != UserStatusIgnore {
		whereSql += " AND T0.`status` = ?"
		args = append(args, param.Status)
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
	var fields = "T0.`id`, T0.`name`" +
		// ", T0.`password`" +
		", T0.`type`, T0.`true_name`, T0.`mobile`, T0.`email`, T0.`status`, T0.`create_time`, T0.`update_time`, T0.`memo`"

	// 关联查询
	var relatedSql string

	// 连表查询
	var sql = fmt.Sprintf("SELECT %s FROM user AS T0 %s %s %s %s", fields, relatedSql, whereSql, orderSql, pageSql)

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
	var countSql = fmt.Sprintf("SELECT count(*) AS count FROM user AS T0 %s %s", relatedSql, whereSql)
	var count RawCount
	err = orm.NewOrm().Raw(countSql, args...).QueryRow(&count)
	if err != nil {
		return
	}
	total = count.Count
	return
}

// 插入单个对象
func InsertUserOne(m User) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(&m)
	return
}

// 插入多个对象
func InsertUserMulti(list []User) (num int64, err error) {
	o := orm.NewOrm()
	num, err = o.InsertMulti(100, list)
	return
}

// 更新单个对象
func UpdateUserOne(m User, fields ...string) (num int64, err error) {
	o := orm.NewOrm()
	if len(fields) == 0 {
		fields = []string{"name", "password", "type", "true_name", "mobile", "email", "status", "create_time", "update_time", "memo"}
	}
	num, err = o.Update(&m, fields...)
	return
}

// 更新多个对象
func UpdateUserMulti(ids []int, params map[string]interface{}) (num int64, err error) {
	o := orm.NewOrm()
	num, err = o.QueryTable(UserTBName()).Filter("id__in", ids).Update(params)
	return
}

// 删除单个对象
func DeleteUserOne(id int) (num int64, err error) {
	o := orm.NewOrm()
	num, err = o.Delete(&User{ID: id})
	return
}

// 删除多个对象
func DeleteUserMulti(ids []int) (num int64, err error) {
	o := orm.NewOrm()
	num, err = o.QueryTable(UserTBName()).Filter("id__in", ids).Delete()
	return
}

// 更新用户角色
func UpdateUserRoleMulti(param UpdateUserRoleParam) (err error) {
	var list = make([]*UserRoleRel, len(param.RoleList))
	for k, v := range param.RoleList {
		list[k] = &UserRoleRel{
			UserID: param.ID,
			RoleID: v,
		}
	}

	o := orm.NewOrm()
	err = o.Begin()
	if err != nil {
		return
	}

	// 删除旧关系
	_, err = o.Raw("DELETE FROM user_role_rel WHERE user_id = ?", param.ID).Exec()
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

// 删除单个对象
func DeleteUserOneWithRel(id int) (num int64, err error) {
	o := orm.NewOrm()

	err = o.Begin()
	if err != nil {
		return
	}

	// 删除用户与角色关系
	_, err = o.Raw("DELETE FROM user_role_rel WHERE user_id = ?", id).Exec()
	if err != nil {
		o.Rollback()
		return
	}

	// 删除用户与班级关系
	_, err = o.Raw("DELETE FROM class_user_rel WHERE user_id = ?", id).Exec()
	if err != nil {
		o.Rollback()
		return
	}

	// 删除用户
	_, err = o.Raw("DELETE FROM user WHERE id = ?", id).Exec()
	if err != nil {
		o.Rollback()
		return
	}

	err = o.Commit()
	return
}

// 删除多个对象
func DeleteUserMultiWithRel(ids []int) (num int64, err error) {
	o := orm.NewOrm()

	err = o.Begin()
	if err != nil {
		return
	}

	// 删除用户与角色关系
	_, err = o.Raw(fmt.Sprintf("DELETE FROM user_role_rel WHERE user_id IN (?%s)", strings.Repeat(", ?", len(ids)-1)), ids).Exec()
	if err != nil {
		o.Rollback()
		return
	}

	// 删除用户与班级关系
	_, err = o.Raw(fmt.Sprintf("DELETE FROM class_user_rel WHERE user_id IN (?%s)", strings.Repeat(", ?", len(ids)-1)), ids).Exec()
	if err != nil {
		o.Rollback()
		return
	}

	// 删除用户
	_, err = o.Raw(fmt.Sprintf("DELETE FROM user WHERE id IN (?%s)", strings.Repeat(", ?", len(ids)-1)), ids).Exec()
	if err != nil {
		o.Rollback()
		return
	}

	err = o.Commit()
	return
}
