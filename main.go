package main

import (
	"encoding/json"
	"fmt"
	"go_sp/database"
	"go_sp/models"
	"log"
	"sync"
	"time"
)

func main() {
	db := database.Connect()

	// Consultar las facturas con sus detalles de manera concurrente
	var facturas []models.Factura
	start := time.Now()

	// Consulta concurrente de facturas
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		result := db.Preload("Detalle").Find(&facturas)
		if result.Error != nil {
			log.Fatal("failed to query invoices: ", result.Error)
		}
	}()
	wg.Wait()

	fmt.Println(ConvertToJson(facturas))
	fmt.Println("Consulta de facturas tomó:", time.Since(start))
}

func ConvertToJson(contenedorFacturas []models.Factura) string {
	jsonFactura, _ := json.Marshal(contenedorFacturas)
	return string(jsonFactura)
}
