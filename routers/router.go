package routers

import (
	"github.com/astaxie/beego"
	"my_blog/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/login/long", &controllers.LoginController{}, "POST:Long")
}
