package Services

import (
	"../dao"
	"bufio"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"os"
	"strconv"
	"unsafe"
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
		_ = fmt.Errorf("user get type fail, %v", err)
	}
	switch getType {
	case 0:
		users, err := myOrm.GetAllUsers()
		if err != nil {
			_ = fmt.Errorf("get all users failed with error, %v", err)
		}
		_ = ctrl.Ctx.Output.JSON(users, true, true)
		break
	case 1:
		users, err := myOrm.GetAllUsersByMoney()
		if err != nil {
			_ = fmt.Errorf("get all users by order of money failed with error, %v", err)
		}
		_ = ctrl.Ctx.Output.JSON(users, true, true)
	case 2:
		users, err := myOrm.GetAllUsersByBuildScore()
		if err != nil {
			_ = fmt.Errorf("get all users by order of build score failed with error, %v", err)
		}
		_ = ctrl.Ctx.Output.JSON(users, true, true)
		break
	case 3:
		users, err := myOrm.GetAllUsersByGameScore()
		if err != nil {
			_ = fmt.Errorf("get all users by order of game score failed with error, %v", err)
		}
		_ = ctrl.Ctx.Output.JSON(users, true, true)
		break
	case 4:
		users, err := myOrm.GetAllUsersByTotalScore()
		if err != nil {
			_ = fmt.Errorf("get all users by order of total score failed with error, %v", err)
		}
		_ = ctrl.Ctx.Output.JSON(users, true, true)
		break
	case 5:
		macAddr := ctrl.GetString("mac")
		user, err := myOrm.GetUserByMacAddr(macAddr)
		if err != nil {
			_ = fmt.Errorf("get user with mac_addr failed with error, %v", err)
		}
		_ = ctrl.Ctx.Output.JSON(user, true, true)
		break
	default:
		_ = fmt.Errorf("undefined type for user get")
		break
	}
}

func (ctrl *UserController) Post() {
	getType, err := ctrl.GetInt("type")
	if err != nil {
		_ = fmt.Errorf("user get type fail, %v", err)
	}
	// type = 0 -> add, type = 1 -> update, type = 2 -> delete
	switch getType {
	case 0:
		name := ctrl.GetString("name")
		if name == "" {
			name = "new_user"
		}
		macAddr := ctrl.GetString("mac")
		if macAddr == "" {
			_ = fmt.Errorf("user with name %s should have mac address", name)
		}
		user, err := myOrm.AddUser(name, macAddr)
		if err != nil{
			_ = fmt.Errorf("add user error, %v", err)
		}
		_ = ctrl.Ctx.Output.JSON(user, true, true)
	case 1:
		userId, _ := ctrl.GetInt("user_id")
		money, _ := ctrl.GetInt("money")
		buildScore, _ := ctrl.GetInt("build_score")
		gameScore, _ := ctrl.GetInt("game_score")
		totalScore, _ := ctrl.GetInt("total_score")
		slotNum, _ := ctrl.GetInt("slot_num")
		user, err := myOrm.UpdateUser(userId, money, buildScore, gameScore, totalScore, slotNum)
		if err != nil {
			_ = fmt.Errorf("update user fail, %v", err)
		}
		_ = ctrl.Ctx.Output.JSON(user, true, true)
	case 2:
		userId, _ := ctrl.GetInt("user_id")
		user, err := myOrm.DeleteUser(userId)
		if err != nil {
			_ = fmt.Errorf("insert user fail, %v", err)
		}
		_ = ctrl.Ctx.Output.JSON(user, true, true)
	default:
		_ = fmt.Errorf("undefined type for user post")
		break
	}
}

func (ctrl *LevelController) Get() {
	getType, err := ctrl.GetInt("type")
	if err != nil {
		_ = fmt.Errorf("level get type fail, %v", err)
	}
	switch getType {
	case 0:
		makerId, _ := ctrl.GetInt("maker_id")
		level, err := myOrm.GetLevelByMakerId(makerId)
		if err != nil {
			_ = fmt.Errorf("GetLevel fail, %v", err)
		}
		_ = ctrl.Ctx.Output.JSON(level, true, true)
		break
	case 1:
		leverId, _:= ctrl.GetInt("level_id")
		level, err := myOrm.GetLevelsByLevelId(leverId)
		if err != nil {
			_ = fmt.Errorf("GetLevel fail, %v", err)
		}
		_ = ctrl.Ctx.Output.JSON(level, true, true)
		break
	case 2:
		level, err := myOrm.GetAllLevels()
		if err != nil {
			_ = fmt.Errorf("GetLevel fail, %v", err)
		}
		_ = ctrl.Ctx.Output.JSON(level, true, true)
		// ctrl.Ctx.Output.Download("Screenshots/Map_Screenshot.png", "sc")
		//li := models.LevelWithImages{level, "wwwww"}
		//_ = ctrl.Ctx.Output.JSON(li, true, true)
		break
	case 3:
		level, err := myOrm.GetAllLevelsByTry()
		if err != nil {
			_ = fmt.Errorf("GetLevel fail, %v", err)
		}
		_ = ctrl.Ctx.Output.JSON(level, true, true)
		break
	case 4:
		level, err := myOrm.GetAllLevelsByPass()
		if err != nil {
			_ = fmt.Errorf("GetLevel fail, %v", err)
		}
		_ = ctrl.Ctx.Output.JSON(level, true, true)
		break
	case 5:
		level, err := myOrm.GetAllLevelsByThumb()
		if err != nil {
			_ = fmt.Errorf("GetLevel fail, %v", err)
		}
		_ = ctrl.Ctx.Output.JSON(level, true, true)
		break
	default:
		{
			_ = fmt.Errorf("undefined type for level get")
			break
		}
	}
}

func (ctrl *LevelController) Post() {
	getType, err := ctrl.GetInt("type")
	if err != nil {
		_ = fmt.Errorf("level post type fail, %v", err)
	}
	// type = 0 -> add, type = 1 -> update, type = 2 -> delete
	switch getType {
	case 0:
		tryNum, _ := ctrl.GetInt("try_num")
		passNum, _ := ctrl.GetInt("pass_num")
		thumbNum, _ := ctrl.GetInt("thumb_num")
		makerId, _ := ctrl.GetInt("id_of_maker")
		mapData := ctrl.GetString("map_data")
		oneStarStep, _ := ctrl.GetInt("one_star_step")
		twoStarStep, _ := ctrl.GetInt("two_star_step")
		threeStarStep, _ := ctrl.GetInt("three_star_step")
		levelName := ctrl.GetString("level_name")
		if mapData == ""{
			mapData = "No Data"
		}
		level, err := myOrm.AddLevel(tryNum, passNum, thumbNum, makerId,
			mapData, oneStarStep, twoStarStep, threeStarStep, levelName)
		if err != nil {
			fmt.Print("insert level fail with ", err)
		}
		_ = ctrl.Ctx.Output.JSON(level, true, true)

		file, _, err := ctrl.GetFile("screenshot")  //返回文件，文件信息头，错误信息
		// fmt.Print(file, information, err)
		if err != nil {
			fmt.Print("File retrieval failure")
		} else{
			defer file.Close()
			_ = ctrl.SaveToFile("screenshot", "Screenshots/Map_" + strconv.Itoa(level.LevelId) + "_Screenshot.png")
		}
		break
	case 1:
		levelId, _ := ctrl.GetInt("level_id")
		try, _ := ctrl.GetBool("try")
		pass, _ := ctrl.GetBool("pass")
		thumb, _ := ctrl.GetBool("thumb")
		oneStarStep, _ := ctrl.GetInt("one_star_step")
		twoStarStep, _ := ctrl.GetInt("two_star_step")
		threeStarStep, _ := ctrl.GetInt("three_star_step")
		fmt.Print(levelId, try, pass, thumb)
		level, err := myOrm.UpdateLevel(levelId, try, pass, thumb, oneStarStep, twoStarStep, threeStarStep)
		if err != nil {
			_ = fmt.Errorf("insert level fail, %v", err)
		}
		_ = ctrl.Ctx.Output.JSON(level, true, true)

		file, _, err := ctrl.GetFile("screenshot")
		if err != nil {
			fmt.Print("File retrieval failure")
		} else{
			defer file.Close()
			_ = ctrl.SaveToFile("screenshot",
				"Screenshots/Map_" + strconv.Itoa(level.LevelId) + "_Screenshot.png")
		}
		break
	case 2:
		levelId, _ := ctrl.GetInt("level_id")
		level, err := myOrm.DeleteLevel(levelId)
		if err != nil {
			_ = fmt.Errorf("delete level fail, %v", err)
		}
		_ = ctrl.Ctx.Output.JSON(level, true, true)
		break
	//case 3:
	//	fmt.Print("type 3")
	//	var user models.User
	//	data := ctrl.Ctx.Input.RequestBody
	//	_ = json.Unmarshal(data, &user)
	//	fmt.Print(user)
	//	_ = ctrl.Ctx.Output.JSON(user, true, true)
	//	break
	case 3:
		// data := ctrl.Ctx.Input.RequestBody
		// str := *(*string)(unsafe.Pointer(&data))
		////_ = ctrl.Ctx.Output.JSON("Ok", true, true)
		//fileStr, _ := RetrieveROM("Screenshot.png")
		//var data = []byte(fileStr)
		//file, _ := os.Create("test.png")
		//defer file.Close()
		//bytes_written, _ := file.Write(data)

		// fmt.Printf("Wrote %d bytes to file \n", bytes_written)
		break
	default:
		fmt.Print("undefined type for level post")
		break
	}
}

func RetrieveROM(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Print(err)
		return "", err
	}
	defer file.Close()
	stats, statsErr := file.Stat()
	if statsErr != nil {
		fmt.Print(statsErr)
		return "", statsErr
	}
	var size = stats.Size()
	bytes := make([]byte, size)
	bufr := bufio.NewReader(file)
	_,err = bufr.Read(bytes)
	var str = *(*string)(unsafe.Pointer(&bytes))
	return str, err
}

func init() {
	myOrm = dao.MyOrm{}
	myOrm.O = orm.NewOrm()
	_ = myOrm.O.Using("default")
	beego.Router("/user", &UserController{}, "get:Get;post:Post")
	beego.Router("/level", &LevelController{}, "get:Get;post:Post")
}
