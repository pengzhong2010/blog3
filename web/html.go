package web

import (
	"github.com/astaxie/beego"
	// "github.com/favframework/debug"
	// "photolimit/models"
	// "strconv"
)

type HtmlController struct {
	beego.Controller
}

func (h *HtmlController) Home() {
	h.TplName = "web/home.html"
}

func (h *HtmlController) Category() {
	h.TplName = "web/category.html"
}

func (h *HtmlController) Blog() {
	h.TplName = "web/blog.html"
}

func (h *HtmlController) Topic() {
	id := h.Ctx.Input.Param(":id")
	// idint, _ := strconv.Atoi(id)
	// t := models.Topic{Id: idint}
	// b := models.GetTopic(&t)
	// if !b {
	// 	h.Ctx.Redirect(302, h.URLFor(".Blog"))
	// 	h.StopRun()
	// }
	// created_month_year := t.Created.Format("Jan 2006")
	// created_day := t.Created.Format("02")
	// created_month_day_year := t.Created.Format("Jan 02,2006")
	// godump.Dump(t)
	h.Data["Topic_id"] = id
	// h.Data["Topic"] = t
	// h.Data["Topic_cmy"] = created_month_year
	// h.Data["Topic_cd"] = created_day
	h.TplName = "web/topic.html"
}
