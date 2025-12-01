package can

import (
	"fmt"
	"log"

	"ev-gitlab.mataelang.net/ev-connect/create-ias/can.git"
	bit29can "github.com/jojohimawan/intelligent-agent-system/internal/can/gen"
	create "github.com/jojohimawan/intelligent-agent-system/internal/can/gen/create"
	j1939 "github.com/jojohimawan/intelligent-agent-system/internal/can/gen/j1939"
	"github.com/jojohimawan/intelligent-agent-system/pkg/models"
)

type OBD2Extractor func(*bit29can.OBD2) (float64, string, string)
type J1939Extractor func(can.Frame) ([]*models.DecodedSignal, error)
type CREATEExtractor func(can.Frame) ([]*models.DecodedSignal, error)

type Decoder struct {
	obd2Handlers   map[bit29can.OBD2_S01PID]OBD2Extractor
	j1939Handlers  map[uint32]J1939Extractor
	createHandlers map[uint32]CREATEExtractor
}

func NewDecoder() *Decoder {
	d := &Decoder{
		obd2Handlers:   make(map[bit29can.OBD2_S01PID]OBD2Extractor),
		j1939Handlers:  make(map[uint32]J1939Extractor),
		createHandlers: make(map[uint32]CREATEExtractor),
	}
	d.registerOBD2Handlers()
	d.registerJ1939Handlers()
	d.registerCREATEHandlers()
	return d
}

func (d *Decoder) Decode(frame can.Frame) ([]*models.DecodedSignal, error) {
	log.Printf("Incoming Frame: %x | Looking for Handler: %x\n", frame.ID, frame.ID)

	if frame.ID == bit29can.Messages().OBD2.ID {
		sig, err := d.decodeOBD2(frame)
		if err != nil {
			return nil, err
		}

		return []*models.DecodedSignal{sig}, nil
	}

	handler, exists := d.j1939Handlers[frame.ID]
	if exists {
		log.Printf("Found handler fod J1939")
		return handler(frame)
	}

	createHandler, exists := d.createHandlers[frame.ID]
	if exists {
		log.Printf("Found handler fod CReATE")
		return createHandler(frame)
	}

	log.Printf("No handler returned.")

	return nil, nil
}

func (d *Decoder) registerOBD2Handlers() {
	msg := bit29can.Messages().OBD2

	d.obd2Handlers[bit29can.OBD2_S01PID_S01PID0CEngineRPM] = func(m *bit29can.OBD2) (float64, string, string) {
		signal := msg.S01PID0C_EngineRPM
		raw := signal.UnmarshalUnsigned(m.Frame().Data)
		phys := signal.ToPhysical(float64(raw))

		return phys, signal.Name, signal.Unit
	}

	d.obd2Handlers[bit29can.OBD2_S01PID_S01PID04CalcEngineLoad] = func(m *bit29can.OBD2) (float64, string, string) {
		signal := msg.S01PID04_CalcEngineLoad
		raw := signal.UnmarshalUnsigned(m.Frame().Data)
		phys := signal.ToPhysical(float64(raw))

		return phys, signal.Name, signal.Unit
	}

	d.obd2Handlers[bit29can.OBD2_S01PID_S01PID11ThrottlePosition] = func(m *bit29can.OBD2) (float64, string, string) {
		signal := msg.S01PID11_ThrottlePosition
		raw := signal.UnmarshalUnsigned(m.Frame().Data)
		phys := signal.ToPhysical(float64(raw))

		return phys, signal.Name, signal.Unit
	}

	d.obd2Handlers[bit29can.OBD2_S01PID_S01PID05EngineCoolantTemp] = func(m *bit29can.OBD2) (float64, string, string) {
		signal := msg.S01PID05_EngineCoolantTemp
		raw := signal.UnmarshalUnsigned(m.Frame().Data)
		phys := signal.ToPhysical(float64(raw))

		return phys, signal.Name, signal.Unit
	}

	d.obd2Handlers[bit29can.OBD2_S01PID_S01PID0FIntakeAirTemperature] = func(m *bit29can.OBD2) (float64, string, string) {
		signal := msg.S01PID0F_IntakeAirTemperature
		raw := signal.UnmarshalUnsigned(m.Frame().Data)
		phys := signal.ToPhysical(float64(raw))

		return phys, signal.Name, signal.Unit
	}

	d.obd2Handlers[bit29can.OBD2_S01PID_S01PID10MAFAirFlowRate] = func(m *bit29can.OBD2) (float64, string, string) {
		signal := msg.S01PID10_MAFAirFlowRate
		raw := signal.UnmarshalUnsigned(m.Frame().Data)
		phys := signal.ToPhysical(float64(raw))

		return phys, signal.Name, signal.Unit
	}

	d.obd2Handlers[bit29can.OBD2_S01PID_S01PID0DVehicleSpeed] = func(m *bit29can.OBD2) (float64, string, string) {
		signal := msg.S01PID0D_VehicleSpeed
		raw := signal.UnmarshalUnsigned(m.Frame().Data)
		phys := signal.ToPhysical(float64(raw))

		return phys, signal.Name, signal.Unit
	}
}

func (d *Decoder) registerJ1939Handlers() {
	eec1Sig := j1939.Messages().EEC1
	ccvs1Sig := j1939.Messages().CCVS1
	lfe1Sig := j1939.Messages().LFE1
	egf1Sig := j1939.Messages().EGF1
	trf1Sig := j1939.Messages().TRF1

	d.j1939Handlers[eec1Sig.ID] = func(f can.Frame) ([]*models.DecodedSignal, error) {
		msg := j1939.NewEEC1()

		if err := msg.UnmarshalFrame(f); err != nil {
			return nil, err
		}

		var results []*models.DecodedSignal
		signal := eec1Sig.EngineSpeed
		raw := signal.UnmarshalUnsigned(f.Data)
		phys := signal.ToPhysical(float64(raw))

		results = append(results, &models.DecodedSignal{
			Source: "J1939",
			Param:  signal.Name,
			Value:  phys,
			Unit:   signal.Unit,
		})

		return results, nil
	}

	d.j1939Handlers[ccvs1Sig.ID] = func(f can.Frame) ([]*models.DecodedSignal, error) {
		msg := j1939.NewCCVS1()

		if err := msg.UnmarshalFrame(f); err != nil {
			return nil, err
		}

		var results []*models.DecodedSignal
		signal := ccvs1Sig.WheelBasedVehicleSpeed
		raw := signal.UnmarshalUnsigned(f.Data)
		phys := signal.ToPhysical(float64(raw))

		results = append(results, &models.DecodedSignal{
			Source: "J1939",
			Param:  signal.Name,
			Value:  phys,
			Unit:   signal.Unit,
		})

		return results, nil
	}

	d.j1939Handlers[lfe1Sig.ID] = func(f can.Frame) ([]*models.DecodedSignal, error) {
		msg := j1939.NewLFE1()

		if err := msg.UnmarshalFrame(f); err != nil {
			return nil, err
		}

		var results []*models.DecodedSignal
		signal := lfe1Sig.EngineFuelRate
		raw := signal.UnmarshalUnsigned(f.Data)
		phys := signal.ToPhysical(float64(raw))

		results = append(results, &models.DecodedSignal{
			Source: "J1939",
			Param:  signal.Name,
			Value:  phys,
			Unit:   signal.Unit,
		})

		return results, nil
	}

	d.j1939Handlers[egf1Sig.ID] = func(f can.Frame) ([]*models.DecodedSignal, error) {
		msg := j1939.NewEGF1()

		if err := msg.UnmarshalFrame(f); err != nil {
			return nil, err
		}

		var results []*models.DecodedSignal
		signal := egf1Sig.EngineIntakeAirMassFlowRate
		raw := signal.UnmarshalUnsigned(f.Data)
		phys := signal.ToPhysical(float64(raw))

		results = append(results, &models.DecodedSignal{
			Source: "J1939",
			Param:  signal.Name,
			Value:  phys,
			Unit:   signal.Unit,
		})

		return results, nil
	}

	d.j1939Handlers[trf1Sig.ID] = func(f can.Frame) ([]*models.DecodedSignal, error) {
		msg := j1939.NewTRF1()

		if err := msg.UnmarshalFrame(f); err != nil {
			return nil, err
		}

		var results []*models.DecodedSignal
		signal := trf1Sig.TransmissionOilTemperature1
		raw := signal.UnmarshalUnsigned(f.Data)
		phys := signal.ToPhysical(float64(raw))

		results = append(results, &models.DecodedSignal{
			Source: "J1939",
			Param:  signal.Name,
			Value:  phys,
			Unit:   signal.Unit,
		})

		return results, nil
	}
}

func (d *Decoder) registerCREATEHandlers() {
	brcmSig := create.Messages().IOV_BackRightCornerMotor1
	blcmSig := create.Messages().IOV_BackLeftCornerMotor1
	frcmSig := create.Messages().IOV_FrontRightCornerMotor1
	flcmSig := create.Messages().IOV_FrontLeftCornerMotor1

	socSig := create.Messages().BMS_DALY_SoCStatus

	d.createHandlers[brcmSig.ID] = func(f can.Frame) ([]*models.DecodedSignal, error) {
		log.Printf("Decoding BackRightCornerMotor frame..")
		msg := create.NewIOV_BackRightCornerMotor1()

		if err := msg.UnmarshalFrame(f); err != nil {
			return nil, err
		}

		var results []*models.DecodedSignal

		rpmsignal := brcmSig.VESC_StatusERPM4
		rpmraw := rpmsignal.UnmarshalSigned(f.Data)
		rpmphys := rpmsignal.ToPhysical(float64(rpmraw))
		results = append(results, &models.DecodedSignal{
			Source: "CReATE_ECU",
			Param:  rpmsignal.Name,
			Value:  rpmphys,
			Unit:   rpmsignal.Unit,
		})

		currentsignal := brcmSig.VESC_StatusCurrent4
		currentraw := currentsignal.UnmarshalSigned(f.Data)
		currentphys := currentsignal.ToPhysical(float64(currentraw))
		results = append(results, &models.DecodedSignal{
			Source: "CReATE_ECU",
			Param:  currentsignal.Name,
			Value:  currentphys,
			Unit:   currentsignal.Unit,
		})

		dutySignal := brcmSig.VESC_StatusDutyCycle4
		dutyRaw := dutySignal.UnmarshalSigned(f.Data)
		dutyPhys := dutySignal.ToPhysical(float64(dutyRaw))
		results = append(results, &models.DecodedSignal{
			Source: "CReATE_ECU",
			Param:  dutySignal.Name,
			Value:  dutyPhys,
			Unit:   dutySignal.Unit,
		})

		return results, nil
	}

	d.createHandlers[blcmSig.ID] = func(f can.Frame) ([]*models.DecodedSignal, error) {
		log.Printf("Decoding BackLeftCornerMotor frame..")
		msg := create.NewIOV_BackLeftCornerMotor1()

		if err := msg.UnmarshalFrame(f); err != nil {
			return nil, err
		}

		var results []*models.DecodedSignal

		rpmsignal := blcmSig.VESC_StatusERPM3
		rpmraw := rpmsignal.UnmarshalSigned(f.Data)
		rpmphys := rpmsignal.ToPhysical(float64(rpmraw))
		results = append(results, &models.DecodedSignal{
			Source: "CReATE_ECU",
			Param:  rpmsignal.Name,
			Value:  rpmphys,
			Unit:   rpmsignal.Unit,
		})

		currentsignal := blcmSig.VESC_StatusCurrent3
		currentraw := currentsignal.UnmarshalSigned(f.Data)
		currentphys := currentsignal.ToPhysical(float64(currentraw))
		results = append(results, &models.DecodedSignal{
			Source: "CReATE_ECU",
			Param:  currentsignal.Name,
			Value:  currentphys,
			Unit:   currentsignal.Unit,
		})

		dutySignal := blcmSig.VESC_StatusDutyCycle3
		dutyRaw := dutySignal.UnmarshalSigned(f.Data)
		dutyPhys := dutySignal.ToPhysical(float64(dutyRaw))
		results = append(results, &models.DecodedSignal{
			Source: "CReATE_ECU",
			Param:  dutySignal.Name,
			Value:  dutyPhys,
			Unit:   dutySignal.Unit,
		})

		return results, nil
	}

	d.createHandlers[frcmSig.ID] = func(f can.Frame) ([]*models.DecodedSignal, error) {
		log.Printf("Decoding FrontRightCornerMotor frame..")
		msg := create.NewIOV_FrontRightCornerMotor1()

		if err := msg.UnmarshalFrame(f); err != nil {
			return nil, err
		}

		var results []*models.DecodedSignal

		rpmsignal := frcmSig.VESC_StatusERPM2
		rpmraw := rpmsignal.UnmarshalSigned(f.Data)
		rpmphys := rpmsignal.ToPhysical(float64(rpmraw))
		results = append(results, &models.DecodedSignal{
			Source: "CReATE_ECU",
			Param:  rpmsignal.Name,
			Value:  rpmphys,
			Unit:   rpmsignal.Unit,
		})

		currentsignal := frcmSig.VESC_StatusCurrent2
		currentraw := currentsignal.UnmarshalSigned(f.Data)
		currentphys := currentsignal.ToPhysical(float64(currentraw))
		results = append(results, &models.DecodedSignal{
			Source: "CReATE_ECU",
			Param:  currentsignal.Name,
			Value:  currentphys,
			Unit:   currentsignal.Unit,
		})

		dutySignal := frcmSig.VESC_StatusDutyCycle2
		dutyRaw := dutySignal.UnmarshalSigned(f.Data)
		dutyPhys := dutySignal.ToPhysical(float64(dutyRaw))
		results = append(results, &models.DecodedSignal{
			Source: "CReATE_ECU",
			Param:  dutySignal.Name,
			Value:  dutyPhys,
			Unit:   dutySignal.Unit,
		})

		return results, nil
	}

	d.createHandlers[flcmSig.ID] = func(f can.Frame) ([]*models.DecodedSignal, error) {
		log.Printf("Decoding FrontLeftCornerMotor frame..")
		msg := create.NewIOV_FrontLeftCornerMotor1()

		if err := msg.UnmarshalFrame(f); err != nil {
			return nil, err
		}

		var results []*models.DecodedSignal

		rpmsignal := flcmSig.VESC_StatusERPM1
		rpmraw := rpmsignal.UnmarshalSigned(f.Data)
		rpmphys := rpmsignal.ToPhysical(float64(rpmraw))
		results = append(results, &models.DecodedSignal{
			Source: "CReATE_ECU",
			Param:  rpmsignal.Name,
			Value:  rpmphys,
			Unit:   rpmsignal.Unit,
		})

		currentsignal := flcmSig.VESC_StatusCurrent1
		currentraw := currentsignal.UnmarshalSigned(f.Data)
		currentphys := currentsignal.ToPhysical(float64(currentraw))
		results = append(results, &models.DecodedSignal{
			Source: "CReATE_ECU",
			Param:  currentsignal.Name,
			Value:  currentphys,
			Unit:   currentsignal.Unit,
		})

		dutySignal := flcmSig.VESC_StatusDutyCycle1
		dutyRaw := dutySignal.UnmarshalSigned(f.Data)
		dutyPhys := dutySignal.ToPhysical(float64(dutyRaw))
		results = append(results, &models.DecodedSignal{
			Source: "CReATE_ECU",
			Param:  dutySignal.Name,
			Value:  dutyPhys,
			Unit:   dutySignal.Unit,
		})

		return results, nil
	}

	d.createHandlers[socSig.ID] = func(f can.Frame) ([]*models.DecodedSignal, error) {
		log.Printf("Decoding SoC frame..")
		msg := create.NewBMS_DALY_SoCStatus()

		if err := msg.UnmarshalFrame(f); err != nil {
			return nil, err
		}

		var results []*models.DecodedSignal

		socSignal := socSig.DALY_BatterySoC
		socRaw := socSignal.UnmarshalUnsigned(f.Data)
		socPhys := socSignal.ToPhysical(float64(socRaw))
		results = append(results, &models.DecodedSignal{
			Source: "CReATE_ECU",
			Param:  socSignal.Name,
			Value:  socPhys,
			Unit:   socSignal.Unit,
		})

		currentSignal := socSig.DALY_BatteryCurrent
		currentRaw := currentSignal.UnmarshalUnsigned(f.Data)
		currentPhys := currentSignal.ToPhysical(float64(currentRaw))
		results = append(results, &models.DecodedSignal{
			Source: "CReATE_ECU",
			Param:  currentSignal.Name,
			Value:  currentPhys,
			Unit:   currentSignal.Unit,
		})

		voltageSignal := socSig.DALY_CollectedTotalVoltage
		voltageRaw := voltageSignal.UnmarshalUnsigned(f.Data)
		voltagePhys := voltageSignal.ToPhysical(float64(voltageRaw))
		results = append(results, &models.DecodedSignal{
			Source: "CReATE_ECU",
			Param:  voltageSignal.Name,
			Value:  voltagePhys,
			Unit:   voltageSignal.Unit,
		})

		pressureSignal := socSig.DALY_AccumulatedPressure
		pressureRaw := pressureSignal.UnmarshalUnsigned(f.Data)
		pressurePhys := pressureSignal.ToPhysical(float64(pressureRaw))
		results = append(results, &models.DecodedSignal{
			Source: "CReATE_ECU",
			Param:  pressureSignal.Name,
			Value:  pressurePhys,
			Unit:   pressureSignal.Unit,
		})

		return results, nil
	}
}

func (d *Decoder) decodeOBD2(frame can.Frame) (*models.DecodedSignal, error) {
	msg := bit29can.NewOBD2()
	msg.UnmarshalFrame(frame)

	handler, exists := d.obd2Handlers[msg.S01PID()]
	if !exists {
		return nil, fmt.Errorf("no handler found for PID 0x%02X", msg.S01PID())
	}

	val, name, unt := handler(msg)

	return &models.DecodedSignal{
		Source: "OBD2",
		Param:  name,
		Value:  val,
		Unit:   unt,
	}, nil
}
