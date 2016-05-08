package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "base.html"
}

func (c *MainController) Base() {
	c.Data["json"] = map[string]interface{}{"name": "astaxie"}
	// c.Data["Email"] = "astaxie@gmail.com"
	c.ServeJSON()
	c.StopRun()
	c.TplName = "index.tpl"
}
