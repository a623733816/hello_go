package controllers

import (
	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) Get() {
	c.Data["Website"] = "beego1111111111111.me"
	c.Data["Email"] = "astaxie1111111111@gmail.com"
	c.Ctx.WriteString("12323121")
	c.TplName = "index.tpl"
}
