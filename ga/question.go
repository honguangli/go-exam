package ga

//// 试题
//type Question struct {
//	ID           int     `orm:"column(id)" form:"id" json:"id"`
//	SubjectID    int     `orm:"column(subject_id)" form:"subject_id" json:"subject_id"`
//	Name         string  `orm:"column(name)" form:"name" json:"name"`
//	Type         int     `orm:"column(type)" form:"type" json:"type"`
//	Content      string  `orm:"column(content)" form:"content" json:"content"`
//	Tips         string  `orm:"column(tips)" form:"tips" json:"tips"`
//	Analysis     string  `orm:"column(analysis)" form:"analysis" json:"analysis"`
//	Difficulty   float64 `orm:"column(difficulty)" form:"difficulty" json:"difficulty"`
//	KnowledgeIds string  `orm:"column(knowledge_ids)" form:"knowledge_ids" json:"knowledge_ids"`
//	Score        int     `orm:"column(score)" form:"score" json:"score"`
//	Status       int     `orm:"column(status)" form:"status" json:"status"`
//	CreateTime   int64   `orm:"column(create_time)" form:"create_time" json:"create_time"`
//	UpdateTime   int64   `orm:"column(update_time)" form:"update_time" json:"update_time"`
//	Memo         string  `orm:"column(memo)" form:"memo" json:"memo"`
//	Points       []int   `json:"points"` // 知识点
//}

//const (
//	QUESTION_TYPE_IGNORE   = iota // 忽略类型筛选
//	QUESTION_CHOICE_SINGLE        // 单选题
//	QUESTION_CHOICE_MULTI         // 多选题
//	QUESTION_JUDGE                // 判断题
//	QUESTION_BLANK_SINGLE         // 填空题
//	QUESTION_BLANK_MULTI          // 多项填空题
//	QUESTION_ANSWER               // 简答题
//	QUESTION_ANSWER_MULTI         // 多项简答题
//	QUESTION_FILE_SINGLE          // 文件题
//	QUESTION_FILE_MULTI           // 多项文件题
//)
//
//// 查询试题列表 参数
//type QueryQuestionListParam struct {
//	Type    int   // 类型
//	Points  []int // 知识点
//	Exclude []int // 排除id
//	Limit   int   // 数量
//}

//var questionMap = make(map[int][]*Question)
//
//func init() {
//	var questionList = make([]*Question, 0)
//	err := json.Unmarshal([]byte(questionData), &questionList)
//	if err != nil {
//		fmt.Printf("解析题库失败: %s\n", err.Error())
//		return
//	}
//
//	for _, v := range questionList {
//		questionMap[v.Type] = append(questionMap[v.Type], v)
//	}
//}

//// 查询试题列表
//func QueryQuestionList(param QueryQuestionListParam) (list []*Question, total int64, err error) {
//	list = make([]*Question, 0)
//
//	questionList, ok := questionMap[param.Type]
//	if !ok {
//		return nil, 0, fmt.Errorf("题库不足，组卷失败")
//	}
//
//	var pointsMap = make(map[int]struct{})
//	for _, id := range param.Points {
//		pointsMap[id] = struct{}{}
//	}
//	var excludeIDMap = make(map[int]struct{})
//	for _, id := range param.Exclude {
//		excludeIDMap[id] = struct{}{}
//	}
//
//	for _, v := range questionList {
//		// 排除id（交集）
//		if _, ok := excludeIDMap[v.ID]; ok {
//			continue
//		}
//
//		// 判断知识点（交集）
//		var use bool
//		for _, id := range v.Points {
//			if _, ok := pointsMap[id]; ok {
//				use = true
//				break
//			}
//		}
//		if !use {
//			continue
//		}
//
//		list = append(list, v)
//	}
//
//	if len(list) < param.Limit {
//		return nil, 0, fmt.Errorf("题库数量不足，组卷失败")
//	}
//
//	swap(list)
//
//	return list[:param.Limit], int64(param.Limit), nil
//}
//
//// 交换位置
//func swap(list []*Question) {
//	for i := 0; i < len(list); i++ {
//		j := rand.Intn(i + 1)
//		list[i], list[j] = list[j], list[i]
//	}
//}
