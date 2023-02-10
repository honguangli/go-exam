package models

import (
	"fmt"
	"go-exam/enums"
	"go-exam/utils"
	"time"

	"github.com/astaxie/beego/logs"
	"github.com/dgrijalva/jwt-go"
)

// web app name
const (
	APP_TOOLS = "app_tools"
)

// web app 通讯秘钥
const (
	AppKey    = "cz74e826ea5485aaa233ceb972a3326bde"
	AppSecret = "2bb011d61507f7eeabd6ca7cc6d551eb"
)

// 认证token key
const AppTokenKey = "welcome to cabinet app"

// 请求体
type Request struct {
	ReqID string `json:"req_id"` // 请求ID
	Data  string `json:"data"`   // 请求数据
	Time  int64  `json:"time"`   // 请求时间
	Sign  string `json:"sign"`   // 签名
}

// 响应体
type Response struct {
	ReqID string             `json:"req_id"` // 请求ID 与请求参数中的req_id保持一致
	Code  enums.ResponseCode `json:"code"`   // 响应代码
	Msg   string             `json:"msg"`    // 响应描述
	Data  string             `json:"data"`   // 响应数据
	Time  int64              `json:"time"`   // 响应时间
	Sign  string             `json:"sign"`   // 签名
}

// 签名
func Sign(reqID string, data string, timeStamp int64) string {
	return utils.StringUtil.MD5(fmt.Sprintf("req%s-key%s+secret%s-data%s+time%d-key%s", reqID, AppKey, AppSecret, data, timeStamp, AppKey))
}

// app token 参数
type AppTokenClaims struct {
	UserID     int    `json:"user_id"`
	UserName   string `json:"user_name"`
	MerchantID string `json:"merchant_id"`
	StoreID    string `json:"store_id"`
	jwt.StandardClaims
}

// 生成 app token
func GenAppToken(userID int, userName string, merchantID string, storeID string, hour int) (tokenString string, expireTime int64, err error) {
	var now = time.Now()
	expireTime = now.Add(time.Hour * time.Duration(hour)).Unix()
	claims := AppTokenClaims{userID, userName, merchantID, storeID,
		jwt.StandardClaims{
			ExpiresAt: expireTime,
			IssuedAt:  now.Unix(),
			Issuer:    "cms",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString([]byte(AppTokenKey))
	if err != nil {
		logs.Info("jwt: sign token string error: err = %s", err.Error())
		return
	}
	return
}

// 校验 app token
func AuthAppToken(tokenString string) (claims *AppTokenClaims, ok bool) {
	token, err := jwt.ParseWithClaims(tokenString, &AppTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(AppTokenKey), nil
	})
	if err != nil {
		return nil, false
	}
	claims, ok = token.Claims.(*AppTokenClaims)
	if !ok || !token.Valid {
		return nil, false
	}
	if claims.UserID <= 0 {
		return nil, false
	}
	return claims, true
}

// web app 登录账号
type AppUser struct {
	UserID     int             `json:"user_id"`
	UserName   string          `json:"user_name"`
	Avatar     string          `json:"avatar"`
	MerchantID string          `json:"merchant_id"`
	StoreID    string          `json:"store_id"`
	Permission []AppPermission `json:"permission"`
	Token      string          `json:"token"`
	ExpireTime int64           `json:"expire_time"`
}

// web app 权限
type AppPermission struct {
	Name string            `json:"name"`
	Args map[string]string `json:"args"`
}

// 获取app权限前缀
func GetAppPermissionPrefix(appName string) (prefix string) {
	switch appName {
	case APP_TOOLS:
		prefix = "AppTools"
	}
	return
}

// 空数组对象
var EmptyArray = make([]struct{}, 0)
