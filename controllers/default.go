package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "balatero.ddns.net"
	c.Data["Email"] = "junjiebalatero@gmail.com"
	c.TplName = "index.tpl"
}
