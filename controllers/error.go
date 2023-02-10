package controllers

import (
	"github.com/astaxie/beego"
)

type ErrorController struct {
	beego.Controller
}

type ErrResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"info"`
	Data    interface{} `json:"data"`
}

func (c *ErrorController) Error404() {
	c.Data["json"] = ErrResponse{
		Code:    404,
		Message: "Not Found",
		Data:    "NULL",
	}
	c.ServeJSON()
	c.StopRun()
}

func (c *ErrorController) Error401() {
	c.Data["json"] = ErrResponse{
		Code:    401,
		Message: "Permission denied",
		Data:    "NULL",
	}
	c.ServeJSON()
	c.StopRun()
}

func (c *ErrorController) Error403() {
	c.Data["json"] = ErrResponse{
		Code:    403,
		Message: "Forbidden",
		Data:    "NULL",
	}
	c.ServeJSON()
	c.StopRun()
}
