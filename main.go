package main

import (
	"github.com/IBM-Cloud/cdn-with-cda-todolist/controllers"
	"github.com/beego/beego"
)

func main() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/task/", &controllers.TaskController{}, "get:ListTasks;post:NewTask")
	beego.Router("/test-dca/", &controllers.TaskController{}, "get:TestDca")
	beego.Router("/task/:id:int", &controllers.TaskController{}, "get:GetTask;put:UpdateTask")
	beego.Run()
}
