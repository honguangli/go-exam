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
			// 科目
			beego.NSNamespace("/subject",
				beego.NSRouter("/list", &controllers.SubjectController{}, "Post:List"),
				beego.NSRouter("/detail", &controllers.SubjectController{}, "Post:Detail"),
				beego.NSRouter("/create", &controllers.SubjectController{}, "Post:Create"),
				beego.NSRouter("/update", &controllers.SubjectController{}, "Post:Update"),
				beego.NSRouter("/delete", &controllers.SubjectController{}, "Post:Delete"),
			),
			// 知识点
			beego.NSNamespace("/knowledge",
				beego.NSRouter("/list", &controllers.KnowledgeController{}, "Post:List"),
				beego.NSRouter("/detail", &controllers.KnowledgeController{}, "Post:Detail"),
				beego.NSRouter("/create", &controllers.KnowledgeController{}, "Post:Create"),
				beego.NSRouter("/update", &controllers.KnowledgeController{}, "Post:Update"),
				beego.NSRouter("/delete", &controllers.KnowledgeController{}, "Post:Delete"),
			),

			// 试题
			beego.NSNamespace("/question",
				beego.NSRouter("/list", &controllers.QuestionController{}, "Post:List"),
			),

			// 权限
			//beego.NSRouter("/getAsyncRoutes", &controllers.UserController{}, "Post:GetAsyncRoutes"),

			// 权限
			beego.NSNamespace("/permission",
				beego.NSRouter("/list", &controllers.PermissionController{}, "Post:List"),
				beego.NSRouter("/detail", &controllers.PermissionController{}, "Post:Detail"),
				beego.NSRouter("/create", &controllers.PermissionController{}, "Post:Create"),
				beego.NSRouter("/update", &controllers.PermissionController{}, "Post:Update"),
				beego.NSRouter("/delete", &controllers.PermissionController{}, "Post:Delete"),
			),
			// 角色
			beego.NSNamespace("/role",
				beego.NSRouter("/list", &controllers.RoleController{}, "Post:List"),
				beego.NSRouter("/detail", &controllers.RoleController{}, "Post:Detail"),
				beego.NSRouter("/create", &controllers.RoleController{}, "Post:Create"),
				beego.NSRouter("/update", &controllers.RoleController{}, "Post:Update"),
				beego.NSRouter("/permission", &controllers.RoleController{}, "Post:UpdatePermission"),
				beego.NSRouter("/delete", &controllers.RoleController{}, "Post:Delete"),
			),
			// 用户
			beego.NSNamespace("/user",
				beego.NSRouter("/list", &controllers.UserController{}, "Post:List"),
				beego.NSRouter("/detail", &controllers.UserController{}, "Post:Detail"),
				beego.NSRouter("/create", &controllers.UserController{}, "Post:Create"),
				beego.NSRouter("/update", &controllers.UserController{}, "Post:Update"),
				beego.NSRouter("/role", &controllers.UserController{}, "Post:UpdateRole"),
				beego.NSRouter("/delete", &controllers.UserController{}, "Post:Delete"),
			),
			// 班级
			beego.NSNamespace("/class",
				beego.NSRouter("/list", &controllers.ClassController{}, "Post:List"),
				beego.NSRouter("/detail", &controllers.ClassController{}, "Post:Detail"),
				beego.NSRouter("/create", &controllers.ClassController{}, "Post:Create"),
				beego.NSRouter("/update", &controllers.ClassController{}, "Post:Update"),
				beego.NSRouter("/user", &controllers.ClassController{}, "Post:UpdateUser"),
				beego.NSRouter("/delete", &controllers.ClassController{}, "Post:Delete"),
			),
		),
	))
}
