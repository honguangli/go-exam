package controllers

import (
	"go-exam/models"
)

// 用户
type UserController struct {
	BaseController
}

// 预执行
func (c *UserController) Prepare() {
	c.BaseController.Prepare()
}

// 查询用户信息
func (c *UserController) QueryUserInfo() {
	type Param struct {
		Path string `json:"path"`
	}
	var param Param
	var err error
	if err = c.ParseParam(&param); err != nil {
		c.Failure("参数错误")
	}
	if len(param.Path) == 0 {
		c.Failure("参数错误")
	}

	c.Success(nil)
}

func (c *UserController) GetAsyncRoutes() {
	//c.Data["json"] = "{\"success\":true,\"data\":[{\"path\":\"/permission\",\"meta\":{\"title\":\"menus.permission\",\"icon\":\"lollipop\",\"rank\":10},\"children\":[{\"path\":\"/permission/page/index\",\"name\":\"PermissionPage\",\"meta\":{\"title\":\"menus.permissionPage\",\"roles\":[\"admin\",\"common\"]}},{\"path\":\"/permission/button/index\",\"name\":\"PermissionButton\",\"meta\":{\"title\":\"menus.permissionButton\",\"roles\":[\"admin\",\"common\"],\"auths\":[\"btn_add\",\"btn_edit\",\"btn_delete\"]}}]}]}"
	//c.ServeJSON()
	//c.StopRun()
	//var str = "[{\"path\":\"/permission\",\"meta\":{\"title\":\"menus.permission\",\"icon\":\"lollipop\",\"rank\":10},\"children\":[{\"path\":\"/permission/page/index\",\"name\":\"PermissionPage\",\"meta\":{\"title\":\"menus.permissionPage\",\"roles\":[\"admin\",\"common\"]}},{\"path\":\"/permission/button/index\",\"name\":\"PermissionButton\",\"meta\":{\"title\":\"menus.permissionButton\",\"roles\":[\"admin\",\"common\"],\"auths\":[\"btn_add\",\"btn_edit\",\"btn_delete\"]}}]}]"
	//
	//var ps = make([]*models.Permission, 0)
	//json.Unmarshal([]byte(str), &ps)

	//for _, v := range ps {
	//	models.SetEmptyArray(v)
	//}
	ps, _, _ := models.QueryOldPermissionList()

	c.Success(ps)
}
