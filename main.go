package main

import (
	_ "./dao"
	"./models"
	"fmt"
	"github.com/astaxie/beego/orm"
)

var (
	DB orm.Ormer
)

func main() {
	DB = orm.NewOrm()
	DB.Using("default")
	user := models.User{UserId: 1}
	err := DB.Read(&user)
	if err == orm.ErrNoRows {
		fmt.Println("查询不到")
	} else if err == orm.ErrMissPK {
		fmt.Println("找不到主键")
	} else {
		fmt.Println(user.UserId, user.UserName)
	}
}
