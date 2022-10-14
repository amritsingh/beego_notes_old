package controllers

import (
	"beego_notes/models"
	"fmt"
	"net/http"

	"github.com/astaxie/beego"
)

type NotesController struct {
	beego.Controller
}

func (c *NotesController) NotesIndex() {
	notes := models.NotesGetAll()
	fmt.Println(notes)
	c.Data["notes"] = notes
	c.TplName = "notes/index.tpl"
}

func (c *NotesController) NotesNewForm() {
	c.TplName = "notes/new.tpl"
}

func (c *NotesController) NotesCreate() {
	name := c.GetString("name")
	content := c.GetString("content")

	models.NotesCreate(name, content)
	c.Redirect("/notes", http.StatusMovedPermanently)
}
