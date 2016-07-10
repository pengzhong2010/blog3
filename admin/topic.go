package admin

import (
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"github.com/favframework/debug"
	// "log"
	"path/filepath"
	// "photolimit/common"
	"photolimit/models"
	"strconv"
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
	res_b, res_id := topicdb.TopicAdd()
	if !res_b {
		c.Ctx.Output.Body([]byte("add faild"))
		c.StopRun()
	}
	c.Ctx.Redirect(302, c.URLFor(".TopicEdit", ":id", res_id))
}

func (c *TopicController) TopicEdit() {
	id := c.Ctx.Input.Param(":id")
	idint, _ := strconv.Atoi(id)
	topic := models.Topic{Id: idint}

	b := topic.TopicRead()
	if !b {
		c.Ctx.Output.Body([]byte("not found"))
		c.StopRun()
	}

	_, categorys := topic.GetTopicCategorys()
	_, images := topic.GetTopicImages()
	godump.Dump(images)
	// godump.Dump(b1)
	// godump.Dump(b2)
	// godump.Dump(category)
	// c.StopRun()
	// categorys_all := []models.Category{}
	// category := models.Category{}
	// category.CategoryList(&categorys_all, "")

	beego.ReadFromRequest(&c.Controller)
	c.Data["topic"] = topic
	c.Data["categorys"] = categorys
	// c.Data["all_categorys"] = categorys_all
	c.Data["images"] = images
	c.Layout = "admin/layout.html"
	c.TplName = "admin/topicedit.html"
}

func (c *TopicController) TopicEditDo() {
	tf := TopicForm{}
	c.ParseForm(&tf)
	valid := validation.Validation{}
	b, _ := valid.Valid(&tf)
	if !b {
		c.Ctx.Output.Body([]byte("valid error"))
		c.StopRun()
	}

	user := c.Ctx.Input.Session("user").(models.User)
	id := c.Ctx.Input.Param(":id")
	idint, _ := strconv.Atoi(id)
	topicdb := models.Topic{Id: idint, Name: tf.Name, Content: tf.Content, User: &user}
	res_b := topicdb.TopicEdit()
	if !res_b {
		c.Ctx.Output.Body([]byte("error"))
		c.StopRun()
	}
	categorys := c.GetStrings("categorys[]")
	categorys_s := make([]int, len(categorys), len(categorys))
	// godump.Dump(categorys_s)
	for k, v := range categorys {
		tmp, _ := strconv.Atoi(v)
		categorys_s[k] = tmp
	}

	// godump.Dump(categorys_s)
	b3 := topicdb.UpdateTopicCategorys(categorys_s)
	godump.Dump(b3)
	// c.StopRun()
	c.Ctx.Redirect(302, c.URLFor(".TopicEdit", ":id", id))

}

func (c *TopicController) TopicAddImg() {
	id := c.Ctx.Input.Param(":id")
	idint, _ := strconv.Atoi(id)
	topic := models.Topic{Id: idint}
	topic.TopicRead()
	if topic.Id == 0 {
		flash := beego.NewFlash()
		flash.Notice("category not match")
		flash.Store(&c.Controller)
		c.Ctx.Redirect(302, c.URLFor(".TopicList"))
		c.StopRun()
	}

	f, _, err := c.GetFile("img")
	defer f.Close()
	if err != nil {
		flash := beego.NewFlash()
		flash.Notice("empty img")
		flash.Store(&c.Controller)
		c.Ctx.Redirect(302, c.URLFor(".TopicEdit", ":id", id))
		c.StopRun()
	}

	basedir := beego.AppConfig.String("basedir")
	sep := string(filepath.Separator)
	imgdir := basedir + sep + "static" + sep + "img" + sep

	image := models.Image{}

	b1 := image.ImageCreate()
	if !b1 {
		flash := beego.NewFlash()
		flash.Notice("img create error")
		flash.Store(&c.Controller)
		c.Ctx.Redirect(302, c.URLFor(".TopicEdit", ":id", id))
		c.StopRun()
	}

	imgdir += strconv.Itoa(image.Id) + ".jpg"
	c.SaveToFile("img", imgdir)
	image.Url = "/static/img/" + strconv.Itoa(image.Id) + ".jpg"
	b2 := image.ImageUpdate()

	if !b2 {
		flash := beego.NewFlash()
		flash.Notice("img update error")
		flash.Store(&c.Controller)
		c.Ctx.Redirect(302, c.URLFor(".TopicEdit", ":id", id))
		c.StopRun()
	}
	// c.StopRun()

	b3 := topic.AddImage(&image)
	// godump.Dump(b3)
	// c.StopRun()
	if !b3 {
		flash := beego.NewFlash()
		flash.Notice("edit faild")
		flash.Store(&c.Controller)
		c.Ctx.Redirect(302, c.URLFor(".TopicEdit", ":id", id))
		c.StopRun()
	}

	flash := beego.NewFlash()
	flash.Notice("success")
	flash.Store(&c.Controller)
	c.Ctx.Redirect(302, c.URLFor(".TopicEdit", ":id", id))
	c.StopRun()
}

func (c *TopicController) TopicDelImg() {
	id := c.Ctx.Input.Param(":id")
	idint, _ := strconv.Atoi(id)
	img_id := c.Ctx.Input.Param(":img_id")
	img_idint, _ := strconv.Atoi(img_id)
	topic := models.Topic{Id: idint}
	topic.TopicRead()
	if topic.Id == 0 {
		flash := beego.NewFlash()
		flash.Notice("topic not match")
		flash.Store(&c.Controller)
		c.Ctx.Redirect(302, c.URLFor(".TopicList"))
		c.StopRun()
	}
	image := models.Image{Id: img_idint}
	b1 := image.ImageRead()
	if !b1 {
		flash := beego.NewFlash()
		flash.Notice("img not exist")
		flash.Store(&c.Controller)
		c.Ctx.Redirect(302, c.URLFor(".TopicEdit", ":id", id))
		c.StopRun()
	}
	b2 := topic.DelImage(&image)
	if !b2 {
		flash := beego.NewFlash()
		flash.Notice("img del error")
		flash.Store(&c.Controller)
		c.Ctx.Redirect(302, c.URLFor(".TopicEdit", ":id", id))
		c.StopRun()
	}
	flash := beego.NewFlash()
	flash.Notice("success")
	flash.Store(&c.Controller)
	c.Ctx.Redirect(302, c.URLFor(".TopicEdit", ":id", id))
	c.StopRun()
}

func (c *TopicController) TopicDel() {
	id := c.Ctx.Input.Param(":id")
	idint, _ := strconv.Atoi(id)
	topic := models.Topic{Id: idint}
	topic.TopicRead()
	if topic.Id == 0 {
		flash := beego.NewFlash()
		flash.Notice("topic not match")
		flash.Store(&c.Controller)
		c.Ctx.Redirect(302, c.URLFor(".TopicList"))
		c.StopRun()
	}
	b := topic.TopicDel()
	if !b {
		flash := beego.NewFlash()
		flash.Notice("img del error")
		flash.Store(&c.Controller)
		c.Ctx.Redirect(302, c.URLFor(".TopicList"))
		c.StopRun()
	}
	flash := beego.NewFlash()
	flash.Notice("success")
	flash.Store(&c.Controller)
	c.Ctx.Redirect(302, c.URLFor(".TopicList"))
	c.StopRun()
}
