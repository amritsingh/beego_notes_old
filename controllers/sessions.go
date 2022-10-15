package controllers

import "github.com/astaxie/beego"

type SessionsController struct {
	beego.Controller
}

func (c *SessionsController) SignupPage() {
	c.TplName = "sessions/signup.tpl"
}

func (c *SessionsController) LoginPage() {
	c.TplName = "sessions/login.tpl"
}
