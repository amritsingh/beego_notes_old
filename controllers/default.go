package controllers

import (
	"beego_notes/models"

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
	currUser := currentUser(c.Ctx)
	if currUser != nil {
		c.Data["email"] = currUser.Username
	}
	c.TplName = "index.tpl"
}

var AuthFilter = func(ctx *context.Context) {
	// session, _ := beego.GlobalSessions.SessionStart(ctx.ResponseWriter, ctx.Request)
	// defer session.SessionRelease(ctx.ResponseWriter)
	session := ctx.Input.CruSession

	sessionID := session.Get("user_id")
	var user *models.User
	if sessionID == nil {
		ctx.Abort(401, "Not authorized")
	} else {
		user = models.UserFind(sessionID.(uint64))
		if user.ID == 0 {
			ctx.Abort(401, "Not authorized")
		}
	}

	var idInterface interface{} = &(user.ID)
	session.Set("user_id", idInterface)
	session.Set("email", user.Username)
}

func currentUser(ctx *context.Context) *models.User {
	session := ctx.Input.CruSession
	sessionID := session.Get("user_id")
	var user *models.User
	if sessionID != nil {
		user = models.UserFind(sessionID.(uint64))
	} else {
		user = nil
	}
	return user
}
