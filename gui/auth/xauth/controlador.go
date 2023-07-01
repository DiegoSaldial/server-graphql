package xauth

import "github.com/beego/beego/v2/server/web"

type AuthController struct {
	web.Controller
}

func (a *AuthController) Get() {
	a.TplName = "auth/index.tpl"
	a.Data["SomeVar"] = "SomeValue"
	a.Data["Title"] = "Add"
}
