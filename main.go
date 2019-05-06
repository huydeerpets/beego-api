package main

import (
	"beego-api/controllers"
	_ "beego-api/routers"
	"beego-api/service"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	beego.InsertFilter("/v1/user/*", beego.BeforeRouter, func(ctx *context.Context) {
		token := ctx.Input.Header("Authorization")
		err, _ := service.CheckToken(token)
		if err != nil {
			respError := &controllers.RespError{
				Code: controllers.CheckError,
				Msg:  err.Error(),
			}
			ctx.Output.JSON(respError, false, false)
		}
	})

	beego.Run()
}
