package controllers

import (
	"beego_notes/models"
	"fmt"
	"net/http"
)

type SessionsController struct {
	BaseController
}

func (c *SessionsController) SignupPage() {
	c.TplName = "sessions/signup.tpl"
}

func (c *SessionsController) LoginPage() {
	c.TplName = "sessions/login.tpl"
}

func (c *SessionsController) Signup() {
	email := c.GetString("email")
	password := c.GetString("password")
	confirmPassword := c.GetString("confirm_password")
	// Check if email already exists
	available := models.UserCheckAvailability(email)
	fmt.Println(email, password, confirmPassword, available)
	if !available {
		flashMessage(&c.Controller, Notice, "User already exists!")
		c.TplName = "sessions/signup.tpl"
		return
	}
	if password != confirmPassword {
		fmt.Println("Password mismatch!")
		flashMessage(&c.Controller, Notice, "Password and Confirm Password mismatch!")
		//flash.Store(&c.Controller)
		//c.Data["alert"] = "Password and Confirm Password mismatch!"
		c.TplName = "sessions/signup.tpl"
		return
	}
	user := models.UserCreate(email, password)
	if user.ID == 0 {
		flashMessage(&c.Controller, Error, "Error!")
		c.TplName = "sessions/signup.tpl"
	} else {
		// Set session
		c.Session.Set("user_id", user.ID)
		// Redirect to home page
		flashMessage(&c.Controller, Success, "User created!")
		c.Redirect("/", http.StatusMovedPermanently)
	}
}

func (c *SessionsController) Login() {
	email := c.GetString("email")
	password := c.GetString("password")
	user := models.UserCheck(email, password)
	if user != nil && user.ID > 0 {
		// Set session
		c.Session.Set("user_id", user.ID)
		// Redirect to home page
		c.Redirect("/", http.StatusMovedPermanently)
	} else {
		flashMessage(&c.Controller, Error, "User and/or password wrong!")
		c.TplName = "sessions/login.tpl"
		return
	}
}

func (c *SessionsController) Logout() {
	session := c.StartSession()
	session.Delete("user_id")
	c.Redirect("/", http.StatusMovedPermanently)
}
