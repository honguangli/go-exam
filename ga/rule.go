package ga

import (
	"fmt"
	"go-exam/models"
)

// 组卷规则
type Rule struct {
	Size              int     `json:"size"`                // 初始种群大小
	MutationRate      float64 `json:"mutation_rate"`       // 变异概率
	Elitism           bool    `json:"elitism"`             // 精英主义
	SubjectID         int     `json:"subject_id"`          // 科目id
	Score             int     `json:"score"`               // 试卷总分
	Difficulty        float64 `json:"difficulty"`          // 试卷期望难度系数
	DifficultyWeight  float64 `json:"difficulty_weight"`   // 难度权重
	Points            []int   `json:"points"`              // 试卷期望包含知识点
	PointsWeight      float64 `json:"points_weight"`       // 知识点权重
	ChoiceSingleNum   int     `json:"choice_single_num"`   // 单选题数量
	ChoiceSingleScore int     `json:"choice_single_score"` // 单选题分值
	ChoiceMultiNum    int     `json:"choice_multi_num"`    // 多选题数量
	ChoiceMultiScore  int     `json:"choice_multi_score"`  // 多选题分值
	JudgeNum          int     `json:"judge_num"`           // 判断题数量
	JudgeScore        int     `json:"judge_score"`         // 判断题分值
	BlankSingleNum    int     `json:"blank_single_num"`    // 填空题数量
	BlankSingleScore  int     `json:"blank_single_score"`  // 填空题分值
	BlankMultiNum     int     `json:"blank_multi_num"`     // 多项填空题数量
	BlankMultiScore   int     `json:"blank_multi_score"`   // 多项填空题分值
	AnswerSingleNum   int     `json:"answer_single_num"`   // 简答题数量
	AnswerSingleScore int     `json:"answer_single_score"` // 简答题分值
	AnswerMultiNum    int     `json:"answer_multi_num"`    // 多项简答题数量
	AnswerMultiScore  int     `json:"answer_multi_score"`  // 多项简答题分值
}

const (
	SIZE_MIN           = 10       // 初始种群大小最小值
	SIZE_MAX           = 999      // 初始种群大小最大值
	MUTATION_RATE_MIN  = 0.000001 // 变异概率最小值
	MUTATION_RATE_MAX  = 0.999999 // 变异概率最大值
	DIFFICULTY_MIN     = 0.01     // 试卷期望难度系数最小值
	DIFFICULTY_MAX     = 0.99     // 试卷期望难度系数最大值
	WEIGHT_MIN         = 0.0      // 难度权重最小值
	WEIGHT_MAX         = 1.0      // 难度权重最大值
	QUESTION_SCORE_MIN = 1        // 试题分数最小值
)

// 校验
func (m *Rule) Check() error {
	// 初始种群大小
	if m.Size < SIZE_MIN {
		return fmt.Errorf("无效参数: 初始种群大小=%d, 参考取值范围: [%d, %d]", m.Size, SIZE_MIN, SIZE_MAX)
	}
	// 变异概率
	if m.MutationRate < MUTATION_RATE_MIN || m.MutationRate > MUTATION_RATE_MAX {
		return fmt.Errorf("无效参数: 变异概率=%f, 参考取值范围: [%f, %f]", m.MutationRate, MUTATION_RATE_MIN, MUTATION_RATE_MAX)
	}
	// 科目id
	if m.SubjectID <= 0 {
		return fmt.Errorf("无效参数: 未设定科目")
	}
	// 难度
	if m.Difficulty < DIFFICULTY_MIN || m.Difficulty > DIFFICULTY_MAX {
		return fmt.Errorf("无效参数: 难度=%f, 参考取值范围: [%f, %f]", m.Difficulty, DIFFICULTY_MIN, DIFFICULTY_MAX)
	}
	// 难度权重
	if m.DifficultyWeight < WEIGHT_MIN || m.DifficultyWeight > WEIGHT_MAX {
		return fmt.Errorf("无效参数: 难度权重=%f, 参考取值范围: [%f, %f]", m.Difficulty, WEIGHT_MIN, WEIGHT_MAX)
	}
	// 知识点权重
	if m.PointsWeight < WEIGHT_MIN || m.PointsWeight > WEIGHT_MAX {
		return fmt.Errorf("无效参数: 知识点权重=%f, 参考取值范围: [%f, %f]", m.PointsWeight, WEIGHT_MIN, WEIGHT_MAX)
	}
	// 各项指标权重
	var totalWeight = m.DifficultyWeight + m.PointsWeight
	if totalWeight != WEIGHT_MAX {
		return fmt.Errorf("无效参数: 各项指标权重=%f, 参考取值范围: %f", m.PointsWeight, WEIGHT_MAX)
	}
	// 试题数量
	var totalNum = m.ChoiceSingleNum + m.ChoiceMultiNum + m.JudgeNum + m.BlankSingleNum + m.BlankMultiNum + m.AnswerSingleNum + m.AnswerMultiNum
	if totalNum <= 0 {
		return fmt.Errorf("无效参数: 未设定试题数量")
	}
	if m.ChoiceSingleNum > 0 && m.ChoiceSingleScore < QUESTION_SCORE_MIN {
		return fmt.Errorf("无效参数: 未设定单选题分数")
	}
	if m.ChoiceMultiNum > 0 && m.ChoiceMultiScore < QUESTION_SCORE_MIN {
		return fmt.Errorf("无效参数: 未设定多选题分数")
	}
	if m.JudgeNum > 0 && m.JudgeScore < QUESTION_SCORE_MIN {
		return fmt.Errorf("无效参数: 未设定判断题分数")
	}
	if m.BlankSingleNum > 0 && m.BlankSingleScore < QUESTION_SCORE_MIN {
		return fmt.Errorf("无效参数: 未设定填空题分数")
	}
	if m.BlankMultiNum > 0 && m.BlankMultiScore < QUESTION_SCORE_MIN {
		return fmt.Errorf("无效参数: 未设定多项填空题分数")
	}
	if m.AnswerSingleNum > 0 && m.AnswerSingleScore < QUESTION_SCORE_MIN {
		return fmt.Errorf("无效参数: 未设定简答题分数")
	}
	if m.AnswerMultiNum > 0 && m.AnswerMultiScore < QUESTION_SCORE_MIN {
		return fmt.Errorf("无效参数: 未设定多项简答题分数")
	}
	// 总分
	//var totalScore = m.ChoiceSingleNum*m.ChoiceSingleScore +
	//	m.ChoiceMultiNum*m.ChoiceMultiScore +
	//	m.JudgeNum*m.JudgeScore +
	//	m.BlankSingleNum*m.BlankSingleScore +
	//	m.BlankMultiNum*m.BlankMultiScore +
	//	m.AnswerSingleNum*m.AnswerSingleScore +
	//	m.AnswerMultiNum*m.AnswerMultiScore
	if m.Score <= 0 {
		return fmt.Errorf("无效参数: 未设定试题分数")
	}
	return nil
}

func (m *Rule) getScore(questionType int) int {
	switch questionType {
	case models.QUESTION_CHOICE_SINGLE: // 单选题
		return m.ChoiceSingleScore
	case models.QUESTION_CHOICE_MULTI: // 多选题
		return m.ChoiceMultiScore
	case models.QUESTION_JUDGE: // 判断题
		return m.JudgeScore
	case models.QUESTION_BLANK_SINGLE: // 填空题
		return m.BlankSingleScore
	case models.QUESTION_BLANK_MULTI: // 多项填空题
		return m.BlankMultiScore
	case models.QUESTION_ANSWER: // 简答题
		return m.AnswerSingleScore
	case models.QUESTION_ANSWER_MULTI: // 多项简答题
		return m.AnswerMultiScore
	}
	return 0
}
