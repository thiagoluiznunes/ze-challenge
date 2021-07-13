package domain

const (
	Point        string = "Point"
	MultiPolygon string = "MultiPolygon"
)

var MapGeometryType map[string]string = map[string]string{
	Point:        Point,
	MultiPolygon: MultiPolygon,
}
