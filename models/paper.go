package models

import (
	"database/sql"
	"fmt"
	"github.com/astaxie/beego/orm"
	"strings"
	"time"
)

// 试卷表
type Paper struct {
	ID                int     `orm:"column(id)" form:"id" json:"id"`
	Name              string  `orm:"column(name)" form:"name" json:"name"`
	SubjectID         int     `orm:"column(subject_id)" form:"subject_id" json:"subject_id"`
	KnowledgeIds      string  `orm:"column(knowledge_ids)" form:"knowledge_ids" json:"knowledge_ids"`
	Score             int     `orm:"column(score)" form:"score" json:"score"`
	PassScore         int     `orm:"column(pass_score)" form:"pass_score" json:"pass_score"`
	Difficulty        float64 `orm:"column(difficulty)" form:"difficulty" json:"difficulty"`
	ChoiceSingleNum   int     `orm:"column(choice_single_num)" form:"choice_single_num" json:"choice_single_num"`
	ChoiceSingleScore int     `orm:"column(choice_single_score)" form:"choice_single_score" json:"choice_single_score"`
	ChoiceMultiNum    int     `orm:"column(choice_multi_num)" form:"choice_multi_num" json:"choice_multi_num"`
	ChoiceMultiScore  int     `orm:"column(choice_multi_score)" form:"choice_multi_score" json:"choice_multi_score"`
	JudgeNum          int     `orm:"column(judge_num)" form:"judge_num" json:"judge_num"`
	JudgeScore        int     `orm:"column(judge_score)" form:"judge_score" json:"judge_score"`
	BlankSingleNum    int     `orm:"column(blank_single_num)" form:"blank_single_num" json:"blank_single_num"`
	BlankSingleScore  int     `orm:"column(blank_single_score)" form:"blank_single_score" json:"blank_single_score"`
	BlankMultiNum     int     `orm:"column(blank_multi_num)" form:"blank_multi_num" json:"blank_multi_num"`
	BlankMultiScore   int     `orm:"column(blank_multi_score)" form:"blank_multi_score" json:"blank_multi_score"`
	AnswerSingleNum   int     `orm:"column(answer_single_num)" form:"answer_single_num" json:"answer_single_num"`
	AnswerSingleScore int     `orm:"column(answer_single_score)" form:"answer_single_score" json:"answer_single_score"`
	AnswerMultiNum    int     `orm:"column(answer_multi_num)" form:"answer_multi_num" json:"answer_multi_num"`
	AnswerMultiScore  int     `orm:"column(answer_multi_score)" form:"answer_multi_score" json:"answer_multi_score"`
	Status            int     `orm:"column(status)" form:"status" json:"status"`
	CreateTime        int64   `orm:"column(create_time)" form:"create_time" json:"create_time"`
	UpdateTime        int64   `orm:"column(update_time)" form:"update_time" json:"update_time"`
	Memo              string  `orm:"column(memo)" form:"memo" json:"memo"`
}

// 状态
const (
	PAPER_EDIT   = 0 // 草稿
	PAPER_FREEZE = 1 // 已发布
)

// 查询详情参数
type ReadPaperDetailParam struct {
	ID int `json:"id"`
}

// 查询列表参数
type ReadPaperListParam struct {
	BaseQueryParam
	SubjectID int    `json:"subject_id"`
	Name      string `json:"name"`
	Status    int    `json:"status"`
	ClosePage bool   `form:"close_page" json:"close_page"`
}

// 删除参数
type DeletePaperParam struct {
	ID   int   `json:"id"`
	List []int `json:"list"`
}

// 初始化
func init() {
	orm.RegisterModel(new(Paper))
}

// 表名
func PaperTBName() (name string) {
	return "paper"
}

// 自定义表名
func (m *Paper) TableName() (name string) {
	return PaperTBName()
}

// 多字段索引
func (m *Paper) TableIndex() [][]string {
	return [][]string{}
}

// 多字段唯一键
func (m *Paper) TableUnique() [][]string {
	return [][]string{}
}

// 自定义引擎
func (m *Paper) TableEngine() string {
	return "INNODB"
}

// 查询单个对象
func ReadPaperOne(id int) (m Paper, err error) {
	o := orm.NewOrm()
	m.ID = id
	err = o.Read(&m)
	return
}

// 查询多个对象
func ReadPaperList(param ReadPaperListParam) (list []*Paper, total int64, err error) {
	list = make([]*Paper, 0)
	query := orm.NewOrm().QueryTable(PaperTBName())

	sortOrder := "id"
	switch param.Sort {
	}
	if param.Order == "desc" {
		sortOrder = "-" + sortOrder
	}

	total, err = query.Count()
	if err != nil {
		return
	}
	_, err = query.OrderBy(sortOrder).Limit(param.Limit, param.Offset).All(&list)
	return
}

// 查询多个对象
func ReadPaperListRaw(param ReadPaperListParam) (list []*Paper, total int64, err error) {
	list = make([]*Paper, 0)
	var args = make([]interface{}, 0)
	var whereSql = "WHERE 1=1"

	// 排序
	var orderSql = "ORDER BY "
	switch param.Sort {
	default:
		orderSql += "T0.id"
	}
	if param.Order == "desc" {
		orderSql += " DESC"
	}

	// 分页
	var pageSql string
	if !param.ClosePage {
		pageSql = fmt.Sprintf("LIMIT %d OFFSET %d", param.Limit, param.Offset)
	}

	// 查询字段
	var fields = "T0.`id`, T0.`name`, T0.`subject_id`, T0.`knowledge_ids`, T0.`score`, T0.`pass_score`, T0.`difficulty`, T0.`choice_single_num`, T0.`choice_single_score`, T0.`choice_multi_num`, T0.`choice_multi_score`, T0.`judge_num`, T0.`judge_score`, T0.`blank_single_num`, T0.`blank_single_score`, T0.`blank_multi_num`, T0.`blank_multi_score`, T0.`answer_single_num`, T0.`answer_single_score`, T0.`answer_multi_num`, T0.`answer_multi_score`, T0.`status`, T0.`create_time`, T0.`update_time`, T0.`memo`"

	// 关联查询
	var relatedSql string

	// 连表查询
	var sql = fmt.Sprintf("SELECT %s FROM paper AS T0 %s %s %s %s", fields, relatedSql, whereSql, orderSql, pageSql)

	// 查询列表
	total, err = orm.NewOrm().Raw(sql, args...).QueryRows(&list)
	if err != nil {
		return
	}

	// 关闭分页时不查询count
	if param.ClosePage {
		return
	}

	// 查询总数
	var countSql = fmt.Sprintf("SELECT count(*) AS count FROM paper AS T0 %s %s", relatedSql, whereSql)
	var count RawCount
	err = orm.NewOrm().Raw(countSql, args...).QueryRow(&count)
	if err != nil {
		return
	}
	total = count.Count
	return
}

// 插入单个对象
func InsertPaperOne(m Paper) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(&m)
	return
}

// 插入多个对象
func InsertPaperMulti(list []Paper) (num int64, err error) {
	o := orm.NewOrm()
	num, err = o.InsertMulti(100, list)
	return
}

// 更新单个对象
func UpdatePaperOne(m Paper, fields ...string) (num int64, err error) {
	o := orm.NewOrm()
	if len(fields) == 0 {
		fields = []string{"name", "subject_id", "knowledge_ids", "score", "pass_score", "difficulty", "choice_single_num", "choice_single_score", "choice_multi_num", "choice_multi_score", "judge_num", "judge_score", "blank_single_num", "blank_single_score", "blank_multi_num", "blank_multi_score", "answer_single_num", "answer_single_score", "answer_multi_num", "answer_multi_score", "status", "create_time", "update_time", "memo"}
	}
	num, err = o.Update(&m, fields...)
	return
}

// 更新多个对象
func UpdatePaperMulti(ids []int, params map[string]interface{}) (num int64, err error) {
	o := orm.NewOrm()
	num, err = o.QueryTable(PaperTBName()).Filter("id__in", ids).Update(params)
	return
}

// 删除单个对象
func DeletePaperOne(id int) (num int64, err error) {
	o := orm.NewOrm()
	num, err = o.Delete(&Paper{ID: id})
	return
}

// 删除多个对象
func DeletePaperMulti(ids []int) (num int64, err error) {
	o := orm.NewOrm()
	num, err = o.QueryTable(PaperTBName()).Filter("id__in", ids).Delete()
	return
}

// 智能组卷试卷
func InsertPaperOneWithQuestionList(m Paper, list []*Question) (id int64, err error) {
	o := orm.NewOrm()

	// 查询试卷试题
	var qid = make([]int, 0)
	for _, v := range list {
		qid = append(qid, v.ID)
	}
	var questionList = make([]*Question, 0)
	var optionList = make([]*QuestionOption, 0)
	var optionMap = make(map[int][]*QuestionOption, 0)
	_, err = o.Raw(fmt.Sprintf("SELECT * FROM question WHERE id IN (?%s)", strings.Repeat(", ?", len(qid)-1)), qid).QueryRows(&questionList)
	if err != nil {
		return
	}
	_, err = o.Raw(fmt.Sprintf("SELECT * FROM question_option WHERE question_id IN (?%s)", strings.Repeat(", ?", len(qid)-1)), qid).QueryRows(&optionList)
	if err != nil {
		return
	}
	for _, v := range optionList {
		if _, ok := optionMap[v.QuestionID]; !ok {
			optionMap[v.QuestionID] = []*QuestionOption{v}
		} else {
			optionMap[v.QuestionID] = append(optionMap[v.QuestionID], v)
		}
	}

	err = o.Begin()
	if err != nil {
		return
	}

	// 保存试卷
	m.Status = PAPER_EDIT
	m.CreateTime = time.Now().Unix()
	id, err = o.Insert(&m)
	if err != nil {
		o.Rollback()
		return
	}

	// 保存试题
	var paperQuestionOptionList = make([]*PaperQuestionOption, 0)
	for _, question := range questionList {
		var paperQuestion = PaperQuestion{
			ID:           0,
			PaperID:      m.ID,
			OriginID:     question.ID,
			SubjectID:    question.SubjectID,
			Name:         question.Name,
			Type:         question.Type,
			Content:      question.Content,
			Tips:         question.Tips,
			Analysis:     question.Analysis,
			Difficulty:   question.Difficulty,
			KnowledgeIds: question.KnowledgeIds,
			Score:        question.Score,
			UpdateTime:   question.UpdateTime,
			Memo:         question.Memo,
		}
		_, err = o.Insert(&paperQuestion)
		if err != nil {
			o.Rollback()
			return
		}
		if optionList, ok := optionMap[question.ID]; ok {
			for _, option := range optionList {
				paperQuestionOptionList = append(paperQuestionOptionList, &PaperQuestionOption{
					ID:         0,
					QuestionID: paperQuestion.ID,
					Tag:        option.Tag,
					Content:    option.Content,
					IsRight:    option.IsRight,
					Memo:       option.Memo,
				})
			}
		}
	}
	_, err = o.InsertMulti(100, paperQuestionOptionList)
	if err != nil {
		o.Rollback()
		return
	}

	err = o.Commit()
	return
}

// 智能组卷试卷
func InsertPaperOneWithQuestionListV2(m Paper, list []*Question) (id int64, err error) {
	o := orm.NewOrm()

	// 查询试卷试题
	var qidl = make([]int, 0)
	for _, v := range list {
		qidl = append(qidl, v.ID)
	}
	var optionList = make([]*QuestionOption, 0)
	var optionMap = make(map[int][]*QuestionOption, 0)
	_, err = o.Raw(fmt.Sprintf("SELECT * FROM question_option WHERE question_id IN (?%s)", strings.Repeat(", ?", len(qidl)-1)), qidl).QueryRows(&optionList)
	if err != nil {
		return
	}
	for _, v := range optionList {
		if _, ok := optionMap[v.QuestionID]; !ok {
			optionMap[v.QuestionID] = []*QuestionOption{v}
		} else {
			optionMap[v.QuestionID] = append(optionMap[v.QuestionID], v)
		}
	}

	err = o.Begin()
	if err != nil {
		return
	}

	// 保存试卷
	m.Status = PAPER_EDIT
	m.CreateTime = time.Now().Unix()
	id, err = o.Insert(&m)
	if err != nil {
		o.Rollback()
		return
	}

	// 保存试题
	var paperQuestionOptionList = make([]*PaperQuestionOption, 0)
	for _, qid := range qidl {
		var rawResult sql.Result
		rawResult, err = o.Raw("INSERT INTO paper_question (`paper_id`, `origin_id`, `subject_id`, `name`, `type`, `content`, `tips`, `analysis`, `difficulty`, `knowledge_ids`, `score`, `update_time`, `memo`)"+
			" SELECT ?, ?, T0.`subject_id`, T0.`name`, T0.`type`, T0.`content`, T0.`tips`, T0.`analysis`, T0.`difficulty`, T0.`knowledge_ids`, T0.`score`, T0.`update_time`, T0.`memo` FROM question AS T0"+
			" WHERE id = ?", m.ID, qid, qid).Exec()
		if err != nil {
			o.Rollback()
			return
		}
		var newID int64
		newID, err = rawResult.LastInsertId()
		if err != nil {
			o.Rollback()
			return
		}
		if optionList, ok := optionMap[qid]; ok {
			for _, option := range optionList {
				paperQuestionOptionList = append(paperQuestionOptionList, &PaperQuestionOption{
					ID:         0,
					QuestionID: int(newID),
					Tag:        option.Tag,
					Content:    option.Content,
					IsRight:    option.IsRight,
					Memo:       option.Memo,
				})
			}
		}
	}
	_, err = o.InsertMulti(100, paperQuestionOptionList)
	if err != nil {
		o.Rollback()
		return
	}

	err = o.Commit()
	return
}
