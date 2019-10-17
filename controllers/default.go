package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.TplName = "index.html"
	this.Ctx.Output.Header("Cache-Control", "no-cache, no-store, must-revalidate")
	this.Render()
}

func (this *MainController) Get2() {
	this.TplName = "detection-test-object.html"
	this.Render()
}

