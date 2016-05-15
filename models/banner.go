package models

import (
	"github.com/astaxie/beego/orm"
	// "github.com/favframework/debug"
	// "strconv"
	// "log"
)

func GetBanners(banners *[]*Banner, q string) {
	o := orm.NewOrm()
	qs := o.QueryTable("banner")
	if len(q) != 0 {
		cond := orm.NewCondition()
		cond1 := cond.And("deleted", 0)
		cond2 := cond.AndCond(cond1).AndCond(cond.And("name", q))
		qs.SetCond(cond2).All(banners)
		// o.QueryTable("user").Filter("deleted", 0).OrderBy("id").All(users)
	} else {
		o.QueryTable("banner").Filter("deleted", 0).OrderBy("id").All(banners)

	}
	for _, v := range *banners {
		o.LoadRelated(v, "Image")
	}

}

func (b *Banner) BannerRead() bool {
	o := orm.NewOrm()
	banner := Banner{Id: b.Id}
	if rerr := o.Read(&banner); rerr != nil {
		b.Id = 0
		return false
	}
	o.Read(b)
	if b.Image != nil {
		o.Read(b.Image)
	}
	return true
}

func (b *Banner) BannerAdd() (bool, int) {
	res_b := false
	res_id := 0
	o := orm.NewOrm()
	if id, rerr := o.Insert(b); rerr == nil {
		res_b = true
		res_id = int(id)
	}
	return res_b, res_id
}

func (b *Banner) BannerEdit() bool {
	res_b := false
	o := orm.NewOrm()
	if _, uerr := o.Update(b); uerr == nil {
		res_b = true
	}

	return res_b
}

func (b *Banner) BannerDel() bool {
	res_b := false
	o := orm.NewOrm()
	b_tmp := b.BannerRead()
	if !b_tmp {
		return false
	}
	if b.Image != nil {
		b.Image.ImageDel()
	}
	b.Deleted = true
	if _, uerr := o.Update(b); uerr == nil {
		res_b = true
	}

	return res_b
}

func (b *Banner) BannerEditImg() bool {
	o := orm.NewOrm()
	banner := Banner{Id: b.Id}
	b_tmp := banner.BannerRead()
	if !b_tmp {
		return false
	}
	if b.Image == nil {
		return false
	}
	if banner.Image != nil && banner.Image.Id != b.Image.Id {
		banner.Image.ImageDel()
	}
	banner.Image = b.Image
	if _, uerr := o.Update(&banner); uerr == nil {
		return true
	}
	return false
}
