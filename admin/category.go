package admin

import (
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"github.com/favframework/debug"
	// "log"
	"path/filepath"
	"photolimit/models"
	"strconv"
	// "reflect"
	// "strings"
)

type CategoryController struct {
	beego.Controller
}

func (c *CategoryController) CategoryList() {
	q := c.Ctx.Input.Query("q")
	categorys := []models.Category{}
	category := models.Category{}
	category.CategoryList(&categorys, q)
	// godump.Dump(categorys)
	c.Data["category"] = categorys
	c.Data["Title"] = "CategoryList"
	c.Layout = "admin/layout.html"
	c.TplName = "admin/categorylist.html"

}

type CategoryForm struct {
	Name string `form:"name" valid:"Required;MaxSize(60)"`
}

func (c *CategoryController) CategoryAdd() {
	c.Layout = "admin/layout.html"
	c.TplName = "admin/categoryadd.html"
}

func (c *CategoryController) CategoryAddDo() {
	cf := CategoryForm{}
	c.ParseForm(&cf)
	valid := validation.Validation{}
	b, _ := valid.Valid(&cf)
	if !b {
		c.Ctx.Output.Body([]byte("valid error"))
		c.StopRun()
	}
	categorydb := models.Category{Name: cf.Name}
	res_b, _ := categorydb.CategoryAdd()
	if !res_b {
		c.Ctx.Output.Body([]byte("name exist"))
		c.StopRun()
	}
	c.Ctx.Redirect(302, c.URLFor(".CategoryList"))
}

func (c *CategoryController) CategoryEdit() {
	id := c.Ctx.Input.Param(":id")
	idint, _ := strconv.Atoi(id)
	category := models.Category{Id: idint}
	category.CategoryRead()
	godump.Dump(category)
	// godump.Dump(categoryimg)
	if category.Id == 0 {
		c.Ctx.Output.Body([]byte("not found"))
		c.StopRun()
	}
	beego.ReadFromRequest(&c.Controller)
	c.Data["category"] = category
	c.Layout = "admin/layout.html"
	c.TplName = "admin/categoryadd.html"
}

func (c *CategoryController) CategoryEditDo() {
	cf := CategoryForm{}
	c.ParseForm(&cf)
	godump.Dump(cf)

	valid := validation.Validation{}
	b, _ := valid.Valid(&cf)
	if !b {
		c.Ctx.Output.Body([]byte("valid error"))
		c.StopRun()
	}
	id := c.Ctx.Input.Param(":id")
	idint, _ := strconv.Atoi(id)
	cm := models.Category{Id: idint, Name: cf.Name}
	res_b := cm.CategoryEdit()
	if !res_b {
		c.Ctx.Output.Body([]byte("error"))
		c.StopRun()
	}
	c.Ctx.Redirect(302, c.URLFor(".CategoryList"))

}

func (c *CategoryController) CategoryEditImg() {
	id := c.Ctx.Input.Param(":id")
	idint, _ := strconv.Atoi(id)
	category := models.Category{Id: idint}
	category.CategoryRead()
	if category.Id == 0 {
		flash := beego.NewFlash()
		flash.Notice("category not match")
		flash.Store(&c.Controller)
		c.Ctx.Redirect(302, c.URLFor(".CategoryList"))
		c.StopRun()
	}

	f, _, err := c.GetFile("img")
	defer f.Close()
	if err != nil {
		flash := beego.NewFlash()
		flash.Notice("empty img")
		flash.Store(&c.Controller)
		c.Ctx.Redirect(302, c.URLFor(".CategoryEdit", ":id", id))
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
		c.Ctx.Redirect(302, c.URLFor(".CategoryEdit", ":id", id))
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
		c.Ctx.Redirect(302, c.URLFor(".CategoryEdit", ":id", id))
		c.StopRun()
	}
	// c.StopRun()
	category.Image = &image
	b3 := category.CategoryEdit()
	// godump.Dump(b3)
	// c.StopRun()
	if !b3 {
		flash := beego.NewFlash()
		flash.Notice("edit faild")
		flash.Store(&c.Controller)
		c.Ctx.Redirect(302, c.URLFor(".CategoryEdit", ":id", id))
		c.StopRun()
	}

	flash := beego.NewFlash()
	flash.Notice("success")
	flash.Store(&c.Controller)
	c.Ctx.Redirect(302, c.URLFor(".CategoryEdit", ":id", id))
	c.StopRun()
}
