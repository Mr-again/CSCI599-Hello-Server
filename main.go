package main

import (
	_ "./Services"
	"github.com/astaxie/beego"
)

func init() {
	beego.Run()
}

func main() {
	//user, _ := myOrm.GetUsersByBuildScore(100)
	//fmt.Println(user)

}
