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
		fmt.Errorf("get user by mac address fail: %v", err)
	}
	return user, err
}

func (myOrm MyOrm) GetUsersByBuildScore(buildScore int) ([]models.User, error) {
	o := myOrm.O
	var user []models.User
	_, err := o.QueryTable("user").Filter("build_score", buildScore).All(&user)
	if err != nil {
		fmt.Errorf("get user by build score fail: %v", err)
	}
	return user, err
}

func (myOrm MyOrm) GetUsersByGameScore(gameScore int) ([]models.User, error) {
	o := myOrm.O
	var user []models.User
	_, err := o.QueryTable("user").Filter("game_score", gameScore).All(&user)
	if err != nil {
		fmt.Errorf("get user by game score fail: %v", err)
	}
	return user, err
}

func (myOrm MyOrm) GetUsersByTotalScore(totalScore int) ([]models.User, error) {
	o := myOrm.O
	var user []models.User
	_, err := o.QueryTable("user").Filter("total_score", totalScore).All(&user)
	if err != nil {
		fmt.Errorf("get user by total score fail: %v", err)
	}
	return user, err
}

func (myOrm MyOrm) GetLevelByMakerId(makerId int) ([]models.Level, error) {
	o := myOrm.O
	var level []models.Level
	_, err := o.QueryTable("level").Filter("maker_id", makerId).All(&level)
	if err != nil {
		fmt.Errorf("get level by maker id fail: %v", err)
	}
	//for _, l := range level {
	//	err := o.QueryTable("user").Filter("user_id", makerId).One(l.Maker)
	//	if err != nil {
	//		fmt.Errorf("do not exist this user with id %d: %v", makerId, err)
	//	}
	//}
	return level, err
}

func (myOrm MyOrm) GetLevelsByLevelId(levelId int) (models.Level, error) {
	o := myOrm.O
	var level models.Level
	err := o.QueryTable("level").Filter("level_id", levelId).One(&level)
	if err != nil {
		fmt.Errorf("get level by level id fail: %v", err)
	}
	//err = o.QueryTable("user").Filter("user_id", level).One(level.Maker)
	//if err != nil {
	//	fmt.Errorf("do not exist this user with id %d: %v", level.Maker.UserId, err)
	//}
	return level, err
}

func (myOrm MyOrm) AddUser(userName string, macAddr string, money int,
	buildScore int, gameScore int, totalScore int, slotNum int) (models.User, error) {
	o := myOrm.O
	user := models.User{UserName:userName, MacAddress:macAddr, Money:money,
		BuildScore:buildScore, GameScore:gameScore, TotalScore:totalScore, SlotNum:slotNum}
	_, err := o.Insert(&user)
	return user, err
}

func (myOrm MyOrm) AddLevel(tryNum int, passNum int, thumbNum int, makerId int, mapData string) (models.Level, error) {
	o := myOrm.O
	level := models.Level{TryNum:tryNum, PassNum:passNum, ThumbNum:thumbNum, MapData:mapData, IdOfMaker:makerId}
	_, err := o.Insert(&level)
	return level, err
}

func (myOrm MyOrm) UpdateUser(user models.User, money int, buildScore int, gameScore int, totalScore int, slotNum int) (models.User, error) {
	return models.User{}, nil
}

func (myOrm MyOrm) UpdateLevel(level models.Level, tryNum int, passNum int, thumbNum int) (models.Level, error) {
	return models.Level{}, nil
}
