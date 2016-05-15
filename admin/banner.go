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

type BannerController struct {
	beego.Controller
}

func (c *BannerController) BannerList() {
	q := c.Ctx.Input.Query("q")
	banners := []*models.Banner{}

	models.GetBanners(&banners, q)
	godump.Dump(banners)
	c.Data["banners"] = banners
	c.Data["Title"] = "BannerList"
	c.Layout = "admin/layout.html"
	c.TplName = "admin/bannerlist.html"

}

type BannerForm struct {
	Name string `form:"name" valid:"Required;MaxSize(60)"`
	Url  string `form:"url" valid:"MaxSize(300)"`
}

func (c *BannerController) BannerAdd() {
	c.Layout = "admin/layout.html"
	c.TplName = "admin/banneradd.html"
}

func (c *BannerController) BannerAddDo() {
	bf := BannerForm{}
	c.ParseForm(&bf)
	valid := validation.Validation{}
	b, _ := valid.Valid(&bf)
	if !b {
		c.Ctx.Output.Body([]byte("valid error"))
		c.StopRun()
	}
	bannerdb := models.Banner{Name: bf.Name, Url: bf.Url}
	res_b, _ := bannerdb.BannerAdd()
	if !res_b {
		c.Ctx.Output.Body([]byte("name exist"))
		c.StopRun()
	}
	c.Ctx.Redirect(302, c.URLFor(".BannerList"))
}

func (c *BannerController) BannerEdit() {
	id := c.Ctx.Input.Param(":id")
	idint, _ := strconv.Atoi(id)
	banner := models.Banner{Id: idint}
	banner.BannerRead()

	godump.Dump(banner)
	if banner.Id == 0 {
		c.Ctx.Output.Body([]byte("not found"))
		c.StopRun()
	}
	beego.ReadFromRequest(&c.Controller)
	c.Data["banner"] = banner
	c.Layout = "admin/layout.html"
	c.TplName = "admin/banneredit.html"
}

func (c *BannerController) BannerEditDo() {
	bf := BannerForm{}
	c.ParseForm(&bf)
	godump.Dump(bf)

	valid := validation.Validation{}
	b, _ := valid.Valid(&bf)
	if !b {
		c.Ctx.Output.Body([]byte("valid error"))
		c.StopRun()
	}
	id := c.Ctx.Input.Param(":id")
	idint, _ := strconv.Atoi(id)
	bm := models.Banner{Id: idint}
	b1 := bm.BannerRead()
	if !b1 {
		c.Ctx.Output.Body([]byte("banner not exist"))
		c.StopRun()
	}
	bm.Name = bf.Name
	bm.Url = bf.Url
	res_b := bm.BannerEdit()
	if !res_b {
		c.Ctx.Output.Body([]byte("error"))
		c.StopRun()
	}
	c.Ctx.Redirect(302, c.URLFor(".BannerEdit", ":id", id))

}

func (c *BannerController) BannerEditImg() {
	id := c.Ctx.Input.Param(":id")
	idint, _ := strconv.Atoi(id)
	banner := models.Banner{Id: idint}
	banner.BannerRead()
	if banner.Id == 0 {
		flash := beego.NewFlash()
		flash.Notice("category not match")
		flash.Store(&c.Controller)
		c.Ctx.Redirect(302, c.URLFor(".BannerList"))
		c.StopRun()
	}

	f, _, err := c.GetFile("img")
	defer f.Close()
	if err != nil {
		flash := beego.NewFlash()
		flash.Notice("empty img")
		flash.Store(&c.Controller)
		c.Ctx.Redirect(302, c.URLFor(".BannerEdit", ":id", id))
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
		c.Ctx.Redirect(302, c.URLFor(".BannerEdit", ":id", id))
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
		c.Ctx.Redirect(302, c.URLFor(".BannerEdit", ":id", id))
		c.StopRun()
	}
	// c.StopRun()

	banner.Image = &image
	godump.Dump(banner)
	b3 := banner.BannerEditImg()
	godump.Dump(banner)
	// c.StopRun()
	if !b3 {
		flash := beego.NewFlash()
		flash.Notice("edit faild")
		flash.Store(&c.Controller)
		c.Ctx.Redirect(302, c.URLFor(".BannerEdit", ":id", id))
		c.StopRun()
	}

	flash := beego.NewFlash()
	flash.Notice("success")
	flash.Store(&c.Controller)
	c.Ctx.Redirect(302, c.URLFor(".BannerEdit", ":id", id))
	c.StopRun()
}

func (c *BannerController) BannerDel() {
	id := c.Ctx.Input.Param(":id")
	idint, _ := strconv.Atoi(id)
	banner := models.Banner{Id: idint}
	banner.BannerDel()
	c.Ctx.Redirect(302, c.URLFor(".BannerList"))
}
