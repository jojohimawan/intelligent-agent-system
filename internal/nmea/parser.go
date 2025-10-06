package nmea

import (
	"fmt"
	"strings"
	"time"

	gonmea "github.com/adrianmo/go-nmea"
	pb "github.com/jojohimawan/intelligent-agent-system/api"
)

func ParseSentence(sentence string) (gonmea.Sentence, error) {
	return gonmea.Parse(strings.TrimSpace(sentence))
}

func SentenceToLocation(s gonmea.Sentence, vin string) (*pb.LocationRequest, error) {
	if s.DataType() != gonmea.TypeRMC {
		return nil, fmt.Errorf("unsupported sentence type: %s", s.DataType())
	}

	rmc, ok := s.(gonmea.RMC)
	if !ok {
		return nil, fmt.Errorf("failed to cast to RMC")
	}

	return &pb.LocationRequest{
		Vin:       vin,
		Lat:       rmc.Latitude,
		Lon:       rmc.Longitude,
		Timestamp: time.Now().Unix(),
	}, nil
}
