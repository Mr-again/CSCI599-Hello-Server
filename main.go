package main

import (
	_ "./Services"
	"github.com/astaxie/beego"
)



func init() {
	beego.Run()
}

func main() {
	//user, _ := myOrm.GetUserByMacAddr("mac2")
	//fmt.Println(user)

}
