package controllers

import (
	"github.com/astaxie/beego/logs"
	"go-exam/models"
	"time"
)

// 试卷
type PaperController struct {
	BaseController
}

// 预执行
func (c *PaperController) Prepare() {
	c.BaseController.Prepare()
}

// 查询列表
func (c *PaperController) List() {
	var param models.ReadPaperListParam
	var err error
	if err = c.ParseParam(&param); err != nil {
		logs.Info("c[Paper][List]: 参数错误, err = %s, req = %s", err.Error(), c.Ctx.Input.RequestBody)
		c.Failure("参数错误")
	}

	// 查询列表
	list, total, err := models.ReadPaperList(param)
	if err != nil {
		logs.Info("c[Paper][List]: 查询列表失败, err = %s", err.Error())
		c.Failure("获取数据失败")
	}

	var res = make(map[string]interface{})
	res["list"] = list
	res["total"] = total

	c.Success(res)
}

// 查询详情
func (c *PaperController) Detail() {
	var param models.ReadPaperDetailParam
	var err error
	if err = c.ParseParam(&param); err != nil {
		logs.Info("c[Paper][Detail]: 参数错误, err = %s, req = %s", err.Error(), c.Ctx.Input.RequestBody)
		c.Failure("参数错误")
	}

	// 查询详情
	m, err := models.ReadPaperOne(param.ID)
	if err != nil {
		logs.Info("c[Paper][Detail]: 查询详情失败, err = %s", err.Error())
		c.Failure("获取数据失败")
	}

	c.Success(m)
}

// 智能组卷
func (c *PaperController) Auto() {
	// TODO paper表新增各题分值字段

	// 获取参数（试卷参数）

	// 调用组卷算法获取试题列表

	// 保存试卷信息

	// 保存试卷试题列表（试题列表、试题选项列表）
}

// 手工组卷
func (c *PaperController) Op() {
	// TODO paper表新增各题分值字段

	// 获取参数（试卷参数、试题集合）

	// 更新试卷信息

	// 删除原有试题列表、选项列表，保存新试卷试题列表（试题列表、试题选项列表）
}

// 创建
func (c *PaperController) Create() {
	var m models.Paper
	var err error
	if err = c.ParseParam(&m); err != nil {
		logs.Info("c[Paper][Create]: 参数错误, err = %s, req = %s", err.Error(), c.Ctx.Input.RequestBody)
		c.Failure("参数错误")
	}

	// 创建
	m.CreateTime = time.Now().Unix()
	id, err := models.InsertPaperOne(m)
	if err != nil {
		logs.Info("c[Paper][Create]: 创建失败, err = %s", err.Error())
		c.Failure("操作失败")
	}

	var res = make(map[string]interface{})
	res["id"] = id

	c.Success(res)
}

// 更新
func (c *PaperController) Update() {
	var m models.Paper
	var err error
	if err = c.ParseParam(&m); err != nil {
		logs.Info("c[Paper][Update]: 参数错误, err = %s, req = %s", err.Error(), c.Ctx.Input.RequestBody)
		c.Failure("参数错误")
	}

	// 更新
	m.UpdateTime = time.Now().Unix()
	_, err = models.UpdatePaperOne(m, "name", "status", "update_time", "memo")
	if err != nil {
		logs.Info("c[Paper][Update]: 更新失败, err = %s", err.Error())
		c.Failure("操作失败")
	}

	c.Success(nil)
}

// 删除
func (c *PaperController) Delete() {
	var param models.DeletePaperParam
	var err error
	if err = c.ParseParam(&param); err != nil {
		logs.Info("c[Paper][Delete]: 参数错误, err = %s, req = %s", err.Error(), c.Ctx.Input.RequestBody)
		c.Failure("参数错误")
	}
	if param.ID < 0 && len(param.List) == 0 {
		logs.Info("c[Paper][Delete]: 参数错误, 无效id或list为空, req = %s", err.Error(), c.Ctx.Input.RequestBody)
		c.Failure("参数错误")
	}
	for _, v := range param.List {
		if v <= 0 {
			logs.Info("c[Paper][Delete]: 参数错误, 无效id, req = %s", err.Error(), c.Ctx.Input.RequestBody)
			c.Failure("参数错误")
		}
	}

	var num int64
	if param.ID > 0 {
		// 删除
		num, err = models.DeletePaperOne(param.ID)
		if err != nil {
			logs.Info("c[Paper][Delete]: 删除失败, err = %s", err.Error())
			c.Failure("操作失败")
		}
	} else if len(param.List) == 1 {
		// 删除
		num, err = models.DeletePaperOne(param.List[0])
		if err != nil {
			logs.Info("c[Paper][Delete]: 删除失败, err = %s", err.Error())
			c.Failure("操作失败")
		}
	} else {
		// 批量删除
		num, err = models.DeletePaperMulti(param.List)
		if err != nil {
			logs.Info("c[Paper][Delete]: 批量删除失败, err = %s", err.Error())
			c.Failure("操作失败")
		}
	}

	var res = make(map[string]interface{})
	res["num"] = num

	c.Success(res)
}
