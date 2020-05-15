package dao

import (
	"../models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"os"
	"strconv"
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
		_ = fmt.Errorf("register driver %v", err1)
	}
	err := orm.RegisterDataBase("default", "mysql", dataSource)
	if err != nil {
		fmt.Println("here2")
		_ = fmt.Errorf("register database %v", err)
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
		_ = fmt.Errorf("run syncdb %v", err2)
	}
	orm.RunCommand()
}

func (myOrm MyOrm) GetAllUsers() ([]models.User, error) {
	o := myOrm.O
	var users []models.User
	_, err := o.QueryTable("user").All(&users)
	if err != nil {
		_ = fmt.Errorf("get all users fail: %v", err)
	}
	return users, err
}

func (myOrm MyOrm) GetAllUsersByMoney() ([]models.User, error) {
	o := myOrm.O
	var users []models.User
	_, err := o.QueryTable("user").OrderBy("-money").All(&users)
	if err != nil {
		_ = fmt.Errorf("get all users fail: %v", err)
	}
	return users, err
}

func (myOrm MyOrm) GetAllUsersByBuildScore() ([]models.User, error) {
	o := myOrm.O
	var users []models.User
	_, err := o.QueryTable("user").OrderBy("-build_score").All(&users)
	if err != nil {
		_ = fmt.Errorf("get all users fail: %v", err)
	}
	return users, err
}

func (myOrm MyOrm) GetAllUsersByGameScore() ([]models.User, error) {
	o := myOrm.O
	var users []models.User
	_, err := o.QueryTable("user").OrderBy("-game_score").All(&users)
	if err != nil {
		_ = fmt.Errorf("get all users fail: %v", err)
	}
	return users, err
}

func (myOrm MyOrm) GetAllUsersByTotalScore() ([]models.User, error) {
	o := myOrm.O
	var users []models.User
	_, err := o.QueryTable("user").OrderBy("-total_score").All(&users)
	if err != nil {
		_ = fmt.Errorf("get all users fail: %v", err)
	}
	return users, err
}

func (myOrm MyOrm) GetUserByMacAddr(macAddr string) (models.User, error) {
	o := myOrm.O
	user := models.User{}
	err := o.QueryTable("user").Filter("mac_address", macAddr).One(&user)
	if err != nil {
		_ = fmt.Errorf("get user by mac address fail: %v", err)
	}
	return user, err
}

func (myOrm MyOrm) AddUser(userName string, macAddr string) (models.User, error) {
	o := myOrm.O
	user := models.User{UserName:userName, MacAddress:macAddr, Money:0,
		BuildScore:0, GameScore:0, TotalScore:0, SlotNum:5}
	_, err := o.Insert(&user)
	return user, err
}

func (myOrm MyOrm) UpdateUser(userId int, money int, buildScore int, gameScore int,
	totalScore int, slotNum int) (models.User, error) {
	o := myOrm.O
	user := models.User{UserId: userId}
	if o.Read(&user) == nil{
		user.Money += money
		user.BuildScore += buildScore
		user.GameScore += gameScore
		user.TotalScore += totalScore
		user.SlotNum += slotNum
		_, err := o.Update(&user)
		if err != nil {
			_ = fmt.Errorf("update user by user id fail: %v", err)
		}
	}
	return user, nil
}

func (myOrm MyOrm) DeleteUser(userId int) (models.User, error) {
	o := myOrm.O
	user := models.User{UserId:userId}
	if o.Read(&user) == nil{
		_, err := o.Delete(&user)
		if err != nil {
			_ = fmt.Errorf("delete level by level id fail: %v", err)
		}
	}
	return user, nil
}




func Exists(path string) bool {
	_, err := os.Stat(path)    //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

func (myOrm MyOrm) GetAllLevels() ([]models.LevelWithImage, error) {
	o := myOrm.O
	var levels []models.Level
	_, err := o.QueryTable("level").All(&levels)
	if err != nil {
		_ = fmt.Errorf("get all levels fail: %v", err)
	}
	var levelsWithImages []models.LevelWithImage
	for _, level := range levels {
		fileName := "Screenshots/Map_" + strconv.Itoa(level.LevelId) + "_Screenshot.png"
		var f []byte
		if Exists(fileName){
			f, _ = ioutil.ReadFile(fileName)
		}
		limage := models.LevelWithImage{LevelId:level.LevelId, TryNum:level.TryNum,
			PassNum:level.PassNum, ThumbNum:level.ThumbNum, IdOfMaker:level.IdOfMaker, MapData:level.MapData,
			OneStarStep:level.OneStarStep, TwoStarStep:level.TwoStarStep, ThreeStarStep:level.ThreeStarStep,
			LevelName:level.LevelName, Pic:f}

		levelsWithImages = append(levelsWithImages, limage)
	}
	return levelsWithImages, err
}

func (myOrm MyOrm) GetAllLevelsByTry() ([]models.Level, error) {
	o := myOrm.O
	var level []models.Level
	_, err := o.QueryTable("level").OrderBy("-try_num").All(&level)
	if err != nil {
		_ = fmt.Errorf("get all levels fail: %v", err)
	}
	return level, err
}

func (myOrm MyOrm) GetAllLevelsByPass() ([]models.Level, error) {
	o := myOrm.O
	var level []models.Level
	_, err := o.QueryTable("level").OrderBy("-pass_num").All(&level)
	if err != nil {
		_ = fmt.Errorf("get all levels fail: %v", err)
	}
	return level, err
}

func (myOrm MyOrm) GetAllLevelsByThumb() ([]models.Level, error) {
	o := myOrm.O
	var level []models.Level
	_, err := o.QueryTable("level").OrderBy("-thumb_num").All(&level)
	if err != nil {
		_ = fmt.Errorf("get all levels fail: %v", err)
	}
	return level, err
}

func (myOrm MyOrm) GetLevelByMakerId(makerId int) ([]models.Level, error) {
	o := myOrm.O
	var level []models.Level
	_, err := o.QueryTable("level").Filter("id_of_maker", makerId).All(&level)
	if err != nil {
		_ = fmt.Errorf("get level by maker id fail: %v", err)
	}
	return level, err
}

func (myOrm MyOrm) GetLevelsByLevelId(levelId int) (models.Level, error) {
	o := myOrm.O
	var level models.Level
	err := o.QueryTable("level").Filter("level_id", levelId).One(&level)
	if err != nil {
		_ = fmt.Errorf("get level by level id fail: %v", err)
	}
	return level, err
}

func (myOrm MyOrm) AddLevel(tryNum int, passNum int, thumbNum int, makerId int,
	mapData string, oneStarStep int, twoStarStep int, threeStarStep int, levelName string) (models.Level, error) {
	o := myOrm.O
	level := models.Level{TryNum:tryNum, PassNum:passNum, ThumbNum:thumbNum,
		MapData:mapData, IdOfMaker:makerId,
		OneStarStep:oneStarStep, TwoStarStep:twoStarStep, ThreeStarStep:threeStarStep, LevelName:levelName}
	// _, _ = myOrm.UpdateUser(makerId, 0, 0, 0, 0, -1)
	_, err := o.Insert(&level)
	return level, err
}

func (myOrm MyOrm) UpdateLevel(levelId int, try bool, pass bool, thumb bool,
	oneStarStep int, twoStarStep int, threeStarStep int) (models.Level, error) {
	o := myOrm.O
	level := models.Level{LevelId: levelId}
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
			_ = fmt.Errorf("update level by level id fail: %v", err)
		}
	}
	return level, nil
}

func (myOrm MyOrm) DeleteLevel(levelId int) (models.Level, error) {
	o := myOrm.O
	level := models.Level{LevelId: levelId}
	if o.Read(&level) == nil{
		_, err := o.Delete(&level)
		if err != nil {
			_ = fmt.Errorf("delete level by level id fail: %v", err)
		}
	}
	return level, nil
}
