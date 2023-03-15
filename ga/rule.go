package ga

// 组卷规则
type Rule struct {
	Size              int     `json:"size"`                // 初始种群大小
	MutationRate      float64 `json:"mutation_rate"`       // 变异概率
	Elitism           bool    `json:"elitism"`             // 精英主义
	Score             int     `json:"score"`               // 试卷总分
	Difficulty        float64 `json:"difficulty"`          // 试卷期望难度系数
	DifficultyWeight  float64 `json:"difficulty_weight"`   // 难度权重
	Points            []int   `json:"points"`              // 试卷期望包含知识点
	PointsWeight      float64 `json:"points_weight"`       // 知识点权重
	ChoiceSingleNum   int     `json:"single_num"`          // 单选题数量
	ChoiceSingleScore int     `json:"single_score"`        // 单选题分值
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

func (m *Rule) getScore(questionType int) int {
	switch questionType {
	case QUESTION_CHOICE_SINGLE: // 单选题
		return m.ChoiceSingleScore
	case QUESTION_CHOICE_MULTI: // 多选题
		return m.ChoiceMultiScore
	case QUESTION_JUDGE: // 判断题
		return m.JudgeScore
	case QUESTION_BLANK_SINGLE: // 填空题
		return m.BlankSingleScore
	case QUESTION_BLANK_MULTI: // 多项填空题
		return m.BlankMultiScore
	case QUESTION_ANSWER: // 简答题
		return m.AnswerSingleScore
	case QUESTION_ANSWER_MULTI: // 多项简答题
		return m.AnswerMultiScore
	}
	return 0
}
