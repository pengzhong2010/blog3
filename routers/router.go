package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"photolimit/admin"
	"photolimit/api"
	"photolimit/controllers"
	"photolimit/web"
)

var LoginCheck = func(ctx *context.Context) {
	user := ctx.Input.Session("user")
	if user == nil {
		url := ctx.Input.URL()
		if url != "/admin/login" {
			if url != "/admin/logout" {
				ctx.Redirect(302, "/admin/login"+"?redirect="+url)
			} else {
				ctx.Redirect(302, "/admin/login")
			}

		}
	}
}
var AccessControllAllow = func(ctx *context.Context) {
	ctx.Output.Header("Access-Control-Allow-Origin", "*")
}

func init() {
	beego.Router("/", &web.HtmlController{}, "get:Home")
	beego.Router("/home", &web.HtmlController{}, "get:Home")
	beego.Router("/category", &web.HtmlController{}, "get:Category")
	beego.Router("/blog", &web.HtmlController{}, "get:Blog")
	beego.Router("/topic/:id([1-9][0-9]*", &web.HtmlController{}, "get:Topic")

	ns := beego.NewNamespace("/admin",
		beego.NSRouter("/", &admin.TopicController{}, "get:TopicList"),
		beego.NSRouter("/login", &admin.UserController{}, "get:UserLogin;post:UserLoginDo"),
		beego.NSRouter("/logout", &admin.UserController{}, "get:UserLogout"),
		beego.NSNamespace("/user",
			beego.NSRouter("/list", &admin.UserController{}, "get:UserList"),
			beego.NSRouter("/add", &admin.UserController{}, "get:UserAdd;post:UserAddDo"),
			beego.NSRouter("/edit/:id([0-9]+)", &admin.UserController{}, "get:UserEdit;post:UserEditDo"),
			beego.NSRouter("/del/:id([0-9]+)", &admin.UserController{}, "get:UserDel"),
		),
		beego.NSNamespace("/banner",
			beego.NSRouter("/list", &admin.BannerController{}, "get:BannerList"),
			beego.NSRouter("/add", &admin.BannerController{}, "get:BannerAdd;post:BannerAddDo"),
			beego.NSRouter("/edit/:id([0-9]+)", &admin.BannerController{}, "get:BannerEdit;post:BannerEditDo"),
			beego.NSRouter("/edit/:id([0-9]+)/img", &admin.BannerController{}, "post:BannerEditImg"),
			beego.NSRouter("/del/:id([0-9]+)", &admin.BannerController{}, "get:BannerDel"),
		),
		beego.NSNamespace("/category",
			beego.NSRouter("/list", &admin.CategoryController{}, "get:CategoryList"),
			beego.NSRouter("/add", &admin.CategoryController{}, "get:CategoryAdd;post:CategoryAddDo"),
			beego.NSRouter("/edit/:id([0-9]+)", &admin.CategoryController{}, "get:CategoryEdit;post:CategoryEditDo"),
			beego.NSRouter("/edit/:id([0-9]+)/img", &admin.CategoryController{}, "post:CategoryEditImg"),
			// beego.NSRouter("/del/:id([0-9]+)", &admin.CategoryController{}, "post:CategoryDel"),
		),
		beego.NSNamespace("/topic",
			beego.NSRouter("/list", &admin.TopicController{}, "get:TopicList"),
			beego.NSRouter("/add", &admin.TopicController{}, "get:TopicAdd;post:TopicAddDo"),
			beego.NSRouter("/edit/:id", &admin.TopicController{}, "get:TopicEdit;post:TopicEditDo"),
			beego.NSRouter("/edit/:id([0-9]+)/img", &admin.TopicController{}, "post:TopicAddImg"),
			beego.NSRouter("/edit/:id([0-9]+)/img_del/:img_id([0-9]+)", &admin.TopicController{}, "get:TopicDelImg"),
			beego.NSRouter("/del/:id", &admin.TopicController{}, "get:TopicDel"),
		),
		beego.NSRouter("/base", &controllers.MainController{}, "get:Get"),
	)
	beego.AddNamespace(ns)

	beego.InsertFilter("/admin/*", beego.BeforeExec, LoginCheck)

	ns1 := beego.NewNamespace("/api",
		beego.NSRouter("/banner", &api.WebController{}, "get:GetBanner"),
		beego.NSRouter("/category", &api.WebController{}, "get:GetCategory"),
		beego.NSRouter("/blog", &api.WebController{}, "get:GetBlog"),
		beego.NSRouter("/recent", &api.WebController{}, "get:GetRecent"),
		beego.NSRouter("/topic/:id([1-9][0-9]*)", &api.WebController{}, "get:GetTopic"),
		// beego.NSRouter("/login", &admin.UserController{}, "get:UserLogin;post:UserLoginDo"),
		// beego.NSRouter("/logout", &admin.UserController{}, "get:UserLogout"),

	)
	beego.AddNamespace(ns1)

	beego.InsertFilter("/api/*", beego.BeforeExec, AccessControllAllow)

}
