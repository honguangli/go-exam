package main

import (
	_ "go-exam/init"
	_ "go-exam/routers"

	"github.com/astaxie/beego"
)

func main() {
	beego.AddViewPath("static")

	beego.Run()
}
