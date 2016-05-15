package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/favframework/debug"
	// "strconv"
)

func GetCategorys(cs *[]*Category, q string) {
	o := orm.NewOrm()
	qs := o.QueryTable("category")
	if len(q) != 0 {
		cond := orm.NewCondition()
		// cond1 := cond.And("deleted", 0)
		cond2 := cond.AndCond(cond.And("name", q))
		qs.SetCond(cond2).All(cs)
		// o.QueryTable("user").Filter("deleted", 0).OrderBy("id").All(users)
	} else {
		o.QueryTable("category").OrderBy("id").All(cs)

	}
	for _, v := range *cs {
		o.LoadRelated(v, "Image")
	}

}

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
	// godump.Dump(categorys)
}

func AllCategorys(cs *[]orm.Params) bool {
	o := orm.NewOrm()
	_, err := o.QueryTable("Category").Values(cs)
	if err != nil {
		return false
	}
	return true
}

func (c *Category) CategoryRead() bool {
	o := orm.NewOrm()
	category := Category{Id: c.Id}
	if rerr := o.Read(&category); rerr != nil {
		c.Id = 0
		return false
	}
	// o.QueryTable("Category").Filter("Id", c.Id).RelatedSel().One(c)
	o.Read(c)
	godump.Dump(c)
	if c.Image != nil {
		err := o.Read(c.Image)
		godump.Dump(err)
	}

	return true

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

	if _, uerr := o.Update(c); uerr == nil {
		res_b = true
	}

	return res_b
}

func (c *Category) CategoryEditImg() bool {
	o := orm.NewOrm()
	category := Category{Id: c.Id}
	b := category.CategoryRead()
	if !b {
		return false
	}
	if c.Image == nil {
		return false
	}
	if category.Image != nil && category.Image.Id != c.Image.Id {
		category.Image.ImageDel()
	}
	category.Image = c.Image
	if _, uerr := o.Update(&category); uerr == nil {
		return true
	}
	return false
}
