package models

import (
	"github.com/astaxie/beego/orm"
	// "github.com/favframework/debug"
	// "strconv"
)

func (c *Category) CategoryList(categorys *[]Category, q string) {
	o := orm.NewOrm()
	qs := o.QueryTable("category")
	if len(q) != 0 {
		cond := orm.NewCondition()
		cond1 := cond.And("name", q)
		// cond2 := cond.AndCond(cond1).AndCond(cond.And("email", q).Or("name", q))
		qs.SetCond(cond1).All(categorys)
		// o.QueryTable("user").Filter("deleted", 0).OrderBy("id").All(users)
	} else {
		o.QueryTable("category").OrderBy("id").All(categorys)
	}
}

func (c *Category) CategoryRead() *CategoryImage {
	o := orm.NewOrm()
	if rerr := o.Read(c); rerr != nil {
		c.Id = 0
	} else {
		ci := &CategoryImage{}
		err := o.QueryTable("CategoryImage").Filter("Category__Id", c.Id).One(ci)
		if err == nil {
			return ci
		}
	}
	return nil
}

func (c *Category) CategoryAdd() (bool, int) {
	res_b := false
	res_id := 0
	o := orm.NewOrm()
	if created, id, rerr := o.ReadOrCreate(c, "Name"); rerr == nil {
		if created {
			res_b = true
			res_id = int(id)
		}
	}
	return res_b, res_id
}

func (c *Category) CategoryEdit() bool {
	res_b := false
	o := orm.NewOrm()
	category := Category{Id: c.Id}
	if rerr := o.Read(&category); rerr == nil {
		category.Name = c.Name
		if _, uerr := o.Update(&category); uerr == nil {
			res_b = true
		}
	}
	return res_b
}
