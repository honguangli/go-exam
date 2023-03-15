package ga

import (
	"encoding/json"
	"fmt"
	"math"
	"math/rand"
	"testing"
	"time"
)

func TestGenQuestionList(t *testing.T) {
	var maxPointId = 30

	var list = make([]*Question, 0)

	// 生成试题数量 v0: 最小题量 v1：最大题量 v2：当前题量
	var ms = map[int][]int{
		QUESTION_CHOICE_SINGLE: []int{1200},
		QUESTION_CHOICE_MULTI:  []int{1000},
		QUESTION_JUDGE:         []int{800},
		QUESTION_BLANK_SINGLE:  []int{600},
		QUESTION_BLANK_MULTI:   []int{500},
		QUESTION_ANSWER:        []int{300},
		QUESTION_ANSWER_MULTI:  []int{100},
	}

	rand.Seed(time.Now().UnixNano())
	for k, v := range ms {
		for i := 0; i < v[0]; i++ {
			var points = rand.Perm(maxPointId)
			list = append(list, &Question{
				ID:         0,                                      // id
				Type:       k,                                      // 题目类型
				Difficulty: float64(int(rand.Float64()*100)) / 100, // 难度系数
				Points:     points[:rand.Intn(3)+1],                // 知识点
			})
		}
	}

	ids := rand.Perm(len(list))
	for k, v := range list {
		v.ID = ids[k]
	}

	fmt.Println("成功", len(list))
	bf, _ := json.Marshal(list)
	fmt.Println(string(bf))
}

func TestGA(t *testing.T) {
	// 创建规则
	var rule = Rule{
		Size:              50,                                           // 初始种群大小
		MutationRate:      0.085,                                        // 变异概率
		Elitism:           true,                                         // 精英主义
		Score:             100,                                          // 试卷总分
		Difficulty:        0.85,                                         // 试卷期望难度系数
		DifficultyWeight:  0.80,                                         // 难度权重
		Points:            []int{1, 3, 4, 6, 9, 12, 15, 20, 23, 24, 26}, // 试卷期望包含知识点
		PointsWeight:      0.20,                                         // 知识点权重
		ChoiceSingleNum:   6,                                            // 单选题数量
		ChoiceSingleScore: 4,                                            // 单选题分值
		ChoiceMultiNum:    4,                                            // 多选题数量
		ChoiceMultiScore:  6,                                            // 多选题分值
		JudgeNum:          6,                                            // 判断题数量
		JudgeScore:        2,                                            // 判断题分值
		BlankSingleNum:    0,                                            // 填空题数量
		BlankSingleScore:  0,                                            // 填空题分值
		BlankMultiNum:     0,                                            // 多项填空题数量
		BlankMultiScore:   0,                                            // 多项填空题分值
		AnswerSingleNum:   4,                                            // 简答题数量
		AnswerSingleScore: 10,                                           // 简答题分值
		AnswerMultiNum:    0,                                            // 多项简答题数量
		AnswerMultiScore:  0,                                            // 多项简答题分值
	}

	// 初始化种群
	population, err := NewPopulation(&rule)
	if err != nil {
		fmt.Printf("初始化种群失败：%s\n", err.Error())
		return
	}

	//for k, v := range population.Papers {
	//	fmt.Printf("k = %d | d = %f | k = %f | fit = %f | calc = %f\n", k, v.Difficulty, v.KPCoverage, v.AdaptationDegree, 1-(1-v.KPCoverage)*rule.PointsWeight-math.Abs(rule.Difficulty-v.Difficulty)*rule.DifficultyWeight)
	//}

	// 迭代次数
	var count = 100

	// 期望适应度
	var expand = 0.98

	var paper *Paper
	for i := 0; i < count; i++ {
		paper = population.GetFitness()
		fmt.Printf("适应度：%f\n", paper.AdaptationDegree)

		if paper.AdaptationDegree >= expand {
			break
		}

		fmt.Printf("第%d次进化\n", i+1)
		err = population.Evolve()
		if err != nil {
			fmt.Printf("进化失败: %s\n", err.Error())
			return
		}
		//for k, v := range population.Papers {
		//	fmt.Printf("k = %d | d = %f(%f) | k = %f(%f) | fit = %f(%f)\n", k, v.Difficulty, v.GetDifficulty(), v.KPCoverage, v.GetKpCoverage(rule.Points), v.AdaptationDegree, v.GetAdaptationDegree(rule.Difficulty, rule.PointsWeight, rule.DifficultyWeight))
		//}
	}
	fmt.Printf("进化完成: 适应度：%f\n", paper.AdaptationDegree)
	bf, err := json.Marshal(paper)
	if err != nil {
		fmt.Printf("json marshal error: %s", err.Error())
	} else {
		fmt.Printf("paper: %s\n", string(bf))
	}

	paper.SetDifficulty()
	paper.SetKpCoverage(rule.Points)
	paper.SetAdaptationDegree(rule.Difficulty, rule.PointsWeight, rule.DifficultyWeight)
	bf, err = json.Marshal(paper)
	if err != nil {
		fmt.Printf("json marshal error: %s", err.Error())
	} else {
		fmt.Printf("paper: %s\n", string(bf))
	}

	fmt.Printf("d1 = %f\n", 1-(1-paper.KPCoverage)*rule.PointsWeight-math.Abs(rule.Difficulty-paper.Difficulty)*rule.DifficultyWeight)
	fmt.Printf("d2 = %f\n", 1/(1-(1-paper.KPCoverage)*rule.PointsWeight-math.Abs(rule.Difficulty-paper.Difficulty)*rule.DifficultyWeight))
}
