package models

import (
	"github.com/astaxie/beego/orm"
	// "github.com/favframework/debug"
)

func (i *Image) ImageCreate() bool {
	o := orm.NewOrm()
	_, err := o.Insert(i)
	if err != nil {
		return false
	}
	return true
}

func (i *Image) ImageUpdate() bool {
	o := orm.NewOrm()
	img := Image{Id: i.Id}
	if o.Read(&img) == nil {
		if _, err := o.Update(i); err == nil {
			return true
		}
	}

	return false
}

func (i *Image) ImageRead() bool {
	o := orm.NewOrm()
	img := Image{Id: i.Id}
	if rerr := o.Read(&img); rerr != nil {
		i.Id = 0
		return false
	}

	o.Read(i)

	return true
}

func (i *Image) ImageDel() bool {
	o := orm.NewOrm()
	img := Image{Id: i.Id}
	if rerr := o.Read(&img); rerr == nil {
		o.Read(i)
		i.Deleted = true
		if _, err := o.Update(i); err == nil {
			return true
		}
	}

	return false
}
