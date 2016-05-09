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
