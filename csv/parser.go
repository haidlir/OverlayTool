package csv

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/spkg/bom"

	nmodel "github.com/haidlir/OverlayTool/model"
)

// ParsePolygonFile parses the polygon input file into polygons structure
func ParsePolygonFile(polygonInputFile, delimiter string) ([]nmodel.Polygon, error) {
	records, err := ReadCSVFile(polygonInputFile, delimiter)
	if err != nil {
		return nil, fmt.Errorf("unable to read all rows of polygon file: %v", err)
	}
	if len(records) <= 0 {
		return nil, fmt.Errorf("the records is empty")
	}
	polygons := []nmodel.Polygon{}
	for irow, row := range records {
		if len(row) != 2 {
			log.Printf("skip at row %v: the column number is not 2 (A and B) but %v", irow, len(row))
			continue
		}
		if row[0] == "" {
			log.Printf("skip at row %v: the ID (A) is empty", irow)
			continue
		}
		if row[1] == "" {
			log.Printf("skip at row %v: the coordinates (B) is empty", irow)
			continue
		}
		coordinates, err := parseCoordinates(row[1])
		if err != nil {
			log.Printf("skip at row %v: the coordinates (B) is unabled to be parsed: %v", irow, err)
			continue
		}
		id := string(bom.Clean([]byte(row[0])))
		polygon := nmodel.Polygon{
			ID:          id,
			Coordinates: coordinates,
		}
		polygons = append(polygons, polygon)
	}
	return polygons, nil
}

// ParsePointFile parses the polygon input file into polygons structure
func ParsePointFile(pointInputFile, delimiter string) ([]nmodel.Point, error) {
	records, err := ReadCSVFile(pointInputFile, delimiter)
	if err != nil {
		return nil, fmt.Errorf("unable to read all rows of point file: %v", err)
	}
	if len(records) <= 0 {
		return nil, fmt.Errorf("the records is empty")
	}
	points := []nmodel.Point{}
	for irow, row := range records {
		if len(row) != 3 {
			log.Printf("skip at row %v: the column number is not 3 (A, B and C) but %v", irow, len(row))
			continue
		}
		if row[0] == "" {
			log.Printf("skip at row %v: the ID (A) is empty", irow)
			continue
		}
		if row[1] == "" {
			log.Printf("skip at row %v: the latitude (B) is empty", irow)
			continue
		}
		if row[2] == "" {
			log.Printf("skip at row %v: the longitude (C) is empty", irow)
			continue
		}
		id := string(bom.Clean([]byte(row[0])))
		lat, err := strconv.ParseFloat(row[1], 64)
		if err != nil {
			return nil, fmt.Errorf("unable to cast to float for latitude '%v'", lat)
		}
		long, err := strconv.ParseFloat(row[2], 64)
		if err != nil {
			return nil, fmt.Errorf("unable to cast to float for longitude '%v'", long)
		}
		coordinate := nmodel.Coordinate{
			Lat:  lat,
			Long: long,
		}
		point := nmodel.Point{
			ID:          id,
			Coordinate: coordinate,
		}
		points = append(points, point)
	}
	return points, nil
}

// ReadCSVFile reads file and pass it into 2D array of string
func ReadCSVFile(fileLoc, delimiter string) ([][]string, error) {
	csvFile, err := os.Open(fileLoc)
	if err != nil {
		return nil, fmt.Errorf("unable to open file: %v", err)
	}
	reader := csv.NewReader(bufio.NewReader(csvFile))
	reader.Comma = rune(delimiter[0])
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("unable to read all rows of file: %v", err)
	}
	return records, nil
}

func parseCoordinates(coordinatesInString string) ([]nmodel.Coordinate, error) {
	coodinateInArrayString := strings.Split(coordinatesInString, ",")
	coordnates := []nmodel.Coordinate{}
	for _, coordinateInString := range coodinateInArrayString {
		coordinateInArray := strings.Split(coordinateInString, " ")
		if len(coordinateInArray) != 2 {
			return nil, fmt.Errorf("wrong coordinate format: '%v'", coordinateInString)
		}
		lat, err := strconv.ParseFloat(coordinateInArray[0], 64)
		if err != nil {
			return nil, fmt.Errorf("unable to cast to float for coordinate '%v'", coordinateInString)
		}
		long, err := strconv.ParseFloat(coordinateInArray[1], 64)
		if err != nil {
			return nil, fmt.Errorf("unable to cast to float for coordinate '%v'", coordinateInString)
		}
		coordinate := nmodel.Coordinate{
			Lat:  lat,
			Long: long,
		}
		coordnates = append(coordnates, coordinate)
	}
	return coordnates, nil
}
