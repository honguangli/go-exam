package ga

import (
	"encoding/json"
	"fmt"
	"math/rand"
)

// 试题
type Question struct {
	ID         int     `json:"id"`         // id
	Type       int     `json:"type"`       // 类型
	Difficulty float64 `json:"difficulty"` // 难度系数
	Points     []int   `json:"points"`     // 知识点
	Score      int     `json:"score"`      // 分数
}

const (
	QUESTION_TYPE_IGNORE   = iota // 忽略类型筛选
	QUESTION_CHOICE_SINGLE        // 单选题
	QUESTION_CHOICE_MULTI         // 多选题
	QUESTION_JUDGE                // 判断题
	QUESTION_BLANK_SINGLE         // 填空题
	QUESTION_BLANK_MULTI          // 多项填空题
	QUESTION_ANSWER               // 简答题
	QUESTION_ANSWER_MULTI         // 多项简答题
	QUESTION_FILE_SINGLE          // 文件题
	QUESTION_FILE_MULTI           // 多项文件题
)

// 查询试题列表 参数
type QueryQuestionListParam struct {
	Type    int   // 类型
	Points  []int // 知识点
	Exclude []int // 排除id
	Limit   int   // 数量
}

var questionMap = make(map[int][]*Question)

func init() {
	var questionList = make([]*Question, 0)
	err := json.Unmarshal([]byte(questionData), &questionList)
	if err != nil {
		fmt.Printf("解析题库失败: %s\n", err.Error())
		return
	}

	for _, v := range questionList {
		questionMap[v.Type] = append(questionMap[v.Type], v)
	}
}

// 查询试题列表
func QueryQuestionList(param QueryQuestionListParam) (list []*Question, total int64, err error) {
	list = make([]*Question, 0)

	questionList, ok := questionMap[param.Type]
	if !ok {
		return nil, 0, fmt.Errorf("题库不足，组卷失败")
	}

	var pointsMap = make(map[int]struct{})
	for _, id := range param.Points {
		pointsMap[id] = struct{}{}
	}
	var excludeIDMap = make(map[int]struct{})
	for _, id := range param.Exclude {
		excludeIDMap[id] = struct{}{}
	}

	for _, v := range questionList {
		// 排除id（交集）
		if _, ok := excludeIDMap[v.ID]; ok {
			continue
		}

		// 判断知识点（交集）
		var use bool
		for _, id := range v.Points {
			if _, ok := pointsMap[id]; ok {
				use = true
				break
			}
		}
		if !use {
			continue
		}

		list = append(list, v)
	}

	if len(list) < param.Limit {
		return nil, 0, fmt.Errorf("题库数量不足，组卷失败")
	}

	swap(list)

	return list[:param.Limit], int64(param.Limit), nil
}

// 交换位置
func swap(list []*Question) {
	for i := 0; i < len(list); i++ {
		j := rand.Intn(i + 1)
		list[i], list[j] = list[j], list[i]
	}
}
