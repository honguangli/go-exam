package controllers

import (
	"github.com/astaxie/beego/logs"
	"go-exam/models"
	"time"
)

// 考试
type ExamController struct {
	BaseController
}

// 预执行
func (c *ExamController) Prepare() {
	c.BaseController.Prepare()
}

// 查询考生待考试列表
func (c *ExamController) List() {
	var param models.ReadGradeRelListParam
	var err error
	if err = c.ParseParam(&param); err != nil {
		logs.Info("c[grade][user_exam_list]: 参数错误, err = %s, req = %s", err.Error(), c.Ctx.Input.RequestBody)
		c.Failure("参数错误")
	}
	if len(param.UserName) == 0 {
		logs.Info("c[grade][user_grade_list]: 参数错误, user_name不能为空, req = %s", c.Ctx.Input.RequestBody)
		c.Failure("参数错误")
	}

	// 查询列表
	param.StatusList = []int{models.GradeDefault, models.GradeUnSubmit}
	param.PlanStatusList = []int{models.PlanPublished}
	param.PlanQueryGrade = -1
	list, total, err := models.ReadGradeRelListRaw(param)
	if err != nil {
		logs.Info("c[grade][user_exam_list]: 查询列表失败, err = %s", err.Error())
		c.Failure("获取数据失败")
	}

	var res = make(map[string]interface{})
	res["list"] = list
	res["total"] = total

	c.Success(res)
}

// 开始考试
func (c *ExamController) Start() {
	var param models.StartExamParam
	var err error
	if err = c.ParseParam(&param); err != nil {
		logs.Info("c[exam][start]: 参数错误, err = %s, req = %s", err.Error(), c.Ctx.Input.RequestBody)
		c.Failure("参数错误")
	}
	if param.ID <= 0 {
		logs.Info("c[exam][start]: 参数错误, id非法, req = %s", c.Ctx.Input.RequestBody)
		c.Failure("参数错误")
	}
	if len(param.UserName) == 0 {
		logs.Info("c[exam][start]: 参数错误, 用户名不能为空, req = %s", c.Ctx.Input.RequestBody)
		c.Failure("参数错误")
	}
	var nowUnix = time.Now().Unix()

	// 查询待考详情
	m, err := models.ReadGradeRelOne(param.ID)
	if err != nil {
		logs.Info("c[exam][start]: 查询待考科目失败, err = %s", err.Error())
		c.Failure("获取数据失败")
	}
	// 考生名称
	if m.UserName != param.UserName {
		logs.Info("c[exam][start]: 用户名不一致")
		c.Failure("获取数据失败")
	}
	// 考生状态
	if m.UserStatus != models.UserEnable {
		logs.Info("c[exam][start]: 账号被禁用")
		c.Failure("账号已禁用")
	}
	// 考试状态
	if m.Status != models.GradeDefault && m.Status != models.GradeUnSubmit {
		logs.Info("c[exam][start]: 非待考状态")
		c.Failure("获取数据失败")
	}
	// 考试计划状态
	if m.PlanStatus != models.PlanPublished {
		logs.Info("c[exam][start]: 考试计划非待考状态")
		c.Failure("获取数据失败")
	}
	// 考试计划时间
	if nowUnix < m.PlanStartTime {
		logs.Info("c[exam][start]: 考试计划未开始")
		c.Failure("考试未开始")
	}
	if nowUnix >= m.PlanEndTime {
		logs.Info("c[exam][start]: 考试计划已结束")
		c.Failure("考试已结束")
	}
	if m.StartTime > 0 && nowUnix >= (m.StartTime+int64(m.PlanDuration*60)) {
		logs.Info("c[exam][start]: 考试时间已超过")
		c.Failure("考试已结束")
	}

	// 查询试题信息
	questionList, _, err := models.ReadPaperQuestionListRaw(models.ReadPaperQuestionListParam{
		PaperID:   m.PaperID,
		ClosePage: true,
	})
	if err != nil {
		logs.Info("c[exam][start]: 查询试题列表失败, err = %s", err.Error())
		c.Failure("获取数据失败")
	}
	if len(questionList) == 0 {
		logs.Info("c[exam][start]: 试题为空")
		c.Failure("获取数据失败")
	}

	// 查询试题选项信息
	questionOptionList, _, err := models.ReadPaperQuestionOptionListRaw(models.ReadPaperQuestionOptionListParam{
		PaperID:   m.PaperID,
		ClosePage: true,
	})
	if err != nil {
		logs.Info("c[exam][start]: 查询试题选项列表失败, err = %s", err.Error())
		c.Failure("获取数据失败")
	}

	// 更新考试
	if m.Status == models.GradeDefault {
		m.StartTime = nowUnix
		_, err = models.UpdateGradeOne(models.Grade{
			ID:        m.ID,
			Status:    models.GradeUnSubmit,
			StartTime: nowUnix,
		}, "status", "start_time")
		if err != nil {
			logs.Info("c[exam][start]: 更新考试数据失败, err = %s", err.Error())
			c.Failure("获取数据失败")
		}
	}

	var result = make(map[string]interface{})
	result["detail"] = m
	result["question_list"] = questionList
	result["option_list"] = questionOptionList

	c.Success(result)
}

// 提交答题卡
func (c *ExamController) Submit() {
	var param models.SubmitExamParam
	var err error
	if err = c.ParseParam(&param); err != nil {
		logs.Info("c[exam][submit]: 参数错误, err = %s, req = %s", err.Error(), c.Ctx.Input.RequestBody)
		c.Failure("参数错误")
	}
	if param.ID <= 0 {
		logs.Info("c[exam][submit]: 参数错误, id非法, req = %s", c.Ctx.Input.RequestBody)
		c.Failure("参数错误")
	}
	if len(param.UserName) == 0 {
		logs.Info("c[exam][submit]: 参数错误, 用户名不能为空, req = %s", c.Ctx.Input.RequestBody)
		c.Failure("参数错误")
	}
	if len(param.Answers) == 0 {
		logs.Info("c[exam][submit]: 参数错误, 答题卡不能为空, req = %s", c.Ctx.Input.RequestBody)
		c.Failure("参数错误")
	}

	var nowUnix = time.Now().Unix()

	// 查询待考详情
	m, err := models.ReadGradeRelOne(param.ID)
	if err != nil {
		logs.Info("c[exam][submit]: 查询待考科目失败, err = %s", err.Error())
		c.Failure("获取数据失败")
	}
	// 考生名称
	if m.UserName != param.UserName {
		logs.Info("c[exam][submit]: 用户名不一致")
		c.Failure("获取数据失败")
	}
	// 考生状态
	if m.UserStatus != models.UserEnable {
		logs.Info("c[exam][submit]: 账号被禁用")
		c.Failure("账号已禁用")
	}
	// 考试状态
	if m.Status != models.GradeDefault && m.Status != models.GradeUnSubmit {
		logs.Info("c[exam][submit]: 非待考状态")
		c.Failure("获取数据失败")
	}
	// 考试计划状态
	if m.PlanStatus != models.PlanPublished {
		logs.Info("c[exam][submit]: 考试计划非待考状态")
		c.Failure("获取数据失败")
	}
	// 考试计划时间
	if nowUnix < m.PlanStartTime {
		logs.Info("c[exam][submit]: 考试计划未开始")
		c.Failure("考试未开始")
	}
	if nowUnix >= m.PlanEndTime {
		logs.Info("c[exam][submit]: 考试计划已结束")
		c.Failure("考试已结束")
	}
	if nowUnix >= (m.StartTime + int64(m.PlanDuration*60)) {
		logs.Info("c[exam][start]: 考试时间已超过")
		c.Failure("考试已结束")
	}

	// 保存答题卡
	err = models.SubmitAnswer(&m, param)
	if err != nil {
		logs.Info("c[exam][submit]: 保存答题卡失败, err = %s", err.Error())
		c.Failure("提交失败")
	}

	c.Success(nil)
}
