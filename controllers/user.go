package controllers

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
