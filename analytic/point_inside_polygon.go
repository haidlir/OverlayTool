package analytic

import (
	"fmt"

	pb "gopkg.in/cheggaaa/pb.v1"
	gogeo "github.com/kellydunn/golang-geo"

	nmodel "github.com/haidlir/OverlayTool/model"
)

// PointInsidePolygon iterates the searching process of points inside polygons
func PointInsidePolygon(polygons []nmodel.Polygon, points []nmodel.Point) ([]nmodel.OverlayPointInsidePolygon, error) {
	if len(polygons) <= 0 || len(points) <= 0 {
		return nil, fmt.Errorf("polygons (%v) or points (%v) are not available", len(polygons), len(points))
	}
	geoPolygons := formGeoPolygons(polygons)
	geoPoints := formGeoPoints(points)
	overlayPointInsidePolygon := []nmodel.OverlayPointInsidePolygon{}
	for _, polygon := range polygons {
		newPolyResult := nmodel.OverlayPointInsidePolygon{
			PolygonID : polygon.ID,
			Points : []nmodel.Point{},
		}
		overlayPointInsidePolygon = append(overlayPointInsidePolygon, newPolyResult)
	}
	if len(geoPolygons) != len(overlayPointInsidePolygon) {
		return nil, fmt.Errorf("the length of geoPolygons is different to polygons")
	}
	bar := pb.StartNew(len(geoPoints))
	for iPoint, geoPoint := range geoPoints {
		for iPoly, geoPolygon := range geoPolygons {
			if !geoPolygon.Contains(geoPoint) {
				continue
			}
			overlayPointInsidePolygon[iPoly].Points = append(overlayPointInsidePolygon[iPoly].Points, points[iPoint])
		}
		bar.Increment()
	}
	bar.FinishPrint("Iteration Complete.")
	return overlayPointInsidePolygon, nil
}

func formGeoPolygons(polygons []nmodel.Polygon) []*gogeo.Polygon {
	geoPolygons := []*gogeo.Polygon{}
	for _, polygon := range(polygons) {
		polygonsBoundaries := []*gogeo.Point{}
		for _, coordinate := range polygon.Coordinates {
			polygonsBoundaries = append(polygonsBoundaries, gogeo.NewPoint(coordinate.Lat, coordinate.Long))
		}
		geoPolygons = append(geoPolygons, gogeo.NewPolygon(polygonsBoundaries))
		
	}
	return geoPolygons
}

func formGeoPoints(points []nmodel.Point) []*gogeo.Point {
	geoPoints := []*gogeo.Point{}
	for _, point := range points {
		geoPoints = append(geoPoints, gogeo.NewPoint(point.Lat, point.Long))
	}
	return geoPoints
}
