// Package j1939can provides primitives for encoding and decoding j1939 CAN messages.
//
// Source: dummy/j1939/j1939.dbc
package j1939can

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"sync"
	"time"

	"ev-gitlab.mataelang.net/ev-connect/create-ias/can.git"
	"ev-gitlab.mataelang.net/ev-connect/create-ias/can.git/pkg/candebug"
	"ev-gitlab.mataelang.net/ev-connect/create-ias/can.git/pkg/canrunner"
	"ev-gitlab.mataelang.net/ev-connect/create-ias/can.git/pkg/cantext"
	"ev-gitlab.mataelang.net/ev-connect/create-ias/can.git/pkg/descriptor"
	"ev-gitlab.mataelang.net/ev-connect/create-ias/can.git/pkg/generated"
	"ev-gitlab.mataelang.net/ev-connect/create-ias/can.git/pkg/socketcan"
)

// prevent unused imports
var (
	_ = context.Background
	_ = fmt.Print
	_ = net.Dial
	_ = http.Error
	_ = sync.Mutex{}
	_ = time.Now
	_ = socketcan.Dial
	_ = candebug.ServeMessagesHTTP
	_ = canrunner.Run
)

// Generated code. DO NOT EDIT.
// EEC1Reader provides read access to a EEC1 message.
type EEC1Reader interface {
	can.FrameMarshaler
	// EngineSpeed returns the physical value of the EngineSpeed signal.
	EngineSpeed() float64
	// RawEngineSpeed returns the raw (encoded) value of the EngineSpeed signal.
	RawEngineSpeed() uint16
}

// EEC1Writer provides write access to a EEC1 message.
type EEC1Writer interface {
	// CopyFrom copies all values from EEC1.
	CopyFrom(EEC1Reader) *EEC1
	// SetEngineSpeed sets the physical value of the EngineSpeed signal.
	SetEngineSpeed(float64) *EEC1
	// SetRawEngineSpeed sets the raw (encoded) value of the EngineSpeed signal.
	SetRawEngineSpeed(uint16) *EEC1
}

type EEC1 struct {
	xxx_EngineSpeed uint16
}

func NewEEC1() *EEC1 {
	m := &EEC1{}
	m.Reset()
	return m
}

func (m *EEC1) Reset() {
	m.xxx_EngineSpeed = 0
}

func (m *EEC1) CopyFrom(o EEC1Reader) *EEC1 {
	f, _ := o.MarshalFrame()
	_ = m.UnmarshalFrame(f)
	return m
}

// Descriptor returns the EEC1 descriptor.
func (m *EEC1) Descriptor() *descriptor.Message {
	return Messages().EEC1.Message
}

// String returns a compact string representation of the message.
func (m *EEC1) String() string {
	return cantext.MessageString(m)
}

func (m *EEC1) EngineSpeed() float64 {
	return Messages().EEC1.EngineSpeed.ToPhysical(float64(m.xxx_EngineSpeed))
}

func (m *EEC1) SetEngineSpeed(v float64) *EEC1 {
	m.xxx_EngineSpeed = uint16(Messages().EEC1.EngineSpeed.FromPhysical(v))
	return m
}

func (m *EEC1) RawEngineSpeed() uint16 {
	return m.xxx_EngineSpeed
}

func (m *EEC1) SetRawEngineSpeed(v uint16) *EEC1 {
	m.xxx_EngineSpeed = uint16(Messages().EEC1.EngineSpeed.SaturatedCastUnsigned(uint64(v)))
	return m
}

// Frame returns a CAN frame representing the message.
func (m *EEC1) Frame() can.Frame {
	md := Messages().EEC1
	f := can.Frame{ID: md.ID, IsExtended: md.IsExtended, Length: md.Length}
	md.EngineSpeed.MarshalUnsigned(&f.Data, uint64(m.xxx_EngineSpeed))
	return f
}

// MarshalFrame encodes the message as a CAN frame.
func (m *EEC1) MarshalFrame() (can.Frame, error) {
	return m.Frame(), nil
}

// UnmarshalFrame decodes the message from a CAN frame.
func (m *EEC1) UnmarshalFrame(f can.Frame) error {
	md := Messages().EEC1
	switch {
	case f.ID != md.ID:
		return fmt.Errorf(
			"unmarshal EEC1: expects ID 217056510 (got %s with ID %d)", f.String(), f.ID,
		)
	case f.Length != md.Length:
		return fmt.Errorf(
			"unmarshal EEC1: expects length 8 (got %s with length %d)", f.String(), f.Length,
		)
	case f.IsRemote:
		return fmt.Errorf(
			"unmarshal EEC1: expects non-remote frame (got remote frame %s)", f.String(),
		)
	case f.IsExtended != md.IsExtended:
		return fmt.Errorf(
			"unmarshal EEC1: expects extended ID (got %s with standard ID)", f.String(),
		)
	}
	m.xxx_EngineSpeed = uint16(md.EngineSpeed.UnmarshalUnsigned(f.Data))
	return nil
}

// EGF1Reader provides read access to a EGF1 message.
type EGF1Reader interface {
	can.FrameMarshaler
	// EngineIntakeAirMassFlowRate returns the physical value of the EngineIntakeAirMassFlowRate signal.
	EngineIntakeAirMassFlowRate() float64
	// RawEngineIntakeAirMassFlowRate returns the raw (encoded) value of the EngineIntakeAirMassFlowRate signal.
	RawEngineIntakeAirMassFlowRate() uint16
}

// EGF1Writer provides write access to a EGF1 message.
type EGF1Writer interface {
	// CopyFrom copies all values from EGF1.
	CopyFrom(EGF1Reader) *EGF1
	// SetEngineIntakeAirMassFlowRate sets the physical value of the EngineIntakeAirMassFlowRate signal.
	SetEngineIntakeAirMassFlowRate(float64) *EGF1
	// SetRawEngineIntakeAirMassFlowRate sets the raw (encoded) value of the EngineIntakeAirMassFlowRate signal.
	SetRawEngineIntakeAirMassFlowRate(uint16) *EGF1
}

type EGF1 struct {
	xxx_EngineIntakeAirMassFlowRate uint16
}

func NewEGF1() *EGF1 {
	m := &EGF1{}
	m.Reset()
	return m
}

func (m *EGF1) Reset() {
	m.xxx_EngineIntakeAirMassFlowRate = 0
}

func (m *EGF1) CopyFrom(o EGF1Reader) *EGF1 {
	f, _ := o.MarshalFrame()
	_ = m.UnmarshalFrame(f)
	return m
}

// Descriptor returns the EGF1 descriptor.
func (m *EGF1) Descriptor() *descriptor.Message {
	return Messages().EGF1.Message
}

// String returns a compact string representation of the message.
func (m *EGF1) String() string {
	return cantext.MessageString(m)
}

func (m *EGF1) EngineIntakeAirMassFlowRate() float64 {
	return Messages().EGF1.EngineIntakeAirMassFlowRate.ToPhysical(float64(m.xxx_EngineIntakeAirMassFlowRate))
}

func (m *EGF1) SetEngineIntakeAirMassFlowRate(v float64) *EGF1 {
	m.xxx_EngineIntakeAirMassFlowRate = uint16(Messages().EGF1.EngineIntakeAirMassFlowRate.FromPhysical(v))
	return m
}

func (m *EGF1) RawEngineIntakeAirMassFlowRate() uint16 {
	return m.xxx_EngineIntakeAirMassFlowRate
}

func (m *EGF1) SetRawEngineIntakeAirMassFlowRate(v uint16) *EGF1 {
	m.xxx_EngineIntakeAirMassFlowRate = uint16(Messages().EGF1.EngineIntakeAirMassFlowRate.SaturatedCastUnsigned(uint64(v)))
	return m
}

// Frame returns a CAN frame representing the message.
func (m *EGF1) Frame() can.Frame {
	md := Messages().EGF1
	f := can.Frame{ID: md.ID, IsExtended: md.IsExtended, Length: md.Length}
	md.EngineIntakeAirMassFlowRate.MarshalUnsigned(&f.Data, uint64(m.xxx_EngineIntakeAirMassFlowRate))
	return f
}

// MarshalFrame encodes the message as a CAN frame.
func (m *EGF1) MarshalFrame() (can.Frame, error) {
	return m.Frame(), nil
}

// UnmarshalFrame decodes the message from a CAN frame.
func (m *EGF1) UnmarshalFrame(f can.Frame) error {
	md := Messages().EGF1
	switch {
	case f.ID != md.ID:
		return fmt.Errorf(
			"unmarshal EGF1: expects ID 217058046 (got %s with ID %d)", f.String(), f.ID,
		)
	case f.Length != md.Length:
		return fmt.Errorf(
			"unmarshal EGF1: expects length 8 (got %s with length %d)", f.String(), f.Length,
		)
	case f.IsRemote:
		return fmt.Errorf(
			"unmarshal EGF1: expects non-remote frame (got remote frame %s)", f.String(),
		)
	case f.IsExtended != md.IsExtended:
		return fmt.Errorf(
			"unmarshal EGF1: expects extended ID (got %s with standard ID)", f.String(),
		)
	}
	m.xxx_EngineIntakeAirMassFlowRate = uint16(md.EngineIntakeAirMassFlowRate.UnmarshalUnsigned(f.Data))
	return nil
}

// CCVS1Reader provides read access to a CCVS1 message.
type CCVS1Reader interface {
	can.FrameMarshaler
	// WheelBasedVehicleSpeed returns the physical value of the WheelBasedVehicleSpeed signal.
	WheelBasedVehicleSpeed() float64
	// RawWheelBasedVehicleSpeed returns the raw (encoded) value of the WheelBasedVehicleSpeed signal.
	RawWheelBasedVehicleSpeed() uint16
}

// CCVS1Writer provides write access to a CCVS1 message.
type CCVS1Writer interface {
	// CopyFrom copies all values from CCVS1.
	CopyFrom(CCVS1Reader) *CCVS1
	// SetWheelBasedVehicleSpeed sets the physical value of the WheelBasedVehicleSpeed signal.
	SetWheelBasedVehicleSpeed(float64) *CCVS1
	// SetRawWheelBasedVehicleSpeed sets the raw (encoded) value of the WheelBasedVehicleSpeed signal.
	SetRawWheelBasedVehicleSpeed(uint16) *CCVS1
}

type CCVS1 struct {
	xxx_WheelBasedVehicleSpeed uint16
}

func NewCCVS1() *CCVS1 {
	m := &CCVS1{}
	m.Reset()
	return m
}

func (m *CCVS1) Reset() {
	m.xxx_WheelBasedVehicleSpeed = 0
}

func (m *CCVS1) CopyFrom(o CCVS1Reader) *CCVS1 {
	f, _ := o.MarshalFrame()
	_ = m.UnmarshalFrame(f)
	return m
}

// Descriptor returns the CCVS1 descriptor.
func (m *CCVS1) Descriptor() *descriptor.Message {
	return Messages().CCVS1.Message
}

// String returns a compact string representation of the message.
func (m *CCVS1) String() string {
	return cantext.MessageString(m)
}

func (m *CCVS1) WheelBasedVehicleSpeed() float64 {
	return Messages().CCVS1.WheelBasedVehicleSpeed.ToPhysical(float64(m.xxx_WheelBasedVehicleSpeed))
}

func (m *CCVS1) SetWheelBasedVehicleSpeed(v float64) *CCVS1 {
	m.xxx_WheelBasedVehicleSpeed = uint16(Messages().CCVS1.WheelBasedVehicleSpeed.FromPhysical(v))
	return m
}

func (m *CCVS1) RawWheelBasedVehicleSpeed() uint16 {
	return m.xxx_WheelBasedVehicleSpeed
}

func (m *CCVS1) SetRawWheelBasedVehicleSpeed(v uint16) *CCVS1 {
	m.xxx_WheelBasedVehicleSpeed = uint16(Messages().CCVS1.WheelBasedVehicleSpeed.SaturatedCastUnsigned(uint64(v)))
	return m
}

// Frame returns a CAN frame representing the message.
func (m *CCVS1) Frame() can.Frame {
	md := Messages().CCVS1
	f := can.Frame{ID: md.ID, IsExtended: md.IsExtended, Length: md.Length}
	md.WheelBasedVehicleSpeed.MarshalUnsigned(&f.Data, uint64(m.xxx_WheelBasedVehicleSpeed))
	return f
}

// MarshalFrame encodes the message as a CAN frame.
func (m *CCVS1) MarshalFrame() (can.Frame, error) {
	return m.Frame(), nil
}

// UnmarshalFrame decodes the message from a CAN frame.
func (m *CCVS1) UnmarshalFrame(f can.Frame) error {
	md := Messages().CCVS1
	switch {
	case f.ID != md.ID:
		return fmt.Errorf(
			"unmarshal CCVS1: expects ID 419361278 (got %s with ID %d)", f.String(), f.ID,
		)
	case f.Length != md.Length:
		return fmt.Errorf(
			"unmarshal CCVS1: expects length 8 (got %s with length %d)", f.String(), f.Length,
		)
	case f.IsRemote:
		return fmt.Errorf(
			"unmarshal CCVS1: expects non-remote frame (got remote frame %s)", f.String(),
		)
	case f.IsExtended != md.IsExtended:
		return fmt.Errorf(
			"unmarshal CCVS1: expects extended ID (got %s with standard ID)", f.String(),
		)
	}
	m.xxx_WheelBasedVehicleSpeed = uint16(md.WheelBasedVehicleSpeed.UnmarshalUnsigned(f.Data))
	return nil
}

// LFE1Reader provides read access to a LFE1 message.
type LFE1Reader interface {
	can.FrameMarshaler
	// EngineFuelRate returns the physical value of the EngineFuelRate signal.
	EngineFuelRate() float64
	// RawEngineFuelRate returns the raw (encoded) value of the EngineFuelRate signal.
	RawEngineFuelRate() uint16
}

// LFE1Writer provides write access to a LFE1 message.
type LFE1Writer interface {
	// CopyFrom copies all values from LFE1.
	CopyFrom(LFE1Reader) *LFE1
	// SetEngineFuelRate sets the physical value of the EngineFuelRate signal.
	SetEngineFuelRate(float64) *LFE1
	// SetRawEngineFuelRate sets the raw (encoded) value of the EngineFuelRate signal.
	SetRawEngineFuelRate(uint16) *LFE1
}

type LFE1 struct {
	xxx_EngineFuelRate uint16
}

func NewLFE1() *LFE1 {
	m := &LFE1{}
	m.Reset()
	return m
}

func (m *LFE1) Reset() {
	m.xxx_EngineFuelRate = 0
}

func (m *LFE1) CopyFrom(o LFE1Reader) *LFE1 {
	f, _ := o.MarshalFrame()
	_ = m.UnmarshalFrame(f)
	return m
}

// Descriptor returns the LFE1 descriptor.
func (m *LFE1) Descriptor() *descriptor.Message {
	return Messages().LFE1.Message
}

// String returns a compact string representation of the message.
func (m *LFE1) String() string {
	return cantext.MessageString(m)
}

func (m *LFE1) EngineFuelRate() float64 {
	return Messages().LFE1.EngineFuelRate.ToPhysical(float64(m.xxx_EngineFuelRate))
}

func (m *LFE1) SetEngineFuelRate(v float64) *LFE1 {
	m.xxx_EngineFuelRate = uint16(Messages().LFE1.EngineFuelRate.FromPhysical(v))
	return m
}

func (m *LFE1) RawEngineFuelRate() uint16 {
	return m.xxx_EngineFuelRate
}

func (m *LFE1) SetRawEngineFuelRate(v uint16) *LFE1 {
	m.xxx_EngineFuelRate = uint16(Messages().LFE1.EngineFuelRate.SaturatedCastUnsigned(uint64(v)))
	return m
}

// Frame returns a CAN frame representing the message.
func (m *LFE1) Frame() can.Frame {
	md := Messages().LFE1
	f := can.Frame{ID: md.ID, IsExtended: md.IsExtended, Length: md.Length}
	md.EngineFuelRate.MarshalUnsigned(&f.Data, uint64(m.xxx_EngineFuelRate))
	return f
}

// MarshalFrame encodes the message as a CAN frame.
func (m *LFE1) MarshalFrame() (can.Frame, error) {
	return m.Frame(), nil
}

// UnmarshalFrame decodes the message from a CAN frame.
func (m *LFE1) UnmarshalFrame(f can.Frame) error {
	md := Messages().LFE1
	switch {
	case f.ID != md.ID:
		return fmt.Errorf(
			"unmarshal LFE1: expects ID 419361534 (got %s with ID %d)", f.String(), f.ID,
		)
	case f.Length != md.Length:
		return fmt.Errorf(
			"unmarshal LFE1: expects length 8 (got %s with length %d)", f.String(), f.Length,
		)
	case f.IsRemote:
		return fmt.Errorf(
			"unmarshal LFE1: expects non-remote frame (got remote frame %s)", f.String(),
		)
	case f.IsExtended != md.IsExtended:
		return fmt.Errorf(
			"unmarshal LFE1: expects extended ID (got %s with standard ID)", f.String(),
		)
	}
	m.xxx_EngineFuelRate = uint16(md.EngineFuelRate.UnmarshalUnsigned(f.Data))
	return nil
}

// VP1Reader provides read access to a VP1 message.
type VP1Reader interface {
	can.FrameMarshaler
	// Latitude returns the physical value of the Latitude signal.
	Latitude() float64
	// RawLatitude returns the raw (encoded) value of the Latitude signal.
	RawLatitude() uint32
	// Longitude returns the physical value of the Longitude signal.
	Longitude() float64
	// RawLongitude returns the raw (encoded) value of the Longitude signal.
	RawLongitude() uint32
}

// VP1Writer provides write access to a VP1 message.
type VP1Writer interface {
	// CopyFrom copies all values from VP1.
	CopyFrom(VP1Reader) *VP1
	// SetLatitude sets the physical value of the Latitude signal.
	SetLatitude(float64) *VP1
	// SetRawLatitude sets the raw (encoded) value of the Latitude signal.
	SetRawLatitude(uint32) *VP1
	// SetLongitude sets the physical value of the Longitude signal.
	SetLongitude(float64) *VP1
	// SetRawLongitude sets the raw (encoded) value of the Longitude signal.
	SetRawLongitude(uint32) *VP1
}

type VP1 struct {
	xxx_Latitude  uint32
	xxx_Longitude uint32
}

func NewVP1() *VP1 {
	m := &VP1{}
	m.Reset()
	return m
}

func (m *VP1) Reset() {
	m.xxx_Latitude = 0
	m.xxx_Longitude = 0
}

func (m *VP1) CopyFrom(o VP1Reader) *VP1 {
	f, _ := o.MarshalFrame()
	_ = m.UnmarshalFrame(f)
	return m
}

// Descriptor returns the VP1 descriptor.
func (m *VP1) Descriptor() *descriptor.Message {
	return Messages().VP1.Message
}

// String returns a compact string representation of the message.
func (m *VP1) String() string {
	return cantext.MessageString(m)
}

func (m *VP1) Latitude() float64 {
	return Messages().VP1.Latitude.ToPhysical(float64(m.xxx_Latitude))
}

func (m *VP1) SetLatitude(v float64) *VP1 {
	m.xxx_Latitude = uint32(Messages().VP1.Latitude.FromPhysical(v))
	return m
}

func (m *VP1) RawLatitude() uint32 {
	return m.xxx_Latitude
}

func (m *VP1) SetRawLatitude(v uint32) *VP1 {
	m.xxx_Latitude = uint32(Messages().VP1.Latitude.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *VP1) Longitude() float64 {
	return Messages().VP1.Longitude.ToPhysical(float64(m.xxx_Longitude))
}

func (m *VP1) SetLongitude(v float64) *VP1 {
	m.xxx_Longitude = uint32(Messages().VP1.Longitude.FromPhysical(v))
	return m
}

func (m *VP1) RawLongitude() uint32 {
	return m.xxx_Longitude
}

func (m *VP1) SetRawLongitude(v uint32) *VP1 {
	m.xxx_Longitude = uint32(Messages().VP1.Longitude.SaturatedCastUnsigned(uint64(v)))
	return m
}

// Frame returns a CAN frame representing the message.
func (m *VP1) Frame() can.Frame {
	md := Messages().VP1
	f := can.Frame{ID: md.ID, IsExtended: md.IsExtended, Length: md.Length}
	md.Latitude.MarshalUnsigned(&f.Data, uint64(m.xxx_Latitude))
	md.Longitude.MarshalUnsigned(&f.Data, uint64(m.xxx_Longitude))
	return f
}

// MarshalFrame encodes the message as a CAN frame.
func (m *VP1) MarshalFrame() (can.Frame, error) {
	return m.Frame(), nil
}

// UnmarshalFrame decodes the message from a CAN frame.
func (m *VP1) UnmarshalFrame(f can.Frame) error {
	md := Messages().VP1
	switch {
	case f.ID != md.ID:
		return fmt.Errorf(
			"unmarshal VP1: expects ID 419361790 (got %s with ID %d)", f.String(), f.ID,
		)
	case f.Length != md.Length:
		return fmt.Errorf(
			"unmarshal VP1: expects length 8 (got %s with length %d)", f.String(), f.Length,
		)
	case f.IsRemote:
		return fmt.Errorf(
			"unmarshal VP1: expects non-remote frame (got remote frame %s)", f.String(),
		)
	case f.IsExtended != md.IsExtended:
		return fmt.Errorf(
			"unmarshal VP1: expects extended ID (got %s with standard ID)", f.String(),
		)
	}
	m.xxx_Latitude = uint32(md.Latitude.UnmarshalUnsigned(f.Data))
	m.xxx_Longitude = uint32(md.Longitude.UnmarshalUnsigned(f.Data))
	return nil
}

// TRF1Reader provides read access to a TRF1 message.
type TRF1Reader interface {
	can.FrameMarshaler
	// TransmissionOilTemperature1 returns the physical value of the TransmissionOilTemperature1 signal.
	TransmissionOilTemperature1() float64
	// RawTransmissionOilTemperature1 returns the raw (encoded) value of the TransmissionOilTemperature1 signal.
	RawTransmissionOilTemperature1() uint16
}

// TRF1Writer provides write access to a TRF1 message.
type TRF1Writer interface {
	// CopyFrom copies all values from TRF1.
	CopyFrom(TRF1Reader) *TRF1
	// SetTransmissionOilTemperature1 sets the physical value of the TransmissionOilTemperature1 signal.
	SetTransmissionOilTemperature1(float64) *TRF1
	// SetRawTransmissionOilTemperature1 sets the raw (encoded) value of the TransmissionOilTemperature1 signal.
	SetRawTransmissionOilTemperature1(uint16) *TRF1
}

type TRF1 struct {
	xxx_TransmissionOilTemperature1 uint16
}

func NewTRF1() *TRF1 {
	m := &TRF1{}
	m.Reset()
	return m
}

func (m *TRF1) Reset() {
	m.xxx_TransmissionOilTemperature1 = 0
}

func (m *TRF1) CopyFrom(o TRF1Reader) *TRF1 {
	f, _ := o.MarshalFrame()
	_ = m.UnmarshalFrame(f)
	return m
}

// Descriptor returns the TRF1 descriptor.
func (m *TRF1) Descriptor() *descriptor.Message {
	return Messages().TRF1.Message
}

// String returns a compact string representation of the message.
func (m *TRF1) String() string {
	return cantext.MessageString(m)
}

func (m *TRF1) TransmissionOilTemperature1() float64 {
	return Messages().TRF1.TransmissionOilTemperature1.ToPhysical(float64(m.xxx_TransmissionOilTemperature1))
}

func (m *TRF1) SetTransmissionOilTemperature1(v float64) *TRF1 {
	m.xxx_TransmissionOilTemperature1 = uint16(Messages().TRF1.TransmissionOilTemperature1.FromPhysical(v))
	return m
}

func (m *TRF1) RawTransmissionOilTemperature1() uint16 {
	return m.xxx_TransmissionOilTemperature1
}

func (m *TRF1) SetRawTransmissionOilTemperature1(v uint16) *TRF1 {
	m.xxx_TransmissionOilTemperature1 = uint16(Messages().TRF1.TransmissionOilTemperature1.SaturatedCastUnsigned(uint64(v)))
	return m
}

// Frame returns a CAN frame representing the message.
func (m *TRF1) Frame() can.Frame {
	md := Messages().TRF1
	f := can.Frame{ID: md.ID, IsExtended: md.IsExtended, Length: md.Length}
	md.TransmissionOilTemperature1.MarshalUnsigned(&f.Data, uint64(m.xxx_TransmissionOilTemperature1))
	return f
}

// MarshalFrame encodes the message as a CAN frame.
func (m *TRF1) MarshalFrame() (can.Frame, error) {
	return m.Frame(), nil
}

// UnmarshalFrame decodes the message from a CAN frame.
func (m *TRF1) UnmarshalFrame(f can.Frame) error {
	md := Messages().TRF1
	switch {
	case f.ID != md.ID:
		return fmt.Errorf(
			"unmarshal TRF1: expects ID 419363070 (got %s with ID %d)", f.String(), f.ID,
		)
	case f.Length != md.Length:
		return fmt.Errorf(
			"unmarshal TRF1: expects length 8 (got %s with length %d)", f.String(), f.Length,
		)
	case f.IsRemote:
		return fmt.Errorf(
			"unmarshal TRF1: expects non-remote frame (got remote frame %s)", f.String(),
		)
	case f.IsExtended != md.IsExtended:
		return fmt.Errorf(
			"unmarshal TRF1: expects extended ID (got %s with standard ID)", f.String(),
		)
	}
	m.xxx_TransmissionOilTemperature1 = uint16(md.TransmissionOilTemperature1.UnmarshalUnsigned(f.Data))
	return nil
}

// Nodes returns the j1939 node descriptors.
func Nodes() *NodesDescriptor {
	return nd
}

// NodesDescriptor contains all j1939 node descriptors.
type NodesDescriptor struct {
}

// Messages returns the j1939 message descriptors.
func Messages() *MessagesDescriptor {
	return md
}

// MessagesDescriptor contains all j1939 message descriptors.
type MessagesDescriptor struct {
	EEC1  *EEC1Descriptor
	EGF1  *EGF1Descriptor
	CCVS1 *CCVS1Descriptor
	LFE1  *LFE1Descriptor
	VP1   *VP1Descriptor
	TRF1  *TRF1Descriptor
}

// UnmarshalFrame unmarshals the provided j1939 CAN frame.
func (md *MessagesDescriptor) UnmarshalFrame(f can.Frame) (generated.Message, error) {
	switch f.ID {
	case md.EEC1.ID:
		var msg EEC1
		if err := msg.UnmarshalFrame(f); err != nil {
			return nil, fmt.Errorf("unmarshal j1939 frame: %w", err)
		}
		return &msg, nil
	case md.EGF1.ID:
		var msg EGF1
		if err := msg.UnmarshalFrame(f); err != nil {
			return nil, fmt.Errorf("unmarshal j1939 frame: %w", err)
		}
		return &msg, nil
	case md.CCVS1.ID:
		var msg CCVS1
		if err := msg.UnmarshalFrame(f); err != nil {
			return nil, fmt.Errorf("unmarshal j1939 frame: %w", err)
		}
		return &msg, nil
	case md.LFE1.ID:
		var msg LFE1
		if err := msg.UnmarshalFrame(f); err != nil {
			return nil, fmt.Errorf("unmarshal j1939 frame: %w", err)
		}
		return &msg, nil
	case md.VP1.ID:
		var msg VP1
		if err := msg.UnmarshalFrame(f); err != nil {
			return nil, fmt.Errorf("unmarshal j1939 frame: %w", err)
		}
		return &msg, nil
	case md.TRF1.ID:
		var msg TRF1
		if err := msg.UnmarshalFrame(f); err != nil {
			return nil, fmt.Errorf("unmarshal j1939 frame: %w", err)
		}
		return &msg, nil
	default:
		return nil, fmt.Errorf("unmarshal j1939 frame: ID not in database: %d", f.ID)
	}
}

type EEC1Descriptor struct {
	*descriptor.Message
	EngineSpeed *descriptor.Signal
}

type EGF1Descriptor struct {
	*descriptor.Message
	EngineIntakeAirMassFlowRate *descriptor.Signal
}

type CCVS1Descriptor struct {
	*descriptor.Message
	WheelBasedVehicleSpeed *descriptor.Signal
}

type LFE1Descriptor struct {
	*descriptor.Message
	EngineFuelRate *descriptor.Signal
}

type VP1Descriptor struct {
	*descriptor.Message
	Latitude  *descriptor.Signal
	Longitude *descriptor.Signal
}

type TRF1Descriptor struct {
	*descriptor.Message
	TransmissionOilTemperature1 *descriptor.Signal
}

// Database returns the j1939 database descriptor.
func (md *MessagesDescriptor) Database() *descriptor.Database {
	return d
}

var nd = &NodesDescriptor{}

var md = &MessagesDescriptor{
	EEC1: &EEC1Descriptor{
		Message:     d.Messages[0],
		EngineSpeed: d.Messages[0].Signals[0],
	},
	EGF1: &EGF1Descriptor{
		Message:                     d.Messages[1],
		EngineIntakeAirMassFlowRate: d.Messages[1].Signals[0],
	},
	CCVS1: &CCVS1Descriptor{
		Message:                d.Messages[2],
		WheelBasedVehicleSpeed: d.Messages[2].Signals[0],
	},
	LFE1: &LFE1Descriptor{
		Message:        d.Messages[3],
		EngineFuelRate: d.Messages[3].Signals[0],
	},
	VP1: &VP1Descriptor{
		Message:   d.Messages[4],
		Latitude:  d.Messages[4].Signals[0],
		Longitude: d.Messages[4].Signals[1],
	},
	TRF1: &TRF1Descriptor{
		Message:                     d.Messages[5],
		TransmissionOilTemperature1: d.Messages[5].Signals[0],
	},
}

var d = (*descriptor.Database)(&descriptor.Database{
	SourceFile: (string)("dummy/j1939/j1939.dbc"),
	Version:    (string)(""),
	Messages: ([]*descriptor.Message)([]*descriptor.Message{
		(*descriptor.Message)(&descriptor.Message{
			Name:        (string)("EEC1"),
			ID:          (uint32)(217056510),
			IsExtended:  (bool)(true),
			Length:      (uint8)(8),
			SendType:    (descriptor.SendType)(0),
			Description: (string)("Electronic Engine Controller 1"),
			Signals: ([]*descriptor.Signal)([]*descriptor.Signal{
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("EngineSpeed"),
					Start:             (uint8)(24),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(false),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.125),
					Min:               (float64)(0),
					Max:               (float64)(8031.875),
					Unit:              (string)("rpm"),
					Description:       (string)("Actual engine speed which is calculated over a minimum crankshaft angle of 720 degrees divided by the number of cylinders.…"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
			}),
			SenderNode: (string)("Vector__XXX"),
			CycleTime:  (time.Duration)(0),
			DelayTime:  (time.Duration)(0),
		}),
		(*descriptor.Message)(&descriptor.Message{
			Name:        (string)("EGF1"),
			ID:          (uint32)(217058046),
			IsExtended:  (bool)(true),
			Length:      (uint8)(8),
			SendType:    (descriptor.SendType)(0),
			Description: (string)("Engine Gas Flow Rate"),
			Signals: ([]*descriptor.Signal)([]*descriptor.Signal{
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("EngineIntakeAirMassFlowRate"),
					Start:             (uint8)(16),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(false),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.05),
					Min:               (float64)(0),
					Max:               (float64)(3212.75),
					Unit:              (string)("kg/h"),
					Description:       (string)("Engine Intake Air Mass Flow Rate: Mass flow rate of fresh air entering the engine air intake, before any EGR mixer, if used.  Flow rate of fresh air conducted to the engine cylinders to support combustion."),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
			}),
			SenderNode: (string)("Vector__XXX"),
			CycleTime:  (time.Duration)(0),
			DelayTime:  (time.Duration)(0),
		}),
		(*descriptor.Message)(&descriptor.Message{
			Name:        (string)("CCVS1"),
			ID:          (uint32)(419361278),
			IsExtended:  (bool)(true),
			Length:      (uint8)(8),
			SendType:    (descriptor.SendType)(0),
			Description: (string)("Cruise Control/Vehicle Speed 1"),
			Signals: ([]*descriptor.Signal)([]*descriptor.Signal{
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("WheelBasedVehicleSpeed"),
					Start:             (uint8)(8),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(false),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.00390625),
					Min:               (float64)(0),
					Max:               (float64)(250.996),
					Unit:              (string)("km/h"),
					Description:       (string)("Wheel-Based Vehicle Speed: Speed of the vehicle as calculated from wheel or tailshaft speed."),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
			}),
			SenderNode: (string)("Vector__XXX"),
			CycleTime:  (time.Duration)(0),
			DelayTime:  (time.Duration)(0),
		}),
		(*descriptor.Message)(&descriptor.Message{
			Name:        (string)("LFE1"),
			ID:          (uint32)(419361534),
			IsExtended:  (bool)(true),
			Length:      (uint8)(8),
			SendType:    (descriptor.SendType)(0),
			Description: (string)("Fuel Economy (Liquid)"),
			Signals: ([]*descriptor.Signal)([]*descriptor.Signal{
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("EngineFuelRate"),
					Start:             (uint8)(0),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(false),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.05),
					Min:               (float64)(0),
					Max:               (float64)(3212.75),
					Unit:              (string)("l/h"),
					Description:       (string)("Engine Fuel Rate: Amount of fuel consumed by engine per unit of time.NOTE - See SPN 1600 for alternate resolution."),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
			}),
			SenderNode: (string)("Vector__XXX"),
			CycleTime:  (time.Duration)(0),
			DelayTime:  (time.Duration)(0),
		}),
		(*descriptor.Message)(&descriptor.Message{
			Name:        (string)("VP1"),
			ID:          (uint32)(419361790),
			IsExtended:  (bool)(true),
			Length:      (uint8)(8),
			SendType:    (descriptor.SendType)(0),
			Description: (string)("Vehicle Position 1"),
			Signals: ([]*descriptor.Signal)([]*descriptor.Signal{
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("Latitude"),
					Start:             (uint8)(0),
					Length:            (uint8)(32),
					IsBigEndian:       (bool)(false),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(-210),
					Scale:             (float64)(1e-07),
					Min:               (float64)(-210),
					Max:               (float64)(211.1081215),
					Unit:              (string)("deg"),
					Description:       (string)("Latitude: Latitude position of the vehicle.Negative values are South and positive values are North."),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("Longitude"),
					Start:             (uint8)(32),
					Length:            (uint8)(32),
					IsBigEndian:       (bool)(false),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(-210),
					Scale:             (float64)(1e-07),
					Min:               (float64)(-210),
					Max:               (float64)(211.1081215),
					Unit:              (string)("deg"),
					Description:       (string)("Longitude: Longitude position of the vehicle.Negative values are West and positive value are East."),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
			}),
			SenderNode: (string)("Vector__XXX"),
			CycleTime:  (time.Duration)(0),
			DelayTime:  (time.Duration)(0),
		}),
		(*descriptor.Message)(&descriptor.Message{
			Name:        (string)("TRF1"),
			ID:          (uint32)(419363070),
			IsExtended:  (bool)(true),
			Length:      (uint8)(8),
			SendType:    (descriptor.SendType)(0),
			Description: (string)("Transmission Fluids 1"),
			Signals: ([]*descriptor.Signal)([]*descriptor.Signal{
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("TransmissionOilTemperature1"),
					Start:             (uint8)(32),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(false),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(-273),
					Scale:             (float64)(0.03125),
					Min:               (float64)(-273),
					Max:               (float64)(1734.96875),
					Unit:              (string)("degC"),
					Description:       (string)("Transmission Oil Temperature 1: First instance of transmission lubricant temperature."),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
			}),
			SenderNode: (string)("Vector__XXX"),
			CycleTime:  (time.Duration)(0),
			DelayTime:  (time.Duration)(0),
		}),
	}),
	Nodes: ([]*descriptor.Node)(nil),
})
