package controllers

import (
	"beego_notes/models"
	"net/http"

	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/session"

	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
	Session session.Store
}

func (c *BaseController) Prepare() {
	c.Session = c.StartSession()
}

type MainController struct {
	BaseController
}

func (c *MainController) Get() {
	currUser := c.currentUser()
	if currUser != nil {
		c.Data["email"] = currUser.Username
	}
	c.TplName = "index.tpl"
}

var AuthFilter = func(ctx *context.Context) {
	session := ctx.Input.CruSession

	sessionID := session.Get("user_id")
	var user *models.User
	if sessionID == nil {
		//ctx.Abort(401, "Not authorized")
		ctx.Redirect(http.StatusPermanentRedirect, "/login")
	} else {
		user = models.UserFind(sessionID.(uint64))
		if user.ID == 0 {
			// ctx.Abort(401, "Not authorized")
			ctx.Redirect(http.StatusPermanentRedirect, "/login")
		}
	}
	session.Set("email", user.Username)
}

func (contr *BaseController) currentUser() *models.User {
	sessionID := contr.Session.Get("user_id")
	var user *models.User
	if sessionID != nil {
		userID, _ := sessionID.(uint64)
		user = models.UserFind(userID)
	} else {
		user = nil
	}
	return user
}

type FlashMessageType int

const (
	Error FlashMessageType = iota
	Warning
	Success
	Notice
)

func flashMessage(contr *beego.Controller, messageType FlashMessageType, message string) *beego.FlashData {
	flash := beego.NewFlash()
	switch messageType {
	case Error:
		flash.Error(message)
	case Warning:
		flash.Warning(message)
	case Success:
		flash.Success(message)
	case Notice:
		flash.Notice(message)
	}
	flash.Store(contr)
	return flash
}
