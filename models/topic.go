package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/favframework/debug"
	// "strconv"
)

func (t *Topic) TopicList(topics *[]Topic, q string) {
	o := orm.NewOrm()
	qs := o.QueryTable("topic")
	if len(q) != 0 {
		cond := orm.NewCondition()
		cond1 := cond.And("deleted", 0)
		cond2 := cond.AndCond(cond1).AndCond(cond.And("name", q))
		qs.SetCond(cond2).All(topics)
		// o.QueryTable("user").Filter("deleted", 0).OrderBy("id").All(users)
	} else {
		o.QueryTable("topic").OrderBy("id").All(topics)
	}
	godump.Dump(topics)
}

func (c *Topic) TopicAdd() (bool, int) {
	res_b := false
	res_id := 0
	o := orm.NewOrm()
	id, err := o.Insert(c)
	if err == nil {
		res_b = true
		res_id = int(id)
	} else {
		godump.Dump(err)
	}

	return res_b, res_id
}

func (c *Topic) TopicRead(categorys *[]Category, images *[]Image) bool {
	o := orm.NewOrm()
	topic := Topic{Id: c.Id}
	if rerr := o.Read(&topic); rerr != nil {
		c.Id = 0
		return false
	}
	// o.QueryTable("Category").Filter("Id", c.Id).RelatedSel().One(c)
	o.Read(c)
	godump.Dump(c)
	_, cerr := o.QueryTable("Category").Filter("Topics__Topic__Id", c.Id).All(categorys)
	godump.Dump(cerr)
	_, ierr := o.QueryTable("Image").Filter("Topics__Topic__Id", c.Id).All(images)
	godump.Dump(ierr)
	return true
}
