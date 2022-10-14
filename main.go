package main

import (
	"beego_notes/models"
	_ "beego_notes/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("mysqlDataSource"))
	orm.RegisterModel(new(models.Note))
}

func main() {
	beego.Run()
}
