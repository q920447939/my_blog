package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"my_blog/connection/mysql"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	db := mysql.Db.Unscoped().Table("test")
	fmt.Println(db)
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
