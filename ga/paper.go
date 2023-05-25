package ga

import (
	"fmt"
	"go-exam/models"
	"math"
	"strconv"
	"strings"
)

// 试卷
type Paper struct {
	Score            int                // 总分
	Difficulty       float64            // 难度系数
	KPCoverage       float64            // 知识点覆盖率
	AdaptationDegree float64            // 适应度
	QuestionList     []*models.Question // 试题列表
}

// 生成试卷
func NewPaper(rule *Rule) (paper *Paper, err error) {
	paper = new(Paper)

	// 单选题
	if rule.ChoiceSingleNum > 0 {
		list, _, err := models.ReadQuestionSimpleListRaw(models.ReadQuestionSimpleListParam{
			SubjectID:       rule.SubjectID,
			Type:            models.QUESTION_CHOICE_SINGLE,
			Status:          models.QUESTION_ENABLE,
			KnowledgeIDList: rule.Points,
			ExcludeIDList:   nil,
			BaseQueryParam: models.BaseQueryParam{
				Limit: rule.ChoiceSingleNum,
			},
		})
		if err != nil {
			return nil, err
		}
		if len(list) != rule.ChoiceSingleNum {
			return nil, fmt.Errorf("单选题不足，组卷失败")
		}
		for _, v := range list {
			v.Score = rule.ChoiceSingleScore
		}
		paper.QuestionList = append(paper.QuestionList, list...)
	}

	// 多选题
	if rule.ChoiceMultiNum > 0 {
		list, _, err := models.ReadQuestionSimpleListRaw(models.ReadQuestionSimpleListParam{
			SubjectID:       rule.SubjectID,
			Type:            models.QUESTION_CHOICE_MULTI,
			Status:          models.QUESTION_ENABLE,
			KnowledgeIDList: rule.Points,
			ExcludeIDList:   nil,
			BaseQueryParam: models.BaseQueryParam{
				Limit: rule.ChoiceMultiNum,
			},
		})
		if err != nil {
			return nil, err
		}
		if len(list) != rule.ChoiceMultiNum {
			return nil, fmt.Errorf("多选题不足，组卷失败")
		}
		for _, v := range list {
			v.Score = rule.ChoiceMultiScore
		}
		paper.QuestionList = append(paper.QuestionList, list...)
	}

	// 判断题
	if rule.JudgeNum > 0 {
		list, _, err := models.ReadQuestionSimpleListRaw(models.ReadQuestionSimpleListParam{
			SubjectID:       rule.SubjectID,
			Type:            models.QUESTION_JUDGE,
			Status:          models.QUESTION_ENABLE,
			KnowledgeIDList: rule.Points,
			ExcludeIDList:   nil,
			BaseQueryParam: models.BaseQueryParam{
				Limit: rule.JudgeNum,
			},
		})
		if err != nil {
			return nil, err
		}
		if len(list) != rule.JudgeNum {
			return nil, fmt.Errorf("判断题不足，组卷失败")
		}
		for _, v := range list {
			v.Score = rule.JudgeScore
		}
		paper.QuestionList = append(paper.QuestionList, list...)
	}

	// 填空题
	if rule.BlankSingleNum > 0 {
		list, _, err := models.ReadQuestionSimpleListRaw(models.ReadQuestionSimpleListParam{
			SubjectID:       rule.SubjectID,
			Type:            models.QUESTION_BLANK_SINGLE,
			Status:          models.QUESTION_ENABLE,
			KnowledgeIDList: rule.Points,
			ExcludeIDList:   nil,
			BaseQueryParam: models.BaseQueryParam{
				Limit: rule.BlankSingleNum,
			},
		})
		if err != nil {
			return nil, err
		}
		if len(list) != rule.BlankSingleNum {
			return nil, fmt.Errorf("填空题不足，组卷失败")
		}
		for _, v := range list {
			v.Score = rule.BlankSingleScore
		}
		paper.QuestionList = append(paper.QuestionList, list...)
	}

	// 多项填空题
	if rule.BlankMultiNum > 0 {
		list, _, err := models.ReadQuestionSimpleListRaw(models.ReadQuestionSimpleListParam{
			SubjectID:       rule.SubjectID,
			Type:            models.QUESTION_BLANK_MULTI,
			Status:          models.QUESTION_ENABLE,
			KnowledgeIDList: rule.Points,
			ExcludeIDList:   nil,
			BaseQueryParam: models.BaseQueryParam{
				Limit: rule.BlankMultiNum,
			},
		})
		if err != nil {
			return nil, err
		}
		if len(list) != rule.BlankMultiNum {
			return nil, fmt.Errorf("多项填空题不足，组卷失败")
		}
		for _, v := range list {
			v.Score = rule.BlankMultiScore
		}
		paper.QuestionList = append(paper.QuestionList, list...)
	}

	// 简答题
	if rule.AnswerSingleNum > 0 {
		list, _, err := models.ReadQuestionSimpleListRaw(models.ReadQuestionSimpleListParam{
			SubjectID:       rule.SubjectID,
			Type:            models.QUESTION_ANSWER,
			Status:          models.QUESTION_ENABLE,
			KnowledgeIDList: rule.Points,
			ExcludeIDList:   nil,
			BaseQueryParam: models.BaseQueryParam{
				Limit: rule.AnswerSingleNum,
			},
		})
		if err != nil {
			return nil, err
		}
		if len(list) != rule.AnswerSingleNum {
			return nil, fmt.Errorf("简答题不足，组卷失败")
		}
		for _, v := range list {
			v.Score = rule.AnswerSingleScore
		}
		paper.QuestionList = append(paper.QuestionList, list...)
	}

	// 多项简答题
	if rule.AnswerMultiNum > 0 {
		list, _, err := models.ReadQuestionSimpleListRaw(models.ReadQuestionSimpleListParam{
			SubjectID:       rule.SubjectID,
			Type:            models.QUESTION_ANSWER_MULTI,
			Status:          models.QUESTION_ENABLE,
			KnowledgeIDList: rule.Points,
			ExcludeIDList:   nil,
			BaseQueryParam: models.BaseQueryParam{
				Limit: rule.AnswerMultiNum,
			},
		})
		if len(list) != rule.AnswerMultiNum {
			return nil, fmt.Errorf("多项简答题不足，组卷失败")
		}
		if err != nil {
			return nil, err
		}
		for _, v := range list {
			v.Score = rule.AnswerMultiScore
		}
		paper.QuestionList = append(paper.QuestionList, list...)
	}

	paper.Score = rule.Score
	paper.SetDifficulty()
	paper.SetKpCoverage(rule.Points)
	paper.SetAdaptationDegree(rule.Difficulty, rule.PointsWeight, rule.DifficultyWeight)
	return
}

// 获取试题id数组
func (m *Paper) GetQuestionIDList() []int {
	var list = make([]int, len(m.QuestionList))
	for k, v := range m.QuestionList {
		if v == nil {
			continue
		}
		list[k] = v.ID
	}
	return list
}

// 判断是否重复
func (m *Paper) CheckQuestionIsExist(id int) bool {
	for _, v := range m.QuestionList {
		if v == nil {
			continue
		}
		if v.ID == id {
			return true
		}
	}
	return false
}

// 计算试卷难度系数
// 计算公式：每题难度 * 分数求和除以总分
func (m *Paper) SetDifficulty() {
	var totalDifficulty float64
	for _, v := range m.QuestionList {
		totalDifficulty += float64(v.Score) * v.Difficulty
	}
	m.Difficulty = math.Round(totalDifficulty/float64(m.Score)*1000) / 1000
}

// 计算试卷知识点覆盖率
// 计算公式：所有试题包含知识点除以期望包含的知识点
func (m *Paper) SetKpCoverage(points []int) {
	var mp = make(map[string]byte)
	for _, v := range points {
		mp[strconv.Itoa(v)] = 0
	}

	var total int
	for _, v := range m.QuestionList {
		for _, id := range strings.Split(v.KnowledgeIds, ",") {
			if c, ok := mp[id]; ok && c == 0 {
				mp[id] = 1
				total++
			}
		}
	}
	m.KPCoverage = math.Round(float64(total)/float64(len(points))*1000) / 1000
}

// 计算试卷适应度
// f1为知识点分布的权重，f2为难度系数所占权重。
// 当f1=0时退化为只限制试题难度系数，当f2=0时退化为只限制知识点分布
func (m *Paper) SetAdaptationDegree(difficulty float64, f1 float64, f2 float64) {
	m.AdaptationDegree = 1 / (1 + (1-m.KPCoverage)*f1 + math.Abs(difficulty-m.Difficulty)*f2)
	return
}
