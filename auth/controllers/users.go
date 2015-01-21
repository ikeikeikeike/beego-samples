package controllers

import "github.com/astaxie/beego"

type UsersController struct {
	BaseController
}

func (c *UsersController) NestPrepare() {
	if !c.IsLogin {
		c.Ctx.Redirect(302, c.LoginPath())
		return
	}
}

// func (c *UsersController) NestFinish() {}

func (c *UsersController) Index() {
	beego.ReadFromRequest(&c.Controller)

	c.TplNames = "users/index.tpl"
}
