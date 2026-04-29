package models

// RespuestaAPI representa el formato estándar de respuesta del API.
type RespuestaAPI struct {
	Data    interface{} `json:"Data"`
	Message string      `json:"Message"`
	Status  string      `json:"Status"`
	Success bool        `json:"Success"`
}
