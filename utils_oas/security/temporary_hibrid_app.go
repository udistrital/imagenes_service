package security

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"errors"

	beego "github.com/beego/beego/v2/server/web"
	"github.com/udistrital/imagenes_service/utils_oas/request"
)

// SecurityHivridApp valida el acceso temporal híbrido mediante usuario, token y hash.
// Nota: se conserva el nombre original "Hivrid" para no romper referencias existentes.
func SecurityHivridApp(user, token, hash string) (answer bool, err error) {
	var hashGenerated string

	// Paso 1: generar hash local con usuario, token y secreto configurado.
	if hashGenerated, err = GeneratedHash(user, token); err != nil {
		return false, err
	}

	// Paso 2: validar que el hash generado coincida con el hash recibido.
	if hashGenerated != hash {
		return false, errors.New("the hash does not match")
	}

	// Paso 3: consultar si existe sesión académica asociada al token.
	if sesion := GetSesionAcademica(token); !sesion {
		return false, errors.New("there is no session")
	}

	return true, nil
}

// GeneratedHash genera un hash SHA1 usando usuario, token y Secret configurado en app.conf.
func GeneratedHash(user, token string) (outputHash string, err error) {
	secret, err := beego.AppConfig.String("Secret")
	if err != nil || secret == "" {
		return "", errors.New("secret not defined")
	}

	var buffer bytes.Buffer
	buffer.WriteString(user)
	buffer.WriteString(token)
	buffer.WriteString(secret)

	h := sha1.New()
	h.Write([]byte(buffer.String()))

	outputHash = hex.EncodeToString(h.Sum(nil))
	return outputHash, nil
}

// GetSesionAcademica consulta en WSO2 si el token tiene una sesión académica activa.
func GetSesionAcademica(token string) (sesion bool) {
	var dataSesion interface{}

	url := "http://jbpm.udistritaloas.edu.co:8280/services/uranoPruebasProxy/get_usuario_session/" + token

	if err := request.GetJsonWSO2(url, &dataSesion); err != nil || dataSesion == nil {
		return false
	}

	dataMap, ok := dataSesion.(map[string]interface{})
	if !ok {
		return false
	}

	usuarios, ok := dataMap["usuarios"].(map[string]interface{})
	if !ok {
		return false
	}

	if _, ok := usuarios["usuario"]; ok {
		return true
	}

	return false
}
