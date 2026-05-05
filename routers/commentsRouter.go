package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/udistrital/imagenes_service/controllers:DetectarRostroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/imagenes_service/controllers:DetectarRostroController"],
        beego.ControllerComments{
            Method: "ValidarRostroHumano",
            Router: `/validar`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
