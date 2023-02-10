package utils

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego/logs"
)

/**
 * 时间工具
 */
const TimeFormatLayout = "2006-01-02 15:04:05"

type timeUtil struct {
}

var TimeUtil timeUtil

// 时间格式化
func (timeUtil) Format(t time.Time, layout ...string) string {
	if len(layout) > 0 {
		format := layout[0]
		format = strings.Replace(format, "YYYY", "2006", -1)
		format = strings.Replace(format, "MM", "01", -1)
		format = strings.Replace(format, "DD", "02", -1)
		format = strings.Replace(format, "HH", "15", -1)
		format = strings.Replace(format, "mm", "04", -1)
		format = strings.Replace(format, "ss", "05", -1)
		return t.Format(format)
	}
	return t.Format(TimeFormatLayout)
}

// 获取零点 00:00:00
func (timeUtil) StartTime(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

// 获取末点 23:59:59
func (timeUtil) EndTime(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 1e9-1, t.Location())
}

/**
 * 金额工具
 */
type moneyUtil struct {
}

var MoneyUtil moneyUtil

// 格式化金额 分转换为元
// @param cents 金额，单位分
func (moneyUtil) Cents2Yuan(cents int64) string {
	return strconv.FormatFloat(float64(cents)/100, 'f', 2, 64)
}

// 格式化金额 元转换为分
// @param yuan 金额，单位元
func (moneyUtil) Yuan2Cents(yuan string) (int64, error) {
	f, err := strconv.ParseFloat(yuan, 64)
	return int64(f * 100), err
}

// 格式化折扣
// @param discount 折扣率
func (moneyUtil) Discount2String(discount int64) string {
	if discount%10 == 0 {
		return fmt.Sprintf("%d折", discount/10)
	}
	return fmt.Sprintf("%d.%d折", discount/10, discount%10)
}

// 折扣 四舍五入
// @param cents 金额，单位分
// @param discount 折扣，单位%，范围[1,100]整数
func (moneyUtil) DiscountRound(cents int64, discount int64) int64 {
	return int64(math.Round(float64(cents*discount) / 100))
}

// 折扣 四舍五入
// @param cents 金额，单位分
// @param discount 折扣，单位%，范围[1,100]整数
func (moneyUtil) DiscountCeil(cents int64, discount int64) int64 {
	return int64(math.Ceil(float64(cents*discount) / 100))
}

// 折扣 四舍五入
// @param cents 金额，单位分
// @param discount 折扣，单位%，范围[1,100]整数
func (moneyUtil) DiscountFloor(cents int64, discount int64) int64 {
	return int64(math.Floor(float64(cents*discount) / 100))
}

// Round 四舍五入，ROUND_HALF_UP 模式实现
// 返回将 val 根据指定精度 precision（十进制小数点后数字的数目）进行四舍五入的结果。precision 也可以是负数或零。
func (moneyUtil) Round(val float64, precision int) float64 {
	p := math.Pow10(precision)
	return math.Floor(val*p+0.5) / p
}

// 金额 四舍五入
// @param cents 金额，单位分
func (moneyUtil) RoundCents(cents int64) int64 {
	val := float64(cents)
	p := math.Pow10(-1)
	return int64(math.Floor(val*p+0.5) / p)
}

/**
 * Excel工具
 */
type excelUtil struct {
}

var ExcelUtil excelUtil

// excel列名称转换为序号
// A -> 1
// B -> 2
// ...
// Z -> 26
// AA -> 27
// ...
func (excelUtil) Title2Index(s string) int {
	var ans int
	for i := 0; i < len(s); i++ {
		ans = ans*26 + int(s[i]-'A'+1)
	}
	return ans
}

// excel列序号转换为名称
// 1 -> A
// 2 -> B
// ...
// 26 -> Z
// 27 -> AA
// ...
func (excelUtil) Index2Title(n int) string {
	var ans string
	for n > 0 {
		n--
		a := n % 26
		ans = string(rune(a+'A')) + ans
		n = n / 26
	}
	return ans
}

/**
 * JSON工具
 */
type jsonUtil struct {
}

var JsonUtil jsonUtil

// 将目标对象格式化为json字符串
func (jsonUtil) Marshal(obj interface{}) string {
	b, err := json.Marshal(obj)
	if err != nil {
		logs.Info("转换json时失败: error = %v, obj = %#v", err, obj)
	}
	return string(b)
}

// 将目标对象格式化为json字符串，缩进
func (jsonUtil) MarshalIndent(obj interface{}) string {
	b, err := json.MarshalIndent(obj, "", "    ")
	if err != nil {
		logs.Info("转换json时失败: error = %v, obj = %#v", err, obj)
	}
	return string(b)
}

/**
 * 字符串工具
 */
type stringUtil struct {
}

var StringUtil stringUtil

// 手机号中间号码段使用*号代替
// @param mobile 手机号 长度11
func (stringUtil) MobileHidden(mobile string) string {
	if len(mobile) < 11 {
		return mobile
	}
	return fmt.Sprintf("%s****%s", mobile[:3], mobile[len(mobile)-4:])
}

// 截取字符串（支持中文）
// @param str 原字符串
// @param begin 起始截取下标
// @param length 截取长度
func (stringUtil) SubString(str string, begin int, length int) string {
	var rs = []rune(str)
	var lth = len(rs)
	if begin < 0 {
		begin = 0
	}
	if begin >= lth {
		begin = lth
	}
	var end = begin + length
	if end > lth {
		end = lth
	}
	return string(rs[begin:end])
}

const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// 生成随机字符串
// @param size 长度
func (stringUtil) RandChar(size int) string {
	rand.NewSource(time.Now().UnixNano()) // 产生随机种子
	var s bytes.Buffer
	for i := 0; i < size; i++ {
		s.WriteByte(chars[rand.Int63()%int64(len(chars))])
	}
	return s.String()
}

func (stringUtil) MD5(s string) string {
	ctx := md5.New()
	ctx.Write([]byte(s))
	return hex.EncodeToString(ctx.Sum(nil))
}

/**
 * 数学工具
 */
type mathUtil struct {
}

var MathUtil mathUtil

// 最小值
// @param a 比较值a
// @param b 比较值b
func (mathUtil) MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 最小值
// @param a 比较值a
// @param b 比较值b
func (mathUtil) MinInt64(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

// 最大值
// @param a 比较值a
// @param b 比较值b
func (mathUtil) MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 最大值
// @param a 比较值a
// @param b 比较值b
func (mathUtil) MaxInt64(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
