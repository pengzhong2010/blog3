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
	Topic     []*Topic `orm:"reverse(many)"`
	Entry     []*Entry `orm:"reverse(many)"`

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
	Id    int
	Name  string   `orm:"unique"`
	Topic []*Topic `orm:"reverse(many)"`
	Image *Image   `orm:"null;rel(one);on_delete(do_nothing)"`
}

type Topic struct {
	Id       int
	Name     string
	Content  string      `orm:"null;type(text)"`
	User     *User       `orm:"rel(fk);ondelete(do_nothing)"`
	Category []*Category `orm:"rel(m2m);rel_table(topic_category);ondelete(do_nothing)"`
	Entry    []*Entry    `orm:"reverse(many)"`
	Image    []*Image    `orm:"rel(m2m);rel_table(topic_image);ondelete(do_nothing)"`

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
	Topic    []*Topic  `orm:"reverse(many)"`

	Deleted bool `orm:"default(false)"`
}

func init() {
	orm.RegisterModel(new(User), new(Category), new(Topic), new(Entry), new(Image))
}
