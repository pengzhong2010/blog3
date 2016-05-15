package api

import (
	// "encoding/json"
	"github.com/astaxie/beego"
	"github.com/favframework/debug"
	"log"
	"photolimit/common"
	"photolimit/models"
	"strconv"
	// "time"
)

type WebController struct {
	beego.Controller
}

func (c *WebController) GetRecent() {
	topics := models.GetRecent()
	// godump.Dump(*topics[0])
	// b := common.Struct2Map(*topics[0])
	// godump.Dump(b)
	s := make([]map[string]interface{}, len(topics), len(topics))

	for k, v := range topics {
		b := common.Struct2Map(*v)
		log.Println(b["Created"])
		// time := v.Created.Format("2006-01-02 15:04:05")
		created_month_year := v.Created.Format("Jan 2006")
		created_day := v.Created.Format("02")
		created_month_day_year := v.Created.Format("Jan 02,2006")
		b["c_m_y"] = created_month_year
		b["c_d"] = created_day
		b["c_m_d_y"] = created_month_day_year

		s[k] = b
	}
	// godump.Dump(s)

	c.Data["json"] = map[string]interface{}{"recent": s}
	c.ServeJSON()
	c.StopRun()
}

func (c *WebController) GetBanner() {
	// c.Ctx.Output.Header("Access-Control-Allow-Origin", "*")
	banners := []*models.Banner{}
	models.GetBanners(&banners, "")
	// godump.Dump(banners)
	c.Data["json"] = map[string]interface{}{"banners": banners}
	c.ServeJSON()

	c.StopRun()
}

func (c *WebController) GetCategory() {
	categorys := []*models.Category{}
	models.GetCategorys(&categorys, "")
	// godump.Dump(banners)
	c.Data["json"] = map[string]interface{}{"categorys": categorys}
	c.ServeJSON()
	c.StopRun()
}

func (c *WebController) GetBlog() {
	cs := []*models.Topic{}
	models.GetTopics(&cs, "")
	s := make([]map[string]interface{}, len(cs), len(cs))

	for k, v := range cs {
		b := common.Struct2Map(*v)
		log.Println(b["Created"])
		// time := v.Created.Format("2006-01-02 15:04:05")
		created_month_year := v.Created.Format("Jan 2006")
		created_day := v.Created.Format("02")
		created_month_day_year := v.Created.Format("Jan 02,2006")
		b["c_m_y"] = created_month_year
		b["c_d"] = created_day
		b["c_m_d_y"] = created_month_day_year

		s[k] = b
	}
	godump.Dump(s)
	c.Data["json"] = map[string]interface{}{"blogs": s}
	c.ServeJSON()
	c.StopRun()
}

func (c *WebController) GetTopic() {
	id := c.Ctx.Input.Param(":id")
	idint, _ := strconv.Atoi(id)
	t := models.Topic{Id: idint}
	b := models.GetTopic(&t)

	if !b {
		c.Data["json"] = map[string]interface{}{"topic": ""}
		c.StopRun()
	}
	t1 := common.Struct2Map(t)
	created_month_year := t.Created.Format("Jan 2006")
	created_day := t.Created.Format("02")
	t1["c_m_y"] = created_month_year
	t1["c_d"] = created_day

	c.Data["json"] = map[string]interface{}{"topic": t1}
	c.ServeJSON()
	c.StopRun()
}

func (c *WebController) GetReply() {

}

func (c *WebController) CustomLogin() {

}

func (c *WebController) CustomLogout() {

}
