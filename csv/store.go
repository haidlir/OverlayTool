package csv

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"

	nmodel "github.com/haidlir/OverlayTool/model"
)

// StoreOverlayPointInsidePolygon stores the result of point inside polygon into file
func StoreOverlayPointInsidePolygon(pointInsidePolygon []nmodel.OverlayPointInsidePolygon, delimiter, outputFilename string) error {
    file, err := os.Create(outputFilename)
    if err != nil {
		return fmt.Errorf("unable to create file: %v", err)
	}
    defer file.Close()
	records := [][]string{}
	for _, row := range(pointInsidePolygon) {
		pointsStr := []string{}
		for _, point := range row.Points {
			pointsStr = append(pointsStr, point.ID)
		}
		record := []string{fmt.Sprintf("%v (%v)", row.PolygonID, len(row.Points)), strings.Join(pointsStr, "|")}
		records = append(records, record)
	}
	writer := csv.NewWriter(file)
    defer writer.Flush()
	writer.Comma = rune(delimiter[0])
	writer.WriteAll(records) // calls Flush internally
	if err := writer.Error(); err != nil {
		return fmt.Errorf("error writing csv: %v", err)
	}
	return nil
}
