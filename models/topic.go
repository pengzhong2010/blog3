package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/favframework/debug"
	// "strconv"
	"photolimit/common"
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
		o.QueryTable("topic").Filter("deleted", 0).OrderBy("id").All(topics)
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

func (c *Topic) TopicRead() bool {
	o := orm.NewOrm()
	topic := Topic{Id: c.Id}
	if rerr := o.Read(&topic); rerr != nil {
		c.Id = 0
		return false
	}
	// o.QueryTable("Category").Filter("Id", c.Id).RelatedSel().One(c)
	rerr := o.Read(c)
	if rerr != nil {
		return false
	}
	// godump.Dump(c)
	// o.QueryTable("Category").Filter("Topics__Topic__Id", c.Id).All(categorys)
	// godump.Dump(cerr)
	// o.QueryTable("Image").Filter("Topics__Topic__Id", c.Id).All(images)
	// godump.Dump(ierr)
	return true
}

func (c *Topic) GetTopicCategorys() (bool, []orm.Params) {
	o := orm.NewOrm()
	var category_lists1 []orm.Params
	var category_lists []orm.Params
	topic := Topic{Id: c.Id}
	if rerr := o.Read(&topic); rerr != nil {
		c.Id = 0
		return false, category_lists
	}

	o.QueryTable("Category").Filter("Topics__Topic__Id", c.Id).Values(&category_lists1)

	o.QueryTable("Category").Values(&category_lists)
	s := make([]int, 1, len(category_lists))
	for _, v := range category_lists1 {
		v_id := int(v["Id"].(int64))
		s = append(s, v_id)
		// godump.Dump(v_id)
	}
	// godump.Dump(s)
	// godump.Dump(category_lists1)
	// godump.Dump(category_lists)
	for k, v := range category_lists {
		v_id := int(v["Id"].(int64))
		b, _ := common.Int_in_array(v_id, s)
		// godump.Dump(b)
		category_lists[k]["b_in"] = b
		// godump.Dump(category_lists[k])
	}
	return true, category_lists

}

func (c *Topic) GetTopicImages() (bool, []orm.Params) {
	o := orm.NewOrm()
	var images []orm.Params
	topic := Topic{Id: c.Id}
	if rerr := o.Read(&topic); rerr != nil {
		c.Id = 0
		return false, images
	}
	o.QueryTable("Image").Filter("Topics__Topic__Id", c.Id).Values(&images)
	// godump.Dump(images)
	return true, images
}

func (c *Topic) TopicEdit() bool {
	o := orm.NewOrm()
	topic := Topic{Id: c.Id}
	if rerr := o.Read(&topic); rerr != nil {
		c.Id = 0
		return false
	}
	if _, uerr := o.Update(c); uerr == nil {
		return true
	}
	return false
}

func (c *Topic) UpdateTopicCategorys(categorys []int) bool {
	o := orm.NewOrm()
	topic := Topic{Id: c.Id}
	if rerr := o.Read(&topic); rerr != nil {
		c.Id = 0
		return false
	}
	m2m := o.QueryM2M(c, "Categorys")
	m2m.Clear()
	godump.Dump(m2m)
	for _, v := range categorys {
		cate := Category{Id: v}
		o.Read(&cate)
		m2m.Add(&cate)
	}

	return true
}

func (c *Topic) AddImage(image *Image) bool {
	o := orm.NewOrm()
	topic := Topic{Id: c.Id}
	if rerr := o.Read(&topic); rerr != nil {
		c.Id = 0
		return false
	}
	m2m := o.QueryM2M(c, "Images")
	m2m.Add(image)
	return true
}

func (c *Topic) DelImage(image *Image) bool {
	o := orm.NewOrm()
	topic := Topic{Id: c.Id}
	if rerr := o.Read(&topic); rerr != nil {
		c.Id = 0
		return false
	}
	m2m := o.QueryM2M(c, "Images")
	m2m.Remove(image)
	image.ImageDel()
	return true
}

func (c *Topic) TopicDel() bool {
	o := orm.NewOrm()
	topic := Topic{Id: c.Id}
	if o.Read(&topic) == nil {
		o.Read(c)
		c.Deleted = true
		if _, err := o.Update(c); err == nil {
			return true
		}
	}

	return false
}
