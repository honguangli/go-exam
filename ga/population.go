package ga

import (
	"fmt"
	"go-exam/models"
	"math/rand"
	"time"
)

// 种群
// 种群时试卷集合
type Population struct {
	Rule   *Rule      `json:"rule"`   // 组卷规则
	Papers []*Paper   `json:"papers"` // 试卷列表
	Rand   *rand.Rand `json:"-"`      // 随机生成器
}

// 生成种群
func NewPopulation(rule *Rule) (population *Population, err error) {
	var papers = make([]*Paper, rule.Size)
	for i := 0; i < rule.Size; i++ {
		paper, err := NewPaper(rule)
		if err != nil {
			return nil, err
		}
		papers[i] = paper
	}
	return &Population{
		Rule:   rule,
		Papers: papers,
		Rand:   rand.New(rand.NewSource(time.Now().UnixNano())),
	}, nil
}

// 获取种群中最优秀个体
func (m *Population) GetFitness() *Paper {
	var paper = m.Papers[0]
	for i := 1; i < len(m.Papers); i++ {
		if m.Papers[i].AdaptationDegree > paper.AdaptationDegree {
			paper = m.Papers[i]
		}
	}
	return paper
}

// 进化
func (m *Population) Evolve() error {
	var papers = make([]*Paper, len(m.Papers))
	// 精英主义
	var offset int
	if m.Rule.Elitism {
		papers[0] = m.GetFitness()
		offset = 1
	}

	// 交叉
	for i := offset; i < len(papers); i++ {
		p1 := m.Select()
		p2 := m.Select()
		for p1 == p2 {
			p2 = m.Select()
		}

		paper, err := m.Crossover(p1, p2)
		if err != nil {
			fmt.Printf("交叉失败: %s\n", err.Error())
			return err
		}
		papers[i] = paper
	}

	// 变异
	for i := offset; i < len(papers); i++ {
		err := m.Mutation(papers[i])
		if err != nil {
			fmt.Printf("变异失败: %s\n", err.Error())
			return err
		}
		papers[i].Score = m.Rule.Score
		papers[i].SetDifficulty()
		papers[i].SetKpCoverage(m.Rule.Points)
		papers[i].SetAdaptationDegree(m.Rule.Difficulty, m.Rule.PointsWeight, m.Rule.DifficultyWeight)
	}

	m.Papers = papers

	return nil
}

// 选择算子
// 算法：轮盘赌
// 个体被选中概率 = 个体适应度 / 种群全部个体的适应度之和
// 通过累积适应度计算选择个体，不需要单独计算个体被选中概率，从而减少计算次数
func (m *Population) Select() *Paper {
	// 种群全部个体的适应度之和
	var totalAdaptationDegree float64

	// 计算种群全部个体的适应度之和
	for _, v := range m.Papers {
		totalAdaptationDegree += v.AdaptationDegree
	}

	// 轮盘指针
	var pointer = m.Rand.Float64() * totalAdaptationDegree

	// 累积适应度
	var sum float64

	// 遍历个体
	// 若pointer <= sum则说明轮盘指针落在当前个体所在区间
	for _, v := range m.Papers {
		sum += v.AdaptationDegree
		if pointer <= sum {
			return v
		}
	}

	return nil
}

// 交叉算子
func (m *Population) Crossover(p1, p2 *Paper) (*Paper, error) {
	var paper = new(Paper)
	paper.QuestionList = make([]*models.Question, len(p1.QuestionList))

	// 随机设置两个交叉点
	a := m.Rand.Intn(len(paper.QuestionList))
	b := m.Rand.Intn(len(paper.QuestionList))
	if a > b {
		a, b = b, a
	}

	// 遗传p1的试题
	for i := 0; i < a; i++ {
		paper.QuestionList[i] = p1.QuestionList[i]
	}
	for i := b; i < len(paper.QuestionList); i++ {
		paper.QuestionList[i] = p1.QuestionList[i]
	}

	// 交叉点a到b遗传p2的试题
	for i := a; i <= b; i++ {
		// 检查试题是否重复
		// 若重复则需要从试题库中重新抽题
		if !paper.CheckQuestionIsExist(p2.QuestionList[i].ID) {
			paper.QuestionList[i] = p2.QuestionList[i]
			continue
		}

		// 重新获取试题
		list, _, err := models.ReadQuestionSimpleListRaw(models.ReadQuestionSimpleListParam{
			SubjectID:       m.Rule.SubjectID,
			Type:            p2.QuestionList[i].Type,
			Status:          models.QUESTION_ENABLE,
			KnowledgeIDList: m.Rule.Points,
			ExcludeIDList:   paper.GetQuestionIDList(),
			BaseQueryParam: models.BaseQueryParam{
				Limit: 1,
			},
		})

		if err != nil {
			fmt.Printf("获取试题失败: %s\n", err.Error())
			return nil, err
		}
		if len(list) == 0 {
			return nil, fmt.Errorf("未查询到符合条件的试题，组卷失败")
		}
		paper.QuestionList[i] = list[0]
		paper.QuestionList[i].Score = m.Rule.getScore(paper.QuestionList[i].Type)
	}

	return paper, nil
}

// 变异算子
func (m *Population) Mutation(paper *Paper) error {
	for k, v := range paper.QuestionList {
		if m.Rand.Float64() < m.Rule.MutationRate {
			// 从题库中重新抽题
			list, _, err := models.ReadQuestionSimpleListRaw(models.ReadQuestionSimpleListParam{
				SubjectID:       m.Rule.SubjectID,
				Type:            v.Type,
				Status:          models.QUESTION_ENABLE,
				KnowledgeIDList: m.Rule.Points,
				ExcludeIDList:   paper.GetQuestionIDList(),
				BaseQueryParam: models.BaseQueryParam{
					Limit: 1,
				},
			})
			if err != nil {
				return err
			}
			if len(list) == 0 {
				continue
			}
			paper.QuestionList[k] = list[0]
			paper.QuestionList[k].Score = m.Rule.getScore(paper.QuestionList[k].Type)
		}
	}

	return nil
}
