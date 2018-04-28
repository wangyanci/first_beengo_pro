package main

import "code.wangyanci.com/benngo_project/swr/controller"
import "github.com/astaxie/beego"

func main() {
	beego.Router("/", &controller.FistController{})
	beego.Router("/abc", &controller.SecondController{})
	beego.Run()
}