package main

import (
	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	_ "my_blog/routers"
)

func main() {
	fmt.Println(123)
	beego.Run()
}
