package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego/logs"
	"go-exam/ga"
	"go-exam/models"
	"strconv"
	"strings"
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
	list, total, err := models.ReadPaperListRaw(param)
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

	// 查询选项列表

	c.Success(m)
}

// 智能组卷
func (c *PaperController) Auto() {
	var m models.Paper
	var err error
	if err = c.ParseParam(&m); err != nil {
		logs.Info("c[Paper][Auto]: 参数错误, err = %s, req = %s", err.Error(), c.Ctx.Input.RequestBody)
		c.Failure("参数错误")
	}

	//var startTime = time.Now()
	p, err := c.genPaper(&m)
	if err != nil {
		logs.Info("c[Paper][Auto]: 创建失败, err = %s", err.Error())
		c.Failure("操作失败")
	}
	m.Difficulty = p.Difficulty
	id, err := models.InsertPaperOneWithQuestionList(m, p.QuestionList)
	if err != nil {
		logs.Info("c[Paper][Auto]: 创建失败, err = %s", err.Error())
		c.Failure("操作失败")
	}
	//logs.Info("总耗时：%v", time.Now().Sub(startTime))

	// 调用组卷算法获取试题列表

	// 保存试卷信息

	// 保存试卷试题列表（试题列表、试题选项列表）

	// 创建
	//m.CreateTime = time.Now().Unix()
	//id, err := models.InsertPaperOne(m)
	//if err != nil {
	//	logs.Info("c[Paper][Create]: 创建失败, err = %s", err.Error())
	//	c.Failure("操作失败")
	//}

	var res = make(map[string]interface{})
	res["id"] = id

	c.Success(res)
}

// 手工组卷
func (c *PaperController) Create() {
	var m models.Paper
	var err error
	if err = c.ParseParam(&m); err != nil {
		logs.Info("c[Paper][Create]: 参数错误, err = %s, req = %s", err.Error(), c.Ctx.Input.RequestBody)
		c.Failure("参数错误")
	}

	// 获取参数（试卷参数、试题集合）

	// 更新试卷信息

	// 删除原有试题列表、选项列表，保存新试卷试题列表（试题列表、试题选项列表）

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

func (c *PaperController) genPaper(param *models.Paper) (*ga.Paper, error) {
	logs.Info("测试组卷算法")
	var startTime = time.Now()
	var points = make([]int, 0)
	for _, v := range strings.Split(param.KnowledgeIds, ",") {
		if id, err := strconv.Atoi(v); err == nil {
			points = append(points, id)
		}
	}
	// 创建规则
	var rule = ga.Rule{
		Size:              10,                      // 初始种群大小
		MutationRate:      0.085,                   // 变异概率
		Elitism:           true,                    // 精英主义
		SubjectID:         param.SubjectID,         // 科目id
		Score:             param.Score,             // 试卷总分
		Difficulty:        param.Difficulty,        // 试卷期望难度系数
		DifficultyWeight:  0.80,                    // 难度权重
		Points:            points,                  // 试卷期望包含知识点
		PointsWeight:      0.20,                    // 知识点权重
		ChoiceSingleNum:   param.ChoiceSingleNum,   // 单选题数量
		ChoiceSingleScore: param.ChoiceSingleScore, // 单选题分值
		ChoiceMultiNum:    param.ChoiceMultiNum,    // 多选题数量
		ChoiceMultiScore:  param.ChoiceMultiScore,  // 多选题分值
		JudgeNum:          param.JudgeNum,          // 判断题数量
		JudgeScore:        param.JudgeScore,        // 判断题分值
		BlankSingleNum:    param.BlankSingleNum,    // 填空题数量
		BlankSingleScore:  param.BlankSingleScore,  // 填空题分值
		BlankMultiNum:     param.BlankMultiNum,     // 多项填空题数量
		BlankMultiScore:   param.BlankMultiScore,   // 多项填空题分值
		AnswerSingleNum:   param.AnswerSingleNum,   // 简答题数量
		AnswerSingleScore: param.AnswerSingleScore, // 简答题分值
		AnswerMultiNum:    param.AnswerMultiNum,    // 多项简答题数量
		AnswerMultiScore:  param.AnswerMultiScore,  // 多项简答题分值
	}
	if err := rule.Check(); err != nil {
		logs.Info("无效参数：%s", err.Error())
		return nil, err
	}

	// 初始化种群
	population, err := ga.NewPopulation(&rule)
	if err != nil {
		logs.Info("初始化种群失败：%s", err.Error())
		return nil, err
	}

	//for k, v := range population.Papers {
	//	fmt.Printf("k = %d | d = %f | k = %f | fit = %f | calc = %f\n", k, v.Difficulty, v.KPCoverage, v.AdaptationDegree, 1-(1-v.KPCoverage)*rule.PointsWeight-math.Abs(rule.Difficulty-v.Difficulty)*rule.DifficultyWeight)
	//}

	// 迭代次数
	var count = 100

	// 期望适应度
	var expand = 0.98

	var paper *ga.Paper
	for i := 0; i < count; i++ {
		paper = population.GetFitness()
		//logs.Info("适应度：%f", paper.AdaptationDegree)

		if paper.AdaptationDegree >= expand {
			break
		}

		//logs.Info("第%d次进化", i+1)
		err = population.Evolve()
		if err != nil {
			logs.Info("进化失败: %s", err.Error())
			return nil, err
		}
		//for k, v := range population.Papers {
		//	fmt.Printf("k = %d | d = %f(%f) | k = %f(%f) | fit = %f(%f)\n", k, v.Difficulty, v.GetDifficulty(), v.KPCoverage, v.GetKpCoverage(rule.Points), v.AdaptationDegree, v.GetAdaptationDegree(rule.Difficulty, rule.PointsWeight, rule.DifficultyWeight))
		//}
	}
	logs.Info("组卷耗时：%v", time.Now().Sub(startTime))
	logs.Info("进化完成: 适应度 = %f | 难度 = %f | 知识点覆盖率 = %f", paper.AdaptationDegree, paper.Difficulty, paper.KPCoverage)

	bf, err := json.Marshal(paper)
	if err != nil {
		logs.Info("json marshal error: %s", err.Error())
	} else {
		logs.Info("paper: %s", string(bf))
	}

	return paper, nil
}
