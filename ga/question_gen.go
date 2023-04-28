package ga

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"go-exam/models"
	"io/ioutil"
	"math/rand"
	"regexp"
	"strings"
	"testing"
	"time"
)

func TestSplitQuestion(t *testing.T) {
	b, err := ioutil.ReadFile("./doc/试题.txt")
	if err != nil {
		fmt.Printf("read file error: %s", err.Error())
		return
	}

	data := string(b)
	data = splitByQuestionSort(data)
	dataMap := splitByQuestionType(data)

	for k, v := range dataMap {
		err = ioutil.WriteFile(".\\data\\"+k+".txt", []byte(v), 0666)
		if err != nil {
			fmt.Printf("write file error: %s\n", err.Error())
		}
	}
}

// 按题号切分
func splitByQuestionSort(data string) string {
	re := regexp.MustCompile(" [0-9]*\\.") // 路由文本搜索正则
	found := re.FindAllStringIndex(data, -1)
	if found == nil || len(found) == 0 {
		fmt.Println("splitByQuestionSort: not found string")
		return data
	}
	fmt.Printf("splitByQuestionSort: %v\n", found)

	for i := len(found) - 1; i >= 0; i-- {
		data = data[:found[i][0]] + " \n\r" + data[found[i][0]:]
	}
	return data
}

// 按题型切分
func splitByQuestionType(data string) map[string]string {
	re := regexp.MustCompile(" *[一二三四五六七八九](、|.) ?(选择|判断|填空|简答|是非|简单|综合)题 *") // 路由文本搜索正则
	found := re.FindAllStringIndex(data, -1)
	if found == nil || len(found) == 0 {
		fmt.Println("splitByQuestionType: not found string")
		return map[string]string{"其他": data}
	}
	fmt.Printf("splitByQuestionType: %v\n", found)

	var dataMap = map[string]string{
		"选择": "",
		"判断": "", // 判断、是非
		"填空": "",
		"简答": "", // 简答、简单
		"综合": "",
		"其他": "", // 其他
	}
	var dataMapSize = make(map[string]int)
	for i := len(found) - 1; i >= 0; i-- {
		var filterName = data[found[i][0]:found[i][1]]
		var filterStr = "\n\r" + data[found[i][0]:found[i][1]] + " \n\r" + data[found[i][1]:]
		if strings.Contains(filterName, "选择") {
			dataMap["选择"] = filterStr + dataMap["选择"]
			dataMapSize["选择"]++
		} else if strings.Contains(filterName, "判断") {
			dataMap["判断"] = filterStr + dataMap["判断"]
			dataMapSize["判断"]++
		} else if strings.Contains(filterName, "填空") {
			dataMap["填空"] = filterStr + dataMap["填空"]
			dataMapSize["填空"]++
		} else if strings.Contains(filterName, "简答") {
			dataMap["简答"] = filterStr + dataMap["简答"]
			dataMapSize["简答"]++
		} else if strings.Contains(filterName, "是非") {
			dataMap["判断"] = filterStr + dataMap["判断"]
			dataMapSize["判断"]++
		} else if strings.Contains(filterName, "简单") {
			dataMap["简答"] = filterStr + dataMap["简答"]
			dataMapSize["简答"]++
		} else if strings.Contains(filterName, "综合") {
			dataMap["综合"] = filterStr + dataMap["综合"]
			dataMapSize["综合"]++
		} else {
			dataMap["其他"] = filterStr + dataMap["其他"]
			dataMapSize["其他"]++
		}
		data = data[:found[i][0]]

		//data = data[:found[i][0]] + "\n\r" + data[found[i][0]:found[i][1]] + " \n\r" + data[found[i][1]:]
	}
	var size int
	for k, v := range dataMapSize {
		size += v
		fmt.Printf("splitByQuestionType: %s | %d\n", k, v)
	}
	fmt.Printf("splitByQuestionType: %d | %d\n", len(found), size)
	return dataMap
}

// 试题
type q struct {
	Type        int                 `json:"type"`
	Content     string              `json:"content"`
	Analysis    string              `json:"analysis"`
	Options     []*o                `json:"options"`
	OptionRight map[string]struct{} `json:"option_right"`
}

// 选项
type o struct {
	Tag     string `json:"tag"`
	Content string `json:"content"`
	IsRight int    `json:"is_right"`
}

// 解析选择题
func TestTransformChoice(t *testing.T) {

}
func TransformChoice() {
	if b, err := ioutil.ReadFile("./doc/选择题.txt"); err != nil {
		logs.Info("read choice file error: %s", err.Error())
	} else {
		transformChoiceQuestion(string(b))
	}

	if b, err := ioutil.ReadFile("./doc/判断题.txt"); err != nil {
		logs.Info("read judge file error: %s", err.Error())
	} else {
		transformJudgeQuestion(string(b))
	}
	return
}

// 解析选择题
func transformChoiceQuestion(data string) {
	// 正则表达式
	// 单行题干：([0-9]+\..+\n)(【[A-Z]+】.*\n)([A-Z]\..*\n)*
	// 多行题干：([0-9]+\.(.+\n)+)(【[A-Z]+】.*\n)([A-Z]\..*\n)*
	// 切分试题列表
	re := regexp.MustCompile("([0-9]+\\..+\\n)(【[A-Z]+】.*\\n)([A-Z]\\..*\\n)*") // 试题文本搜索正则
	found := re.FindAllString(data, -1)
	if len(found) == 0 {
		logs.Info("transformChoiceQuestion: not found string")
		return
	}
	logs.Info("transformChoiceQuestion: len = %d", len(found))

	var questionList = make([]*q, 0)
	for k, v := range found {
		//logs.Info("k = %d | v = %s", k, v)
		lines := strings.Split(v, "\n")

		// 解析题目
		// 去掉题干前缀的题号、符号、空格以及后缀空格
		var questionContent = strings.TrimSpace(lines[0][strings.Index(lines[0], ".")+1:])

		// 解析答案
		// 截取"【"和"】"间的文本
		var optionRights = strings.Split(lines[1][strings.Index(lines[1], "【")+len("【"):strings.Index(lines[1], "】")], "")
		var optionRightMap = make(map[string]struct{})
		for _, r := range optionRights {
			optionRightMap[r] = struct{}{}
		}

		// 解析选项列表
		// 选项tag：截取前缀
		// 选项内容：去掉选项前缀的tag、.、空格以及后缀空格
		// 是否正确选项：根据答案匹配
		var options = make([]*o, 0)
		var rightSize int
		for i := 2; i < len(lines)-1; i++ {
			var optionTag = strings.TrimSpace(lines[i][:strings.Index(lines[i], ".")])
			var optionContent = strings.TrimSpace(lines[i][strings.Index(lines[i], ".")+1:])
			var optionIsRight = models.QUESTION_OPTION_IS_NOT_RIGHT
			if _, ok := optionRightMap[optionTag]; ok {
				optionIsRight = models.QUESTION_OPTION_IS_RIGHT
				rightSize++
			}
			options = append(options, &o{
				Tag:     optionTag,
				Content: optionContent,
				IsRight: optionIsRight,
			})
		}
		if rightSize == 0 {
			logs.Info("transformChoiceQuestion: not right option, k = %d | content = %s | or = %v | lp = %d | rp = %d \n%s",
				k, questionContent, optionRights, strings.Index(lines[1], "【")+1, strings.Index(lines[1], "】"), v)
		}

		// 判断题目类型
		// 单选题或多选题
		var questionType = models.QUESTION_CHOICE_SINGLE
		if rightSize > 1 {
			questionType = models.QUESTION_CHOICE_MULTI
		}

		questionList = append(questionList, &q{
			Type:        questionType,
			Content:     questionContent,
			Analysis:    "",
			Options:     options,
			OptionRight: optionRightMap,
		})
	}

	//for k, v := range questionList {
	//	logs.Info("k = %d | type = %d | %s | len(options) = %d | %v", k, v.Type, v.Content, len(v.Options), v.OptionRight)
	//	for _, vv := range v.Options {
	//		logs.Info("right = %t | %s、%s", vv.IsRight == 1, vv.Tag, vv.Content)
	//	}
	//}
	//insertOption(questionList)
}

// 解析判断题
func transformJudgeQuestion(data string) {
	// 正则表达式
	// ([0-9]+\..+\n)(【(True|False)】.*\n)
	// 切分试题列表
	re := regexp.MustCompile("([0-9]+\\..+\\n)(【(True|False)】.*\\n)") // 试题文本搜索正则
	found := re.FindAllString(data, -1)
	if len(found) == 0 {
		logs.Info("transformJudgeQuestion: not found string")
		return
	}
	logs.Info("transformJudgeQuestion: len = %d", len(found))

	var questionList = make([]*q, 0)
	for _, v := range found {
		//logs.Info("k = %d | v = %s", k, v)
		lines := strings.Split(v, "\n")

		// 解析题目
		// 去掉题干前缀的题号、符号、空格以及后缀空格
		var questionContent = strings.TrimSpace(lines[0][strings.Index(lines[0], ".")+1:])

		var optionRightMap = make(map[string]struct{})

		// 生成选项
		var options = []*o{
			&o{
				Tag:     "A",
				Content: "正确",
				IsRight: models.QUESTION_OPTION_IS_NOT_RIGHT,
			},
			&o{
				Tag:     "B",
				Content: "错误",
				IsRight: models.QUESTION_OPTION_IS_NOT_RIGHT,
			},
		}
		if !strings.Contains(lines[1], "True") {
			optionRightMap["False"] = struct{}{}
			options[1].IsRight = models.QUESTION_OPTION_IS_RIGHT
		} else {
			optionRightMap["True"] = struct{}{}
			options[0].IsRight = models.QUESTION_OPTION_IS_RIGHT
		}

		questionList = append(questionList, &q{
			Type:        models.QUESTION_JUDGE,
			Content:     questionContent,
			Analysis:    "",
			Options:     options,
			OptionRight: optionRightMap,
		})
	}

	//for k, v := range questionList {
	//	logs.Info("k = %d | type = %d | %s | len(options) = %d | %v", k, v.Type, v.Content, len(v.Options), v.OptionRight)
	//	for _, vv := range v.Options {
	//		logs.Info("right = %t | %s、%s", vv.IsRight == 1, vv.Tag, vv.Content)
	//	}
	//}
	// insertOption(questionList)
}

func insertOption(ql []*q) {
	o := orm.NewOrm()
	err := o.Begin()
	if err != nil {
		logs.Info("begin error: %s", err.Error())
		return
	}

	nowUnix := time.Now().Unix()
	rand.Seed(time.Now().UnixNano())

	for _, v := range ql {
		rp := len(v.Content)
		if rp > 500 {
			rp = 500
		}
		var question = models.Question{
			ID:           0,
			SubjectID:    0,
			Name:         v.Content[:rp],
			Type:         v.Type,
			Content:      v.Content,
			Tips:         "",
			Analysis:     v.Analysis,
			Difficulty:   float64(rand.Intn(50)+38) / 100,
			KnowledgeIds: "",
			Score:        4,
			Status:       models.QUESTION_ENABLE,
			CreateTime:   nowUnix,
			UpdateTime:   0,
			Memo:         "",
		}
		id, err := o.Insert(&question)
		if err != nil {
			logs.Info("insert question error: %s", err.Error())
			o.Rollback()
			return
		}

		var options = make([]*models.QuestionOption, 0)
		for _, vv := range v.Options {
			options = append(options, &models.QuestionOption{
				ID:         0,
				QuestionID: int(id),
				Tag:        vv.Tag,
				Content:    vv.Content,
				IsRight:    vv.IsRight,
				Memo:       "",
			})
		}
		if len(options) > 0 {
			_, err = o.InsertMulti(100, options)
			if err != nil {
				logs.Info("insert options error: %s", err.Error())
				o.Rollback()
				return
			}
		}
	}

	err = o.Commit()
	if err != nil {
		logs.Info("commit error: %s", err.Error())
	}
}
