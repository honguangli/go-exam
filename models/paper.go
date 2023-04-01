package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

// 试卷表
type Paper struct {
	ID                int     `orm:"column(id)" form:"id" json:"id"`
	Name              string  `orm:"column(name)" form:"name" json:"name"`
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

// 查询详情参数
type ReadPaperDetailParam struct {
	ID int `json:"id"`
}

// 查询列表参数
type ReadPaperListParam struct {
	BaseQueryParam
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
	var fields = "T0.`id`, T0.`name`, T0.`score`, T0.`pass_score`, T0.`difficulty`, T0.`choice_single_num`, T0.`choice_single_score`, T0.`choice_multi_num`, T0.`choice_multi_score`, T0.`judge_num`, T0.`judge_score`, T0.`blank_single_num`, T0.`blank_single_score`, T0.`blank_multi_num`, T0.`blank_multi_score`, T0.`answer_single_num`, T0.`answer_single_score`, T0.`answer_multi_num`, T0.`answer_multi_score`, T0.`status`, T0.`create_time`, T0.`update_time`, T0.`memo`"

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
		fields = []string{"name", "score", "pass_score", "difficulty", "choice_single_num", "choice_single_score", "choice_multi_num", "choice_multi_score", "judge_num", "judge_score", "blank_single_num", "blank_single_score", "blank_multi_num", "blank_multi_score", "answer_single_num", "answer_single_score", "answer_multi_num", "answer_multi_score", "status", "create_time", "update_time", "memo"}
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
