package controllers

import (
	"github.com/astaxie/beego/logs"
	"go-exam/models"
	"go-exam/utils"
	"time"
)

// 用户
type UserController struct {
	BaseController
}

// 预执行
func (c *UserController) Prepare() {
	c.BaseController.Prepare()
}

// 查询列表
func (c *UserController) List() {
	var param models.ReadUserListParam
	var err error
	if err = c.ParseParam(&param); err != nil {
		logs.Info("c[User][List]: 参数错误, err = %s, req = %s", err.Error(), c.Ctx.Input.RequestBody)
		c.Failure("参数错误")
	}

	// 查询列表
	list, total, err := models.ReadUserList(param)
	if err != nil {
		logs.Info("c[User][List]: 查询列表失败, err = %s", err.Error())
		c.Failure("获取数据失败")
	}

	var res = make(map[string]interface{})
	res["list"] = list
	res["total"] = total

	c.Success(res)
}

// 查询详情
func (c *UserController) Detail() {
	var param models.ReadUserDetailParam
	var err error
	if err = c.ParseParam(&param); err != nil {
		logs.Info("c[User][Detail]: 参数错误, err = %s, req = %s", err.Error(), c.Ctx.Input.RequestBody)
		c.Failure("参数错误")
	}

	// 查询详情
	m, err := models.ReadUserOne(param.ID)
	if err != nil {
		logs.Info("c[User][Detail]: 查询详情失败, err = %s", err.Error())
		c.Failure("获取数据失败")
	}

	c.Success(m)
}

// 创建
func (c *UserController) Create() {
	var m models.User
	var err error
	if err = c.ParseParam(&m); err != nil {
		logs.Info("c[User][Create]: 参数错误, err = %s, req = %s", err.Error(), c.Ctx.Input.RequestBody)
		c.Failure("参数错误")
	}

	// 创建
	m.CreateTime = time.Now().Unix()
	id, err := models.InsertUserOne(m)
	if utils.IsUniqueConstraintError(err) {
		c.Failure("用户名已被占用")
	}
	if err != nil {
		logs.Info("c[User][Create]: 创建失败, err = %s", err.Error())
		c.Failure("操作失败")
	}

	var res = make(map[string]interface{})
	res["id"] = id

	c.Success(res)
}

// 更新
func (c *UserController) Update() {
	var m models.User
	var err error
	if err = c.ParseParam(&m); err != nil {
		logs.Info("c[User][Update]: 参数错误, err = %s, req = %s", err.Error(), c.Ctx.Input.RequestBody)
		c.Failure("参数错误")
	}

	// 更新
	m.UpdateTime = time.Now().Unix()
	_, err = models.UpdateUserOne(m, "name", "true_name", "mobile", "email", "status", "update_time", "memo")
	if utils.IsUniqueConstraintError(err) {
		c.Failure("用户名已被占用")
	}
	if err != nil {
		logs.Info("c[User][Update]: 更新失败, err = %s", err.Error())
		c.Failure("操作失败")
	}

	c.Success(nil)
}

// 删除
func (c *UserController) Delete() {
	var param models.DeleteUserParam
	var err error
	if err = c.ParseParam(&param); err != nil {
		logs.Info("c[User][Delete]: 参数错误, err = %s, req = %s", err.Error(), c.Ctx.Input.RequestBody)
		c.Failure("参数错误")
	}
	if param.ID < 0 && len(param.List) == 0 {
		logs.Info("c[User][Delete]: 参数错误, 无效id或list为空, req = %s", err.Error(), c.Ctx.Input.RequestBody)
		c.Failure("参数错误")
	}
	for _, v := range param.List {
		if v <= 0 {
			logs.Info("c[User][Delete]: 参数错误, 无效id, req = %s", err.Error(), c.Ctx.Input.RequestBody)
			c.Failure("参数错误")
		}
	}

	var num int64
	if param.ID > 0 {
		// 删除
		num, err = models.DeleteUserOne(param.ID)
		if err != nil {
			logs.Info("c[User][Delete]: 删除失败, err = %s", err.Error())
			c.Failure("操作失败")
		}
	} else if len(param.List) == 1 {
		// 删除
		num, err = models.DeleteUserOne(param.List[0])
		if err != nil {
			logs.Info("c[User][Delete]: 删除失败, err = %s", err.Error())
			c.Failure("操作失败")
		}
	} else {
		// 批量删除
		num, err = models.DeleteUserMulti(param.List)
		if err != nil {
			logs.Info("c[User][Delete]: 批量删除失败, err = %s", err.Error())
			c.Failure("操作失败")
		}
	}

	var res = make(map[string]interface{})
	res["num"] = num

	c.Success(res)
}

// 更新账号类型
func (c *UserController) UpdateType() {
	var m models.User
	var err error
	if err = c.ParseParam(&m); err != nil {
		logs.Info("c[User][UpdateType]: 参数错误, err = %s, req = %s", err.Error(), c.Ctx.Input.RequestBody)
		c.Failure("参数错误")
	}

	// 更新
	m.UpdateTime = time.Now().Unix()
	_, err = models.UpdateUserOne(m, "type", "update_time", "memo")
	if err != nil {
		logs.Info("c[User][UpdateType]: 更新失败, err = %s", err.Error())
		c.Failure("操作失败")
	}

	c.Success(nil)
}

// 更新密码
func (c *UserController) UpdatePassword() {
	var m models.User
	var err error
	if err = c.ParseParam(&m); err != nil {
		logs.Info("c[User][UpdatePassword]: 参数错误, err = %s, req = %s", err.Error(), c.Ctx.Input.RequestBody)
		c.Failure("参数错误")
	}

	// 更新
	m.UpdateTime = time.Now().Unix()
	_, err = models.UpdateUserOne(m, "password", "update_time", "memo")
	if err != nil {
		logs.Info("c[User][UpdatePassword]: 更新失败, err = %s", err.Error())
		c.Failure("操作失败")
	}

	c.Success(nil)
}

// 查询授权列表
func (c *UserController) RoleList() {
	var param models.ReadUserRoleRelListParam
	var err error
	if err = c.ParseParam(&param); err != nil {
		logs.Info("c[User][RoleList]: 参数错误, err = %s, req = %s", err.Error(), c.Ctx.Input.RequestBody)
		c.Failure("参数错误")
	}
	if param.UserID <= 0 {
		logs.Info("c[User][RoleList]: 参数错误, 用户id不能为空, req = %s", c.Ctx.Input.RequestBody)
		c.Failure("参数错误")
	}

	// 查询列表
	param.ClosePage = true
	list, total, err := models.ReadUserRoleRelListRaw(param)
	if err != nil {
		logs.Info("c[User][RoleList]: 查询列表失败, err = %s", err.Error())
		c.Failure("获取数据失败")
	}

	var res = make(map[string]interface{})
	res["list"] = list
	res["total"] = total

	c.Success(res)
}

// 更新角色
func (c *UserController) AuthRole() {
	var param models.UpdateUserRoleParam
	var err error
	if err = c.ParseParam(&param); err != nil {
		logs.Info("c[User][AuthRole]: 参数错误, err = %s, req = %s", err.Error(), c.Ctx.Input.RequestBody)
		c.Failure("参数错误")
	}

	// 更新
	err = models.UpdateUserRoleMulti(param)
	if err != nil {
		logs.Info("c[User][AuthRole]: 更新失败, err = %s", err.Error())
		c.Failure("操作失败")
	}

	c.Success(nil)
}
