package gui

import (
	"opentaxi/gui/auth/xauth"

	"github.com/beego/beego/v2/server/web"
)

func MainBeego() {
	web.BConfig.CopyRequestBody = true
	web.BConfig.Listen.EnableAdmin = true
	web.BConfig.Listen.AdminAddr = "localhost"
	web.BConfig.Listen.AdminPort = 8015

	web.Router("/", &xauth.AuthController{})
	web.Run()
}
