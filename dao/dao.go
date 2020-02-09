package dao

import (
	_ "../models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)



func init() {
	mysqlUser := beego.AppConfig.String("mysqluser")
	mysqlPass := beego.AppConfig.String("mysqlpass")
	mysqlDb := beego.AppConfig.String("mysqldb")
	mysqlUrl := beego.AppConfig.String("mysqlurl")
	dataSource := mysqlUser + ":" + mysqlPass + "@" + "tcp(" + mysqlUrl + ")/" + mysqlDb + "?charset=utf8"
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", dataSource)
	orm.RunSyncdb("default", false, false)
	orm.RunCommand()
}
