package controllers

import (
		"github.com/astaxie/beego"
	"fmt"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
	this.TplName = "login.html"
}

func (this *LoginController)  Post() {
	username := this.Input().Get("username")
	password := this.Input().Get("password")
	if username == "" || password == "" {
		controllers.BaseResultController{}.JsonResult(4000,"账号或密码为空")
		return
	}
	cUsername := beego.AppConfig.String("username")
	cPassword := beego.AppConfig.String("password")
	if username != cUsername || password != cPassword {
		controllers.BaseResultController{}.JsonResult(4000,"账号或密码错误")
		return
	}
	//设置cookie
	this.Ctx.SetCookie("username", username, fmt.Sprint(beego.AppConfig.String("COOKIE::MAX_TIME")), "/")
	this.Ctx.SetCookie("password", password, fmt.Sprint(beego.AppConfig.String("COOKIE::MAX_TIME")), "/")
	this.Redirect("/", 302)
	return
}


