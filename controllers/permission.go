package controllers

import (
	"github.com/astaxie/beego/logs"
	"go-exam/models"
	"time"
)

// 权限
type PermissionController struct {
	BaseController
}

// 预执行
func (c *PermissionController) Prepare() {
	c.BaseController.Prepare()
}

// 查询列表
func (c *PermissionController) List() {
	var param models.ReadPermissionListParam
	var err error
	if err = c.ParseParam(&param); err != nil {
		logs.Info("c[permission][list]: 参数错误, err = %s, req = %s", err.Error(), c.Ctx.Input.RequestBody)
		c.Failure("参数错误")
	}

	// 查询列表
	list, total, err := models.ReadPermissionList(param)
	if err != nil {
		logs.Info("c[permission][list]: 查询列表失败, err = %s", err.Error())
		c.Failure("获取数据失败")
	}

	var res = make(map[string]interface{})
	res["list"] = list
	res["total"] = total

	c.Success(res)
}

// 查询列表
func (c *PermissionController) All() {
	var param models.ReadPermissionListParam
	var err error
	if err = c.ParseParam(&param); err != nil {
		logs.Info("c[permission][list]: 参数错误, err = %s, req = %s", err.Error(), c.Ctx.Input.RequestBody)
		c.Failure("参数错误")
	}

	// 查询列表
	param.ClosePage = true
	list, total, err := models.ReadPermissionListRaw(param)
	if err != nil {
		logs.Info("c[permission][list]: 查询列表失败, err = %s", err.Error())
		c.Failure("获取数据失败")
	}

	var res = make(map[string]interface{})
	res["list"] = list
	res["total"] = total

	c.Success(res)
}

// 查询详情
func (c *PermissionController) Detail() {
	var param models.ReadPermissionDetailParam
	var err error
	if err = c.ParseParam(&param); err != nil {
		logs.Info("c[permission][detail]: 参数错误, err = %s, req = %s", err.Error(), c.Ctx.Input.RequestBody)
		c.Failure("参数错误")
	}

	// 查询详情
	m, err := models.ReadPermissionOne(param.ID)
	if err != nil {
		logs.Info("c[permission][detail]: 查询详情失败, err = %s", err.Error())
		c.Failure("获取数据失败")
	}

	c.Success(m)
}

// 创建
func (c *PermissionController) Create() {
	var m models.Permission
	var err error
	if err = c.ParseParam(&m); err != nil {
		logs.Info("c[permission][create]: 参数错误, err = %s, req = %s", err.Error(), c.Ctx.Input.RequestBody)
		c.Failure("参数错误")
	}

	// 创建
	m.CreateTime = time.Now().Unix()
	id, err := models.InsertPermissionOne(m)
	if err != nil {
		logs.Info("c[permission][create]: 创建失败, err = %s", err.Error())
		c.Failure("操作失败")
	}

	var res = make(map[string]interface{})
	res["id"] = id

	c.Success(res)
}

// 更新
func (c *PermissionController) Update() {
	var m models.Permission
	var err error
	if err = c.ParseParam(&m); err != nil {
		logs.Info("c[permission][update]: 参数错误, err = %s, req = %s", err.Error(), c.Ctx.Input.RequestBody)
		c.Failure("参数错误")
	}

	// 更新
	m.UpdateTime = time.Now().Unix()
	_, err = models.UpdatePermissionOne(m, "type", "pid", "code", "status", "path", "name", "component", "redirect", "meta_title", "meta_icon", "meta_extra_icon", "meta_show_link", "meta_show_parent", "meta_keep_alive", "meta_frame_src", "meta_frame_loading", "meta_hidden_tag", "meta_rank", "update_time", "memo")
	if err != nil {
		logs.Info("c[permission][update]: 更新失败, err = %s", err.Error())
		c.Failure("操作失败")
	}

	c.Success(nil)
}

// 删除
func (c *PermissionController) Delete() {
	var param models.DeletePermissionParam
	var err error
	if err = c.ParseParam(&param); err != nil {
		logs.Info("c[permission][delete]: 参数错误, err = %s, req = %s", err.Error(), c.Ctx.Input.RequestBody)
		c.Failure("参数错误")
	}
	if param.ID < 0 && len(param.List) == 0 {
		logs.Info("c[permission][delete]: 参数错误, 无效id或list为空, req = %s", err.Error(), c.Ctx.Input.RequestBody)
		c.Failure("参数错误")
	}
	for _, v := range param.List {
		if v <= 0 {
			logs.Info("c[permission][delete]: 参数错误, 无效id, req = %s", err.Error(), c.Ctx.Input.RequestBody)
			c.Failure("参数错误")
		}
	}

	var num int64
	if param.ID > 0 {
		// 删除
		num, err = models.DeletePermissionMultiWithChildren([]int{param.ID})
		if err != nil {
			logs.Info("c[permission][delete]: 删除失败, err = %s", err.Error())
			c.Failure("操作失败")
		}
	} else {
		// 批量删除
		num, err = models.DeletePermissionMultiWithChildren(param.List)
		if err != nil {
			logs.Info("c[permission][delete]: 批量删除失败, err = %s", err.Error())
			c.Failure("操作失败")
		}
	}

	var res = make(map[string]interface{})
	res["num"] = num

	c.Success(res)
}
