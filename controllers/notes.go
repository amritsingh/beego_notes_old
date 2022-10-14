package controllers

import (
	"beego_notes/models"

	"github.com/astaxie/beego"
)

type NotesController struct {
	beego.Controller
}

func (c *NotesController) NotesIndex() {
	notes := models.NotesGetAll()
	c.Data["notes"] = notes
	c.TplName = "notes/index.tpl"
}
