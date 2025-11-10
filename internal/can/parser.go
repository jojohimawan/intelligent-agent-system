package can

import (
	"fmt"

	pb "github.com/jojohimawan/intelligent-agent-system/api"
)

func MessageToOBD(vin string, msg *OBD2) (*pb.VehicleOBD, error) {
	return &pb.VehicleOBD{
		Vin:     vin,
		Service: fmt.Sprintf("0x%02X", msg.service),
		Pid:     fmt.Sprintf("0x%02X", msg.pid),
		Param:   msg.param,
		Value:   msg.value,
		Unit:    msg.unit,
	}, nil
}
