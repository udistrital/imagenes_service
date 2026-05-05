package helpers

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"
	"mime/multipart"
	"strings"

	"github.com/beego/beego/v2/server/web/context"
	"github.com/udistrital/imagenes_service/models"
)

var (
	ErrImagenNoEnviada          = errors.New("debe enviar una imagen en el campo 'imagen' o 'imagen_base64'")
	ErrImagenBase64Vacia        = errors.New("el campo 'imagen_base64' no puede estar vacío")
	ErrImagenBase64Invalida     = errors.New("la imagen enviada en base64 no es válida")
	ErrImagenVacia              = errors.New("la imagen enviada está vacía")
	ErrFormatoImagenNoPermitido = errors.New("formato no permitido, use JPG, JPEG o PNG")
)

// ObtenerImagenDesdeSolicitud obtiene la imagen desde multipart/form-data o desde body JSON en base64.
func ObtenerImagenDesdeSolicitud(ctx *context.Context) ([]byte, error) {
	bytesImagen, err := obtenerImagenMultipart(ctx)
	if err == nil && len(bytesImagen) > 0 {
		return bytesImagen, nil
	}

	return obtenerImagenBase64(ctx)
}

// obtenerImagenMultipart obtiene la imagen enviada en el campo multipart "imagen".
func obtenerImagenMultipart(ctx *context.Context) ([]byte, error) {
	archivo, encabezado, err := ctx.Request.FormFile("imagen")
	if err != nil {
		return nil, err
	}
	defer archivo.Close()

	if !esTipoImagenPermitido(encabezado) {
		return nil, ErrFormatoImagenNoPermitido
	}

	bytesImagen, err := io.ReadAll(archivo)
	if err != nil {
		return nil, err
	}

	if len(bytesImagen) == 0 {
		return nil, ErrImagenVacia
	}

	return bytesImagen, nil
}

// obtenerImagenBase64 obtiene la imagen desde un body JSON con el campo imagen_base64.
func obtenerImagenBase64(ctx *context.Context) ([]byte, error) {
	var solicitud models.SolicitudValidacionRostro

	if ctx.Request.Body == nil {
		return nil, ErrImagenNoEnviada
	}

	if err := json.NewDecoder(ctx.Request.Body).Decode(&solicitud); err != nil {
		return nil, ErrImagenNoEnviada
	}

	if strings.TrimSpace(solicitud.ImagenBase64) == "" {
		return nil, ErrImagenNoEnviada
	}

	return DecodificarImagenBase64(solicitud.ImagenBase64)
}

// DecodificarImagenBase64 convierte una imagen base64 en bytes.
func DecodificarImagenBase64(imagenBase64 string) ([]byte, error) {
	if strings.TrimSpace(imagenBase64) == "" {
		return nil, ErrImagenBase64Vacia
	}

	imagenBase64 = limpiarPrefijoBase64(imagenBase64)

	bytesImagen, err := base64.StdEncoding.DecodeString(imagenBase64)
	if err != nil {
		return nil, ErrImagenBase64Invalida
	}

	if len(bytesImagen) == 0 {
		return nil, ErrImagenVacia
	}

	return bytesImagen, nil
}

// limpiarPrefijoBase64 permite recibir base64 puro o con prefijo data URI.
func limpiarPrefijoBase64(valor string) string {
	valor = strings.TrimSpace(valor)

	if strings.Contains(valor, ",") {
		partes := strings.SplitN(valor, ",", 2)
		return partes[1]
	}

	return valor
}

// esTipoImagenPermitido valida el tipo MIME de la imagen recibida por multipart.
func esTipoImagenPermitido(encabezado *multipart.FileHeader) bool {
	tipoContenido := strings.ToLower(encabezado.Header.Get("Content-Type"))

	return tipoContenido == "image/jpeg" ||
		tipoContenido == "image/jpg" ||
		tipoContenido == "image/png"
}
