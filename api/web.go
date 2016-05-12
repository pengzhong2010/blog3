package api

import (
	"github.com/astaxie/beego"
	"github.com/favframework/debug"
	"photolimit/models"
)

type WebController struct {
	beego.Controller
}

func (c *WebController) GetRecent() {
	topics := models.GetRecent()
	godump.Dump(topics)
	c.Data["json"] = map[string]interface{}{"recent": topics}
	c.ServeJSON()
	c.StopRun()
}

func (c *WebController) GetBanner() {

}

func (c *WebController) GetCategory() {

}

func (c *WebController) GetBlog() {

}

func (c *WebController) GetTopic() {

}

func (c *WebController) GetReply() {

}

func (c *WebController) CustomLogin() {

}

func (c *WebController) CustomLogout() {

}
