package models

// Factura representa la tabla "facturas" en la base de datos
type Factura struct {
	ID                   uint `gorm:"primaryKey;index"`
	Facturaproforma      int64
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
	Detalle              []Detalle `gorm:"foreignKey:FacturaID"`
}

// Detalle representa la tabla "detalles" en la base de datos
type Detalle struct {
	ID               uint `gorm:"primaryKey"`
	FacturaID        uint `gorm:"index"`
	Facturaproforma  int64
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
