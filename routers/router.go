package routers

import (
	"beego_notes/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	beego.Router("/notes", &controllers.NotesController{}, "get:NotesIndex")
}
