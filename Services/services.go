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
		fmt.Errorf("UserGet GetType fail, %v", err)
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
			break
		}
	case 2:
		{
			break
		}
	case 3:
		{
			break
		}
	default:
		{
			fmt.Errorf("undefined Type")
			break
		}
	}
}

func (ctrl *UserController) Post() {
}

func (ctrl *LevelController) Get() {

}

func (ctrl *LevelController) Post() {

}

func init() {
	myOrm = dao.MyOrm{}
	myOrm.O = orm.NewOrm()
	myOrm.O.Using("default")
	beego.Router("/user", &UserController{}, "get:Get;post:Post")
}
