package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

// 权限表
type Permission struct {
	ID               int    `orm:"column(id)" form:"id" json:"id"`
	Type             int    `orm:"column(type)" form:"type" json:"type"`
	Pid              int    `orm:"column(pid)" form:"pid" json:"pid"`
	Status           int    `orm:"column(status)" form:"status" json:"status"`
	Path             string `orm:"column(path)" form:"path" json:"path"`
	Name             string `orm:"column(name)" form:"name" json:"name"`
	Component        string `orm:"column(component)" form:"component" json:"component"`
	Redirect         string `orm:"column(redirect)" form:"redirect" json:"redirect"`
	MetaTitle        string `orm:"column(meta_title)" form:"meta_title" json:"meta_title"`
	MetaIcon         string `orm:"column(meta_icon)" form:"meta_icon" json:"meta_icon"`
	MetaExtraIcon    string `orm:"column(meta_extra_icon)" form:"meta_extra_icon" json:"meta_extra_icon"`
	MetaShowLink     int8   `orm:"column(meta_show_link)" form:"meta_show_link" json:"meta_show_link"`
	MetaShowParent   int8   `orm:"column(meta_show_parent)" form:"meta_show_parent" json:"meta_show_parent"`
	MetaKeepAlive    int8   `orm:"column(meta_keep_alive)" form:"meta_keep_alive" json:"meta_keep_alive"`
	MetaFrameSrc     string `orm:"column(meta_frame_src)" form:"meta_frame_src" json:"meta_frame_src"`
	MetaFrameLoading int8   `orm:"column(meta_frame_loading)" form:"meta_frame_loading" json:"meta_frame_loading"`
	MetaHiddenTag    int8   `orm:"column(meta_hidden_tag)" form:"meta_hidden_tag" json:"meta_hidden_tag"`
	MetaRank         int    `orm:"column(meta_rank)" form:"meta_rank" json:"meta_rank"`
	CreateTime       int64  `orm:"column(create_time)" form:"create_time" json:"create_time"`
	UpdateTime       int64  `orm:"column(update_time)" form:"update_time" json:"update_time"`
	Memo             string `orm:"column(memo)" form:"memo" json:"memo"`
}

// 类型
const (
	PERMISSION_TYPE_IGNORE = -1 // 忽略类型
	PERMISSION_MENU        = 1  // 菜单权限
	PERMISSION_PAGE        = 2  // 页面权限
	PERMISSION_COMPONENT   = 3  // 组件权限
	PERMISSION_OP          = 4  // 操作权限
	PERMISSION_BUTTON      = 5  // 按钮权限
	PERMISSION_DATA        = 6  // 数据权限
)

// 状态
const (
	PERMISSION_STATUS_IGNORE = -1 // 忽略状态
	PERMISSION_DISABLE       = 0  // 禁用
	PERMISSION_ENABLE        = 1  // 正常
)

// 是否显示
const (
	PERMISSION_META_SHOW_LINE_DISABLE = 0 // 隐藏
	PERMISSION_META_SHOW_LINE_ENABLE  = 1 // 显示
)

// 是否显示父级菜单
const (
	PERMISSION_META_SHOW_PARENT_DISABLE = 0 // 隐藏
	PERMISSION_META_SHOW_PARENT_ENABLE  = 1 // 显示
)

// 是否缓存
const (
	PERMISSION_META_KEEP_ALIVE_DISABLE = 0 // 关闭
	PERMISSION_META_KEEP_ALIVE_ENABLE  = 1 // 开启
)

// 是否开启iframe首次加载动画
const (
	PERMISSION_META_IFRAME_LOADING_DISABLE = 0 // 关闭
	PERMISSION_META_IFRAME_LOADING_ENABLE  = 1 // 开启
)

// 是否不添加信息到标签页
const (
	PERMISSION_META_HIDDEN_TAG_DISABLE = 0 // 添加
	PERMISSION_META_HIDDEN_TAG_ENABLE  = 1 // 不添加
)

// 查询详情参数
type ReadPermissionDetailParam struct {
	ID int `json:"id"`
}

// 查询列表参数
type ReadPermissionListParam struct {
	BaseQueryParam
	Type      int    `json:"type"`
	Status    int    `json:"status"`
	Path      string `json:"path"`
	Name      string `json:"name"`
	MetaTitle string `json:"meta_title"`
	ClosePage bool   `form:"close_page" json:"close_page"`
}

// 删除参数
type DeletePermissionDetailParam struct {
	ID   int   `json:"id"`
	List []int `json:"list"`
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
	return [][]string{
		{"name"},
	}
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

	if param.Type != PERMISSION_TYPE_IGNORE {
		query = query.Filter("type", param.Type)
	}

	if param.Status != PERMISSION_STATUS_IGNORE {
		query = query.Filter("status", param.Status)
	}

	if len(param.Path) > 0 {
		query = query.Filter("path__icontains", param.Path)
	}

	if len(param.Name) > 0 {
		query = query.Filter("name__icontains", param.Name)
	}

	if len(param.MetaTitle) > 0 {
		query = query.Filter("meta_title__icontains", param.MetaTitle)
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
	var fields = "T0.`id`, T0.`type`, T0.`pid`, T0.`status`, T0.`path`, T0.`name`, T0.`component`, T0.`redirect`, T0.`meta_title`, T0.`meta_icon`, T0.`meta_extra_icon`, T0.`meta_show_link`, T0.`meta_show_parent`, T0.`meta_keep_alive`, T0.`meta_frame_src`, T0.`meta_frame_loading`, T0.`meta_hidden_tag`, T0.`meta_rank`, T0.`create_time`, T0.`update_time`, T0.`memo`"

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
		fields = []string{"type", "pid", "status", "path", "name", "component", "redirect", "meta_title", "meta_icon", "meta_extra_icon", "meta_show_link", "meta_show_parent", "meta_keep_alive", "meta_frame_src", "meta_frame_loading", "meta_hidden_tag", "meta_rank", "create_time", "update_time", "memo"}
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
