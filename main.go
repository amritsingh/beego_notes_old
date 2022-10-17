package main

import (
	"beego_notes/models"
	_ "beego_notes/routers"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/session"

	//_ "github.com/astaxie/session/providers/memory"
	_ "github.com/go-sql-driver/mysql"
)

func initSessions() {
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.GlobalSessions, _ = session.NewManager("memory", &session.ManagerConfig{CookieName: "notes",
		Gclifetime: 3600, EnableSetCookie: true, SessionNameInHTTPHeader: "beego_notes",
		Maxlifetime: 36000, Secure: false})
	fmt.Println(beego.GlobalSessions)
	go beego.GlobalSessions.GC()
}

func initDB() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("mysqlDataSource"))
	orm.RegisterModel(new(models.User), new(models.Note))
}

func init() {
	initDB()
	initSessions()
}

func main() {
	beego.Run()
}
