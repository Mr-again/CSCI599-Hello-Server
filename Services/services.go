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
	update, _ := ctrl.GetBool("update")
	if update {
		user_id, _ := ctrl.GetInt("user_id")
		money, _ := ctrl.GetInt("money")
		buildScore, _ := ctrl.GetInt("build_score")
		gameScore, _ := ctrl.GetInt("game_score")
		totalScore, _ := ctrl.GetInt("total_score")
		slotNum, _ := ctrl.GetInt("slot_num")
		_, err := myOrm.UpdateUser(user_id, money, buildScore, gameScore, totalScore, slotNum)
		if err != nil {
			fmt.Errorf("update user fail, %v", err)
		}
	} else {
		name := ctrl.GetString("name")
		if name == "" {
			name = "new_user"
		}
		macAddr := ctrl.GetString("mac")
		if macAddr == "" {
			fmt.Errorf("user with name %s should have mac address", name)
		}
		money, _ := ctrl.GetInt("money")
		buildScore, _ := ctrl.GetInt("build_score")
		gameScore, _ := ctrl.GetInt("game_score")
		totalScore, _ := ctrl.GetInt("total_score")
		slotNum, _ := ctrl.GetInt("slot_num")
		_, err := myOrm.AddUser(name, macAddr, money, buildScore, gameScore, totalScore, slotNum)
		if err != nil {
			fmt.Errorf("insert user fail, %v", err)
		}
	}
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
			//fmt.Println(level)
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
	case 2:
		{
			level, err := myOrm.GetAllLevels()
			if err != nil {
				fmt.Errorf("GetLevel fail, %v", err)
			}
			ctrl.Ctx.Output.JSON(level, true, true)
			break
		}
	case 3:
		{
			update, _ := ctrl.GetBool("update")
			delete, _ := ctrl.GetBool("delete")
			if update{
				level_id, _ := ctrl.GetInt("level_id")
				try, _ := ctrl.GetBool("try")
				pass, _ := ctrl.GetBool("pass")
				thumb, _ := ctrl.GetBool("thumb")
				oneStarStep, _ := ctrl.GetInt("one_star_step")
				twoStarStep, _ := ctrl.GetInt("two_star_step")
				threeStarStep, _ := ctrl.GetInt("three_star_step")
				fmt.Print(level_id, try, pass, thumb)
				level, err := myOrm.UpdateLevel(level_id, try, pass, thumb, oneStarStep, twoStarStep, threeStarStep)
				if err != nil {
					fmt.Errorf("insert level fail, %v", err)
				}
				ctrl.Ctx.Output.JSON(level, true, true)
				break
			} else if delete{
				level_id, _ := ctrl.GetInt("level_id")
				level, err := myOrm.DeleteLevel(level_id)
				if err != nil {
					fmt.Errorf("delete level fail, %v", err)
				}
				ctrl.Ctx.Output.JSON(level, true, true)
				break
			} else {
				tryNum, _ := ctrl.GetInt("try_num")
				passNum, _ := ctrl.GetInt("pass_num")
				thumbNum, _ := ctrl.GetInt("thumb_num")
				makerId := ctrl.GetString("id_of_maker")
				mapData := ctrl.GetString("map_data")
				oneStarStep, _ := ctrl.GetInt("one_star_step")
				twoStarStep, _ := ctrl.GetInt("two_star_step")
				threeStarStep, _ := ctrl.GetInt("three_star_step")
				levelName := ctrl.GetString("level_name")
				if mapData == ""{
					mapData = "No Data"
				}
				level, err := myOrm.AddLevel(tryNum, passNum, thumbNum,
					makerId, mapData, oneStarStep, twoStarStep, threeStarStep, levelName)
				if err != nil {
					fmt.Errorf("insert level fail, %v", err)
				}
				ctrl.Ctx.Output.JSON(level, true, true)
				break
			}
		}
	default:
		{
			fmt.Errorf("undefined type for level")
			break
		}
	}
}

func (ctrl *LevelController) Post() {
	update, _ := ctrl.GetBool("update")
	delete, _ := ctrl.GetBool("delete")
	if update{
		level_id, _ := ctrl.GetInt("level_id")
		try, _ := ctrl.GetBool("try")
		pass, _ := ctrl.GetBool("pass")
		thumb, _ := ctrl.GetBool("thumb")
		oneStarStep, _ := ctrl.GetInt("one_star_step")
		twoStarStep, _ := ctrl.GetInt("two_star_step")
		threeStarStep, _ := ctrl.GetInt("three_star_step")
		fmt.Print(level_id, try, pass, thumb)
		_, err := myOrm.UpdateLevel(level_id, try, pass, thumb, oneStarStep, twoStarStep, threeStarStep)
		if err != nil {
			fmt.Errorf("insert level fail, %v", err)
		}
	} else if delete{
		level_id, _ := ctrl.GetInt("level_id")
		_, err := myOrm.DeleteLevel(level_id)
		if err != nil {
			fmt.Errorf("delete level fail, %v", err)
		}
	} else {
		tryNum, _ := ctrl.GetInt("try_num")
		passNum, _ := ctrl.GetInt("pass_num")
		thumbNum, _ := ctrl.GetInt("thumb_num")
		makerId := ctrl.GetString("id_of_maker")
		mapData := ctrl.GetString("map_data")
		oneStarStep, _ := ctrl.GetInt("one_star_step")
		twoStarStep, _ := ctrl.GetInt("two_star_step")
		threeStarStep, _ := ctrl.GetInt("three_star_step")
		level_name := ctrl.GetString("level_name")
		if mapData == ""{
			mapData = "No Data"
		}
		_, err := myOrm.AddLevel(tryNum, passNum, thumbNum, makerId,
			mapData, oneStarStep, twoStarStep, threeStarStep, level_name)
		if err != nil {
			fmt.Errorf("insert level fail, %v", err)
		}
	}
}

func init() {
	myOrm = dao.MyOrm{}
	myOrm.O = orm.NewOrm()
	myOrm.O.Using("default")
	beego.Router("/user", &UserController{}, "get:Get;post:Post")
	beego.Router("/level", &LevelController{}, "get:Get;post:Post")
}
