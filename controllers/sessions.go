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
		c.Redirect("/signup", http.StatusMovedPermanently)
		return
	}
	if password != confirmPassword {
		c.Redirect("/signup", http.StatusMovedPermanently)
		return
	}
	user := models.UserCreate(email, password)
	if user.ID == 0 {
		c.Redirect("/signup", http.StatusMovedPermanently)
	} else {
		// Set session
		fmt.Println(user)
		//session := beego.GlobalSessions.SessionStart()
		//session := c.StartSession()``
		c.Session.Set("user_id", user.ID)
		// Redirect to home page
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
		c.Redirect("/login", http.StatusMovedPermanently)
	}
}

func (c *SessionsController) Logout() {
	session := c.StartSession()
	session.Delete("user_id")
	c.Redirect("/", http.StatusMovedPermanently)
}
