package model

// Polygon is the area inside a set of points
type Polygon struct {
	ID          string
	Coordinates []Coordinate
}

// Point is represented by geo coordinate
type Point struct {
	ID string
	Coordinate
}

// Coordinate is combination of latitude and longitude
type Coordinate struct {
	Lat  float64
	Long float64
}

// OverlayPointInsidePolygon is the result structure of Point Inside Polygon script
// Content: ID of Polygon and list of points inside the polygon.
type OverlayPointInsidePolygon struct {
	PolygonID string
	Points    []Point
}
