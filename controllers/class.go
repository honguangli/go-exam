package controllers

import (
	"github.com/astaxie/beego/logs"
	"go-exam/models"
)

// 班级
type ClassController struct {
	BaseController
}

// 预执行
func (c *ClassController) Prepare() {
	c.BaseController.Prepare()
}

// 查询列表
func (c *ClassController) List() {
	var param models.ReadClassListParam
	var err error
	if err = c.ParseParam(&param); err != nil {
		logs.Info("c[Class][List]: 参数错误, err = %s, req = %s", err.Error(), c.Ctx.Input.RequestBody)
		c.Failure("参数错误")
	}

	// 查询列表
	list, total, err := models.ReadClassList(param)
	if err != nil {
		logs.Info("c[Class][List]: 查询列表失败, err = %s", err.Error())
		c.Failure("获取数据失败")
	}

	var res = make(map[string]interface{})
	res["list"] = list
	res["total"] = total

	c.Success(res)
}

// 查询详情
func (c *ClassController) Detail() {
	var param models.ReadClassDetailParam
	var err error
	if err = c.ParseParam(&param); err != nil {
		logs.Info("c[Class][Detail]: 参数错误, err = %s, req = %s", err.Error(), c.Ctx.Input.RequestBody)
		c.Failure("参数错误")
	}

	// 查询详情
	m, err := models.ReadClassOne(param.ID)
	if err != nil {
		logs.Info("c[Class][Detail]: 查询详情失败, err = %s", err.Error())
		c.Failure("获取数据失败")
	}

	c.Success(m)
}

// 创建
func (c *ClassController) Create() {
	var m models.Class
	var err error
	if err = c.ParseParam(&m); err != nil {
		logs.Info("c[Class][Create]: 参数错误, err = %s, req = %s", err.Error(), c.Ctx.Input.RequestBody)
		c.Failure("参数错误")
	}

	// 创建
	id, err := models.InsertClassOne(m)
	if err != nil {
		logs.Info("c[Class][Create]: 创建失败, err = %s", err.Error())
		c.Failure("获取数据失败")
	}

	var res = make(map[string]interface{})
	res["id"] = id

	c.Success(res)
}

// 更新
func (c *ClassController) Update() {
	var m models.Class
	var err error
	if err = c.ParseParam(&m); err != nil {
		logs.Info("c[Class][Update]: 参数错误, err = %s, req = %s", err.Error(), c.Ctx.Input.RequestBody)
		c.Failure("参数错误")
	}

	// 更新
	_, err = models.UpdateClassOne(m, "name", "status", "desc")
	if err != nil {
		logs.Info("c[Class][Update]: 更新失败, err = %s", err.Error())
		c.Failure("操作失败")
	}

	c.Success(nil)
}

// 删除
func (c *ClassController) Delete() {
	var param models.DeleteClassParam
	var err error
	if err = c.ParseParam(&param); err != nil {
		logs.Info("c[Class][Delete]: 参数错误, err = %s, req = %s", err.Error(), c.Ctx.Input.RequestBody)
		c.Failure("参数错误")
	}
	if param.ID < 0 && len(param.List) == 0 {
		logs.Info("c[Class][Delete]: 参数错误, 无效id或list为空, req = %s", err.Error(), c.Ctx.Input.RequestBody)
		c.Failure("参数错误")
	}
	for _, v := range param.List {
		if v <= 0 {
			logs.Info("c[Class][Delete]: 参数错误, 无效id, req = %s", err.Error(), c.Ctx.Input.RequestBody)
			c.Failure("参数错误")
		}
	}

	var num int64
	if param.ID > 0 {
		// 删除
		num, err = models.DeleteClassOne(param.ID)
		if err != nil {
			logs.Info("c[Class][Delete]: 删除失败, err = %s", err.Error())
			c.Failure("操作失败")
		}
	} else if len(param.List) == 1 {
		// 删除
		num, err = models.DeleteClassOne(param.List[0])
		if err != nil {
			logs.Info("c[Class][Delete]: 删除失败, err = %s", err.Error())
			c.Failure("操作失败")
		}
	} else {
		// 批量删除
		num, err = models.DeleteClassMulti(param.List)
		if err != nil {
			logs.Info("c[Class][Delete]: 批量删除失败, err = %s", err.Error())
			c.Failure("操作失败")
		}
	}

	var res = make(map[string]interface{})
	res["num"] = num

	c.Success(res)
}

// 查询班级用户列表
func (c *ClassController) UserList() {
	var param models.ReadClassUserRelModelListParam
	var err error
	if err = c.ParseParam(&param); err != nil {
		logs.Info("c[Class][UserList]: 参数错误, err = %s, req = %s", err.Error(), c.Ctx.Input.RequestBody)
		c.Failure("参数错误")
	}

	// 查询列表
	list, total, err := models.ReadClassUserRelModelListRaw(param)
	if err != nil {
		logs.Info("c[Class][UserList]: 查询列表失败, err = %s", err.Error())
		c.Failure("获取数据失败")
	}

	var res = make(map[string]interface{})
	res["list"] = list
	res["total"] = total

	c.Success(res)
}

// 添加用户
func (c *ClassController) PushUser() {
	var param models.InsertClassUserRelParam
	var err error
	if err = c.ParseParam(&param); err != nil {
		logs.Info("c[Class][PushUser]: 参数错误, err = %s, req = %s", err.Error(), c.Ctx.Input.RequestBody)
		c.Failure("参数错误")
	}

	// 添加用户
	num, err := models.InsertOrUpdateClassUserRelMulti(param)
	if err != nil {
		logs.Info("c[Class][PushUser]: 更新失败, err = %s", err.Error())
		c.Failure("操作失败")
	}

	var res = make(map[string]interface{})
	res["num"] = num

	c.Success(res)
}

// 删除用户
func (c *ClassController) DeleteUser() {
	var param models.DeleteClassUserRelParam
	var err error
	if err = c.ParseParam(&param); err != nil {
		logs.Info("c[Class][DeleteUser]: 参数错误, err = %s, req = %s", err.Error(), c.Ctx.Input.RequestBody)
		c.Failure("参数错误")
	}
	if param.ID < 0 && len(param.List) == 0 {
		logs.Info("c[Class][DeleteUser]: 参数错误, 无效id或list为空, req = %s", err.Error(), c.Ctx.Input.RequestBody)
		c.Failure("参数错误")
	}
	for _, v := range param.List {
		if v <= 0 {
			logs.Info("c[Class][DeleteUser]: 参数错误, 无效id, req = %s", err.Error(), c.Ctx.Input.RequestBody)
			c.Failure("参数错误")
		}
	}

	var num int64
	if param.ID > 0 {
		// 删除
		num, err = models.DeleteClassUserRelOne(param.ID)
		if err != nil {
			logs.Info("c[Class][DeleteUser]: 删除失败, err = %s", err.Error())
			c.Failure("操作失败")
		}
	} else if len(param.List) == 1 {
		// 删除
		num, err = models.DeleteClassUserRelOne(param.List[0])
		if err != nil {
			logs.Info("c[Class][DeleteUser]: 删除失败, err = %s", err.Error())
			c.Failure("操作失败")
		}
	} else {
		// 批量删除
		num, err = models.DeleteClassUserRelMulti(param.List)
		if err != nil {
			logs.Info("c[Class][DeleteUser]: 批量删除失败, err = %s", err.Error())
			c.Failure("操作失败")
		}
	}

	var res = make(map[string]interface{})
	res["num"] = num

	c.Success(res)
}
