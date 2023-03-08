package controllers

import (
	"github.com/astaxie/beego/logs"
	"go-exam/models"
)

// 试题
type QuestionController struct {
	BaseController
}

// 预执行
func (c *QuestionController) Prepare() {
	c.BaseController.Prepare()
}

// 查询用户信息
func (c *QuestionController) QueryUserInfo() {
	type Param struct {
		Path string `json:"path"`
	}
	var param Param
	var err error
	if err = c.ParseParam(&param); err != nil {
		c.Failure("参数错误")
	}
	if len(param.Path) == 0 {
		c.Failure("参数错误")
	}

	c.Success(nil)
}

// 查询试题列表
func (c *QuestionController) List() {
	var param models.ReadQuestionListParam
	var err error
	if err = c.ParseParam(&param); err != nil {
		c.Failure("参数错误")
	}

	// 查询试题列表
	list, total, err := models.ReadQuestionListRaw(param)
	if err != nil {
		logs.Info("c[question][list]: 查询试题列表失败, err = %s", err.Error())
		c.Failure("获取数据失败")
	}

	var res = make(map[string]interface{})
	res["list"] = list
	res["total"] = total

	c.Success(res)
}
