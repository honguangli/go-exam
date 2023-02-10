package controllers

// 用户端
type MainController struct {
	BaseController
}

// 预执行
func (c *MainController) Prepare() {
	c.BaseController.Prepare()
}

// 首页
func (c *MainController) Index() {
	// 如果是异步请求，则说明路由未正确匹配，此时响应404即可
	if c.Ctx.Input.IsAjax() {
		c.Ctx.Output.Status = 404
		return
	}

	c.Redirect("http://192.168.0.103:8848", 302)
	return

	c.ViewPath = "static"
	c.TplName = "app/cabinet/index.html"
}
