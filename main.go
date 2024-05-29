package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Factura representa la tabla "facturas" en la base de datos
type Factura struct {
	ID                   uint  `gorm:"primaryKey"`
	Facturaproforma      int64 `gorm:"unique"`
	Fechafacturaproforma string
	Guiamaster           string
	Guiahija             string
	Tipoguia             string
	ContenedorNo         string
	NoBl                 string
	TerminosNegoc        string
	Idclienteexportaddor string
	Idclienteimportador  string
	Agencia              string
	Aerolineaonaviera    string
	PesoNeto             float64
	PesoBruto            float64
	Ggn                  string
	IdGuia               int64
	TotalRemisiones      int64
	Detalle              []Detalle `gorm:"foreignKey:Facturaproforma;references:Facturaproforma"`
}

// Detalle representa la tabla "detalles" en la base de datos
type Detalle struct {
	ID               uint  `gorm:"primaryKey"`
	Facturaproforma  int64 `gorm:"index"`
	ItemFactura      int64
	Po               string
	IdproducoMaestro string
	Nopallet         string
	CajasXPallet     float64
	ClamshellXCajas  float64
	GramosXClamshell float64
	Tipocaja         string
	TipoTapacaja     string
	PrecioKilo       float64
}

func main() {
	// Datos de conexión a la base de datos
	dsn := "root:root@tcp(localhost:3306)/ventas?charset=utf8mb4&parseTime=True&loc=Local"

	// Abrir conexión a la base de datos
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database: ", err)
	}

	// Migrar el esquema
	db.AutoMigrate(&Factura{}, &Detalle{})

	// Consultar las facturas con sus detalles
	var facturas []Factura
	result := db.Preload("Detalle").Find(&facturas)
	if result.Error != nil {
		log.Fatal("failed to query invoices: ", result.Error)
	}

	// Imprimir las facturas con sus detalles
	for _, factura := range facturas {
		fmt.Printf("Factura ID: %d, Facturaproforma: %d\n", factura.ID, factura.Facturaproforma)
		for _, detalle := range factura.Detalle {
			fmt.Printf("  Detalle ID: %d, ItemFactura: %d, Producto: %s\n", detalle.ID, detalle.ItemFactura, detalle.IdproducoMaestro)
		}
	}

	// registerDummyData(db)
}

func registerDummyData(db *gorm.DB) error { //Registra data tipo dummy
	rand.Seed(time.Now().UnixNano())

	for i := 102; i < 5000; i++ {
		factura := Factura{
			Facturaproforma:      int64(i + 1),
			Fechafacturaproforma: time.Now().Format("2006-01-02"),
			Guiamaster:           fmt.Sprintf("GM%d", i+1),
			Guiahija:             fmt.Sprintf("GH%d", i+1),
			Tipoguia:             "TIPO1",
			ContenedorNo:         fmt.Sprintf("CONT%d", i+1),
			NoBl:                 fmt.Sprintf("BL%d", i+1),
			TerminosNegoc:        "FOB",
			Idclienteexportaddor: fmt.Sprintf("EX%d", i+1),
			Idclienteimportador:  fmt.Sprintf("IM%d", i+1),
			Agencia:              fmt.Sprintf("AG%d", i+1),
			Aerolineaonaviera:    fmt.Sprintf("AIR%d", i+1),
			PesoNeto:             rand.Float64() * 100,
			PesoBruto:            rand.Float64() * 100,
			Ggn:                  fmt.Sprintf("GGN%d", i+1),
			IdGuia:               int64(rand.Intn(100)),
			TotalRemisiones:      int64(rand.Intn(100)),
		}

		for j := 0; j < 20; j++ { // 5 detalles por factura
			detalle := Detalle{
				Facturaproforma:  int64(i + 1),
				ItemFactura:      int64(j + 1),
				Po:               fmt.Sprintf("PO%d", j+1),
				IdproducoMaestro: fmt.Sprintf("PRD%d", j+1),
				Nopallet:         fmt.Sprintf("PAL%d", j+1),
				CajasXPallet:     rand.Float64() * 10,
				ClamshellXCajas:  rand.Float64() * 100,
				GramosXClamshell: rand.Float64() * 500,
				Tipocaja:         "TIPO1",
				TipoTapacaja:     "TIPO2",
				PrecioKilo:       rand.Float64() * 10,
			}
			factura.Detalle = append(factura.Detalle, detalle)
		}

		if err := db.Create(&factura).Error; err != nil {
			return err
		}
	}

	return nil
}
