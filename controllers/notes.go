package controllers

import (
	"beego_notes/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/astaxie/beego"
)

type NotesController struct {
	BaseController
}

func init() {
	beego.InsertFilter("/notes*", beego.BeforeRouter, AuthFilter)
}

func (c *NotesController) NotesIndex() {
	currUser := c.currentUser()
	fmt.Println(currUser)
	notes := models.NotesGetAll(currUser)
	c.Data["email"] = currUser.Username
	c.Data["notes"] = notes
	c.TplName = "notes/index.tpl"
}

func (c *NotesController) NotesNewForm() {
	currUser := c.currentUser()
	c.Data["email"] = currUser.Username
	c.TplName = "notes/new.tpl"
}

func (c *NotesController) NotesCreate() {
	name := c.GetString("name")
	content := c.GetString("content")
	currUser := c.currentUser()

	models.NotesCreate(currUser, name, content)
	c.Redirect("/notes", http.StatusMovedPermanently)
}

func (c *NotesController) NotesShow() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	currUser := c.currentUser()
	note := models.NotesFind(currUser, id)
	fmt.Println(note)
	c.Data["email"] = currUser.Username
	c.Data["note"] = note
	c.TplName = "notes/show.tpl"
}

func (c *NotesController) NotesEditForm() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	currUser := c.currentUser()
	note := models.NotesFind(currUser, id)
	c.Data["email"] = currUser.Username
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
	currUser := c.currentUser()
	note := models.NotesFind(currUser, id)
	c.Data["email"] = currUser.Username
	note.Update(name, content)
	c.Redirect("/notes", http.StatusMovedPermanently)
}

func (c *NotesController) NotesDelete() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	currUser := c.currentUser()
	note := models.NotesFind(currUser, id)
	note.MarkDelete()
	c.Redirect("/notes", http.StatusMovedPermanently)
}
