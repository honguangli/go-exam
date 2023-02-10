package routers

import (
	"go-exam/controllers"

	"github.com/astaxie/beego"
)

func init() {
	// 错误处理
	beego.ErrorController(&controllers.ErrorController{})

	beego.AddNamespace(beego.NewNamespace("/exam",
		beego.NSRouter("*", &controllers.MainController{}, "*:Index"),

		// api
		beego.NSNamespace("/api",
			// 用户
			beego.NSNamespace("/user",
				beego.NSRouter("/info", &controllers.UserController{}, "Get,Post:QueryUserInfo"),
			),
		),
	))
}
