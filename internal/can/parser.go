package can

import (
	"time"

	timestamppb "google.golang.org/protobuf/types/known/timestamppb"

	pb "github.com/jojohimawan/intelligent-agent-system/api"
)

func MarshalSignal(vin string, signals []*DecodedSignal) (*pb.TelematicsBatch, error) {
	var protoSignals []*pb.Telematics

	for _, s := range signals {
		protoSignals = append(protoSignals, &pb.Telematics{
			Source: s.source,
			Param:  s.param,
			Value:  &pb.Telematics_DoubleVal{DoubleVal: s.value},
			Unit:   s.unit,
		})
	}

	return &pb.TelematicsBatch{
		Vin:         vin,
		CaptureTime: timestamppb.New(time.Now()),
		Signals:     protoSignals,
	}, nil
}
