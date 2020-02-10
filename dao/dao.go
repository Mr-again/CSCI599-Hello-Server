package dao

import (
	"../models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type MyOrm struct {
	O orm.Ormer
}

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

func (myOrm MyOrm) GetUserByMacAddr(macAddr string) (models.User, error) {
	o := myOrm.O
	user := models.User{}
	err := o.QueryTable("user").Filter("mac_address", macAddr).One(&user)
	if err != nil {
		fmt.Errorf("GetUserByMacAddr fail: %v", err)
	}
	return user, err
}

func (myOrm MyOrm) GetUsersByBuildScore(page int) ([]models.User, error) {
	return nil, nil
}

func (myOrm MyOrm) GetUsersByGameScore(page int) ([]models.User, error) {
	return nil, nil
}

func (myOrm MyOrm) GetUsersByTotalScore(page int) ([]models.User, error) {
	return nil, nil
}

func (myOrm MyOrm) GetLevelByMakerId(makerId string) (models.Level, error) {
	return models.Level{}, nil
}

func (myOrm MyOrm) GetLevels(page int) ([]models.Level, error) {
	return nil, nil
}

func (myOrm MyOrm) AddUser(userName string, macAddr string, money int, buildScore int, gameScore int, totalScore int, slotNum int) (models.User, error) {
	return models.User{}, nil
}

func (myOrm MyOrm) AddLevel(tryNum int, passNum int, thumbNum int, makerId int, mapData string) (models.Level, error) {
	return models.Level{}, nil
}

func (myOrm MyOrm) UpdateUser(user models.User, money int, buildScore int, gameScore int, totalScore int, slotNum int) (models.User, error) {
	return models.User{}, nil
}

func (myOrm MyOrm) UpdateLevel(level models.Level, tryNum int, passNum int, thumbNum int) (models.Level, error) {
	return models.Level{}, nil
}
