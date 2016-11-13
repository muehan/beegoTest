package controllers

import (
	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
}

func (c *HomeController) Get() {
	c.Data["name"] = "Andre"
	c.Data["list"] = []string{"a", "b", "c"}
	c.TplName = "home.html"
}
