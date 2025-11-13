package can

import (
	"fmt"

	"ev-gitlab.mataelang.net/ev-connect/create-ias/can.git"
	bit29can "github.com/jojohimawan/intelligent-agent-system/internal/can/gen"
)

type OBD2 struct {
	message *bit29can.OBD2
	service byte
	pid     byte
	param   string
	value   float64
	unit    string
}

func DecodeMode01PID(frame can.Frame) (*OBD2, error) {
	if frame.ID != bit29can.Messages().OBD2.ID || !frame.IsExtended {
		fmt.Printf("Unknown frame ID or frame is not extended. Skipping...")
		return nil, fmt.Errorf("Unknown frame ID or frame is not extended. Skipping...")
	}

	o := &OBD2{
		message: bit29can.NewOBD2(),
	}

	if err := o.message.UnmarshalFrame(frame); err != nil {
		fmt.Printf("   Raw data: % X\n", frame.Data[:frame.Length])
		return nil, fmt.Errorf("⚠️  Failed to unmarshal frame: %v", err)
	}

	pid := o.message.S01PID()
	fmt.Printf("📋 S01PID raw: %d (0x%02X)\n", uint8(o.message.S01PID()), uint8(o.message.S01PID()))

	switch pid {
	case bit29can.OBD2_S01PID_S01PID0CEngineRPM:
		rpmSignal := bit29can.Messages().OBD2.S01PID0C_EngineRPM
		rpmRaw := rpmSignal.UnmarshalUnsigned(o.message.Frame().Data)
		rpmPhys := rpmSignal.ToPhysical(float64(rpmRaw))

		fmt.Printf("RPM (from high-level accessor): %.2f\n", o.message.S01PID0C_EngineRPM())
		fmt.Printf("RPM (from manual unmarshal): %.2f %s\n", rpmPhys, rpmSignal.Unit)

		return &OBD2{
			service: uint8(o.message.Service()),
			pid:     uint8(pid),
			param:   "RPM",
			value:   rpmPhys,
			unit:    rpmSignal.Unit,
		}, nil
	case bit29can.OBD2_S01PID_S01PID04CalcEngineLoad:
		engLoadSignal := bit29can.Messages().OBD2.S01PID04_CalcEngineLoad
		engLoadRaw := engLoadSignal.UnmarshalUnsigned(o.message.Frame().Data)
		engLoadPhys := engLoadSignal.ToPhysical(float64(engLoadRaw))

		fmt.Printf("Engine Load (from high-level accessor): %.2f\n", o.message.S01PID04_CalcEngineLoad())
		fmt.Printf("Engine Load (from manual unmarshal): %.2f%s\n", engLoadPhys, engLoadSignal.Unit)

		return &OBD2{
			service: uint8(o.message.Service()),
			pid:     uint8(pid),
			param:   "Engine Load",
			value:   engLoadPhys,
			unit:    engLoadSignal.Unit,
		}, nil

	case bit29can.OBD2_S01PID_S01PID11ThrottlePosition:
		throtSignal := bit29can.Messages().OBD2.S01PID11_ThrottlePosition
		throtRaw := throtSignal.UnmarshalUnsigned(o.message.Frame().Data)
		throtPhys := throtSignal.ToPhysical(float64(throtRaw))

		fmt.Printf("Throttle Manifold (from high-level accessor): %.2f\n", o.message.S01PID11_ThrottlePosition())
		fmt.Printf("Throttle Manifold (from manual unmarshal): %.2f%s\n", throtPhys, throtSignal.Unit)

		return &OBD2{
			service: uint8(o.message.Service()),
			pid:     uint8(pid),
			param:   "Throttle Manifold",
			value:   throtPhys,
			unit:    throtSignal.Unit,
		}, nil

	case bit29can.OBD2_S01PID_S01PID05EngineCoolantTemp:
		coolantSignal := bit29can.Messages().OBD2.S01PID05_EngineCoolantTemp
		coolantRaw := coolantSignal.UnmarshalUnsigned(o.message.Frame().Data)
		coolantPhys := coolantSignal.ToPhysical(float64(coolantRaw))

		fmt.Printf("Engine Coolant Temp (from high-level accessor): %.2f\n", o.message.S01PID05_EngineCoolantTemp())
		fmt.Printf("Engine Coolant Temp (from manual unmarshal): %.2f%s\n", coolantPhys, coolantSignal.Unit)

		return &OBD2{
			service: uint8(o.message.Service()),
			pid:     uint8(pid),
			param:   "Engine Coolant Temperature",
			value:   coolantPhys,
			unit:    coolantSignal.Unit,
		}, nil
	case bit29can.OBD2_S01PID_S01PID0FIntakeAirTemperature:
		airSignal := bit29can.Messages().OBD2.S01PID0F_IntakeAirTemperature
		airRaw := airSignal.UnmarshalUnsigned(o.message.Frame().Data)
		airPhys := airSignal.ToPhysical(float64(airRaw))

		fmt.Printf("Intake Air Temp (from high-level accessor): %.2f\n", o.message.S01PID0F_IntakeAirTemperature())
		fmt.Printf("Intake Air Temp (from manual unmarshal): %.2f%s\n", airPhys, airSignal.Unit)

		return &OBD2{
			service: uint8(o.message.Service()),
			pid:     uint8(pid),
			param:   "Intake Air Temperature",
			value:   airPhys,
			unit:    airSignal.Unit,
		}, nil
	case bit29can.OBD2_S01PID_S01PID10MAFAirFlowRate:
		airflowSignal := bit29can.Messages().OBD2.S01PID10_MAFAirFlowRate
		airflowRaw := airflowSignal.UnmarshalUnsigned(o.message.Frame().Data)
		airflowPhys := airflowSignal.ToPhysical(float64(airflowRaw))

		fmt.Printf("Airflow Rate (from high-level accessor): %.2f\n", o.message.S01PID10_MAFAirFlowRate())
		fmt.Printf("Airflow Rate (from manual unmarshal): %.2f%s\n", airflowPhys, airflowSignal.Unit)

		return &OBD2{
			service: uint8(o.message.Service()),
			pid:     uint8(pid),
			param:   "Airflow Rate",
			value:   airflowPhys,
			unit:    airflowSignal.Unit,
		}, nil
	case bit29can.OBD2_S01PID_S01PID0DVehicleSpeed:
		speedSignal := bit29can.Messages().OBD2.S01PID0D_VehicleSpeed
		speedRaw := speedSignal.UnmarshalUnsigned(o.message.Frame().Data)
		speedPhys := speedSignal.ToPhysical(float64(speedRaw))

		fmt.Printf("Speed (from high-level accessor): %d\n", o.message.S01PID0D_VehicleSpeed())
		fmt.Printf("Speed (from manual unmarshal): %.2f%s\n", speedPhys, speedSignal.Unit)

		return &OBD2{
			service: uint8(o.message.Service()),
			pid:     uint8(pid),
			param:   "Speed",
			value:   speedPhys,
			unit:    speedSignal.Unit,
		}, nil
	default:
		fmt.Printf("unhandled Mode 01 PID: 0x%02X", pid)
		return nil, fmt.Errorf("unhandled Mode 01 PID: 0x%02X", pid)
	}

}
