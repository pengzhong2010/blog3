package admin

import (
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"github.com/favframework/debug"
	// "log"
	// "path/filepath"
	"photolimit/models"
	// "strconv"
	// "reflect"
	// "strings"
)

type TopicController struct {
	beego.Controller
}

func (c *TopicController) TopicList() {
	q := c.Ctx.Input.Query("q")
	topics := []models.Topic{}
	topic := models.Topic{}
	topic.TopicList(&topics, q)
	godump.Dump(topics)
	c.Data["topics"] = topics
	c.Data["Title"] = "TopicList"
	c.Layout = "admin/layout.html"
	c.TplName = "admin/topiclist.html"

}

type TopicForm struct {
	Name    string `form:"name" valid:"Required;MaxSize(180)"`
	Content string `form:"content"`
}

func (c *TopicController) TopicAdd() {
	c.Layout = "admin/layout.html"
	c.TplName = "admin/topicadd.html"
}

func (c *TopicController) TopicAddDo() {
	tf := TopicForm{}
	c.ParseForm(&tf)
	valid := validation.Validation{}
	b, _ := valid.Valid(&tf)
	if !b {
		c.Ctx.Output.Body([]byte("valid error"))
		c.StopRun()
	}

	user := c.Ctx.Input.Session("user").(models.User)
	topicdb := models.Topic{Name: tf.Name, Content: tf.Content, User: &user}
	res_b, _ := topicdb.TopicAdd()
	if !res_b {
		c.Ctx.Output.Body([]byte("add faild"))
		c.StopRun()
	}
	c.Ctx.Redirect(302, c.URLFor(".TopicList"))
}

func (c *TopicController) TopicEdit() {
	id := c.Ctx.Input.Param(":id")
	idint, _ := strconv.Atoi(id)
	topic := models.Topic{Id: idint}
	topic.TopicRead()
	godump.Dump(topic)
	// // godump.Dump(categoryimg)
	// if category.Id == 0 {
	// 	c.Ctx.Output.Body([]byte("not found"))
	// 	c.StopRun()
	// }
	// beego.ReadFromRequest(&c.Controller)
	// c.Data["category"] = category
	// c.Layout = "admin/layout.html"
	// c.TplName = "admin/topicedit.html"
}

// func (t *TopicController) TopicEditDo() {
// 	cf := CategoryForm{}
// 	c.ParseForm(&cf)
// 	godump.Dump(cf)

// 	valid := validation.Validation{}
// 	b, _ := valid.Valid(&cf)
// 	if !b {
// 		c.Ctx.Output.Body([]byte("valid error"))
// 		c.StopRun()
// 	}
// 	id := c.Ctx.Input.Param(":id")
// 	idint, _ := strconv.Atoi(id)
// 	cm := models.Category{Id: idint, Name: cf.Name}
// 	res_b := cm.CategoryEdit()
// 	if !res_b {
// 		c.Ctx.Output.Body([]byte("error"))
// 		c.StopRun()
// 	}
// 	c.Ctx.Redirect(302, c.URLFor(".CategoryList"))

// }

// func (t *TopicController) TopicEditImg() {
// 	id := c.Ctx.Input.Param(":id")
// 	idint, _ := strconv.Atoi(id)
// 	category := models.Category{Id: idint}
// 	category.CategoryRead()
// 	if category.Id == 0 {
// 		flash := beego.NewFlash()
// 		flash.Notice("category not match")
// 		flash.Store(&c.Controller)
// 		c.Ctx.Redirect(302, c.URLFor(".CategoryList"))
// 		c.StopRun()
// 	}

// 	f, _, err := c.GetFile("img")
// 	defer f.Close()
// 	if err != nil {
// 		flash := beego.NewFlash()
// 		flash.Notice("empty img")
// 		flash.Store(&c.Controller)
// 		c.Ctx.Redirect(302, c.URLFor(".CategoryEdit", ":id", id))
// 		c.StopRun()
// 	}

// 	basedir := beego.AppConfig.String("basedir")
// 	sep := string(filepath.Separator)
// 	imgdir := basedir + sep + "static" + sep + "img" + sep

// 	image := models.Image{}

// 	b1 := image.ImageCreate()
// 	if !b1 {
// 		flash := beego.NewFlash()
// 		flash.Notice("img create error")
// 		flash.Store(&c.Controller)
// 		c.Ctx.Redirect(302, c.URLFor(".CategoryEdit", ":id", id))
// 		c.StopRun()
// 	}

// 	imgdir += strconv.Itoa(image.Id) + ".jpg"
// 	c.SaveToFile("img", imgdir)
// 	image.Url = "/static/img/" + strconv.Itoa(image.Id) + ".jpg"
// 	b2 := image.ImageUpdate()

// 	if !b2 {
// 		flash := beego.NewFlash()
// 		flash.Notice("img update error")
// 		flash.Store(&c.Controller)
// 		c.Ctx.Redirect(302, c.URLFor(".CategoryEdit", ":id", id))
// 		c.StopRun()
// 	}
// 	// c.StopRun()
// 	category.Image = &image
// 	b3 := category.CategoryEdit()
// 	// godump.Dump(b3)
// 	// c.StopRun()
// 	if !b3 {
// 		flash := beego.NewFlash()
// 		flash.Notice("edit faild")
// 		flash.Store(&c.Controller)
// 		c.Ctx.Redirect(302, c.URLFor(".CategoryEdit", ":id", id))
// 		c.StopRun()
// 	}

// 	flash := beego.NewFlash()
// 	flash.Notice("success")
// 	flash.Store(&c.Controller)
// 	c.Ctx.Redirect(302, c.URLFor(".CategoryEdit", ":id", id))
// 	c.StopRun()
// }
