package main

import (
	"flag"
	"fmt"
	"log"

	nanalytic "github.com/haidlir/OverlayTool/analytic"
	ncsv "github.com/haidlir/OverlayTool/csv"
)

const (
	version = "0.0.1"
)

var polygonInputFile = flag.String("input-polygon", "polygon.csv", "the path of csv file containing list of polygon")
var pointInputFile = flag.String("input-point", "point.csv", "the path of csv file container list of point")
var outputFile = flag.String("output", "output.csv", "the path of csv file container list of result")
var inputDelimiter = flag.String("input-delimiter", ";", "the delimiter of inputted csv file")

func main() {
	flag.Parse()

	// Parse CSV
	log.Println("Parse Polygon File...")
	polygons, err := ncsv.ParsePolygonFile(*polygonInputFile, *inputDelimiter)
	if err != nil {
		log.Fatalf("unable to parse polygon input file: %v", err)
	}
	log.Println("Parse Point File...")
	points, err := ncsv.ParsePointFile(*pointInputFile, *inputDelimiter)
	if err != nil {
		log.Fatalf("unable to parse point input file: %v", err)
	}

	// Process Overlay Iteration
	log.Println("Start Iteration Process...")
	pointInsidePolygon, err := nanalytic.PointInsidePolygon(polygons, points)
	if err != nil {
		log.Printf("find error(s) while iterate point inside polygon searching process: %v", err)
	}
	if pointInsidePolygon == nil {
		log.Fatalf("found no result")
	}

	// Store results into output file
	log.Println("Store Result...")
	err = ncsv.StoreOverlayPointInsidePolygon(pointInsidePolygon, *inputDelimiter, *outputFile)
	if err != nil {
		log.Fatalf("unable to store point-inside-polygon's result: %v", err)
	}
	fmt.Printf("All Done. Check Your Result: %v", *outputFile)
}
