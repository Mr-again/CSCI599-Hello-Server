package Services

import (
	"../dao"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

var (
	myOrm dao.MyOrm
)

type UserController struct {
	beego.Controller
}

type LevelController struct {
	beego.Controller
}

func (ctrl *UserController) Get() {
	getType, err := ctrl.GetInt("type")
	if err != nil {
		fmt.Errorf("user get type fail, %v", err)
	}
	switch getType {
	case 0:
		{
			macAddr := ctrl.GetString("mac")
			user, err := myOrm.GetUserByMacAddr(macAddr)
			if err != nil {
				fmt.Errorf("GetUser fail, %v", err)
			}
			ctrl.Ctx.Output.JSON(user, true, true)
			break
		}
	case 1:
		{
			buildScore, _:= ctrl.GetInt("build_score")
			user, err := myOrm.GetUsersByBuildScore(buildScore)
			if err != nil {
				fmt.Errorf("GetUser fail, %v", err)
			}
			ctrl.Ctx.Output.JSON(user, true, true)
			break
		}
	case 2:
		{
			gameScore, _:= ctrl.GetInt("game_score")
			user, err := myOrm.GetUsersByGameScore(gameScore)
			if err != nil {
				fmt.Errorf("GetUser fail, %v", err)
			}
			ctrl.Ctx.Output.JSON(user, true, true)
			break
		}
	case 3:
		{
			totalScore, _:= ctrl.GetInt("total_score")
			user, err := myOrm.GetUsersByTotalScore(totalScore)
			if err != nil {
				fmt.Errorf("GetUser fail, %v", err)
			}
			ctrl.Ctx.Output.JSON(user, true, true)
			break
		}
	default:
		{
			fmt.Errorf("undefined type for user")
			break
		}
	}
}

func (ctrl *UserController) Post() {
	//ctrl.Ctx.Request()
}

func (ctrl *LevelController) Get() {
	getType, err := ctrl.GetInt("type")
	if err != nil {
		fmt.Errorf("level get type fail, %v", err)
	}
	switch getType {
	case 0:
		{
			makerId, _ := ctrl.GetInt("maker_id")
			level, err := myOrm.GetLevelByMakerId(makerId)
			if err != nil {
				fmt.Errorf("GetLevel fail, %v", err)
			}
			ctrl.Ctx.Output.JSON(level, true, true)
			break
		}
	case 1:
		{
			leverId, _:= ctrl.GetInt("level_id")
			level, err := myOrm.GetLevelsByLevelId(leverId)
			if err != nil {
				fmt.Errorf("GetLevel fail, %v", err)
			}
			ctrl.Ctx.Output.JSON(level, true, true)
			break
		}
	default:
		{
			fmt.Errorf("undefined type for level")
			break
		}
	}
}

func (ctrl *LevelController) Post() {

}

func init() {
	myOrm = dao.MyOrm{}
	myOrm.O = orm.NewOrm()
	myOrm.O.Using("default")
	beego.Router("/user", &UserController{}, "get:Get;post:Post")
	beego.Router("/level", &LevelController{}, "get:Get;post:Post")
}
