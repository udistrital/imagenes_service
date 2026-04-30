package main

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/filter/cors"
	_ "github.com/udistrital/imagenes_service/routers"
	"github.com/udistrital/imagenes_service/utils_oas/apistatus"
	"github.com/udistrital/imagenes_service/utils_oas/auditoria"
	"github.com/udistrital/imagenes_service/utils_oas/customerror"
)

func main() {
	allowedOrigins := []string{"*.udistrital.edu.co"}
	if beego.BConfig.RunMode == "dev" {
		allowedOrigins = []string{"*"}
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins: allowedOrigins,
		AllowMethods: []string{"PUT", "PATCH", "GET", "POST", "OPTIONS", "DELETE"},
		AllowHeaders: []string{"Origin", "x-requested-with",
			"content-type",
			"accept",
			"origin",
			"authorization",
			"x-csrftoken"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	beego.ErrorController(&customerror.CustomErrorController{})
	apistatus.Init()
	auditoria.InitMiddleware()
	beego.Run()
}
