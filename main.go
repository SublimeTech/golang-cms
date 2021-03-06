package main

import (
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/session/redis"
	_ "github.com/dionyself/golang-cms/routers"
	"github.com/dionyself/golang-cms/utils"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

func init() {
	currentEnvironment := beego.AppConfig.String("RunMode")
	utils.SessionInit(currentEnvironment)
}

func main() {
	beego.Run()
}
