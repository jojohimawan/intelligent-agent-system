package data

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"

	pb "github.com/jojohimawan/intelligent-agent-system/api"
)

func LoadDummyLocations(path string) ([] *pb.LocationRequest, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("Failed to open CSV: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("Failed to read CSV: %w", err)
	}

	var locations []*pb.LocationRequest
	for index, row := range rows {
		if index == 0 && (row[0] == "vin" || row[1] == "lat") {
			continue
		}

		if len(row) < 3 {
			continue
		}

		vin := row[0]

		lat, err := strconv.ParseFloat(row[1], 64)
		if err != nil {
			return nil, fmt.Errorf("Invalid latitude on line %d: %w", index+1, err)
		}

		lon, err := strconv.ParseFloat(row[2], 64)
		if err != nil {
			return nil, fmt.Errorf("Invalid longitude on line %d: %w", index+1, err)
		}

		locations = append(locations, &pb.LocationRequest{
			Vin: vin,
			Lat: lat,
			Lon: lon,
			Timestamp: time.Now().Unix(),
		})
	}

	return locations, nil
}