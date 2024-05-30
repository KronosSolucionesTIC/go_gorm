package utils

import (
	"encoding/json"
	"go_sp/models"
)

func ConvertToJson(contenedorFacturas []models.Factura) string {
	jsonFactura, _ := json.Marshal(contenedorFacturas)
	return string(jsonFactura)
}
