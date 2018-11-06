package controllers

import (
		"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["isIndex"] = true
	c.Data["isLogin"] = checkAccount(c.Ctx)
	c.TplName = "index.html"
}




func checkAccount(ctx *context.Context) bool {
	ck, err := ctx.Request.Cookie("username")
	if err != nil {
		return false
	}

	uname := ck.Value

	ck, err = ctx.Request.Cookie("password")
	if err != nil {
		return false
	}

	pwd := ck.Value
	return uname == beego.AppConfig.String("username") &&
		pwd == beego.AppConfig.String("password")
}


