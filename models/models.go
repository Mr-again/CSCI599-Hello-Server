package models

import (
	"github.com/astaxie/beego/orm"
)

type User struct {
	UserId     int `orm:"PK;auto"`
	UserName   string
	MacAddress string
	Money      int
	BuildScore int
	GameScore  int
	TotalScore int
	SlotNum    int
	//Levels     []*Level `orm:"reverse(many)"`
}

type Level struct {
	LevelId  int `orm:"PK;auto"`
	TryNum   int
	PassNum  int
	ThumbNum int
	IdOfMaker  int
	//Maker    *User `orm:"rel(fk)"`
	MapData  string
	OneStarStep int
	TwoStarStep int
	ThreeStarStep int
	LevelName string
}

type LevelWithImage struct {
	LevelId  int
	TryNum   int
	PassNum  int
	ThumbNum int
	IdOfMaker  int
	MapData  string
	OneStarStep int
	TwoStarStep int
	ThreeStarStep int
	LevelName string
	Pic []byte
}


func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(User), new(Level))
}
