package main

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/filter/cors"
	_ "github.com/udistrital/imagenes_service/routers"
	"github.com/udistrital/imagenes_service/utils_oas/apistatus"
	"github.com/udistrital/imagenes_service/utils_oas/auditoria"
	"github.com/udistrital/imagenes_service/utils_oas/customerror"
	"github.com/udistrital/imagenes_service/utils_oas/security"
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
		AllowMethods: []string{"POST"},
		AllowHeaders: []string{
			"Accept",
			"Authorization",
			"Content-Type",
			"User-Agent",
			"X-Amzn-Trace-Id"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	apistatus.Init()
	auditoria.InitMiddleware()
	beego.ErrorController(&customerror.CustomErrorController{})
	security.SetSecurityHeaders()
	beego.Run()
}
