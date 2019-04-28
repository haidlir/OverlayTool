package csv_test

import (
	"testing"

	ncsv "github.com/haidlir/OverlayTool/csv"
	nmodel "github.com/haidlir/OverlayTool/model"
)

func TestStoreOverlayPointInsidePolygonOK(t *testing.T) {
	// Prepare dummy points
	point1 := nmodel.Point{
		ID: "point1",
	}
	point2 := nmodel.Point{
		ID: "point2",
	}
	point3 := nmodel.Point{
		ID: "point3",
	}
	// Pre-defined OverlayPointInsidePolygon
	poly1 := nmodel.OverlayPointInsidePolygon{
		PolygonID : "polygon1",
		Points : []nmodel.Point{point1},
	}
	poly2 := nmodel.OverlayPointInsidePolygon{
		PolygonID : "polygon2",
		Points : []nmodel.Point{point1, point2},
	}
	poly3 := nmodel.OverlayPointInsidePolygon{
		PolygonID : "polygon3",
		Points : []nmodel.Point{point1, point2, point3},
	}
	poly4 := nmodel.OverlayPointInsidePolygon{
		PolygonID : "polygon4",
		Points : []nmodel.Point{},
	}
	overlayPointInsidePolygon := []nmodel.OverlayPointInsidePolygon{poly1, poly2, poly3, poly4}
	delimiter := ";"
	outputFilename := "_fixtures/test-output.csv"
	err := ncsv.StoreOverlayPointInsidePolygon(overlayPointInsidePolygon, delimiter, outputFilename)
	if err != nil {
		t.Fatalf("It should be OK: %v", err)
	}
	// let's check it
	records, err := ncsv.ReadCSVFile(outputFilename, delimiter)
	if err != nil {
		t.Fatalf("It should be OK: %v", err)
	}
	if len(records) != 4 {
		t.Fatalf("The number of row of the output file should be 4 instead of %v", len(records))
	}
	for i, record := range records{
		if len(record) != 2 {
			t.Errorf("The number of column of of row %v should be 2 instead of %v", i+1, len(record))
		}
	}
}

func OverlayPointInsidePolygonCSVFile(filename, delimiter string) ([][]string, error) {
	return ncsv.ReadCSVFile(filename, delimiter)
}