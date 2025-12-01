package can

import (
	"log"
	"time"

	timestamppb "google.golang.org/protobuf/types/known/timestamppb"

	pb "github.com/jojohimawan/intelligent-agent-system/api"
	"github.com/jojohimawan/intelligent-agent-system/internal/mapper"
	"github.com/jojohimawan/intelligent-agent-system/pkg/models"
)

func MarshalSignal(vin string, mapper *mapper.Service, signals []*models.DecodedSignal) (*pb.TelematicsBatch, error) {
	var protoSignals []*pb.Telematics

	for _, s := range signals {
		vssPoint, found := mapper.Translate(s)
		if !found {
			log.Printf("[WARN]MarshalSignal: Unknown signal, skipping...")
			continue
		}

		protoSignals = append(protoSignals, &pb.Telematics{
			Source: s.Source,
			Param:  vssPoint.Path,
			Value:  &pb.Telematics_DoubleVal{DoubleVal: s.Value.(float64)},
			Unit:   vssPoint.Unit,
		})
	}

	return &pb.TelematicsBatch{
		Vin:         vin,
		CaptureTime: timestamppb.New(time.Now()),
		Signals:     protoSignals,
	}, nil
}
