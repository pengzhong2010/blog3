package api

import (
	"github.com/astaxie/beego"
	"photolimit/models"
)

type WebController struct {
	beego.Controller
}

func (c *WebController) GetRecent() {
	models.GetRecent()
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
