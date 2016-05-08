package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"photolimit/admin"
	"photolimit/controllers"
)

var LoginCheck = func(ctx *context.Context) {
	user := ctx.Input.Session("user")
	if user == nil {
		url := ctx.Input.URL()
		if url != "/admin/login" {
			ctx.Redirect(302, "/admin/login")
		}
	}
}

func init() {
	beego.Router("/", &controllers.MainController{})

	ns := beego.NewNamespace("/admin",
		beego.NSRouter("/login", &admin.UserController{}, "get:UserLogin;post:UserLoginDo"),
		beego.NSRouter("/logout", &admin.UserController{}, "get:UserLogout"),
		beego.NSNamespace("/user",
			beego.NSRouter("/list", &admin.UserController{}, "get:UserList"),
			beego.NSRouter("/add", &admin.UserController{}, "get:UserAdd;post:UserAddDo"),
			beego.NSRouter("/edit/:id([0-9]+)", &admin.UserController{}, "get:UserEdit;post:UserEditDo"),
			beego.NSRouter("/del/:id([0-9]+)", &admin.UserController{}, "get:UserDel"),
		),
		beego.NSNamespace("/category",
			beego.NSRouter("/list", &admin.CategoryController{}, "get:CategoryList"),
			beego.NSRouter("/add", &admin.CategoryController{}, "get:CategoryAdd;post:CategoryAddDo"),
			beego.NSRouter("/edit/:id([0-9]+)", &admin.CategoryController{}, "get:CategoryEdit;post:CategoryEditDo"),
			beego.NSRouter("/edit/:id([0-9]+)/img", &admin.CategoryController{}, "post:CategoryEditImg"),
			// beego.NSRouter("/del/:id([0-9]+)", &admin.CategoryController{}, "post:CategoryDel"),
		),
		// beego.NSNamespace("/topic",
		// 	beego.NSRouter("/list", &controllers.MainController{}, "get:Get"),
		// 	beego.NSRouter("/add", &controllers.MainController{}, "get:Get;post:Get"),
		// 	beego.NSRouter("/edit/:id", &controllers.MainController{}, "get:Get;post:Get"),
		// 	beego.NSRouter("/del/:id", &controllers.MainController{}, "post:Get"),
		// ),
		beego.NSRouter("/base", &controllers.MainController{}, "get:Get"),
	)
	beego.AddNamespace(ns)

	// beego.InsertFilter("/admin/*", beego.BeforeExec, LoginCheck)

}
