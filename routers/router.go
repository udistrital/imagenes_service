package routers

import (
	beego "github.com/beego/beego/v2/server/web"

	"github.com/udistrital/imagenes_service/controllers"
)

func init() {
	beego.Router("/v1/rostro/validar", &controllers.DetectarRostroController{}, "post:ValidarRostroHumano")
}
