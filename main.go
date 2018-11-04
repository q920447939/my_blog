package main

import (
	_ "github.com/go-sql-driver/mysql"
	_ "my_blog/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}

