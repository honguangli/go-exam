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
				beego.NSRouter("/detail", &controllers.QuestionController{}, "Post:Detail"),
				beego.NSRouter("/create", &controllers.QuestionController{}, "Post:Create"),
				beego.NSRouter("/update", &controllers.QuestionController{}, "Post:Update"),
				beego.NSRouter("/delete", &controllers.QuestionController{}, "Post:Delete"),
			),
			// 试卷
			beego.NSNamespace("/paper",
				beego.NSRouter("/list", &controllers.PaperController{}, "Post:List"),
				beego.NSRouter("/detail", &controllers.PaperController{}, "Post:Detail"),
				beego.NSRouter("/create", &controllers.PaperController{}, "Post:Create"),
				beego.NSRouter("/update", &controllers.PaperController{}, "Post:Update"),
				beego.NSRouter("/delete", &controllers.PaperController{}, "Post:Delete"),
			),
			// 考试计划
			beego.NSNamespace("/plan",
				beego.NSRouter("/list", &controllers.PlanController{}, "Post:List"),
				beego.NSRouter("/detail", &controllers.PlanController{}, "Post:Detail"),
				beego.NSRouter("/create", &controllers.PlanController{}, "Post:Create"),
				beego.NSRouter("/update", &controllers.PlanController{}, "Post:Update"),
				beego.NSRouter("/delete", &controllers.PlanController{}, "Post:Delete"),
			),

			// 权限
			beego.NSNamespace("/permission",
				beego.NSRouter("/list", &controllers.PermissionController{}, "Post:List"),
				beego.NSRouter("/all", &controllers.PermissionController{}, "Post:All"),
				beego.NSRouter("/detail", &controllers.PermissionController{}, "Post:Detail"),
				beego.NSRouter("/create", &controllers.PermissionController{}, "Post:Create"),
				beego.NSRouter("/update", &controllers.PermissionController{}, "Post:Update"),
				beego.NSRouter("/delete", &controllers.PermissionController{}, "Post:Delete"),
			),
			// 角色
			beego.NSNamespace("/role",
				beego.NSRouter("/list", &controllers.RoleController{}, "Post:List"),
				beego.NSRouter("/all", &controllers.RoleController{}, "Post:All"),
				beego.NSRouter("/detail", &controllers.RoleController{}, "Post:Detail"),
				beego.NSRouter("/create", &controllers.RoleController{}, "Post:Create"),
				beego.NSRouter("/update", &controllers.RoleController{}, "Post:Update"),
				beego.NSRouter("/delete", &controllers.RoleController{}, "Post:Delete"),
				beego.NSRouter("/permission/list", &controllers.RoleController{}, "Post:PermissionList"),
				beego.NSRouter("/permission/auth", &controllers.RoleController{}, "Post:AuthPermission"),
			),
			// 用户
			beego.NSNamespace("/user",
				beego.NSRouter("/list", &controllers.UserController{}, "Post:List"),
				beego.NSRouter("/detail", &controllers.UserController{}, "Post:Detail"),
				beego.NSRouter("/create", &controllers.UserController{}, "Post:Create"),
				beego.NSRouter("/update", &controllers.UserController{}, "Post:Update"),
				beego.NSRouter("/delete", &controllers.UserController{}, "Post:Delete"),
				beego.NSRouter("/update_type", &controllers.UserController{}, "Post:UpdateType"),
				beego.NSRouter("/update_password", &controllers.UserController{}, "Post:UpdatePassword"),
				beego.NSRouter("/role/list", &controllers.UserController{}, "Post:RoleList"),
				beego.NSRouter("/role/auth", &controllers.UserController{}, "Post:AuthRole"),
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
