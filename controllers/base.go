package controllers

import (
	"encoding/json"
	"time"

	"go-exam/enums"
	"go-exam/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

// app
type BaseController struct {
	beego.Controller
	controllerName string
	actionName     string
	Request        models.Request
	user           models.AppUser
}

// 预执行
func (c *BaseController) Prepare() {
	c.controllerName, c.actionName = c.GetControllerAndAction()
}

// 鉴权
func (c *BaseController) Auth() {
	c.CheckToken()

	c.Success(nil)
}

// 校验token
func (c *BaseController) CheckToken() {
	var token = c.Ctx.Request.Header.Get("Authorization")
	if len(token) == 0 {
		c.NotLogin()
		return
	}

	claims, ok := models.AuthAppToken(token)
	if !ok {
		c.NotLogin()
		return
	}
	c.user = models.AppUser{
		UserID:     claims.UserID,
		UserName:   claims.UserName,
		MerchantID: claims.MerchantID,
		StoreID:    claims.StoreID,
	}
}

// 校验权限
func (c *BaseController) CheckPermission() {
	c.CheckCustomPermission(c.controllerName, c.actionName)
}

// 校验权限
func (c *BaseController) CheckCustomPermission(controllerName string, actionName string) {
	// 校验token
	c.CheckToken()

	// 查询用户权限
	// ok, err := models.CheckResourceByApp(c.user.UserID, fmt.Sprintf("%s.%s", controllerName, actionName))
	// if err != nil {
	// 	logs.Info("app[校验权限]: 查询权限失败, err = %s", err.Error())
	// 	c.Forbidden()
	// 	return
	// }
	// if !ok {
	// 	c.Forbidden()
	// 	return
	// }
}

// Bad Request
func (c *BaseController) BadRequest(msg string) {
	c.Ctx.Output.SetStatus(int(enums.BADREQUEST))
	c.Failure(msg, enums.BADREQUEST)
}

// 未登录
func (c *BaseController) NotLogin() {
	c.Ctx.Output.SetStatus(int(enums.NOTLOGIN))
	c.Failure("请登录", enums.NOTLOGIN)
}

// 无权限
func (c *BaseController) Forbidden() {
	c.Ctx.Output.SetStatus(int(enums.FORBIDDEN))
	c.Failure("未授权", enums.FORBIDDEN)
}

// 解析请求
func (c *BaseController) ParseParam(param interface{}) error {
	var req models.Request
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &req); err != nil {
		logs.Info("请求参数错误: error = %v, req = %s", err, c.Ctx.Input.RequestBody)
		c.BadRequest("参数错误")
	}

	// 请求ID
	c.Request = req

	c.ValidRequest()

	// 签名校验
	if req.Sign != models.Sign(req.ReqID, req.Data, req.Time) {
		c.BadRequest("签名错误")
	}

	return json.Unmarshal([]byte(req.Data), param)
}

// 校验参数
func (c *BaseController) ValidRequest() {
	// 请求ID
	if len(c.Request.ReqID) == 0 {
		c.Error("无效ID")
	}

	// 时间校验
	if c.Request.Time <= 0 {
		c.Error("请求时间错误")
	}

	// 有效时间范围：当前时间 +- duration分钟
	var now = time.Now()
	var duration = time.Minute * 5
	if c.Request.Time < now.Add(duration*-1).Unix() || c.Request.Time > now.Add(duration).Unix() {
		c.Error("请求超时")
	}

	// 签名
	if len(c.Request.Sign) == 0 {
		c.Error("签名错误")
	}
}

// 响应 异常
func (c *BaseController) Error(msg string) {
	var resp = c.NewResponse(enums.ERROR, msg, nil)
	c.Data["json"] = resp
	c.ServeJSON()
	c.StopRun()
}

// 响应 失败
func (c *BaseController) Failure(msg string, code ...enums.ResponseCode) {
	var resp *models.Response
	if len(code) == 0 {
		resp = c.NewResponse(enums.FAILURE, msg, nil)
	} else {
		resp = c.NewResponse(code[0], msg, nil)
	}
	c.Data["json"] = resp
	c.ServeJSON()
	c.StopRun()
}

// 响应 成功
func (c *BaseController) Success(data interface{}, msg ...string) {
	var resp *models.Response
	if len(msg) == 0 {
		resp = c.NewResponse(enums.SUCCESS, "SUCCESS", data)
	} else {
		resp = c.NewResponse(enums.SUCCESS, msg[0], data)
	}
	c.Data["json"] = resp
	c.ServeJSON()
	c.StopRun()
}

// 响应 自定义
func (c *BaseController) DataResult(resp *models.Response) {
	c.Data["json"] = resp
	c.ServeJSON()
	c.StopRun()
}

// 响应
func (c *BaseController) NewResponse(code enums.ResponseCode, msg string, data interface{}) *models.Response {
	var resp = models.Response{
		ReqID: c.Request.ReqID,
		Time:  time.Now().Unix(),
	}

	if data == nil || data == "" {
		resp.Data = "{}"
	} else {
		jsonData, err := json.Marshal(data)
		if err != nil {
			logs.Info("返回数据格式异常：error = %s，data = %v，code = %d，msg = %s", err.Error(), data, code, msg)
			code = enums.FAILURE
			msg = "SERVICE ERROR"
			jsonData = []byte("{}")
		}
		resp.Data = string(jsonData)
	}

	resp.Code = code
	resp.Msg = msg
	resp.Sign = models.Sign(resp.ReqID, resp.Data, resp.Time)
	return &resp
}
