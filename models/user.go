package models

import (
	"github.com/astaxie/beego/orm"
	// "github.com/favframework/debug"
	// "strconv"
)

func (u *User) UserList(users *[]User, q string) {
	o := orm.NewOrm()
	qs := o.QueryTable("user")
	if len(q) != 0 {
		cond := orm.NewCondition()
		cond1 := cond.And("deleted", 0)
		cond2 := cond.AndCond(cond1).AndCond(cond.And("email", q).Or("name", q))
		qs.SetCond(cond2).All(users)
		// o.QueryTable("user").Filter("deleted", 0).OrderBy("id").All(users)
	} else {
		o.QueryTable("user").Filter("deleted", 0).OrderBy("id").All(users)
	}
}
func (u *User) UserRead() bool {
	o := orm.NewOrm()
	if rerr := o.Read(u); rerr != nil {
		u.Id = 0
		return false
	}
	return true
}
func (u *User) UserAdd() (bool, int) {
	res_b := false
	res_id := 0
	o := orm.NewOrm()
	if created, id, rerr := o.ReadOrCreate(u, "Email"); rerr == nil {
		if created {
			res_b = true
			res_id = int(id)
		}
	}
	return res_b, res_id
}

func (u *User) UserEdit() bool {
	res_b := false
	o := orm.NewOrm()
	user := User{Id: u.Id}
	if rerr := o.Read(&user); rerr == nil {

		user.Pwd = u.Pwd
		user.Name = u.Name
		if _, uerr := o.Update(&user); uerr == nil {
			res_b = true
		}
	}
	return res_b
}

func (u *User) UserDel() bool {
	res_b := false
	o := orm.NewOrm()
	if rerr := o.Read(u); rerr == nil {
		u.Deleted = true
		if _, uerr := o.Update(u); uerr == nil {
			res_b = true
		}
	}
	return res_b
}

func (u *User) AdminLogin() bool {
	// godump.Dump(u)
	res_b := false
	user := User{Email: u.Email, Pwd: u.Pwd}
	o := orm.NewOrm()
	o.QueryTable("user").Filter("Email", u.Email).Filter("deleted", 0).Filter("AdminRole", 1).One(u)
	if u.Id != 0 {
		if user.Pwd == u.Pwd {
			res_b = true
		}
	}
	// godump.Dump(u)
	// godump.Dump(res_b)
	return res_b
}
