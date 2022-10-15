package routers

import (
	"beego_notes/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	beego.Router("/notes", &controllers.NotesController{}, "get:NotesIndex")
	beego.Router("/notes/new", &controllers.NotesController{}, "get:NotesNewForm")
	beego.Router("/notes", &controllers.NotesController{}, "post:NotesCreate")
	beego.Router("/notes/:id([0-9]+)", &controllers.NotesController{}, "get:NotesShow")
	beego.Router("/notes/edit/:id", &controllers.NotesController{}, "get:NotesEditForm")
	beego.Router("/notes/:id([0-9]+)", &controllers.NotesController{}, "post:NotesUpdate")
	beego.Router("/notes/:id", &controllers.NotesController{}, "delete:NotesDelete")

	beego.Router("/signup", &controllers.SessionsController{}, "get:SignupPage")
	beego.Router("/login", &controllers.SessionsController{}, "get:LoginPage")
}
