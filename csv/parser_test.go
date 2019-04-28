package csv_test

import (
	"testing"

	ncsv "github.com/haidlir/OverlayTool/csv"
)

func TestParsePolygonFileOK(t *testing.T) {
	csvFileLocation := "_fixtures/polygon_ok_001.csv"
	delimiter := ";"
	polygons, err := ncsv.ParsePolygonFile(csvFileLocation, delimiter)
	if err != nil {
		t.Fatalf("ParsePolygonFile should be OK: %v", err)
	}
	if polygons == nil {
		t.Fatal("Polygons should exist")
	}
	// Check Len
	if len(polygons) != 9 {
		t.Fatalf("The polygons's length should be 9 instead of %v", len(polygons))
	}
	// Check First Row
	if &polygons[0] == nil {
		t.Fatalf("The polygons first row should be exists")
	}
	if polygons[0].ID != "16335047" {
		t.Errorf("The first row ID should be '16335047' instead of '%v'", polygons[0].ID)
	}
	firstRowCoordinateLen := len(polygons[0].Coordinates)
	if firstRowCoordinateLen != 55 {
		t.Fatalf("The first row coordinates's length should be 55 instead of %v", firstRowCoordinateLen)
	}
	if polygons[0].Coordinates[0].Lat != -7.0015284193312715 {
		t.Fatalf("The first coordinate's latitude of the first row should be '-7.0015284193312715' instead of '%v'", polygons[0].Coordinates[0].Lat)
	}
	if polygons[0].Coordinates[54].Long != 110.3313191793859 {
		t.Fatalf("The first coordinate's latitude of the first row should be '110.3313191793859' instead of '%v'", polygons[0].Coordinates[54].Long)
	}
	// Check Last Row
	if &polygons[8] == nil {
		t.Fatalf("The polygons last row should be exists")
	}
	if polygons[8].ID != "16752041" {
		t.Errorf("The last row ID should be '16752041' instead of '%v'", polygons[8].ID)
	}
	lastRowCoordinateLen := len(polygons[8].Coordinates)
	if lastRowCoordinateLen != 2 {
		t.Fatalf("The first row coordinates's length should be 2 instead of %v", lastRowCoordinateLen)
	}
	if polygons[8].Coordinates[0].Lat != -7.53422314787327 {
		t.Fatalf("The first coordinate's latitude of the first row should be '-7.53422314787327' instead of '%v'", polygons[8].Coordinates[0].Lat)
	}
	if polygons[8].Coordinates[1].Long != 110.75915124180835 {
		t.Fatalf("The first coordinate's latitude of the first row should be '110.75915124180835' instead of '%v'", polygons[8].Coordinates[1].Long)
	}
}

func TestParsePolygonFileEmpty(t *testing.T) {
	csvFileLocation := "_fixtures/empty.csv"
	delimiter := ";"
	polygons, err := ncsv.ParsePolygonFile(csvFileLocation, delimiter)
	if err == nil {
		t.Error("It should be NOK, since the file is empty")
	}
	if polygons != nil {
		t.Errorf("It should be nil, since the file is empty")
	}
}

func TestParsePolygonFileSkipped(t *testing.T) {
	csvFileLocation := "_fixtures/polygon_some-skipped_001.csv"
	delimiter := ";"
	polygons, err := ncsv.ParsePolygonFile(csvFileLocation, delimiter)
	if err != nil {
		t.Fatalf("ParsePolygonFile should be OK: %v", err)
	}
	if polygons == nil {
		t.Fatal("Polygons should exist")
	}
	// Check Len
	if len(polygons) != 5 {
		t.Fatalf("The polygons's length should be 5 instead of %v", len(polygons))
	}
}

func TestParsePolygonFileNOK(t *testing.T) {
	csvFileLocation := "_fixtures/polygon_nok_001.csv"
	delimiter := ";"
	polygons, err := ncsv.ParsePolygonFile(csvFileLocation, delimiter)
	if err == nil {
		t.Error("It should be NOK, since the file is empty")
	}
	if polygons != nil {
		t.Errorf("It should be nil, since the file is empty")
	}
}

func TestParsePointFileOK(t *testing.T) {
	csvFileLocation := "_fixtures/point_ok_001.csv"
	delimiter := ";"
	points, err := ncsv.ParsePointFile(csvFileLocation, delimiter)
	if err != nil {
		t.Fatalf("ParsePointFile should be OK: %v", err)
	}
	if points == nil {
		t.Fatal("Points should exist")
	}
	// Check Len
	if len(points) != 2 {
		t.Fatalf("The points's length should be 2 instead of %v", len(points))
	}
	// Check First Row
	if &points[0] == nil {
		t.Fatalf("The points first row should be exists")
	}
	if points[0].ID != "ODP-KUD-FH/026 FH/D02/026.01" {
		t.Errorf("The first row ID should be 'ODP-KUD-FH/026 FH/D02/026.01' instead of '%v'", points[0].ID)
	}
	if points[0].Lat != -6.81575 {
		t.Fatalf("The first row's latitude should be '-6.81575' instead of '%v'", points[0].Lat)
	}
	if points[0].Long != 110.83834 {
		t.Fatalf("The first row's longitude should be '110.83834' instead of '%v'", points[0].Long)
	}
	// Check First Row
	if &points[1] == nil {
		t.Fatalf("The points first row should be exists")
	}
	if points[1].ID != "ODP-DMA-FN/026" {
		t.Errorf("The first row ID should be 'ODP-DMA-FN/026' instead of '%v'", points[1].ID)
	}
	if points[1].Lat != -6.85746 {
		t.Fatalf("The first row's latitude should be '-6.85746' instead of '%v'", points[1].Lat)
	}
	if points[1].Long != 110.70497 {
		t.Fatalf("The first row's longitude should be '110.70497' instead of '%v'", points[1].Long)
	}
}

func TestParsePointFileEmpty(t *testing.T) {
	csvFileLocation := "_fixtures/empty.csv"
	delimiter := ";"
	points, err := ncsv.ParsePointFile(csvFileLocation, delimiter)
	if err == nil {
		t.Error("It should be NOK, since the file is empty")
	}
	if points != nil {
		t.Errorf("It should be nil, since the file is empty")
	}
}

func TestParsePointFileSkipped(t *testing.T) {
	csvFileLocation := "_fixtures/point_some-skipped_001.csv"
	delimiter := ";"
	points, err := ncsv.ParsePointFile(csvFileLocation, delimiter)
	if err != nil {
		t.Fatalf("ParsepointFile should be OK: %v", err)
	}
	if points == nil {
		t.Fatal("points should exist")
	}
	// Check Len
	if len(points) != 2 {
		t.Fatalf("The points's length should be 2 instead of %v", len(points))
	}
}

func TestParsePointFileNOK(t *testing.T) {
	csvFileLocation := "_fixtures/point_nok_001.csv"
	delimiter := ";"
	points, err := ncsv.ParsePointFile(csvFileLocation, delimiter)
	if err == nil {
		t.Error("It should be NOK, since the file is empty")
	}
	if points != nil {
		t.Errorf("It should be nil, since the file is empty")
	}
}