package routers

import (
	"github.com/astaxie/beego"
	ctl "github.com/ikeikeikeike/beego-samples/auth/controllers"
)

func init() {
	beego.Router("/", &ctl.UsersController{}, "get:Index")
	beego.Router("/login", &ctl.LoginController{}, "get,post:Login")
	beego.Router("/logout", &ctl.LoginController{}, "get:Logout")
	beego.Router("/signup", &ctl.LoginController{}, "get,post:Signup")
}
