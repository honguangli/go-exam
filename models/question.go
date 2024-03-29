package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"strconv"
	"strings"
	"time"
)

// 试题表
type Question struct {
	ID           int     `orm:"column(id)" form:"id" json:"id"`
	SubjectID    int     `orm:"column(subject_id)" form:"subject_id" json:"subject_id"`
	Name         string  `orm:"column(name)" form:"name" json:"name"`
	Type         int     `orm:"column(type)" form:"type" json:"type"`
	Content      string  `orm:"column(content)" form:"content" json:"content"`
	Tips         string  `orm:"column(tips)" form:"tips" json:"tips"`
	Analysis     string  `orm:"column(analysis)" form:"analysis" json:"analysis"`
	Difficulty   float64 `orm:"column(difficulty)" form:"difficulty" json:"difficulty"`
	KnowledgeIds string  `orm:"column(knowledge_ids)" form:"knowledge_ids" json:"knowledge_ids"`
	Score        int     `orm:"column(score)" form:"score" json:"score"`
	Status       int     `orm:"column(status)" form:"status" json:"status"`
	CreateTime   int64   `orm:"column(create_time)" form:"create_time" json:"create_time"`
	UpdateTime   int64   `orm:"column(update_time)" form:"update_time" json:"update_time"`
	Memo         string  `orm:"column(memo)" form:"memo" json:"memo"`
}

// 类型
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

// 状态
const (
	QUESTION_STATUS_IGNORE = -1 // 忽略状态
	QUESTION_DISABLE       = 0  // 禁用
	QUESTION_ENABLE        = 1  // 正常
)

// 查询详情参数
type ReadQuestionDetailParam struct {
	ID int `json:"id"`
}

// 查询列表参数
type ReadQuestionListParam struct {
	BaseQueryParam
	SubjectID int    `json:"subject_id"`
	Name      string `json:"name"`
	Type      int    `json:"type"`
	Status    int    `json:"status"`
	ClosePage bool   `form:"close_page" json:"close_page"`
}

// 查询列表参数 简化版
type ReadQuestionSimpleListParam struct {
	BaseQueryParam
	SubjectID       int    `json:"subject_id"`
	Name            string `json:"name"`
	Type            int    `json:"type"`
	Status          int    `json:"status"`
	KnowledgeIDList []int  `json:"knowledge_id_list"`
	ExcludeIDList   []int  `json:"exclude_id_list"`
	ClosePage       bool   `form:"close_page" json:"close_page"`
}

// 创建参数
type InsertQuestionParam struct {
	Question Question          `json:"question"`
	Options  []*QuestionOption `json:"options"`
}

// 更新参数
type UpdateQuestionParam struct {
	Question Question          `json:"question"`
	Options  []*QuestionOption `json:"options"`
}

// 删除参数
type DeleteQuestionParam struct {
	ID   int   `json:"id"`
	List []int `json:"list"`
}

// 初始化
func init() {
	orm.RegisterModel(new(Question))
}

// 表名
func QuestionTBName() (name string) {
	return "question"
}

// 自定义表名
func (m *Question) TableName() (name string) {
	return QuestionTBName()
}

// 多字段索引
func (m *Question) TableIndex() [][]string {
	return [][]string{}
}

// 多字段唯一键
func (m *Question) TableUnique() [][]string {
	return [][]string{}
}

// 自定义引擎
func (m *Question) TableEngine() string {
	return "INNODB"
}

// 查询单个对象
func ReadQuestionOne(id int) (m Question, err error) {
	o := orm.NewOrm()
	m.ID = id
	err = o.Read(&m)
	return
}

// 查询多个对象
func ReadQuestionList(param ReadQuestionListParam) (list []*Question, total int64, err error) {
	list = make([]*Question, 0)
	query := orm.NewOrm().QueryTable(QuestionTBName())

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
func ReadQuestionListRaw(param ReadQuestionListParam) (list []*Question, total int64, err error) {
	list = make([]*Question, 0)
	var args = make([]interface{}, 0)
	var whereSql = "WHERE 1=1"

	if param.SubjectID > 0 {
		whereSql += " AND T0.subject_id = ?"
		args = append(args, param.SubjectID)
	}

	if len(param.Name) > 0 {
		whereSql += " AND T0.name LIKE ?"
		args = append(args, fmt.Sprintf("%%%s%%", param.Name))
	}

	if param.Type != QUESTION_TYPE_IGNORE {
		whereSql += " AND T0.type = ?"
		args = append(args, param.Type)
	}

	if param.Status != QUESTION_STATUS_IGNORE {
		whereSql += " AND T0.status = ?"
		args = append(args, param.Status)
	}

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
	var fields = "T0.`id`, T0.`subject_id`, T0.`name`, T0.`type`, T0.`content`, T0.`tips`, T0.`analysis`, T0.`difficulty`, T0.`knowledge_ids`, T0.`score`, T0.`status`, T0.`create_time`, T0.`update_time`, T0.`memo`"

	// 关联查询
	var relatedSql string

	// 连表查询
	var sql = fmt.Sprintf("SELECT %s FROM question AS T0 %s %s %s %s", fields, relatedSql, whereSql, orderSql, pageSql)

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
	var countSql = fmt.Sprintf("SELECT count(*) AS count FROM question AS T0 %s %s", relatedSql, whereSql)
	var count RawCount
	err = orm.NewOrm().Raw(countSql, args...).QueryRow(&count)
	if err != nil {
		return
	}
	total = count.Count
	return
}

// 查询多个对象
func ReadQuestionSimpleListRaw(param ReadQuestionSimpleListParam) (list []*Question, total int64, err error) {
	list = make([]*Question, 0)
	var args = make([]interface{}, 0)
	var whereSql = "WHERE 1=1"

	if param.SubjectID > 0 {
		whereSql += " AND T0.subject_id = ?"
		args = append(args, param.SubjectID)
	}

	if len(param.Name) > 0 {
		whereSql += " AND T0.name LIKE ?"
		args = append(args, fmt.Sprintf("%%%s%%", param.Name))
	}

	if param.Type != QUESTION_TYPE_IGNORE {
		whereSql += " AND T0.type = ?"
		args = append(args, param.Type)
	}

	if param.Status != QUESTION_STATUS_IGNORE {
		whereSql += " AND T0.status = ?"
		args = append(args, param.Status)
	}

	if len(param.ExcludeIDList) > 0 {
		whereSql += fmt.Sprintf(" AND T0.id NOT IN (?%s)", strings.Repeat(", ?", len(param.ExcludeIDList)-1))
		args = append(args, param.ExcludeIDList)
	}

	if len(param.KnowledgeIDList) > 0 {
		whereSql += fmt.Sprintf(" AND (FIND_IN_SET(?, T0.knowledge_ids)%s)", strings.Repeat(" OR FIND_IN_SET(?, T0.knowledge_ids)", len(param.KnowledgeIDList)-1))
		for _, v := range param.KnowledgeIDList {
			args = append(args, strconv.Itoa(v))
		}
	}

	// 排序
	var orderSql = "ORDER BY RAND()"

	// 分页
	var pageSql string
	if !param.ClosePage {
		pageSql = fmt.Sprintf("LIMIT %d OFFSET %d", param.Limit, param.Offset)
	}

	// 查询字段
	var fields = "T0.`id`, T0.`subject_id`, T0.`type`, T0.`difficulty`, T0.`knowledge_ids`"

	// 关联查询
	var relatedSql string

	// 连表查询
	var sql = fmt.Sprintf("SELECT %s FROM question AS T0 %s %s %s %s", fields, relatedSql, whereSql, orderSql, pageSql)

	// 查询列表
	total, err = orm.NewOrm().Raw(sql, args...).QueryRows(&list)
	if err != nil {
		return
	}
	return
}

// 插入单个对象
func InsertQuestionOne(m Question) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(&m)
	return
}

// 插入多个对象
func InsertQuestionMulti(list []Question) (num int64, err error) {
	o := orm.NewOrm()
	num, err = o.InsertMulti(100, list)
	return
}

// 更新单个对象
func UpdateQuestionOne(m Question, fields ...string) (num int64, err error) {
	o := orm.NewOrm()
	if len(fields) == 0 {
		fields = []string{"subject_id", "name", "type", "content", "tips", "analysis", "difficulty", "knowledge_ids", "score", "status", "create_time", "update_time", "memo"}
	}
	num, err = o.Update(&m, fields...)
	return
}

// 更新多个对象
func UpdateQuestionMulti(ids []int, params map[string]interface{}) (num int64, err error) {
	o := orm.NewOrm()
	num, err = o.QueryTable(QuestionTBName()).Filter("id__in", ids).Update(params)
	return
}

// 删除单个对象
func DeleteQuestionOne(id int) (num int64, err error) {
	o := orm.NewOrm()
	num, err = o.Delete(&Question{ID: id})
	return
}

// 删除多个对象
func DeleteQuestionMulti(ids []int) (num int64, err error) {
	o := orm.NewOrm()
	num, err = o.QueryTable(QuestionTBName()).Filter("id__in", ids).Delete()
	return
}

// 插入单个对象
func InsertQuestionOneWithOptions(param InsertQuestionParam) (id int64, err error) {
	o := orm.NewOrm()

	err = o.Begin()
	if err != nil {
		return
	}

	// 保存试题
	param.Question.Status = QUESTION_ENABLE
	param.Question.CreateTime = time.Now().Unix()
	id, err = o.Insert(&param.Question)
	if err != nil {
		o.Rollback()
		return
	}

	// 保存选项
	if len(param.Options) > 0 {
		for _, v := range param.Options {
			v.ID = 0
			v.QuestionID = param.Question.ID
		}
		_, err = o.InsertMulti(100, param.Options)
		if err != nil {
			o.Rollback()
			return
		}
	}

	err = o.Commit()
	return
}

// 更新单个对象
func UpdateQuestionOneWithOptions(param UpdateQuestionParam) (num int64, err error) {
	o := orm.NewOrm()

	err = o.Begin()
	if err != nil {
		return
	}

	// 更新试题
	param.Question.UpdateTime = time.Now().Unix()
	num, err = o.Update(&param.Question, "subject_id", "name", "type", "content", "tips", "analysis", "difficulty", "knowledge_ids", "score", "status", "update_time", "memo")
	if err != nil {
		o.Rollback()
		return
	}

	// 删除选项
	_, err = o.QueryTable(QuestionOptionTBName()).Filter("question_id", param.Question.ID).Delete()
	if err != nil {
		o.Rollback()
		return
	}

	// 保存选项
	if len(param.Options) > 0 {
		for _, v := range param.Options {
			v.ID = 0
			v.QuestionID = param.Question.ID
		}
		_, err = o.InsertMulti(100, param.Options)
		if err != nil {
			o.Rollback()
			return
		}
	}

	err = o.Commit()
	return
}
