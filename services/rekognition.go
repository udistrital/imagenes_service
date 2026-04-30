package services

import (
	"context"
	"errors"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/rekognition"
	"github.com/aws/aws-sdk-go-v2/service/rekognition/types"
	"github.com/udistrital/imagenes_service/helpers"
	"github.com/udistrital/imagenes_service/models"
)

// ValidarRostroHumano consulta Amazon Rekognition y determina si la imagen contiene un rostro humano.
func ValidarRostroHumano(bytesImagen []byte) (*models.RespuestaValidacionRostro, error) {
	if len(bytesImagen) == 0 {
		return nil, helpers.ErrImagenVacia
	}

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return nil, errors.New("error cargando configuración de AWS: " + err.Error())
	}

	cliente := rekognition.NewFromConfig(cfg)

	resultado, err := cliente.DetectFaces(context.TODO(), &rekognition.DetectFacesInput{
		Image: &types.Image{
			Bytes: bytesImagen,
		},
		Attributes: []types.Attribute{
			types.AttributeDefault,
		},
	})
	if err != nil {
		return nil, errors.New("error consultando Amazon Rekognition: " + err.Error())
	}

	respuesta := &models.RespuestaValidacionRostro{
		Valida:          false,
		TieneRostro:     false,
		EsRostroHumano:  false,
		ConfianzaRostro: 0,
		CantidadRostros: len(resultado.FaceDetails),
		Mensaje:         "No se detectó un rostro humano en la imagen",
	}

	if len(resultado.FaceDetails) == 0 {
		return respuesta, nil
	}

	rostro := resultado.FaceDetails[0]
	respuesta.TieneRostro = true

	if rostro.Confidence != nil {
		respuesta.ConfianzaRostro = *rostro.Confidence
	}

	if respuesta.ConfianzaRostro >= 90 {
		respuesta.Valida = true
		respuesta.EsRostroHumano = true
		respuesta.Mensaje = "La imagen contiene un rostro humano detectable"
		return respuesta, nil
	}

	respuesta.Mensaje = "Se detectó un rostro, pero con baja confianza"
	return respuesta, nil
}
