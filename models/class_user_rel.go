package models

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/orm"
)

// 班级考生表
type ClassUserRel struct {
	ID      int `orm:"column(id)" form:"id" json:"id"`
	ClassID int `orm:"column(class_id)" form:"class_id" json:"class_id"`
	UserID  int `orm:"column(user_id)" form:"user_id" json:"user_id"`
}

// 查询详情参数
type ReadClassUserRelDetailParam struct {
	ID int `json:"id"`
}

// 查询列表参数
type ReadClassUserRelListParam struct {
	BaseQueryParam
	ClosePage bool `form:"close_page" json:"close_page"`
}

// 创建参数
type InsertClassUserRelParam struct {
	ClassID  int   `json:"class_id"`
	UserList []int `json:"user_list"`
}

// 删除参数
type DeleteClassUserRelParam struct {
	ID   int   `json:"id"`
	List []int `json:"list"`
}

// 初始化
func init() {
	orm.RegisterModel(new(ClassUserRel))
}

// 表名
func ClassUserRelTBName() (name string) {
	return "class_user_rel"
}

// 自定义表名
func (m *ClassUserRel) TableName() (name string) {
	return ClassUserRelTBName()
}

// 多字段索引
func (m *ClassUserRel) TableIndex() [][]string {
	return [][]string{}
}

// 多字段唯一键
func (m *ClassUserRel) TableUnique() [][]string {
	return [][]string{
		{"user_id", "class_id"},
	}
}

// 自定义引擎
func (m *ClassUserRel) TableEngine() string {
	return "INNODB"
}

// 查询单个对象
func ReadClassUserRelOne(id int) (m ClassUserRel, err error) {
	o := orm.NewOrm()
	m.ID = id
	err = o.Read(&m)
	return
}

// 查询多个对象
func ReadClassUserRelList(param ReadClassUserRelListParam) (list []*ClassUserRel, total int64, err error) {
	list = make([]*ClassUserRel, 0)
	query := orm.NewOrm().QueryTable(ClassUserRelTBName())

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
func ReadClassUserRelListRaw(param ReadClassUserRelListParam) (list []*ClassUserRel, total int64, err error) {
	list = make([]*ClassUserRel, 0)
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
	var fields = "T0.`id`, T0.`class_id`, T0.`user_id`"

	// 关联查询
	var relatedSql string

	// 连表查询
	var sql = fmt.Sprintf("SELECT %s FROM class_user_rel AS T0 %s %s %s %s", fields, relatedSql, whereSql, orderSql, pageSql)

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
	var countSql = fmt.Sprintf("SELECT count(*) AS count FROM class_user_rel AS T0 %s %s", relatedSql, whereSql)
	var count RawCount
	err = orm.NewOrm().Raw(countSql, args...).QueryRow(&count)
	if err != nil {
		return
	}
	total = count.Count
	return
}

// 插入单个对象
func InsertClassUserRelOne(m ClassUserRel) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(&m)
	return
}

// 插入多个对象
func InsertClassUserRelMulti(list []ClassUserRel) (num int64, err error) {
	o := orm.NewOrm()
	num, err = o.InsertMulti(100, list)
	return
}

// 更新单个对象
func UpdateClassUserRelOne(m ClassUserRel, fields ...string) (num int64, err error) {
	o := orm.NewOrm()
	if len(fields) == 0 {
		fields = []string{"class_id", "user_id"}
	}
	num, err = o.Update(&m, fields...)
	return
}

// 更新多个对象
func UpdateClassUserRelMulti(ids []int, params map[string]interface{}) (num int64, err error) {
	o := orm.NewOrm()
	num, err = o.QueryTable(ClassUserRelTBName()).Filter("id__in", ids).Update(params)
	return
}

// 删除单个对象
func DeleteClassUserRelOne(id int) (num int64, err error) {
	o := orm.NewOrm()
	num, err = o.Delete(&ClassUserRel{ID: id})
	return
}

// 删除多个对象
func DeleteClassUserRelMulti(ids []int) (num int64, err error) {
	o := orm.NewOrm()
	num, err = o.QueryTable(ClassUserRelTBName()).Filter("id__in", ids).Delete()
	return
}

// 插入或更新多个对象
func InsertOrUpdateClassUserRelMulti(param InsertClassUserRelParam) (num int64, err error) {
	var m = make(map[int]struct{})
	for _, v := range param.UserList {
		m[v] = struct{}{}
	}
	var list = make([]ClassUserRel, 0)
	for k := range m {
		list = append(list, ClassUserRel{ClassID: param.ClassID, UserID: k})
	}
	if len(list) == 0 {
		return 0, errors.New("参数错误: 用户列表不能为空")
	}
	num = int64(len(list))

	o := orm.NewOrm()
	err = o.Begin()
	if err != nil {
		return
	}

	for _, v := range list {
		_, err = o.InsertOrUpdate(&v)
		if err != nil {
			o.Rollback()
			return
		}
	}

	err = o.Commit()
	return
}
