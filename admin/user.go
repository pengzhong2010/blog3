package admin

import (
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"github.com/favframework/debug"
	// "log"
	"photolimit/models"
	"strconv"
	// "reflect"
	// "strings"
)

var UserModel models.User

type UserController struct {
	beego.Controller
}

func (c *UserController) UserList() {
	q := c.Ctx.Input.Query("q")
	users := []models.User{}
	u := models.User{}
	u.UserList(&users, q)
	godump.Dump(users)
	// c.Data["Website"] = "beego.me"
	// c.Data["Email"] = "astaxie@gmail.com"
	c.Data["users"] = users
	c.Data["Title"] = "UserList"
	c.Layout = "admin/layout.html"
	c.TplName = "admin/userlist.html"

}

type UserForm struct {
	Email string `form:"email" valid:"Required;Email;MaxSize(50)"`
	Pwd   string `form:"pwd" valid:"Required;MaxSize(30)"`
	Name  string `form:"name" valid:"Required;MaxSize(60)"`
}

func (u *UserForm) Valid(v *validation.Validation) {
	// if strings.Index(u.Name, "admin1") != -1 {
	// 	v.SetError("Name", "名称里不能含有 admin1")
	// }
}

func (c *UserController) UserAdd() {
	c.Layout = "admin/layout.html"
	c.TplName = "admin/useradd.html"
}

func (c *UserController) UserAddDo() {
	uf := UserForm{}
	c.ParseForm(&uf)
	valid := validation.Validation{}
	b, _ := valid.Valid(&uf)
	if !b {
		c.Ctx.Output.Body([]byte("valid error"))
		c.StopRun()
	}
	userdb := models.User{Email: uf.Email, Pwd: uf.Pwd, Name: uf.Name}
	res_b, _ := userdb.UserAdd()
	if !res_b {
		c.Ctx.Output.Body([]byte("email exist"))
		c.StopRun()
	}
	c.Ctx.Redirect(302, c.URLFor(".UserList"))
}

func (c *UserController) UserEdit() {
	id := c.Ctx.Input.Param(":id")
	idint, _ := strconv.Atoi(id)
	u := models.User{Id: idint}
	u.UserRead()
	if u.Id == 0 {
		c.Ctx.Output.Body([]byte("not found"))
		c.StopRun()
	}
	c.Data["user"] = u
	c.Layout = "admin/layout.html"
	c.TplName = "admin/useradd.html"
}

func (c *UserController) UserEditDo() {
	uf := UserForm{}
	c.ParseForm(&uf)
	godump.Dump(uf)

	valid := validation.Validation{}
	b, _ := valid.Valid(&uf)
	if !b {
		c.Ctx.Output.Body([]byte("valid error"))
		c.StopRun()
	}
	id := c.Ctx.Input.Param(":id")
	idint, _ := strconv.Atoi(id)
	um := models.User{Id: idint, Pwd: uf.Pwd, Name: uf.Name}
	res_b := um.UserEdit()
	if !res_b {
		c.Ctx.Output.Body([]byte("error"))
		c.StopRun()
	}
	c.Ctx.Redirect(302, c.URLFor(".UserList"))

}

func (c *UserController) UserDel() {
	id := c.Ctx.Input.Param(":id")
	idint, _ := strconv.Atoi(id)
	u := models.User{Id: idint}
	b := u.UserDel()
	if !b {
		c.Ctx.Output.Body([]byte("del error"))
		c.StopRun()
	}
	c.Ctx.Redirect(302, c.URLFor(".UserList"))
}

type LoginForm struct {
	Email  string `form:"email" valid:"Required;Email;MaxSize(50)"`
	Pwd    string `form:"pwd" valid:"Required;MaxSize(30)"`
	Cookie bool   `form:"cookie"`
}

func (c *UserController) UserLogin() {
	// aa := c.Ctx.GetCookie("user")
	bb := c.Ctx.Input.Session("user")
	// godump.Dump(aa)
	godump.Dump("bb")
	godump.Dump(bb)
	// aaa := c.Ctx.Input.URL()
	// godump.Dump(aaa)
	beego.ReadFromRequest(&c.Controller)
	c.Layout = "admin/layout.html"
	c.TplName = "admin/login.html"
}

func (c *UserController) UserLoginDo() {
	lf := LoginForm{}
	c.ParseForm(&lf)
	// godump.Dump(lf)

	valid := validation.Validation{}
	b, _ := valid.Valid(&lf)
	if !b {
		c.Ctx.Output.Body([]byte("valid error"))
		c.StopRun()
	}

	u := models.User{Email: lf.Email, Pwd: lf.Pwd}
	b_login := u.AdminLogin()
	if !b_login {
		flash := beego.NewFlash()
		flash.Error("Email or Password error")
		flash.Store(&c.Controller)
		c.Ctx.Redirect(302, c.URLFor(".UserLogin"))
		c.StopRun()
	}
	c.Ctx.Output.Session("user", u)
	// c.Ctx.SetCookie("user", string(u.Id), 3600, "/")
	redirect := c.Ctx.Input.Query("redirect")
	if len(redirect) == 0 {
		c.Ctx.Redirect(302, c.URLFor(".UserList"))
	} else {
		c.Ctx.Redirect(302, redirect)
	}

	c.StopRun()

}

func (c *UserController) UserLogout() {
	c.Ctx.Output.Session("user", nil)
	c.Ctx.Redirect(302, c.URLFor(".UserLogin"))
}
