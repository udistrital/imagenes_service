package models

// SolicitudValidacionRostro representa la solicitud cuando la imagen se envía en base64.
type SolicitudValidacionRostro struct {
	ImagenBase64 string `json:"imagen_base64"`
}

// RespuestaValidacionRostro representa el resultado de validación de rostro humano.
type RespuestaValidacionRostro struct {
	Valida          bool    `json:"valida"`
	TieneRostro     bool    `json:"tiene_rostro"`
	EsRostroHumano  bool    `json:"es_rostro_humano"`
	ConfianzaRostro float32 `json:"confianza_rostro"`
	CantidadRostros int     `json:"cantidad_rostros"`
	Mensaje         string  `json:"mensaje"`
}
