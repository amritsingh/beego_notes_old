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

func init() {
	beego.InsertFilter("/notes/*", beego.BeforeRouter, AuthFilter)
}

func (c *NotesController) NotesIndex() {
	currUser := currentUser(c.Ctx)
	notes := models.NotesGetAll(currUser)
	c.Data["notes"] = notes
	c.TplName = "notes/index.tpl"
}

func (c *NotesController) NotesNewForm() {
	c.TplName = "notes/new.tpl"
}

func (c *NotesController) NotesCreate() {
	name := c.GetString("name")
	content := c.GetString("content")
	currUser := currentUser(c.Ctx)

	models.NotesCreate(currUser, name, content)
	c.Redirect("/notes", http.StatusMovedPermanently)
}

func (c *NotesController) NotesShow() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	currUser := currentUser(c.Ctx)
	note := models.NotesFind(currUser, id)
	fmt.Println(note)
	c.Data["note"] = note
	c.TplName = "notes/show.tpl"
}

func (c *NotesController) NotesEditForm() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	currUser := currentUser(c.Ctx)
	note := models.NotesFind(currUser, id)
	c.Data["note"] = note
	c.TplName = "notes/edit.tpl"
}

func (c *NotesController) NotesUpdate() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	name := c.GetString("name")
	content := c.GetString("content")
	currUser := currentUser(c.Ctx)
	note := models.NotesFind(currUser, id)
	note.Update(name, content)
	c.Redirect("/notes", http.StatusMovedPermanently)
}

func (c *NotesController) NotesDelete() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	currUser := currentUser(c.Ctx)
	note := models.NotesFind(currUser, id)
	note.MarkDelete()
	c.Redirect("/notes", http.StatusMovedPermanently)
}
