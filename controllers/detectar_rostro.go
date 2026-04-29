package controllers

import (
	"net/http"

	beego "github.com/beego/beego/v2/server/web"
	"github.com/udistrital/imagenes_mid/helpers"
	"github.com/udistrital/imagenes_mid/models"
	"github.com/udistrital/imagenes_mid/services"
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
	bytesImagen, err := helpers.ObtenerImagenDesdeSolicitud(c.Ctx)
	if err != nil {
		c.responder(http.StatusBadRequest, false, "400", err.Error(), nil)
		return
	}

	respuesta, err := services.ValidarRostroHumano(bytesImagen)
	if err != nil {
		c.responder(http.StatusInternalServerError, false, "500", err.Error(), nil)
		return
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
