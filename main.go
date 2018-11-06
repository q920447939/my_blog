package main

import (
		"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	_ "my_blog/routers"
)

func main() {
	//beego.AutoRender
	beego.Run()
}
