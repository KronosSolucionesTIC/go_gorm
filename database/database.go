package database

import (
	"fmt"
	"go_sp/models"
	"log"
	"math/rand"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	// Datos de conexión a la base de datos
	dsn := "root:root@tcp(localhost:3306)/ventas?charset=utf8mb4&parseTime=True&loc=Local"

	// Abrir conexión a la base de datos
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database: ", err)
	}

	return db
}

func RegisterDummyData(db *gorm.DB) error { //Registra data tipo dummy
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 5000; i++ {
		factura := models.Factura{
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
			detalle := models.Detalle{
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
