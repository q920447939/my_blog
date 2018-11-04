package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
	this.TplName = "login.html"
}

func (this *LoginController) Post() {
	username := this.Input().Get("username")
	password := this.Input().Get("password")
	if username == "" || password == "" {
		return
	}
	cUsername := beego.AppConfig.String("username")
	cPassword := beego.AppConfig.String("password")
	if username != cUsername || password != cPassword {
		return
	}
	//设置cookie
	this.Ctx.SetCookie("username", username, fmt.Sprint(beego.AppConfig.String("COOKIE::MAX_TIME")), "/")
	this.Ctx.SetCookie("password", password, fmt.Sprint(beego.AppConfig.String("COOKIE::MAX_TIME")), "/")
	this.Redirect("/", 302)
}

func (this *LoginController) Long() {
	fmt.Println(123)
}
