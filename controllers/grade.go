package controllers

import (
	"github.com/astaxie/beego/logs"
	"go-exam/models"
)

// 成绩
type GradeController struct {
	BaseController
}

// 预执行
func (c *GradeController) Prepare() {
	c.BaseController.Prepare()
}

// 查询列表
func (c *GradeController) List() {
	var param models.ReadGradeRelListParam
	var err error
	if err = c.ParseParam(&param); err != nil {
		logs.Info("c[grade][list]: 参数错误, err = %s, req = %s", err.Error(), c.Ctx.Input.RequestBody)
		c.Failure("参数错误")
	}

	// 查询列表
	param.PlanQueryGrade = -1
	list, total, err := models.ReadGradeRelListRaw(param)
	if err != nil {
		logs.Info("c[grade][list]: 查询列表失败, err = %s", err.Error())
		c.Failure("获取数据失败")
	}

	var res = make(map[string]interface{})
	res["list"] = list
	res["total"] = total

	c.Success(res)
}

// 查询考生成绩列表
func (c *GradeController) UserGradeList() {
	var param models.ReadGradeRelListParam
	var err error
	if err = c.ParseParam(&param); err != nil {
		logs.Info("c[grade][user_grade_list]: 参数错误, err = %s, req = %s", err.Error(), c.Ctx.Input.RequestBody)
		c.Failure("参数错误")
	}
	if len(param.UserName) == 0 {
		logs.Info("c[grade][user_grade_list]: 参数错误, user_name不能为空, req = %s", c.Ctx.Input.RequestBody)
		c.Failure("参数错误")
	}

	// 查询列表
	param.PlanStatusList = []int{models.PlanPublished, models.PlanCanceled, models.PlanEnded}
	param.PlanQueryGrade = models.PlanGradeEnable
	list, total, err := models.ReadGradeRelListRaw(param)
	if err != nil {
		logs.Info("c[grade][user_grade_list]: 查询列表失败, err = %s", err.Error())
		c.Failure("获取数据失败")
	}

	var res = make(map[string]interface{})
	res["list"] = list
	res["total"] = total

	c.Success(res)
}

// 查询详情
func (c *GradeController) Detail() {
	var param models.ReadGradeDetailParam
	var err error
	if err = c.ParseParam(&param); err != nil {
		logs.Info("c[grade][detail]: 参数错误, err = %s, req = %s", err.Error(), c.Ctx.Input.RequestBody)
		c.Failure("参数错误")
	}

	// 查询详情
	m, err := models.ReadGradeOne(param.ID)
	if err != nil {
		logs.Info("c[grade][detail]: 查询详情失败, err = %s", err.Error())
		c.Failure("获取数据失败")
	}

	c.Success(m)
}
