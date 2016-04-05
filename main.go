package main

import (
	_ "awesome/routers"
	_ "awesome/viewfunc"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
    _ "github.com/go-sql-driver/mysql"
)

func init() {
    orm.RegisterDriver("mysql", orm.DR_MySQL)

    orm.RegisterDataBase("default", "mysql", "root:@/bee?charset=utf8", 30)
}

func main() {
	
	beego.Run()
}

