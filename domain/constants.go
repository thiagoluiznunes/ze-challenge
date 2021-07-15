package domain

// NRTransactionKey indicates the key used by NewRelic when executing a new transaction
const NRTransactionKey = "NRTransaction"

const (
	Point        string = "Point"
	MultiPolygon string = "MultiPolygon"
)

var MapGeometryType map[string]string = map[string]string{
	Point:        Point,
	MultiPolygon: MultiPolygon,
}
