package controllers

import (
	"github.com/astaxie/beego/logs"
	"go-exam/models"
)

// 知识点
type KnowledgeController struct {
	BaseController
}

// 预执行
func (c *KnowledgeController) Prepare() {
	c.BaseController.Prepare()
}

// 查询列表
func (c *KnowledgeController) List() {
	var param models.ReadKnowledgeListParam
	var err error
	if err = c.ParseParam(&param); err != nil {
		logs.Info("c[knowledge][list]: 参数错误, err = %s, req = %s", err.Error(), c.Ctx.Input.RequestBody)
		c.Failure("参数错误")
	}

	// 查询列表
	list, total, err := models.ReadKnowledgeList(param)
	if err != nil {
		logs.Info("c[knowledge][list]: 查询列表失败, err = %s", err.Error())
		c.Failure("获取数据失败")
	}

	var res = make(map[string]interface{})
	res["list"] = list
	res["total"] = total

	c.Success(res)
}

// 查询详情
func (c *KnowledgeController) Detail() {
	var param models.ReadKnowledgeDetailParam
	var err error
	if err = c.ParseParam(&param); err != nil {
		logs.Info("c[knowledge][detail]: 参数错误, err = %s, req = %s", err.Error(), c.Ctx.Input.RequestBody)
		c.Failure("参数错误")
	}

	// 查询详情
	m, err := models.ReadKnowledgeOne(param.ID)
	if err != nil {
		logs.Info("c[knowledge][detail]: 查询详情失败, err = %s", err.Error())
		c.Failure("获取数据失败")
	}

	c.Success(m)
}

// 创建
func (c *KnowledgeController) Create() {
	var m models.Knowledge
	var err error
	if err = c.ParseParam(&m); err != nil {
		logs.Info("c[knowledge][create]: 参数错误, err = %s, req = %s", err.Error(), c.Ctx.Input.RequestBody)
		c.Failure("参数错误")
	}

	// 创建
	id, err := models.InsertKnowledgeOne(m)
	if err != nil {
		logs.Info("c[knowledge][create]: 创建失败, err = %s", err.Error())
		c.Failure("获取数据失败")
	}

	var res = make(map[string]interface{})
	res["id"] = id

	c.Success(res)
}

// 更新
func (c *KnowledgeController) Update() {
	var m models.Knowledge
	var err error
	if err = c.ParseParam(&m); err != nil {
		logs.Info("c[knowledge][update]: 参数错误, err = %s, req = %s", err.Error(), c.Ctx.Input.RequestBody)
		c.Failure("参数错误")
	}

	// 更新
	_, err = models.UpdateKnowledgeOne(m, "name", "desc")
	if err != nil {
		logs.Info("c[knowledge][update]: 更新失败, err = %s", err.Error())
		c.Failure("获取数据失败")
	}

	c.Success(nil)
}

// 删除
func (c *KnowledgeController) Delete() {
	var param models.DeleteKnowledgeDetailParam
	var err error
	if err = c.ParseParam(&param); err != nil {
		logs.Info("c[knowledge][delete]: 参数错误, err = %s, req = %s", err.Error(), c.Ctx.Input.RequestBody)
		c.Failure("参数错误")
	}
	if param.ID < 0 && len(param.List) == 0 {
		logs.Info("c[knowledge][delete]: 参数错误, 无效id或list为空, req = %s", err.Error(), c.Ctx.Input.RequestBody)
		c.Failure("参数错误")
	}
	for _, v := range param.List {
		if v <= 0 {
			logs.Info("c[knowledge][delete]: 参数错误, 无效id, req = %s", err.Error(), c.Ctx.Input.RequestBody)
			c.Failure("参数错误")
		}
	}

	var num int64
	if param.ID > 0 {
		// 删除
		num, err = models.DeleteKnowledgeOne(param.ID)
		if err != nil {
			logs.Info("c[knowledge][delete]: 删除失败, err = %s", err.Error())
			c.Failure("获取数据失败")
		}
	} else if len(param.List) == 1 {
		// 删除
		num, err = models.DeleteKnowledgeOne(param.List[0])
		if err != nil {
			logs.Info("c[knowledge][delete]: 删除失败, err = %s", err.Error())
			c.Failure("获取数据失败")
		}
	} else {
		// 批量删除
		num, err = models.DeleteKnowledgeMulti(param.List)
		if err != nil {
			logs.Info("c[knowledge][delete]: 批量删除失败, err = %s", err.Error())
			c.Failure("获取数据失败")
		}
	}

	var res = make(map[string]interface{})
	res["num"] = num

	c.Success(res)
}
