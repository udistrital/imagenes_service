package controllers

import (
	"net/http"

	beego "github.com/beego/beego/v2/server/web"
	"github.com/udistrital/imagenes_service/helpers"
	"github.com/udistrital/imagenes_service/models"
	"github.com/udistrital/imagenes_service/services"
	"github.com/udistrital/imagenes_service/utils_oas/errorctrl"
)

// DetectarRostroController gestiona la validación de rostros humanos mediante Amazon Rekognition.
type DetectarRostroController struct {
	beego.Controller
}

// ValidarRostroHumano ...
// @Title ValidarRostroHumano
// @Description Valida si una imagen contiene un rostro humano detectable. Permite multipart/form-data o base64.
// @Param	imagen	formData	file	false	"Imagen a validar en formato JPG, JPEG o PNG"
// @Param	body	body	models.SolicitudValidacionRostro	false	"Imagen enviada en base64"
// @Success 200 {object} models.RespuestaAPI
// @Failure 400 {object} models.RespuestaAPI
// @Failure 500 {object} models.RespuestaAPI
// @router /rostro/validar [post]
func (c *DetectarRostroController) ValidarRostroHumano() {
	defer errorctrl.ErrorControlController(c.Controller, "DetectarRostroController/ValidarRostroHumano")

	bytesImagen, err := helpers.ObtenerImagenDesdeSolicitud(c.Ctx)
	if err != nil {
		panic(errorctrl.Error("ValidarRostroHumano - helpers.ObtenerImagenDesdeSolicitud", err, "400"))
	}

	respuesta, err := services.ValidarRostroHumano(bytesImagen)
	if err != nil {
		panic(errorctrl.Error("ValidarRostroHumano - services.ValidarRostroHumano", err, "500"))
	}

	c.responder(http.StatusOK, true, "200", "Request successful", respuesta)
}

// responder centraliza el formato estándar de respuesta del API.
func (c *DetectarRostroController) responder(httpStatus int, success bool, status string, message string, data interface{}) {
	c.Ctx.Output.SetStatus(httpStatus)
	c.Data["json"] = models.RespuestaAPI{
		Data:    data,
		Message: message,
		Status:  status,
		Success: success,
	}
	_ = c.ServeJSON()
}
