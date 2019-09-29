package main

import (
	"github.com/astaxie/beego"
	"github.com/beego/samples/todo/controllers"
)

func main() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/task/", &controllers.TaskController{}, "get:ListTasks;post:NewTask")
	beego.Router("/test-dca/", &controllers.TaskController{}, "get:TestDca")
	beego.Router("/task/:id:int", &controllers.TaskController{}, "get:GetTask;put:UpdateTask")
	beego.Run()
}
