package controllers

import (
	"github.com/astaxie/beego/logs"
	"go-exam/models"
	"time"
)

// 考试计划
type PlanController struct {
	BaseController
}

// 预执行
func (c *PlanController) Prepare() {
	c.BaseController.Prepare()
}

// 查询列表
func (c *PlanController) List() {
	var param models.ReadPlanListParam
	var err error
	if err = c.ParseParam(&param); err != nil {
		logs.Info("c[Plan][List]: 参数错误, err = %s, req = %s", err.Error(), c.Ctx.Input.RequestBody)
		c.Failure("参数错误")
	}

	// 查询列表
	list, total, err := models.ReadPlanRelListRaw(param)
	if err != nil {
		logs.Info("c[Plan][List]: 查询列表失败, err = %s", err.Error())
		c.Failure("获取数据失败")
	}

	var res = make(map[string]interface{})
	res["list"] = list
	res["total"] = total

	c.Success(res)
}

// 查询详情
func (c *PlanController) Detail() {
	var param models.ReadPlanDetailParam
	var err error
	if err = c.ParseParam(&param); err != nil {
		logs.Info("c[Plan][Detail]: 参数错误, err = %s, req = %s", err.Error(), c.Ctx.Input.RequestBody)
		c.Failure("参数错误")
	}

	// 查询详情
	m, err := models.ReadPlanRelOne(param.ID)
	if err != nil {
		logs.Info("c[Plan][Detail]: 查询详情失败, err = %s", err.Error())
		c.Failure("获取数据失败")
	}

	c.Success(m)
}

// 创建
func (c *PlanController) Create() {
	var m models.Plan
	var err error
	if err = c.ParseParam(&m); err != nil {
		logs.Info("c[Plan][Create]: 参数错误, err = %s, req = %s", err.Error(), c.Ctx.Input.RequestBody)
		c.Failure("参数错误")
	}

	// 创建
	m.CreateTime = time.Now().Unix()
	id, err := models.InsertPlanOne(m)
	if err != nil {
		logs.Info("c[Plan][Create]: 创建失败, err = %s", err.Error())
		c.Failure("操作失败")
	}

	var res = make(map[string]interface{})
	res["id"] = id

	c.Success(res)
}

// 更新
func (c *PlanController) Update() {
	var m models.Plan
	var err error
	if err = c.ParseParam(&m); err != nil {
		logs.Info("c[Plan][Update]: 参数错误, err = %s, req = %s", err.Error(), c.Ctx.Input.RequestBody)
		c.Failure("参数错误")
	}

	// 查询详情
	om, err := models.ReadPlanOne(m.ID)
	if err != nil {
		logs.Info("c[Plan][Update]: 查询详情失败, err = %s, req = %s", err.Error(), c.Ctx.Input.RequestBody)
		c.Failure("参数错误")
	}
	if om.Status != models.PlanDefault {
		c.Failure("操作失败")
	}

	// 更新
	m.UpdateTime = time.Now().Unix()
	_, err = models.UpdatePlanOne(m, "name", "paper_id", "start_time", "end_time", "duration", "update_time", "memo")
	if err != nil {
		logs.Info("c[Plan][Update]: 更新失败, err = %s", err.Error())
		c.Failure("操作失败")
	}

	c.Success(nil)
}

// 发布
func (c *PlanController) Publish() {
	var param models.PublishPlanParam
	var err error
	if err = c.ParseParam(&param); err != nil {
		logs.Info("c[Plan][Publish]: 参数错误, err = %s, req = %s", err.Error(), c.Ctx.Input.RequestBody)
		c.Failure("参数错误")
	}
	if param.ID <= 0 {
		logs.Info("c[Plan][Publish]: 参数错误, id非法, req = %s", c.Ctx.Input.RequestBody)
		c.Failure("参数错误")
	}

	// 发布
	err = models.PublishPlan(param)
	if err == models.ErrStudentsEmpty {
		c.Failure(err.Error())
	}
	if err != nil {
		logs.Info("c[Plan][Publish]: 更新失败, err = %s", err.Error())
		c.Failure("操作失败")
	}

	c.Success(nil)
}

// 删除
func (c *PlanController) Delete() {
	var param models.DeletePlanParam
	var err error
	if err = c.ParseParam(&param); err != nil {
		logs.Info("c[Plan][Delete]: 参数错误, err = %s, req = %s", err.Error(), c.Ctx.Input.RequestBody)
		c.Failure("参数错误")
	}
	if param.ID < 0 {
		logs.Info("c[Plan][Delete]: 参数错误, 无效id, req = %s", err.Error(), c.Ctx.Input.RequestBody)
		c.Failure("参数错误")
	}

	// 查询详情
	om, err := models.ReadPlanOne(param.ID)
	if err != nil {
		logs.Info("c[Plan][Delete]: 查询详情失败, err = %s, req = %s", err.Error(), c.Ctx.Input.RequestBody)
		c.Failure("参数错误")
	}
	if om.Status != models.PlanDefault {
		c.Failure("操作失败")
	}

	// 删除
	num, err := models.DeletePlanOne(param.ID)
	if err != nil {
		logs.Info("c[Plan][Delete]: 删除失败, err = %s", err.Error())
		c.Failure("操作失败")
	}

	var res = make(map[string]interface{})
	res["num"] = num

	c.Success(res)
}

// 查询考试班级列表
func (c *PlanController) ClassList() {
	var param models.ReadPlanClassRelModelListParam
	var err error
	if err = c.ParseParam(&param); err != nil {
		logs.Info("c[Class][ClassList]: 参数错误, err = %s, req = %s", err.Error(), c.Ctx.Input.RequestBody)
		c.Failure("参数错误")
	}

	// 查询列表
	list, total, err := models.ReadPlanClassRelModelListRaw(param)
	if err != nil {
		logs.Info("c[Class][ClassList]: 查询列表失败, err = %s", err.Error())
		c.Failure("获取数据失败")
	}

	var res = make(map[string]interface{})
	res["list"] = list
	res["total"] = total

	c.Success(res)
}

// 添加考试班级
func (c *PlanController) PushClass() {
	var param models.InsertPlanClassRelParam
	var err error
	if err = c.ParseParam(&param); err != nil {
		logs.Info("c[Class][PushClass]: 参数错误, err = %s, req = %s", err.Error(), c.Ctx.Input.RequestBody)
		c.Failure("参数错误")
	}

	// 查询详情
	om, err := models.ReadPlanOne(param.PlanID)
	if err != nil {
		logs.Info("c[Plan][PushClass]: 查询详情失败, err = %s, req = %s", err.Error(), c.Ctx.Input.RequestBody)
		c.Failure("参数错误")
	}
	if om.Status != models.PlanDefault {
		c.Failure("操作失败")
	}

	// 添加班级
	num, err := models.InsertOrUpdatePlanClassRelMulti(param)
	if err != nil {
		logs.Info("c[Class][PushClass]: 更新失败, err = %s", err.Error())
		c.Failure("操作失败")
	}

	var res = make(map[string]interface{})
	res["num"] = num

	c.Success(res)
}

// 删除考试班级
func (c *PlanController) DeleteClass() {
	var param models.DeletePlanClassRelParam
	var err error
	if err = c.ParseParam(&param); err != nil {
		logs.Info("c[Class][DeleteClass]: 参数错误, err = %s, req = %s", err.Error(), c.Ctx.Input.RequestBody)
		c.Failure("参数错误")
	}
	if param.ID < 0 && len(param.List) == 0 {
		logs.Info("c[Class][DeleteClass]: 参数错误, 无效id或list为空, req = %s", err.Error(), c.Ctx.Input.RequestBody)
		c.Failure("参数错误")
	}
	for _, v := range param.List {
		if v <= 0 {
			logs.Info("c[Class][DeleteClass]: 参数错误, 无效id, req = %s", err.Error(), c.Ctx.Input.RequestBody)
			c.Failure("参数错误")
		}
	}

	// 查询详情
	om, err := models.ReadPlanOne(param.ID)
	if err != nil {
		logs.Info("c[Plan][DeleteClass]: 查询详情失败, err = %s, req = %s", err.Error(), c.Ctx.Input.RequestBody)
		c.Failure("参数错误")
	}
	if om.Status != models.PlanDefault {
		c.Failure("操作失败")
	}

	var num int64
	if param.ID > 0 {
		// 删除
		num, err = models.DeletePlanClassRelOne(param.ID)
		if err != nil {
			logs.Info("c[Class][DeleteClass]: 删除失败, err = %s", err.Error())
			c.Failure("操作失败")
		}
	} else if len(param.List) == 1 {
		// 删除
		num, err = models.DeletePlanClassRelOne(param.List[0])
		if err != nil {
			logs.Info("c[Class][DeleteClass]: 删除失败, err = %s", err.Error())
			c.Failure("操作失败")
		}
	} else {
		// 批量删除
		num, err = models.DeletePlanClassRelMulti(param.List)
		if err != nil {
			logs.Info("c[Class][DeleteClass]: 批量删除失败, err = %s", err.Error())
			c.Failure("操作失败")
		}
	}

	var res = make(map[string]interface{})
	res["num"] = num

	c.Success(res)
}
