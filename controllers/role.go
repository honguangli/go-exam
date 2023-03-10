package controllers

import (
	"github.com/astaxie/beego/logs"
	"go-exam/models"
	"go-exam/utils"
	"time"
)

// 角色
type RoleController struct {
	BaseController
}

// 预执行
func (c *RoleController) Prepare() {
	c.BaseController.Prepare()
}

// 查询列表
func (c *RoleController) List() {
	var param models.ReadRoleListParam
	var err error
	if err = c.ParseParam(&param); err != nil {
		logs.Info("c[Role][List]: 参数错误, err = %s, req = %s", err.Error(), c.Ctx.Input.RequestBody)
		c.Failure("参数错误")
	}

	// 查询列表
	list, total, err := models.ReadRoleList(param)
	if err != nil {
		logs.Info("c[Role][List]: 查询列表失败, err = %s", err.Error())
		c.Failure("获取数据失败")
	}

	var res = make(map[string]interface{})
	res["list"] = list
	res["total"] = total

	c.Success(res)
}

// 查询列表
func (c *RoleController) All() {
	var param models.ReadRoleListParam
	var err error
	if err = c.ParseParam(&param); err != nil {
		logs.Info("c[Role][All]: 参数错误, err = %s, req = %s", err.Error(), c.Ctx.Input.RequestBody)
		c.Failure("参数错误")
	}

	// 查询列表
	param.ClosePage = true
	param.Sort = "seq"
	param.Order = "asc"
	list, total, err := models.ReadRoleListRaw(param)
	if err != nil {
		logs.Info("c[Role][All]: 查询列表失败, err = %s", err.Error())
		c.Failure("获取数据失败")
	}

	var res = make(map[string]interface{})
	res["list"] = list
	res["total"] = total

	c.Success(res)
}

// 查询详情
func (c *RoleController) Detail() {
	var param models.ReadRoleDetailParam
	var err error
	if err = c.ParseParam(&param); err != nil {
		logs.Info("c[Role][Detail]: 参数错误, err = %s, req = %s", err.Error(), c.Ctx.Input.RequestBody)
		c.Failure("参数错误")
	}

	// 查询详情
	m, err := models.ReadRoleOne(param.ID)
	if err != nil {
		logs.Info("c[Role][Detail]: 查询详情失败, err = %s", err.Error())
		c.Failure("获取数据失败")
	}

	c.Success(m)
}

// 创建
func (c *RoleController) Create() {
	var m models.Role
	var err error
	if err = c.ParseParam(&m); err != nil {
		logs.Info("c[Role][Create]: 参数错误, err = %s, req = %s", err.Error(), c.Ctx.Input.RequestBody)
		c.Failure("参数错误")
	}

	// 创建
	m.CreateTime = time.Now().Unix()
	id, err := models.InsertRoleOne(m)
	if utils.IsUniqueConstraintError(err) {
		c.Failure("角色代码已被占用")
	}
	if err != nil {
		logs.Info("c[Role][Create]: 创建失败, err = %s", err.Error())
		c.Failure("操作失败")
	}

	var res = make(map[string]interface{})
	res["id"] = id

	c.Success(res)
}

// 更新
func (c *RoleController) Update() {
	var m models.Role
	var err error
	if err = c.ParseParam(&m); err != nil {
		logs.Info("c[Role][Update]: 参数错误, err = %s, req = %s", err.Error(), c.Ctx.Input.RequestBody)
		c.Failure("参数错误")
	}

	// 更新
	m.UpdateTime = time.Now().Unix()
	_, err = models.UpdateRoleOne(m, "name", "code", "seq", "status", "update_time", "memo")
	if utils.IsUniqueConstraintError(err) {
		c.Failure("角色代码已被占用")
	}
	if err != nil {
		logs.Info("c[Role][Update]: 更新失败, err = %s", err.Error())
		c.Failure("操作失败")
	}

	c.Success(nil)
}

// 删除
func (c *RoleController) Delete() {
	var param models.DeleteRoleParam
	var err error
	if err = c.ParseParam(&param); err != nil {
		logs.Info("c[Role][Delete]: 参数错误, err = %s, req = %s", err.Error(), c.Ctx.Input.RequestBody)
		c.Failure("参数错误")
	}
	if param.ID < 0 && len(param.List) == 0 {
		logs.Info("c[Role][Delete]: 参数错误, 无效id或list为空, req = %s", err.Error(), c.Ctx.Input.RequestBody)
		c.Failure("参数错误")
	}
	for _, v := range param.List {
		if v <= 0 {
			logs.Info("c[Role][Delete]: 参数错误, 无效id, req = %s", err.Error(), c.Ctx.Input.RequestBody)
			c.Failure("参数错误")
		}
	}

	var num int64
	if param.ID > 0 {
		// 删除
		num, err = models.DeleteRoleOneWithRel(param.ID)
		if err != nil {
			logs.Info("c[Role][Delete]: 删除失败, err = %s", err.Error())
			c.Failure("操作失败")
		}
	} else if len(param.List) == 1 {
		// 删除
		num, err = models.DeleteRoleOneWithRel(param.List[0])
		if err != nil {
			logs.Info("c[Role][Delete]: 删除失败, err = %s", err.Error())
			c.Failure("操作失败")
		}
	} else {
		// 批量删除
		num, err = models.DeleteRoleMultiWithRel(param.List)
		if err != nil {
			logs.Info("c[Role][Delete]: 批量删除失败, err = %s", err.Error())
			c.Failure("操作失败")
		}
	}

	var res = make(map[string]interface{})
	res["num"] = num

	c.Success(res)
}

// 查询授权列表
func (c *RoleController) PermissionList() {
	var param models.ReadRolePermissionRelListParam
	var err error
	if err = c.ParseParam(&param); err != nil {
		logs.Info("c[Role][PermissionList]: 参数错误, err = %s, req = %s", err.Error(), c.Ctx.Input.RequestBody)
		c.Failure("参数错误")
	}
	if param.RoleID <= 0 {
		logs.Info("c[Role][PermissionList]: 参数错误, 角色id不能为空, req = %s", c.Ctx.Input.RequestBody)
		c.Failure("参数错误")
	}

	// 查询列表
	param.ClosePage = true
	list, total, err := models.ReadRolePermissionRelListRaw(param)
	if err != nil {
		logs.Info("c[Role][PermissionList]: 查询列表失败, err = %s", err.Error())
		c.Failure("获取数据失败")
	}

	var res = make(map[string]interface{})
	res["list"] = list
	res["total"] = total

	c.Success(res)
}

// 授权
func (c *RoleController) AuthPermission() {
	var param models.UpdateRolePermissionParam
	var err error
	if err = c.ParseParam(&param); err != nil {
		logs.Info("c[Role][AuthPermission]: 参数错误, err = %s, req = %s", err.Error(), c.Ctx.Input.RequestBody)
		c.Failure("参数错误")
	}

	// 更新
	err = models.UpdateRolePermissionMulti(param)
	if err != nil {
		logs.Info("c[Role][AuthPermission]: 更新失败, err = %s", err.Error())
		c.Failure("操作失败")
	}

	c.Success(nil)
}
