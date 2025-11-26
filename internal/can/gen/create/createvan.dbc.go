// Package createvancan provides primitives for encoding and decoding createvan CAN messages.
//
// Source: dummy/create/createvan.dbc
package createvancan

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
// IVI_DashboardStatusReader provides read access to a IVI_DashboardStatus message.
type IVI_DashboardStatusReader interface {
	can.FrameMarshaler
}

// IVI_DashboardStatusWriter provides write access to a IVI_DashboardStatus message.
type IVI_DashboardStatusWriter interface {
	// CopyFrom copies all values from IVI_DashboardStatus.
	CopyFrom(IVI_DashboardStatusReader) *IVI_DashboardStatus
}

type IVI_DashboardStatus struct {
}

func NewIVI_DashboardStatus() *IVI_DashboardStatus {
	m := &IVI_DashboardStatus{}
	m.Reset()
	return m
}

func (m *IVI_DashboardStatus) Reset() {
}

func (m *IVI_DashboardStatus) CopyFrom(o IVI_DashboardStatusReader) *IVI_DashboardStatus {
	f, _ := o.MarshalFrame()
	_ = m.UnmarshalFrame(f)
	return m
}

// Descriptor returns the IVI_DashboardStatus descriptor.
func (m *IVI_DashboardStatus) Descriptor() *descriptor.Message {
	return Messages().IVI_DashboardStatus.Message
}

// String returns a compact string representation of the message.
func (m *IVI_DashboardStatus) String() string {
	return cantext.MessageString(m)
}

// Frame returns a CAN frame representing the message.
func (m *IVI_DashboardStatus) Frame() can.Frame {
	md := Messages().IVI_DashboardStatus
	f := can.Frame{ID: md.ID, IsExtended: md.IsExtended, Length: md.Length}
	return f
}

// MarshalFrame encodes the message as a CAN frame.
func (m *IVI_DashboardStatus) MarshalFrame() (can.Frame, error) {
	return m.Frame(), nil
}

// UnmarshalFrame decodes the message from a CAN frame.
func (m *IVI_DashboardStatus) UnmarshalFrame(f can.Frame) error {
	md := Messages().IVI_DashboardStatus
	switch {
	case f.ID != md.ID:
		return fmt.Errorf(
			"unmarshal IVI_DashboardStatus: expects ID 217056510 (got %s with ID %d)", f.String(), f.ID,
		)
	case f.Length != md.Length:
		return fmt.Errorf(
			"unmarshal IVI_DashboardStatus: expects length 8 (got %s with length %d)", f.String(), f.Length,
		)
	case f.IsRemote:
		return fmt.Errorf(
			"unmarshal IVI_DashboardStatus: expects non-remote frame (got remote frame %s)", f.String(),
		)
	case f.IsExtended != md.IsExtended:
		return fmt.Errorf(
			"unmarshal IVI_DashboardStatus: expects extended ID (got %s with standard ID)", f.String(),
		)
	}
	return nil
}

// IVI_CornerMotorStatusReader provides read access to a IVI_CornerMotorStatus message.
type IVI_CornerMotorStatusReader interface {
	can.FrameMarshaler
}

// IVI_CornerMotorStatusWriter provides write access to a IVI_CornerMotorStatus message.
type IVI_CornerMotorStatusWriter interface {
	// CopyFrom copies all values from IVI_CornerMotorStatus.
	CopyFrom(IVI_CornerMotorStatusReader) *IVI_CornerMotorStatus
}

type IVI_CornerMotorStatus struct {
}

func NewIVI_CornerMotorStatus() *IVI_CornerMotorStatus {
	m := &IVI_CornerMotorStatus{}
	m.Reset()
	return m
}

func (m *IVI_CornerMotorStatus) Reset() {
}

func (m *IVI_CornerMotorStatus) CopyFrom(o IVI_CornerMotorStatusReader) *IVI_CornerMotorStatus {
	f, _ := o.MarshalFrame()
	_ = m.UnmarshalFrame(f)
	return m
}

// Descriptor returns the IVI_CornerMotorStatus descriptor.
func (m *IVI_CornerMotorStatus) Descriptor() *descriptor.Message {
	return Messages().IVI_CornerMotorStatus.Message
}

// String returns a compact string representation of the message.
func (m *IVI_CornerMotorStatus) String() string {
	return cantext.MessageString(m)
}

// Frame returns a CAN frame representing the message.
func (m *IVI_CornerMotorStatus) Frame() can.Frame {
	md := Messages().IVI_CornerMotorStatus
	f := can.Frame{ID: md.ID, IsExtended: md.IsExtended, Length: md.Length}
	return f
}

// MarshalFrame encodes the message as a CAN frame.
func (m *IVI_CornerMotorStatus) MarshalFrame() (can.Frame, error) {
	return m.Frame(), nil
}

// UnmarshalFrame decodes the message from a CAN frame.
func (m *IVI_CornerMotorStatus) UnmarshalFrame(f can.Frame) error {
	md := Messages().IVI_CornerMotorStatus
	switch {
	case f.ID != md.ID:
		return fmt.Errorf(
			"unmarshal IVI_CornerMotorStatus: expects ID 217056766 (got %s with ID %d)", f.String(), f.ID,
		)
	case f.Length != md.Length:
		return fmt.Errorf(
			"unmarshal IVI_CornerMotorStatus: expects length 8 (got %s with length %d)", f.String(), f.Length,
		)
	case f.IsRemote:
		return fmt.Errorf(
			"unmarshal IVI_CornerMotorStatus: expects non-remote frame (got remote frame %s)", f.String(),
		)
	case f.IsExtended != md.IsExtended:
		return fmt.Errorf(
			"unmarshal IVI_CornerMotorStatus: expects extended ID (got %s with standard ID)", f.String(),
		)
	}
	return nil
}

// IOV_FrontLeftCornerMotor1Reader provides read access to a IOV_FrontLeftCornerMotor1 message.
type IOV_FrontLeftCornerMotor1Reader interface {
	can.FrameMarshaler
	// VESC_StatusERPM1 returns the value of the VESC_StatusERPM1 signal.
	VESC_StatusERPM1() int32
	// VESC_StatusCurrent1 returns the value of the VESC_StatusCurrent1 signal.
	VESC_StatusCurrent1() int16
	// VESC_StatusDutyCycle1 returns the value of the VESC_StatusDutyCycle1 signal.
	VESC_StatusDutyCycle1() int16
}

// IOV_FrontLeftCornerMotor1Writer provides write access to a IOV_FrontLeftCornerMotor1 message.
type IOV_FrontLeftCornerMotor1Writer interface {
	// CopyFrom copies all values from IOV_FrontLeftCornerMotor1.
	CopyFrom(IOV_FrontLeftCornerMotor1Reader) *IOV_FrontLeftCornerMotor1
	// SetVESC_StatusERPM1 sets the value of the VESC_StatusERPM1 signal.
	SetVESC_StatusERPM1(int32) *IOV_FrontLeftCornerMotor1
	// SetVESC_StatusCurrent1 sets the value of the VESC_StatusCurrent1 signal.
	SetVESC_StatusCurrent1(int16) *IOV_FrontLeftCornerMotor1
	// SetVESC_StatusDutyCycle1 sets the value of the VESC_StatusDutyCycle1 signal.
	SetVESC_StatusDutyCycle1(int16) *IOV_FrontLeftCornerMotor1
}

type IOV_FrontLeftCornerMotor1 struct {
	xxx_VESC_StatusERPM1      int32
	xxx_VESC_StatusCurrent1   int16
	xxx_VESC_StatusDutyCycle1 int16
}

func NewIOV_FrontLeftCornerMotor1() *IOV_FrontLeftCornerMotor1 {
	m := &IOV_FrontLeftCornerMotor1{}
	m.Reset()
	return m
}

func (m *IOV_FrontLeftCornerMotor1) Reset() {
	m.xxx_VESC_StatusERPM1 = 0
	m.xxx_VESC_StatusCurrent1 = 0
	m.xxx_VESC_StatusDutyCycle1 = 0
}

func (m *IOV_FrontLeftCornerMotor1) CopyFrom(o IOV_FrontLeftCornerMotor1Reader) *IOV_FrontLeftCornerMotor1 {
	f, _ := o.MarshalFrame()
	_ = m.UnmarshalFrame(f)
	return m
}

// Descriptor returns the IOV_FrontLeftCornerMotor1 descriptor.
func (m *IOV_FrontLeftCornerMotor1) Descriptor() *descriptor.Message {
	return Messages().IOV_FrontLeftCornerMotor1.Message
}

// String returns a compact string representation of the message.
func (m *IOV_FrontLeftCornerMotor1) String() string {
	return cantext.MessageString(m)
}

func (m *IOV_FrontLeftCornerMotor1) VESC_StatusERPM1() int32 {
	return m.xxx_VESC_StatusERPM1
}

func (m *IOV_FrontLeftCornerMotor1) SetVESC_StatusERPM1(v int32) *IOV_FrontLeftCornerMotor1 {
	m.xxx_VESC_StatusERPM1 = int32(Messages().IOV_FrontLeftCornerMotor1.VESC_StatusERPM1.SaturatedCastSigned(int64(v)))
	return m
}

func (m *IOV_FrontLeftCornerMotor1) VESC_StatusCurrent1() int16 {
	return m.xxx_VESC_StatusCurrent1
}

func (m *IOV_FrontLeftCornerMotor1) SetVESC_StatusCurrent1(v int16) *IOV_FrontLeftCornerMotor1 {
	m.xxx_VESC_StatusCurrent1 = int16(Messages().IOV_FrontLeftCornerMotor1.VESC_StatusCurrent1.SaturatedCastSigned(int64(v)))
	return m
}

func (m *IOV_FrontLeftCornerMotor1) VESC_StatusDutyCycle1() int16 {
	return m.xxx_VESC_StatusDutyCycle1
}

func (m *IOV_FrontLeftCornerMotor1) SetVESC_StatusDutyCycle1(v int16) *IOV_FrontLeftCornerMotor1 {
	m.xxx_VESC_StatusDutyCycle1 = int16(Messages().IOV_FrontLeftCornerMotor1.VESC_StatusDutyCycle1.SaturatedCastSigned(int64(v)))
	return m
}

// Frame returns a CAN frame representing the message.
func (m *IOV_FrontLeftCornerMotor1) Frame() can.Frame {
	md := Messages().IOV_FrontLeftCornerMotor1
	f := can.Frame{ID: md.ID, IsExtended: md.IsExtended, Length: md.Length}
	md.VESC_StatusERPM1.MarshalSigned(&f.Data, int64(m.xxx_VESC_StatusERPM1))
	md.VESC_StatusCurrent1.MarshalSigned(&f.Data, int64(m.xxx_VESC_StatusCurrent1))
	md.VESC_StatusDutyCycle1.MarshalSigned(&f.Data, int64(m.xxx_VESC_StatusDutyCycle1))
	return f
}

// MarshalFrame encodes the message as a CAN frame.
func (m *IOV_FrontLeftCornerMotor1) MarshalFrame() (can.Frame, error) {
	return m.Frame(), nil
}

// UnmarshalFrame decodes the message from a CAN frame.
func (m *IOV_FrontLeftCornerMotor1) UnmarshalFrame(f can.Frame) error {
	md := Messages().IOV_FrontLeftCornerMotor1
	switch {
	case f.ID != md.ID:
		return fmt.Errorf(
			"unmarshal IOV_FrontLeftCornerMotor1: expects ID 217063422 (got %s with ID %d)", f.String(), f.ID,
		)
	case f.Length != md.Length:
		return fmt.Errorf(
			"unmarshal IOV_FrontLeftCornerMotor1: expects length 8 (got %s with length %d)", f.String(), f.Length,
		)
	case f.IsRemote:
		return fmt.Errorf(
			"unmarshal IOV_FrontLeftCornerMotor1: expects non-remote frame (got remote frame %s)", f.String(),
		)
	case f.IsExtended != md.IsExtended:
		return fmt.Errorf(
			"unmarshal IOV_FrontLeftCornerMotor1: expects extended ID (got %s with standard ID)", f.String(),
		)
	}
	m.xxx_VESC_StatusERPM1 = int32(md.VESC_StatusERPM1.UnmarshalSigned(f.Data))
	m.xxx_VESC_StatusCurrent1 = int16(md.VESC_StatusCurrent1.UnmarshalSigned(f.Data))
	m.xxx_VESC_StatusDutyCycle1 = int16(md.VESC_StatusDutyCycle1.UnmarshalSigned(f.Data))
	return nil
}

// IOV_FrontLeftCornerMotor2Reader provides read access to a IOV_FrontLeftCornerMotor2 message.
type IOV_FrontLeftCornerMotor2Reader interface {
	can.FrameMarshaler
	// VESC_StatusAmpHours1 returns the value of the VESC_StatusAmpHours1 signal.
	VESC_StatusAmpHours1() int32
	// VESC_StatusAmpHoursChg1 returns the value of the VESC_StatusAmpHoursChg1 signal.
	VESC_StatusAmpHoursChg1() int32
}

// IOV_FrontLeftCornerMotor2Writer provides write access to a IOV_FrontLeftCornerMotor2 message.
type IOV_FrontLeftCornerMotor2Writer interface {
	// CopyFrom copies all values from IOV_FrontLeftCornerMotor2.
	CopyFrom(IOV_FrontLeftCornerMotor2Reader) *IOV_FrontLeftCornerMotor2
	// SetVESC_StatusAmpHours1 sets the value of the VESC_StatusAmpHours1 signal.
	SetVESC_StatusAmpHours1(int32) *IOV_FrontLeftCornerMotor2
	// SetVESC_StatusAmpHoursChg1 sets the value of the VESC_StatusAmpHoursChg1 signal.
	SetVESC_StatusAmpHoursChg1(int32) *IOV_FrontLeftCornerMotor2
}

type IOV_FrontLeftCornerMotor2 struct {
	xxx_VESC_StatusAmpHours1    int32
	xxx_VESC_StatusAmpHoursChg1 int32
}

func NewIOV_FrontLeftCornerMotor2() *IOV_FrontLeftCornerMotor2 {
	m := &IOV_FrontLeftCornerMotor2{}
	m.Reset()
	return m
}

func (m *IOV_FrontLeftCornerMotor2) Reset() {
	m.xxx_VESC_StatusAmpHours1 = 0
	m.xxx_VESC_StatusAmpHoursChg1 = 0
}

func (m *IOV_FrontLeftCornerMotor2) CopyFrom(o IOV_FrontLeftCornerMotor2Reader) *IOV_FrontLeftCornerMotor2 {
	f, _ := o.MarshalFrame()
	_ = m.UnmarshalFrame(f)
	return m
}

// Descriptor returns the IOV_FrontLeftCornerMotor2 descriptor.
func (m *IOV_FrontLeftCornerMotor2) Descriptor() *descriptor.Message {
	return Messages().IOV_FrontLeftCornerMotor2.Message
}

// String returns a compact string representation of the message.
func (m *IOV_FrontLeftCornerMotor2) String() string {
	return cantext.MessageString(m)
}

func (m *IOV_FrontLeftCornerMotor2) VESC_StatusAmpHours1() int32 {
	return m.xxx_VESC_StatusAmpHours1
}

func (m *IOV_FrontLeftCornerMotor2) SetVESC_StatusAmpHours1(v int32) *IOV_FrontLeftCornerMotor2 {
	m.xxx_VESC_StatusAmpHours1 = int32(Messages().IOV_FrontLeftCornerMotor2.VESC_StatusAmpHours1.SaturatedCastSigned(int64(v)))
	return m
}

func (m *IOV_FrontLeftCornerMotor2) VESC_StatusAmpHoursChg1() int32 {
	return m.xxx_VESC_StatusAmpHoursChg1
}

func (m *IOV_FrontLeftCornerMotor2) SetVESC_StatusAmpHoursChg1(v int32) *IOV_FrontLeftCornerMotor2 {
	m.xxx_VESC_StatusAmpHoursChg1 = int32(Messages().IOV_FrontLeftCornerMotor2.VESC_StatusAmpHoursChg1.SaturatedCastSigned(int64(v)))
	return m
}

// Frame returns a CAN frame representing the message.
func (m *IOV_FrontLeftCornerMotor2) Frame() can.Frame {
	md := Messages().IOV_FrontLeftCornerMotor2
	f := can.Frame{ID: md.ID, IsExtended: md.IsExtended, Length: md.Length}
	md.VESC_StatusAmpHours1.MarshalSigned(&f.Data, int64(m.xxx_VESC_StatusAmpHours1))
	md.VESC_StatusAmpHoursChg1.MarshalSigned(&f.Data, int64(m.xxx_VESC_StatusAmpHoursChg1))
	return f
}

// MarshalFrame encodes the message as a CAN frame.
func (m *IOV_FrontLeftCornerMotor2) MarshalFrame() (can.Frame, error) {
	return m.Frame(), nil
}

// UnmarshalFrame decodes the message from a CAN frame.
func (m *IOV_FrontLeftCornerMotor2) UnmarshalFrame(f can.Frame) error {
	md := Messages().IOV_FrontLeftCornerMotor2
	switch {
	case f.ID != md.ID:
		return fmt.Errorf(
			"unmarshal IOV_FrontLeftCornerMotor2: expects ID 217067518 (got %s with ID %d)", f.String(), f.ID,
		)
	case f.Length != md.Length:
		return fmt.Errorf(
			"unmarshal IOV_FrontLeftCornerMotor2: expects length 8 (got %s with length %d)", f.String(), f.Length,
		)
	case f.IsRemote:
		return fmt.Errorf(
			"unmarshal IOV_FrontLeftCornerMotor2: expects non-remote frame (got remote frame %s)", f.String(),
		)
	case f.IsExtended != md.IsExtended:
		return fmt.Errorf(
			"unmarshal IOV_FrontLeftCornerMotor2: expects extended ID (got %s with standard ID)", f.String(),
		)
	}
	m.xxx_VESC_StatusAmpHours1 = int32(md.VESC_StatusAmpHours1.UnmarshalSigned(f.Data))
	m.xxx_VESC_StatusAmpHoursChg1 = int32(md.VESC_StatusAmpHoursChg1.UnmarshalSigned(f.Data))
	return nil
}

// IOV_FrontLeftCornerMotor3Reader provides read access to a IOV_FrontLeftCornerMotor3 message.
type IOV_FrontLeftCornerMotor3Reader interface {
	can.FrameMarshaler
	// VESC_StatusWattHours1 returns the value of the VESC_StatusWattHours1 signal.
	VESC_StatusWattHours1() int32
	// VESC_StatusWattHoursChg1 returns the value of the VESC_StatusWattHoursChg1 signal.
	VESC_StatusWattHoursChg1() int32
}

// IOV_FrontLeftCornerMotor3Writer provides write access to a IOV_FrontLeftCornerMotor3 message.
type IOV_FrontLeftCornerMotor3Writer interface {
	// CopyFrom copies all values from IOV_FrontLeftCornerMotor3.
	CopyFrom(IOV_FrontLeftCornerMotor3Reader) *IOV_FrontLeftCornerMotor3
	// SetVESC_StatusWattHours1 sets the value of the VESC_StatusWattHours1 signal.
	SetVESC_StatusWattHours1(int32) *IOV_FrontLeftCornerMotor3
	// SetVESC_StatusWattHoursChg1 sets the value of the VESC_StatusWattHoursChg1 signal.
	SetVESC_StatusWattHoursChg1(int32) *IOV_FrontLeftCornerMotor3
}

type IOV_FrontLeftCornerMotor3 struct {
	xxx_VESC_StatusWattHours1    int32
	xxx_VESC_StatusWattHoursChg1 int32
}

func NewIOV_FrontLeftCornerMotor3() *IOV_FrontLeftCornerMotor3 {
	m := &IOV_FrontLeftCornerMotor3{}
	m.Reset()
	return m
}

func (m *IOV_FrontLeftCornerMotor3) Reset() {
	m.xxx_VESC_StatusWattHours1 = 0
	m.xxx_VESC_StatusWattHoursChg1 = 0
}

func (m *IOV_FrontLeftCornerMotor3) CopyFrom(o IOV_FrontLeftCornerMotor3Reader) *IOV_FrontLeftCornerMotor3 {
	f, _ := o.MarshalFrame()
	_ = m.UnmarshalFrame(f)
	return m
}

// Descriptor returns the IOV_FrontLeftCornerMotor3 descriptor.
func (m *IOV_FrontLeftCornerMotor3) Descriptor() *descriptor.Message {
	return Messages().IOV_FrontLeftCornerMotor3.Message
}

// String returns a compact string representation of the message.
func (m *IOV_FrontLeftCornerMotor3) String() string {
	return cantext.MessageString(m)
}

func (m *IOV_FrontLeftCornerMotor3) VESC_StatusWattHours1() int32 {
	return m.xxx_VESC_StatusWattHours1
}

func (m *IOV_FrontLeftCornerMotor3) SetVESC_StatusWattHours1(v int32) *IOV_FrontLeftCornerMotor3 {
	m.xxx_VESC_StatusWattHours1 = int32(Messages().IOV_FrontLeftCornerMotor3.VESC_StatusWattHours1.SaturatedCastSigned(int64(v)))
	return m
}

func (m *IOV_FrontLeftCornerMotor3) VESC_StatusWattHoursChg1() int32 {
	return m.xxx_VESC_StatusWattHoursChg1
}

func (m *IOV_FrontLeftCornerMotor3) SetVESC_StatusWattHoursChg1(v int32) *IOV_FrontLeftCornerMotor3 {
	m.xxx_VESC_StatusWattHoursChg1 = int32(Messages().IOV_FrontLeftCornerMotor3.VESC_StatusWattHoursChg1.SaturatedCastSigned(int64(v)))
	return m
}

// Frame returns a CAN frame representing the message.
func (m *IOV_FrontLeftCornerMotor3) Frame() can.Frame {
	md := Messages().IOV_FrontLeftCornerMotor3
	f := can.Frame{ID: md.ID, IsExtended: md.IsExtended, Length: md.Length}
	md.VESC_StatusWattHours1.MarshalSigned(&f.Data, int64(m.xxx_VESC_StatusWattHours1))
	md.VESC_StatusWattHoursChg1.MarshalSigned(&f.Data, int64(m.xxx_VESC_StatusWattHoursChg1))
	return f
}

// MarshalFrame encodes the message as a CAN frame.
func (m *IOV_FrontLeftCornerMotor3) MarshalFrame() (can.Frame, error) {
	return m.Frame(), nil
}

// UnmarshalFrame decodes the message from a CAN frame.
func (m *IOV_FrontLeftCornerMotor3) UnmarshalFrame(f can.Frame) error {
	md := Messages().IOV_FrontLeftCornerMotor3
	switch {
	case f.ID != md.ID:
		return fmt.Errorf(
			"unmarshal IOV_FrontLeftCornerMotor3: expects ID 217071614 (got %s with ID %d)", f.String(), f.ID,
		)
	case f.Length != md.Length:
		return fmt.Errorf(
			"unmarshal IOV_FrontLeftCornerMotor3: expects length 8 (got %s with length %d)", f.String(), f.Length,
		)
	case f.IsRemote:
		return fmt.Errorf(
			"unmarshal IOV_FrontLeftCornerMotor3: expects non-remote frame (got remote frame %s)", f.String(),
		)
	case f.IsExtended != md.IsExtended:
		return fmt.Errorf(
			"unmarshal IOV_FrontLeftCornerMotor3: expects extended ID (got %s with standard ID)", f.String(),
		)
	}
	m.xxx_VESC_StatusWattHours1 = int32(md.VESC_StatusWattHours1.UnmarshalSigned(f.Data))
	m.xxx_VESC_StatusWattHoursChg1 = int32(md.VESC_StatusWattHoursChg1.UnmarshalSigned(f.Data))
	return nil
}

// IOV_FrontLeftCornerMotor4Reader provides read access to a IOV_FrontLeftCornerMotor4 message.
type IOV_FrontLeftCornerMotor4Reader interface {
	can.FrameMarshaler
	// VESC_StatusTemFET1 returns the value of the VESC_StatusTemFET1 signal.
	VESC_StatusTemFET1() int16
	// VESC_StatusTempMotor1 returns the value of the VESC_StatusTempMotor1 signal.
	VESC_StatusTempMotor1() int16
	// VESC_StatusCurrentIn1 returns the value of the VESC_StatusCurrentIn1 signal.
	VESC_StatusCurrentIn1() int16
	// VESC_StatusPIDPos1 returns the value of the VESC_StatusPIDPos1 signal.
	VESC_StatusPIDPos1() int16
}

// IOV_FrontLeftCornerMotor4Writer provides write access to a IOV_FrontLeftCornerMotor4 message.
type IOV_FrontLeftCornerMotor4Writer interface {
	// CopyFrom copies all values from IOV_FrontLeftCornerMotor4.
	CopyFrom(IOV_FrontLeftCornerMotor4Reader) *IOV_FrontLeftCornerMotor4
	// SetVESC_StatusTemFET1 sets the value of the VESC_StatusTemFET1 signal.
	SetVESC_StatusTemFET1(int16) *IOV_FrontLeftCornerMotor4
	// SetVESC_StatusTempMotor1 sets the value of the VESC_StatusTempMotor1 signal.
	SetVESC_StatusTempMotor1(int16) *IOV_FrontLeftCornerMotor4
	// SetVESC_StatusCurrentIn1 sets the value of the VESC_StatusCurrentIn1 signal.
	SetVESC_StatusCurrentIn1(int16) *IOV_FrontLeftCornerMotor4
	// SetVESC_StatusPIDPos1 sets the value of the VESC_StatusPIDPos1 signal.
	SetVESC_StatusPIDPos1(int16) *IOV_FrontLeftCornerMotor4
}

type IOV_FrontLeftCornerMotor4 struct {
	xxx_VESC_StatusTemFET1    int16
	xxx_VESC_StatusTempMotor1 int16
	xxx_VESC_StatusCurrentIn1 int16
	xxx_VESC_StatusPIDPos1    int16
}

func NewIOV_FrontLeftCornerMotor4() *IOV_FrontLeftCornerMotor4 {
	m := &IOV_FrontLeftCornerMotor4{}
	m.Reset()
	return m
}

func (m *IOV_FrontLeftCornerMotor4) Reset() {
	m.xxx_VESC_StatusTemFET1 = 0
	m.xxx_VESC_StatusTempMotor1 = 0
	m.xxx_VESC_StatusCurrentIn1 = 0
	m.xxx_VESC_StatusPIDPos1 = 0
}

func (m *IOV_FrontLeftCornerMotor4) CopyFrom(o IOV_FrontLeftCornerMotor4Reader) *IOV_FrontLeftCornerMotor4 {
	f, _ := o.MarshalFrame()
	_ = m.UnmarshalFrame(f)
	return m
}

// Descriptor returns the IOV_FrontLeftCornerMotor4 descriptor.
func (m *IOV_FrontLeftCornerMotor4) Descriptor() *descriptor.Message {
	return Messages().IOV_FrontLeftCornerMotor4.Message
}

// String returns a compact string representation of the message.
func (m *IOV_FrontLeftCornerMotor4) String() string {
	return cantext.MessageString(m)
}

func (m *IOV_FrontLeftCornerMotor4) VESC_StatusTemFET1() int16 {
	return m.xxx_VESC_StatusTemFET1
}

func (m *IOV_FrontLeftCornerMotor4) SetVESC_StatusTemFET1(v int16) *IOV_FrontLeftCornerMotor4 {
	m.xxx_VESC_StatusTemFET1 = int16(Messages().IOV_FrontLeftCornerMotor4.VESC_StatusTemFET1.SaturatedCastSigned(int64(v)))
	return m
}

func (m *IOV_FrontLeftCornerMotor4) VESC_StatusTempMotor1() int16 {
	return m.xxx_VESC_StatusTempMotor1
}

func (m *IOV_FrontLeftCornerMotor4) SetVESC_StatusTempMotor1(v int16) *IOV_FrontLeftCornerMotor4 {
	m.xxx_VESC_StatusTempMotor1 = int16(Messages().IOV_FrontLeftCornerMotor4.VESC_StatusTempMotor1.SaturatedCastSigned(int64(v)))
	return m
}

func (m *IOV_FrontLeftCornerMotor4) VESC_StatusCurrentIn1() int16 {
	return m.xxx_VESC_StatusCurrentIn1
}

func (m *IOV_FrontLeftCornerMotor4) SetVESC_StatusCurrentIn1(v int16) *IOV_FrontLeftCornerMotor4 {
	m.xxx_VESC_StatusCurrentIn1 = int16(Messages().IOV_FrontLeftCornerMotor4.VESC_StatusCurrentIn1.SaturatedCastSigned(int64(v)))
	return m
}

func (m *IOV_FrontLeftCornerMotor4) VESC_StatusPIDPos1() int16 {
	return m.xxx_VESC_StatusPIDPos1
}

func (m *IOV_FrontLeftCornerMotor4) SetVESC_StatusPIDPos1(v int16) *IOV_FrontLeftCornerMotor4 {
	m.xxx_VESC_StatusPIDPos1 = int16(Messages().IOV_FrontLeftCornerMotor4.VESC_StatusPIDPos1.SaturatedCastSigned(int64(v)))
	return m
}

// Frame returns a CAN frame representing the message.
func (m *IOV_FrontLeftCornerMotor4) Frame() can.Frame {
	md := Messages().IOV_FrontLeftCornerMotor4
	f := can.Frame{ID: md.ID, IsExtended: md.IsExtended, Length: md.Length}
	md.VESC_StatusTemFET1.MarshalSigned(&f.Data, int64(m.xxx_VESC_StatusTemFET1))
	md.VESC_StatusTempMotor1.MarshalSigned(&f.Data, int64(m.xxx_VESC_StatusTempMotor1))
	md.VESC_StatusCurrentIn1.MarshalSigned(&f.Data, int64(m.xxx_VESC_StatusCurrentIn1))
	md.VESC_StatusPIDPos1.MarshalSigned(&f.Data, int64(m.xxx_VESC_StatusPIDPos1))
	return f
}

// MarshalFrame encodes the message as a CAN frame.
func (m *IOV_FrontLeftCornerMotor4) MarshalFrame() (can.Frame, error) {
	return m.Frame(), nil
}

// UnmarshalFrame decodes the message from a CAN frame.
func (m *IOV_FrontLeftCornerMotor4) UnmarshalFrame(f can.Frame) error {
	md := Messages().IOV_FrontLeftCornerMotor4
	switch {
	case f.ID != md.ID:
		return fmt.Errorf(
			"unmarshal IOV_FrontLeftCornerMotor4: expects ID 217075710 (got %s with ID %d)", f.String(), f.ID,
		)
	case f.Length != md.Length:
		return fmt.Errorf(
			"unmarshal IOV_FrontLeftCornerMotor4: expects length 8 (got %s with length %d)", f.String(), f.Length,
		)
	case f.IsRemote:
		return fmt.Errorf(
			"unmarshal IOV_FrontLeftCornerMotor4: expects non-remote frame (got remote frame %s)", f.String(),
		)
	case f.IsExtended != md.IsExtended:
		return fmt.Errorf(
			"unmarshal IOV_FrontLeftCornerMotor4: expects extended ID (got %s with standard ID)", f.String(),
		)
	}
	m.xxx_VESC_StatusTemFET1 = int16(md.VESC_StatusTemFET1.UnmarshalSigned(f.Data))
	m.xxx_VESC_StatusTempMotor1 = int16(md.VESC_StatusTempMotor1.UnmarshalSigned(f.Data))
	m.xxx_VESC_StatusCurrentIn1 = int16(md.VESC_StatusCurrentIn1.UnmarshalSigned(f.Data))
	m.xxx_VESC_StatusPIDPos1 = int16(md.VESC_StatusPIDPos1.UnmarshalSigned(f.Data))
	return nil
}

// IOV_FrontRightCornerMotor1Reader provides read access to a IOV_FrontRightCornerMotor1 message.
type IOV_FrontRightCornerMotor1Reader interface {
	can.FrameMarshaler
	// VESC_StatusERPM2 returns the value of the VESC_StatusERPM2 signal.
	VESC_StatusERPM2() int32
	// VESC_StatusCurrent2 returns the value of the VESC_StatusCurrent2 signal.
	VESC_StatusCurrent2() int16
	// VESC_StatusDutyCycle2 returns the value of the VESC_StatusDutyCycle2 signal.
	VESC_StatusDutyCycle2() int16
}

// IOV_FrontRightCornerMotor1Writer provides write access to a IOV_FrontRightCornerMotor1 message.
type IOV_FrontRightCornerMotor1Writer interface {
	// CopyFrom copies all values from IOV_FrontRightCornerMotor1.
	CopyFrom(IOV_FrontRightCornerMotor1Reader) *IOV_FrontRightCornerMotor1
	// SetVESC_StatusERPM2 sets the value of the VESC_StatusERPM2 signal.
	SetVESC_StatusERPM2(int32) *IOV_FrontRightCornerMotor1
	// SetVESC_StatusCurrent2 sets the value of the VESC_StatusCurrent2 signal.
	SetVESC_StatusCurrent2(int16) *IOV_FrontRightCornerMotor1
	// SetVESC_StatusDutyCycle2 sets the value of the VESC_StatusDutyCycle2 signal.
	SetVESC_StatusDutyCycle2(int16) *IOV_FrontRightCornerMotor1
}

type IOV_FrontRightCornerMotor1 struct {
	xxx_VESC_StatusERPM2      int32
	xxx_VESC_StatusCurrent2   int16
	xxx_VESC_StatusDutyCycle2 int16
}

func NewIOV_FrontRightCornerMotor1() *IOV_FrontRightCornerMotor1 {
	m := &IOV_FrontRightCornerMotor1{}
	m.Reset()
	return m
}

func (m *IOV_FrontRightCornerMotor1) Reset() {
	m.xxx_VESC_StatusERPM2 = 0
	m.xxx_VESC_StatusCurrent2 = 0
	m.xxx_VESC_StatusDutyCycle2 = 0
}

func (m *IOV_FrontRightCornerMotor1) CopyFrom(o IOV_FrontRightCornerMotor1Reader) *IOV_FrontRightCornerMotor1 {
	f, _ := o.MarshalFrame()
	_ = m.UnmarshalFrame(f)
	return m
}

// Descriptor returns the IOV_FrontRightCornerMotor1 descriptor.
func (m *IOV_FrontRightCornerMotor1) Descriptor() *descriptor.Message {
	return Messages().IOV_FrontRightCornerMotor1.Message
}

// String returns a compact string representation of the message.
func (m *IOV_FrontRightCornerMotor1) String() string {
	return cantext.MessageString(m)
}

func (m *IOV_FrontRightCornerMotor1) VESC_StatusERPM2() int32 {
	return m.xxx_VESC_StatusERPM2
}

func (m *IOV_FrontRightCornerMotor1) SetVESC_StatusERPM2(v int32) *IOV_FrontRightCornerMotor1 {
	m.xxx_VESC_StatusERPM2 = int32(Messages().IOV_FrontRightCornerMotor1.VESC_StatusERPM2.SaturatedCastSigned(int64(v)))
	return m
}

func (m *IOV_FrontRightCornerMotor1) VESC_StatusCurrent2() int16 {
	return m.xxx_VESC_StatusCurrent2
}

func (m *IOV_FrontRightCornerMotor1) SetVESC_StatusCurrent2(v int16) *IOV_FrontRightCornerMotor1 {
	m.xxx_VESC_StatusCurrent2 = int16(Messages().IOV_FrontRightCornerMotor1.VESC_StatusCurrent2.SaturatedCastSigned(int64(v)))
	return m
}

func (m *IOV_FrontRightCornerMotor1) VESC_StatusDutyCycle2() int16 {
	return m.xxx_VESC_StatusDutyCycle2
}

func (m *IOV_FrontRightCornerMotor1) SetVESC_StatusDutyCycle2(v int16) *IOV_FrontRightCornerMotor1 {
	m.xxx_VESC_StatusDutyCycle2 = int16(Messages().IOV_FrontRightCornerMotor1.VESC_StatusDutyCycle2.SaturatedCastSigned(int64(v)))
	return m
}

// Frame returns a CAN frame representing the message.
func (m *IOV_FrontRightCornerMotor1) Frame() can.Frame {
	md := Messages().IOV_FrontRightCornerMotor1
	f := can.Frame{ID: md.ID, IsExtended: md.IsExtended, Length: md.Length}
	md.VESC_StatusERPM2.MarshalSigned(&f.Data, int64(m.xxx_VESC_StatusERPM2))
	md.VESC_StatusCurrent2.MarshalSigned(&f.Data, int64(m.xxx_VESC_StatusCurrent2))
	md.VESC_StatusDutyCycle2.MarshalSigned(&f.Data, int64(m.xxx_VESC_StatusDutyCycle2))
	return f
}

// MarshalFrame encodes the message as a CAN frame.
func (m *IOV_FrontRightCornerMotor1) MarshalFrame() (can.Frame, error) {
	return m.Frame(), nil
}

// UnmarshalFrame decodes the message from a CAN frame.
func (m *IOV_FrontRightCornerMotor1) UnmarshalFrame(f can.Frame) error {
	md := Messages().IOV_FrontRightCornerMotor1
	switch {
	case f.ID != md.ID:
		return fmt.Errorf(
			"unmarshal IOV_FrontRightCornerMotor1: expects ID 217079806 (got %s with ID %d)", f.String(), f.ID,
		)
	case f.Length != md.Length:
		return fmt.Errorf(
			"unmarshal IOV_FrontRightCornerMotor1: expects length 8 (got %s with length %d)", f.String(), f.Length,
		)
	case f.IsRemote:
		return fmt.Errorf(
			"unmarshal IOV_FrontRightCornerMotor1: expects non-remote frame (got remote frame %s)", f.String(),
		)
	case f.IsExtended != md.IsExtended:
		return fmt.Errorf(
			"unmarshal IOV_FrontRightCornerMotor1: expects extended ID (got %s with standard ID)", f.String(),
		)
	}
	m.xxx_VESC_StatusERPM2 = int32(md.VESC_StatusERPM2.UnmarshalSigned(f.Data))
	m.xxx_VESC_StatusCurrent2 = int16(md.VESC_StatusCurrent2.UnmarshalSigned(f.Data))
	m.xxx_VESC_StatusDutyCycle2 = int16(md.VESC_StatusDutyCycle2.UnmarshalSigned(f.Data))
	return nil
}

// IOV_FrontRightCornerMotor2Reader provides read access to a IOV_FrontRightCornerMotor2 message.
type IOV_FrontRightCornerMotor2Reader interface {
	can.FrameMarshaler
	// VESC_StatusAmpHours2 returns the value of the VESC_StatusAmpHours2 signal.
	VESC_StatusAmpHours2() int32
	// VESC_StatusAmpHoursChg2 returns the value of the VESC_StatusAmpHoursChg2 signal.
	VESC_StatusAmpHoursChg2() int32
}

// IOV_FrontRightCornerMotor2Writer provides write access to a IOV_FrontRightCornerMotor2 message.
type IOV_FrontRightCornerMotor2Writer interface {
	// CopyFrom copies all values from IOV_FrontRightCornerMotor2.
	CopyFrom(IOV_FrontRightCornerMotor2Reader) *IOV_FrontRightCornerMotor2
	// SetVESC_StatusAmpHours2 sets the value of the VESC_StatusAmpHours2 signal.
	SetVESC_StatusAmpHours2(int32) *IOV_FrontRightCornerMotor2
	// SetVESC_StatusAmpHoursChg2 sets the value of the VESC_StatusAmpHoursChg2 signal.
	SetVESC_StatusAmpHoursChg2(int32) *IOV_FrontRightCornerMotor2
}

type IOV_FrontRightCornerMotor2 struct {
	xxx_VESC_StatusAmpHours2    int32
	xxx_VESC_StatusAmpHoursChg2 int32
}

func NewIOV_FrontRightCornerMotor2() *IOV_FrontRightCornerMotor2 {
	m := &IOV_FrontRightCornerMotor2{}
	m.Reset()
	return m
}

func (m *IOV_FrontRightCornerMotor2) Reset() {
	m.xxx_VESC_StatusAmpHours2 = 0
	m.xxx_VESC_StatusAmpHoursChg2 = 0
}

func (m *IOV_FrontRightCornerMotor2) CopyFrom(o IOV_FrontRightCornerMotor2Reader) *IOV_FrontRightCornerMotor2 {
	f, _ := o.MarshalFrame()
	_ = m.UnmarshalFrame(f)
	return m
}

// Descriptor returns the IOV_FrontRightCornerMotor2 descriptor.
func (m *IOV_FrontRightCornerMotor2) Descriptor() *descriptor.Message {
	return Messages().IOV_FrontRightCornerMotor2.Message
}

// String returns a compact string representation of the message.
func (m *IOV_FrontRightCornerMotor2) String() string {
	return cantext.MessageString(m)
}

func (m *IOV_FrontRightCornerMotor2) VESC_StatusAmpHours2() int32 {
	return m.xxx_VESC_StatusAmpHours2
}

func (m *IOV_FrontRightCornerMotor2) SetVESC_StatusAmpHours2(v int32) *IOV_FrontRightCornerMotor2 {
	m.xxx_VESC_StatusAmpHours2 = int32(Messages().IOV_FrontRightCornerMotor2.VESC_StatusAmpHours2.SaturatedCastSigned(int64(v)))
	return m
}

func (m *IOV_FrontRightCornerMotor2) VESC_StatusAmpHoursChg2() int32 {
	return m.xxx_VESC_StatusAmpHoursChg2
}

func (m *IOV_FrontRightCornerMotor2) SetVESC_StatusAmpHoursChg2(v int32) *IOV_FrontRightCornerMotor2 {
	m.xxx_VESC_StatusAmpHoursChg2 = int32(Messages().IOV_FrontRightCornerMotor2.VESC_StatusAmpHoursChg2.SaturatedCastSigned(int64(v)))
	return m
}

// Frame returns a CAN frame representing the message.
func (m *IOV_FrontRightCornerMotor2) Frame() can.Frame {
	md := Messages().IOV_FrontRightCornerMotor2
	f := can.Frame{ID: md.ID, IsExtended: md.IsExtended, Length: md.Length}
	md.VESC_StatusAmpHours2.MarshalSigned(&f.Data, int64(m.xxx_VESC_StatusAmpHours2))
	md.VESC_StatusAmpHoursChg2.MarshalSigned(&f.Data, int64(m.xxx_VESC_StatusAmpHoursChg2))
	return f
}

// MarshalFrame encodes the message as a CAN frame.
func (m *IOV_FrontRightCornerMotor2) MarshalFrame() (can.Frame, error) {
	return m.Frame(), nil
}

// UnmarshalFrame decodes the message from a CAN frame.
func (m *IOV_FrontRightCornerMotor2) UnmarshalFrame(f can.Frame) error {
	md := Messages().IOV_FrontRightCornerMotor2
	switch {
	case f.ID != md.ID:
		return fmt.Errorf(
			"unmarshal IOV_FrontRightCornerMotor2: expects ID 217083902 (got %s with ID %d)", f.String(), f.ID,
		)
	case f.Length != md.Length:
		return fmt.Errorf(
			"unmarshal IOV_FrontRightCornerMotor2: expects length 8 (got %s with length %d)", f.String(), f.Length,
		)
	case f.IsRemote:
		return fmt.Errorf(
			"unmarshal IOV_FrontRightCornerMotor2: expects non-remote frame (got remote frame %s)", f.String(),
		)
	case f.IsExtended != md.IsExtended:
		return fmt.Errorf(
			"unmarshal IOV_FrontRightCornerMotor2: expects extended ID (got %s with standard ID)", f.String(),
		)
	}
	m.xxx_VESC_StatusAmpHours2 = int32(md.VESC_StatusAmpHours2.UnmarshalSigned(f.Data))
	m.xxx_VESC_StatusAmpHoursChg2 = int32(md.VESC_StatusAmpHoursChg2.UnmarshalSigned(f.Data))
	return nil
}

// IOV_FrontRightCornerMotor3Reader provides read access to a IOV_FrontRightCornerMotor3 message.
type IOV_FrontRightCornerMotor3Reader interface {
	can.FrameMarshaler
	// VESC_StatusWattHours2 returns the value of the VESC_StatusWattHours2 signal.
	VESC_StatusWattHours2() int32
	// VESC_StatusWattHoursChg2 returns the value of the VESC_StatusWattHoursChg2 signal.
	VESC_StatusWattHoursChg2() int32
}

// IOV_FrontRightCornerMotor3Writer provides write access to a IOV_FrontRightCornerMotor3 message.
type IOV_FrontRightCornerMotor3Writer interface {
	// CopyFrom copies all values from IOV_FrontRightCornerMotor3.
	CopyFrom(IOV_FrontRightCornerMotor3Reader) *IOV_FrontRightCornerMotor3
	// SetVESC_StatusWattHours2 sets the value of the VESC_StatusWattHours2 signal.
	SetVESC_StatusWattHours2(int32) *IOV_FrontRightCornerMotor3
	// SetVESC_StatusWattHoursChg2 sets the value of the VESC_StatusWattHoursChg2 signal.
	SetVESC_StatusWattHoursChg2(int32) *IOV_FrontRightCornerMotor3
}

type IOV_FrontRightCornerMotor3 struct {
	xxx_VESC_StatusWattHours2    int32
	xxx_VESC_StatusWattHoursChg2 int32
}

func NewIOV_FrontRightCornerMotor3() *IOV_FrontRightCornerMotor3 {
	m := &IOV_FrontRightCornerMotor3{}
	m.Reset()
	return m
}

func (m *IOV_FrontRightCornerMotor3) Reset() {
	m.xxx_VESC_StatusWattHours2 = 0
	m.xxx_VESC_StatusWattHoursChg2 = 0
}

func (m *IOV_FrontRightCornerMotor3) CopyFrom(o IOV_FrontRightCornerMotor3Reader) *IOV_FrontRightCornerMotor3 {
	f, _ := o.MarshalFrame()
	_ = m.UnmarshalFrame(f)
	return m
}

// Descriptor returns the IOV_FrontRightCornerMotor3 descriptor.
func (m *IOV_FrontRightCornerMotor3) Descriptor() *descriptor.Message {
	return Messages().IOV_FrontRightCornerMotor3.Message
}

// String returns a compact string representation of the message.
func (m *IOV_FrontRightCornerMotor3) String() string {
	return cantext.MessageString(m)
}

func (m *IOV_FrontRightCornerMotor3) VESC_StatusWattHours2() int32 {
	return m.xxx_VESC_StatusWattHours2
}

func (m *IOV_FrontRightCornerMotor3) SetVESC_StatusWattHours2(v int32) *IOV_FrontRightCornerMotor3 {
	m.xxx_VESC_StatusWattHours2 = int32(Messages().IOV_FrontRightCornerMotor3.VESC_StatusWattHours2.SaturatedCastSigned(int64(v)))
	return m
}

func (m *IOV_FrontRightCornerMotor3) VESC_StatusWattHoursChg2() int32 {
	return m.xxx_VESC_StatusWattHoursChg2
}

func (m *IOV_FrontRightCornerMotor3) SetVESC_StatusWattHoursChg2(v int32) *IOV_FrontRightCornerMotor3 {
	m.xxx_VESC_StatusWattHoursChg2 = int32(Messages().IOV_FrontRightCornerMotor3.VESC_StatusWattHoursChg2.SaturatedCastSigned(int64(v)))
	return m
}

// Frame returns a CAN frame representing the message.
func (m *IOV_FrontRightCornerMotor3) Frame() can.Frame {
	md := Messages().IOV_FrontRightCornerMotor3
	f := can.Frame{ID: md.ID, IsExtended: md.IsExtended, Length: md.Length}
	md.VESC_StatusWattHours2.MarshalSigned(&f.Data, int64(m.xxx_VESC_StatusWattHours2))
	md.VESC_StatusWattHoursChg2.MarshalSigned(&f.Data, int64(m.xxx_VESC_StatusWattHoursChg2))
	return f
}

// MarshalFrame encodes the message as a CAN frame.
func (m *IOV_FrontRightCornerMotor3) MarshalFrame() (can.Frame, error) {
	return m.Frame(), nil
}

// UnmarshalFrame decodes the message from a CAN frame.
func (m *IOV_FrontRightCornerMotor3) UnmarshalFrame(f can.Frame) error {
	md := Messages().IOV_FrontRightCornerMotor3
	switch {
	case f.ID != md.ID:
		return fmt.Errorf(
			"unmarshal IOV_FrontRightCornerMotor3: expects ID 217087998 (got %s with ID %d)", f.String(), f.ID,
		)
	case f.Length != md.Length:
		return fmt.Errorf(
			"unmarshal IOV_FrontRightCornerMotor3: expects length 8 (got %s with length %d)", f.String(), f.Length,
		)
	case f.IsRemote:
		return fmt.Errorf(
			"unmarshal IOV_FrontRightCornerMotor3: expects non-remote frame (got remote frame %s)", f.String(),
		)
	case f.IsExtended != md.IsExtended:
		return fmt.Errorf(
			"unmarshal IOV_FrontRightCornerMotor3: expects extended ID (got %s with standard ID)", f.String(),
		)
	}
	m.xxx_VESC_StatusWattHours2 = int32(md.VESC_StatusWattHours2.UnmarshalSigned(f.Data))
	m.xxx_VESC_StatusWattHoursChg2 = int32(md.VESC_StatusWattHoursChg2.UnmarshalSigned(f.Data))
	return nil
}

// IOV_FrontRightCornerMotor4Reader provides read access to a IOV_FrontRightCornerMotor4 message.
type IOV_FrontRightCornerMotor4Reader interface {
	can.FrameMarshaler
	// VESC_StatusTempFET2 returns the value of the VESC_StatusTempFET2 signal.
	VESC_StatusTempFET2() int16
	// VESC_StatusTempMotor2 returns the value of the VESC_StatusTempMotor2 signal.
	VESC_StatusTempMotor2() int16
	// VESC_StatusCurrentIn2 returns the value of the VESC_StatusCurrentIn2 signal.
	VESC_StatusCurrentIn2() int16
	// VESC_StatusPIDPos2 returns the value of the VESC_StatusPIDPos2 signal.
	VESC_StatusPIDPos2() int16
}

// IOV_FrontRightCornerMotor4Writer provides write access to a IOV_FrontRightCornerMotor4 message.
type IOV_FrontRightCornerMotor4Writer interface {
	// CopyFrom copies all values from IOV_FrontRightCornerMotor4.
	CopyFrom(IOV_FrontRightCornerMotor4Reader) *IOV_FrontRightCornerMotor4
	// SetVESC_StatusTempFET2 sets the value of the VESC_StatusTempFET2 signal.
	SetVESC_StatusTempFET2(int16) *IOV_FrontRightCornerMotor4
	// SetVESC_StatusTempMotor2 sets the value of the VESC_StatusTempMotor2 signal.
	SetVESC_StatusTempMotor2(int16) *IOV_FrontRightCornerMotor4
	// SetVESC_StatusCurrentIn2 sets the value of the VESC_StatusCurrentIn2 signal.
	SetVESC_StatusCurrentIn2(int16) *IOV_FrontRightCornerMotor4
	// SetVESC_StatusPIDPos2 sets the value of the VESC_StatusPIDPos2 signal.
	SetVESC_StatusPIDPos2(int16) *IOV_FrontRightCornerMotor4
}

type IOV_FrontRightCornerMotor4 struct {
	xxx_VESC_StatusTempFET2   int16
	xxx_VESC_StatusTempMotor2 int16
	xxx_VESC_StatusCurrentIn2 int16
	xxx_VESC_StatusPIDPos2    int16
}

func NewIOV_FrontRightCornerMotor4() *IOV_FrontRightCornerMotor4 {
	m := &IOV_FrontRightCornerMotor4{}
	m.Reset()
	return m
}

func (m *IOV_FrontRightCornerMotor4) Reset() {
	m.xxx_VESC_StatusTempFET2 = 0
	m.xxx_VESC_StatusTempMotor2 = 0
	m.xxx_VESC_StatusCurrentIn2 = 0
	m.xxx_VESC_StatusPIDPos2 = 0
}

func (m *IOV_FrontRightCornerMotor4) CopyFrom(o IOV_FrontRightCornerMotor4Reader) *IOV_FrontRightCornerMotor4 {
	f, _ := o.MarshalFrame()
	_ = m.UnmarshalFrame(f)
	return m
}

// Descriptor returns the IOV_FrontRightCornerMotor4 descriptor.
func (m *IOV_FrontRightCornerMotor4) Descriptor() *descriptor.Message {
	return Messages().IOV_FrontRightCornerMotor4.Message
}

// String returns a compact string representation of the message.
func (m *IOV_FrontRightCornerMotor4) String() string {
	return cantext.MessageString(m)
}

func (m *IOV_FrontRightCornerMotor4) VESC_StatusTempFET2() int16 {
	return m.xxx_VESC_StatusTempFET2
}

func (m *IOV_FrontRightCornerMotor4) SetVESC_StatusTempFET2(v int16) *IOV_FrontRightCornerMotor4 {
	m.xxx_VESC_StatusTempFET2 = int16(Messages().IOV_FrontRightCornerMotor4.VESC_StatusTempFET2.SaturatedCastSigned(int64(v)))
	return m
}

func (m *IOV_FrontRightCornerMotor4) VESC_StatusTempMotor2() int16 {
	return m.xxx_VESC_StatusTempMotor2
}

func (m *IOV_FrontRightCornerMotor4) SetVESC_StatusTempMotor2(v int16) *IOV_FrontRightCornerMotor4 {
	m.xxx_VESC_StatusTempMotor2 = int16(Messages().IOV_FrontRightCornerMotor4.VESC_StatusTempMotor2.SaturatedCastSigned(int64(v)))
	return m
}

func (m *IOV_FrontRightCornerMotor4) VESC_StatusCurrentIn2() int16 {
	return m.xxx_VESC_StatusCurrentIn2
}

func (m *IOV_FrontRightCornerMotor4) SetVESC_StatusCurrentIn2(v int16) *IOV_FrontRightCornerMotor4 {
	m.xxx_VESC_StatusCurrentIn2 = int16(Messages().IOV_FrontRightCornerMotor4.VESC_StatusCurrentIn2.SaturatedCastSigned(int64(v)))
	return m
}

func (m *IOV_FrontRightCornerMotor4) VESC_StatusPIDPos2() int16 {
	return m.xxx_VESC_StatusPIDPos2
}

func (m *IOV_FrontRightCornerMotor4) SetVESC_StatusPIDPos2(v int16) *IOV_FrontRightCornerMotor4 {
	m.xxx_VESC_StatusPIDPos2 = int16(Messages().IOV_FrontRightCornerMotor4.VESC_StatusPIDPos2.SaturatedCastSigned(int64(v)))
	return m
}

// Frame returns a CAN frame representing the message.
func (m *IOV_FrontRightCornerMotor4) Frame() can.Frame {
	md := Messages().IOV_FrontRightCornerMotor4
	f := can.Frame{ID: md.ID, IsExtended: md.IsExtended, Length: md.Length}
	md.VESC_StatusTempFET2.MarshalSigned(&f.Data, int64(m.xxx_VESC_StatusTempFET2))
	md.VESC_StatusTempMotor2.MarshalSigned(&f.Data, int64(m.xxx_VESC_StatusTempMotor2))
	md.VESC_StatusCurrentIn2.MarshalSigned(&f.Data, int64(m.xxx_VESC_StatusCurrentIn2))
	md.VESC_StatusPIDPos2.MarshalSigned(&f.Data, int64(m.xxx_VESC_StatusPIDPos2))
	return f
}

// MarshalFrame encodes the message as a CAN frame.
func (m *IOV_FrontRightCornerMotor4) MarshalFrame() (can.Frame, error) {
	return m.Frame(), nil
}

// UnmarshalFrame decodes the message from a CAN frame.
func (m *IOV_FrontRightCornerMotor4) UnmarshalFrame(f can.Frame) error {
	md := Messages().IOV_FrontRightCornerMotor4
	switch {
	case f.ID != md.ID:
		return fmt.Errorf(
			"unmarshal IOV_FrontRightCornerMotor4: expects ID 217092094 (got %s with ID %d)", f.String(), f.ID,
		)
	case f.Length != md.Length:
		return fmt.Errorf(
			"unmarshal IOV_FrontRightCornerMotor4: expects length 8 (got %s with length %d)", f.String(), f.Length,
		)
	case f.IsRemote:
		return fmt.Errorf(
			"unmarshal IOV_FrontRightCornerMotor4: expects non-remote frame (got remote frame %s)", f.String(),
		)
	case f.IsExtended != md.IsExtended:
		return fmt.Errorf(
			"unmarshal IOV_FrontRightCornerMotor4: expects extended ID (got %s with standard ID)", f.String(),
		)
	}
	m.xxx_VESC_StatusTempFET2 = int16(md.VESC_StatusTempFET2.UnmarshalSigned(f.Data))
	m.xxx_VESC_StatusTempMotor2 = int16(md.VESC_StatusTempMotor2.UnmarshalSigned(f.Data))
	m.xxx_VESC_StatusCurrentIn2 = int16(md.VESC_StatusCurrentIn2.UnmarshalSigned(f.Data))
	m.xxx_VESC_StatusPIDPos2 = int16(md.VESC_StatusPIDPos2.UnmarshalSigned(f.Data))
	return nil
}

// IOV_BackLeftCornerMotor1Reader provides read access to a IOV_BackLeftCornerMotor1 message.
type IOV_BackLeftCornerMotor1Reader interface {
	can.FrameMarshaler
	// VESC_StatusERPM3 returns the value of the VESC_StatusERPM3 signal.
	VESC_StatusERPM3() int32
	// VESC_StatusCurrent3 returns the value of the VESC_StatusCurrent3 signal.
	VESC_StatusCurrent3() int16
	// VESC_StatusDutyCycle3 returns the value of the VESC_StatusDutyCycle3 signal.
	VESC_StatusDutyCycle3() int16
}

// IOV_BackLeftCornerMotor1Writer provides write access to a IOV_BackLeftCornerMotor1 message.
type IOV_BackLeftCornerMotor1Writer interface {
	// CopyFrom copies all values from IOV_BackLeftCornerMotor1.
	CopyFrom(IOV_BackLeftCornerMotor1Reader) *IOV_BackLeftCornerMotor1
	// SetVESC_StatusERPM3 sets the value of the VESC_StatusERPM3 signal.
	SetVESC_StatusERPM3(int32) *IOV_BackLeftCornerMotor1
	// SetVESC_StatusCurrent3 sets the value of the VESC_StatusCurrent3 signal.
	SetVESC_StatusCurrent3(int16) *IOV_BackLeftCornerMotor1
	// SetVESC_StatusDutyCycle3 sets the value of the VESC_StatusDutyCycle3 signal.
	SetVESC_StatusDutyCycle3(int16) *IOV_BackLeftCornerMotor1
}

type IOV_BackLeftCornerMotor1 struct {
	xxx_VESC_StatusERPM3      int32
	xxx_VESC_StatusCurrent3   int16
	xxx_VESC_StatusDutyCycle3 int16
}

func NewIOV_BackLeftCornerMotor1() *IOV_BackLeftCornerMotor1 {
	m := &IOV_BackLeftCornerMotor1{}
	m.Reset()
	return m
}

func (m *IOV_BackLeftCornerMotor1) Reset() {
	m.xxx_VESC_StatusERPM3 = 0
	m.xxx_VESC_StatusCurrent3 = 0
	m.xxx_VESC_StatusDutyCycle3 = 0
}

func (m *IOV_BackLeftCornerMotor1) CopyFrom(o IOV_BackLeftCornerMotor1Reader) *IOV_BackLeftCornerMotor1 {
	f, _ := o.MarshalFrame()
	_ = m.UnmarshalFrame(f)
	return m
}

// Descriptor returns the IOV_BackLeftCornerMotor1 descriptor.
func (m *IOV_BackLeftCornerMotor1) Descriptor() *descriptor.Message {
	return Messages().IOV_BackLeftCornerMotor1.Message
}

// String returns a compact string representation of the message.
func (m *IOV_BackLeftCornerMotor1) String() string {
	return cantext.MessageString(m)
}

func (m *IOV_BackLeftCornerMotor1) VESC_StatusERPM3() int32 {
	return m.xxx_VESC_StatusERPM3
}

func (m *IOV_BackLeftCornerMotor1) SetVESC_StatusERPM3(v int32) *IOV_BackLeftCornerMotor1 {
	m.xxx_VESC_StatusERPM3 = int32(Messages().IOV_BackLeftCornerMotor1.VESC_StatusERPM3.SaturatedCastSigned(int64(v)))
	return m
}

func (m *IOV_BackLeftCornerMotor1) VESC_StatusCurrent3() int16 {
	return m.xxx_VESC_StatusCurrent3
}

func (m *IOV_BackLeftCornerMotor1) SetVESC_StatusCurrent3(v int16) *IOV_BackLeftCornerMotor1 {
	m.xxx_VESC_StatusCurrent3 = int16(Messages().IOV_BackLeftCornerMotor1.VESC_StatusCurrent3.SaturatedCastSigned(int64(v)))
	return m
}

func (m *IOV_BackLeftCornerMotor1) VESC_StatusDutyCycle3() int16 {
	return m.xxx_VESC_StatusDutyCycle3
}

func (m *IOV_BackLeftCornerMotor1) SetVESC_StatusDutyCycle3(v int16) *IOV_BackLeftCornerMotor1 {
	m.xxx_VESC_StatusDutyCycle3 = int16(Messages().IOV_BackLeftCornerMotor1.VESC_StatusDutyCycle3.SaturatedCastSigned(int64(v)))
	return m
}

// Frame returns a CAN frame representing the message.
func (m *IOV_BackLeftCornerMotor1) Frame() can.Frame {
	md := Messages().IOV_BackLeftCornerMotor1
	f := can.Frame{ID: md.ID, IsExtended: md.IsExtended, Length: md.Length}
	md.VESC_StatusERPM3.MarshalSigned(&f.Data, int64(m.xxx_VESC_StatusERPM3))
	md.VESC_StatusCurrent3.MarshalSigned(&f.Data, int64(m.xxx_VESC_StatusCurrent3))
	md.VESC_StatusDutyCycle3.MarshalSigned(&f.Data, int64(m.xxx_VESC_StatusDutyCycle3))
	return f
}

// MarshalFrame encodes the message as a CAN frame.
func (m *IOV_BackLeftCornerMotor1) MarshalFrame() (can.Frame, error) {
	return m.Frame(), nil
}

// UnmarshalFrame decodes the message from a CAN frame.
func (m *IOV_BackLeftCornerMotor1) UnmarshalFrame(f can.Frame) error {
	md := Messages().IOV_BackLeftCornerMotor1
	switch {
	case f.ID != md.ID:
		return fmt.Errorf(
			"unmarshal IOV_BackLeftCornerMotor1: expects ID 217096190 (got %s with ID %d)", f.String(), f.ID,
		)
	case f.Length != md.Length:
		return fmt.Errorf(
			"unmarshal IOV_BackLeftCornerMotor1: expects length 8 (got %s with length %d)", f.String(), f.Length,
		)
	case f.IsRemote:
		return fmt.Errorf(
			"unmarshal IOV_BackLeftCornerMotor1: expects non-remote frame (got remote frame %s)", f.String(),
		)
	case f.IsExtended != md.IsExtended:
		return fmt.Errorf(
			"unmarshal IOV_BackLeftCornerMotor1: expects extended ID (got %s with standard ID)", f.String(),
		)
	}
	m.xxx_VESC_StatusERPM3 = int32(md.VESC_StatusERPM3.UnmarshalSigned(f.Data))
	m.xxx_VESC_StatusCurrent3 = int16(md.VESC_StatusCurrent3.UnmarshalSigned(f.Data))
	m.xxx_VESC_StatusDutyCycle3 = int16(md.VESC_StatusDutyCycle3.UnmarshalSigned(f.Data))
	return nil
}

// IOV_BackLeftCornerMotor2Reader provides read access to a IOV_BackLeftCornerMotor2 message.
type IOV_BackLeftCornerMotor2Reader interface {
	can.FrameMarshaler
	// VESC_StatusAmpHours3 returns the value of the VESC_StatusAmpHours3 signal.
	VESC_StatusAmpHours3() int32
	// VESC_StatusAmpHoursChg3 returns the value of the VESC_StatusAmpHoursChg3 signal.
	VESC_StatusAmpHoursChg3() int32
}

// IOV_BackLeftCornerMotor2Writer provides write access to a IOV_BackLeftCornerMotor2 message.
type IOV_BackLeftCornerMotor2Writer interface {
	// CopyFrom copies all values from IOV_BackLeftCornerMotor2.
	CopyFrom(IOV_BackLeftCornerMotor2Reader) *IOV_BackLeftCornerMotor2
	// SetVESC_StatusAmpHours3 sets the value of the VESC_StatusAmpHours3 signal.
	SetVESC_StatusAmpHours3(int32) *IOV_BackLeftCornerMotor2
	// SetVESC_StatusAmpHoursChg3 sets the value of the VESC_StatusAmpHoursChg3 signal.
	SetVESC_StatusAmpHoursChg3(int32) *IOV_BackLeftCornerMotor2
}

type IOV_BackLeftCornerMotor2 struct {
	xxx_VESC_StatusAmpHours3    int32
	xxx_VESC_StatusAmpHoursChg3 int32
}

func NewIOV_BackLeftCornerMotor2() *IOV_BackLeftCornerMotor2 {
	m := &IOV_BackLeftCornerMotor2{}
	m.Reset()
	return m
}

func (m *IOV_BackLeftCornerMotor2) Reset() {
	m.xxx_VESC_StatusAmpHours3 = 0
	m.xxx_VESC_StatusAmpHoursChg3 = 0
}

func (m *IOV_BackLeftCornerMotor2) CopyFrom(o IOV_BackLeftCornerMotor2Reader) *IOV_BackLeftCornerMotor2 {
	f, _ := o.MarshalFrame()
	_ = m.UnmarshalFrame(f)
	return m
}

// Descriptor returns the IOV_BackLeftCornerMotor2 descriptor.
func (m *IOV_BackLeftCornerMotor2) Descriptor() *descriptor.Message {
	return Messages().IOV_BackLeftCornerMotor2.Message
}

// String returns a compact string representation of the message.
func (m *IOV_BackLeftCornerMotor2) String() string {
	return cantext.MessageString(m)
}

func (m *IOV_BackLeftCornerMotor2) VESC_StatusAmpHours3() int32 {
	return m.xxx_VESC_StatusAmpHours3
}

func (m *IOV_BackLeftCornerMotor2) SetVESC_StatusAmpHours3(v int32) *IOV_BackLeftCornerMotor2 {
	m.xxx_VESC_StatusAmpHours3 = int32(Messages().IOV_BackLeftCornerMotor2.VESC_StatusAmpHours3.SaturatedCastSigned(int64(v)))
	return m
}

func (m *IOV_BackLeftCornerMotor2) VESC_StatusAmpHoursChg3() int32 {
	return m.xxx_VESC_StatusAmpHoursChg3
}

func (m *IOV_BackLeftCornerMotor2) SetVESC_StatusAmpHoursChg3(v int32) *IOV_BackLeftCornerMotor2 {
	m.xxx_VESC_StatusAmpHoursChg3 = int32(Messages().IOV_BackLeftCornerMotor2.VESC_StatusAmpHoursChg3.SaturatedCastSigned(int64(v)))
	return m
}

// Frame returns a CAN frame representing the message.
func (m *IOV_BackLeftCornerMotor2) Frame() can.Frame {
	md := Messages().IOV_BackLeftCornerMotor2
	f := can.Frame{ID: md.ID, IsExtended: md.IsExtended, Length: md.Length}
	md.VESC_StatusAmpHours3.MarshalSigned(&f.Data, int64(m.xxx_VESC_StatusAmpHours3))
	md.VESC_StatusAmpHoursChg3.MarshalSigned(&f.Data, int64(m.xxx_VESC_StatusAmpHoursChg3))
	return f
}

// MarshalFrame encodes the message as a CAN frame.
func (m *IOV_BackLeftCornerMotor2) MarshalFrame() (can.Frame, error) {
	return m.Frame(), nil
}

// UnmarshalFrame decodes the message from a CAN frame.
func (m *IOV_BackLeftCornerMotor2) UnmarshalFrame(f can.Frame) error {
	md := Messages().IOV_BackLeftCornerMotor2
	switch {
	case f.ID != md.ID:
		return fmt.Errorf(
			"unmarshal IOV_BackLeftCornerMotor2: expects ID 217100286 (got %s with ID %d)", f.String(), f.ID,
		)
	case f.Length != md.Length:
		return fmt.Errorf(
			"unmarshal IOV_BackLeftCornerMotor2: expects length 8 (got %s with length %d)", f.String(), f.Length,
		)
	case f.IsRemote:
		return fmt.Errorf(
			"unmarshal IOV_BackLeftCornerMotor2: expects non-remote frame (got remote frame %s)", f.String(),
		)
	case f.IsExtended != md.IsExtended:
		return fmt.Errorf(
			"unmarshal IOV_BackLeftCornerMotor2: expects extended ID (got %s with standard ID)", f.String(),
		)
	}
	m.xxx_VESC_StatusAmpHours3 = int32(md.VESC_StatusAmpHours3.UnmarshalSigned(f.Data))
	m.xxx_VESC_StatusAmpHoursChg3 = int32(md.VESC_StatusAmpHoursChg3.UnmarshalSigned(f.Data))
	return nil
}

// IOV_BackLeftCornerMotor3Reader provides read access to a IOV_BackLeftCornerMotor3 message.
type IOV_BackLeftCornerMotor3Reader interface {
	can.FrameMarshaler
	// VESC_StatusWattHours3 returns the value of the VESC_StatusWattHours3 signal.
	VESC_StatusWattHours3() int32
	// VESC_StatusWattHoursChg3 returns the value of the VESC_StatusWattHoursChg3 signal.
	VESC_StatusWattHoursChg3() int32
}

// IOV_BackLeftCornerMotor3Writer provides write access to a IOV_BackLeftCornerMotor3 message.
type IOV_BackLeftCornerMotor3Writer interface {
	// CopyFrom copies all values from IOV_BackLeftCornerMotor3.
	CopyFrom(IOV_BackLeftCornerMotor3Reader) *IOV_BackLeftCornerMotor3
	// SetVESC_StatusWattHours3 sets the value of the VESC_StatusWattHours3 signal.
	SetVESC_StatusWattHours3(int32) *IOV_BackLeftCornerMotor3
	// SetVESC_StatusWattHoursChg3 sets the value of the VESC_StatusWattHoursChg3 signal.
	SetVESC_StatusWattHoursChg3(int32) *IOV_BackLeftCornerMotor3
}

type IOV_BackLeftCornerMotor3 struct {
	xxx_VESC_StatusWattHours3    int32
	xxx_VESC_StatusWattHoursChg3 int32
}

func NewIOV_BackLeftCornerMotor3() *IOV_BackLeftCornerMotor3 {
	m := &IOV_BackLeftCornerMotor3{}
	m.Reset()
	return m
}

func (m *IOV_BackLeftCornerMotor3) Reset() {
	m.xxx_VESC_StatusWattHours3 = 0
	m.xxx_VESC_StatusWattHoursChg3 = 0
}

func (m *IOV_BackLeftCornerMotor3) CopyFrom(o IOV_BackLeftCornerMotor3Reader) *IOV_BackLeftCornerMotor3 {
	f, _ := o.MarshalFrame()
	_ = m.UnmarshalFrame(f)
	return m
}

// Descriptor returns the IOV_BackLeftCornerMotor3 descriptor.
func (m *IOV_BackLeftCornerMotor3) Descriptor() *descriptor.Message {
	return Messages().IOV_BackLeftCornerMotor3.Message
}

// String returns a compact string representation of the message.
func (m *IOV_BackLeftCornerMotor3) String() string {
	return cantext.MessageString(m)
}

func (m *IOV_BackLeftCornerMotor3) VESC_StatusWattHours3() int32 {
	return m.xxx_VESC_StatusWattHours3
}

func (m *IOV_BackLeftCornerMotor3) SetVESC_StatusWattHours3(v int32) *IOV_BackLeftCornerMotor3 {
	m.xxx_VESC_StatusWattHours3 = int32(Messages().IOV_BackLeftCornerMotor3.VESC_StatusWattHours3.SaturatedCastSigned(int64(v)))
	return m
}

func (m *IOV_BackLeftCornerMotor3) VESC_StatusWattHoursChg3() int32 {
	return m.xxx_VESC_StatusWattHoursChg3
}

func (m *IOV_BackLeftCornerMotor3) SetVESC_StatusWattHoursChg3(v int32) *IOV_BackLeftCornerMotor3 {
	m.xxx_VESC_StatusWattHoursChg3 = int32(Messages().IOV_BackLeftCornerMotor3.VESC_StatusWattHoursChg3.SaturatedCastSigned(int64(v)))
	return m
}

// Frame returns a CAN frame representing the message.
func (m *IOV_BackLeftCornerMotor3) Frame() can.Frame {
	md := Messages().IOV_BackLeftCornerMotor3
	f := can.Frame{ID: md.ID, IsExtended: md.IsExtended, Length: md.Length}
	md.VESC_StatusWattHours3.MarshalSigned(&f.Data, int64(m.xxx_VESC_StatusWattHours3))
	md.VESC_StatusWattHoursChg3.MarshalSigned(&f.Data, int64(m.xxx_VESC_StatusWattHoursChg3))
	return f
}

// MarshalFrame encodes the message as a CAN frame.
func (m *IOV_BackLeftCornerMotor3) MarshalFrame() (can.Frame, error) {
	return m.Frame(), nil
}

// UnmarshalFrame decodes the message from a CAN frame.
func (m *IOV_BackLeftCornerMotor3) UnmarshalFrame(f can.Frame) error {
	md := Messages().IOV_BackLeftCornerMotor3
	switch {
	case f.ID != md.ID:
		return fmt.Errorf(
			"unmarshal IOV_BackLeftCornerMotor3: expects ID 217104382 (got %s with ID %d)", f.String(), f.ID,
		)
	case f.Length != md.Length:
		return fmt.Errorf(
			"unmarshal IOV_BackLeftCornerMotor3: expects length 8 (got %s with length %d)", f.String(), f.Length,
		)
	case f.IsRemote:
		return fmt.Errorf(
			"unmarshal IOV_BackLeftCornerMotor3: expects non-remote frame (got remote frame %s)", f.String(),
		)
	case f.IsExtended != md.IsExtended:
		return fmt.Errorf(
			"unmarshal IOV_BackLeftCornerMotor3: expects extended ID (got %s with standard ID)", f.String(),
		)
	}
	m.xxx_VESC_StatusWattHours3 = int32(md.VESC_StatusWattHours3.UnmarshalSigned(f.Data))
	m.xxx_VESC_StatusWattHoursChg3 = int32(md.VESC_StatusWattHoursChg3.UnmarshalSigned(f.Data))
	return nil
}

// IOV_BackLeftCornerMotor4Reader provides read access to a IOV_BackLeftCornerMotor4 message.
type IOV_BackLeftCornerMotor4Reader interface {
	can.FrameMarshaler
	// VESC_StatusTempFET3 returns the value of the VESC_StatusTempFET3 signal.
	VESC_StatusTempFET3() int16
	// VESC_StatusTempMotor3 returns the value of the VESC_StatusTempMotor3 signal.
	VESC_StatusTempMotor3() int16
	// VESC_StatusCurrentIn3 returns the value of the VESC_StatusCurrentIn3 signal.
	VESC_StatusCurrentIn3() int16
	// VESC_StatusPIDPos3 returns the value of the VESC_StatusPIDPos3 signal.
	VESC_StatusPIDPos3() int16
}

// IOV_BackLeftCornerMotor4Writer provides write access to a IOV_BackLeftCornerMotor4 message.
type IOV_BackLeftCornerMotor4Writer interface {
	// CopyFrom copies all values from IOV_BackLeftCornerMotor4.
	CopyFrom(IOV_BackLeftCornerMotor4Reader) *IOV_BackLeftCornerMotor4
	// SetVESC_StatusTempFET3 sets the value of the VESC_StatusTempFET3 signal.
	SetVESC_StatusTempFET3(int16) *IOV_BackLeftCornerMotor4
	// SetVESC_StatusTempMotor3 sets the value of the VESC_StatusTempMotor3 signal.
	SetVESC_StatusTempMotor3(int16) *IOV_BackLeftCornerMotor4
	// SetVESC_StatusCurrentIn3 sets the value of the VESC_StatusCurrentIn3 signal.
	SetVESC_StatusCurrentIn3(int16) *IOV_BackLeftCornerMotor4
	// SetVESC_StatusPIDPos3 sets the value of the VESC_StatusPIDPos3 signal.
	SetVESC_StatusPIDPos3(int16) *IOV_BackLeftCornerMotor4
}

type IOV_BackLeftCornerMotor4 struct {
	xxx_VESC_StatusTempFET3   int16
	xxx_VESC_StatusTempMotor3 int16
	xxx_VESC_StatusCurrentIn3 int16
	xxx_VESC_StatusPIDPos3    int16
}

func NewIOV_BackLeftCornerMotor4() *IOV_BackLeftCornerMotor4 {
	m := &IOV_BackLeftCornerMotor4{}
	m.Reset()
	return m
}

func (m *IOV_BackLeftCornerMotor4) Reset() {
	m.xxx_VESC_StatusTempFET3 = 0
	m.xxx_VESC_StatusTempMotor3 = 0
	m.xxx_VESC_StatusCurrentIn3 = 0
	m.xxx_VESC_StatusPIDPos3 = 0
}

func (m *IOV_BackLeftCornerMotor4) CopyFrom(o IOV_BackLeftCornerMotor4Reader) *IOV_BackLeftCornerMotor4 {
	f, _ := o.MarshalFrame()
	_ = m.UnmarshalFrame(f)
	return m
}

// Descriptor returns the IOV_BackLeftCornerMotor4 descriptor.
func (m *IOV_BackLeftCornerMotor4) Descriptor() *descriptor.Message {
	return Messages().IOV_BackLeftCornerMotor4.Message
}

// String returns a compact string representation of the message.
func (m *IOV_BackLeftCornerMotor4) String() string {
	return cantext.MessageString(m)
}

func (m *IOV_BackLeftCornerMotor4) VESC_StatusTempFET3() int16 {
	return m.xxx_VESC_StatusTempFET3
}

func (m *IOV_BackLeftCornerMotor4) SetVESC_StatusTempFET3(v int16) *IOV_BackLeftCornerMotor4 {
	m.xxx_VESC_StatusTempFET3 = int16(Messages().IOV_BackLeftCornerMotor4.VESC_StatusTempFET3.SaturatedCastSigned(int64(v)))
	return m
}

func (m *IOV_BackLeftCornerMotor4) VESC_StatusTempMotor3() int16 {
	return m.xxx_VESC_StatusTempMotor3
}

func (m *IOV_BackLeftCornerMotor4) SetVESC_StatusTempMotor3(v int16) *IOV_BackLeftCornerMotor4 {
	m.xxx_VESC_StatusTempMotor3 = int16(Messages().IOV_BackLeftCornerMotor4.VESC_StatusTempMotor3.SaturatedCastSigned(int64(v)))
	return m
}

func (m *IOV_BackLeftCornerMotor4) VESC_StatusCurrentIn3() int16 {
	return m.xxx_VESC_StatusCurrentIn3
}

func (m *IOV_BackLeftCornerMotor4) SetVESC_StatusCurrentIn3(v int16) *IOV_BackLeftCornerMotor4 {
	m.xxx_VESC_StatusCurrentIn3 = int16(Messages().IOV_BackLeftCornerMotor4.VESC_StatusCurrentIn3.SaturatedCastSigned(int64(v)))
	return m
}

func (m *IOV_BackLeftCornerMotor4) VESC_StatusPIDPos3() int16 {
	return m.xxx_VESC_StatusPIDPos3
}

func (m *IOV_BackLeftCornerMotor4) SetVESC_StatusPIDPos3(v int16) *IOV_BackLeftCornerMotor4 {
	m.xxx_VESC_StatusPIDPos3 = int16(Messages().IOV_BackLeftCornerMotor4.VESC_StatusPIDPos3.SaturatedCastSigned(int64(v)))
	return m
}

// Frame returns a CAN frame representing the message.
func (m *IOV_BackLeftCornerMotor4) Frame() can.Frame {
	md := Messages().IOV_BackLeftCornerMotor4
	f := can.Frame{ID: md.ID, IsExtended: md.IsExtended, Length: md.Length}
	md.VESC_StatusTempFET3.MarshalSigned(&f.Data, int64(m.xxx_VESC_StatusTempFET3))
	md.VESC_StatusTempMotor3.MarshalSigned(&f.Data, int64(m.xxx_VESC_StatusTempMotor3))
	md.VESC_StatusCurrentIn3.MarshalSigned(&f.Data, int64(m.xxx_VESC_StatusCurrentIn3))
	md.VESC_StatusPIDPos3.MarshalSigned(&f.Data, int64(m.xxx_VESC_StatusPIDPos3))
	return f
}

// MarshalFrame encodes the message as a CAN frame.
func (m *IOV_BackLeftCornerMotor4) MarshalFrame() (can.Frame, error) {
	return m.Frame(), nil
}

// UnmarshalFrame decodes the message from a CAN frame.
func (m *IOV_BackLeftCornerMotor4) UnmarshalFrame(f can.Frame) error {
	md := Messages().IOV_BackLeftCornerMotor4
	switch {
	case f.ID != md.ID:
		return fmt.Errorf(
			"unmarshal IOV_BackLeftCornerMotor4: expects ID 217108478 (got %s with ID %d)", f.String(), f.ID,
		)
	case f.Length != md.Length:
		return fmt.Errorf(
			"unmarshal IOV_BackLeftCornerMotor4: expects length 8 (got %s with length %d)", f.String(), f.Length,
		)
	case f.IsRemote:
		return fmt.Errorf(
			"unmarshal IOV_BackLeftCornerMotor4: expects non-remote frame (got remote frame %s)", f.String(),
		)
	case f.IsExtended != md.IsExtended:
		return fmt.Errorf(
			"unmarshal IOV_BackLeftCornerMotor4: expects extended ID (got %s with standard ID)", f.String(),
		)
	}
	m.xxx_VESC_StatusTempFET3 = int16(md.VESC_StatusTempFET3.UnmarshalSigned(f.Data))
	m.xxx_VESC_StatusTempMotor3 = int16(md.VESC_StatusTempMotor3.UnmarshalSigned(f.Data))
	m.xxx_VESC_StatusCurrentIn3 = int16(md.VESC_StatusCurrentIn3.UnmarshalSigned(f.Data))
	m.xxx_VESC_StatusPIDPos3 = int16(md.VESC_StatusPIDPos3.UnmarshalSigned(f.Data))
	return nil
}

// IOV_BackRightCornerMotor1Reader provides read access to a IOV_BackRightCornerMotor1 message.
type IOV_BackRightCornerMotor1Reader interface {
	can.FrameMarshaler
	// VESC_StatusERPM4 returns the value of the VESC_StatusERPM4 signal.
	VESC_StatusERPM4() int32
	// VESC_StatusCurrent4 returns the value of the VESC_StatusCurrent4 signal.
	VESC_StatusCurrent4() int16
	// VESC_StatusDutyCycle4 returns the value of the VESC_StatusDutyCycle4 signal.
	VESC_StatusDutyCycle4() int16
}

// IOV_BackRightCornerMotor1Writer provides write access to a IOV_BackRightCornerMotor1 message.
type IOV_BackRightCornerMotor1Writer interface {
	// CopyFrom copies all values from IOV_BackRightCornerMotor1.
	CopyFrom(IOV_BackRightCornerMotor1Reader) *IOV_BackRightCornerMotor1
	// SetVESC_StatusERPM4 sets the value of the VESC_StatusERPM4 signal.
	SetVESC_StatusERPM4(int32) *IOV_BackRightCornerMotor1
	// SetVESC_StatusCurrent4 sets the value of the VESC_StatusCurrent4 signal.
	SetVESC_StatusCurrent4(int16) *IOV_BackRightCornerMotor1
	// SetVESC_StatusDutyCycle4 sets the value of the VESC_StatusDutyCycle4 signal.
	SetVESC_StatusDutyCycle4(int16) *IOV_BackRightCornerMotor1
}

type IOV_BackRightCornerMotor1 struct {
	xxx_VESC_StatusERPM4      int32
	xxx_VESC_StatusCurrent4   int16
	xxx_VESC_StatusDutyCycle4 int16
}

func NewIOV_BackRightCornerMotor1() *IOV_BackRightCornerMotor1 {
	m := &IOV_BackRightCornerMotor1{}
	m.Reset()
	return m
}

func (m *IOV_BackRightCornerMotor1) Reset() {
	m.xxx_VESC_StatusERPM4 = 0
	m.xxx_VESC_StatusCurrent4 = 0
	m.xxx_VESC_StatusDutyCycle4 = 0
}

func (m *IOV_BackRightCornerMotor1) CopyFrom(o IOV_BackRightCornerMotor1Reader) *IOV_BackRightCornerMotor1 {
	f, _ := o.MarshalFrame()
	_ = m.UnmarshalFrame(f)
	return m
}

// Descriptor returns the IOV_BackRightCornerMotor1 descriptor.
func (m *IOV_BackRightCornerMotor1) Descriptor() *descriptor.Message {
	return Messages().IOV_BackRightCornerMotor1.Message
}

// String returns a compact string representation of the message.
func (m *IOV_BackRightCornerMotor1) String() string {
	return cantext.MessageString(m)
}

func (m *IOV_BackRightCornerMotor1) VESC_StatusERPM4() int32 {
	return m.xxx_VESC_StatusERPM4
}

func (m *IOV_BackRightCornerMotor1) SetVESC_StatusERPM4(v int32) *IOV_BackRightCornerMotor1 {
	m.xxx_VESC_StatusERPM4 = int32(Messages().IOV_BackRightCornerMotor1.VESC_StatusERPM4.SaturatedCastSigned(int64(v)))
	return m
}

func (m *IOV_BackRightCornerMotor1) VESC_StatusCurrent4() int16 {
	return m.xxx_VESC_StatusCurrent4
}

func (m *IOV_BackRightCornerMotor1) SetVESC_StatusCurrent4(v int16) *IOV_BackRightCornerMotor1 {
	m.xxx_VESC_StatusCurrent4 = int16(Messages().IOV_BackRightCornerMotor1.VESC_StatusCurrent4.SaturatedCastSigned(int64(v)))
	return m
}

func (m *IOV_BackRightCornerMotor1) VESC_StatusDutyCycle4() int16 {
	return m.xxx_VESC_StatusDutyCycle4
}

func (m *IOV_BackRightCornerMotor1) SetVESC_StatusDutyCycle4(v int16) *IOV_BackRightCornerMotor1 {
	m.xxx_VESC_StatusDutyCycle4 = int16(Messages().IOV_BackRightCornerMotor1.VESC_StatusDutyCycle4.SaturatedCastSigned(int64(v)))
	return m
}

// Frame returns a CAN frame representing the message.
func (m *IOV_BackRightCornerMotor1) Frame() can.Frame {
	md := Messages().IOV_BackRightCornerMotor1
	f := can.Frame{ID: md.ID, IsExtended: md.IsExtended, Length: md.Length}
	md.VESC_StatusERPM4.MarshalSigned(&f.Data, int64(m.xxx_VESC_StatusERPM4))
	md.VESC_StatusCurrent4.MarshalSigned(&f.Data, int64(m.xxx_VESC_StatusCurrent4))
	md.VESC_StatusDutyCycle4.MarshalSigned(&f.Data, int64(m.xxx_VESC_StatusDutyCycle4))
	return f
}

// MarshalFrame encodes the message as a CAN frame.
func (m *IOV_BackRightCornerMotor1) MarshalFrame() (can.Frame, error) {
	return m.Frame(), nil
}

// UnmarshalFrame decodes the message from a CAN frame.
func (m *IOV_BackRightCornerMotor1) UnmarshalFrame(f can.Frame) error {
	md := Messages().IOV_BackRightCornerMotor1
	switch {
	case f.ID != md.ID:
		return fmt.Errorf(
			"unmarshal IOV_BackRightCornerMotor1: expects ID 217112574 (got %s with ID %d)", f.String(), f.ID,
		)
	case f.Length != md.Length:
		return fmt.Errorf(
			"unmarshal IOV_BackRightCornerMotor1: expects length 8 (got %s with length %d)", f.String(), f.Length,
		)
	case f.IsRemote:
		return fmt.Errorf(
			"unmarshal IOV_BackRightCornerMotor1: expects non-remote frame (got remote frame %s)", f.String(),
		)
	case f.IsExtended != md.IsExtended:
		return fmt.Errorf(
			"unmarshal IOV_BackRightCornerMotor1: expects extended ID (got %s with standard ID)", f.String(),
		)
	}
	m.xxx_VESC_StatusERPM4 = int32(md.VESC_StatusERPM4.UnmarshalSigned(f.Data))
	m.xxx_VESC_StatusCurrent4 = int16(md.VESC_StatusCurrent4.UnmarshalSigned(f.Data))
	m.xxx_VESC_StatusDutyCycle4 = int16(md.VESC_StatusDutyCycle4.UnmarshalSigned(f.Data))
	return nil
}

// IOV_BackRightCornerMotor2Reader provides read access to a IOV_BackRightCornerMotor2 message.
type IOV_BackRightCornerMotor2Reader interface {
	can.FrameMarshaler
	// VESC_StatusAmpHours4 returns the value of the VESC_StatusAmpHours4 signal.
	VESC_StatusAmpHours4() int32
	// VESC_StatusAmpHoursChg4 returns the value of the VESC_StatusAmpHoursChg4 signal.
	VESC_StatusAmpHoursChg4() int32
}

// IOV_BackRightCornerMotor2Writer provides write access to a IOV_BackRightCornerMotor2 message.
type IOV_BackRightCornerMotor2Writer interface {
	// CopyFrom copies all values from IOV_BackRightCornerMotor2.
	CopyFrom(IOV_BackRightCornerMotor2Reader) *IOV_BackRightCornerMotor2
	// SetVESC_StatusAmpHours4 sets the value of the VESC_StatusAmpHours4 signal.
	SetVESC_StatusAmpHours4(int32) *IOV_BackRightCornerMotor2
	// SetVESC_StatusAmpHoursChg4 sets the value of the VESC_StatusAmpHoursChg4 signal.
	SetVESC_StatusAmpHoursChg4(int32) *IOV_BackRightCornerMotor2
}

type IOV_BackRightCornerMotor2 struct {
	xxx_VESC_StatusAmpHours4    int32
	xxx_VESC_StatusAmpHoursChg4 int32
}

func NewIOV_BackRightCornerMotor2() *IOV_BackRightCornerMotor2 {
	m := &IOV_BackRightCornerMotor2{}
	m.Reset()
	return m
}

func (m *IOV_BackRightCornerMotor2) Reset() {
	m.xxx_VESC_StatusAmpHours4 = 0
	m.xxx_VESC_StatusAmpHoursChg4 = 0
}

func (m *IOV_BackRightCornerMotor2) CopyFrom(o IOV_BackRightCornerMotor2Reader) *IOV_BackRightCornerMotor2 {
	f, _ := o.MarshalFrame()
	_ = m.UnmarshalFrame(f)
	return m
}

// Descriptor returns the IOV_BackRightCornerMotor2 descriptor.
func (m *IOV_BackRightCornerMotor2) Descriptor() *descriptor.Message {
	return Messages().IOV_BackRightCornerMotor2.Message
}

// String returns a compact string representation of the message.
func (m *IOV_BackRightCornerMotor2) String() string {
	return cantext.MessageString(m)
}

func (m *IOV_BackRightCornerMotor2) VESC_StatusAmpHours4() int32 {
	return m.xxx_VESC_StatusAmpHours4
}

func (m *IOV_BackRightCornerMotor2) SetVESC_StatusAmpHours4(v int32) *IOV_BackRightCornerMotor2 {
	m.xxx_VESC_StatusAmpHours4 = int32(Messages().IOV_BackRightCornerMotor2.VESC_StatusAmpHours4.SaturatedCastSigned(int64(v)))
	return m
}

func (m *IOV_BackRightCornerMotor2) VESC_StatusAmpHoursChg4() int32 {
	return m.xxx_VESC_StatusAmpHoursChg4
}

func (m *IOV_BackRightCornerMotor2) SetVESC_StatusAmpHoursChg4(v int32) *IOV_BackRightCornerMotor2 {
	m.xxx_VESC_StatusAmpHoursChg4 = int32(Messages().IOV_BackRightCornerMotor2.VESC_StatusAmpHoursChg4.SaturatedCastSigned(int64(v)))
	return m
}

// Frame returns a CAN frame representing the message.
func (m *IOV_BackRightCornerMotor2) Frame() can.Frame {
	md := Messages().IOV_BackRightCornerMotor2
	f := can.Frame{ID: md.ID, IsExtended: md.IsExtended, Length: md.Length}
	md.VESC_StatusAmpHours4.MarshalSigned(&f.Data, int64(m.xxx_VESC_StatusAmpHours4))
	md.VESC_StatusAmpHoursChg4.MarshalSigned(&f.Data, int64(m.xxx_VESC_StatusAmpHoursChg4))
	return f
}

// MarshalFrame encodes the message as a CAN frame.
func (m *IOV_BackRightCornerMotor2) MarshalFrame() (can.Frame, error) {
	return m.Frame(), nil
}

// UnmarshalFrame decodes the message from a CAN frame.
func (m *IOV_BackRightCornerMotor2) UnmarshalFrame(f can.Frame) error {
	md := Messages().IOV_BackRightCornerMotor2
	switch {
	case f.ID != md.ID:
		return fmt.Errorf(
			"unmarshal IOV_BackRightCornerMotor2: expects ID 217116670 (got %s with ID %d)", f.String(), f.ID,
		)
	case f.Length != md.Length:
		return fmt.Errorf(
			"unmarshal IOV_BackRightCornerMotor2: expects length 8 (got %s with length %d)", f.String(), f.Length,
		)
	case f.IsRemote:
		return fmt.Errorf(
			"unmarshal IOV_BackRightCornerMotor2: expects non-remote frame (got remote frame %s)", f.String(),
		)
	case f.IsExtended != md.IsExtended:
		return fmt.Errorf(
			"unmarshal IOV_BackRightCornerMotor2: expects extended ID (got %s with standard ID)", f.String(),
		)
	}
	m.xxx_VESC_StatusAmpHours4 = int32(md.VESC_StatusAmpHours4.UnmarshalSigned(f.Data))
	m.xxx_VESC_StatusAmpHoursChg4 = int32(md.VESC_StatusAmpHoursChg4.UnmarshalSigned(f.Data))
	return nil
}

// IOV_BackRightCornerMotor3Reader provides read access to a IOV_BackRightCornerMotor3 message.
type IOV_BackRightCornerMotor3Reader interface {
	can.FrameMarshaler
	// VESC_StatusWattHours4 returns the value of the VESC_StatusWattHours4 signal.
	VESC_StatusWattHours4() int32
	// VESC_StatusWattHoursChg4 returns the value of the VESC_StatusWattHoursChg4 signal.
	VESC_StatusWattHoursChg4() int32
}

// IOV_BackRightCornerMotor3Writer provides write access to a IOV_BackRightCornerMotor3 message.
type IOV_BackRightCornerMotor3Writer interface {
	// CopyFrom copies all values from IOV_BackRightCornerMotor3.
	CopyFrom(IOV_BackRightCornerMotor3Reader) *IOV_BackRightCornerMotor3
	// SetVESC_StatusWattHours4 sets the value of the VESC_StatusWattHours4 signal.
	SetVESC_StatusWattHours4(int32) *IOV_BackRightCornerMotor3
	// SetVESC_StatusWattHoursChg4 sets the value of the VESC_StatusWattHoursChg4 signal.
	SetVESC_StatusWattHoursChg4(int32) *IOV_BackRightCornerMotor3
}

type IOV_BackRightCornerMotor3 struct {
	xxx_VESC_StatusWattHours4    int32
	xxx_VESC_StatusWattHoursChg4 int32
}

func NewIOV_BackRightCornerMotor3() *IOV_BackRightCornerMotor3 {
	m := &IOV_BackRightCornerMotor3{}
	m.Reset()
	return m
}

func (m *IOV_BackRightCornerMotor3) Reset() {
	m.xxx_VESC_StatusWattHours4 = 0
	m.xxx_VESC_StatusWattHoursChg4 = 0
}

func (m *IOV_BackRightCornerMotor3) CopyFrom(o IOV_BackRightCornerMotor3Reader) *IOV_BackRightCornerMotor3 {
	f, _ := o.MarshalFrame()
	_ = m.UnmarshalFrame(f)
	return m
}

// Descriptor returns the IOV_BackRightCornerMotor3 descriptor.
func (m *IOV_BackRightCornerMotor3) Descriptor() *descriptor.Message {
	return Messages().IOV_BackRightCornerMotor3.Message
}

// String returns a compact string representation of the message.
func (m *IOV_BackRightCornerMotor3) String() string {
	return cantext.MessageString(m)
}

func (m *IOV_BackRightCornerMotor3) VESC_StatusWattHours4() int32 {
	return m.xxx_VESC_StatusWattHours4
}

func (m *IOV_BackRightCornerMotor3) SetVESC_StatusWattHours4(v int32) *IOV_BackRightCornerMotor3 {
	m.xxx_VESC_StatusWattHours4 = int32(Messages().IOV_BackRightCornerMotor3.VESC_StatusWattHours4.SaturatedCastSigned(int64(v)))
	return m
}

func (m *IOV_BackRightCornerMotor3) VESC_StatusWattHoursChg4() int32 {
	return m.xxx_VESC_StatusWattHoursChg4
}

func (m *IOV_BackRightCornerMotor3) SetVESC_StatusWattHoursChg4(v int32) *IOV_BackRightCornerMotor3 {
	m.xxx_VESC_StatusWattHoursChg4 = int32(Messages().IOV_BackRightCornerMotor3.VESC_StatusWattHoursChg4.SaturatedCastSigned(int64(v)))
	return m
}

// Frame returns a CAN frame representing the message.
func (m *IOV_BackRightCornerMotor3) Frame() can.Frame {
	md := Messages().IOV_BackRightCornerMotor3
	f := can.Frame{ID: md.ID, IsExtended: md.IsExtended, Length: md.Length}
	md.VESC_StatusWattHours4.MarshalSigned(&f.Data, int64(m.xxx_VESC_StatusWattHours4))
	md.VESC_StatusWattHoursChg4.MarshalSigned(&f.Data, int64(m.xxx_VESC_StatusWattHoursChg4))
	return f
}

// MarshalFrame encodes the message as a CAN frame.
func (m *IOV_BackRightCornerMotor3) MarshalFrame() (can.Frame, error) {
	return m.Frame(), nil
}

// UnmarshalFrame decodes the message from a CAN frame.
func (m *IOV_BackRightCornerMotor3) UnmarshalFrame(f can.Frame) error {
	md := Messages().IOV_BackRightCornerMotor3
	switch {
	case f.ID != md.ID:
		return fmt.Errorf(
			"unmarshal IOV_BackRightCornerMotor3: expects ID 217120766 (got %s with ID %d)", f.String(), f.ID,
		)
	case f.Length != md.Length:
		return fmt.Errorf(
			"unmarshal IOV_BackRightCornerMotor3: expects length 8 (got %s with length %d)", f.String(), f.Length,
		)
	case f.IsRemote:
		return fmt.Errorf(
			"unmarshal IOV_BackRightCornerMotor3: expects non-remote frame (got remote frame %s)", f.String(),
		)
	case f.IsExtended != md.IsExtended:
		return fmt.Errorf(
			"unmarshal IOV_BackRightCornerMotor3: expects extended ID (got %s with standard ID)", f.String(),
		)
	}
	m.xxx_VESC_StatusWattHours4 = int32(md.VESC_StatusWattHours4.UnmarshalSigned(f.Data))
	m.xxx_VESC_StatusWattHoursChg4 = int32(md.VESC_StatusWattHoursChg4.UnmarshalSigned(f.Data))
	return nil
}

// IOV_BackRightCornerMotor4Reader provides read access to a IOV_BackRightCornerMotor4 message.
type IOV_BackRightCornerMotor4Reader interface {
	can.FrameMarshaler
	// VESC_StatusTempFET4 returns the value of the VESC_StatusTempFET4 signal.
	VESC_StatusTempFET4() int16
	// VESC_StatusTempMotor4 returns the value of the VESC_StatusTempMotor4 signal.
	VESC_StatusTempMotor4() int16
	// VESC_StatusCurrentIn4 returns the value of the VESC_StatusCurrentIn4 signal.
	VESC_StatusCurrentIn4() int16
	// VESC_StatusPIDPos4 returns the value of the VESC_StatusPIDPos4 signal.
	VESC_StatusPIDPos4() int16
}

// IOV_BackRightCornerMotor4Writer provides write access to a IOV_BackRightCornerMotor4 message.
type IOV_BackRightCornerMotor4Writer interface {
	// CopyFrom copies all values from IOV_BackRightCornerMotor4.
	CopyFrom(IOV_BackRightCornerMotor4Reader) *IOV_BackRightCornerMotor4
	// SetVESC_StatusTempFET4 sets the value of the VESC_StatusTempFET4 signal.
	SetVESC_StatusTempFET4(int16) *IOV_BackRightCornerMotor4
	// SetVESC_StatusTempMotor4 sets the value of the VESC_StatusTempMotor4 signal.
	SetVESC_StatusTempMotor4(int16) *IOV_BackRightCornerMotor4
	// SetVESC_StatusCurrentIn4 sets the value of the VESC_StatusCurrentIn4 signal.
	SetVESC_StatusCurrentIn4(int16) *IOV_BackRightCornerMotor4
	// SetVESC_StatusPIDPos4 sets the value of the VESC_StatusPIDPos4 signal.
	SetVESC_StatusPIDPos4(int16) *IOV_BackRightCornerMotor4
}

type IOV_BackRightCornerMotor4 struct {
	xxx_VESC_StatusTempFET4   int16
	xxx_VESC_StatusTempMotor4 int16
	xxx_VESC_StatusCurrentIn4 int16
	xxx_VESC_StatusPIDPos4    int16
}

func NewIOV_BackRightCornerMotor4() *IOV_BackRightCornerMotor4 {
	m := &IOV_BackRightCornerMotor4{}
	m.Reset()
	return m
}

func (m *IOV_BackRightCornerMotor4) Reset() {
	m.xxx_VESC_StatusTempFET4 = 0
	m.xxx_VESC_StatusTempMotor4 = 0
	m.xxx_VESC_StatusCurrentIn4 = 0
	m.xxx_VESC_StatusPIDPos4 = 0
}

func (m *IOV_BackRightCornerMotor4) CopyFrom(o IOV_BackRightCornerMotor4Reader) *IOV_BackRightCornerMotor4 {
	f, _ := o.MarshalFrame()
	_ = m.UnmarshalFrame(f)
	return m
}

// Descriptor returns the IOV_BackRightCornerMotor4 descriptor.
func (m *IOV_BackRightCornerMotor4) Descriptor() *descriptor.Message {
	return Messages().IOV_BackRightCornerMotor4.Message
}

// String returns a compact string representation of the message.
func (m *IOV_BackRightCornerMotor4) String() string {
	return cantext.MessageString(m)
}

func (m *IOV_BackRightCornerMotor4) VESC_StatusTempFET4() int16 {
	return m.xxx_VESC_StatusTempFET4
}

func (m *IOV_BackRightCornerMotor4) SetVESC_StatusTempFET4(v int16) *IOV_BackRightCornerMotor4 {
	m.xxx_VESC_StatusTempFET4 = int16(Messages().IOV_BackRightCornerMotor4.VESC_StatusTempFET4.SaturatedCastSigned(int64(v)))
	return m
}

func (m *IOV_BackRightCornerMotor4) VESC_StatusTempMotor4() int16 {
	return m.xxx_VESC_StatusTempMotor4
}

func (m *IOV_BackRightCornerMotor4) SetVESC_StatusTempMotor4(v int16) *IOV_BackRightCornerMotor4 {
	m.xxx_VESC_StatusTempMotor4 = int16(Messages().IOV_BackRightCornerMotor4.VESC_StatusTempMotor4.SaturatedCastSigned(int64(v)))
	return m
}

func (m *IOV_BackRightCornerMotor4) VESC_StatusCurrentIn4() int16 {
	return m.xxx_VESC_StatusCurrentIn4
}

func (m *IOV_BackRightCornerMotor4) SetVESC_StatusCurrentIn4(v int16) *IOV_BackRightCornerMotor4 {
	m.xxx_VESC_StatusCurrentIn4 = int16(Messages().IOV_BackRightCornerMotor4.VESC_StatusCurrentIn4.SaturatedCastSigned(int64(v)))
	return m
}

func (m *IOV_BackRightCornerMotor4) VESC_StatusPIDPos4() int16 {
	return m.xxx_VESC_StatusPIDPos4
}

func (m *IOV_BackRightCornerMotor4) SetVESC_StatusPIDPos4(v int16) *IOV_BackRightCornerMotor4 {
	m.xxx_VESC_StatusPIDPos4 = int16(Messages().IOV_BackRightCornerMotor4.VESC_StatusPIDPos4.SaturatedCastSigned(int64(v)))
	return m
}

// Frame returns a CAN frame representing the message.
func (m *IOV_BackRightCornerMotor4) Frame() can.Frame {
	md := Messages().IOV_BackRightCornerMotor4
	f := can.Frame{ID: md.ID, IsExtended: md.IsExtended, Length: md.Length}
	md.VESC_StatusTempFET4.MarshalSigned(&f.Data, int64(m.xxx_VESC_StatusTempFET4))
	md.VESC_StatusTempMotor4.MarshalSigned(&f.Data, int64(m.xxx_VESC_StatusTempMotor4))
	md.VESC_StatusCurrentIn4.MarshalSigned(&f.Data, int64(m.xxx_VESC_StatusCurrentIn4))
	md.VESC_StatusPIDPos4.MarshalSigned(&f.Data, int64(m.xxx_VESC_StatusPIDPos4))
	return f
}

// MarshalFrame encodes the message as a CAN frame.
func (m *IOV_BackRightCornerMotor4) MarshalFrame() (can.Frame, error) {
	return m.Frame(), nil
}

// UnmarshalFrame decodes the message from a CAN frame.
func (m *IOV_BackRightCornerMotor4) UnmarshalFrame(f can.Frame) error {
	md := Messages().IOV_BackRightCornerMotor4
	switch {
	case f.ID != md.ID:
		return fmt.Errorf(
			"unmarshal IOV_BackRightCornerMotor4: expects ID 217186302 (got %s with ID %d)", f.String(), f.ID,
		)
	case f.Length != md.Length:
		return fmt.Errorf(
			"unmarshal IOV_BackRightCornerMotor4: expects length 8 (got %s with length %d)", f.String(), f.Length,
		)
	case f.IsRemote:
		return fmt.Errorf(
			"unmarshal IOV_BackRightCornerMotor4: expects non-remote frame (got remote frame %s)", f.String(),
		)
	case f.IsExtended != md.IsExtended:
		return fmt.Errorf(
			"unmarshal IOV_BackRightCornerMotor4: expects extended ID (got %s with standard ID)", f.String(),
		)
	}
	m.xxx_VESC_StatusTempFET4 = int16(md.VESC_StatusTempFET4.UnmarshalSigned(f.Data))
	m.xxx_VESC_StatusTempMotor4 = int16(md.VESC_StatusTempMotor4.UnmarshalSigned(f.Data))
	m.xxx_VESC_StatusCurrentIn4 = int16(md.VESC_StatusCurrentIn4.UnmarshalSigned(f.Data))
	m.xxx_VESC_StatusPIDPos4 = int16(md.VESC_StatusPIDPos4.UnmarshalSigned(f.Data))
	return nil
}

// OBC_CommandReader provides read access to a OBC_Command message.
type OBC_CommandReader interface {
	can.FrameMarshaler
	// OBC_MaxAllwChargVolt returns the physical value of the OBC_MaxAllwChargVolt signal.
	OBC_MaxAllwChargVolt() float64
	// RawOBC_MaxAllwChargVolt returns the raw (encoded) value of the OBC_MaxAllwChargVolt signal.
	RawOBC_MaxAllwChargVolt() uint16
	// OBC_MaxAllowChargAmp returns the physical value of the OBC_MaxAllowChargAmp signal.
	OBC_MaxAllowChargAmp() float64
	// RawOBC_MaxAllowChargAmp returns the raw (encoded) value of the OBC_MaxAllowChargAmp signal.
	RawOBC_MaxAllowChargAmp() uint16
	// OBC_ControlWorkEnable returns the value of the OBC_ControlWorkEnable signal.
	OBC_ControlWorkEnable() OBC_Command_OBC_ControlWorkEnable
	// OBC_ControlOperatingMode returns the value of the OBC_ControlOperatingMode signal.
	OBC_ControlOperatingMode() OBC_Command_OBC_ControlOperatingMode
}

// OBC_CommandWriter provides write access to a OBC_Command message.
type OBC_CommandWriter interface {
	// CopyFrom copies all values from OBC_Command.
	CopyFrom(OBC_CommandReader) *OBC_Command
	// SetOBC_MaxAllwChargVolt sets the physical value of the OBC_MaxAllwChargVolt signal.
	SetOBC_MaxAllwChargVolt(float64) *OBC_Command
	// SetRawOBC_MaxAllwChargVolt sets the raw (encoded) value of the OBC_MaxAllwChargVolt signal.
	SetRawOBC_MaxAllwChargVolt(uint16) *OBC_Command
	// SetOBC_MaxAllowChargAmp sets the physical value of the OBC_MaxAllowChargAmp signal.
	SetOBC_MaxAllowChargAmp(float64) *OBC_Command
	// SetRawOBC_MaxAllowChargAmp sets the raw (encoded) value of the OBC_MaxAllowChargAmp signal.
	SetRawOBC_MaxAllowChargAmp(uint16) *OBC_Command
	// SetOBC_ControlWorkEnable sets the value of the OBC_ControlWorkEnable signal.
	SetOBC_ControlWorkEnable(OBC_Command_OBC_ControlWorkEnable) *OBC_Command
	// SetOBC_ControlOperatingMode sets the value of the OBC_ControlOperatingMode signal.
	SetOBC_ControlOperatingMode(OBC_Command_OBC_ControlOperatingMode) *OBC_Command
}

type OBC_Command struct {
	xxx_OBC_MaxAllwChargVolt     uint16
	xxx_OBC_MaxAllowChargAmp     uint16
	xxx_OBC_ControlWorkEnable    OBC_Command_OBC_ControlWorkEnable
	xxx_OBC_ControlOperatingMode OBC_Command_OBC_ControlOperatingMode
}

func NewOBC_Command() *OBC_Command {
	m := &OBC_Command{}
	m.Reset()
	return m
}

func (m *OBC_Command) Reset() {
	m.xxx_OBC_MaxAllwChargVolt = 0
	m.xxx_OBC_MaxAllowChargAmp = 0
	m.xxx_OBC_ControlWorkEnable = 0
	m.xxx_OBC_ControlOperatingMode = 0
}

func (m *OBC_Command) CopyFrom(o OBC_CommandReader) *OBC_Command {
	f, _ := o.MarshalFrame()
	_ = m.UnmarshalFrame(f)
	return m
}

// Descriptor returns the OBC_Command descriptor.
func (m *OBC_Command) Descriptor() *descriptor.Message {
	return Messages().OBC_Command.Message
}

// String returns a compact string representation of the message.
func (m *OBC_Command) String() string {
	return cantext.MessageString(m)
}

func (m *OBC_Command) OBC_MaxAllwChargVolt() float64 {
	return Messages().OBC_Command.OBC_MaxAllwChargVolt.ToPhysical(float64(m.xxx_OBC_MaxAllwChargVolt))
}

func (m *OBC_Command) SetOBC_MaxAllwChargVolt(v float64) *OBC_Command {
	m.xxx_OBC_MaxAllwChargVolt = uint16(Messages().OBC_Command.OBC_MaxAllwChargVolt.FromPhysical(v))
	return m
}

func (m *OBC_Command) RawOBC_MaxAllwChargVolt() uint16 {
	return m.xxx_OBC_MaxAllwChargVolt
}

func (m *OBC_Command) SetRawOBC_MaxAllwChargVolt(v uint16) *OBC_Command {
	m.xxx_OBC_MaxAllwChargVolt = uint16(Messages().OBC_Command.OBC_MaxAllwChargVolt.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBC_Command) OBC_MaxAllowChargAmp() float64 {
	return Messages().OBC_Command.OBC_MaxAllowChargAmp.ToPhysical(float64(m.xxx_OBC_MaxAllowChargAmp))
}

func (m *OBC_Command) SetOBC_MaxAllowChargAmp(v float64) *OBC_Command {
	m.xxx_OBC_MaxAllowChargAmp = uint16(Messages().OBC_Command.OBC_MaxAllowChargAmp.FromPhysical(v))
	return m
}

func (m *OBC_Command) RawOBC_MaxAllowChargAmp() uint16 {
	return m.xxx_OBC_MaxAllowChargAmp
}

func (m *OBC_Command) SetRawOBC_MaxAllowChargAmp(v uint16) *OBC_Command {
	m.xxx_OBC_MaxAllowChargAmp = uint16(Messages().OBC_Command.OBC_MaxAllowChargAmp.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBC_Command) OBC_ControlWorkEnable() OBC_Command_OBC_ControlWorkEnable {
	return m.xxx_OBC_ControlWorkEnable
}

func (m *OBC_Command) SetOBC_ControlWorkEnable(v OBC_Command_OBC_ControlWorkEnable) *OBC_Command {
	m.xxx_OBC_ControlWorkEnable = OBC_Command_OBC_ControlWorkEnable(Messages().OBC_Command.OBC_ControlWorkEnable.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBC_Command) OBC_ControlOperatingMode() OBC_Command_OBC_ControlOperatingMode {
	return m.xxx_OBC_ControlOperatingMode
}

func (m *OBC_Command) SetOBC_ControlOperatingMode(v OBC_Command_OBC_ControlOperatingMode) *OBC_Command {
	m.xxx_OBC_ControlOperatingMode = OBC_Command_OBC_ControlOperatingMode(Messages().OBC_Command.OBC_ControlOperatingMode.SaturatedCastUnsigned(uint64(v)))
	return m
}

// OBC_Command_OBC_ControlWorkEnable models the OBC_ControlWorkEnable signal of the OBC_Command message.
type OBC_Command_OBC_ControlWorkEnable uint8

// Value descriptions for the OBC_ControlWorkEnable signal of the OBC_Command message.
const (
	OBC_Command_OBC_ControlWorkEnable_ChargerIsStartingtoCharge OBC_Command_OBC_ControlWorkEnable = 0
	OBC_Command_OBC_ControlWorkEnable_ChargerCloseTheOutput     OBC_Command_OBC_ControlWorkEnable = 1
	OBC_Command_OBC_ControlWorkEnable_ChargeEnd                 OBC_Command_OBC_ControlWorkEnable = 2
)

func (v OBC_Command_OBC_ControlWorkEnable) String() string {
	switch v {
	case 0:
		return "ChargerIsStartingtoCharge"
	case 1:
		return "ChargerCloseTheOutput"
	case 2:
		return "ChargeEnd"
	default:
		return fmt.Sprintf("OBC_Command_OBC_ControlWorkEnable(%d)", v)
	}
}

// OBC_Command_OBC_ControlOperatingMode models the OBC_ControlOperatingMode signal of the OBC_Command message.
type OBC_Command_OBC_ControlOperatingMode uint8

// Value descriptions for the OBC_ControlOperatingMode signal of the OBC_Command message.
const (
	OBC_Command_OBC_ControlOperatingMode_ChargingMode OBC_Command_OBC_ControlOperatingMode = 0
	OBC_Command_OBC_ControlOperatingMode_HeatingModel OBC_Command_OBC_ControlOperatingMode = 1
)

func (v OBC_Command_OBC_ControlOperatingMode) String() string {
	switch v {
	case 0:
		return "ChargingMode"
	case 1:
		return "HeatingModel"
	default:
		return fmt.Sprintf("OBC_Command_OBC_ControlOperatingMode(%d)", v)
	}
}

// Frame returns a CAN frame representing the message.
func (m *OBC_Command) Frame() can.Frame {
	md := Messages().OBC_Command
	f := can.Frame{ID: md.ID, IsExtended: md.IsExtended, Length: md.Length}
	md.OBC_MaxAllwChargVolt.MarshalUnsigned(&f.Data, uint64(m.xxx_OBC_MaxAllwChargVolt))
	md.OBC_MaxAllowChargAmp.MarshalUnsigned(&f.Data, uint64(m.xxx_OBC_MaxAllowChargAmp))
	md.OBC_ControlWorkEnable.MarshalUnsigned(&f.Data, uint64(m.xxx_OBC_ControlWorkEnable))
	md.OBC_ControlOperatingMode.MarshalUnsigned(&f.Data, uint64(m.xxx_OBC_ControlOperatingMode))
	return f
}

// MarshalFrame encodes the message as a CAN frame.
func (m *OBC_Command) MarshalFrame() (can.Frame, error) {
	return m.Frame(), nil
}

// UnmarshalFrame decodes the message from a CAN frame.
func (m *OBC_Command) UnmarshalFrame(f can.Frame) error {
	md := Messages().OBC_Command
	switch {
	case f.ID != md.ID:
		return fmt.Errorf(
			"unmarshal OBC_Command: expects ID 403105268 (got %s with ID %d)", f.String(), f.ID,
		)
	case f.Length != md.Length:
		return fmt.Errorf(
			"unmarshal OBC_Command: expects length 8 (got %s with length %d)", f.String(), f.Length,
		)
	case f.IsRemote:
		return fmt.Errorf(
			"unmarshal OBC_Command: expects non-remote frame (got remote frame %s)", f.String(),
		)
	case f.IsExtended != md.IsExtended:
		return fmt.Errorf(
			"unmarshal OBC_Command: expects extended ID (got %s with standard ID)", f.String(),
		)
	}
	m.xxx_OBC_MaxAllwChargVolt = uint16(md.OBC_MaxAllwChargVolt.UnmarshalUnsigned(f.Data))
	m.xxx_OBC_MaxAllowChargAmp = uint16(md.OBC_MaxAllowChargAmp.UnmarshalUnsigned(f.Data))
	m.xxx_OBC_ControlWorkEnable = OBC_Command_OBC_ControlWorkEnable(md.OBC_ControlWorkEnable.UnmarshalUnsigned(f.Data))
	m.xxx_OBC_ControlOperatingMode = OBC_Command_OBC_ControlOperatingMode(md.OBC_ControlOperatingMode.UnmarshalUnsigned(f.Data))
	return nil
}

// BMS_DALY_SoCStatusReader provides read access to a BMS_DALY_SoCStatus message.
type BMS_DALY_SoCStatusReader interface {
	can.FrameMarshaler
	// DALY_AccumulatedPressure returns the physical value of the DALY_AccumulatedPressure signal.
	DALY_AccumulatedPressure() float64
	// RawDALY_AccumulatedPressure returns the raw (encoded) value of the DALY_AccumulatedPressure signal.
	RawDALY_AccumulatedPressure() uint16
	// DALY_CollectedTotalVoltage returns the physical value of the DALY_CollectedTotalVoltage signal.
	DALY_CollectedTotalVoltage() float64
	// RawDALY_CollectedTotalVoltage returns the raw (encoded) value of the DALY_CollectedTotalVoltage signal.
	RawDALY_CollectedTotalVoltage() uint16
	// DALY_BatteryCurrent returns the physical value of the DALY_BatteryCurrent signal.
	DALY_BatteryCurrent() float64
	// RawDALY_BatteryCurrent returns the raw (encoded) value of the DALY_BatteryCurrent signal.
	RawDALY_BatteryCurrent() uint16
	// DALY_BatterySoC returns the physical value of the DALY_BatterySoC signal.
	DALY_BatterySoC() float64
	// RawDALY_BatterySoC returns the raw (encoded) value of the DALY_BatterySoC signal.
	RawDALY_BatterySoC() uint16
}

// BMS_DALY_SoCStatusWriter provides write access to a BMS_DALY_SoCStatus message.
type BMS_DALY_SoCStatusWriter interface {
	// CopyFrom copies all values from BMS_DALY_SoCStatus.
	CopyFrom(BMS_DALY_SoCStatusReader) *BMS_DALY_SoCStatus
	// SetDALY_AccumulatedPressure sets the physical value of the DALY_AccumulatedPressure signal.
	SetDALY_AccumulatedPressure(float64) *BMS_DALY_SoCStatus
	// SetRawDALY_AccumulatedPressure sets the raw (encoded) value of the DALY_AccumulatedPressure signal.
	SetRawDALY_AccumulatedPressure(uint16) *BMS_DALY_SoCStatus
	// SetDALY_CollectedTotalVoltage sets the physical value of the DALY_CollectedTotalVoltage signal.
	SetDALY_CollectedTotalVoltage(float64) *BMS_DALY_SoCStatus
	// SetRawDALY_CollectedTotalVoltage sets the raw (encoded) value of the DALY_CollectedTotalVoltage signal.
	SetRawDALY_CollectedTotalVoltage(uint16) *BMS_DALY_SoCStatus
	// SetDALY_BatteryCurrent sets the physical value of the DALY_BatteryCurrent signal.
	SetDALY_BatteryCurrent(float64) *BMS_DALY_SoCStatus
	// SetRawDALY_BatteryCurrent sets the raw (encoded) value of the DALY_BatteryCurrent signal.
	SetRawDALY_BatteryCurrent(uint16) *BMS_DALY_SoCStatus
	// SetDALY_BatterySoC sets the physical value of the DALY_BatterySoC signal.
	SetDALY_BatterySoC(float64) *BMS_DALY_SoCStatus
	// SetRawDALY_BatterySoC sets the raw (encoded) value of the DALY_BatterySoC signal.
	SetRawDALY_BatterySoC(uint16) *BMS_DALY_SoCStatus
}

type BMS_DALY_SoCStatus struct {
	xxx_DALY_AccumulatedPressure   uint16
	xxx_DALY_CollectedTotalVoltage uint16
	xxx_DALY_BatteryCurrent        uint16
	xxx_DALY_BatterySoC            uint16
}

func NewBMS_DALY_SoCStatus() *BMS_DALY_SoCStatus {
	m := &BMS_DALY_SoCStatus{}
	m.Reset()
	return m
}

func (m *BMS_DALY_SoCStatus) Reset() {
	m.xxx_DALY_AccumulatedPressure = 0
	m.xxx_DALY_CollectedTotalVoltage = 0
	m.xxx_DALY_BatteryCurrent = 0
	m.xxx_DALY_BatterySoC = 0
}

func (m *BMS_DALY_SoCStatus) CopyFrom(o BMS_DALY_SoCStatusReader) *BMS_DALY_SoCStatus {
	f, _ := o.MarshalFrame()
	_ = m.UnmarshalFrame(f)
	return m
}

// Descriptor returns the BMS_DALY_SoCStatus descriptor.
func (m *BMS_DALY_SoCStatus) Descriptor() *descriptor.Message {
	return Messages().BMS_DALY_SoCStatus.Message
}

// String returns a compact string representation of the message.
func (m *BMS_DALY_SoCStatus) String() string {
	return cantext.MessageString(m)
}

func (m *BMS_DALY_SoCStatus) DALY_AccumulatedPressure() float64 {
	return Messages().BMS_DALY_SoCStatus.DALY_AccumulatedPressure.ToPhysical(float64(m.xxx_DALY_AccumulatedPressure))
}

func (m *BMS_DALY_SoCStatus) SetDALY_AccumulatedPressure(v float64) *BMS_DALY_SoCStatus {
	m.xxx_DALY_AccumulatedPressure = uint16(Messages().BMS_DALY_SoCStatus.DALY_AccumulatedPressure.FromPhysical(v))
	return m
}

func (m *BMS_DALY_SoCStatus) RawDALY_AccumulatedPressure() uint16 {
	return m.xxx_DALY_AccumulatedPressure
}

func (m *BMS_DALY_SoCStatus) SetRawDALY_AccumulatedPressure(v uint16) *BMS_DALY_SoCStatus {
	m.xxx_DALY_AccumulatedPressure = uint16(Messages().BMS_DALY_SoCStatus.DALY_AccumulatedPressure.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *BMS_DALY_SoCStatus) DALY_CollectedTotalVoltage() float64 {
	return Messages().BMS_DALY_SoCStatus.DALY_CollectedTotalVoltage.ToPhysical(float64(m.xxx_DALY_CollectedTotalVoltage))
}

func (m *BMS_DALY_SoCStatus) SetDALY_CollectedTotalVoltage(v float64) *BMS_DALY_SoCStatus {
	m.xxx_DALY_CollectedTotalVoltage = uint16(Messages().BMS_DALY_SoCStatus.DALY_CollectedTotalVoltage.FromPhysical(v))
	return m
}

func (m *BMS_DALY_SoCStatus) RawDALY_CollectedTotalVoltage() uint16 {
	return m.xxx_DALY_CollectedTotalVoltage
}

func (m *BMS_DALY_SoCStatus) SetRawDALY_CollectedTotalVoltage(v uint16) *BMS_DALY_SoCStatus {
	m.xxx_DALY_CollectedTotalVoltage = uint16(Messages().BMS_DALY_SoCStatus.DALY_CollectedTotalVoltage.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *BMS_DALY_SoCStatus) DALY_BatteryCurrent() float64 {
	return Messages().BMS_DALY_SoCStatus.DALY_BatteryCurrent.ToPhysical(float64(m.xxx_DALY_BatteryCurrent))
}

func (m *BMS_DALY_SoCStatus) SetDALY_BatteryCurrent(v float64) *BMS_DALY_SoCStatus {
	m.xxx_DALY_BatteryCurrent = uint16(Messages().BMS_DALY_SoCStatus.DALY_BatteryCurrent.FromPhysical(v))
	return m
}

func (m *BMS_DALY_SoCStatus) RawDALY_BatteryCurrent() uint16 {
	return m.xxx_DALY_BatteryCurrent
}

func (m *BMS_DALY_SoCStatus) SetRawDALY_BatteryCurrent(v uint16) *BMS_DALY_SoCStatus {
	m.xxx_DALY_BatteryCurrent = uint16(Messages().BMS_DALY_SoCStatus.DALY_BatteryCurrent.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *BMS_DALY_SoCStatus) DALY_BatterySoC() float64 {
	return Messages().BMS_DALY_SoCStatus.DALY_BatterySoC.ToPhysical(float64(m.xxx_DALY_BatterySoC))
}

func (m *BMS_DALY_SoCStatus) SetDALY_BatterySoC(v float64) *BMS_DALY_SoCStatus {
	m.xxx_DALY_BatterySoC = uint16(Messages().BMS_DALY_SoCStatus.DALY_BatterySoC.FromPhysical(v))
	return m
}

func (m *BMS_DALY_SoCStatus) RawDALY_BatterySoC() uint16 {
	return m.xxx_DALY_BatterySoC
}

func (m *BMS_DALY_SoCStatus) SetRawDALY_BatterySoC(v uint16) *BMS_DALY_SoCStatus {
	m.xxx_DALY_BatterySoC = uint16(Messages().BMS_DALY_SoCStatus.DALY_BatterySoC.SaturatedCastUnsigned(uint64(v)))
	return m
}

// Frame returns a CAN frame representing the message.
func (m *BMS_DALY_SoCStatus) Frame() can.Frame {
	md := Messages().BMS_DALY_SoCStatus
	f := can.Frame{ID: md.ID, IsExtended: md.IsExtended, Length: md.Length}
	md.DALY_AccumulatedPressure.MarshalUnsigned(&f.Data, uint64(m.xxx_DALY_AccumulatedPressure))
	md.DALY_CollectedTotalVoltage.MarshalUnsigned(&f.Data, uint64(m.xxx_DALY_CollectedTotalVoltage))
	md.DALY_BatteryCurrent.MarshalUnsigned(&f.Data, uint64(m.xxx_DALY_BatteryCurrent))
	md.DALY_BatterySoC.MarshalUnsigned(&f.Data, uint64(m.xxx_DALY_BatterySoC))
	return f
}

// MarshalFrame encodes the message as a CAN frame.
func (m *BMS_DALY_SoCStatus) MarshalFrame() (can.Frame, error) {
	return m.Frame(), nil
}

// UnmarshalFrame decodes the message from a CAN frame.
func (m *BMS_DALY_SoCStatus) UnmarshalFrame(f can.Frame) error {
	md := Messages().BMS_DALY_SoCStatus
	switch {
	case f.ID != md.ID:
		return fmt.Errorf(
			"unmarshal BMS_DALY_SoCStatus: expects ID 412106753 (got %s with ID %d)", f.String(), f.ID,
		)
	case f.Length != md.Length:
		return fmt.Errorf(
			"unmarshal BMS_DALY_SoCStatus: expects length 8 (got %s with length %d)", f.String(), f.Length,
		)
	case f.IsRemote:
		return fmt.Errorf(
			"unmarshal BMS_DALY_SoCStatus: expects non-remote frame (got remote frame %s)", f.String(),
		)
	case f.IsExtended != md.IsExtended:
		return fmt.Errorf(
			"unmarshal BMS_DALY_SoCStatus: expects extended ID (got %s with standard ID)", f.String(),
		)
	}
	m.xxx_DALY_AccumulatedPressure = uint16(md.DALY_AccumulatedPressure.UnmarshalUnsigned(f.Data))
	m.xxx_DALY_CollectedTotalVoltage = uint16(md.DALY_CollectedTotalVoltage.UnmarshalUnsigned(f.Data))
	m.xxx_DALY_BatteryCurrent = uint16(md.DALY_BatteryCurrent.UnmarshalUnsigned(f.Data))
	m.xxx_DALY_BatterySoC = uint16(md.DALY_BatterySoC.UnmarshalUnsigned(f.Data))
	return nil
}

// BMS_DALY_RangeVoltageReader provides read access to a BMS_DALY_RangeVoltage message.
type BMS_DALY_RangeVoltageReader interface {
	can.FrameMarshaler
	// DALY_MaxMonomerVoltage returns the value of the DALY_MaxMonomerVoltage signal.
	DALY_MaxMonomerVoltage() uint16
	// DALY_MaxUnitVoltageCellNo returns the value of the DALY_MaxUnitVoltageCellNo signal.
	DALY_MaxUnitVoltageCellNo() uint8
	// DALY_MinMonomerVoltage returns the value of the DALY_MinMonomerVoltage signal.
	DALY_MinMonomerVoltage() uint16
	// DALY_MinUnitVoltageCellNo returns the value of the DALY_MinUnitVoltageCellNo signal.
	DALY_MinUnitVoltageCellNo() uint8
}

// BMS_DALY_RangeVoltageWriter provides write access to a BMS_DALY_RangeVoltage message.
type BMS_DALY_RangeVoltageWriter interface {
	// CopyFrom copies all values from BMS_DALY_RangeVoltage.
	CopyFrom(BMS_DALY_RangeVoltageReader) *BMS_DALY_RangeVoltage
	// SetDALY_MaxMonomerVoltage sets the value of the DALY_MaxMonomerVoltage signal.
	SetDALY_MaxMonomerVoltage(uint16) *BMS_DALY_RangeVoltage
	// SetDALY_MaxUnitVoltageCellNo sets the value of the DALY_MaxUnitVoltageCellNo signal.
	SetDALY_MaxUnitVoltageCellNo(uint8) *BMS_DALY_RangeVoltage
	// SetDALY_MinMonomerVoltage sets the value of the DALY_MinMonomerVoltage signal.
	SetDALY_MinMonomerVoltage(uint16) *BMS_DALY_RangeVoltage
	// SetDALY_MinUnitVoltageCellNo sets the value of the DALY_MinUnitVoltageCellNo signal.
	SetDALY_MinUnitVoltageCellNo(uint8) *BMS_DALY_RangeVoltage
}

type BMS_DALY_RangeVoltage struct {
	xxx_DALY_MaxMonomerVoltage    uint16
	xxx_DALY_MaxUnitVoltageCellNo uint8
	xxx_DALY_MinMonomerVoltage    uint16
	xxx_DALY_MinUnitVoltageCellNo uint8
}

func NewBMS_DALY_RangeVoltage() *BMS_DALY_RangeVoltage {
	m := &BMS_DALY_RangeVoltage{}
	m.Reset()
	return m
}

func (m *BMS_DALY_RangeVoltage) Reset() {
	m.xxx_DALY_MaxMonomerVoltage = 0
	m.xxx_DALY_MaxUnitVoltageCellNo = 0
	m.xxx_DALY_MinMonomerVoltage = 0
	m.xxx_DALY_MinUnitVoltageCellNo = 0
}

func (m *BMS_DALY_RangeVoltage) CopyFrom(o BMS_DALY_RangeVoltageReader) *BMS_DALY_RangeVoltage {
	f, _ := o.MarshalFrame()
	_ = m.UnmarshalFrame(f)
	return m
}

// Descriptor returns the BMS_DALY_RangeVoltage descriptor.
func (m *BMS_DALY_RangeVoltage) Descriptor() *descriptor.Message {
	return Messages().BMS_DALY_RangeVoltage.Message
}

// String returns a compact string representation of the message.
func (m *BMS_DALY_RangeVoltage) String() string {
	return cantext.MessageString(m)
}

func (m *BMS_DALY_RangeVoltage) DALY_MaxMonomerVoltage() uint16 {
	return m.xxx_DALY_MaxMonomerVoltage
}

func (m *BMS_DALY_RangeVoltage) SetDALY_MaxMonomerVoltage(v uint16) *BMS_DALY_RangeVoltage {
	m.xxx_DALY_MaxMonomerVoltage = uint16(Messages().BMS_DALY_RangeVoltage.DALY_MaxMonomerVoltage.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *BMS_DALY_RangeVoltage) DALY_MaxUnitVoltageCellNo() uint8 {
	return m.xxx_DALY_MaxUnitVoltageCellNo
}

func (m *BMS_DALY_RangeVoltage) SetDALY_MaxUnitVoltageCellNo(v uint8) *BMS_DALY_RangeVoltage {
	m.xxx_DALY_MaxUnitVoltageCellNo = uint8(Messages().BMS_DALY_RangeVoltage.DALY_MaxUnitVoltageCellNo.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *BMS_DALY_RangeVoltage) DALY_MinMonomerVoltage() uint16 {
	return m.xxx_DALY_MinMonomerVoltage
}

func (m *BMS_DALY_RangeVoltage) SetDALY_MinMonomerVoltage(v uint16) *BMS_DALY_RangeVoltage {
	m.xxx_DALY_MinMonomerVoltage = uint16(Messages().BMS_DALY_RangeVoltage.DALY_MinMonomerVoltage.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *BMS_DALY_RangeVoltage) DALY_MinUnitVoltageCellNo() uint8 {
	return m.xxx_DALY_MinUnitVoltageCellNo
}

func (m *BMS_DALY_RangeVoltage) SetDALY_MinUnitVoltageCellNo(v uint8) *BMS_DALY_RangeVoltage {
	m.xxx_DALY_MinUnitVoltageCellNo = uint8(Messages().BMS_DALY_RangeVoltage.DALY_MinUnitVoltageCellNo.SaturatedCastUnsigned(uint64(v)))
	return m
}

// Frame returns a CAN frame representing the message.
func (m *BMS_DALY_RangeVoltage) Frame() can.Frame {
	md := Messages().BMS_DALY_RangeVoltage
	f := can.Frame{ID: md.ID, IsExtended: md.IsExtended, Length: md.Length}
	md.DALY_MaxMonomerVoltage.MarshalUnsigned(&f.Data, uint64(m.xxx_DALY_MaxMonomerVoltage))
	md.DALY_MaxUnitVoltageCellNo.MarshalUnsigned(&f.Data, uint64(m.xxx_DALY_MaxUnitVoltageCellNo))
	md.DALY_MinMonomerVoltage.MarshalUnsigned(&f.Data, uint64(m.xxx_DALY_MinMonomerVoltage))
	md.DALY_MinUnitVoltageCellNo.MarshalUnsigned(&f.Data, uint64(m.xxx_DALY_MinUnitVoltageCellNo))
	return f
}

// MarshalFrame encodes the message as a CAN frame.
func (m *BMS_DALY_RangeVoltage) MarshalFrame() (can.Frame, error) {
	return m.Frame(), nil
}

// UnmarshalFrame decodes the message from a CAN frame.
func (m *BMS_DALY_RangeVoltage) UnmarshalFrame(f can.Frame) error {
	md := Messages().BMS_DALY_RangeVoltage
	switch {
	case f.ID != md.ID:
		return fmt.Errorf(
			"unmarshal BMS_DALY_RangeVoltage: expects ID 412172289 (got %s with ID %d)", f.String(), f.ID,
		)
	case f.Length != md.Length:
		return fmt.Errorf(
			"unmarshal BMS_DALY_RangeVoltage: expects length 8 (got %s with length %d)", f.String(), f.Length,
		)
	case f.IsRemote:
		return fmt.Errorf(
			"unmarshal BMS_DALY_RangeVoltage: expects non-remote frame (got remote frame %s)", f.String(),
		)
	case f.IsExtended != md.IsExtended:
		return fmt.Errorf(
			"unmarshal BMS_DALY_RangeVoltage: expects extended ID (got %s with standard ID)", f.String(),
		)
	}
	m.xxx_DALY_MaxMonomerVoltage = uint16(md.DALY_MaxMonomerVoltage.UnmarshalUnsigned(f.Data))
	m.xxx_DALY_MaxUnitVoltageCellNo = uint8(md.DALY_MaxUnitVoltageCellNo.UnmarshalUnsigned(f.Data))
	m.xxx_DALY_MinMonomerVoltage = uint16(md.DALY_MinMonomerVoltage.UnmarshalUnsigned(f.Data))
	m.xxx_DALY_MinUnitVoltageCellNo = uint8(md.DALY_MinUnitVoltageCellNo.UnmarshalUnsigned(f.Data))
	return nil
}

// BMS_DALY_RangeTemperatureReader provides read access to a BMS_DALY_RangeTemperature message.
type BMS_DALY_RangeTemperatureReader interface {
	can.FrameMarshaler
	// DALY_MaxMonomerTemp returns the physical value of the DALY_MaxMonomerTemp signal.
	DALY_MaxMonomerTemp() float64
	// RawDALY_MaxMonomerTemp returns the raw (encoded) value of the DALY_MaxMonomerTemp signal.
	RawDALY_MaxMonomerTemp() uint8
	// DALY_MaxUnitTempCellNo returns the value of the DALY_MaxUnitTempCellNo signal.
	DALY_MaxUnitTempCellNo() uint8
	// DALY_MinMonomerTemp returns the physical value of the DALY_MinMonomerTemp signal.
	DALY_MinMonomerTemp() float64
	// RawDALY_MinMonomerTemp returns the raw (encoded) value of the DALY_MinMonomerTemp signal.
	RawDALY_MinMonomerTemp() uint8
	// DALY_MinUnitTempCellNo returns the value of the DALY_MinUnitTempCellNo signal.
	DALY_MinUnitTempCellNo() uint8
}

// BMS_DALY_RangeTemperatureWriter provides write access to a BMS_DALY_RangeTemperature message.
type BMS_DALY_RangeTemperatureWriter interface {
	// CopyFrom copies all values from BMS_DALY_RangeTemperature.
	CopyFrom(BMS_DALY_RangeTemperatureReader) *BMS_DALY_RangeTemperature
	// SetDALY_MaxMonomerTemp sets the physical value of the DALY_MaxMonomerTemp signal.
	SetDALY_MaxMonomerTemp(float64) *BMS_DALY_RangeTemperature
	// SetRawDALY_MaxMonomerTemp sets the raw (encoded) value of the DALY_MaxMonomerTemp signal.
	SetRawDALY_MaxMonomerTemp(uint8) *BMS_DALY_RangeTemperature
	// SetDALY_MaxUnitTempCellNo sets the value of the DALY_MaxUnitTempCellNo signal.
	SetDALY_MaxUnitTempCellNo(uint8) *BMS_DALY_RangeTemperature
	// SetDALY_MinMonomerTemp sets the physical value of the DALY_MinMonomerTemp signal.
	SetDALY_MinMonomerTemp(float64) *BMS_DALY_RangeTemperature
	// SetRawDALY_MinMonomerTemp sets the raw (encoded) value of the DALY_MinMonomerTemp signal.
	SetRawDALY_MinMonomerTemp(uint8) *BMS_DALY_RangeTemperature
	// SetDALY_MinUnitTempCellNo sets the value of the DALY_MinUnitTempCellNo signal.
	SetDALY_MinUnitTempCellNo(uint8) *BMS_DALY_RangeTemperature
}

type BMS_DALY_RangeTemperature struct {
	xxx_DALY_MaxMonomerTemp    uint8
	xxx_DALY_MaxUnitTempCellNo uint8
	xxx_DALY_MinMonomerTemp    uint8
	xxx_DALY_MinUnitTempCellNo uint8
}

func NewBMS_DALY_RangeTemperature() *BMS_DALY_RangeTemperature {
	m := &BMS_DALY_RangeTemperature{}
	m.Reset()
	return m
}

func (m *BMS_DALY_RangeTemperature) Reset() {
	m.xxx_DALY_MaxMonomerTemp = 0
	m.xxx_DALY_MaxUnitTempCellNo = 0
	m.xxx_DALY_MinMonomerTemp = 0
	m.xxx_DALY_MinUnitTempCellNo = 0
}

func (m *BMS_DALY_RangeTemperature) CopyFrom(o BMS_DALY_RangeTemperatureReader) *BMS_DALY_RangeTemperature {
	f, _ := o.MarshalFrame()
	_ = m.UnmarshalFrame(f)
	return m
}

// Descriptor returns the BMS_DALY_RangeTemperature descriptor.
func (m *BMS_DALY_RangeTemperature) Descriptor() *descriptor.Message {
	return Messages().BMS_DALY_RangeTemperature.Message
}

// String returns a compact string representation of the message.
func (m *BMS_DALY_RangeTemperature) String() string {
	return cantext.MessageString(m)
}

func (m *BMS_DALY_RangeTemperature) DALY_MaxMonomerTemp() float64 {
	return Messages().BMS_DALY_RangeTemperature.DALY_MaxMonomerTemp.ToPhysical(float64(m.xxx_DALY_MaxMonomerTemp))
}

func (m *BMS_DALY_RangeTemperature) SetDALY_MaxMonomerTemp(v float64) *BMS_DALY_RangeTemperature {
	m.xxx_DALY_MaxMonomerTemp = uint8(Messages().BMS_DALY_RangeTemperature.DALY_MaxMonomerTemp.FromPhysical(v))
	return m
}

func (m *BMS_DALY_RangeTemperature) RawDALY_MaxMonomerTemp() uint8 {
	return m.xxx_DALY_MaxMonomerTemp
}

func (m *BMS_DALY_RangeTemperature) SetRawDALY_MaxMonomerTemp(v uint8) *BMS_DALY_RangeTemperature {
	m.xxx_DALY_MaxMonomerTemp = uint8(Messages().BMS_DALY_RangeTemperature.DALY_MaxMonomerTemp.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *BMS_DALY_RangeTemperature) DALY_MaxUnitTempCellNo() uint8 {
	return m.xxx_DALY_MaxUnitTempCellNo
}

func (m *BMS_DALY_RangeTemperature) SetDALY_MaxUnitTempCellNo(v uint8) *BMS_DALY_RangeTemperature {
	m.xxx_DALY_MaxUnitTempCellNo = uint8(Messages().BMS_DALY_RangeTemperature.DALY_MaxUnitTempCellNo.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *BMS_DALY_RangeTemperature) DALY_MinMonomerTemp() float64 {
	return Messages().BMS_DALY_RangeTemperature.DALY_MinMonomerTemp.ToPhysical(float64(m.xxx_DALY_MinMonomerTemp))
}

func (m *BMS_DALY_RangeTemperature) SetDALY_MinMonomerTemp(v float64) *BMS_DALY_RangeTemperature {
	m.xxx_DALY_MinMonomerTemp = uint8(Messages().BMS_DALY_RangeTemperature.DALY_MinMonomerTemp.FromPhysical(v))
	return m
}

func (m *BMS_DALY_RangeTemperature) RawDALY_MinMonomerTemp() uint8 {
	return m.xxx_DALY_MinMonomerTemp
}

func (m *BMS_DALY_RangeTemperature) SetRawDALY_MinMonomerTemp(v uint8) *BMS_DALY_RangeTemperature {
	m.xxx_DALY_MinMonomerTemp = uint8(Messages().BMS_DALY_RangeTemperature.DALY_MinMonomerTemp.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *BMS_DALY_RangeTemperature) DALY_MinUnitTempCellNo() uint8 {
	return m.xxx_DALY_MinUnitTempCellNo
}

func (m *BMS_DALY_RangeTemperature) SetDALY_MinUnitTempCellNo(v uint8) *BMS_DALY_RangeTemperature {
	m.xxx_DALY_MinUnitTempCellNo = uint8(Messages().BMS_DALY_RangeTemperature.DALY_MinUnitTempCellNo.SaturatedCastUnsigned(uint64(v)))
	return m
}

// Frame returns a CAN frame representing the message.
func (m *BMS_DALY_RangeTemperature) Frame() can.Frame {
	md := Messages().BMS_DALY_RangeTemperature
	f := can.Frame{ID: md.ID, IsExtended: md.IsExtended, Length: md.Length}
	md.DALY_MaxMonomerTemp.MarshalUnsigned(&f.Data, uint64(m.xxx_DALY_MaxMonomerTemp))
	md.DALY_MaxUnitTempCellNo.MarshalUnsigned(&f.Data, uint64(m.xxx_DALY_MaxUnitTempCellNo))
	md.DALY_MinMonomerTemp.MarshalUnsigned(&f.Data, uint64(m.xxx_DALY_MinMonomerTemp))
	md.DALY_MinUnitTempCellNo.MarshalUnsigned(&f.Data, uint64(m.xxx_DALY_MinUnitTempCellNo))
	return f
}

// MarshalFrame encodes the message as a CAN frame.
func (m *BMS_DALY_RangeTemperature) MarshalFrame() (can.Frame, error) {
	return m.Frame(), nil
}

// UnmarshalFrame decodes the message from a CAN frame.
func (m *BMS_DALY_RangeTemperature) UnmarshalFrame(f can.Frame) error {
	md := Messages().BMS_DALY_RangeTemperature
	switch {
	case f.ID != md.ID:
		return fmt.Errorf(
			"unmarshal BMS_DALY_RangeTemperature: expects ID 412237825 (got %s with ID %d)", f.String(), f.ID,
		)
	case f.Length != md.Length:
		return fmt.Errorf(
			"unmarshal BMS_DALY_RangeTemperature: expects length 8 (got %s with length %d)", f.String(), f.Length,
		)
	case f.IsRemote:
		return fmt.Errorf(
			"unmarshal BMS_DALY_RangeTemperature: expects non-remote frame (got remote frame %s)", f.String(),
		)
	case f.IsExtended != md.IsExtended:
		return fmt.Errorf(
			"unmarshal BMS_DALY_RangeTemperature: expects extended ID (got %s with standard ID)", f.String(),
		)
	}
	m.xxx_DALY_MaxMonomerTemp = uint8(md.DALY_MaxMonomerTemp.UnmarshalUnsigned(f.Data))
	m.xxx_DALY_MaxUnitTempCellNo = uint8(md.DALY_MaxUnitTempCellNo.UnmarshalUnsigned(f.Data))
	m.xxx_DALY_MinMonomerTemp = uint8(md.DALY_MinMonomerTemp.UnmarshalUnsigned(f.Data))
	m.xxx_DALY_MinUnitTempCellNo = uint8(md.DALY_MinUnitTempCellNo.UnmarshalUnsigned(f.Data))
	return nil
}

// BMS_DALY_MOSStatusReader provides read access to a BMS_DALY_MOSStatus message.
type BMS_DALY_MOSStatusReader interface {
	can.FrameMarshaler
	// DALY_DischargeStatus returns the value of the DALY_DischargeStatus signal.
	DALY_DischargeStatus() BMS_DALY_MOSStatus_DALY_DischargeStatus
	// DALY_ChargingMOSTube returns the value of the DALY_ChargingMOSTube signal.
	DALY_ChargingMOSTube() uint8
	// DALY_DischargeMOSTube returns the value of the DALY_DischargeMOSTube signal.
	DALY_DischargeMOSTube() uint8
	// DALY_BMSLife returns the value of the DALY_BMSLife signal.
	DALY_BMSLife() uint8
	// DALY_ResidualCapacity returns the value of the DALY_ResidualCapacity signal.
	DALY_ResidualCapacity() uint32
}

// BMS_DALY_MOSStatusWriter provides write access to a BMS_DALY_MOSStatus message.
type BMS_DALY_MOSStatusWriter interface {
	// CopyFrom copies all values from BMS_DALY_MOSStatus.
	CopyFrom(BMS_DALY_MOSStatusReader) *BMS_DALY_MOSStatus
	// SetDALY_DischargeStatus sets the value of the DALY_DischargeStatus signal.
	SetDALY_DischargeStatus(BMS_DALY_MOSStatus_DALY_DischargeStatus) *BMS_DALY_MOSStatus
	// SetDALY_ChargingMOSTube sets the value of the DALY_ChargingMOSTube signal.
	SetDALY_ChargingMOSTube(uint8) *BMS_DALY_MOSStatus
	// SetDALY_DischargeMOSTube sets the value of the DALY_DischargeMOSTube signal.
	SetDALY_DischargeMOSTube(uint8) *BMS_DALY_MOSStatus
	// SetDALY_BMSLife sets the value of the DALY_BMSLife signal.
	SetDALY_BMSLife(uint8) *BMS_DALY_MOSStatus
	// SetDALY_ResidualCapacity sets the value of the DALY_ResidualCapacity signal.
	SetDALY_ResidualCapacity(uint32) *BMS_DALY_MOSStatus
}

type BMS_DALY_MOSStatus struct {
	xxx_DALY_DischargeStatus  BMS_DALY_MOSStatus_DALY_DischargeStatus
	xxx_DALY_ChargingMOSTube  uint8
	xxx_DALY_DischargeMOSTube uint8
	xxx_DALY_BMSLife          uint8
	xxx_DALY_ResidualCapacity uint32
}

func NewBMS_DALY_MOSStatus() *BMS_DALY_MOSStatus {
	m := &BMS_DALY_MOSStatus{}
	m.Reset()
	return m
}

func (m *BMS_DALY_MOSStatus) Reset() {
	m.xxx_DALY_DischargeStatus = 0
	m.xxx_DALY_ChargingMOSTube = 0
	m.xxx_DALY_DischargeMOSTube = 0
	m.xxx_DALY_BMSLife = 0
	m.xxx_DALY_ResidualCapacity = 0
}

func (m *BMS_DALY_MOSStatus) CopyFrom(o BMS_DALY_MOSStatusReader) *BMS_DALY_MOSStatus {
	f, _ := o.MarshalFrame()
	_ = m.UnmarshalFrame(f)
	return m
}

// Descriptor returns the BMS_DALY_MOSStatus descriptor.
func (m *BMS_DALY_MOSStatus) Descriptor() *descriptor.Message {
	return Messages().BMS_DALY_MOSStatus.Message
}

// String returns a compact string representation of the message.
func (m *BMS_DALY_MOSStatus) String() string {
	return cantext.MessageString(m)
}

func (m *BMS_DALY_MOSStatus) DALY_DischargeStatus() BMS_DALY_MOSStatus_DALY_DischargeStatus {
	return m.xxx_DALY_DischargeStatus
}

func (m *BMS_DALY_MOSStatus) SetDALY_DischargeStatus(v BMS_DALY_MOSStatus_DALY_DischargeStatus) *BMS_DALY_MOSStatus {
	m.xxx_DALY_DischargeStatus = BMS_DALY_MOSStatus_DALY_DischargeStatus(Messages().BMS_DALY_MOSStatus.DALY_DischargeStatus.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *BMS_DALY_MOSStatus) DALY_ChargingMOSTube() uint8 {
	return m.xxx_DALY_ChargingMOSTube
}

func (m *BMS_DALY_MOSStatus) SetDALY_ChargingMOSTube(v uint8) *BMS_DALY_MOSStatus {
	m.xxx_DALY_ChargingMOSTube = uint8(Messages().BMS_DALY_MOSStatus.DALY_ChargingMOSTube.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *BMS_DALY_MOSStatus) DALY_DischargeMOSTube() uint8 {
	return m.xxx_DALY_DischargeMOSTube
}

func (m *BMS_DALY_MOSStatus) SetDALY_DischargeMOSTube(v uint8) *BMS_DALY_MOSStatus {
	m.xxx_DALY_DischargeMOSTube = uint8(Messages().BMS_DALY_MOSStatus.DALY_DischargeMOSTube.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *BMS_DALY_MOSStatus) DALY_BMSLife() uint8 {
	return m.xxx_DALY_BMSLife
}

func (m *BMS_DALY_MOSStatus) SetDALY_BMSLife(v uint8) *BMS_DALY_MOSStatus {
	m.xxx_DALY_BMSLife = uint8(Messages().BMS_DALY_MOSStatus.DALY_BMSLife.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *BMS_DALY_MOSStatus) DALY_ResidualCapacity() uint32 {
	return m.xxx_DALY_ResidualCapacity
}

func (m *BMS_DALY_MOSStatus) SetDALY_ResidualCapacity(v uint32) *BMS_DALY_MOSStatus {
	m.xxx_DALY_ResidualCapacity = uint32(Messages().BMS_DALY_MOSStatus.DALY_ResidualCapacity.SaturatedCastUnsigned(uint64(v)))
	return m
}

// BMS_DALY_MOSStatus_DALY_DischargeStatus models the DALY_DischargeStatus signal of the BMS_DALY_MOSStatus message.
type BMS_DALY_MOSStatus_DALY_DischargeStatus uint8

// Value descriptions for the DALY_DischargeStatus signal of the BMS_DALY_MOSStatus message.
const (
	BMS_DALY_MOSStatus_DALY_DischargeStatus_Stationary BMS_DALY_MOSStatus_DALY_DischargeStatus = 0
	BMS_DALY_MOSStatus_DALY_DischargeStatus_Charged    BMS_DALY_MOSStatus_DALY_DischargeStatus = 1
	BMS_DALY_MOSStatus_DALY_DischargeStatus_Discharged BMS_DALY_MOSStatus_DALY_DischargeStatus = 2
)

func (v BMS_DALY_MOSStatus_DALY_DischargeStatus) String() string {
	switch v {
	case 0:
		return "Stationary"
	case 1:
		return "Charged"
	case 2:
		return "Discharged"
	default:
		return fmt.Sprintf("BMS_DALY_MOSStatus_DALY_DischargeStatus(%d)", v)
	}
}

// Frame returns a CAN frame representing the message.
func (m *BMS_DALY_MOSStatus) Frame() can.Frame {
	md := Messages().BMS_DALY_MOSStatus
	f := can.Frame{ID: md.ID, IsExtended: md.IsExtended, Length: md.Length}
	md.DALY_DischargeStatus.MarshalUnsigned(&f.Data, uint64(m.xxx_DALY_DischargeStatus))
	md.DALY_ChargingMOSTube.MarshalUnsigned(&f.Data, uint64(m.xxx_DALY_ChargingMOSTube))
	md.DALY_DischargeMOSTube.MarshalUnsigned(&f.Data, uint64(m.xxx_DALY_DischargeMOSTube))
	md.DALY_BMSLife.MarshalUnsigned(&f.Data, uint64(m.xxx_DALY_BMSLife))
	md.DALY_ResidualCapacity.MarshalUnsigned(&f.Data, uint64(m.xxx_DALY_ResidualCapacity))
	return f
}

// MarshalFrame encodes the message as a CAN frame.
func (m *BMS_DALY_MOSStatus) MarshalFrame() (can.Frame, error) {
	return m.Frame(), nil
}

// UnmarshalFrame decodes the message from a CAN frame.
func (m *BMS_DALY_MOSStatus) UnmarshalFrame(f can.Frame) error {
	md := Messages().BMS_DALY_MOSStatus
	switch {
	case f.ID != md.ID:
		return fmt.Errorf(
			"unmarshal BMS_DALY_MOSStatus: expects ID 412303361 (got %s with ID %d)", f.String(), f.ID,
		)
	case f.Length != md.Length:
		return fmt.Errorf(
			"unmarshal BMS_DALY_MOSStatus: expects length 8 (got %s with length %d)", f.String(), f.Length,
		)
	case f.IsRemote:
		return fmt.Errorf(
			"unmarshal BMS_DALY_MOSStatus: expects non-remote frame (got remote frame %s)", f.String(),
		)
	case f.IsExtended != md.IsExtended:
		return fmt.Errorf(
			"unmarshal BMS_DALY_MOSStatus: expects extended ID (got %s with standard ID)", f.String(),
		)
	}
	m.xxx_DALY_DischargeStatus = BMS_DALY_MOSStatus_DALY_DischargeStatus(md.DALY_DischargeStatus.UnmarshalUnsigned(f.Data))
	m.xxx_DALY_ChargingMOSTube = uint8(md.DALY_ChargingMOSTube.UnmarshalUnsigned(f.Data))
	m.xxx_DALY_DischargeMOSTube = uint8(md.DALY_DischargeMOSTube.UnmarshalUnsigned(f.Data))
	m.xxx_DALY_BMSLife = uint8(md.DALY_BMSLife.UnmarshalUnsigned(f.Data))
	m.xxx_DALY_ResidualCapacity = uint32(md.DALY_ResidualCapacity.UnmarshalUnsigned(f.Data))
	return nil
}

// BMS_DALY_StatusInformationReader provides read access to a BMS_DALY_StatusInformation message.
type BMS_DALY_StatusInformationReader interface {
	can.FrameMarshaler
	// DALY_BatteryString returns the value of the DALY_BatteryString signal.
	DALY_BatteryString() uint8
	// DALY_TemperatureStatus returns the value of the DALY_TemperatureStatus signal.
	DALY_TemperatureStatus() uint8
	// DALY_ChagerStatus returns the value of the DALY_ChagerStatus signal.
	DALY_ChagerStatus() BMS_DALY_StatusInformation_DALY_ChagerStatus
	// DALY_LoadStatus returns the value of the DALY_LoadStatus signal.
	DALY_LoadStatus() BMS_DALY_StatusInformation_DALY_LoadStatus
	// DALY_DigitalInput1State returns the value of the DALY_DigitalInput1State signal.
	DALY_DigitalInput1State() bool
	// DALY_DigitalInput2State returns the value of the DALY_DigitalInput2State signal.
	DALY_DigitalInput2State() bool
	// DALY_DigitalInput3State returns the value of the DALY_DigitalInput3State signal.
	DALY_DigitalInput3State() bool
	// DALY_DigitalInput4State returns the value of the DALY_DigitalInput4State signal.
	DALY_DigitalInput4State() bool
	// DALY_DigitalOutput1State returns the value of the DALY_DigitalOutput1State signal.
	DALY_DigitalOutput1State() bool
	// DALY_DigitalOutput2State returns the value of the DALY_DigitalOutput2State signal.
	DALY_DigitalOutput2State() bool
	// DALY_DigitalOutput3State returns the value of the DALY_DigitalOutput3State signal.
	DALY_DigitalOutput3State() bool
	// DALY_DigitalOutput4State returns the value of the DALY_DigitalOutput4State signal.
	DALY_DigitalOutput4State() bool
	// DALY_DischargeCycles returns the value of the DALY_DischargeCycles signal.
	DALY_DischargeCycles() uint16
}

// BMS_DALY_StatusInformationWriter provides write access to a BMS_DALY_StatusInformation message.
type BMS_DALY_StatusInformationWriter interface {
	// CopyFrom copies all values from BMS_DALY_StatusInformation.
	CopyFrom(BMS_DALY_StatusInformationReader) *BMS_DALY_StatusInformation
	// SetDALY_BatteryString sets the value of the DALY_BatteryString signal.
	SetDALY_BatteryString(uint8) *BMS_DALY_StatusInformation
	// SetDALY_TemperatureStatus sets the value of the DALY_TemperatureStatus signal.
	SetDALY_TemperatureStatus(uint8) *BMS_DALY_StatusInformation
	// SetDALY_ChagerStatus sets the value of the DALY_ChagerStatus signal.
	SetDALY_ChagerStatus(BMS_DALY_StatusInformation_DALY_ChagerStatus) *BMS_DALY_StatusInformation
	// SetDALY_LoadStatus sets the value of the DALY_LoadStatus signal.
	SetDALY_LoadStatus(BMS_DALY_StatusInformation_DALY_LoadStatus) *BMS_DALY_StatusInformation
	// SetDALY_DigitalInput1State sets the value of the DALY_DigitalInput1State signal.
	SetDALY_DigitalInput1State(bool) *BMS_DALY_StatusInformation
	// SetDALY_DigitalInput2State sets the value of the DALY_DigitalInput2State signal.
	SetDALY_DigitalInput2State(bool) *BMS_DALY_StatusInformation
	// SetDALY_DigitalInput3State sets the value of the DALY_DigitalInput3State signal.
	SetDALY_DigitalInput3State(bool) *BMS_DALY_StatusInformation
	// SetDALY_DigitalInput4State sets the value of the DALY_DigitalInput4State signal.
	SetDALY_DigitalInput4State(bool) *BMS_DALY_StatusInformation
	// SetDALY_DigitalOutput1State sets the value of the DALY_DigitalOutput1State signal.
	SetDALY_DigitalOutput1State(bool) *BMS_DALY_StatusInformation
	// SetDALY_DigitalOutput2State sets the value of the DALY_DigitalOutput2State signal.
	SetDALY_DigitalOutput2State(bool) *BMS_DALY_StatusInformation
	// SetDALY_DigitalOutput3State sets the value of the DALY_DigitalOutput3State signal.
	SetDALY_DigitalOutput3State(bool) *BMS_DALY_StatusInformation
	// SetDALY_DigitalOutput4State sets the value of the DALY_DigitalOutput4State signal.
	SetDALY_DigitalOutput4State(bool) *BMS_DALY_StatusInformation
	// SetDALY_DischargeCycles sets the value of the DALY_DischargeCycles signal.
	SetDALY_DischargeCycles(uint16) *BMS_DALY_StatusInformation
}

type BMS_DALY_StatusInformation struct {
	xxx_DALY_BatteryString       uint8
	xxx_DALY_TemperatureStatus   uint8
	xxx_DALY_ChagerStatus        BMS_DALY_StatusInformation_DALY_ChagerStatus
	xxx_DALY_LoadStatus          BMS_DALY_StatusInformation_DALY_LoadStatus
	xxx_DALY_DigitalInput1State  bool
	xxx_DALY_DigitalInput2State  bool
	xxx_DALY_DigitalInput3State  bool
	xxx_DALY_DigitalInput4State  bool
	xxx_DALY_DigitalOutput1State bool
	xxx_DALY_DigitalOutput2State bool
	xxx_DALY_DigitalOutput3State bool
	xxx_DALY_DigitalOutput4State bool
	xxx_DALY_DischargeCycles     uint16
}

func NewBMS_DALY_StatusInformation() *BMS_DALY_StatusInformation {
	m := &BMS_DALY_StatusInformation{}
	m.Reset()
	return m
}

func (m *BMS_DALY_StatusInformation) Reset() {
	m.xxx_DALY_BatteryString = 0
	m.xxx_DALY_TemperatureStatus = 0
	m.xxx_DALY_ChagerStatus = 0
	m.xxx_DALY_LoadStatus = 0
	m.xxx_DALY_DigitalInput1State = false
	m.xxx_DALY_DigitalInput2State = false
	m.xxx_DALY_DigitalInput3State = false
	m.xxx_DALY_DigitalInput4State = false
	m.xxx_DALY_DigitalOutput1State = false
	m.xxx_DALY_DigitalOutput2State = false
	m.xxx_DALY_DigitalOutput3State = false
	m.xxx_DALY_DigitalOutput4State = false
	m.xxx_DALY_DischargeCycles = 0
}

func (m *BMS_DALY_StatusInformation) CopyFrom(o BMS_DALY_StatusInformationReader) *BMS_DALY_StatusInformation {
	f, _ := o.MarshalFrame()
	_ = m.UnmarshalFrame(f)
	return m
}

// Descriptor returns the BMS_DALY_StatusInformation descriptor.
func (m *BMS_DALY_StatusInformation) Descriptor() *descriptor.Message {
	return Messages().BMS_DALY_StatusInformation.Message
}

// String returns a compact string representation of the message.
func (m *BMS_DALY_StatusInformation) String() string {
	return cantext.MessageString(m)
}

func (m *BMS_DALY_StatusInformation) DALY_BatteryString() uint8 {
	return m.xxx_DALY_BatteryString
}

func (m *BMS_DALY_StatusInformation) SetDALY_BatteryString(v uint8) *BMS_DALY_StatusInformation {
	m.xxx_DALY_BatteryString = uint8(Messages().BMS_DALY_StatusInformation.DALY_BatteryString.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *BMS_DALY_StatusInformation) DALY_TemperatureStatus() uint8 {
	return m.xxx_DALY_TemperatureStatus
}

func (m *BMS_DALY_StatusInformation) SetDALY_TemperatureStatus(v uint8) *BMS_DALY_StatusInformation {
	m.xxx_DALY_TemperatureStatus = uint8(Messages().BMS_DALY_StatusInformation.DALY_TemperatureStatus.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *BMS_DALY_StatusInformation) DALY_ChagerStatus() BMS_DALY_StatusInformation_DALY_ChagerStatus {
	return m.xxx_DALY_ChagerStatus
}

func (m *BMS_DALY_StatusInformation) SetDALY_ChagerStatus(v BMS_DALY_StatusInformation_DALY_ChagerStatus) *BMS_DALY_StatusInformation {
	m.xxx_DALY_ChagerStatus = BMS_DALY_StatusInformation_DALY_ChagerStatus(Messages().BMS_DALY_StatusInformation.DALY_ChagerStatus.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *BMS_DALY_StatusInformation) DALY_LoadStatus() BMS_DALY_StatusInformation_DALY_LoadStatus {
	return m.xxx_DALY_LoadStatus
}

func (m *BMS_DALY_StatusInformation) SetDALY_LoadStatus(v BMS_DALY_StatusInformation_DALY_LoadStatus) *BMS_DALY_StatusInformation {
	m.xxx_DALY_LoadStatus = BMS_DALY_StatusInformation_DALY_LoadStatus(Messages().BMS_DALY_StatusInformation.DALY_LoadStatus.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *BMS_DALY_StatusInformation) DALY_DigitalInput1State() bool {
	return m.xxx_DALY_DigitalInput1State
}

func (m *BMS_DALY_StatusInformation) SetDALY_DigitalInput1State(v bool) *BMS_DALY_StatusInformation {
	m.xxx_DALY_DigitalInput1State = v
	return m
}

func (m *BMS_DALY_StatusInformation) DALY_DigitalInput2State() bool {
	return m.xxx_DALY_DigitalInput2State
}

func (m *BMS_DALY_StatusInformation) SetDALY_DigitalInput2State(v bool) *BMS_DALY_StatusInformation {
	m.xxx_DALY_DigitalInput2State = v
	return m
}

func (m *BMS_DALY_StatusInformation) DALY_DigitalInput3State() bool {
	return m.xxx_DALY_DigitalInput3State
}

func (m *BMS_DALY_StatusInformation) SetDALY_DigitalInput3State(v bool) *BMS_DALY_StatusInformation {
	m.xxx_DALY_DigitalInput3State = v
	return m
}

func (m *BMS_DALY_StatusInformation) DALY_DigitalInput4State() bool {
	return m.xxx_DALY_DigitalInput4State
}

func (m *BMS_DALY_StatusInformation) SetDALY_DigitalInput4State(v bool) *BMS_DALY_StatusInformation {
	m.xxx_DALY_DigitalInput4State = v
	return m
}

func (m *BMS_DALY_StatusInformation) DALY_DigitalOutput1State() bool {
	return m.xxx_DALY_DigitalOutput1State
}

func (m *BMS_DALY_StatusInformation) SetDALY_DigitalOutput1State(v bool) *BMS_DALY_StatusInformation {
	m.xxx_DALY_DigitalOutput1State = v
	return m
}

func (m *BMS_DALY_StatusInformation) DALY_DigitalOutput2State() bool {
	return m.xxx_DALY_DigitalOutput2State
}

func (m *BMS_DALY_StatusInformation) SetDALY_DigitalOutput2State(v bool) *BMS_DALY_StatusInformation {
	m.xxx_DALY_DigitalOutput2State = v
	return m
}

func (m *BMS_DALY_StatusInformation) DALY_DigitalOutput3State() bool {
	return m.xxx_DALY_DigitalOutput3State
}

func (m *BMS_DALY_StatusInformation) SetDALY_DigitalOutput3State(v bool) *BMS_DALY_StatusInformation {
	m.xxx_DALY_DigitalOutput3State = v
	return m
}

func (m *BMS_DALY_StatusInformation) DALY_DigitalOutput4State() bool {
	return m.xxx_DALY_DigitalOutput4State
}

func (m *BMS_DALY_StatusInformation) SetDALY_DigitalOutput4State(v bool) *BMS_DALY_StatusInformation {
	m.xxx_DALY_DigitalOutput4State = v
	return m
}

func (m *BMS_DALY_StatusInformation) DALY_DischargeCycles() uint16 {
	return m.xxx_DALY_DischargeCycles
}

func (m *BMS_DALY_StatusInformation) SetDALY_DischargeCycles(v uint16) *BMS_DALY_StatusInformation {
	m.xxx_DALY_DischargeCycles = uint16(Messages().BMS_DALY_StatusInformation.DALY_DischargeCycles.SaturatedCastUnsigned(uint64(v)))
	return m
}

// BMS_DALY_StatusInformation_DALY_ChagerStatus models the DALY_ChagerStatus signal of the BMS_DALY_StatusInformation message.
type BMS_DALY_StatusInformation_DALY_ChagerStatus uint8

// Value descriptions for the DALY_ChagerStatus signal of the BMS_DALY_StatusInformation message.
const (
	BMS_DALY_StatusInformation_DALY_ChagerStatus_Disconnected BMS_DALY_StatusInformation_DALY_ChagerStatus = 0
	BMS_DALY_StatusInformation_DALY_ChagerStatus_Connected    BMS_DALY_StatusInformation_DALY_ChagerStatus = 1
)

func (v BMS_DALY_StatusInformation_DALY_ChagerStatus) String() string {
	switch v {
	case 0:
		return "Disconnected"
	case 1:
		return "Connected"
	default:
		return fmt.Sprintf("BMS_DALY_StatusInformation_DALY_ChagerStatus(%d)", v)
	}
}

// BMS_DALY_StatusInformation_DALY_LoadStatus models the DALY_LoadStatus signal of the BMS_DALY_StatusInformation message.
type BMS_DALY_StatusInformation_DALY_LoadStatus uint8

// Value descriptions for the DALY_LoadStatus signal of the BMS_DALY_StatusInformation message.
const (
	BMS_DALY_StatusInformation_DALY_LoadStatus_Disconnected BMS_DALY_StatusInformation_DALY_LoadStatus = 0
	BMS_DALY_StatusInformation_DALY_LoadStatus_Access       BMS_DALY_StatusInformation_DALY_LoadStatus = 1
)

func (v BMS_DALY_StatusInformation_DALY_LoadStatus) String() string {
	switch v {
	case 0:
		return "Disconnected"
	case 1:
		return "Access"
	default:
		return fmt.Sprintf("BMS_DALY_StatusInformation_DALY_LoadStatus(%d)", v)
	}
}

// Frame returns a CAN frame representing the message.
func (m *BMS_DALY_StatusInformation) Frame() can.Frame {
	md := Messages().BMS_DALY_StatusInformation
	f := can.Frame{ID: md.ID, IsExtended: md.IsExtended, Length: md.Length}
	md.DALY_BatteryString.MarshalUnsigned(&f.Data, uint64(m.xxx_DALY_BatteryString))
	md.DALY_TemperatureStatus.MarshalUnsigned(&f.Data, uint64(m.xxx_DALY_TemperatureStatus))
	md.DALY_ChagerStatus.MarshalUnsigned(&f.Data, uint64(m.xxx_DALY_ChagerStatus))
	md.DALY_LoadStatus.MarshalUnsigned(&f.Data, uint64(m.xxx_DALY_LoadStatus))
	md.DALY_DigitalInput1State.MarshalBool(&f.Data, bool(m.xxx_DALY_DigitalInput1State))
	md.DALY_DigitalInput2State.MarshalBool(&f.Data, bool(m.xxx_DALY_DigitalInput2State))
	md.DALY_DigitalInput3State.MarshalBool(&f.Data, bool(m.xxx_DALY_DigitalInput3State))
	md.DALY_DigitalInput4State.MarshalBool(&f.Data, bool(m.xxx_DALY_DigitalInput4State))
	md.DALY_DigitalOutput1State.MarshalBool(&f.Data, bool(m.xxx_DALY_DigitalOutput1State))
	md.DALY_DigitalOutput2State.MarshalBool(&f.Data, bool(m.xxx_DALY_DigitalOutput2State))
	md.DALY_DigitalOutput3State.MarshalBool(&f.Data, bool(m.xxx_DALY_DigitalOutput3State))
	md.DALY_DigitalOutput4State.MarshalBool(&f.Data, bool(m.xxx_DALY_DigitalOutput4State))
	md.DALY_DischargeCycles.MarshalUnsigned(&f.Data, uint64(m.xxx_DALY_DischargeCycles))
	return f
}

// MarshalFrame encodes the message as a CAN frame.
func (m *BMS_DALY_StatusInformation) MarshalFrame() (can.Frame, error) {
	return m.Frame(), nil
}

// UnmarshalFrame decodes the message from a CAN frame.
func (m *BMS_DALY_StatusInformation) UnmarshalFrame(f can.Frame) error {
	md := Messages().BMS_DALY_StatusInformation
	switch {
	case f.ID != md.ID:
		return fmt.Errorf(
			"unmarshal BMS_DALY_StatusInformation: expects ID 412368897 (got %s with ID %d)", f.String(), f.ID,
		)
	case f.Length != md.Length:
		return fmt.Errorf(
			"unmarshal BMS_DALY_StatusInformation: expects length 8 (got %s with length %d)", f.String(), f.Length,
		)
	case f.IsRemote:
		return fmt.Errorf(
			"unmarshal BMS_DALY_StatusInformation: expects non-remote frame (got remote frame %s)", f.String(),
		)
	case f.IsExtended != md.IsExtended:
		return fmt.Errorf(
			"unmarshal BMS_DALY_StatusInformation: expects extended ID (got %s with standard ID)", f.String(),
		)
	}
	m.xxx_DALY_BatteryString = uint8(md.DALY_BatteryString.UnmarshalUnsigned(f.Data))
	m.xxx_DALY_TemperatureStatus = uint8(md.DALY_TemperatureStatus.UnmarshalUnsigned(f.Data))
	m.xxx_DALY_ChagerStatus = BMS_DALY_StatusInformation_DALY_ChagerStatus(md.DALY_ChagerStatus.UnmarshalUnsigned(f.Data))
	m.xxx_DALY_LoadStatus = BMS_DALY_StatusInformation_DALY_LoadStatus(md.DALY_LoadStatus.UnmarshalUnsigned(f.Data))
	m.xxx_DALY_DigitalInput1State = bool(md.DALY_DigitalInput1State.UnmarshalBool(f.Data))
	m.xxx_DALY_DigitalInput2State = bool(md.DALY_DigitalInput2State.UnmarshalBool(f.Data))
	m.xxx_DALY_DigitalInput3State = bool(md.DALY_DigitalInput3State.UnmarshalBool(f.Data))
	m.xxx_DALY_DigitalInput4State = bool(md.DALY_DigitalInput4State.UnmarshalBool(f.Data))
	m.xxx_DALY_DigitalOutput1State = bool(md.DALY_DigitalOutput1State.UnmarshalBool(f.Data))
	m.xxx_DALY_DigitalOutput2State = bool(md.DALY_DigitalOutput2State.UnmarshalBool(f.Data))
	m.xxx_DALY_DigitalOutput3State = bool(md.DALY_DigitalOutput3State.UnmarshalBool(f.Data))
	m.xxx_DALY_DigitalOutput4State = bool(md.DALY_DigitalOutput4State.UnmarshalBool(f.Data))
	m.xxx_DALY_DischargeCycles = uint16(md.DALY_DischargeCycles.UnmarshalUnsigned(f.Data))
	return nil
}

// BMS_DALY_MonoCellVoltageReader provides read access to a BMS_DALY_MonoCellVoltage message.
type BMS_DALY_MonoCellVoltageReader interface {
	can.FrameMarshaler
	// DALY_CellFrameNumber returns the value of the DALY_CellFrameNumber signal.
	DALY_CellFrameNumber() uint8
	// DALY_MonomerVoltage1 returns the value of the DALY_MonomerVoltage1 signal.
	DALY_MonomerVoltage1() uint16
	// DALY_MonomerVoltage2 returns the value of the DALY_MonomerVoltage2 signal.
	DALY_MonomerVoltage2() uint16
	// DALY_MonomerVoltage3 returns the value of the DALY_MonomerVoltage3 signal.
	DALY_MonomerVoltage3() uint16
}

// BMS_DALY_MonoCellVoltageWriter provides write access to a BMS_DALY_MonoCellVoltage message.
type BMS_DALY_MonoCellVoltageWriter interface {
	// CopyFrom copies all values from BMS_DALY_MonoCellVoltage.
	CopyFrom(BMS_DALY_MonoCellVoltageReader) *BMS_DALY_MonoCellVoltage
	// SetDALY_CellFrameNumber sets the value of the DALY_CellFrameNumber signal.
	SetDALY_CellFrameNumber(uint8) *BMS_DALY_MonoCellVoltage
	// SetDALY_MonomerVoltage1 sets the value of the DALY_MonomerVoltage1 signal.
	SetDALY_MonomerVoltage1(uint16) *BMS_DALY_MonoCellVoltage
	// SetDALY_MonomerVoltage2 sets the value of the DALY_MonomerVoltage2 signal.
	SetDALY_MonomerVoltage2(uint16) *BMS_DALY_MonoCellVoltage
	// SetDALY_MonomerVoltage3 sets the value of the DALY_MonomerVoltage3 signal.
	SetDALY_MonomerVoltage3(uint16) *BMS_DALY_MonoCellVoltage
}

type BMS_DALY_MonoCellVoltage struct {
	xxx_DALY_CellFrameNumber uint8
	xxx_DALY_MonomerVoltage1 uint16
	xxx_DALY_MonomerVoltage2 uint16
	xxx_DALY_MonomerVoltage3 uint16
}

func NewBMS_DALY_MonoCellVoltage() *BMS_DALY_MonoCellVoltage {
	m := &BMS_DALY_MonoCellVoltage{}
	m.Reset()
	return m
}

func (m *BMS_DALY_MonoCellVoltage) Reset() {
	m.xxx_DALY_CellFrameNumber = 0
	m.xxx_DALY_MonomerVoltage1 = 0
	m.xxx_DALY_MonomerVoltage2 = 0
	m.xxx_DALY_MonomerVoltage3 = 0
}

func (m *BMS_DALY_MonoCellVoltage) CopyFrom(o BMS_DALY_MonoCellVoltageReader) *BMS_DALY_MonoCellVoltage {
	f, _ := o.MarshalFrame()
	_ = m.UnmarshalFrame(f)
	return m
}

// Descriptor returns the BMS_DALY_MonoCellVoltage descriptor.
func (m *BMS_DALY_MonoCellVoltage) Descriptor() *descriptor.Message {
	return Messages().BMS_DALY_MonoCellVoltage.Message
}

// String returns a compact string representation of the message.
func (m *BMS_DALY_MonoCellVoltage) String() string {
	return cantext.MessageString(m)
}

func (m *BMS_DALY_MonoCellVoltage) DALY_CellFrameNumber() uint8 {
	return m.xxx_DALY_CellFrameNumber
}

func (m *BMS_DALY_MonoCellVoltage) SetDALY_CellFrameNumber(v uint8) *BMS_DALY_MonoCellVoltage {
	m.xxx_DALY_CellFrameNumber = uint8(Messages().BMS_DALY_MonoCellVoltage.DALY_CellFrameNumber.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *BMS_DALY_MonoCellVoltage) DALY_MonomerVoltage1() uint16 {
	return m.xxx_DALY_MonomerVoltage1
}

func (m *BMS_DALY_MonoCellVoltage) SetDALY_MonomerVoltage1(v uint16) *BMS_DALY_MonoCellVoltage {
	m.xxx_DALY_MonomerVoltage1 = uint16(Messages().BMS_DALY_MonoCellVoltage.DALY_MonomerVoltage1.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *BMS_DALY_MonoCellVoltage) DALY_MonomerVoltage2() uint16 {
	return m.xxx_DALY_MonomerVoltage2
}

func (m *BMS_DALY_MonoCellVoltage) SetDALY_MonomerVoltage2(v uint16) *BMS_DALY_MonoCellVoltage {
	m.xxx_DALY_MonomerVoltage2 = uint16(Messages().BMS_DALY_MonoCellVoltage.DALY_MonomerVoltage2.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *BMS_DALY_MonoCellVoltage) DALY_MonomerVoltage3() uint16 {
	return m.xxx_DALY_MonomerVoltage3
}

func (m *BMS_DALY_MonoCellVoltage) SetDALY_MonomerVoltage3(v uint16) *BMS_DALY_MonoCellVoltage {
	m.xxx_DALY_MonomerVoltage3 = uint16(Messages().BMS_DALY_MonoCellVoltage.DALY_MonomerVoltage3.SaturatedCastUnsigned(uint64(v)))
	return m
}

// Frame returns a CAN frame representing the message.
func (m *BMS_DALY_MonoCellVoltage) Frame() can.Frame {
	md := Messages().BMS_DALY_MonoCellVoltage
	f := can.Frame{ID: md.ID, IsExtended: md.IsExtended, Length: md.Length}
	md.DALY_CellFrameNumber.MarshalUnsigned(&f.Data, uint64(m.xxx_DALY_CellFrameNumber))
	md.DALY_MonomerVoltage1.MarshalUnsigned(&f.Data, uint64(m.xxx_DALY_MonomerVoltage1))
	md.DALY_MonomerVoltage2.MarshalUnsigned(&f.Data, uint64(m.xxx_DALY_MonomerVoltage2))
	md.DALY_MonomerVoltage3.MarshalUnsigned(&f.Data, uint64(m.xxx_DALY_MonomerVoltage3))
	return f
}

// MarshalFrame encodes the message as a CAN frame.
func (m *BMS_DALY_MonoCellVoltage) MarshalFrame() (can.Frame, error) {
	return m.Frame(), nil
}

// UnmarshalFrame decodes the message from a CAN frame.
func (m *BMS_DALY_MonoCellVoltage) UnmarshalFrame(f can.Frame) error {
	md := Messages().BMS_DALY_MonoCellVoltage
	switch {
	case f.ID != md.ID:
		return fmt.Errorf(
			"unmarshal BMS_DALY_MonoCellVoltage: expects ID 412434433 (got %s with ID %d)", f.String(), f.ID,
		)
	case f.Length != md.Length:
		return fmt.Errorf(
			"unmarshal BMS_DALY_MonoCellVoltage: expects length 8 (got %s with length %d)", f.String(), f.Length,
		)
	case f.IsRemote:
		return fmt.Errorf(
			"unmarshal BMS_DALY_MonoCellVoltage: expects non-remote frame (got remote frame %s)", f.String(),
		)
	case f.IsExtended != md.IsExtended:
		return fmt.Errorf(
			"unmarshal BMS_DALY_MonoCellVoltage: expects extended ID (got %s with standard ID)", f.String(),
		)
	}
	m.xxx_DALY_CellFrameNumber = uint8(md.DALY_CellFrameNumber.UnmarshalUnsigned(f.Data))
	m.xxx_DALY_MonomerVoltage1 = uint16(md.DALY_MonomerVoltage1.UnmarshalUnsigned(f.Data))
	m.xxx_DALY_MonomerVoltage2 = uint16(md.DALY_MonomerVoltage2.UnmarshalUnsigned(f.Data))
	m.xxx_DALY_MonomerVoltage3 = uint16(md.DALY_MonomerVoltage3.UnmarshalUnsigned(f.Data))
	return nil
}

// BMS_DALY_MonoCellTempReader provides read access to a BMS_DALY_MonoCellTemp message.
type BMS_DALY_MonoCellTempReader interface {
	can.FrameMarshaler
	// DALY_TempFrameNumber returns the value of the DALY_TempFrameNumber signal.
	DALY_TempFrameNumber() uint8
	// DALY_MonomerTemp1 returns the physical value of the DALY_MonomerTemp1 signal.
	DALY_MonomerTemp1() float64
	// RawDALY_MonomerTemp1 returns the raw (encoded) value of the DALY_MonomerTemp1 signal.
	RawDALY_MonomerTemp1() uint8
	// DALY_MonomerTemp2 returns the physical value of the DALY_MonomerTemp2 signal.
	DALY_MonomerTemp2() float64
	// RawDALY_MonomerTemp2 returns the raw (encoded) value of the DALY_MonomerTemp2 signal.
	RawDALY_MonomerTemp2() uint8
	// DALY_MonomerTemp3 returns the physical value of the DALY_MonomerTemp3 signal.
	DALY_MonomerTemp3() float64
	// RawDALY_MonomerTemp3 returns the raw (encoded) value of the DALY_MonomerTemp3 signal.
	RawDALY_MonomerTemp3() uint8
	// DALY_MonomerTemp4 returns the physical value of the DALY_MonomerTemp4 signal.
	DALY_MonomerTemp4() float64
	// RawDALY_MonomerTemp4 returns the raw (encoded) value of the DALY_MonomerTemp4 signal.
	RawDALY_MonomerTemp4() uint8
	// DALY_MonomerTemp5 returns the physical value of the DALY_MonomerTemp5 signal.
	DALY_MonomerTemp5() float64
	// RawDALY_MonomerTemp5 returns the raw (encoded) value of the DALY_MonomerTemp5 signal.
	RawDALY_MonomerTemp5() uint8
	// DALY_MonomerTemp6 returns the physical value of the DALY_MonomerTemp6 signal.
	DALY_MonomerTemp6() float64
	// RawDALY_MonomerTemp6 returns the raw (encoded) value of the DALY_MonomerTemp6 signal.
	RawDALY_MonomerTemp6() uint8
	// DALY_MonomerTemp7 returns the physical value of the DALY_MonomerTemp7 signal.
	DALY_MonomerTemp7() float64
	// RawDALY_MonomerTemp7 returns the raw (encoded) value of the DALY_MonomerTemp7 signal.
	RawDALY_MonomerTemp7() uint8
}

// BMS_DALY_MonoCellTempWriter provides write access to a BMS_DALY_MonoCellTemp message.
type BMS_DALY_MonoCellTempWriter interface {
	// CopyFrom copies all values from BMS_DALY_MonoCellTemp.
	CopyFrom(BMS_DALY_MonoCellTempReader) *BMS_DALY_MonoCellTemp
	// SetDALY_TempFrameNumber sets the value of the DALY_TempFrameNumber signal.
	SetDALY_TempFrameNumber(uint8) *BMS_DALY_MonoCellTemp
	// SetDALY_MonomerTemp1 sets the physical value of the DALY_MonomerTemp1 signal.
	SetDALY_MonomerTemp1(float64) *BMS_DALY_MonoCellTemp
	// SetRawDALY_MonomerTemp1 sets the raw (encoded) value of the DALY_MonomerTemp1 signal.
	SetRawDALY_MonomerTemp1(uint8) *BMS_DALY_MonoCellTemp
	// SetDALY_MonomerTemp2 sets the physical value of the DALY_MonomerTemp2 signal.
	SetDALY_MonomerTemp2(float64) *BMS_DALY_MonoCellTemp
	// SetRawDALY_MonomerTemp2 sets the raw (encoded) value of the DALY_MonomerTemp2 signal.
	SetRawDALY_MonomerTemp2(uint8) *BMS_DALY_MonoCellTemp
	// SetDALY_MonomerTemp3 sets the physical value of the DALY_MonomerTemp3 signal.
	SetDALY_MonomerTemp3(float64) *BMS_DALY_MonoCellTemp
	// SetRawDALY_MonomerTemp3 sets the raw (encoded) value of the DALY_MonomerTemp3 signal.
	SetRawDALY_MonomerTemp3(uint8) *BMS_DALY_MonoCellTemp
	// SetDALY_MonomerTemp4 sets the physical value of the DALY_MonomerTemp4 signal.
	SetDALY_MonomerTemp4(float64) *BMS_DALY_MonoCellTemp
	// SetRawDALY_MonomerTemp4 sets the raw (encoded) value of the DALY_MonomerTemp4 signal.
	SetRawDALY_MonomerTemp4(uint8) *BMS_DALY_MonoCellTemp
	// SetDALY_MonomerTemp5 sets the physical value of the DALY_MonomerTemp5 signal.
	SetDALY_MonomerTemp5(float64) *BMS_DALY_MonoCellTemp
	// SetRawDALY_MonomerTemp5 sets the raw (encoded) value of the DALY_MonomerTemp5 signal.
	SetRawDALY_MonomerTemp5(uint8) *BMS_DALY_MonoCellTemp
	// SetDALY_MonomerTemp6 sets the physical value of the DALY_MonomerTemp6 signal.
	SetDALY_MonomerTemp6(float64) *BMS_DALY_MonoCellTemp
	// SetRawDALY_MonomerTemp6 sets the raw (encoded) value of the DALY_MonomerTemp6 signal.
	SetRawDALY_MonomerTemp6(uint8) *BMS_DALY_MonoCellTemp
	// SetDALY_MonomerTemp7 sets the physical value of the DALY_MonomerTemp7 signal.
	SetDALY_MonomerTemp7(float64) *BMS_DALY_MonoCellTemp
	// SetRawDALY_MonomerTemp7 sets the raw (encoded) value of the DALY_MonomerTemp7 signal.
	SetRawDALY_MonomerTemp7(uint8) *BMS_DALY_MonoCellTemp
}

type BMS_DALY_MonoCellTemp struct {
	xxx_DALY_TempFrameNumber uint8
	xxx_DALY_MonomerTemp1    uint8
	xxx_DALY_MonomerTemp2    uint8
	xxx_DALY_MonomerTemp3    uint8
	xxx_DALY_MonomerTemp4    uint8
	xxx_DALY_MonomerTemp5    uint8
	xxx_DALY_MonomerTemp6    uint8
	xxx_DALY_MonomerTemp7    uint8
}

func NewBMS_DALY_MonoCellTemp() *BMS_DALY_MonoCellTemp {
	m := &BMS_DALY_MonoCellTemp{}
	m.Reset()
	return m
}

func (m *BMS_DALY_MonoCellTemp) Reset() {
	m.xxx_DALY_TempFrameNumber = 0
	m.xxx_DALY_MonomerTemp1 = 0
	m.xxx_DALY_MonomerTemp2 = 0
	m.xxx_DALY_MonomerTemp3 = 0
	m.xxx_DALY_MonomerTemp4 = 0
	m.xxx_DALY_MonomerTemp5 = 0
	m.xxx_DALY_MonomerTemp6 = 0
	m.xxx_DALY_MonomerTemp7 = 0
}

func (m *BMS_DALY_MonoCellTemp) CopyFrom(o BMS_DALY_MonoCellTempReader) *BMS_DALY_MonoCellTemp {
	f, _ := o.MarshalFrame()
	_ = m.UnmarshalFrame(f)
	return m
}

// Descriptor returns the BMS_DALY_MonoCellTemp descriptor.
func (m *BMS_DALY_MonoCellTemp) Descriptor() *descriptor.Message {
	return Messages().BMS_DALY_MonoCellTemp.Message
}

// String returns a compact string representation of the message.
func (m *BMS_DALY_MonoCellTemp) String() string {
	return cantext.MessageString(m)
}

func (m *BMS_DALY_MonoCellTemp) DALY_TempFrameNumber() uint8 {
	return m.xxx_DALY_TempFrameNumber
}

func (m *BMS_DALY_MonoCellTemp) SetDALY_TempFrameNumber(v uint8) *BMS_DALY_MonoCellTemp {
	m.xxx_DALY_TempFrameNumber = uint8(Messages().BMS_DALY_MonoCellTemp.DALY_TempFrameNumber.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *BMS_DALY_MonoCellTemp) DALY_MonomerTemp1() float64 {
	return Messages().BMS_DALY_MonoCellTemp.DALY_MonomerTemp1.ToPhysical(float64(m.xxx_DALY_MonomerTemp1))
}

func (m *BMS_DALY_MonoCellTemp) SetDALY_MonomerTemp1(v float64) *BMS_DALY_MonoCellTemp {
	m.xxx_DALY_MonomerTemp1 = uint8(Messages().BMS_DALY_MonoCellTemp.DALY_MonomerTemp1.FromPhysical(v))
	return m
}

func (m *BMS_DALY_MonoCellTemp) RawDALY_MonomerTemp1() uint8 {
	return m.xxx_DALY_MonomerTemp1
}

func (m *BMS_DALY_MonoCellTemp) SetRawDALY_MonomerTemp1(v uint8) *BMS_DALY_MonoCellTemp {
	m.xxx_DALY_MonomerTemp1 = uint8(Messages().BMS_DALY_MonoCellTemp.DALY_MonomerTemp1.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *BMS_DALY_MonoCellTemp) DALY_MonomerTemp2() float64 {
	return Messages().BMS_DALY_MonoCellTemp.DALY_MonomerTemp2.ToPhysical(float64(m.xxx_DALY_MonomerTemp2))
}

func (m *BMS_DALY_MonoCellTemp) SetDALY_MonomerTemp2(v float64) *BMS_DALY_MonoCellTemp {
	m.xxx_DALY_MonomerTemp2 = uint8(Messages().BMS_DALY_MonoCellTemp.DALY_MonomerTemp2.FromPhysical(v))
	return m
}

func (m *BMS_DALY_MonoCellTemp) RawDALY_MonomerTemp2() uint8 {
	return m.xxx_DALY_MonomerTemp2
}

func (m *BMS_DALY_MonoCellTemp) SetRawDALY_MonomerTemp2(v uint8) *BMS_DALY_MonoCellTemp {
	m.xxx_DALY_MonomerTemp2 = uint8(Messages().BMS_DALY_MonoCellTemp.DALY_MonomerTemp2.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *BMS_DALY_MonoCellTemp) DALY_MonomerTemp3() float64 {
	return Messages().BMS_DALY_MonoCellTemp.DALY_MonomerTemp3.ToPhysical(float64(m.xxx_DALY_MonomerTemp3))
}

func (m *BMS_DALY_MonoCellTemp) SetDALY_MonomerTemp3(v float64) *BMS_DALY_MonoCellTemp {
	m.xxx_DALY_MonomerTemp3 = uint8(Messages().BMS_DALY_MonoCellTemp.DALY_MonomerTemp3.FromPhysical(v))
	return m
}

func (m *BMS_DALY_MonoCellTemp) RawDALY_MonomerTemp3() uint8 {
	return m.xxx_DALY_MonomerTemp3
}

func (m *BMS_DALY_MonoCellTemp) SetRawDALY_MonomerTemp3(v uint8) *BMS_DALY_MonoCellTemp {
	m.xxx_DALY_MonomerTemp3 = uint8(Messages().BMS_DALY_MonoCellTemp.DALY_MonomerTemp3.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *BMS_DALY_MonoCellTemp) DALY_MonomerTemp4() float64 {
	return Messages().BMS_DALY_MonoCellTemp.DALY_MonomerTemp4.ToPhysical(float64(m.xxx_DALY_MonomerTemp4))
}

func (m *BMS_DALY_MonoCellTemp) SetDALY_MonomerTemp4(v float64) *BMS_DALY_MonoCellTemp {
	m.xxx_DALY_MonomerTemp4 = uint8(Messages().BMS_DALY_MonoCellTemp.DALY_MonomerTemp4.FromPhysical(v))
	return m
}

func (m *BMS_DALY_MonoCellTemp) RawDALY_MonomerTemp4() uint8 {
	return m.xxx_DALY_MonomerTemp4
}

func (m *BMS_DALY_MonoCellTemp) SetRawDALY_MonomerTemp4(v uint8) *BMS_DALY_MonoCellTemp {
	m.xxx_DALY_MonomerTemp4 = uint8(Messages().BMS_DALY_MonoCellTemp.DALY_MonomerTemp4.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *BMS_DALY_MonoCellTemp) DALY_MonomerTemp5() float64 {
	return Messages().BMS_DALY_MonoCellTemp.DALY_MonomerTemp5.ToPhysical(float64(m.xxx_DALY_MonomerTemp5))
}

func (m *BMS_DALY_MonoCellTemp) SetDALY_MonomerTemp5(v float64) *BMS_DALY_MonoCellTemp {
	m.xxx_DALY_MonomerTemp5 = uint8(Messages().BMS_DALY_MonoCellTemp.DALY_MonomerTemp5.FromPhysical(v))
	return m
}

func (m *BMS_DALY_MonoCellTemp) RawDALY_MonomerTemp5() uint8 {
	return m.xxx_DALY_MonomerTemp5
}

func (m *BMS_DALY_MonoCellTemp) SetRawDALY_MonomerTemp5(v uint8) *BMS_DALY_MonoCellTemp {
	m.xxx_DALY_MonomerTemp5 = uint8(Messages().BMS_DALY_MonoCellTemp.DALY_MonomerTemp5.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *BMS_DALY_MonoCellTemp) DALY_MonomerTemp6() float64 {
	return Messages().BMS_DALY_MonoCellTemp.DALY_MonomerTemp6.ToPhysical(float64(m.xxx_DALY_MonomerTemp6))
}

func (m *BMS_DALY_MonoCellTemp) SetDALY_MonomerTemp6(v float64) *BMS_DALY_MonoCellTemp {
	m.xxx_DALY_MonomerTemp6 = uint8(Messages().BMS_DALY_MonoCellTemp.DALY_MonomerTemp6.FromPhysical(v))
	return m
}

func (m *BMS_DALY_MonoCellTemp) RawDALY_MonomerTemp6() uint8 {
	return m.xxx_DALY_MonomerTemp6
}

func (m *BMS_DALY_MonoCellTemp) SetRawDALY_MonomerTemp6(v uint8) *BMS_DALY_MonoCellTemp {
	m.xxx_DALY_MonomerTemp6 = uint8(Messages().BMS_DALY_MonoCellTemp.DALY_MonomerTemp6.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *BMS_DALY_MonoCellTemp) DALY_MonomerTemp7() float64 {
	return Messages().BMS_DALY_MonoCellTemp.DALY_MonomerTemp7.ToPhysical(float64(m.xxx_DALY_MonomerTemp7))
}

func (m *BMS_DALY_MonoCellTemp) SetDALY_MonomerTemp7(v float64) *BMS_DALY_MonoCellTemp {
	m.xxx_DALY_MonomerTemp7 = uint8(Messages().BMS_DALY_MonoCellTemp.DALY_MonomerTemp7.FromPhysical(v))
	return m
}

func (m *BMS_DALY_MonoCellTemp) RawDALY_MonomerTemp7() uint8 {
	return m.xxx_DALY_MonomerTemp7
}

func (m *BMS_DALY_MonoCellTemp) SetRawDALY_MonomerTemp7(v uint8) *BMS_DALY_MonoCellTemp {
	m.xxx_DALY_MonomerTemp7 = uint8(Messages().BMS_DALY_MonoCellTemp.DALY_MonomerTemp7.SaturatedCastUnsigned(uint64(v)))
	return m
}

// Frame returns a CAN frame representing the message.
func (m *BMS_DALY_MonoCellTemp) Frame() can.Frame {
	md := Messages().BMS_DALY_MonoCellTemp
	f := can.Frame{ID: md.ID, IsExtended: md.IsExtended, Length: md.Length}
	md.DALY_TempFrameNumber.MarshalUnsigned(&f.Data, uint64(m.xxx_DALY_TempFrameNumber))
	md.DALY_MonomerTemp1.MarshalUnsigned(&f.Data, uint64(m.xxx_DALY_MonomerTemp1))
	md.DALY_MonomerTemp2.MarshalUnsigned(&f.Data, uint64(m.xxx_DALY_MonomerTemp2))
	md.DALY_MonomerTemp3.MarshalUnsigned(&f.Data, uint64(m.xxx_DALY_MonomerTemp3))
	md.DALY_MonomerTemp4.MarshalUnsigned(&f.Data, uint64(m.xxx_DALY_MonomerTemp4))
	md.DALY_MonomerTemp5.MarshalUnsigned(&f.Data, uint64(m.xxx_DALY_MonomerTemp5))
	md.DALY_MonomerTemp6.MarshalUnsigned(&f.Data, uint64(m.xxx_DALY_MonomerTemp6))
	md.DALY_MonomerTemp7.MarshalUnsigned(&f.Data, uint64(m.xxx_DALY_MonomerTemp7))
	return f
}

// MarshalFrame encodes the message as a CAN frame.
func (m *BMS_DALY_MonoCellTemp) MarshalFrame() (can.Frame, error) {
	return m.Frame(), nil
}

// UnmarshalFrame decodes the message from a CAN frame.
func (m *BMS_DALY_MonoCellTemp) UnmarshalFrame(f can.Frame) error {
	md := Messages().BMS_DALY_MonoCellTemp
	switch {
	case f.ID != md.ID:
		return fmt.Errorf(
			"unmarshal BMS_DALY_MonoCellTemp: expects ID 412499969 (got %s with ID %d)", f.String(), f.ID,
		)
	case f.Length != md.Length:
		return fmt.Errorf(
			"unmarshal BMS_DALY_MonoCellTemp: expects length 8 (got %s with length %d)", f.String(), f.Length,
		)
	case f.IsRemote:
		return fmt.Errorf(
			"unmarshal BMS_DALY_MonoCellTemp: expects non-remote frame (got remote frame %s)", f.String(),
		)
	case f.IsExtended != md.IsExtended:
		return fmt.Errorf(
			"unmarshal BMS_DALY_MonoCellTemp: expects extended ID (got %s with standard ID)", f.String(),
		)
	}
	m.xxx_DALY_TempFrameNumber = uint8(md.DALY_TempFrameNumber.UnmarshalUnsigned(f.Data))
	m.xxx_DALY_MonomerTemp1 = uint8(md.DALY_MonomerTemp1.UnmarshalUnsigned(f.Data))
	m.xxx_DALY_MonomerTemp2 = uint8(md.DALY_MonomerTemp2.UnmarshalUnsigned(f.Data))
	m.xxx_DALY_MonomerTemp3 = uint8(md.DALY_MonomerTemp3.UnmarshalUnsigned(f.Data))
	m.xxx_DALY_MonomerTemp4 = uint8(md.DALY_MonomerTemp4.UnmarshalUnsigned(f.Data))
	m.xxx_DALY_MonomerTemp5 = uint8(md.DALY_MonomerTemp5.UnmarshalUnsigned(f.Data))
	m.xxx_DALY_MonomerTemp6 = uint8(md.DALY_MonomerTemp6.UnmarshalUnsigned(f.Data))
	m.xxx_DALY_MonomerTemp7 = uint8(md.DALY_MonomerTemp7.UnmarshalUnsigned(f.Data))
	return nil
}

// OBC_ResponseReader provides read access to a OBC_Response message.
type OBC_ResponseReader interface {
	can.FrameMarshaler
	// OBC_OutputVoltage returns the physical value of the OBC_OutputVoltage signal.
	OBC_OutputVoltage() float64
	// RawOBC_OutputVoltage returns the raw (encoded) value of the OBC_OutputVoltage signal.
	RawOBC_OutputVoltage() uint16
	// OBC_OutputCurrent returns the physical value of the OBC_OutputCurrent signal.
	OBC_OutputCurrent() float64
	// RawOBC_OutputCurrent returns the raw (encoded) value of the OBC_OutputCurrent signal.
	RawOBC_OutputCurrent() uint16
	// OBC_HardwareProtection returns the value of the OBC_HardwareProtection signal.
	OBC_HardwareProtection() OBC_Response_OBC_HardwareProtection
	// OBC_TemperatureProtection returns the value of the OBC_TemperatureProtection signal.
	OBC_TemperatureProtection() OBC_Response_OBC_TemperatureProtection
	// OBC_InputVoltageStatus returns the value of the OBC_InputVoltageStatus signal.
	OBC_InputVoltageStatus() OBC_Response_OBC_InputVoltageStatus
	// OBC_OutputUnderVoltage returns the value of the OBC_OutputUnderVoltage signal.
	OBC_OutputUnderVoltage() OBC_Response_OBC_OutputUnderVoltage
	// OBC_OutputOverVoltage returns the value of the OBC_OutputOverVoltage signal.
	OBC_OutputOverVoltage() OBC_Response_OBC_OutputOverVoltage
	// OBC_OutputOvercurrent returns the value of the OBC_OutputOvercurrent signal.
	OBC_OutputOvercurrent() OBC_Response_OBC_OutputOvercurrent
	// OBC_OutputShortCircuit returns the value of the OBC_OutputShortCircuit signal.
	OBC_OutputShortCircuit() OBC_Response_OBC_OutputShortCircuit
	// OBC_CommunicationStatus returns the value of the OBC_CommunicationStatus signal.
	OBC_CommunicationStatus() OBC_Response_OBC_CommunicationStatus
	// OBC_WorkingStatus returns the value of the OBC_WorkingStatus signal.
	OBC_WorkingStatus() OBC_Response_OBC_WorkingStatus
	// OBC_ComOfInitialization returns the value of the OBC_ComOfInitialization signal.
	OBC_ComOfInitialization() OBC_Response_OBC_ComOfInitialization
	// OBC_FanWorkingEnable returns the value of the OBC_FanWorkingEnable signal.
	OBC_FanWorkingEnable() OBC_Response_OBC_FanWorkingEnable
	// OBC_CoolingPumpEnable returns the value of the OBC_CoolingPumpEnable signal.
	OBC_CoolingPumpEnable() OBC_Response_OBC_CoolingPumpEnable
	// OBC_CCSignalStatus returns the value of the OBC_CCSignalStatus signal.
	OBC_CCSignalStatus() OBC_Response_OBC_CCSignalStatus
	// OBC_CPSignalStatus returns the value of the OBC_CPSignalStatus signal.
	OBC_CPSignalStatus() OBC_Response_OBC_CPSignalStatus
	// OBC_SocketOvheatFault returns the value of the OBC_SocketOvheatFault signal.
	OBC_SocketOvheatFault() OBC_Response_OBC_SocketOvheatFault
	// OBC_ElectronicLockState returns the value of the OBC_ElectronicLockState signal.
	OBC_ElectronicLockState() OBC_Response_OBC_ElectronicLockState
	// OBC_S2SwitchControlBitStatus returns the value of the OBC_S2SwitchControlBitStatus signal.
	OBC_S2SwitchControlBitStatus() OBC_Response_OBC_S2SwitchControlBitStatus
	// OBC_Temperature returns the physical value of the OBC_Temperature signal.
	OBC_Temperature() float64
	// RawOBC_Temperature returns the raw (encoded) value of the OBC_Temperature signal.
	RawOBC_Temperature() uint8
}

// OBC_ResponseWriter provides write access to a OBC_Response message.
type OBC_ResponseWriter interface {
	// CopyFrom copies all values from OBC_Response.
	CopyFrom(OBC_ResponseReader) *OBC_Response
	// SetOBC_OutputVoltage sets the physical value of the OBC_OutputVoltage signal.
	SetOBC_OutputVoltage(float64) *OBC_Response
	// SetRawOBC_OutputVoltage sets the raw (encoded) value of the OBC_OutputVoltage signal.
	SetRawOBC_OutputVoltage(uint16) *OBC_Response
	// SetOBC_OutputCurrent sets the physical value of the OBC_OutputCurrent signal.
	SetOBC_OutputCurrent(float64) *OBC_Response
	// SetRawOBC_OutputCurrent sets the raw (encoded) value of the OBC_OutputCurrent signal.
	SetRawOBC_OutputCurrent(uint16) *OBC_Response
	// SetOBC_HardwareProtection sets the value of the OBC_HardwareProtection signal.
	SetOBC_HardwareProtection(OBC_Response_OBC_HardwareProtection) *OBC_Response
	// SetOBC_TemperatureProtection sets the value of the OBC_TemperatureProtection signal.
	SetOBC_TemperatureProtection(OBC_Response_OBC_TemperatureProtection) *OBC_Response
	// SetOBC_InputVoltageStatus sets the value of the OBC_InputVoltageStatus signal.
	SetOBC_InputVoltageStatus(OBC_Response_OBC_InputVoltageStatus) *OBC_Response
	// SetOBC_OutputUnderVoltage sets the value of the OBC_OutputUnderVoltage signal.
	SetOBC_OutputUnderVoltage(OBC_Response_OBC_OutputUnderVoltage) *OBC_Response
	// SetOBC_OutputOverVoltage sets the value of the OBC_OutputOverVoltage signal.
	SetOBC_OutputOverVoltage(OBC_Response_OBC_OutputOverVoltage) *OBC_Response
	// SetOBC_OutputOvercurrent sets the value of the OBC_OutputOvercurrent signal.
	SetOBC_OutputOvercurrent(OBC_Response_OBC_OutputOvercurrent) *OBC_Response
	// SetOBC_OutputShortCircuit sets the value of the OBC_OutputShortCircuit signal.
	SetOBC_OutputShortCircuit(OBC_Response_OBC_OutputShortCircuit) *OBC_Response
	// SetOBC_CommunicationStatus sets the value of the OBC_CommunicationStatus signal.
	SetOBC_CommunicationStatus(OBC_Response_OBC_CommunicationStatus) *OBC_Response
	// SetOBC_WorkingStatus sets the value of the OBC_WorkingStatus signal.
	SetOBC_WorkingStatus(OBC_Response_OBC_WorkingStatus) *OBC_Response
	// SetOBC_ComOfInitialization sets the value of the OBC_ComOfInitialization signal.
	SetOBC_ComOfInitialization(OBC_Response_OBC_ComOfInitialization) *OBC_Response
	// SetOBC_FanWorkingEnable sets the value of the OBC_FanWorkingEnable signal.
	SetOBC_FanWorkingEnable(OBC_Response_OBC_FanWorkingEnable) *OBC_Response
	// SetOBC_CoolingPumpEnable sets the value of the OBC_CoolingPumpEnable signal.
	SetOBC_CoolingPumpEnable(OBC_Response_OBC_CoolingPumpEnable) *OBC_Response
	// SetOBC_CCSignalStatus sets the value of the OBC_CCSignalStatus signal.
	SetOBC_CCSignalStatus(OBC_Response_OBC_CCSignalStatus) *OBC_Response
	// SetOBC_CPSignalStatus sets the value of the OBC_CPSignalStatus signal.
	SetOBC_CPSignalStatus(OBC_Response_OBC_CPSignalStatus) *OBC_Response
	// SetOBC_SocketOvheatFault sets the value of the OBC_SocketOvheatFault signal.
	SetOBC_SocketOvheatFault(OBC_Response_OBC_SocketOvheatFault) *OBC_Response
	// SetOBC_ElectronicLockState sets the value of the OBC_ElectronicLockState signal.
	SetOBC_ElectronicLockState(OBC_Response_OBC_ElectronicLockState) *OBC_Response
	// SetOBC_S2SwitchControlBitStatus sets the value of the OBC_S2SwitchControlBitStatus signal.
	SetOBC_S2SwitchControlBitStatus(OBC_Response_OBC_S2SwitchControlBitStatus) *OBC_Response
	// SetOBC_Temperature sets the physical value of the OBC_Temperature signal.
	SetOBC_Temperature(float64) *OBC_Response
	// SetRawOBC_Temperature sets the raw (encoded) value of the OBC_Temperature signal.
	SetRawOBC_Temperature(uint8) *OBC_Response
}

type OBC_Response struct {
	xxx_OBC_OutputVoltage            uint16
	xxx_OBC_OutputCurrent            uint16
	xxx_OBC_HardwareProtection       OBC_Response_OBC_HardwareProtection
	xxx_OBC_TemperatureProtection    OBC_Response_OBC_TemperatureProtection
	xxx_OBC_InputVoltageStatus       OBC_Response_OBC_InputVoltageStatus
	xxx_OBC_OutputUnderVoltage       OBC_Response_OBC_OutputUnderVoltage
	xxx_OBC_OutputOverVoltage        OBC_Response_OBC_OutputOverVoltage
	xxx_OBC_OutputOvercurrent        OBC_Response_OBC_OutputOvercurrent
	xxx_OBC_OutputShortCircuit       OBC_Response_OBC_OutputShortCircuit
	xxx_OBC_CommunicationStatus      OBC_Response_OBC_CommunicationStatus
	xxx_OBC_WorkingStatus            OBC_Response_OBC_WorkingStatus
	xxx_OBC_ComOfInitialization      OBC_Response_OBC_ComOfInitialization
	xxx_OBC_FanWorkingEnable         OBC_Response_OBC_FanWorkingEnable
	xxx_OBC_CoolingPumpEnable        OBC_Response_OBC_CoolingPumpEnable
	xxx_OBC_CCSignalStatus           OBC_Response_OBC_CCSignalStatus
	xxx_OBC_CPSignalStatus           OBC_Response_OBC_CPSignalStatus
	xxx_OBC_SocketOvheatFault        OBC_Response_OBC_SocketOvheatFault
	xxx_OBC_ElectronicLockState      OBC_Response_OBC_ElectronicLockState
	xxx_OBC_S2SwitchControlBitStatus OBC_Response_OBC_S2SwitchControlBitStatus
	xxx_OBC_Temperature              uint8
}

func NewOBC_Response() *OBC_Response {
	m := &OBC_Response{}
	m.Reset()
	return m
}

func (m *OBC_Response) Reset() {
	m.xxx_OBC_OutputVoltage = 0
	m.xxx_OBC_OutputCurrent = 0
	m.xxx_OBC_HardwareProtection = false
	m.xxx_OBC_TemperatureProtection = false
	m.xxx_OBC_InputVoltageStatus = 0
	m.xxx_OBC_OutputUnderVoltage = false
	m.xxx_OBC_OutputOverVoltage = false
	m.xxx_OBC_OutputOvercurrent = false
	m.xxx_OBC_OutputShortCircuit = false
	m.xxx_OBC_CommunicationStatus = false
	m.xxx_OBC_WorkingStatus = 0
	m.xxx_OBC_ComOfInitialization = false
	m.xxx_OBC_FanWorkingEnable = false
	m.xxx_OBC_CoolingPumpEnable = false
	m.xxx_OBC_CCSignalStatus = 0
	m.xxx_OBC_CPSignalStatus = false
	m.xxx_OBC_SocketOvheatFault = false
	m.xxx_OBC_ElectronicLockState = 0
	m.xxx_OBC_S2SwitchControlBitStatus = false
	m.xxx_OBC_Temperature = 0
}

func (m *OBC_Response) CopyFrom(o OBC_ResponseReader) *OBC_Response {
	f, _ := o.MarshalFrame()
	_ = m.UnmarshalFrame(f)
	return m
}

// Descriptor returns the OBC_Response descriptor.
func (m *OBC_Response) Descriptor() *descriptor.Message {
	return Messages().OBC_Response.Message
}

// String returns a compact string representation of the message.
func (m *OBC_Response) String() string {
	return cantext.MessageString(m)
}

func (m *OBC_Response) OBC_OutputVoltage() float64 {
	return Messages().OBC_Response.OBC_OutputVoltage.ToPhysical(float64(m.xxx_OBC_OutputVoltage))
}

func (m *OBC_Response) SetOBC_OutputVoltage(v float64) *OBC_Response {
	m.xxx_OBC_OutputVoltage = uint16(Messages().OBC_Response.OBC_OutputVoltage.FromPhysical(v))
	return m
}

func (m *OBC_Response) RawOBC_OutputVoltage() uint16 {
	return m.xxx_OBC_OutputVoltage
}

func (m *OBC_Response) SetRawOBC_OutputVoltage(v uint16) *OBC_Response {
	m.xxx_OBC_OutputVoltage = uint16(Messages().OBC_Response.OBC_OutputVoltage.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBC_Response) OBC_OutputCurrent() float64 {
	return Messages().OBC_Response.OBC_OutputCurrent.ToPhysical(float64(m.xxx_OBC_OutputCurrent))
}

func (m *OBC_Response) SetOBC_OutputCurrent(v float64) *OBC_Response {
	m.xxx_OBC_OutputCurrent = uint16(Messages().OBC_Response.OBC_OutputCurrent.FromPhysical(v))
	return m
}

func (m *OBC_Response) RawOBC_OutputCurrent() uint16 {
	return m.xxx_OBC_OutputCurrent
}

func (m *OBC_Response) SetRawOBC_OutputCurrent(v uint16) *OBC_Response {
	m.xxx_OBC_OutputCurrent = uint16(Messages().OBC_Response.OBC_OutputCurrent.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBC_Response) OBC_HardwareProtection() OBC_Response_OBC_HardwareProtection {
	return m.xxx_OBC_HardwareProtection
}

func (m *OBC_Response) SetOBC_HardwareProtection(v OBC_Response_OBC_HardwareProtection) *OBC_Response {
	m.xxx_OBC_HardwareProtection = v
	return m
}

func (m *OBC_Response) OBC_TemperatureProtection() OBC_Response_OBC_TemperatureProtection {
	return m.xxx_OBC_TemperatureProtection
}

func (m *OBC_Response) SetOBC_TemperatureProtection(v OBC_Response_OBC_TemperatureProtection) *OBC_Response {
	m.xxx_OBC_TemperatureProtection = v
	return m
}

func (m *OBC_Response) OBC_InputVoltageStatus() OBC_Response_OBC_InputVoltageStatus {
	return m.xxx_OBC_InputVoltageStatus
}

func (m *OBC_Response) SetOBC_InputVoltageStatus(v OBC_Response_OBC_InputVoltageStatus) *OBC_Response {
	m.xxx_OBC_InputVoltageStatus = OBC_Response_OBC_InputVoltageStatus(Messages().OBC_Response.OBC_InputVoltageStatus.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBC_Response) OBC_OutputUnderVoltage() OBC_Response_OBC_OutputUnderVoltage {
	return m.xxx_OBC_OutputUnderVoltage
}

func (m *OBC_Response) SetOBC_OutputUnderVoltage(v OBC_Response_OBC_OutputUnderVoltage) *OBC_Response {
	m.xxx_OBC_OutputUnderVoltage = v
	return m
}

func (m *OBC_Response) OBC_OutputOverVoltage() OBC_Response_OBC_OutputOverVoltage {
	return m.xxx_OBC_OutputOverVoltage
}

func (m *OBC_Response) SetOBC_OutputOverVoltage(v OBC_Response_OBC_OutputOverVoltage) *OBC_Response {
	m.xxx_OBC_OutputOverVoltage = v
	return m
}

func (m *OBC_Response) OBC_OutputOvercurrent() OBC_Response_OBC_OutputOvercurrent {
	return m.xxx_OBC_OutputOvercurrent
}

func (m *OBC_Response) SetOBC_OutputOvercurrent(v OBC_Response_OBC_OutputOvercurrent) *OBC_Response {
	m.xxx_OBC_OutputOvercurrent = v
	return m
}

func (m *OBC_Response) OBC_OutputShortCircuit() OBC_Response_OBC_OutputShortCircuit {
	return m.xxx_OBC_OutputShortCircuit
}

func (m *OBC_Response) SetOBC_OutputShortCircuit(v OBC_Response_OBC_OutputShortCircuit) *OBC_Response {
	m.xxx_OBC_OutputShortCircuit = v
	return m
}

func (m *OBC_Response) OBC_CommunicationStatus() OBC_Response_OBC_CommunicationStatus {
	return m.xxx_OBC_CommunicationStatus
}

func (m *OBC_Response) SetOBC_CommunicationStatus(v OBC_Response_OBC_CommunicationStatus) *OBC_Response {
	m.xxx_OBC_CommunicationStatus = v
	return m
}

func (m *OBC_Response) OBC_WorkingStatus() OBC_Response_OBC_WorkingStatus {
	return m.xxx_OBC_WorkingStatus
}

func (m *OBC_Response) SetOBC_WorkingStatus(v OBC_Response_OBC_WorkingStatus) *OBC_Response {
	m.xxx_OBC_WorkingStatus = OBC_Response_OBC_WorkingStatus(Messages().OBC_Response.OBC_WorkingStatus.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBC_Response) OBC_ComOfInitialization() OBC_Response_OBC_ComOfInitialization {
	return m.xxx_OBC_ComOfInitialization
}

func (m *OBC_Response) SetOBC_ComOfInitialization(v OBC_Response_OBC_ComOfInitialization) *OBC_Response {
	m.xxx_OBC_ComOfInitialization = v
	return m
}

func (m *OBC_Response) OBC_FanWorkingEnable() OBC_Response_OBC_FanWorkingEnable {
	return m.xxx_OBC_FanWorkingEnable
}

func (m *OBC_Response) SetOBC_FanWorkingEnable(v OBC_Response_OBC_FanWorkingEnable) *OBC_Response {
	m.xxx_OBC_FanWorkingEnable = v
	return m
}

func (m *OBC_Response) OBC_CoolingPumpEnable() OBC_Response_OBC_CoolingPumpEnable {
	return m.xxx_OBC_CoolingPumpEnable
}

func (m *OBC_Response) SetOBC_CoolingPumpEnable(v OBC_Response_OBC_CoolingPumpEnable) *OBC_Response {
	m.xxx_OBC_CoolingPumpEnable = v
	return m
}

func (m *OBC_Response) OBC_CCSignalStatus() OBC_Response_OBC_CCSignalStatus {
	return m.xxx_OBC_CCSignalStatus
}

func (m *OBC_Response) SetOBC_CCSignalStatus(v OBC_Response_OBC_CCSignalStatus) *OBC_Response {
	m.xxx_OBC_CCSignalStatus = OBC_Response_OBC_CCSignalStatus(Messages().OBC_Response.OBC_CCSignalStatus.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBC_Response) OBC_CPSignalStatus() OBC_Response_OBC_CPSignalStatus {
	return m.xxx_OBC_CPSignalStatus
}

func (m *OBC_Response) SetOBC_CPSignalStatus(v OBC_Response_OBC_CPSignalStatus) *OBC_Response {
	m.xxx_OBC_CPSignalStatus = v
	return m
}

func (m *OBC_Response) OBC_SocketOvheatFault() OBC_Response_OBC_SocketOvheatFault {
	return m.xxx_OBC_SocketOvheatFault
}

func (m *OBC_Response) SetOBC_SocketOvheatFault(v OBC_Response_OBC_SocketOvheatFault) *OBC_Response {
	m.xxx_OBC_SocketOvheatFault = v
	return m
}

func (m *OBC_Response) OBC_ElectronicLockState() OBC_Response_OBC_ElectronicLockState {
	return m.xxx_OBC_ElectronicLockState
}

func (m *OBC_Response) SetOBC_ElectronicLockState(v OBC_Response_OBC_ElectronicLockState) *OBC_Response {
	m.xxx_OBC_ElectronicLockState = OBC_Response_OBC_ElectronicLockState(Messages().OBC_Response.OBC_ElectronicLockState.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBC_Response) OBC_S2SwitchControlBitStatus() OBC_Response_OBC_S2SwitchControlBitStatus {
	return m.xxx_OBC_S2SwitchControlBitStatus
}

func (m *OBC_Response) SetOBC_S2SwitchControlBitStatus(v OBC_Response_OBC_S2SwitchControlBitStatus) *OBC_Response {
	m.xxx_OBC_S2SwitchControlBitStatus = v
	return m
}

func (m *OBC_Response) OBC_Temperature() float64 {
	return Messages().OBC_Response.OBC_Temperature.ToPhysical(float64(m.xxx_OBC_Temperature))
}

func (m *OBC_Response) SetOBC_Temperature(v float64) *OBC_Response {
	m.xxx_OBC_Temperature = uint8(Messages().OBC_Response.OBC_Temperature.FromPhysical(v))
	return m
}

func (m *OBC_Response) RawOBC_Temperature() uint8 {
	return m.xxx_OBC_Temperature
}

func (m *OBC_Response) SetRawOBC_Temperature(v uint8) *OBC_Response {
	m.xxx_OBC_Temperature = uint8(Messages().OBC_Response.OBC_Temperature.SaturatedCastUnsigned(uint64(v)))
	return m
}

// OBC_Response_OBC_HardwareProtection models the OBC_HardwareProtection signal of the OBC_Response message.
type OBC_Response_OBC_HardwareProtection bool

// Value descriptions for the OBC_HardwareProtection signal of the OBC_Response message.
const (
	OBC_Response_OBC_HardwareProtection_Normal             OBC_Response_OBC_HardwareProtection = false
	OBC_Response_OBC_HardwareProtection_HardwareProtection OBC_Response_OBC_HardwareProtection = true
)

func (v OBC_Response_OBC_HardwareProtection) String() string {
	switch bool(v) {
	case false:
		return "Normal"
	case true:
		return "HardwareProtection"
	}
	return fmt.Sprintf("OBC_Response_OBC_HardwareProtection(%t)", v)
}

// OBC_Response_OBC_TemperatureProtection models the OBC_TemperatureProtection signal of the OBC_Response message.
type OBC_Response_OBC_TemperatureProtection bool

// Value descriptions for the OBC_TemperatureProtection signal of the OBC_Response message.
const (
	OBC_Response_OBC_TemperatureProtection_Normal                        OBC_Response_OBC_TemperatureProtection = false
	OBC_Response_OBC_TemperatureProtection_InternalTemperatureProtection OBC_Response_OBC_TemperatureProtection = true
)

func (v OBC_Response_OBC_TemperatureProtection) String() string {
	switch bool(v) {
	case false:
		return "Normal"
	case true:
		return "InternalTemperatureProtection"
	}
	return fmt.Sprintf("OBC_Response_OBC_TemperatureProtection(%t)", v)
}

// OBC_Response_OBC_InputVoltageStatus models the OBC_InputVoltageStatus signal of the OBC_Response message.
type OBC_Response_OBC_InputVoltageStatus uint8

// Value descriptions for the OBC_InputVoltageStatus signal of the OBC_Response message.
const (
	OBC_Response_OBC_InputVoltageStatus_TheVoltageIsNormal OBC_Response_OBC_InputVoltageStatus = 0
	OBC_Response_OBC_InputVoltageStatus_InputUnderVoltage  OBC_Response_OBC_InputVoltageStatus = 1
	OBC_Response_OBC_InputVoltageStatus_InputOverVoltage   OBC_Response_OBC_InputVoltageStatus = 2
	OBC_Response_OBC_InputVoltageStatus_NoInputVoltage     OBC_Response_OBC_InputVoltageStatus = 3
)

func (v OBC_Response_OBC_InputVoltageStatus) String() string {
	switch v {
	case 0:
		return "TheVoltageIsNormal"
	case 1:
		return "InputUnderVoltage"
	case 2:
		return "InputOverVoltage"
	case 3:
		return "NoInputVoltage"
	default:
		return fmt.Sprintf("OBC_Response_OBC_InputVoltageStatus(%d)", v)
	}
}

// OBC_Response_OBC_OutputUnderVoltage models the OBC_OutputUnderVoltage signal of the OBC_Response message.
type OBC_Response_OBC_OutputUnderVoltage bool

// Value descriptions for the OBC_OutputUnderVoltage signal of the OBC_Response message.
const (
	OBC_Response_OBC_OutputUnderVoltage_Normal OBC_Response_OBC_OutputUnderVoltage = false
	OBC_Response_OBC_OutputUnderVoltage_Fault  OBC_Response_OBC_OutputUnderVoltage = true
)

func (v OBC_Response_OBC_OutputUnderVoltage) String() string {
	switch bool(v) {
	case false:
		return "Normal"
	case true:
		return "Fault"
	}
	return fmt.Sprintf("OBC_Response_OBC_OutputUnderVoltage(%t)", v)
}

// OBC_Response_OBC_OutputOverVoltage models the OBC_OutputOverVoltage signal of the OBC_Response message.
type OBC_Response_OBC_OutputOverVoltage bool

// Value descriptions for the OBC_OutputOverVoltage signal of the OBC_Response message.
const (
	OBC_Response_OBC_OutputOverVoltage_Normal OBC_Response_OBC_OutputOverVoltage = false
	OBC_Response_OBC_OutputOverVoltage_Fault  OBC_Response_OBC_OutputOverVoltage = true
)

func (v OBC_Response_OBC_OutputOverVoltage) String() string {
	switch bool(v) {
	case false:
		return "Normal"
	case true:
		return "Fault"
	}
	return fmt.Sprintf("OBC_Response_OBC_OutputOverVoltage(%t)", v)
}

// OBC_Response_OBC_OutputOvercurrent models the OBC_OutputOvercurrent signal of the OBC_Response message.
type OBC_Response_OBC_OutputOvercurrent bool

// Value descriptions for the OBC_OutputOvercurrent signal of the OBC_Response message.
const (
	OBC_Response_OBC_OutputOvercurrent_Normal OBC_Response_OBC_OutputOvercurrent = false
	OBC_Response_OBC_OutputOvercurrent_Fault  OBC_Response_OBC_OutputOvercurrent = true
)

func (v OBC_Response_OBC_OutputOvercurrent) String() string {
	switch bool(v) {
	case false:
		return "Normal"
	case true:
		return "Fault"
	}
	return fmt.Sprintf("OBC_Response_OBC_OutputOvercurrent(%t)", v)
}

// OBC_Response_OBC_OutputShortCircuit models the OBC_OutputShortCircuit signal of the OBC_Response message.
type OBC_Response_OBC_OutputShortCircuit bool

// Value descriptions for the OBC_OutputShortCircuit signal of the OBC_Response message.
const (
	OBC_Response_OBC_OutputShortCircuit_Normal OBC_Response_OBC_OutputShortCircuit = false
	OBC_Response_OBC_OutputShortCircuit_Fault  OBC_Response_OBC_OutputShortCircuit = true
)

func (v OBC_Response_OBC_OutputShortCircuit) String() string {
	switch bool(v) {
	case false:
		return "Normal"
	case true:
		return "Fault"
	}
	return fmt.Sprintf("OBC_Response_OBC_OutputShortCircuit(%t)", v)
}

// OBC_Response_OBC_CommunicationStatus models the OBC_CommunicationStatus signal of the OBC_Response message.
type OBC_Response_OBC_CommunicationStatus bool

// Value descriptions for the OBC_CommunicationStatus signal of the OBC_Response message.
const (
	OBC_Response_OBC_CommunicationStatus_CommunicationIsNormal       OBC_Response_OBC_CommunicationStatus = false
	OBC_Response_OBC_CommunicationStatus_ReceiveCommunicationTimeout OBC_Response_OBC_CommunicationStatus = true
)

func (v OBC_Response_OBC_CommunicationStatus) String() string {
	switch bool(v) {
	case false:
		return "CommunicationIsNormal"
	case true:
		return "ReceiveCommunicationTimeout"
	}
	return fmt.Sprintf("OBC_Response_OBC_CommunicationStatus(%t)", v)
}

// OBC_Response_OBC_WorkingStatus models the OBC_WorkingStatus signal of the OBC_Response message.
type OBC_Response_OBC_WorkingStatus uint8

// Value descriptions for the OBC_WorkingStatus signal of the OBC_Response message.
const (
	OBC_Response_OBC_WorkingStatus_Undefined     OBC_Response_OBC_WorkingStatus = 0
	OBC_Response_OBC_WorkingStatus_Work          OBC_Response_OBC_WorkingStatus = 1
	OBC_Response_OBC_WorkingStatus_Stop          OBC_Response_OBC_WorkingStatus = 2
	OBC_Response_OBC_WorkingStatus_StopOrStandBy OBC_Response_OBC_WorkingStatus = 3
)

func (v OBC_Response_OBC_WorkingStatus) String() string {
	switch v {
	case 0:
		return "Undefined"
	case 1:
		return "Work"
	case 2:
		return "Stop"
	case 3:
		return "StopOrStandBy"
	default:
		return fmt.Sprintf("OBC_Response_OBC_WorkingStatus(%d)", v)
	}
}

// OBC_Response_OBC_ComOfInitialization models the OBC_ComOfInitialization signal of the OBC_Response message.
type OBC_Response_OBC_ComOfInitialization bool

// Value descriptions for the OBC_ComOfInitialization signal of the OBC_Response message.
const (
	OBC_Response_OBC_ComOfInitialization_IsNotComplete OBC_Response_OBC_ComOfInitialization = false
	OBC_Response_OBC_ComOfInitialization_Complete      OBC_Response_OBC_ComOfInitialization = true
)

func (v OBC_Response_OBC_ComOfInitialization) String() string {
	switch bool(v) {
	case false:
		return "IsNotComplete"
	case true:
		return "Complete"
	}
	return fmt.Sprintf("OBC_Response_OBC_ComOfInitialization(%t)", v)
}

// OBC_Response_OBC_FanWorkingEnable models the OBC_FanWorkingEnable signal of the OBC_Response message.
type OBC_Response_OBC_FanWorkingEnable bool

// Value descriptions for the OBC_FanWorkingEnable signal of the OBC_Response message.
const (
	OBC_Response_OBC_FanWorkingEnable_Close OBC_Response_OBC_FanWorkingEnable = false
	OBC_Response_OBC_FanWorkingEnable_Open  OBC_Response_OBC_FanWorkingEnable = true
)

func (v OBC_Response_OBC_FanWorkingEnable) String() string {
	switch bool(v) {
	case false:
		return "Close"
	case true:
		return "Open"
	}
	return fmt.Sprintf("OBC_Response_OBC_FanWorkingEnable(%t)", v)
}

// OBC_Response_OBC_CoolingPumpEnable models the OBC_CoolingPumpEnable signal of the OBC_Response message.
type OBC_Response_OBC_CoolingPumpEnable bool

// Value descriptions for the OBC_CoolingPumpEnable signal of the OBC_Response message.
const (
	OBC_Response_OBC_CoolingPumpEnable_Close OBC_Response_OBC_CoolingPumpEnable = false
	OBC_Response_OBC_CoolingPumpEnable_Open  OBC_Response_OBC_CoolingPumpEnable = true
)

func (v OBC_Response_OBC_CoolingPumpEnable) String() string {
	switch bool(v) {
	case false:
		return "Close"
	case true:
		return "Open"
	}
	return fmt.Sprintf("OBC_Response_OBC_CoolingPumpEnable(%t)", v)
}

// OBC_Response_OBC_CCSignalStatus models the OBC_CCSignalStatus signal of the OBC_Response message.
type OBC_Response_OBC_CCSignalStatus uint8

// Value descriptions for the OBC_CCSignalStatus signal of the OBC_Response message.
const (
	OBC_Response_OBC_CCSignalStatus_NotConnected             OBC_Response_OBC_CCSignalStatus = 0
	OBC_Response_OBC_CCSignalStatus_HalfConnected            OBC_Response_OBC_CCSignalStatus = 1
	OBC_Response_OBC_CCSignalStatus_NormalConnected          OBC_Response_OBC_CCSignalStatus = 2
	OBC_Response_OBC_CCSignalStatus_ResistanceDetectionError OBC_Response_OBC_CCSignalStatus = 3
)

func (v OBC_Response_OBC_CCSignalStatus) String() string {
	switch v {
	case 0:
		return "NotConnected"
	case 1:
		return "HalfConnected"
	case 2:
		return "NormalConnected"
	case 3:
		return "ResistanceDetectionError"
	default:
		return fmt.Sprintf("OBC_Response_OBC_CCSignalStatus(%d)", v)
	}
}

// OBC_Response_OBC_CPSignalStatus models the OBC_CPSignalStatus signal of the OBC_Response message.
type OBC_Response_OBC_CPSignalStatus bool

// Value descriptions for the OBC_CPSignalStatus signal of the OBC_Response message.
const (
	OBC_Response_OBC_CPSignalStatus_NoCPSignalWasDetected OBC_Response_OBC_CPSignalStatus = false
	OBC_Response_OBC_CPSignalStatus_Normal                OBC_Response_OBC_CPSignalStatus = true
)

func (v OBC_Response_OBC_CPSignalStatus) String() string {
	switch bool(v) {
	case false:
		return "NoCPSignalWasDetected"
	case true:
		return "Normal"
	}
	return fmt.Sprintf("OBC_Response_OBC_CPSignalStatus(%t)", v)
}

// OBC_Response_OBC_SocketOvheatFault models the OBC_SocketOvheatFault signal of the OBC_Response message.
type OBC_Response_OBC_SocketOvheatFault bool

// Value descriptions for the OBC_SocketOvheatFault signal of the OBC_Response message.
const (
	OBC_Response_OBC_SocketOvheatFault_Normal                   OBC_Response_OBC_SocketOvheatFault = false
	OBC_Response_OBC_SocketOvheatFault_ChargingSocketIsOverheat OBC_Response_OBC_SocketOvheatFault = true
)

func (v OBC_Response_OBC_SocketOvheatFault) String() string {
	switch bool(v) {
	case false:
		return "Normal"
	case true:
		return "ChargingSocketIsOverheat"
	}
	return fmt.Sprintf("OBC_Response_OBC_SocketOvheatFault(%t)", v)
}

// OBC_Response_OBC_ElectronicLockState models the OBC_ElectronicLockState signal of the OBC_Response message.
type OBC_Response_OBC_ElectronicLockState uint8

// Value descriptions for the OBC_ElectronicLockState signal of the OBC_Response message.
const (
	OBC_Response_OBC_ElectronicLockState_InJudgment  OBC_Response_OBC_ElectronicLockState = 0
	OBC_Response_OBC_ElectronicLockState_Locked      OBC_Response_OBC_ElectronicLockState = 1
	OBC_Response_OBC_ElectronicLockState_Unlocked    OBC_Response_OBC_ElectronicLockState = 2
	OBC_Response_OBC_ElectronicLockState_UnlockFault OBC_Response_OBC_ElectronicLockState = 3
	OBC_Response_OBC_ElectronicLockState_LockedFault OBC_Response_OBC_ElectronicLockState = 4
)

func (v OBC_Response_OBC_ElectronicLockState) String() string {
	switch v {
	case 0:
		return "InJudgment"
	case 1:
		return "Locked"
	case 2:
		return "Unlocked"
	case 3:
		return "UnlockFault"
	case 4:
		return "LockedFault"
	default:
		return fmt.Sprintf("OBC_Response_OBC_ElectronicLockState(%d)", v)
	}
}

// OBC_Response_OBC_S2SwitchControlBitStatus models the OBC_S2SwitchControlBitStatus signal of the OBC_Response message.
type OBC_Response_OBC_S2SwitchControlBitStatus bool

// Value descriptions for the OBC_S2SwitchControlBitStatus signal of the OBC_Response message.
const (
	OBC_Response_OBC_S2SwitchControlBitStatus_SwitchOff OBC_Response_OBC_S2SwitchControlBitStatus = false
	OBC_Response_OBC_S2SwitchControlBitStatus_CloseUp   OBC_Response_OBC_S2SwitchControlBitStatus = true
)

func (v OBC_Response_OBC_S2SwitchControlBitStatus) String() string {
	switch bool(v) {
	case false:
		return "SwitchOff"
	case true:
		return "CloseUp"
	}
	return fmt.Sprintf("OBC_Response_OBC_S2SwitchControlBitStatus(%t)", v)
}

// Frame returns a CAN frame representing the message.
func (m *OBC_Response) Frame() can.Frame {
	md := Messages().OBC_Response
	f := can.Frame{ID: md.ID, IsExtended: md.IsExtended, Length: md.Length}
	md.OBC_OutputVoltage.MarshalUnsigned(&f.Data, uint64(m.xxx_OBC_OutputVoltage))
	md.OBC_OutputCurrent.MarshalUnsigned(&f.Data, uint64(m.xxx_OBC_OutputCurrent))
	md.OBC_HardwareProtection.MarshalBool(&f.Data, bool(m.xxx_OBC_HardwareProtection))
	md.OBC_TemperatureProtection.MarshalBool(&f.Data, bool(m.xxx_OBC_TemperatureProtection))
	md.OBC_InputVoltageStatus.MarshalUnsigned(&f.Data, uint64(m.xxx_OBC_InputVoltageStatus))
	md.OBC_OutputUnderVoltage.MarshalBool(&f.Data, bool(m.xxx_OBC_OutputUnderVoltage))
	md.OBC_OutputOverVoltage.MarshalBool(&f.Data, bool(m.xxx_OBC_OutputOverVoltage))
	md.OBC_OutputOvercurrent.MarshalBool(&f.Data, bool(m.xxx_OBC_OutputOvercurrent))
	md.OBC_OutputShortCircuit.MarshalBool(&f.Data, bool(m.xxx_OBC_OutputShortCircuit))
	md.OBC_CommunicationStatus.MarshalBool(&f.Data, bool(m.xxx_OBC_CommunicationStatus))
	md.OBC_WorkingStatus.MarshalUnsigned(&f.Data, uint64(m.xxx_OBC_WorkingStatus))
	md.OBC_ComOfInitialization.MarshalBool(&f.Data, bool(m.xxx_OBC_ComOfInitialization))
	md.OBC_FanWorkingEnable.MarshalBool(&f.Data, bool(m.xxx_OBC_FanWorkingEnable))
	md.OBC_CoolingPumpEnable.MarshalBool(&f.Data, bool(m.xxx_OBC_CoolingPumpEnable))
	md.OBC_CCSignalStatus.MarshalUnsigned(&f.Data, uint64(m.xxx_OBC_CCSignalStatus))
	md.OBC_CPSignalStatus.MarshalBool(&f.Data, bool(m.xxx_OBC_CPSignalStatus))
	md.OBC_SocketOvheatFault.MarshalBool(&f.Data, bool(m.xxx_OBC_SocketOvheatFault))
	md.OBC_ElectronicLockState.MarshalUnsigned(&f.Data, uint64(m.xxx_OBC_ElectronicLockState))
	md.OBC_S2SwitchControlBitStatus.MarshalBool(&f.Data, bool(m.xxx_OBC_S2SwitchControlBitStatus))
	md.OBC_Temperature.MarshalUnsigned(&f.Data, uint64(m.xxx_OBC_Temperature))
	return f
}

// MarshalFrame encodes the message as a CAN frame.
func (m *OBC_Response) MarshalFrame() (can.Frame, error) {
	return m.Frame(), nil
}

// UnmarshalFrame decodes the message from a CAN frame.
func (m *OBC_Response) UnmarshalFrame(f can.Frame) error {
	md := Messages().OBC_Response
	switch {
	case f.ID != md.ID:
		return fmt.Errorf(
			"unmarshal OBC_Response: expects ID 419385573 (got %s with ID %d)", f.String(), f.ID,
		)
	case f.Length != md.Length:
		return fmt.Errorf(
			"unmarshal OBC_Response: expects length 8 (got %s with length %d)", f.String(), f.Length,
		)
	case f.IsRemote:
		return fmt.Errorf(
			"unmarshal OBC_Response: expects non-remote frame (got remote frame %s)", f.String(),
		)
	case f.IsExtended != md.IsExtended:
		return fmt.Errorf(
			"unmarshal OBC_Response: expects extended ID (got %s with standard ID)", f.String(),
		)
	}
	m.xxx_OBC_OutputVoltage = uint16(md.OBC_OutputVoltage.UnmarshalUnsigned(f.Data))
	m.xxx_OBC_OutputCurrent = uint16(md.OBC_OutputCurrent.UnmarshalUnsigned(f.Data))
	m.xxx_OBC_HardwareProtection = OBC_Response_OBC_HardwareProtection(md.OBC_HardwareProtection.UnmarshalBool(f.Data))
	m.xxx_OBC_TemperatureProtection = OBC_Response_OBC_TemperatureProtection(md.OBC_TemperatureProtection.UnmarshalBool(f.Data))
	m.xxx_OBC_InputVoltageStatus = OBC_Response_OBC_InputVoltageStatus(md.OBC_InputVoltageStatus.UnmarshalUnsigned(f.Data))
	m.xxx_OBC_OutputUnderVoltage = OBC_Response_OBC_OutputUnderVoltage(md.OBC_OutputUnderVoltage.UnmarshalBool(f.Data))
	m.xxx_OBC_OutputOverVoltage = OBC_Response_OBC_OutputOverVoltage(md.OBC_OutputOverVoltage.UnmarshalBool(f.Data))
	m.xxx_OBC_OutputOvercurrent = OBC_Response_OBC_OutputOvercurrent(md.OBC_OutputOvercurrent.UnmarshalBool(f.Data))
	m.xxx_OBC_OutputShortCircuit = OBC_Response_OBC_OutputShortCircuit(md.OBC_OutputShortCircuit.UnmarshalBool(f.Data))
	m.xxx_OBC_CommunicationStatus = OBC_Response_OBC_CommunicationStatus(md.OBC_CommunicationStatus.UnmarshalBool(f.Data))
	m.xxx_OBC_WorkingStatus = OBC_Response_OBC_WorkingStatus(md.OBC_WorkingStatus.UnmarshalUnsigned(f.Data))
	m.xxx_OBC_ComOfInitialization = OBC_Response_OBC_ComOfInitialization(md.OBC_ComOfInitialization.UnmarshalBool(f.Data))
	m.xxx_OBC_FanWorkingEnable = OBC_Response_OBC_FanWorkingEnable(md.OBC_FanWorkingEnable.UnmarshalBool(f.Data))
	m.xxx_OBC_CoolingPumpEnable = OBC_Response_OBC_CoolingPumpEnable(md.OBC_CoolingPumpEnable.UnmarshalBool(f.Data))
	m.xxx_OBC_CCSignalStatus = OBC_Response_OBC_CCSignalStatus(md.OBC_CCSignalStatus.UnmarshalUnsigned(f.Data))
	m.xxx_OBC_CPSignalStatus = OBC_Response_OBC_CPSignalStatus(md.OBC_CPSignalStatus.UnmarshalBool(f.Data))
	m.xxx_OBC_SocketOvheatFault = OBC_Response_OBC_SocketOvheatFault(md.OBC_SocketOvheatFault.UnmarshalBool(f.Data))
	m.xxx_OBC_ElectronicLockState = OBC_Response_OBC_ElectronicLockState(md.OBC_ElectronicLockState.UnmarshalUnsigned(f.Data))
	m.xxx_OBC_S2SwitchControlBitStatus = OBC_Response_OBC_S2SwitchControlBitStatus(md.OBC_S2SwitchControlBitStatus.UnmarshalBool(f.Data))
	m.xxx_OBC_Temperature = uint8(md.OBC_Temperature.UnmarshalUnsigned(f.Data))
	return nil
}

// Nodes returns the createvan node descriptors.
func Nodes() *NodesDescriptor {
	return nd
}

// NodesDescriptor contains all createvan node descriptors.
type NodesDescriptor struct {
	CReATE_ECU *descriptor.Node
}

// Messages returns the createvan message descriptors.
func Messages() *MessagesDescriptor {
	return md
}

// MessagesDescriptor contains all createvan message descriptors.
type MessagesDescriptor struct {
	IVI_DashboardStatus        *IVI_DashboardStatusDescriptor
	IVI_CornerMotorStatus      *IVI_CornerMotorStatusDescriptor
	IOV_FrontLeftCornerMotor1  *IOV_FrontLeftCornerMotor1Descriptor
	IOV_FrontLeftCornerMotor2  *IOV_FrontLeftCornerMotor2Descriptor
	IOV_FrontLeftCornerMotor3  *IOV_FrontLeftCornerMotor3Descriptor
	IOV_FrontLeftCornerMotor4  *IOV_FrontLeftCornerMotor4Descriptor
	IOV_FrontRightCornerMotor1 *IOV_FrontRightCornerMotor1Descriptor
	IOV_FrontRightCornerMotor2 *IOV_FrontRightCornerMotor2Descriptor
	IOV_FrontRightCornerMotor3 *IOV_FrontRightCornerMotor3Descriptor
	IOV_FrontRightCornerMotor4 *IOV_FrontRightCornerMotor4Descriptor
	IOV_BackLeftCornerMotor1   *IOV_BackLeftCornerMotor1Descriptor
	IOV_BackLeftCornerMotor2   *IOV_BackLeftCornerMotor2Descriptor
	IOV_BackLeftCornerMotor3   *IOV_BackLeftCornerMotor3Descriptor
	IOV_BackLeftCornerMotor4   *IOV_BackLeftCornerMotor4Descriptor
	IOV_BackRightCornerMotor1  *IOV_BackRightCornerMotor1Descriptor
	IOV_BackRightCornerMotor2  *IOV_BackRightCornerMotor2Descriptor
	IOV_BackRightCornerMotor3  *IOV_BackRightCornerMotor3Descriptor
	IOV_BackRightCornerMotor4  *IOV_BackRightCornerMotor4Descriptor
	OBC_Command                *OBC_CommandDescriptor
	BMS_DALY_SoCStatus         *BMS_DALY_SoCStatusDescriptor
	BMS_DALY_RangeVoltage      *BMS_DALY_RangeVoltageDescriptor
	BMS_DALY_RangeTemperature  *BMS_DALY_RangeTemperatureDescriptor
	BMS_DALY_MOSStatus         *BMS_DALY_MOSStatusDescriptor
	BMS_DALY_StatusInformation *BMS_DALY_StatusInformationDescriptor
	BMS_DALY_MonoCellVoltage   *BMS_DALY_MonoCellVoltageDescriptor
	BMS_DALY_MonoCellTemp      *BMS_DALY_MonoCellTempDescriptor
	OBC_Response               *OBC_ResponseDescriptor
}

// UnmarshalFrame unmarshals the provided createvan CAN frame.
func (md *MessagesDescriptor) UnmarshalFrame(f can.Frame) (generated.Message, error) {
	switch f.ID {
	case md.IVI_DashboardStatus.ID:
		var msg IVI_DashboardStatus
		if err := msg.UnmarshalFrame(f); err != nil {
			return nil, fmt.Errorf("unmarshal createvan frame: %w", err)
		}
		return &msg, nil
	case md.IVI_CornerMotorStatus.ID:
		var msg IVI_CornerMotorStatus
		if err := msg.UnmarshalFrame(f); err != nil {
			return nil, fmt.Errorf("unmarshal createvan frame: %w", err)
		}
		return &msg, nil
	case md.IOV_FrontLeftCornerMotor1.ID:
		var msg IOV_FrontLeftCornerMotor1
		if err := msg.UnmarshalFrame(f); err != nil {
			return nil, fmt.Errorf("unmarshal createvan frame: %w", err)
		}
		return &msg, nil
	case md.IOV_FrontLeftCornerMotor2.ID:
		var msg IOV_FrontLeftCornerMotor2
		if err := msg.UnmarshalFrame(f); err != nil {
			return nil, fmt.Errorf("unmarshal createvan frame: %w", err)
		}
		return &msg, nil
	case md.IOV_FrontLeftCornerMotor3.ID:
		var msg IOV_FrontLeftCornerMotor3
		if err := msg.UnmarshalFrame(f); err != nil {
			return nil, fmt.Errorf("unmarshal createvan frame: %w", err)
		}
		return &msg, nil
	case md.IOV_FrontLeftCornerMotor4.ID:
		var msg IOV_FrontLeftCornerMotor4
		if err := msg.UnmarshalFrame(f); err != nil {
			return nil, fmt.Errorf("unmarshal createvan frame: %w", err)
		}
		return &msg, nil
	case md.IOV_FrontRightCornerMotor1.ID:
		var msg IOV_FrontRightCornerMotor1
		if err := msg.UnmarshalFrame(f); err != nil {
			return nil, fmt.Errorf("unmarshal createvan frame: %w", err)
		}
		return &msg, nil
	case md.IOV_FrontRightCornerMotor2.ID:
		var msg IOV_FrontRightCornerMotor2
		if err := msg.UnmarshalFrame(f); err != nil {
			return nil, fmt.Errorf("unmarshal createvan frame: %w", err)
		}
		return &msg, nil
	case md.IOV_FrontRightCornerMotor3.ID:
		var msg IOV_FrontRightCornerMotor3
		if err := msg.UnmarshalFrame(f); err != nil {
			return nil, fmt.Errorf("unmarshal createvan frame: %w", err)
		}
		return &msg, nil
	case md.IOV_FrontRightCornerMotor4.ID:
		var msg IOV_FrontRightCornerMotor4
		if err := msg.UnmarshalFrame(f); err != nil {
			return nil, fmt.Errorf("unmarshal createvan frame: %w", err)
		}
		return &msg, nil
	case md.IOV_BackLeftCornerMotor1.ID:
		var msg IOV_BackLeftCornerMotor1
		if err := msg.UnmarshalFrame(f); err != nil {
			return nil, fmt.Errorf("unmarshal createvan frame: %w", err)
		}
		return &msg, nil
	case md.IOV_BackLeftCornerMotor2.ID:
		var msg IOV_BackLeftCornerMotor2
		if err := msg.UnmarshalFrame(f); err != nil {
			return nil, fmt.Errorf("unmarshal createvan frame: %w", err)
		}
		return &msg, nil
	case md.IOV_BackLeftCornerMotor3.ID:
		var msg IOV_BackLeftCornerMotor3
		if err := msg.UnmarshalFrame(f); err != nil {
			return nil, fmt.Errorf("unmarshal createvan frame: %w", err)
		}
		return &msg, nil
	case md.IOV_BackLeftCornerMotor4.ID:
		var msg IOV_BackLeftCornerMotor4
		if err := msg.UnmarshalFrame(f); err != nil {
			return nil, fmt.Errorf("unmarshal createvan frame: %w", err)
		}
		return &msg, nil
	case md.IOV_BackRightCornerMotor1.ID:
		var msg IOV_BackRightCornerMotor1
		if err := msg.UnmarshalFrame(f); err != nil {
			return nil, fmt.Errorf("unmarshal createvan frame: %w", err)
		}
		return &msg, nil
	case md.IOV_BackRightCornerMotor2.ID:
		var msg IOV_BackRightCornerMotor2
		if err := msg.UnmarshalFrame(f); err != nil {
			return nil, fmt.Errorf("unmarshal createvan frame: %w", err)
		}
		return &msg, nil
	case md.IOV_BackRightCornerMotor3.ID:
		var msg IOV_BackRightCornerMotor3
		if err := msg.UnmarshalFrame(f); err != nil {
			return nil, fmt.Errorf("unmarshal createvan frame: %w", err)
		}
		return &msg, nil
	case md.IOV_BackRightCornerMotor4.ID:
		var msg IOV_BackRightCornerMotor4
		if err := msg.UnmarshalFrame(f); err != nil {
			return nil, fmt.Errorf("unmarshal createvan frame: %w", err)
		}
		return &msg, nil
	case md.OBC_Command.ID:
		var msg OBC_Command
		if err := msg.UnmarshalFrame(f); err != nil {
			return nil, fmt.Errorf("unmarshal createvan frame: %w", err)
		}
		return &msg, nil
	case md.BMS_DALY_SoCStatus.ID:
		var msg BMS_DALY_SoCStatus
		if err := msg.UnmarshalFrame(f); err != nil {
			return nil, fmt.Errorf("unmarshal createvan frame: %w", err)
		}
		return &msg, nil
	case md.BMS_DALY_RangeVoltage.ID:
		var msg BMS_DALY_RangeVoltage
		if err := msg.UnmarshalFrame(f); err != nil {
			return nil, fmt.Errorf("unmarshal createvan frame: %w", err)
		}
		return &msg, nil
	case md.BMS_DALY_RangeTemperature.ID:
		var msg BMS_DALY_RangeTemperature
		if err := msg.UnmarshalFrame(f); err != nil {
			return nil, fmt.Errorf("unmarshal createvan frame: %w", err)
		}
		return &msg, nil
	case md.BMS_DALY_MOSStatus.ID:
		var msg BMS_DALY_MOSStatus
		if err := msg.UnmarshalFrame(f); err != nil {
			return nil, fmt.Errorf("unmarshal createvan frame: %w", err)
		}
		return &msg, nil
	case md.BMS_DALY_StatusInformation.ID:
		var msg BMS_DALY_StatusInformation
		if err := msg.UnmarshalFrame(f); err != nil {
			return nil, fmt.Errorf("unmarshal createvan frame: %w", err)
		}
		return &msg, nil
	case md.BMS_DALY_MonoCellVoltage.ID:
		var msg BMS_DALY_MonoCellVoltage
		if err := msg.UnmarshalFrame(f); err != nil {
			return nil, fmt.Errorf("unmarshal createvan frame: %w", err)
		}
		return &msg, nil
	case md.BMS_DALY_MonoCellTemp.ID:
		var msg BMS_DALY_MonoCellTemp
		if err := msg.UnmarshalFrame(f); err != nil {
			return nil, fmt.Errorf("unmarshal createvan frame: %w", err)
		}
		return &msg, nil
	case md.OBC_Response.ID:
		var msg OBC_Response
		if err := msg.UnmarshalFrame(f); err != nil {
			return nil, fmt.Errorf("unmarshal createvan frame: %w", err)
		}
		return &msg, nil
	default:
		return nil, fmt.Errorf("unmarshal createvan frame: ID not in database: %d", f.ID)
	}
}

type IVI_DashboardStatusDescriptor struct {
	*descriptor.Message
}

type IVI_CornerMotorStatusDescriptor struct {
	*descriptor.Message
}

type IOV_FrontLeftCornerMotor1Descriptor struct {
	*descriptor.Message
	VESC_StatusERPM1      *descriptor.Signal
	VESC_StatusCurrent1   *descriptor.Signal
	VESC_StatusDutyCycle1 *descriptor.Signal
}

type IOV_FrontLeftCornerMotor2Descriptor struct {
	*descriptor.Message
	VESC_StatusAmpHours1    *descriptor.Signal
	VESC_StatusAmpHoursChg1 *descriptor.Signal
}

type IOV_FrontLeftCornerMotor3Descriptor struct {
	*descriptor.Message
	VESC_StatusWattHours1    *descriptor.Signal
	VESC_StatusWattHoursChg1 *descriptor.Signal
}

type IOV_FrontLeftCornerMotor4Descriptor struct {
	*descriptor.Message
	VESC_StatusTemFET1    *descriptor.Signal
	VESC_StatusTempMotor1 *descriptor.Signal
	VESC_StatusCurrentIn1 *descriptor.Signal
	VESC_StatusPIDPos1    *descriptor.Signal
}

type IOV_FrontRightCornerMotor1Descriptor struct {
	*descriptor.Message
	VESC_StatusERPM2      *descriptor.Signal
	VESC_StatusCurrent2   *descriptor.Signal
	VESC_StatusDutyCycle2 *descriptor.Signal
}

type IOV_FrontRightCornerMotor2Descriptor struct {
	*descriptor.Message
	VESC_StatusAmpHours2    *descriptor.Signal
	VESC_StatusAmpHoursChg2 *descriptor.Signal
}

type IOV_FrontRightCornerMotor3Descriptor struct {
	*descriptor.Message
	VESC_StatusWattHours2    *descriptor.Signal
	VESC_StatusWattHoursChg2 *descriptor.Signal
}

type IOV_FrontRightCornerMotor4Descriptor struct {
	*descriptor.Message
	VESC_StatusTempFET2   *descriptor.Signal
	VESC_StatusTempMotor2 *descriptor.Signal
	VESC_StatusCurrentIn2 *descriptor.Signal
	VESC_StatusPIDPos2    *descriptor.Signal
}

type IOV_BackLeftCornerMotor1Descriptor struct {
	*descriptor.Message
	VESC_StatusERPM3      *descriptor.Signal
	VESC_StatusCurrent3   *descriptor.Signal
	VESC_StatusDutyCycle3 *descriptor.Signal
}

type IOV_BackLeftCornerMotor2Descriptor struct {
	*descriptor.Message
	VESC_StatusAmpHours3    *descriptor.Signal
	VESC_StatusAmpHoursChg3 *descriptor.Signal
}

type IOV_BackLeftCornerMotor3Descriptor struct {
	*descriptor.Message
	VESC_StatusWattHours3    *descriptor.Signal
	VESC_StatusWattHoursChg3 *descriptor.Signal
}

type IOV_BackLeftCornerMotor4Descriptor struct {
	*descriptor.Message
	VESC_StatusTempFET3   *descriptor.Signal
	VESC_StatusTempMotor3 *descriptor.Signal
	VESC_StatusCurrentIn3 *descriptor.Signal
	VESC_StatusPIDPos3    *descriptor.Signal
}

type IOV_BackRightCornerMotor1Descriptor struct {
	*descriptor.Message
	VESC_StatusERPM4      *descriptor.Signal
	VESC_StatusCurrent4   *descriptor.Signal
	VESC_StatusDutyCycle4 *descriptor.Signal
}

type IOV_BackRightCornerMotor2Descriptor struct {
	*descriptor.Message
	VESC_StatusAmpHours4    *descriptor.Signal
	VESC_StatusAmpHoursChg4 *descriptor.Signal
}

type IOV_BackRightCornerMotor3Descriptor struct {
	*descriptor.Message
	VESC_StatusWattHours4    *descriptor.Signal
	VESC_StatusWattHoursChg4 *descriptor.Signal
}

type IOV_BackRightCornerMotor4Descriptor struct {
	*descriptor.Message
	VESC_StatusTempFET4   *descriptor.Signal
	VESC_StatusTempMotor4 *descriptor.Signal
	VESC_StatusCurrentIn4 *descriptor.Signal
	VESC_StatusPIDPos4    *descriptor.Signal
}

type OBC_CommandDescriptor struct {
	*descriptor.Message
	OBC_MaxAllwChargVolt     *descriptor.Signal
	OBC_MaxAllowChargAmp     *descriptor.Signal
	OBC_ControlWorkEnable    *descriptor.Signal
	OBC_ControlOperatingMode *descriptor.Signal
}

type BMS_DALY_SoCStatusDescriptor struct {
	*descriptor.Message
	DALY_AccumulatedPressure   *descriptor.Signal
	DALY_CollectedTotalVoltage *descriptor.Signal
	DALY_BatteryCurrent        *descriptor.Signal
	DALY_BatterySoC            *descriptor.Signal
}

type BMS_DALY_RangeVoltageDescriptor struct {
	*descriptor.Message
	DALY_MaxMonomerVoltage    *descriptor.Signal
	DALY_MaxUnitVoltageCellNo *descriptor.Signal
	DALY_MinMonomerVoltage    *descriptor.Signal
	DALY_MinUnitVoltageCellNo *descriptor.Signal
}

type BMS_DALY_RangeTemperatureDescriptor struct {
	*descriptor.Message
	DALY_MaxMonomerTemp    *descriptor.Signal
	DALY_MaxUnitTempCellNo *descriptor.Signal
	DALY_MinMonomerTemp    *descriptor.Signal
	DALY_MinUnitTempCellNo *descriptor.Signal
}

type BMS_DALY_MOSStatusDescriptor struct {
	*descriptor.Message
	DALY_DischargeStatus  *descriptor.Signal
	DALY_ChargingMOSTube  *descriptor.Signal
	DALY_DischargeMOSTube *descriptor.Signal
	DALY_BMSLife          *descriptor.Signal
	DALY_ResidualCapacity *descriptor.Signal
}

type BMS_DALY_StatusInformationDescriptor struct {
	*descriptor.Message
	DALY_BatteryString       *descriptor.Signal
	DALY_TemperatureStatus   *descriptor.Signal
	DALY_ChagerStatus        *descriptor.Signal
	DALY_LoadStatus          *descriptor.Signal
	DALY_DigitalInput1State  *descriptor.Signal
	DALY_DigitalInput2State  *descriptor.Signal
	DALY_DigitalInput3State  *descriptor.Signal
	DALY_DigitalInput4State  *descriptor.Signal
	DALY_DigitalOutput1State *descriptor.Signal
	DALY_DigitalOutput2State *descriptor.Signal
	DALY_DigitalOutput3State *descriptor.Signal
	DALY_DigitalOutput4State *descriptor.Signal
	DALY_DischargeCycles     *descriptor.Signal
}

type BMS_DALY_MonoCellVoltageDescriptor struct {
	*descriptor.Message
	DALY_CellFrameNumber *descriptor.Signal
	DALY_MonomerVoltage1 *descriptor.Signal
	DALY_MonomerVoltage2 *descriptor.Signal
	DALY_MonomerVoltage3 *descriptor.Signal
}

type BMS_DALY_MonoCellTempDescriptor struct {
	*descriptor.Message
	DALY_TempFrameNumber *descriptor.Signal
	DALY_MonomerTemp1    *descriptor.Signal
	DALY_MonomerTemp2    *descriptor.Signal
	DALY_MonomerTemp3    *descriptor.Signal
	DALY_MonomerTemp4    *descriptor.Signal
	DALY_MonomerTemp5    *descriptor.Signal
	DALY_MonomerTemp6    *descriptor.Signal
	DALY_MonomerTemp7    *descriptor.Signal
}

type OBC_ResponseDescriptor struct {
	*descriptor.Message
	OBC_OutputVoltage            *descriptor.Signal
	OBC_OutputCurrent            *descriptor.Signal
	OBC_HardwareProtection       *descriptor.Signal
	OBC_TemperatureProtection    *descriptor.Signal
	OBC_InputVoltageStatus       *descriptor.Signal
	OBC_OutputUnderVoltage       *descriptor.Signal
	OBC_OutputOverVoltage        *descriptor.Signal
	OBC_OutputOvercurrent        *descriptor.Signal
	OBC_OutputShortCircuit       *descriptor.Signal
	OBC_CommunicationStatus      *descriptor.Signal
	OBC_WorkingStatus            *descriptor.Signal
	OBC_ComOfInitialization      *descriptor.Signal
	OBC_FanWorkingEnable         *descriptor.Signal
	OBC_CoolingPumpEnable        *descriptor.Signal
	OBC_CCSignalStatus           *descriptor.Signal
	OBC_CPSignalStatus           *descriptor.Signal
	OBC_SocketOvheatFault        *descriptor.Signal
	OBC_ElectronicLockState      *descriptor.Signal
	OBC_S2SwitchControlBitStatus *descriptor.Signal
	OBC_Temperature              *descriptor.Signal
}

// Database returns the createvan database descriptor.
func (md *MessagesDescriptor) Database() *descriptor.Database {
	return d
}

var nd = &NodesDescriptor{
	CReATE_ECU: d.Nodes[0],
}

var md = &MessagesDescriptor{
	IVI_DashboardStatus: &IVI_DashboardStatusDescriptor{
		Message: d.Messages[0],
	},
	IVI_CornerMotorStatus: &IVI_CornerMotorStatusDescriptor{
		Message: d.Messages[1],
	},
	IOV_FrontLeftCornerMotor1: &IOV_FrontLeftCornerMotor1Descriptor{
		Message:               d.Messages[2],
		VESC_StatusERPM1:      d.Messages[2].Signals[0],
		VESC_StatusCurrent1:   d.Messages[2].Signals[1],
		VESC_StatusDutyCycle1: d.Messages[2].Signals[2],
	},
	IOV_FrontLeftCornerMotor2: &IOV_FrontLeftCornerMotor2Descriptor{
		Message:                 d.Messages[3],
		VESC_StatusAmpHours1:    d.Messages[3].Signals[0],
		VESC_StatusAmpHoursChg1: d.Messages[3].Signals[1],
	},
	IOV_FrontLeftCornerMotor3: &IOV_FrontLeftCornerMotor3Descriptor{
		Message:                  d.Messages[4],
		VESC_StatusWattHours1:    d.Messages[4].Signals[0],
		VESC_StatusWattHoursChg1: d.Messages[4].Signals[1],
	},
	IOV_FrontLeftCornerMotor4: &IOV_FrontLeftCornerMotor4Descriptor{
		Message:               d.Messages[5],
		VESC_StatusTemFET1:    d.Messages[5].Signals[0],
		VESC_StatusTempMotor1: d.Messages[5].Signals[1],
		VESC_StatusCurrentIn1: d.Messages[5].Signals[2],
		VESC_StatusPIDPos1:    d.Messages[5].Signals[3],
	},
	IOV_FrontRightCornerMotor1: &IOV_FrontRightCornerMotor1Descriptor{
		Message:               d.Messages[6],
		VESC_StatusERPM2:      d.Messages[6].Signals[0],
		VESC_StatusCurrent2:   d.Messages[6].Signals[1],
		VESC_StatusDutyCycle2: d.Messages[6].Signals[2],
	},
	IOV_FrontRightCornerMotor2: &IOV_FrontRightCornerMotor2Descriptor{
		Message:                 d.Messages[7],
		VESC_StatusAmpHours2:    d.Messages[7].Signals[0],
		VESC_StatusAmpHoursChg2: d.Messages[7].Signals[1],
	},
	IOV_FrontRightCornerMotor3: &IOV_FrontRightCornerMotor3Descriptor{
		Message:                  d.Messages[8],
		VESC_StatusWattHours2:    d.Messages[8].Signals[0],
		VESC_StatusWattHoursChg2: d.Messages[8].Signals[1],
	},
	IOV_FrontRightCornerMotor4: &IOV_FrontRightCornerMotor4Descriptor{
		Message:               d.Messages[9],
		VESC_StatusTempFET2:   d.Messages[9].Signals[0],
		VESC_StatusTempMotor2: d.Messages[9].Signals[1],
		VESC_StatusCurrentIn2: d.Messages[9].Signals[2],
		VESC_StatusPIDPos2:    d.Messages[9].Signals[3],
	},
	IOV_BackLeftCornerMotor1: &IOV_BackLeftCornerMotor1Descriptor{
		Message:               d.Messages[10],
		VESC_StatusERPM3:      d.Messages[10].Signals[0],
		VESC_StatusCurrent3:   d.Messages[10].Signals[1],
		VESC_StatusDutyCycle3: d.Messages[10].Signals[2],
	},
	IOV_BackLeftCornerMotor2: &IOV_BackLeftCornerMotor2Descriptor{
		Message:                 d.Messages[11],
		VESC_StatusAmpHours3:    d.Messages[11].Signals[0],
		VESC_StatusAmpHoursChg3: d.Messages[11].Signals[1],
	},
	IOV_BackLeftCornerMotor3: &IOV_BackLeftCornerMotor3Descriptor{
		Message:                  d.Messages[12],
		VESC_StatusWattHours3:    d.Messages[12].Signals[0],
		VESC_StatusWattHoursChg3: d.Messages[12].Signals[1],
	},
	IOV_BackLeftCornerMotor4: &IOV_BackLeftCornerMotor4Descriptor{
		Message:               d.Messages[13],
		VESC_StatusTempFET3:   d.Messages[13].Signals[0],
		VESC_StatusTempMotor3: d.Messages[13].Signals[1],
		VESC_StatusCurrentIn3: d.Messages[13].Signals[2],
		VESC_StatusPIDPos3:    d.Messages[13].Signals[3],
	},
	IOV_BackRightCornerMotor1: &IOV_BackRightCornerMotor1Descriptor{
		Message:               d.Messages[14],
		VESC_StatusERPM4:      d.Messages[14].Signals[0],
		VESC_StatusCurrent4:   d.Messages[14].Signals[1],
		VESC_StatusDutyCycle4: d.Messages[14].Signals[2],
	},
	IOV_BackRightCornerMotor2: &IOV_BackRightCornerMotor2Descriptor{
		Message:                 d.Messages[15],
		VESC_StatusAmpHours4:    d.Messages[15].Signals[0],
		VESC_StatusAmpHoursChg4: d.Messages[15].Signals[1],
	},
	IOV_BackRightCornerMotor3: &IOV_BackRightCornerMotor3Descriptor{
		Message:                  d.Messages[16],
		VESC_StatusWattHours4:    d.Messages[16].Signals[0],
		VESC_StatusWattHoursChg4: d.Messages[16].Signals[1],
	},
	IOV_BackRightCornerMotor4: &IOV_BackRightCornerMotor4Descriptor{
		Message:               d.Messages[17],
		VESC_StatusTempFET4:   d.Messages[17].Signals[0],
		VESC_StatusTempMotor4: d.Messages[17].Signals[1],
		VESC_StatusCurrentIn4: d.Messages[17].Signals[2],
		VESC_StatusPIDPos4:    d.Messages[17].Signals[3],
	},
	OBC_Command: &OBC_CommandDescriptor{
		Message:                  d.Messages[18],
		OBC_MaxAllwChargVolt:     d.Messages[18].Signals[0],
		OBC_MaxAllowChargAmp:     d.Messages[18].Signals[1],
		OBC_ControlWorkEnable:    d.Messages[18].Signals[2],
		OBC_ControlOperatingMode: d.Messages[18].Signals[3],
	},
	BMS_DALY_SoCStatus: &BMS_DALY_SoCStatusDescriptor{
		Message:                    d.Messages[19],
		DALY_AccumulatedPressure:   d.Messages[19].Signals[0],
		DALY_CollectedTotalVoltage: d.Messages[19].Signals[1],
		DALY_BatteryCurrent:        d.Messages[19].Signals[2],
		DALY_BatterySoC:            d.Messages[19].Signals[3],
	},
	BMS_DALY_RangeVoltage: &BMS_DALY_RangeVoltageDescriptor{
		Message:                   d.Messages[20],
		DALY_MaxMonomerVoltage:    d.Messages[20].Signals[0],
		DALY_MaxUnitVoltageCellNo: d.Messages[20].Signals[1],
		DALY_MinMonomerVoltage:    d.Messages[20].Signals[2],
		DALY_MinUnitVoltageCellNo: d.Messages[20].Signals[3],
	},
	BMS_DALY_RangeTemperature: &BMS_DALY_RangeTemperatureDescriptor{
		Message:                d.Messages[21],
		DALY_MaxMonomerTemp:    d.Messages[21].Signals[0],
		DALY_MaxUnitTempCellNo: d.Messages[21].Signals[1],
		DALY_MinMonomerTemp:    d.Messages[21].Signals[2],
		DALY_MinUnitTempCellNo: d.Messages[21].Signals[3],
	},
	BMS_DALY_MOSStatus: &BMS_DALY_MOSStatusDescriptor{
		Message:               d.Messages[22],
		DALY_DischargeStatus:  d.Messages[22].Signals[0],
		DALY_ChargingMOSTube:  d.Messages[22].Signals[1],
		DALY_DischargeMOSTube: d.Messages[22].Signals[2],
		DALY_BMSLife:          d.Messages[22].Signals[3],
		DALY_ResidualCapacity: d.Messages[22].Signals[4],
	},
	BMS_DALY_StatusInformation: &BMS_DALY_StatusInformationDescriptor{
		Message:                  d.Messages[23],
		DALY_BatteryString:       d.Messages[23].Signals[0],
		DALY_TemperatureStatus:   d.Messages[23].Signals[1],
		DALY_ChagerStatus:        d.Messages[23].Signals[2],
		DALY_LoadStatus:          d.Messages[23].Signals[3],
		DALY_DigitalInput1State:  d.Messages[23].Signals[4],
		DALY_DigitalInput2State:  d.Messages[23].Signals[5],
		DALY_DigitalInput3State:  d.Messages[23].Signals[6],
		DALY_DigitalInput4State:  d.Messages[23].Signals[7],
		DALY_DigitalOutput1State: d.Messages[23].Signals[8],
		DALY_DigitalOutput2State: d.Messages[23].Signals[9],
		DALY_DigitalOutput3State: d.Messages[23].Signals[10],
		DALY_DigitalOutput4State: d.Messages[23].Signals[11],
		DALY_DischargeCycles:     d.Messages[23].Signals[12],
	},
	BMS_DALY_MonoCellVoltage: &BMS_DALY_MonoCellVoltageDescriptor{
		Message:              d.Messages[24],
		DALY_CellFrameNumber: d.Messages[24].Signals[0],
		DALY_MonomerVoltage1: d.Messages[24].Signals[1],
		DALY_MonomerVoltage2: d.Messages[24].Signals[2],
		DALY_MonomerVoltage3: d.Messages[24].Signals[3],
	},
	BMS_DALY_MonoCellTemp: &BMS_DALY_MonoCellTempDescriptor{
		Message:              d.Messages[25],
		DALY_TempFrameNumber: d.Messages[25].Signals[0],
		DALY_MonomerTemp1:    d.Messages[25].Signals[1],
		DALY_MonomerTemp2:    d.Messages[25].Signals[2],
		DALY_MonomerTemp3:    d.Messages[25].Signals[3],
		DALY_MonomerTemp4:    d.Messages[25].Signals[4],
		DALY_MonomerTemp5:    d.Messages[25].Signals[5],
		DALY_MonomerTemp6:    d.Messages[25].Signals[6],
		DALY_MonomerTemp7:    d.Messages[25].Signals[7],
	},
	OBC_Response: &OBC_ResponseDescriptor{
		Message:                      d.Messages[26],
		OBC_OutputVoltage:            d.Messages[26].Signals[0],
		OBC_OutputCurrent:            d.Messages[26].Signals[1],
		OBC_HardwareProtection:       d.Messages[26].Signals[2],
		OBC_TemperatureProtection:    d.Messages[26].Signals[3],
		OBC_InputVoltageStatus:       d.Messages[26].Signals[4],
		OBC_OutputUnderVoltage:       d.Messages[26].Signals[5],
		OBC_OutputOverVoltage:        d.Messages[26].Signals[6],
		OBC_OutputOvercurrent:        d.Messages[26].Signals[7],
		OBC_OutputShortCircuit:       d.Messages[26].Signals[8],
		OBC_CommunicationStatus:      d.Messages[26].Signals[9],
		OBC_WorkingStatus:            d.Messages[26].Signals[10],
		OBC_ComOfInitialization:      d.Messages[26].Signals[11],
		OBC_FanWorkingEnable:         d.Messages[26].Signals[12],
		OBC_CoolingPumpEnable:        d.Messages[26].Signals[13],
		OBC_CCSignalStatus:           d.Messages[26].Signals[14],
		OBC_CPSignalStatus:           d.Messages[26].Signals[15],
		OBC_SocketOvheatFault:        d.Messages[26].Signals[16],
		OBC_ElectronicLockState:      d.Messages[26].Signals[17],
		OBC_S2SwitchControlBitStatus: d.Messages[26].Signals[18],
		OBC_Temperature:              d.Messages[26].Signals[19],
	},
}

var d = (*descriptor.Database)(&descriptor.Database{
	SourceFile: (string)("dummy/create/createvan.dbc"),
	Version:    (string)(""),
	Messages: ([]*descriptor.Message)([]*descriptor.Message{
		(*descriptor.Message)(&descriptor.Message{
			Name:        (string)("IVI_DashboardStatus"),
			ID:          (uint32)(217056510),
			IsExtended:  (bool)(true),
			Length:      (uint8)(8),
			SendType:    (descriptor.SendType)(0),
			Description: (string)(""),
			Signals:     ([]*descriptor.Signal)(nil),
			SenderNode:  (string)("Vector__XXX"),
			CycleTime:   (time.Duration)(0),
			DelayTime:   (time.Duration)(0),
		}),
		(*descriptor.Message)(&descriptor.Message{
			Name:        (string)("IVI_CornerMotorStatus"),
			ID:          (uint32)(217056766),
			IsExtended:  (bool)(true),
			Length:      (uint8)(8),
			SendType:    (descriptor.SendType)(0),
			Description: (string)(""),
			Signals:     ([]*descriptor.Signal)(nil),
			SenderNode:  (string)("Vector__XXX"),
			CycleTime:   (time.Duration)(0),
			DelayTime:   (time.Duration)(0),
		}),
		(*descriptor.Message)(&descriptor.Message{
			Name:        (string)("IOV_FrontLeftCornerMotor1"),
			ID:          (uint32)(217063422),
			IsExtended:  (bool)(true),
			Length:      (uint8)(8),
			SendType:    (descriptor.SendType)(0),
			Description: (string)(""),
			Signals: ([]*descriptor.Signal)([]*descriptor.Signal{
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("VESC_StatusERPM1"),
					Start:             (uint8)(7),
					Length:            (uint8)(32),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(-2.147483648e+09),
					Max:               (float64)(2.147483647e+09),
					Unit:              (string)("RPM"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("VESC_StatusCurrent1"),
					Start:             (uint8)(39),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(-32768),
					Max:               (float64)(32767),
					Unit:              (string)("A"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("VESC_StatusDutyCycle1"),
					Start:             (uint8)(55),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(-32768),
					Max:               (float64)(32767),
					Unit:              (string)("%"),
					Description:       (string)(""),
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
			Name:        (string)("IOV_FrontLeftCornerMotor2"),
			ID:          (uint32)(217067518),
			IsExtended:  (bool)(true),
			Length:      (uint8)(8),
			SendType:    (descriptor.SendType)(0),
			Description: (string)(""),
			Signals: ([]*descriptor.Signal)([]*descriptor.Signal{
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("VESC_StatusAmpHours1"),
					Start:             (uint8)(7),
					Length:            (uint8)(32),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(-2.147483648e+09),
					Max:               (float64)(2.147483647e+09),
					Unit:              (string)("Ah"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("VESC_StatusAmpHoursChg1"),
					Start:             (uint8)(39),
					Length:            (uint8)(32),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(-2.147483648e+09),
					Max:               (float64)(2.147483647e+09),
					Unit:              (string)("Ah"),
					Description:       (string)(""),
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
			Name:        (string)("IOV_FrontLeftCornerMotor3"),
			ID:          (uint32)(217071614),
			IsExtended:  (bool)(true),
			Length:      (uint8)(8),
			SendType:    (descriptor.SendType)(0),
			Description: (string)(""),
			Signals: ([]*descriptor.Signal)([]*descriptor.Signal{
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("VESC_StatusWattHours1"),
					Start:             (uint8)(7),
					Length:            (uint8)(32),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(-2.147483648e+09),
					Max:               (float64)(2.147483647e+09),
					Unit:              (string)("Wh"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("VESC_StatusWattHoursChg1"),
					Start:             (uint8)(39),
					Length:            (uint8)(32),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(-2.147483648e+09),
					Max:               (float64)(2.147483647e+09),
					Unit:              (string)("Wh"),
					Description:       (string)(""),
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
			Name:        (string)("IOV_FrontLeftCornerMotor4"),
			ID:          (uint32)(217075710),
			IsExtended:  (bool)(true),
			Length:      (uint8)(8),
			SendType:    (descriptor.SendType)(0),
			Description: (string)(""),
			Signals: ([]*descriptor.Signal)([]*descriptor.Signal{
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("VESC_StatusTemFET1"),
					Start:             (uint8)(7),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(-32768),
					Max:               (float64)(32767),
					Unit:              (string)(""),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("VESC_StatusTempMotor1"),
					Start:             (uint8)(23),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(-32768),
					Max:               (float64)(32767),
					Unit:              (string)(""),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("VESC_StatusCurrentIn1"),
					Start:             (uint8)(39),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(-32768),
					Max:               (float64)(32767),
					Unit:              (string)(""),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("VESC_StatusPIDPos1"),
					Start:             (uint8)(55),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(-32768),
					Max:               (float64)(32767),
					Unit:              (string)(""),
					Description:       (string)(""),
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
			Name:        (string)("IOV_FrontRightCornerMotor1"),
			ID:          (uint32)(217079806),
			IsExtended:  (bool)(true),
			Length:      (uint8)(8),
			SendType:    (descriptor.SendType)(0),
			Description: (string)(""),
			Signals: ([]*descriptor.Signal)([]*descriptor.Signal{
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("VESC_StatusERPM2"),
					Start:             (uint8)(7),
					Length:            (uint8)(32),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(-2.147483648e+09),
					Max:               (float64)(2.147483647e+09),
					Unit:              (string)("RPM"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("VESC_StatusCurrent2"),
					Start:             (uint8)(39),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(-32768),
					Max:               (float64)(32767),
					Unit:              (string)("A"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("VESC_StatusDutyCycle2"),
					Start:             (uint8)(55),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(-32768),
					Max:               (float64)(32767),
					Unit:              (string)("%"),
					Description:       (string)(""),
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
			Name:        (string)("IOV_FrontRightCornerMotor2"),
			ID:          (uint32)(217083902),
			IsExtended:  (bool)(true),
			Length:      (uint8)(8),
			SendType:    (descriptor.SendType)(0),
			Description: (string)(""),
			Signals: ([]*descriptor.Signal)([]*descriptor.Signal{
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("VESC_StatusAmpHours2"),
					Start:             (uint8)(7),
					Length:            (uint8)(32),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(-2.147483648e+09),
					Max:               (float64)(2.147483647e+09),
					Unit:              (string)("Ah"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("VESC_StatusAmpHoursChg2"),
					Start:             (uint8)(39),
					Length:            (uint8)(32),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(-2.147483648e+09),
					Max:               (float64)(2.147483647e+09),
					Unit:              (string)("Ah"),
					Description:       (string)(""),
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
			Name:        (string)("IOV_FrontRightCornerMotor3"),
			ID:          (uint32)(217087998),
			IsExtended:  (bool)(true),
			Length:      (uint8)(8),
			SendType:    (descriptor.SendType)(0),
			Description: (string)(""),
			Signals: ([]*descriptor.Signal)([]*descriptor.Signal{
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("VESC_StatusWattHours2"),
					Start:             (uint8)(7),
					Length:            (uint8)(32),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(-2.147483648e+09),
					Max:               (float64)(2.147483647e+09),
					Unit:              (string)("Wh"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("VESC_StatusWattHoursChg2"),
					Start:             (uint8)(39),
					Length:            (uint8)(32),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(-2.147483648e+09),
					Max:               (float64)(2.147483647e+09),
					Unit:              (string)("Wh"),
					Description:       (string)(""),
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
			Name:        (string)("IOV_FrontRightCornerMotor4"),
			ID:          (uint32)(217092094),
			IsExtended:  (bool)(true),
			Length:      (uint8)(8),
			SendType:    (descriptor.SendType)(0),
			Description: (string)(""),
			Signals: ([]*descriptor.Signal)([]*descriptor.Signal{
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("VESC_StatusTempFET2"),
					Start:             (uint8)(7),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(-32768),
					Max:               (float64)(32767),
					Unit:              (string)(""),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("VESC_StatusTempMotor2"),
					Start:             (uint8)(23),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(-32768),
					Max:               (float64)(32767),
					Unit:              (string)(""),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("VESC_StatusCurrentIn2"),
					Start:             (uint8)(39),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(-32768),
					Max:               (float64)(32767),
					Unit:              (string)(""),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("VESC_StatusPIDPos2"),
					Start:             (uint8)(55),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(-32768),
					Max:               (float64)(32767),
					Unit:              (string)(""),
					Description:       (string)(""),
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
			Name:        (string)("IOV_BackLeftCornerMotor1"),
			ID:          (uint32)(217096190),
			IsExtended:  (bool)(true),
			Length:      (uint8)(8),
			SendType:    (descriptor.SendType)(0),
			Description: (string)(""),
			Signals: ([]*descriptor.Signal)([]*descriptor.Signal{
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("VESC_StatusERPM3"),
					Start:             (uint8)(7),
					Length:            (uint8)(32),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(-2.147483648e+09),
					Max:               (float64)(2.147483647e+09),
					Unit:              (string)("RPM"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("VESC_StatusCurrent3"),
					Start:             (uint8)(39),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(-32768),
					Max:               (float64)(32767),
					Unit:              (string)("A"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("VESC_StatusDutyCycle3"),
					Start:             (uint8)(55),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(-32768),
					Max:               (float64)(32767),
					Unit:              (string)("%"),
					Description:       (string)(""),
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
			Name:        (string)("IOV_BackLeftCornerMotor2"),
			ID:          (uint32)(217100286),
			IsExtended:  (bool)(true),
			Length:      (uint8)(8),
			SendType:    (descriptor.SendType)(0),
			Description: (string)(""),
			Signals: ([]*descriptor.Signal)([]*descriptor.Signal{
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("VESC_StatusAmpHours3"),
					Start:             (uint8)(7),
					Length:            (uint8)(32),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(-2.147483648e+09),
					Max:               (float64)(2.147483647e+09),
					Unit:              (string)("Ah"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("VESC_StatusAmpHoursChg3"),
					Start:             (uint8)(39),
					Length:            (uint8)(32),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(-2.147483648e+09),
					Max:               (float64)(2.147483647e+09),
					Unit:              (string)("Ah"),
					Description:       (string)(""),
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
			Name:        (string)("IOV_BackLeftCornerMotor3"),
			ID:          (uint32)(217104382),
			IsExtended:  (bool)(true),
			Length:      (uint8)(8),
			SendType:    (descriptor.SendType)(0),
			Description: (string)(""),
			Signals: ([]*descriptor.Signal)([]*descriptor.Signal{
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("VESC_StatusWattHours3"),
					Start:             (uint8)(7),
					Length:            (uint8)(32),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(-2.147483648e+09),
					Max:               (float64)(2.147483647e+09),
					Unit:              (string)("Wh"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("VESC_StatusWattHoursChg3"),
					Start:             (uint8)(39),
					Length:            (uint8)(32),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(-2.147483648e+09),
					Max:               (float64)(2.147483647e+09),
					Unit:              (string)("Wh"),
					Description:       (string)(""),
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
			Name:        (string)("IOV_BackLeftCornerMotor4"),
			ID:          (uint32)(217108478),
			IsExtended:  (bool)(true),
			Length:      (uint8)(8),
			SendType:    (descriptor.SendType)(0),
			Description: (string)(""),
			Signals: ([]*descriptor.Signal)([]*descriptor.Signal{
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("VESC_StatusTempFET3"),
					Start:             (uint8)(7),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(-32768),
					Max:               (float64)(32767),
					Unit:              (string)(""),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("VESC_StatusTempMotor3"),
					Start:             (uint8)(23),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(-32768),
					Max:               (float64)(32767),
					Unit:              (string)(""),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("VESC_StatusCurrentIn3"),
					Start:             (uint8)(39),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)(""),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("VESC_StatusPIDPos3"),
					Start:             (uint8)(55),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(-32768),
					Max:               (float64)(32767),
					Unit:              (string)(""),
					Description:       (string)(""),
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
			Name:        (string)("IOV_BackRightCornerMotor1"),
			ID:          (uint32)(217112574),
			IsExtended:  (bool)(true),
			Length:      (uint8)(8),
			SendType:    (descriptor.SendType)(0),
			Description: (string)(""),
			Signals: ([]*descriptor.Signal)([]*descriptor.Signal{
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("VESC_StatusERPM4"),
					Start:             (uint8)(7),
					Length:            (uint8)(32),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(-2.147483648e+09),
					Max:               (float64)(2.147483647e+09),
					Unit:              (string)("RPM"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("VESC_StatusCurrent4"),
					Start:             (uint8)(39),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(-32768),
					Max:               (float64)(32767),
					Unit:              (string)("A"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("VESC_StatusDutyCycle4"),
					Start:             (uint8)(55),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(-32768),
					Max:               (float64)(32767),
					Unit:              (string)("%"),
					Description:       (string)(""),
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
			Name:        (string)("IOV_BackRightCornerMotor2"),
			ID:          (uint32)(217116670),
			IsExtended:  (bool)(true),
			Length:      (uint8)(8),
			SendType:    (descriptor.SendType)(0),
			Description: (string)(""),
			Signals: ([]*descriptor.Signal)([]*descriptor.Signal{
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("VESC_StatusAmpHours4"),
					Start:             (uint8)(7),
					Length:            (uint8)(32),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(-2.147483648e+09),
					Max:               (float64)(2.147483647e+09),
					Unit:              (string)("Ah"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("VESC_StatusAmpHoursChg4"),
					Start:             (uint8)(39),
					Length:            (uint8)(32),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(-2.147483648e+09),
					Max:               (float64)(2.147483647e+09),
					Unit:              (string)("Ah"),
					Description:       (string)(""),
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
			Name:        (string)("IOV_BackRightCornerMotor3"),
			ID:          (uint32)(217120766),
			IsExtended:  (bool)(true),
			Length:      (uint8)(8),
			SendType:    (descriptor.SendType)(0),
			Description: (string)(""),
			Signals: ([]*descriptor.Signal)([]*descriptor.Signal{
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("VESC_StatusWattHours4"),
					Start:             (uint8)(7),
					Length:            (uint8)(32),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("Wh"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("VESC_StatusWattHoursChg4"),
					Start:             (uint8)(39),
					Length:            (uint8)(32),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(-2.147483648e+09),
					Max:               (float64)(2.147483647e+09),
					Unit:              (string)("Wh"),
					Description:       (string)(""),
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
			Name:        (string)("IOV_BackRightCornerMotor4"),
			ID:          (uint32)(217186302),
			IsExtended:  (bool)(true),
			Length:      (uint8)(8),
			SendType:    (descriptor.SendType)(0),
			Description: (string)(""),
			Signals: ([]*descriptor.Signal)([]*descriptor.Signal{
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("VESC_StatusTempFET4"),
					Start:             (uint8)(7),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(-32768),
					Max:               (float64)(32767),
					Unit:              (string)(""),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("VESC_StatusTempMotor4"),
					Start:             (uint8)(23),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(-32768),
					Max:               (float64)(32767),
					Unit:              (string)(""),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("VESC_StatusCurrentIn4"),
					Start:             (uint8)(39),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(-32768),
					Max:               (float64)(32767),
					Unit:              (string)(""),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("VESC_StatusPIDPos4"),
					Start:             (uint8)(56),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(-32768),
					Max:               (float64)(32767),
					Unit:              (string)(""),
					Description:       (string)(""),
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
			Name:        (string)("OBC_Command"),
			ID:          (uint32)(403105268),
			IsExtended:  (bool)(true),
			Length:      (uint8)(8),
			SendType:    (descriptor.SendType)(0),
			Description: (string)(""),
			Signals: ([]*descriptor.Signal)([]*descriptor.Signal{
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("OBC_MaxAllwChargVolt"),
					Start:             (uint8)(7),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(6553.5),
					Unit:              (string)("V"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("OBC_MaxAllowChargAmp"),
					Start:             (uint8)(23),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(6553.5),
					Unit:              (string)("A"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:             (string)("OBC_ControlWorkEnable"),
					Start:            (uint8)(39),
					Length:           (uint8)(8),
					IsBigEndian:      (bool)(true),
					IsSigned:         (bool)(false),
					IsFloat:          (bool)(false),
					IsMultiplexer:    (bool)(false),
					IsMultiplexed:    (bool)(false),
					MultiplexerValue: (uint)(0),
					Offset:           (float64)(0),
					Scale:            (float64)(1),
					Min:              (float64)(0),
					Max:              (float64)(255),
					Unit:             (string)(""),
					Description:      (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)([]*descriptor.ValueDescription{
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(0),
							Description: (string)("ChargerIsStartingtoCharge"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(1),
							Description: (string)("ChargerCloseTheOutput"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(2),
							Description: (string)("ChargeEnd"),
						}),
					}),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:             (string)("OBC_ControlOperatingMode"),
					Start:            (uint8)(47),
					Length:           (uint8)(8),
					IsBigEndian:      (bool)(true),
					IsSigned:         (bool)(false),
					IsFloat:          (bool)(false),
					IsMultiplexer:    (bool)(false),
					IsMultiplexed:    (bool)(false),
					MultiplexerValue: (uint)(0),
					Offset:           (float64)(0),
					Scale:            (float64)(1),
					Min:              (float64)(0),
					Max:              (float64)(255),
					Unit:             (string)(""),
					Description:      (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)([]*descriptor.ValueDescription{
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(0),
							Description: (string)("ChargingMode"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(1),
							Description: (string)("HeatingModel"),
						}),
					}),
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
			Name:        (string)("BMS_DALY_SoCStatus"),
			ID:          (uint32)(412106753),
			IsExtended:  (bool)(true),
			Length:      (uint8)(8),
			SendType:    (descriptor.SendType)(0),
			Description: (string)(""),
			Signals: ([]*descriptor.Signal)([]*descriptor.Signal{
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("DALY_AccumulatedPressure"),
					Start:             (uint8)(7),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(6553.5),
					Unit:              (string)("V"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("DALY_CollectedTotalVoltage"),
					Start:             (uint8)(23),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(6553.5),
					Unit:              (string)("V"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("DALY_BatteryCurrent"),
					Start:             (uint8)(39),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(30000),
					Scale:             (float64)(0.1),
					Min:               (float64)(30000),
					Max:               (float64)(36553.5),
					Unit:              (string)("A"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("DALY_BatterySoC"),
					Start:             (uint8)(55),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(6553.5),
					Unit:              (string)("%"),
					Description:       (string)(""),
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
			Name:        (string)("BMS_DALY_RangeVoltage"),
			ID:          (uint32)(412172289),
			IsExtended:  (bool)(true),
			Length:      (uint8)(8),
			SendType:    (descriptor.SendType)(0),
			Description: (string)(""),
			Signals: ([]*descriptor.Signal)([]*descriptor.Signal{
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("DALY_MaxMonomerVoltage"),
					Start:             (uint8)(7),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(65535),
					Unit:              (string)("mV"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("DALY_MaxUnitVoltageCellNo"),
					Start:             (uint8)(23),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(255),
					Unit:              (string)(""),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("DALY_MinMonomerVoltage"),
					Start:             (uint8)(31),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(65535),
					Unit:              (string)("mV"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("DALY_MinUnitVoltageCellNo"),
					Start:             (uint8)(47),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(255),
					Unit:              (string)(""),
					Description:       (string)(""),
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
			Name:        (string)("BMS_DALY_RangeTemperature"),
			ID:          (uint32)(412237825),
			IsExtended:  (bool)(true),
			Length:      (uint8)(8),
			SendType:    (descriptor.SendType)(0),
			Description: (string)(""),
			Signals: ([]*descriptor.Signal)([]*descriptor.Signal{
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("DALY_MaxMonomerTemp"),
					Start:             (uint8)(7),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(40),
					Scale:             (float64)(1),
					Min:               (float64)(40),
					Max:               (float64)(295),
					Unit:              (string)("degC"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("DALY_MaxUnitTempCellNo"),
					Start:             (uint8)(15),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(255),
					Unit:              (string)(""),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("DALY_MinMonomerTemp"),
					Start:             (uint8)(23),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(40),
					Scale:             (float64)(1),
					Min:               (float64)(40),
					Max:               (float64)(295),
					Unit:              (string)("degC"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("DALY_MinUnitTempCellNo"),
					Start:             (uint8)(31),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(255),
					Unit:              (string)(""),
					Description:       (string)(""),
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
			Name:        (string)("BMS_DALY_MOSStatus"),
			ID:          (uint32)(412303361),
			IsExtended:  (bool)(true),
			Length:      (uint8)(8),
			SendType:    (descriptor.SendType)(0),
			Description: (string)(""),
			Signals: ([]*descriptor.Signal)([]*descriptor.Signal{
				(*descriptor.Signal)(&descriptor.Signal{
					Name:             (string)("DALY_DischargeStatus"),
					Start:            (uint8)(7),
					Length:           (uint8)(8),
					IsBigEndian:      (bool)(true),
					IsSigned:         (bool)(false),
					IsFloat:          (bool)(false),
					IsMultiplexer:    (bool)(false),
					IsMultiplexed:    (bool)(false),
					MultiplexerValue: (uint)(0),
					Offset:           (float64)(0),
					Scale:            (float64)(1),
					Min:              (float64)(0),
					Max:              (float64)(255),
					Unit:             (string)(""),
					Description:      (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)([]*descriptor.ValueDescription{
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(0),
							Description: (string)("Stationary"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(1),
							Description: (string)("Charged"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(2),
							Description: (string)("Discharged"),
						}),
					}),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("DALY_ChargingMOSTube"),
					Start:             (uint8)(15),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(255),
					Unit:              (string)(""),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("DALY_DischargeMOSTube"),
					Start:             (uint8)(23),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(255),
					Unit:              (string)(""),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("DALY_BMSLife"),
					Start:             (uint8)(31),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(255),
					Unit:              (string)("cycles"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("DALY_ResidualCapacity"),
					Start:             (uint8)(39),
					Length:            (uint8)(32),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(4.294967295e+09),
					Unit:              (string)("mAH"),
					Description:       (string)(""),
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
			Name:        (string)("BMS_DALY_StatusInformation"),
			ID:          (uint32)(412368897),
			IsExtended:  (bool)(true),
			Length:      (uint8)(8),
			SendType:    (descriptor.SendType)(0),
			Description: (string)(""),
			Signals: ([]*descriptor.Signal)([]*descriptor.Signal{
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("DALY_BatteryString"),
					Start:             (uint8)(7),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(255),
					Unit:              (string)(""),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("DALY_TemperatureStatus"),
					Start:             (uint8)(15),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(255),
					Unit:              (string)(""),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:             (string)("DALY_ChagerStatus"),
					Start:            (uint8)(23),
					Length:           (uint8)(8),
					IsBigEndian:      (bool)(true),
					IsSigned:         (bool)(false),
					IsFloat:          (bool)(false),
					IsMultiplexer:    (bool)(false),
					IsMultiplexed:    (bool)(false),
					MultiplexerValue: (uint)(0),
					Offset:           (float64)(0),
					Scale:            (float64)(1),
					Min:              (float64)(0),
					Max:              (float64)(255),
					Unit:             (string)(""),
					Description:      (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)([]*descriptor.ValueDescription{
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(0),
							Description: (string)("Disconnected"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(1),
							Description: (string)("Connected"),
						}),
					}),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:             (string)("DALY_LoadStatus"),
					Start:            (uint8)(31),
					Length:           (uint8)(8),
					IsBigEndian:      (bool)(true),
					IsSigned:         (bool)(false),
					IsFloat:          (bool)(false),
					IsMultiplexer:    (bool)(false),
					IsMultiplexed:    (bool)(false),
					MultiplexerValue: (uint)(0),
					Offset:           (float64)(0),
					Scale:            (float64)(1),
					Min:              (float64)(0),
					Max:              (float64)(255),
					Unit:             (string)(""),
					Description:      (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)([]*descriptor.ValueDescription{
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(0),
							Description: (string)("Disconnected"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(1),
							Description: (string)("Access"),
						}),
					}),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("DALY_DigitalInput1State"),
					Start:             (uint8)(32),
					Length:            (uint8)(1),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(1),
					Unit:              (string)(""),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("DALY_DigitalInput2State"),
					Start:             (uint8)(33),
					Length:            (uint8)(1),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(1),
					Unit:              (string)(""),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("DALY_DigitalInput3State"),
					Start:             (uint8)(34),
					Length:            (uint8)(1),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(1),
					Unit:              (string)(""),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("DALY_DigitalInput4State"),
					Start:             (uint8)(35),
					Length:            (uint8)(1),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(1),
					Unit:              (string)(""),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("DALY_DigitalOutput1State"),
					Start:             (uint8)(36),
					Length:            (uint8)(1),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(1),
					Unit:              (string)(""),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("DALY_DigitalOutput2State"),
					Start:             (uint8)(37),
					Length:            (uint8)(1),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(1),
					Unit:              (string)(""),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("DALY_DigitalOutput3State"),
					Start:             (uint8)(38),
					Length:            (uint8)(1),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(1),
					Unit:              (string)(""),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("DALY_DigitalOutput4State"),
					Start:             (uint8)(39),
					Length:            (uint8)(1),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(1),
					Unit:              (string)(""),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("DALY_DischargeCycles"),
					Start:             (uint8)(47),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(65535),
					Unit:              (string)(""),
					Description:       (string)(""),
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
			Name:        (string)("BMS_DALY_MonoCellVoltage"),
			ID:          (uint32)(412434433),
			IsExtended:  (bool)(true),
			Length:      (uint8)(8),
			SendType:    (descriptor.SendType)(0),
			Description: (string)(""),
			Signals: ([]*descriptor.Signal)([]*descriptor.Signal{
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("DALY_CellFrameNumber"),
					Start:             (uint8)(7),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(255),
					Unit:              (string)(""),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("DALY_MonomerVoltage1"),
					Start:             (uint8)(15),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(65535),
					Unit:              (string)("mV"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("DALY_MonomerVoltage2"),
					Start:             (uint8)(31),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(65535),
					Unit:              (string)("mV"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("DALY_MonomerVoltage3"),
					Start:             (uint8)(47),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(65535),
					Unit:              (string)("mV"),
					Description:       (string)(""),
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
			Name:        (string)("BMS_DALY_MonoCellTemp"),
			ID:          (uint32)(412499969),
			IsExtended:  (bool)(true),
			Length:      (uint8)(8),
			SendType:    (descriptor.SendType)(0),
			Description: (string)(""),
			Signals: ([]*descriptor.Signal)([]*descriptor.Signal{
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("DALY_TempFrameNumber"),
					Start:             (uint8)(7),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(255),
					Unit:              (string)(""),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("DALY_MonomerTemp1"),
					Start:             (uint8)(15),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(40),
					Scale:             (float64)(1),
					Min:               (float64)(40),
					Max:               (float64)(295),
					Unit:              (string)("degC"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("DALY_MonomerTemp2"),
					Start:             (uint8)(23),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(40),
					Scale:             (float64)(1),
					Min:               (float64)(40),
					Max:               (float64)(295),
					Unit:              (string)("degC"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("DALY_MonomerTemp3"),
					Start:             (uint8)(31),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(40),
					Scale:             (float64)(1),
					Min:               (float64)(40),
					Max:               (float64)(295),
					Unit:              (string)("degC"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("DALY_MonomerTemp4"),
					Start:             (uint8)(39),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(40),
					Scale:             (float64)(1),
					Min:               (float64)(40),
					Max:               (float64)(295),
					Unit:              (string)("degC"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("DALY_MonomerTemp5"),
					Start:             (uint8)(47),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(40),
					Scale:             (float64)(1),
					Min:               (float64)(40),
					Max:               (float64)(295),
					Unit:              (string)("degC"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("DALY_MonomerTemp6"),
					Start:             (uint8)(55),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(40),
					Scale:             (float64)(1),
					Min:               (float64)(40),
					Max:               (float64)(295),
					Unit:              (string)("degC"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("DALY_MonomerTemp7"),
					Start:             (uint8)(63),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(40),
					Scale:             (float64)(1),
					Min:               (float64)(40),
					Max:               (float64)(295),
					Unit:              (string)("degC"),
					Description:       (string)(""),
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
			Name:        (string)("OBC_Response"),
			ID:          (uint32)(419385573),
			IsExtended:  (bool)(true),
			Length:      (uint8)(8),
			SendType:    (descriptor.SendType)(0),
			Description: (string)(""),
			Signals: ([]*descriptor.Signal)([]*descriptor.Signal{
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("OBC_OutputVoltage"),
					Start:             (uint8)(7),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(6553.5),
					Unit:              (string)("V"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("OBC_OutputCurrent"),
					Start:             (uint8)(23),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(6553.5),
					Unit:              (string)("A"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:             (string)("OBC_HardwareProtection"),
					Start:            (uint8)(32),
					Length:           (uint8)(1),
					IsBigEndian:      (bool)(true),
					IsSigned:         (bool)(false),
					IsFloat:          (bool)(false),
					IsMultiplexer:    (bool)(false),
					IsMultiplexed:    (bool)(false),
					MultiplexerValue: (uint)(0),
					Offset:           (float64)(0),
					Scale:            (float64)(1),
					Min:              (float64)(0),
					Max:              (float64)(1),
					Unit:             (string)(""),
					Description:      (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)([]*descriptor.ValueDescription{
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(0),
							Description: (string)("Normal"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(1),
							Description: (string)("HardwareProtection"),
						}),
					}),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:             (string)("OBC_TemperatureProtection"),
					Start:            (uint8)(33),
					Length:           (uint8)(1),
					IsBigEndian:      (bool)(true),
					IsSigned:         (bool)(false),
					IsFloat:          (bool)(false),
					IsMultiplexer:    (bool)(false),
					IsMultiplexed:    (bool)(false),
					MultiplexerValue: (uint)(0),
					Offset:           (float64)(0),
					Scale:            (float64)(1),
					Min:              (float64)(0),
					Max:              (float64)(1),
					Unit:             (string)(""),
					Description:      (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)([]*descriptor.ValueDescription{
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(0),
							Description: (string)("Normal"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(1),
							Description: (string)("InternalTemperatureProtection"),
						}),
					}),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:             (string)("OBC_InputVoltageStatus"),
					Start:            (uint8)(35),
					Length:           (uint8)(2),
					IsBigEndian:      (bool)(true),
					IsSigned:         (bool)(false),
					IsFloat:          (bool)(false),
					IsMultiplexer:    (bool)(false),
					IsMultiplexed:    (bool)(false),
					MultiplexerValue: (uint)(0),
					Offset:           (float64)(0),
					Scale:            (float64)(1),
					Min:              (float64)(0),
					Max:              (float64)(3),
					Unit:             (string)(""),
					Description:      (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)([]*descriptor.ValueDescription{
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(0),
							Description: (string)("TheVoltageIsNormal"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(1),
							Description: (string)("InputUnderVoltage"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(2),
							Description: (string)("InputOverVoltage"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(3),
							Description: (string)("NoInputVoltage"),
						}),
					}),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:             (string)("OBC_OutputUnderVoltage"),
					Start:            (uint8)(36),
					Length:           (uint8)(1),
					IsBigEndian:      (bool)(true),
					IsSigned:         (bool)(false),
					IsFloat:          (bool)(false),
					IsMultiplexer:    (bool)(false),
					IsMultiplexed:    (bool)(false),
					MultiplexerValue: (uint)(0),
					Offset:           (float64)(0),
					Scale:            (float64)(1),
					Min:              (float64)(0),
					Max:              (float64)(1),
					Unit:             (string)(""),
					Description:      (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)([]*descriptor.ValueDescription{
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(0),
							Description: (string)("Normal"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(1),
							Description: (string)("Fault"),
						}),
					}),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:             (string)("OBC_OutputOverVoltage"),
					Start:            (uint8)(37),
					Length:           (uint8)(1),
					IsBigEndian:      (bool)(true),
					IsSigned:         (bool)(false),
					IsFloat:          (bool)(false),
					IsMultiplexer:    (bool)(false),
					IsMultiplexed:    (bool)(false),
					MultiplexerValue: (uint)(0),
					Offset:           (float64)(0),
					Scale:            (float64)(1),
					Min:              (float64)(0),
					Max:              (float64)(1),
					Unit:             (string)(""),
					Description:      (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)([]*descriptor.ValueDescription{
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(0),
							Description: (string)("Normal"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(1),
							Description: (string)("Fault"),
						}),
					}),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:             (string)("OBC_OutputOvercurrent"),
					Start:            (uint8)(38),
					Length:           (uint8)(1),
					IsBigEndian:      (bool)(true),
					IsSigned:         (bool)(false),
					IsFloat:          (bool)(false),
					IsMultiplexer:    (bool)(false),
					IsMultiplexed:    (bool)(false),
					MultiplexerValue: (uint)(0),
					Offset:           (float64)(0),
					Scale:            (float64)(1),
					Min:              (float64)(0),
					Max:              (float64)(1),
					Unit:             (string)(""),
					Description:      (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)([]*descriptor.ValueDescription{
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(0),
							Description: (string)("Normal"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(1),
							Description: (string)("Fault"),
						}),
					}),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:             (string)("OBC_OutputShortCircuit"),
					Start:            (uint8)(39),
					Length:           (uint8)(1),
					IsBigEndian:      (bool)(true),
					IsSigned:         (bool)(false),
					IsFloat:          (bool)(false),
					IsMultiplexer:    (bool)(false),
					IsMultiplexed:    (bool)(false),
					MultiplexerValue: (uint)(0),
					Offset:           (float64)(0),
					Scale:            (float64)(1),
					Min:              (float64)(0),
					Max:              (float64)(1),
					Unit:             (string)(""),
					Description:      (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)([]*descriptor.ValueDescription{
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(0),
							Description: (string)("Normal"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(1),
							Description: (string)("Fault"),
						}),
					}),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:             (string)("OBC_CommunicationStatus"),
					Start:            (uint8)(40),
					Length:           (uint8)(1),
					IsBigEndian:      (bool)(true),
					IsSigned:         (bool)(false),
					IsFloat:          (bool)(false),
					IsMultiplexer:    (bool)(false),
					IsMultiplexed:    (bool)(false),
					MultiplexerValue: (uint)(0),
					Offset:           (float64)(0),
					Scale:            (float64)(1),
					Min:              (float64)(0),
					Max:              (float64)(1),
					Unit:             (string)(""),
					Description:      (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)([]*descriptor.ValueDescription{
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(0),
							Description: (string)("CommunicationIsNormal"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(1),
							Description: (string)("ReceiveCommunicationTimeout"),
						}),
					}),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:             (string)("OBC_WorkingStatus"),
					Start:            (uint8)(42),
					Length:           (uint8)(2),
					IsBigEndian:      (bool)(true),
					IsSigned:         (bool)(false),
					IsFloat:          (bool)(false),
					IsMultiplexer:    (bool)(false),
					IsMultiplexed:    (bool)(false),
					MultiplexerValue: (uint)(0),
					Offset:           (float64)(0),
					Scale:            (float64)(1),
					Min:              (float64)(0),
					Max:              (float64)(3),
					Unit:             (string)(""),
					Description:      (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)([]*descriptor.ValueDescription{
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(0),
							Description: (string)("Undefined"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(1),
							Description: (string)("Work"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(2),
							Description: (string)("Stop"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(3),
							Description: (string)("StopOrStandBy"),
						}),
					}),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:             (string)("OBC_ComOfInitialization"),
					Start:            (uint8)(43),
					Length:           (uint8)(1),
					IsBigEndian:      (bool)(true),
					IsSigned:         (bool)(false),
					IsFloat:          (bool)(false),
					IsMultiplexer:    (bool)(false),
					IsMultiplexed:    (bool)(false),
					MultiplexerValue: (uint)(0),
					Offset:           (float64)(0),
					Scale:            (float64)(1),
					Min:              (float64)(0),
					Max:              (float64)(1),
					Unit:             (string)(""),
					Description:      (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)([]*descriptor.ValueDescription{
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(0),
							Description: (string)("IsNotComplete"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(1),
							Description: (string)("Complete"),
						}),
					}),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:             (string)("OBC_FanWorkingEnable"),
					Start:            (uint8)(44),
					Length:           (uint8)(1),
					IsBigEndian:      (bool)(true),
					IsSigned:         (bool)(false),
					IsFloat:          (bool)(false),
					IsMultiplexer:    (bool)(false),
					IsMultiplexed:    (bool)(false),
					MultiplexerValue: (uint)(0),
					Offset:           (float64)(0),
					Scale:            (float64)(1),
					Min:              (float64)(0),
					Max:              (float64)(1),
					Unit:             (string)(""),
					Description:      (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)([]*descriptor.ValueDescription{
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(0),
							Description: (string)("Close"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(1),
							Description: (string)("Open"),
						}),
					}),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:             (string)("OBC_CoolingPumpEnable"),
					Start:            (uint8)(45),
					Length:           (uint8)(1),
					IsBigEndian:      (bool)(true),
					IsSigned:         (bool)(false),
					IsFloat:          (bool)(false),
					IsMultiplexer:    (bool)(false),
					IsMultiplexed:    (bool)(false),
					MultiplexerValue: (uint)(0),
					Offset:           (float64)(0),
					Scale:            (float64)(1),
					Min:              (float64)(0),
					Max:              (float64)(1),
					Unit:             (string)(""),
					Description:      (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)([]*descriptor.ValueDescription{
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(0),
							Description: (string)("Close"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(1),
							Description: (string)("Open"),
						}),
					}),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:             (string)("OBC_CCSignalStatus"),
					Start:            (uint8)(49),
					Length:           (uint8)(2),
					IsBigEndian:      (bool)(true),
					IsSigned:         (bool)(false),
					IsFloat:          (bool)(false),
					IsMultiplexer:    (bool)(false),
					IsMultiplexed:    (bool)(false),
					MultiplexerValue: (uint)(0),
					Offset:           (float64)(0),
					Scale:            (float64)(1),
					Min:              (float64)(0),
					Max:              (float64)(3),
					Unit:             (string)(""),
					Description:      (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)([]*descriptor.ValueDescription{
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(0),
							Description: (string)("NotConnected"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(1),
							Description: (string)("HalfConnected"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(2),
							Description: (string)("NormalConnected"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(3),
							Description: (string)("ResistanceDetectionError"),
						}),
					}),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:             (string)("OBC_CPSignalStatus"),
					Start:            (uint8)(50),
					Length:           (uint8)(1),
					IsBigEndian:      (bool)(true),
					IsSigned:         (bool)(false),
					IsFloat:          (bool)(false),
					IsMultiplexer:    (bool)(false),
					IsMultiplexed:    (bool)(false),
					MultiplexerValue: (uint)(0),
					Offset:           (float64)(0),
					Scale:            (float64)(1),
					Min:              (float64)(0),
					Max:              (float64)(1),
					Unit:             (string)(""),
					Description:      (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)([]*descriptor.ValueDescription{
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(0),
							Description: (string)("NoCPSignalWasDetected"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(1),
							Description: (string)("Normal"),
						}),
					}),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:             (string)("OBC_SocketOvheatFault"),
					Start:            (uint8)(51),
					Length:           (uint8)(1),
					IsBigEndian:      (bool)(true),
					IsSigned:         (bool)(false),
					IsFloat:          (bool)(false),
					IsMultiplexer:    (bool)(false),
					IsMultiplexed:    (bool)(false),
					MultiplexerValue: (uint)(0),
					Offset:           (float64)(0),
					Scale:            (float64)(1),
					Min:              (float64)(0),
					Max:              (float64)(1),
					Unit:             (string)(""),
					Description:      (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)([]*descriptor.ValueDescription{
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(0),
							Description: (string)("Normal"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(1),
							Description: (string)("ChargingSocketIsOverheat"),
						}),
					}),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:             (string)("OBC_ElectronicLockState"),
					Start:            (uint8)(54),
					Length:           (uint8)(3),
					IsBigEndian:      (bool)(true),
					IsSigned:         (bool)(false),
					IsFloat:          (bool)(false),
					IsMultiplexer:    (bool)(false),
					IsMultiplexed:    (bool)(false),
					MultiplexerValue: (uint)(0),
					Offset:           (float64)(0),
					Scale:            (float64)(1),
					Min:              (float64)(0),
					Max:              (float64)(7),
					Unit:             (string)(""),
					Description:      (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)([]*descriptor.ValueDescription{
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(0),
							Description: (string)("InJudgment"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(1),
							Description: (string)("Locked"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(2),
							Description: (string)("Unlocked"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(3),
							Description: (string)("UnlockFault"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(4),
							Description: (string)("LockedFault"),
						}),
					}),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:             (string)("OBC_S2SwitchControlBitStatus"),
					Start:            (uint8)(55),
					Length:           (uint8)(1),
					IsBigEndian:      (bool)(true),
					IsSigned:         (bool)(false),
					IsFloat:          (bool)(false),
					IsMultiplexer:    (bool)(false),
					IsMultiplexed:    (bool)(false),
					MultiplexerValue: (uint)(0),
					Offset:           (float64)(0),
					Scale:            (float64)(1),
					Min:              (float64)(0),
					Max:              (float64)(1),
					Unit:             (string)(""),
					Description:      (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)([]*descriptor.ValueDescription{
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(0),
							Description: (string)("SwitchOff"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(1),
							Description: (string)("CloseUp"),
						}),
					}),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("OBC_Temperature"),
					Start:             (uint8)(63),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(-40),
					Scale:             (float64)(1),
					Min:               (float64)(-40),
					Max:               (float64)(215),
					Unit:              (string)(""),
					Description:       (string)(""),
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
	Nodes: ([]*descriptor.Node)([]*descriptor.Node{
		(*descriptor.Node)(&descriptor.Node{
			Name:        (string)("CReATE_ECU"),
			Description: (string)(""),
		}),
	}),
})
