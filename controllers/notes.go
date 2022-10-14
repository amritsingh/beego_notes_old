package controllers

import (
	"beego_notes/models"
	"fmt"
	"net/http"
	"strconv"

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

func (c *NotesController) NotesShow() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	note := models.NotesFind(id)
	fmt.Println(note)
	c.Data["note"] = note
	c.TplName = "notes/show.tpl"
}
