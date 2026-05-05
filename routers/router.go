// Package routers contiene la configuración de rutas del servicio de imágenes.
//
// @APIVersion 1.0.0
// @Title Imágenes Service API
// @Description API para validación de imágenes mediante Amazon Rekognition.
// @TermsOfServiceUrl http://swagger.io/terms/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	beego "github.com/beego/beego/v2/server/web"

	"github.com/udistrital/imagenes_service/controllers"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/rostro",
			beego.NSInclude(
				&controllers.DetectarRostroController{},
			),
		),
	)

	beego.AddNamespace(ns)
}
