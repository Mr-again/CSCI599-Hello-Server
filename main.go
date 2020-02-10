package main

import (
	"./dao"
	"fmt"
	"github.com/astaxie/beego/orm"
)

var (
	myOrm dao.MyOrm
)

func init() {
	myOrm = dao.MyOrm{}
	myOrm.O = orm.NewOrm()
	myOrm.O.Using("default")
}

func main() {
	user, _ := myOrm.GetUserByMacAddr("mac2")
	fmt.Println(user)
	//user := models.User{UserId: 1}
	//err := myOrm.O.Read(&user)
	//if err == orm.ErrNoRows {
	//	fmt.Println("查询不到")
	//} else if err == orm.ErrMissPK {
	//	fmt.Println("找不到主键")
	//} else {
	//	fmt.Println(user.UserId, user.UserName)
	//}
}
