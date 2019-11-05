package main

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:root@/orm_test?charset=utf8")
}

func main() {
	o := orm.NewOrm()
	o.Using("default")

	profile := new(Profile)
	profile.Age = 25

	user := new(User)
	user.Profile = profile
	user.Name = "xsky"

	fmt.Println(o.Insert(profile))
	fmt.Println(o.Insert(user))
}
