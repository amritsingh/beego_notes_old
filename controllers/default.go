package controllers

import (
	"github.com/astaxie/beego/context"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

var AuthFilter = func(ctx *context.Context) {
	// sessionID := ctx.Input.Session("id").(uint64)
	ctx.Abort(401, "Not authorized")
}
