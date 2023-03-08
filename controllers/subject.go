package controllers

import (
	"github.com/astaxie/beego/logs"
	"go-exam/models"
)

// 科目
type SubjectController struct {
	BaseController
}

// 预执行
func (c *SubjectController) Prepare() {
	c.BaseController.Prepare()
}

// 查询列表
func (c *SubjectController) List() {
	var param models.ReadSubjectListParam
	var err error
	if err = c.ParseParam(&param); err != nil {
		logs.Info("c[subject][list]: 参数错误, err = %s, req = %s", err.Error(), c.Ctx.Input.RequestBody)
		c.Failure("参数错误")
	}

	// 查询列表
	list, total, err := models.ReadSubjectList(param)
	if err != nil {
		logs.Info("c[subject][list]: 查询列表失败, err = %s", err.Error())
		c.Failure("获取数据失败")
	}

	var res = make(map[string]interface{})
	res["list"] = list
	res["total"] = total

	c.Success(res)
}

// 查询详情
func (c *SubjectController) Detail() {
	var param models.ReadSubjectDetailParam
	var err error
	if err = c.ParseParam(&param); err != nil {
		logs.Info("c[subject][detail]: 参数错误, err = %s, req = %s", err.Error(), c.Ctx.Input.RequestBody)
		c.Failure("参数错误")
	}

	// 查询详情
	m, err := models.ReadSubjectOne(param.ID)
	if err != nil {
		logs.Info("c[subject][detail]: 查询详情失败, err = %s", err.Error())
		c.Failure("获取数据失败")
	}

	c.Success(m)
}

// 创建
func (c *SubjectController) Create() {
	var m models.Subject
	var err error
	if err = c.ParseParam(&m); err != nil {
		logs.Info("c[subject][create]: 参数错误, err = %s, req = %s", err.Error(), c.Ctx.Input.RequestBody)
		c.Failure("参数错误")
	}

	// 创建
	id, err := models.InsertSubjectOne(m)
	if err != nil {
		logs.Info("c[subject][create]: 创建失败, err = %s", err.Error())
		c.Failure("获取数据失败")
	}

	var res = make(map[string]interface{})
	res["id"] = id

	c.Success(res)
}

// 更新
func (c *SubjectController) Update() {
	var m models.Subject
	var err error
	if err = c.ParseParam(&m); err != nil {
		logs.Info("c[subject][update]: 参数错误, err = %s, req = %s", err.Error(), c.Ctx.Input.RequestBody)
		c.Failure("参数错误")
	}

	// 更新
	_, err = models.UpdateSubjectOne(m, "name", "desc")
	if err != nil {
		logs.Info("c[subject][update]: 更新失败, err = %s", err.Error())
		c.Failure("获取数据失败")
	}

	c.Success(nil)
}

// 删除
func (c *SubjectController) Delete() {
	var param models.DeleteSubjectDetailParam
	var err error
	if err = c.ParseParam(&param); err != nil {
		logs.Info("c[subject][delete]: 参数错误, err = %s, req = %s", err.Error(), c.Ctx.Input.RequestBody)
		c.Failure("参数错误")
	}
	if param.ID < 0 && len(param.List) == 0 {
		logs.Info("c[subject][delete]: 参数错误, 无效id或list为空, req = %s", err.Error(), c.Ctx.Input.RequestBody)
		c.Failure("参数错误")
	}
	for _, v := range param.List {
		if v <= 0 {
			logs.Info("c[subject][delete]: 参数错误, 无效id, req = %s", err.Error(), c.Ctx.Input.RequestBody)
			c.Failure("参数错误")
		}
	}

	var num int64
	if param.ID > 0 {
		// 删除
		num, err = models.DeleteSubjectOne(param.ID)
		if err != nil {
			logs.Info("c[subject][delete]: 删除失败, err = %s", err.Error())
			c.Failure("获取数据失败")
		}
	} else if len(param.List) == 1 {
		// 删除
		num, err = models.DeleteSubjectOne(param.List[0])
		if err != nil {
			logs.Info("c[subject][delete]: 删除失败, err = %s", err.Error())
			c.Failure("获取数据失败")
		}
	} else {
		// 批量删除
		num, err = models.DeleteSubjectMulti(param.List)
		if err != nil {
			logs.Info("c[subject][delete]: 批量删除失败, err = %s", err.Error())
			c.Failure("获取数据失败")
		}
	}

	var res = make(map[string]interface{})
	res["num"] = num

	c.Success(res)
}
