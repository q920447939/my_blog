package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type LoginController struct {
	BaseResultController
}

func (this *LoginController) Get() {
	if this.Input().Get("exit")  =="true" {
		this.TplName = "index.html"
		//设置cookie
		this.Ctx.SetCookie("username", "", fmt.Sprint(beego.AppConfig.String("COOKIE::MAX_TIME")), "/")
		this.Ctx.SetCookie("password", "", fmt.Sprint(beego.AppConfig.String("COOKIE::EXPIRE_TIME")), "/")
		this.Data["isIndex"] = true

		return
	}
	this.TplName = "login.html"
}

func (this *LoginController) Post() {
	this.TplName = "demo.html"
	username := this.Input().Get("username")
	password := this.Input().Get("password")
	if username == "" || password == "" {
		this.JsonResult(4000,"账号或密码为空")
		return
	}
	cUsername := beego.AppConfig.String("username")
	cPassword := beego.AppConfig.String("password")
	if username != cUsername || password != cPassword {
		this.JsonResult(4000,"账号或密码错误")
		return
	}
	//设置cookie
	this.Ctx.SetCookie("username", username, fmt.Sprint(beego.AppConfig.String("COOKIE::MAX_TIME")), "/")
	this.Ctx.SetCookie("password", password, fmt.Sprint(beego.AppConfig.String("COOKIE::MAX_TIME")), "/")
	this.Redirect("/", 302)
	return
}


