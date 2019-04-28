package analytic_test

import (
	"testing"

	nanalytic "github.com/haidlir/OverlayTool/analytic"
	nmodel "github.com/haidlir/OverlayTool/model"
)

func TestPointInsidePolygonOK(t *testing.T) {
	// Pre-Defined Polygon
	// 1,1; 1,-1; -1,-1; -1,1
	polygon1 := nmodel.Polygon{
		ID: "polygon1",
		Coordinates: []nmodel.Coordinate {
			nmodel.Coordinate{1.0, 1.0},
			nmodel.Coordinate{1.0, -1.0},
			nmodel.Coordinate{-1.0, -1.0},
			nmodel.Coordinate{-1.0, 1.0},
		},
	}
	polygon2 := nmodel.Polygon{
		ID: "polygon2",
		Coordinates: []nmodel.Coordinate {
			nmodel.Coordinate{2.0, 2.0},
			nmodel.Coordinate{2.0, -2.0},
			nmodel.Coordinate{-2.0, -2.0},
			nmodel.Coordinate{-2.0, 2.0},
		},
	}
	polygon3 := nmodel.Polygon{
		ID: "polygon3",
		Coordinates: []nmodel.Coordinate {
			nmodel.Coordinate{4.0, 1.0},
			nmodel.Coordinate{4.0, -1.0},
			nmodel.Coordinate{3.0, -1.0},
			nmodel.Coordinate{3.0, 1.0},
		},
	}
	polygons := []nmodel.Polygon{polygon1, polygon2, polygon3}
	// Pre-Defined Point
	point1 := nmodel.Point{
		ID: "point1",
		Coordinate: nmodel.Coordinate{0.0, 0.0},
	}
	point2 := nmodel.Point{
		ID: "point2",
		Coordinate: nmodel.Coordinate{1.5, 1.5},
	}
	point3 := nmodel.Point{
		ID: "point3",
		Coordinate: nmodel.Coordinate{-1.5, -1.5},
	}
	point4 := nmodel.Point{
		ID: "point4",
		Coordinate: nmodel.Coordinate{10., 10.},
	}
	points := []nmodel.Point{point1, point2, point3, point4}
	// Iterate
	pointInsidePolygon, err := nanalytic.PointInsidePolygon(polygons, points)
	if err != nil {
		t.Fatalf("find error(s) while interate point inside polygon searching process: %v", err)
	}
	if pointInsidePolygon == nil {
		t.Fatalf("found no result")
	}
	if len(pointInsidePolygon) != 3 {
		t.Errorf("The result should be 3, exactly same to the polygons")
	}
	for _, polygon := range pointInsidePolygon {
		switch polygon.PolygonID {
		case "polygon1":
			if len(polygon.Points) != 1 {
				t.Errorf("polygon1 should has 1 point inside it.")
			}
		case "polygon2":
			if len(polygon.Points) != 3 {
				t.Errorf("polygon2 should has 3 point inside it.")
			}
		case "polygon3":
			if len(polygon.Points) != 0 {
				t.Errorf("polygon3 should has no point inside it.")
			}
		}
	}
}
