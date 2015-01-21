package controllers

import (
	"html/template"

	"github.com/astaxie/beego"

	"github.com/ikeikeikeike/beego-samples/auth/lib"
	"github.com/ikeikeikeike/beego-samples/auth/models"
)

type LoginController struct {
	BaseController
}

func (c *LoginController) Login() {

	if c.IsLogin {
		c.Ctx.Redirect(302, c.UrlFor("UsersController.Index"))
		return
	}

	c.TplNames = "login/login.tpl"
	c.Data["xsrfdata"] = template.HTML(c.XsrfFormHtml())

	if !c.Ctx.Input.IsPost() {
		return
	}

	flash := beego.NewFlash()
	email := c.GetString("Email")
	password := c.GetString("Password")

	user, err := lib.Authenticate(email, password)
	if err != nil || user.Id < 1 {
		flash.Warning(err.Error())
		flash.Store(&c.Controller)
		return
	}

	flash.Success("Success logged in")
	flash.Store(&c.Controller)

	c.SetLogin(user)

	c.Redirect(c.UrlFor("UsersController.Index"), 303)
}

func (c *LoginController) Logout() {
	c.DelLogin()
	flash := beego.NewFlash()
	flash.Success("Success logged out")
	flash.Store(&c.Controller)

	c.Ctx.Redirect(302, c.UrlFor("LoginController.Login"))
}

func (c *LoginController) Signup() {
	c.TplNames = "login/signup.tpl"
	c.Data["xsrfdata"] = template.HTML(c.XsrfFormHtml())

	if !c.Ctx.Input.IsPost() {
		return
	}

	var err error
	flash := beego.NewFlash()

	u := &models.User{}
	if err = c.ParseForm(u); err != nil {
		flash.Error("Signup invalid!")
		flash.Store(&c.Controller)
		return
	}
	if err = models.IsValid(u); err != nil {
		flash.Error(err.Error())
		flash.Store(&c.Controller)
		return
	}

	id, err := lib.SignupUser(u)
	if err != nil || id < 1 {
		flash.Warning(err.Error())
		flash.Store(&c.Controller)
		return
	}

	flash.Success("Register user")
	flash.Store(&c.Controller)

	c.SetLogin(u)

	c.Redirect(c.UrlFor("UsersController.Index"), 303)
}
