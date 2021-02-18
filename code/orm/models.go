package models

import (
	"github.com/astaxie/beego/orm" // orm: object relational mapping
)

type User struct {
	Id      int
	Name    string
	Profile *Profile `orm:"rel(one)"`
	Post    []*Post  `orm:"reverse(many)"`
}

type Profile struct {
	Id   int
	Age  int16
	User *User `orm:"reverse(one)"`
}

type Post struct {
	Id    int
	Title string
	User  *User  `orm:"rel(fk)"`
	Tag   []*Tag `orm:"rel(m2m)"`
}

type Tag struct {
	Id   int
	Name string
	Post []*Post `orm:"reverse(many)"`
}

func init() {
	orm.RegisterModel(new(User), new(Profile), new(Post), new(Tag))
}
