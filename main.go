package main

import (
	_ "api-app/routers"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego"
)

func init() {
	corsHandler := func(ctx *context.Context) {
		beego.Alert("header settings")
		ctx.Output.Header("Access-Control-Allow-Origin", "*")
		ctx.Output.Header("Access-Control-Allow-Methods", "*")
	}
	beego.InsertFilter("*", beego.BeforeRouter, corsHandler)
}


func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
