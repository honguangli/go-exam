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

// 查询列表
func (c *QuestionController) List() {
	var param models.ReadQuestionListParam
	var err error
	if err = c.ParseParam(&param); err != nil {
		logs.Info("c[Question][List]: 参数错误, err = %s, req = %s", err.Error(), c.Ctx.Input.RequestBody)
		c.Failure("参数错误")
	}

	// 查询列表
	list, total, err := models.ReadQuestionListRaw(param)
	if err != nil {
		logs.Info("c[Question][List]: 查询列表失败, err = %s", err.Error())
		c.Failure("获取数据失败")
	}

	var res = make(map[string]interface{})
	res["list"] = list
	res["total"] = total

	c.Success(res)
}

// 查询详情
func (c *QuestionController) Detail() {
	var param models.ReadQuestionDetailParam
	var err error
	if err = c.ParseParam(&param); err != nil {
		logs.Info("c[Question][Detail]: 参数错误, err = %s, req = %s", err.Error(), c.Ctx.Input.RequestBody)
		c.Failure("参数错误")
	}

	// 查询详情
	m, err := models.ReadQuestionOne(param.ID)
	if err != nil {
		logs.Info("c[Question][Detail]: 查询详情失败, err = %s", err.Error())
		c.Failure("获取数据失败")
	}

	// 查询选项
	options, _, err := models.ReadQuestionOptionListRaw(models.ReadQuestionOptionListParam{
		QuestionID: m.ID,
		ClosePage:  true,
	})
	if err != nil {
		logs.Info("c[Question][Detail]: 查询选项列表失败, err = %s", err.Error())
		c.Failure("获取数据失败")
	}

	var res = make(map[string]interface{})
	res["detail"] = m
	res["options"] = options

	c.Success(res)
}

// 创建
func (c *QuestionController) Create() {
	var param models.InsertQuestionParam
	var err error
	if err = c.ParseParam(&param); err != nil {
		logs.Info("c[Question][Create]: 参数错误, err = %s, req = %s", err.Error(), c.Ctx.Input.RequestBody)
		c.Failure("参数错误")
	}

	// 创建
	id, err := models.InsertQuestionOneWithOptions(param)
	if err != nil {
		logs.Info("c[Question][Create]: 创建失败, err = %s", err.Error())
		c.Failure("操作失败")
	}

	var res = make(map[string]interface{})
	res["id"] = id

	c.Success(res)
}

// 更新
func (c *QuestionController) Update() {
	var param models.UpdateQuestionParam
	var err error
	if err = c.ParseParam(&param); err != nil {
		logs.Info("c[Question][Update]: 参数错误, err = %s, req = %s", err.Error(), c.Ctx.Input.RequestBody)
		c.Failure("参数错误")
	}

	// 更新
	_, err = models.UpdateQuestionOneWithOptions(param)
	if err != nil {
		logs.Info("c[Question][Update]: 更新失败, err = %s", err.Error())
		c.Failure("操作失败")
	}

	c.Success(nil)
}

// 删除
func (c *QuestionController) Delete() {
	var param models.DeleteQuestionParam
	var err error
	if err = c.ParseParam(&param); err != nil {
		logs.Info("c[Question][Delete]: 参数错误, err = %s, req = %s", err.Error(), c.Ctx.Input.RequestBody)
		c.Failure("参数错误")
	}
	if param.ID < 0 && len(param.List) == 0 {
		logs.Info("c[Question][Delete]: 参数错误, 无效id或list为空, req = %s", err.Error(), c.Ctx.Input.RequestBody)
		c.Failure("参数错误")
	}
	for _, v := range param.List {
		if v <= 0 {
			logs.Info("c[Question][Delete]: 参数错误, 无效id, req = %s", err.Error(), c.Ctx.Input.RequestBody)
			c.Failure("参数错误")
		}
	}

	var num int64
	if param.ID > 0 {
		// 删除
		num, err = models.DeleteQuestionOne(param.ID)
		if err != nil {
			logs.Info("c[Question][Delete]: 删除失败, err = %s", err.Error())
			c.Failure("操作失败")
		}
	} else if len(param.List) == 1 {
		// 删除
		num, err = models.DeleteQuestionOne(param.List[0])
		if err != nil {
			logs.Info("c[Question][Delete]: 删除失败, err = %s", err.Error())
			c.Failure("操作失败")
		}
	} else {
		// 批量删除
		num, err = models.DeleteQuestionMulti(param.List)
		if err != nil {
			logs.Info("c[Question][Delete]: 批量删除失败, err = %s", err.Error())
			c.Failure("操作失败")
		}
	}

	var res = make(map[string]interface{})
	res["num"] = num

	c.Success(res)
}
