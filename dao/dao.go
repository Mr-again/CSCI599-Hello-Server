package dao

import (
	"../models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
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
	fmt.Println(dataSource)
	err1 := orm.RegisterDriver("mysql", orm.DRMySQL)
	if err1 != nil {
		fmt.Println("here1")
		fmt.Errorf("Register Driver %v", err1)
	}
	err := orm.RegisterDataBase("default", "mysql", dataSource)
	if err != nil {
		fmt.Println("here2")
		fmt.Errorf("Register Database %v", err)
	}
	mdb, err := orm.GetDB("default")
	if err != nil {
		panic(fmt.Errorf("get db error:%s", err))
	}
	mdb.SetConnMaxLifetime(time.Duration(4 * 3600) * time.Second)
	mdb.SetMaxIdleConns(10)
	mdb.SetMaxOpenConns(30)
	err2 := orm.RunSyncdb("default", false, false)
	if err2 != nil {
		fmt.Println("here3")
		fmt.Errorf("Run Syncdb %v", err2)
	}
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
	_, err := o.QueryTable("level").Filter("id_of_maker", makerId).All(&level)
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

func (myOrm MyOrm) GetAllLevels() ([]models.Level, error) {
	o := myOrm.O
	var level []models.Level
	_, err := o.QueryTable("level").All(&level)
	if err != nil {
		fmt.Errorf("get all levels fail: %v", err)
	}
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

func (myOrm MyOrm) AddLevel(tryNum int, passNum int, thumbNum int, makerId int, mapData string, oneStarStep int, twoStarStep int, threeStarStep int) (models.Level, error) {
	o := myOrm.O
	level := models.Level{TryNum:tryNum, PassNum:passNum, ThumbNum:thumbNum, MapData:mapData, IdOfMaker:makerId, OneStarStep:oneStarStep, TwoStarStep:twoStarStep, ThreeStarStep:threeStarStep}
	//myOrm.UpdateUser(makerId, 0, 0, 0, 0, -1)
	_, err := o.Insert(&level)
	return level, err
}

func (myOrm MyOrm) UpdateUser(user_id int, money int, build_score int, game_score int, total_score int, slot_num int) (models.User, error) {
	o := myOrm.O
	user := models.User{UserId:user_id}
	if o.Read(&user) == nil{
		user.Money += money
		user.BuildScore += build_score
		user.GameScore += game_score
		user.TotalScore += total_score
		user.SlotNum += slot_num
		_, err := o.Update(&user)
		if err != nil {
			fmt.Errorf("update user by user id fail: %v", err)
		}
	}
	return user, nil
}

func (myOrm MyOrm) UpdateLevel(level_id int, try bool, pass bool, thumb bool, oneStarStep int, twoStarStep int, threeStarStep int) (models.Level, error) {
	o := myOrm.O
	level := models.Level{LevelId:level_id}
	if o.Read(&level) == nil{
		if try{
			level.TryNum++
		}
		if pass{
			level.PassNum++
		}
		if thumb{
			level.ThumbNum++
		}
		if oneStarStep != 0 {
			level.OneStarStep = oneStarStep
		}
		if twoStarStep != 0 {
			level.TwoStarStep = twoStarStep
		}
		if threeStarStep != 0 {
			level.ThreeStarStep = threeStarStep
		}
		_, err := o.Update(&level)
		if err != nil {
			fmt.Errorf("update level by level id fail: %v", err)
		}
	}
	return level, nil
}

func (myOrm MyOrm) DeleteLevel(level_id int) (models.Level, error) {
	o := myOrm.O
	level := models.Level{LevelId:level_id}
	if o.Read(&level) == nil{
		_, err := o.Delete(&level)
		if err != nil {
			fmt.Errorf("delete level by level id fail: %v", err)
		}
	}
	return level, nil
}
