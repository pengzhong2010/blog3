package models

import (
	_ "github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
)

type User struct {
	Id        int
	Email     string `orm:"size(50);unique"`
	Name      string `orm:"size(60)"`
	Pwd       string
	AdminRole bool     `orm:"default(false)"`
	Topics    []*Topic `orm:"reverse(many)"`
	Entrys    []*Entry `orm:"reverse(many)"`

	Created time.Time `orm:"auto_now_add;type(datetime)"`
	Updated time.Time `orm:"auto_now;type(datetime)"`
	Deleted bool      `orm:"default(false)"`
}

func (u *User) TableIndex() [][]string {
	return [][]string{
		[]string{"Name"},
	}
}

type Category struct {
	Id     int
	Name   string   `orm:"unique"`
	Topics []*Topic `orm:"reverse(many)"`
	Image  *Image   `orm:"null;rel(one);on_delete(do_nothing)"`
}

type Topic struct {
	Id        int
	Name      string
	Content   string      `orm:"null;type(text)"`
	User      *User       `orm:"rel(fk);ondelete(do_nothing)"`
	Categorys []*Category `orm:"rel(m2m);rel_table(topic_category);ondelete(do_nothing)"`
	Entry     []*Entry    `orm:"reverse(many)"`
	Images    []*Image    `orm:"rel(m2m);rel_table(topic_image);ondelete(do_nothing)"`

	Created time.Time `orm:"auto_now_add;type(datetime)"`
	Updated time.Time `orm:"auto_now;type(datetime)"`
	Deleted bool      `orm:"default(false)"`
}

func (t *Topic) TableIndex() [][]string {
	return [][]string{
		[]string{"Name"},
	}
}

type Entry struct {
	Id      int
	Content string `orm:"type(text)"`
	User    *User  `orm:"rel(fk);ondelete(do_nothing)"`
	Topic   *Topic `orm:"rel(fk);ondelete(do_nothing)"`

	Created time.Time `orm:"auto_now_add;type(datetime)"`
	Updated time.Time `orm:"auto_now;type(datetime)"`
	Deleted bool      `orm:"default(false)"`
}

type Image struct {
	Id       int
	Url      string    `orm:"null;size(300)"`
	Category *Category `orm:"reverse(one)"`
	Topics   []*Topic  `orm:"reverse(many)"`
	Banner   *Banner   `orm:"reverse(one)"`

	Deleted bool `orm:"default(false)"`
}

type Banner struct {
	Id      int
	Name    string
	Url     string `orm:"null;size(300)"`
	Image   *Image `orm:"null;rel(one);on_delete(do_nothing)"`
	Deleted bool   `orm:"default(false)"`
}

func init() {
	orm.RegisterModel(new(User), new(Category), new(Topic), new(Entry), new(Image), new(Banner))
}
