// Package bit29can provides primitives for encoding and decoding bit29 CAN messages.
//
// Source: dummy/obd2/bit29.dbc
package bit29can

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
// OBD2Reader provides read access to a OBD2 message.
type OBD2Reader interface {
	can.FrameMarshaler
	// Length returns the value of the Length signal.
	Length() uint8
	// Service returns the value of the Service signal.
	Service() OBD2_Service
	// Response returns the value of the Response signal.
	Response() uint8
	// S01PID returns the value of the S01PID signal.
	S01PID() OBD2_S01PID
	// S02PID returns the value of the S02PID signal.
	S02PID() OBD2_S02PID
	// S01PID00_PIDsSupported_01_20 returns the value of the S01PID00_PIDsSupported_01_20 signal.
	S01PID00_PIDsSupported_01_20() uint32
	// S01PID01_MonitorStatus returns the value of the S01PID01_MonitorStatus signal.
	S01PID01_MonitorStatus() uint32
	// S02PID02_FreezeDTC returns the value of the S02PID02_FreezeDTC signal.
	S02PID02_FreezeDTC() uint16
	// S01PID02_FreezeDTC returns the value of the S01PID02_FreezeDTC signal.
	S01PID02_FreezeDTC() uint16
	// S01PID03_FuelSystemStatus returns the value of the S01PID03_FuelSystemStatus signal.
	S01PID03_FuelSystemStatus() OBD2_S01PID03_FuelSystemStatus
	// S01PID04_CalcEngineLoad returns the physical value of the S01PID04_CalcEngineLoad signal.
	S01PID04_CalcEngineLoad() float64
	// RawS01PID04_CalcEngineLoad returns the raw (encoded) value of the S01PID04_CalcEngineLoad signal.
	RawS01PID04_CalcEngineLoad() uint8
	// S01PID05_EngineCoolantTemp returns the physical value of the S01PID05_EngineCoolantTemp signal.
	S01PID05_EngineCoolantTemp() float64
	// RawS01PID05_EngineCoolantTemp returns the raw (encoded) value of the S01PID05_EngineCoolantTemp signal.
	RawS01PID05_EngineCoolantTemp() uint8
	// S01PID06_ShortFuelTrimBank1 returns the physical value of the S01PID06_ShortFuelTrimBank1 signal.
	S01PID06_ShortFuelTrimBank1() float64
	// RawS01PID06_ShortFuelTrimBank1 returns the raw (encoded) value of the S01PID06_ShortFuelTrimBank1 signal.
	RawS01PID06_ShortFuelTrimBank1() uint8
	// S01PID07_LongFuelTrimBank1 returns the physical value of the S01PID07_LongFuelTrimBank1 signal.
	S01PID07_LongFuelTrimBank1() float64
	// RawS01PID07_LongFuelTrimBank1 returns the raw (encoded) value of the S01PID07_LongFuelTrimBank1 signal.
	RawS01PID07_LongFuelTrimBank1() uint8
	// S01PID08_ShortFuelTrimBank2 returns the physical value of the S01PID08_ShortFuelTrimBank2 signal.
	S01PID08_ShortFuelTrimBank2() float64
	// RawS01PID08_ShortFuelTrimBank2 returns the raw (encoded) value of the S01PID08_ShortFuelTrimBank2 signal.
	RawS01PID08_ShortFuelTrimBank2() uint8
	// S01PID09_LongFuelTrimBank2 returns the physical value of the S01PID09_LongFuelTrimBank2 signal.
	S01PID09_LongFuelTrimBank2() float64
	// RawS01PID09_LongFuelTrimBank2 returns the raw (encoded) value of the S01PID09_LongFuelTrimBank2 signal.
	RawS01PID09_LongFuelTrimBank2() uint8
	// S01PID0A_FuelPressure returns the physical value of the S01PID0A_FuelPressure signal.
	S01PID0A_FuelPressure() float64
	// RawS01PID0A_FuelPressure returns the raw (encoded) value of the S01PID0A_FuelPressure signal.
	RawS01PID0A_FuelPressure() uint8
	// S01PID0B_IntakeManiAbsPress returns the value of the S01PID0B_IntakeManiAbsPress signal.
	S01PID0B_IntakeManiAbsPress() uint8
	// S01PID0C_EngineRPM returns the physical value of the S01PID0C_EngineRPM signal.
	S01PID0C_EngineRPM() float64
	// RawS01PID0C_EngineRPM returns the raw (encoded) value of the S01PID0C_EngineRPM signal.
	RawS01PID0C_EngineRPM() uint16
	// S01PID0D_VehicleSpeed returns the value of the S01PID0D_VehicleSpeed signal.
	S01PID0D_VehicleSpeed() uint8
	// S01PID0E_TimingAdvance returns the physical value of the S01PID0E_TimingAdvance signal.
	S01PID0E_TimingAdvance() float64
	// RawS01PID0E_TimingAdvance returns the raw (encoded) value of the S01PID0E_TimingAdvance signal.
	RawS01PID0E_TimingAdvance() uint8
	// S01PID0F_IntakeAirTemperature returns the physical value of the S01PID0F_IntakeAirTemperature signal.
	S01PID0F_IntakeAirTemperature() float64
	// RawS01PID0F_IntakeAirTemperature returns the raw (encoded) value of the S01PID0F_IntakeAirTemperature signal.
	RawS01PID0F_IntakeAirTemperature() uint8
	// S01PID10_MAFAirFlowRate returns the physical value of the S01PID10_MAFAirFlowRate signal.
	S01PID10_MAFAirFlowRate() float64
	// RawS01PID10_MAFAirFlowRate returns the raw (encoded) value of the S01PID10_MAFAirFlowRate signal.
	RawS01PID10_MAFAirFlowRate() uint16
	// S01PID11_ThrottlePosition returns the physical value of the S01PID11_ThrottlePosition signal.
	S01PID11_ThrottlePosition() float64
	// RawS01PID11_ThrottlePosition returns the raw (encoded) value of the S01PID11_ThrottlePosition signal.
	RawS01PID11_ThrottlePosition() uint8
	// S01PID12_CmdSecAirStatus returns the value of the S01PID12_CmdSecAirStatus signal.
	S01PID12_CmdSecAirStatus() OBD2_S01PID12_CmdSecAirStatus
	// S01PID14_OxySensor1_Volt returns the physical value of the S01PID14_OxySensor1_Volt signal.
	S01PID14_OxySensor1_Volt() float64
	// RawS01PID14_OxySensor1_Volt returns the raw (encoded) value of the S01PID14_OxySensor1_Volt signal.
	RawS01PID14_OxySensor1_Volt() uint8
	// S01PID15_OxySensor2_Volt returns the physical value of the S01PID15_OxySensor2_Volt signal.
	S01PID15_OxySensor2_Volt() float64
	// RawS01PID15_OxySensor2_Volt returns the raw (encoded) value of the S01PID15_OxySensor2_Volt signal.
	RawS01PID15_OxySensor2_Volt() uint8
	// S01PID16_OxySensor3_Volt returns the physical value of the S01PID16_OxySensor3_Volt signal.
	S01PID16_OxySensor3_Volt() float64
	// RawS01PID16_OxySensor3_Volt returns the raw (encoded) value of the S01PID16_OxySensor3_Volt signal.
	RawS01PID16_OxySensor3_Volt() uint8
	// S01PID17_OxySensor4_Volt returns the physical value of the S01PID17_OxySensor4_Volt signal.
	S01PID17_OxySensor4_Volt() float64
	// RawS01PID17_OxySensor4_Volt returns the raw (encoded) value of the S01PID17_OxySensor4_Volt signal.
	RawS01PID17_OxySensor4_Volt() uint8
	// S01PID18_OxySensor5_Volt returns the physical value of the S01PID18_OxySensor5_Volt signal.
	S01PID18_OxySensor5_Volt() float64
	// RawS01PID18_OxySensor5_Volt returns the raw (encoded) value of the S01PID18_OxySensor5_Volt signal.
	RawS01PID18_OxySensor5_Volt() uint8
	// S01PID19_OxySensor6_Volt returns the physical value of the S01PID19_OxySensor6_Volt signal.
	S01PID19_OxySensor6_Volt() float64
	// RawS01PID19_OxySensor6_Volt returns the raw (encoded) value of the S01PID19_OxySensor6_Volt signal.
	RawS01PID19_OxySensor6_Volt() uint8
	// S01PID14_OxySensor1_STFT returns the physical value of the S01PID14_OxySensor1_STFT signal.
	S01PID14_OxySensor1_STFT() float64
	// RawS01PID14_OxySensor1_STFT returns the raw (encoded) value of the S01PID14_OxySensor1_STFT signal.
	RawS01PID14_OxySensor1_STFT() uint8
	// S01PID15_OxySensor2_STFT returns the physical value of the S01PID15_OxySensor2_STFT signal.
	S01PID15_OxySensor2_STFT() float64
	// RawS01PID15_OxySensor2_STFT returns the raw (encoded) value of the S01PID15_OxySensor2_STFT signal.
	RawS01PID15_OxySensor2_STFT() uint8
	// S01PID16_OxySensor3_STFT returns the physical value of the S01PID16_OxySensor3_STFT signal.
	S01PID16_OxySensor3_STFT() float64
	// RawS01PID16_OxySensor3_STFT returns the raw (encoded) value of the S01PID16_OxySensor3_STFT signal.
	RawS01PID16_OxySensor3_STFT() uint8
	// S01PID17_OxySensor4_STFT returns the physical value of the S01PID17_OxySensor4_STFT signal.
	S01PID17_OxySensor4_STFT() float64
	// RawS01PID17_OxySensor4_STFT returns the raw (encoded) value of the S01PID17_OxySensor4_STFT signal.
	RawS01PID17_OxySensor4_STFT() uint8
	// S01PID18_OxySensor5_STFT returns the physical value of the S01PID18_OxySensor5_STFT signal.
	S01PID18_OxySensor5_STFT() float64
	// RawS01PID18_OxySensor5_STFT returns the raw (encoded) value of the S01PID18_OxySensor5_STFT signal.
	RawS01PID18_OxySensor5_STFT() uint8
	// S01PID19_OxySensor6_STFT returns the physical value of the S01PID19_OxySensor6_STFT signal.
	S01PID19_OxySensor6_STFT() float64
	// RawS01PID19_OxySensor6_STFT returns the raw (encoded) value of the S01PID19_OxySensor6_STFT signal.
	RawS01PID19_OxySensor6_STFT() uint8
	// S01PID1A_OxySensor7_Volt returns the physical value of the S01PID1A_OxySensor7_Volt signal.
	S01PID1A_OxySensor7_Volt() float64
	// RawS01PID1A_OxySensor7_Volt returns the raw (encoded) value of the S01PID1A_OxySensor7_Volt signal.
	RawS01PID1A_OxySensor7_Volt() uint8
	// S01PID1A_OxySensor7_STFT returns the physical value of the S01PID1A_OxySensor7_STFT signal.
	S01PID1A_OxySensor7_STFT() float64
	// RawS01PID1A_OxySensor7_STFT returns the raw (encoded) value of the S01PID1A_OxySensor7_STFT signal.
	RawS01PID1A_OxySensor7_STFT() uint8
	// S01PID1B_OxySensor8_Volt returns the physical value of the S01PID1B_OxySensor8_Volt signal.
	S01PID1B_OxySensor8_Volt() float64
	// RawS01PID1B_OxySensor8_Volt returns the raw (encoded) value of the S01PID1B_OxySensor8_Volt signal.
	RawS01PID1B_OxySensor8_Volt() uint8
	// S01PID1B_OxySensor8_STFT returns the physical value of the S01PID1B_OxySensor8_STFT signal.
	S01PID1B_OxySensor8_STFT() float64
	// RawS01PID1B_OxySensor8_STFT returns the raw (encoded) value of the S01PID1B_OxySensor8_STFT signal.
	RawS01PID1B_OxySensor8_STFT() uint8
	// S01PID1C_OBDStandard returns the value of the S01PID1C_OBDStandard signal.
	S01PID1C_OBDStandard() OBD2_S01PID1C_OBDStandard
	// S01PID1F_TimeSinceEngStart returns the value of the S01PID1F_TimeSinceEngStart signal.
	S01PID1F_TimeSinceEngStart() uint16
	// S01PID20_PIDsSupported_21_40 returns the value of the S01PID20_PIDsSupported_21_40 signal.
	S01PID20_PIDsSupported_21_40() uint32
	// S01PID21_DistanceMILOn returns the value of the S01PID21_DistanceMILOn signal.
	S01PID21_DistanceMILOn() uint16
	// S01PID22_FuelRailPres returns the physical value of the S01PID22_FuelRailPres signal.
	S01PID22_FuelRailPres() float64
	// RawS01PID22_FuelRailPres returns the raw (encoded) value of the S01PID22_FuelRailPres signal.
	RawS01PID22_FuelRailPres() uint16
	// S01PID23_FuelRailGaug returns the physical value of the S01PID23_FuelRailGaug signal.
	S01PID23_FuelRailGaug() float64
	// RawS01PID23_FuelRailGaug returns the raw (encoded) value of the S01PID23_FuelRailGaug signal.
	RawS01PID23_FuelRailGaug() uint16
	// S01PID24_OxySensor1_FAER returns the physical value of the S01PID24_OxySensor1_FAER signal.
	S01PID24_OxySensor1_FAER() float64
	// RawS01PID24_OxySensor1_FAER returns the raw (encoded) value of the S01PID24_OxySensor1_FAER signal.
	RawS01PID24_OxySensor1_FAER() uint16
	// S01PID24_OxySensor1_Volt returns the physical value of the S01PID24_OxySensor1_Volt signal.
	S01PID24_OxySensor1_Volt() float64
	// RawS01PID24_OxySensor1_Volt returns the raw (encoded) value of the S01PID24_OxySensor1_Volt signal.
	RawS01PID24_OxySensor1_Volt() uint16
	// S01PID25_OxySensor2_FAER returns the physical value of the S01PID25_OxySensor2_FAER signal.
	S01PID25_OxySensor2_FAER() float64
	// RawS01PID25_OxySensor2_FAER returns the raw (encoded) value of the S01PID25_OxySensor2_FAER signal.
	RawS01PID25_OxySensor2_FAER() uint16
	// S01PID25_OxySensor2_Volt returns the physical value of the S01PID25_OxySensor2_Volt signal.
	S01PID25_OxySensor2_Volt() float64
	// RawS01PID25_OxySensor2_Volt returns the raw (encoded) value of the S01PID25_OxySensor2_Volt signal.
	RawS01PID25_OxySensor2_Volt() uint16
	// S01PID26_OxySensor3_FAER returns the physical value of the S01PID26_OxySensor3_FAER signal.
	S01PID26_OxySensor3_FAER() float64
	// RawS01PID26_OxySensor3_FAER returns the raw (encoded) value of the S01PID26_OxySensor3_FAER signal.
	RawS01PID26_OxySensor3_FAER() uint16
	// S01PID26_OxySensor3_Volt returns the physical value of the S01PID26_OxySensor3_Volt signal.
	S01PID26_OxySensor3_Volt() float64
	// RawS01PID26_OxySensor3_Volt returns the raw (encoded) value of the S01PID26_OxySensor3_Volt signal.
	RawS01PID26_OxySensor3_Volt() uint16
	// S01PID27_OxySensor4_FAER returns the physical value of the S01PID27_OxySensor4_FAER signal.
	S01PID27_OxySensor4_FAER() float64
	// RawS01PID27_OxySensor4_FAER returns the raw (encoded) value of the S01PID27_OxySensor4_FAER signal.
	RawS01PID27_OxySensor4_FAER() uint16
	// S01PID28_OxySensor5_FAER returns the physical value of the S01PID28_OxySensor5_FAER signal.
	S01PID28_OxySensor5_FAER() float64
	// RawS01PID28_OxySensor5_FAER returns the raw (encoded) value of the S01PID28_OxySensor5_FAER signal.
	RawS01PID28_OxySensor5_FAER() uint16
	// S01PID29_OxySensor6_FAER returns the physical value of the S01PID29_OxySensor6_FAER signal.
	S01PID29_OxySensor6_FAER() float64
	// RawS01PID29_OxySensor6_FAER returns the raw (encoded) value of the S01PID29_OxySensor6_FAER signal.
	RawS01PID29_OxySensor6_FAER() uint16
	// S01PID27_OxySensor4_Volt returns the physical value of the S01PID27_OxySensor4_Volt signal.
	S01PID27_OxySensor4_Volt() float64
	// RawS01PID27_OxySensor4_Volt returns the raw (encoded) value of the S01PID27_OxySensor4_Volt signal.
	RawS01PID27_OxySensor4_Volt() uint16
	// S01PID28_OxySensor5_Volt returns the physical value of the S01PID28_OxySensor5_Volt signal.
	S01PID28_OxySensor5_Volt() float64
	// RawS01PID28_OxySensor5_Volt returns the raw (encoded) value of the S01PID28_OxySensor5_Volt signal.
	RawS01PID28_OxySensor5_Volt() uint16
	// S01PID29_OxySensor6_Volt returns the physical value of the S01PID29_OxySensor6_Volt signal.
	S01PID29_OxySensor6_Volt() float64
	// RawS01PID29_OxySensor6_Volt returns the raw (encoded) value of the S01PID29_OxySensor6_Volt signal.
	RawS01PID29_OxySensor6_Volt() uint16
	// S01PID2A_OxySensor7_FAER returns the physical value of the S01PID2A_OxySensor7_FAER signal.
	S01PID2A_OxySensor7_FAER() float64
	// RawS01PID2A_OxySensor7_FAER returns the raw (encoded) value of the S01PID2A_OxySensor7_FAER signal.
	RawS01PID2A_OxySensor7_FAER() uint16
	// S01PID2A_OxySensor7_Volt returns the physical value of the S01PID2A_OxySensor7_Volt signal.
	S01PID2A_OxySensor7_Volt() float64
	// RawS01PID2A_OxySensor7_Volt returns the raw (encoded) value of the S01PID2A_OxySensor7_Volt signal.
	RawS01PID2A_OxySensor7_Volt() uint16
	// S01PID2B_OxySensor8_FAER returns the physical value of the S01PID2B_OxySensor8_FAER signal.
	S01PID2B_OxySensor8_FAER() float64
	// RawS01PID2B_OxySensor8_FAER returns the raw (encoded) value of the S01PID2B_OxySensor8_FAER signal.
	RawS01PID2B_OxySensor8_FAER() uint16
	// S01PID2B_OxySensor8_Volt returns the physical value of the S01PID2B_OxySensor8_Volt signal.
	S01PID2B_OxySensor8_Volt() float64
	// RawS01PID2B_OxySensor8_Volt returns the raw (encoded) value of the S01PID2B_OxySensor8_Volt signal.
	RawS01PID2B_OxySensor8_Volt() uint16
	// S01PID2C_CmdEGR returns the physical value of the S01PID2C_CmdEGR signal.
	S01PID2C_CmdEGR() float64
	// RawS01PID2C_CmdEGR returns the raw (encoded) value of the S01PID2C_CmdEGR signal.
	RawS01PID2C_CmdEGR() uint8
	// S01PID2D_EGRError returns the physical value of the S01PID2D_EGRError signal.
	S01PID2D_EGRError() float64
	// RawS01PID2D_EGRError returns the raw (encoded) value of the S01PID2D_EGRError signal.
	RawS01PID2D_EGRError() uint8
	// S01PID2E_CmdEvapPurge returns the physical value of the S01PID2E_CmdEvapPurge signal.
	S01PID2E_CmdEvapPurge() float64
	// RawS01PID2E_CmdEvapPurge returns the raw (encoded) value of the S01PID2E_CmdEvapPurge signal.
	RawS01PID2E_CmdEvapPurge() uint8
	// S01PID2F_FuelTankLevel returns the physical value of the S01PID2F_FuelTankLevel signal.
	S01PID2F_FuelTankLevel() float64
	// RawS01PID2F_FuelTankLevel returns the raw (encoded) value of the S01PID2F_FuelTankLevel signal.
	RawS01PID2F_FuelTankLevel() uint8
	// S01PID30_WarmUpsSinceCodeClear returns the value of the S01PID30_WarmUpsSinceCodeClear signal.
	S01PID30_WarmUpsSinceCodeClear() uint8
	// S01PID31_DistanceSinceCodeClear returns the value of the S01PID31_DistanceSinceCodeClear signal.
	S01PID31_DistanceSinceCodeClear() uint16
	// S01PID32_EvapSysVaporPres returns the physical value of the S01PID32_EvapSysVaporPres signal.
	S01PID32_EvapSysVaporPres() float64
	// RawS01PID32_EvapSysVaporPres returns the raw (encoded) value of the S01PID32_EvapSysVaporPres signal.
	RawS01PID32_EvapSysVaporPres() int16
	// S01PID33_AbsBaroPres returns the value of the S01PID33_AbsBaroPres signal.
	S01PID33_AbsBaroPres() uint8
	// S01PID34_OxySensor1_FAER returns the physical value of the S01PID34_OxySensor1_FAER signal.
	S01PID34_OxySensor1_FAER() float64
	// RawS01PID34_OxySensor1_FAER returns the raw (encoded) value of the S01PID34_OxySensor1_FAER signal.
	RawS01PID34_OxySensor1_FAER() uint16
	// S01PID34_OxySensor1_Crnt returns the physical value of the S01PID34_OxySensor1_Crnt signal.
	S01PID34_OxySensor1_Crnt() float64
	// RawS01PID34_OxySensor1_Crnt returns the raw (encoded) value of the S01PID34_OxySensor1_Crnt signal.
	RawS01PID34_OxySensor1_Crnt() uint16
	// S01PID35_OxySensor2_FAER returns the physical value of the S01PID35_OxySensor2_FAER signal.
	S01PID35_OxySensor2_FAER() float64
	// RawS01PID35_OxySensor2_FAER returns the raw (encoded) value of the S01PID35_OxySensor2_FAER signal.
	RawS01PID35_OxySensor2_FAER() uint16
	// S01PID35_OxySensor2_Crnt returns the physical value of the S01PID35_OxySensor2_Crnt signal.
	S01PID35_OxySensor2_Crnt() float64
	// RawS01PID35_OxySensor2_Crnt returns the raw (encoded) value of the S01PID35_OxySensor2_Crnt signal.
	RawS01PID35_OxySensor2_Crnt() uint16
	// S01PID36_OxySensor3_FAER returns the physical value of the S01PID36_OxySensor3_FAER signal.
	S01PID36_OxySensor3_FAER() float64
	// RawS01PID36_OxySensor3_FAER returns the raw (encoded) value of the S01PID36_OxySensor3_FAER signal.
	RawS01PID36_OxySensor3_FAER() uint16
	// S01PID36_OxySensor3_Crnt returns the physical value of the S01PID36_OxySensor3_Crnt signal.
	S01PID36_OxySensor3_Crnt() float64
	// RawS01PID36_OxySensor3_Crnt returns the raw (encoded) value of the S01PID36_OxySensor3_Crnt signal.
	RawS01PID36_OxySensor3_Crnt() uint16
	// S01PID37_OxySensor4_FAER returns the physical value of the S01PID37_OxySensor4_FAER signal.
	S01PID37_OxySensor4_FAER() float64
	// RawS01PID37_OxySensor4_FAER returns the raw (encoded) value of the S01PID37_OxySensor4_FAER signal.
	RawS01PID37_OxySensor4_FAER() uint16
	// S01PID38_OxySensor5_FAER returns the physical value of the S01PID38_OxySensor5_FAER signal.
	S01PID38_OxySensor5_FAER() float64
	// RawS01PID38_OxySensor5_FAER returns the raw (encoded) value of the S01PID38_OxySensor5_FAER signal.
	RawS01PID38_OxySensor5_FAER() uint16
	// S01PID39_OxySensor6_FAER returns the physical value of the S01PID39_OxySensor6_FAER signal.
	S01PID39_OxySensor6_FAER() float64
	// RawS01PID39_OxySensor6_FAER returns the raw (encoded) value of the S01PID39_OxySensor6_FAER signal.
	RawS01PID39_OxySensor6_FAER() uint16
	// S01PID37_OxySensor4_Crnt returns the physical value of the S01PID37_OxySensor4_Crnt signal.
	S01PID37_OxySensor4_Crnt() float64
	// RawS01PID37_OxySensor4_Crnt returns the raw (encoded) value of the S01PID37_OxySensor4_Crnt signal.
	RawS01PID37_OxySensor4_Crnt() uint16
	// S01PID38_OxySensor5_Crnt returns the physical value of the S01PID38_OxySensor5_Crnt signal.
	S01PID38_OxySensor5_Crnt() float64
	// RawS01PID38_OxySensor5_Crnt returns the raw (encoded) value of the S01PID38_OxySensor5_Crnt signal.
	RawS01PID38_OxySensor5_Crnt() uint16
	// S01PID39_OxySensor6_Crnt returns the physical value of the S01PID39_OxySensor6_Crnt signal.
	S01PID39_OxySensor6_Crnt() float64
	// RawS01PID39_OxySensor6_Crnt returns the raw (encoded) value of the S01PID39_OxySensor6_Crnt signal.
	RawS01PID39_OxySensor6_Crnt() uint16
	// S01PID3A_OxySensor7_FAER returns the physical value of the S01PID3A_OxySensor7_FAER signal.
	S01PID3A_OxySensor7_FAER() float64
	// RawS01PID3A_OxySensor7_FAER returns the raw (encoded) value of the S01PID3A_OxySensor7_FAER signal.
	RawS01PID3A_OxySensor7_FAER() uint16
	// S01PID3B_OxySensor8_FAER returns the physical value of the S01PID3B_OxySensor8_FAER signal.
	S01PID3B_OxySensor8_FAER() float64
	// RawS01PID3B_OxySensor8_FAER returns the raw (encoded) value of the S01PID3B_OxySensor8_FAER signal.
	RawS01PID3B_OxySensor8_FAER() uint16
	// S01PID3C_CatTempBank1Sens1 returns the physical value of the S01PID3C_CatTempBank1Sens1 signal.
	S01PID3C_CatTempBank1Sens1() float64
	// RawS01PID3C_CatTempBank1Sens1 returns the raw (encoded) value of the S01PID3C_CatTempBank1Sens1 signal.
	RawS01PID3C_CatTempBank1Sens1() uint16
	// S01PID3D_CatTempBank2Sens1 returns the physical value of the S01PID3D_CatTempBank2Sens1 signal.
	S01PID3D_CatTempBank2Sens1() float64
	// RawS01PID3D_CatTempBank2Sens1 returns the raw (encoded) value of the S01PID3D_CatTempBank2Sens1 signal.
	RawS01PID3D_CatTempBank2Sens1() uint16
	// S01PID3A_OxySensor7_Crnt returns the physical value of the S01PID3A_OxySensor7_Crnt signal.
	S01PID3A_OxySensor7_Crnt() float64
	// RawS01PID3A_OxySensor7_Crnt returns the raw (encoded) value of the S01PID3A_OxySensor7_Crnt signal.
	RawS01PID3A_OxySensor7_Crnt() uint16
	// S01PID3B_OxySensor8_Crnt returns the physical value of the S01PID3B_OxySensor8_Crnt signal.
	S01PID3B_OxySensor8_Crnt() float64
	// RawS01PID3B_OxySensor8_Crnt returns the raw (encoded) value of the S01PID3B_OxySensor8_Crnt signal.
	RawS01PID3B_OxySensor8_Crnt() uint16
	// S01PID3E_CatTempBank1Sens2 returns the physical value of the S01PID3E_CatTempBank1Sens2 signal.
	S01PID3E_CatTempBank1Sens2() float64
	// RawS01PID3E_CatTempBank1Sens2 returns the raw (encoded) value of the S01PID3E_CatTempBank1Sens2 signal.
	RawS01PID3E_CatTempBank1Sens2() uint16
	// S01PID3F_CatTempBank2Sens2 returns the physical value of the S01PID3F_CatTempBank2Sens2 signal.
	S01PID3F_CatTempBank2Sens2() float64
	// RawS01PID3F_CatTempBank2Sens2 returns the raw (encoded) value of the S01PID3F_CatTempBank2Sens2 signal.
	RawS01PID3F_CatTempBank2Sens2() uint16
	// S01PID40_PIDsSupported_41_60 returns the value of the S01PID40_PIDsSupported_41_60 signal.
	S01PID40_PIDsSupported_41_60() uint32
	// S01PID41_MonStatusDriveCycle returns the value of the S01PID41_MonStatusDriveCycle signal.
	S01PID41_MonStatusDriveCycle() uint32
	// S01PID42_ControlModuleVolt returns the physical value of the S01PID42_ControlModuleVolt signal.
	S01PID42_ControlModuleVolt() float64
	// RawS01PID42_ControlModuleVolt returns the raw (encoded) value of the S01PID42_ControlModuleVolt signal.
	RawS01PID42_ControlModuleVolt() uint16
	// S01PID43_AbsLoadValue returns the physical value of the S01PID43_AbsLoadValue signal.
	S01PID43_AbsLoadValue() float64
	// RawS01PID43_AbsLoadValue returns the raw (encoded) value of the S01PID43_AbsLoadValue signal.
	RawS01PID43_AbsLoadValue() uint16
	// S01PID44_FuelAirCmdEquiv returns the physical value of the S01PID44_FuelAirCmdEquiv signal.
	S01PID44_FuelAirCmdEquiv() float64
	// RawS01PID44_FuelAirCmdEquiv returns the raw (encoded) value of the S01PID44_FuelAirCmdEquiv signal.
	RawS01PID44_FuelAirCmdEquiv() uint16
	// S01PID45_RelThrottlePos returns the physical value of the S01PID45_RelThrottlePos signal.
	S01PID45_RelThrottlePos() float64
	// RawS01PID45_RelThrottlePos returns the raw (encoded) value of the S01PID45_RelThrottlePos signal.
	RawS01PID45_RelThrottlePos() uint8
	// S01PID46_AmbientAirTemp returns the physical value of the S01PID46_AmbientAirTemp signal.
	S01PID46_AmbientAirTemp() float64
	// RawS01PID46_AmbientAirTemp returns the raw (encoded) value of the S01PID46_AmbientAirTemp signal.
	RawS01PID46_AmbientAirTemp() uint8
	// S01PID47_AbsThrottlePosB returns the physical value of the S01PID47_AbsThrottlePosB signal.
	S01PID47_AbsThrottlePosB() float64
	// RawS01PID47_AbsThrottlePosB returns the raw (encoded) value of the S01PID47_AbsThrottlePosB signal.
	RawS01PID47_AbsThrottlePosB() uint8
	// S01PID48_AbsThrottlePosC returns the physical value of the S01PID48_AbsThrottlePosC signal.
	S01PID48_AbsThrottlePosC() float64
	// RawS01PID48_AbsThrottlePosC returns the raw (encoded) value of the S01PID48_AbsThrottlePosC signal.
	RawS01PID48_AbsThrottlePosC() uint8
	// S01PID49_AbsThrottlePosD returns the physical value of the S01PID49_AbsThrottlePosD signal.
	S01PID49_AbsThrottlePosD() float64
	// RawS01PID49_AbsThrottlePosD returns the raw (encoded) value of the S01PID49_AbsThrottlePosD signal.
	RawS01PID49_AbsThrottlePosD() uint8
	// S01PID4A_AbsThrottlePosE returns the physical value of the S01PID4A_AbsThrottlePosE signal.
	S01PID4A_AbsThrottlePosE() float64
	// RawS01PID4A_AbsThrottlePosE returns the raw (encoded) value of the S01PID4A_AbsThrottlePosE signal.
	RawS01PID4A_AbsThrottlePosE() uint8
	// S01PID4B_AbsThrottlePosF returns the physical value of the S01PID4B_AbsThrottlePosF signal.
	S01PID4B_AbsThrottlePosF() float64
	// RawS01PID4B_AbsThrottlePosF returns the raw (encoded) value of the S01PID4B_AbsThrottlePosF signal.
	RawS01PID4B_AbsThrottlePosF() uint8
	// S01PID4C_CmdThrottleAct returns the physical value of the S01PID4C_CmdThrottleAct signal.
	S01PID4C_CmdThrottleAct() float64
	// RawS01PID4C_CmdThrottleAct returns the raw (encoded) value of the S01PID4C_CmdThrottleAct signal.
	RawS01PID4C_CmdThrottleAct() uint8
	// S01PID4D_TimeRunMILOn returns the value of the S01PID4D_TimeRunMILOn signal.
	S01PID4D_TimeRunMILOn() uint16
	// S01PID4E_TimeSinceCodeClear returns the value of the S01PID4E_TimeSinceCodeClear signal.
	S01PID4E_TimeSinceCodeClear() uint16
	// S01PID4F_Max_FAER returns the value of the S01PID4F_Max_FAER signal.
	S01PID4F_Max_FAER() uint8
	// S01PID4F_Max_OxySensVol returns the value of the S01PID4F_Max_OxySensVol signal.
	S01PID4F_Max_OxySensVol() uint8
	// S01PID4F_Max_OxySensCrnt returns the value of the S01PID4F_Max_OxySensCrnt signal.
	S01PID4F_Max_OxySensCrnt() uint8
	// S01PID4F_Max_IntManiAbsPres returns the physical value of the S01PID4F_Max_IntManiAbsPres signal.
	S01PID4F_Max_IntManiAbsPres() float64
	// RawS01PID4F_Max_IntManiAbsPres returns the raw (encoded) value of the S01PID4F_Max_IntManiAbsPres signal.
	RawS01PID4F_Max_IntManiAbsPres() uint8
	// S01PID50_Max_AirFlowMAF returns the physical value of the S01PID50_Max_AirFlowMAF signal.
	S01PID50_Max_AirFlowMAF() float64
	// RawS01PID50_Max_AirFlowMAF returns the raw (encoded) value of the S01PID50_Max_AirFlowMAF signal.
	RawS01PID50_Max_AirFlowMAF() uint8
	// S01PID51_FuelType returns the value of the S01PID51_FuelType signal.
	S01PID51_FuelType() OBD2_S01PID51_FuelType
	// S01PID52_EthanolFuelPct returns the physical value of the S01PID52_EthanolFuelPct signal.
	S01PID52_EthanolFuelPct() float64
	// RawS01PID52_EthanolFuelPct returns the raw (encoded) value of the S01PID52_EthanolFuelPct signal.
	RawS01PID52_EthanolFuelPct() uint8
	// S01PID53_AbsEvapSysVapPres returns the physical value of the S01PID53_AbsEvapSysVapPres signal.
	S01PID53_AbsEvapSysVapPres() float64
	// RawS01PID53_AbsEvapSysVapPres returns the raw (encoded) value of the S01PID53_AbsEvapSysVapPres signal.
	RawS01PID53_AbsEvapSysVapPres() uint16
	// S01PID54_EvapSysVapPres returns the physical value of the S01PID54_EvapSysVapPres signal.
	S01PID54_EvapSysVapPres() float64
	// RawS01PID54_EvapSysVapPres returns the raw (encoded) value of the S01PID54_EvapSysVapPres signal.
	RawS01PID54_EvapSysVapPres() uint16
	// S01PID55_ShortSecOxyTrimBank1 returns the physical value of the S01PID55_ShortSecOxyTrimBank1 signal.
	S01PID55_ShortSecOxyTrimBank1() float64
	// RawS01PID55_ShortSecOxyTrimBank1 returns the raw (encoded) value of the S01PID55_ShortSecOxyTrimBank1 signal.
	RawS01PID55_ShortSecOxyTrimBank1() uint8
	// S01PID56_LongSecOxyTrimBank1 returns the physical value of the S01PID56_LongSecOxyTrimBank1 signal.
	S01PID56_LongSecOxyTrimBank1() float64
	// RawS01PID56_LongSecOxyTrimBank1 returns the raw (encoded) value of the S01PID56_LongSecOxyTrimBank1 signal.
	RawS01PID56_LongSecOxyTrimBank1() uint8
	// S01PID55_ShortSecOxyTrimBank3 returns the physical value of the S01PID55_ShortSecOxyTrimBank3 signal.
	S01PID55_ShortSecOxyTrimBank3() float64
	// RawS01PID55_ShortSecOxyTrimBank3 returns the raw (encoded) value of the S01PID55_ShortSecOxyTrimBank3 signal.
	RawS01PID55_ShortSecOxyTrimBank3() uint8
	// S01PID56_LongSecOxyTrimBank3 returns the physical value of the S01PID56_LongSecOxyTrimBank3 signal.
	S01PID56_LongSecOxyTrimBank3() float64
	// RawS01PID56_LongSecOxyTrimBank3 returns the raw (encoded) value of the S01PID56_LongSecOxyTrimBank3 signal.
	RawS01PID56_LongSecOxyTrimBank3() uint8
	// S01PID57_ShortSecOxyTrimBank2 returns the physical value of the S01PID57_ShortSecOxyTrimBank2 signal.
	S01PID57_ShortSecOxyTrimBank2() float64
	// RawS01PID57_ShortSecOxyTrimBank2 returns the raw (encoded) value of the S01PID57_ShortSecOxyTrimBank2 signal.
	RawS01PID57_ShortSecOxyTrimBank2() uint8
	// S01PID58_LongSecOxyTrimBank2 returns the physical value of the S01PID58_LongSecOxyTrimBank2 signal.
	S01PID58_LongSecOxyTrimBank2() float64
	// RawS01PID58_LongSecOxyTrimBank2 returns the raw (encoded) value of the S01PID58_LongSecOxyTrimBank2 signal.
	RawS01PID58_LongSecOxyTrimBank2() uint8
	// S01PID59_FuelRailAbsPres returns the physical value of the S01PID59_FuelRailAbsPres signal.
	S01PID59_FuelRailAbsPres() float64
	// RawS01PID59_FuelRailAbsPres returns the raw (encoded) value of the S01PID59_FuelRailAbsPres signal.
	RawS01PID59_FuelRailAbsPres() uint16
	// S01PID5A_RelAccelPedalPos returns the physical value of the S01PID5A_RelAccelPedalPos signal.
	S01PID5A_RelAccelPedalPos() float64
	// RawS01PID5A_RelAccelPedalPos returns the raw (encoded) value of the S01PID5A_RelAccelPedalPos signal.
	RawS01PID5A_RelAccelPedalPos() uint8
	// S01PID57_ShortSecOxyTrimBank4 returns the physical value of the S01PID57_ShortSecOxyTrimBank4 signal.
	S01PID57_ShortSecOxyTrimBank4() float64
	// RawS01PID57_ShortSecOxyTrimBank4 returns the raw (encoded) value of the S01PID57_ShortSecOxyTrimBank4 signal.
	RawS01PID57_ShortSecOxyTrimBank4() uint8
	// S01PID58_LongSecOxyTrimBank4 returns the physical value of the S01PID58_LongSecOxyTrimBank4 signal.
	S01PID58_LongSecOxyTrimBank4() float64
	// RawS01PID58_LongSecOxyTrimBank4 returns the raw (encoded) value of the S01PID58_LongSecOxyTrimBank4 signal.
	RawS01PID58_LongSecOxyTrimBank4() uint8
	// S01PID5B_HybrBatPackRemLife returns the physical value of the S01PID5B_HybrBatPackRemLife signal.
	S01PID5B_HybrBatPackRemLife() float64
	// RawS01PID5B_HybrBatPackRemLife returns the raw (encoded) value of the S01PID5B_HybrBatPackRemLife signal.
	RawS01PID5B_HybrBatPackRemLife() uint8
	// S01PID5C_EngineOilTemp returns the physical value of the S01PID5C_EngineOilTemp signal.
	S01PID5C_EngineOilTemp() float64
	// RawS01PID5C_EngineOilTemp returns the raw (encoded) value of the S01PID5C_EngineOilTemp signal.
	RawS01PID5C_EngineOilTemp() uint8
	// S01PID5D_FuelInjectionTiming returns the physical value of the S01PID5D_FuelInjectionTiming signal.
	S01PID5D_FuelInjectionTiming() float64
	// RawS01PID5D_FuelInjectionTiming returns the raw (encoded) value of the S01PID5D_FuelInjectionTiming signal.
	RawS01PID5D_FuelInjectionTiming() uint16
	// S01PID5E_EngineFuelRate returns the physical value of the S01PID5E_EngineFuelRate signal.
	S01PID5E_EngineFuelRate() float64
	// RawS01PID5E_EngineFuelRate returns the raw (encoded) value of the S01PID5E_EngineFuelRate signal.
	RawS01PID5E_EngineFuelRate() uint16
	// S01PID5F_EmissionReq returns the value of the S01PID5F_EmissionReq signal.
	S01PID5F_EmissionReq() uint8
	// S01PID60_PIDsSupported_61_80 returns the value of the S01PID60_PIDsSupported_61_80 signal.
	S01PID60_PIDsSupported_61_80() uint32
	// S01PID61_DemandEngTorqPct returns the physical value of the S01PID61_DemandEngTorqPct signal.
	S01PID61_DemandEngTorqPct() float64
	// RawS01PID61_DemandEngTorqPct returns the raw (encoded) value of the S01PID61_DemandEngTorqPct signal.
	RawS01PID61_DemandEngTorqPct() uint8
	// S01PID62_ActualEngTorqPct returns the physical value of the S01PID62_ActualEngTorqPct signal.
	S01PID62_ActualEngTorqPct() float64
	// RawS01PID62_ActualEngTorqPct returns the raw (encoded) value of the S01PID62_ActualEngTorqPct signal.
	RawS01PID62_ActualEngTorqPct() uint8
	// S01PID63_EngRefTorq returns the value of the S01PID63_EngRefTorq signal.
	S01PID63_EngRefTorq() uint16
	// S01PID64_EngPctTorq_Idle returns the physical value of the S01PID64_EngPctTorq_Idle signal.
	S01PID64_EngPctTorq_Idle() float64
	// RawS01PID64_EngPctTorq_Idle returns the raw (encoded) value of the S01PID64_EngPctTorq_Idle signal.
	RawS01PID64_EngPctTorq_Idle() uint8
	// S01PID64_EngPctTorq_EP1 returns the physical value of the S01PID64_EngPctTorq_EP1 signal.
	S01PID64_EngPctTorq_EP1() float64
	// RawS01PID64_EngPctTorq_EP1 returns the raw (encoded) value of the S01PID64_EngPctTorq_EP1 signal.
	RawS01PID64_EngPctTorq_EP1() uint8
	// S01PID64_EngPctTorq_EP2 returns the physical value of the S01PID64_EngPctTorq_EP2 signal.
	S01PID64_EngPctTorq_EP2() float64
	// RawS01PID64_EngPctTorq_EP2 returns the raw (encoded) value of the S01PID64_EngPctTorq_EP2 signal.
	RawS01PID64_EngPctTorq_EP2() uint8
	// S01PID64_EngPctTorq_EP3 returns the physical value of the S01PID64_EngPctTorq_EP3 signal.
	S01PID64_EngPctTorq_EP3() float64
	// RawS01PID64_EngPctTorq_EP3 returns the raw (encoded) value of the S01PID64_EngPctTorq_EP3 signal.
	RawS01PID64_EngPctTorq_EP3() uint8
	// S01PID65_AuxInputOutput returns the value of the S01PID65_AuxInputOutput signal.
	S01PID65_AuxInputOutput() uint8
	// S01PID66_MAFSensor returns the value of the S01PID66_MAFSensor signal.
	S01PID66_MAFSensor() uint8
	// S01PID64_EngPctTorq_EP4 returns the physical value of the S01PID64_EngPctTorq_EP4 signal.
	S01PID64_EngPctTorq_EP4() float64
	// RawS01PID64_EngPctTorq_EP4 returns the raw (encoded) value of the S01PID64_EngPctTorq_EP4 signal.
	RawS01PID64_EngPctTorq_EP4() uint8
	// S01PID67_EngineCoolantTemp returns the value of the S01PID67_EngineCoolantTemp signal.
	S01PID67_EngineCoolantTemp() uint8
	// S01PID68_IntakeAirTempSens returns the value of the S01PID68_IntakeAirTempSens signal.
	S01PID68_IntakeAirTempSens() uint8
	// S01PID69_CmdEGR_EGRError returns the value of the S01PID69_CmdEGR_EGRError signal.
	S01PID69_CmdEGR_EGRError() uint8
	// S01PID6A_CmdDieselIntAir returns the value of the S01PID6A_CmdDieselIntAir signal.
	S01PID6A_CmdDieselIntAir() uint8
	// S01PID6B_ExhaustGasTemp returns the value of the S01PID6B_ExhaustGasTemp signal.
	S01PID6B_ExhaustGasTemp() uint8
	// S01PID6C_CmdThrottleActRel returns the value of the S01PID6C_CmdThrottleActRel signal.
	S01PID6C_CmdThrottleActRel() uint8
	// S01PID6D_FuelPresContrSys returns the value of the S01PID6D_FuelPresContrSys signal.
	S01PID6D_FuelPresContrSys() uint8
	// S01PID6E_InjPresContrSys returns the value of the S01PID6E_InjPresContrSys signal.
	S01PID6E_InjPresContrSys() uint8
	// S01PID6F_TurboComprPres returns the value of the S01PID6F_TurboComprPres signal.
	S01PID6F_TurboComprPres() uint8
	// S01PID70_BoostPresCntrl returns the value of the S01PID70_BoostPresCntrl signal.
	S01PID70_BoostPresCntrl() uint8
	// S01PID80_PIDsSupported_81_A0 returns the value of the S01PID80_PIDsSupported_81_A0 signal.
	S01PID80_PIDsSupported_81_A0() uint32
	// S01PID8E_EngFrictionPctTorq returns the physical value of the S01PID8E_EngFrictionPctTorq signal.
	S01PID8E_EngFrictionPctTorq() float64
	// RawS01PID8E_EngFrictionPctTorq returns the raw (encoded) value of the S01PID8E_EngFrictionPctTorq signal.
	RawS01PID8E_EngFrictionPctTorq() uint8
	// S01PIDA0_PIDsSupported_A1_C0 returns the value of the S01PIDA0_PIDsSupported_A1_C0 signal.
	S01PIDA0_PIDsSupported_A1_C0() uint32
	// S01PIDC0_PIDsSupported_C1_E0 returns the value of the S01PIDC0_PIDsSupported_C1_E0 signal.
	S01PIDC0_PIDsSupported_C1_E0() uint32
}

// OBD2Writer provides write access to a OBD2 message.
type OBD2Writer interface {
	// CopyFrom copies all values from OBD2.
	CopyFrom(OBD2Reader) *OBD2
	// SetLength sets the value of the Length signal.
	SetLength(uint8) *OBD2
	// SetService sets the value of the Service signal.
	SetService(OBD2_Service) *OBD2
	// SetResponse sets the value of the Response signal.
	SetResponse(uint8) *OBD2
	// SetS01PID sets the value of the S01PID signal.
	SetS01PID(OBD2_S01PID) *OBD2
	// SetS02PID sets the value of the S02PID signal.
	SetS02PID(OBD2_S02PID) *OBD2
	// SetS01PID00_PIDsSupported_01_20 sets the value of the S01PID00_PIDsSupported_01_20 signal.
	SetS01PID00_PIDsSupported_01_20(uint32) *OBD2
	// SetS01PID01_MonitorStatus sets the value of the S01PID01_MonitorStatus signal.
	SetS01PID01_MonitorStatus(uint32) *OBD2
	// SetS02PID02_FreezeDTC sets the value of the S02PID02_FreezeDTC signal.
	SetS02PID02_FreezeDTC(uint16) *OBD2
	// SetS01PID02_FreezeDTC sets the value of the S01PID02_FreezeDTC signal.
	SetS01PID02_FreezeDTC(uint16) *OBD2
	// SetS01PID03_FuelSystemStatus sets the value of the S01PID03_FuelSystemStatus signal.
	SetS01PID03_FuelSystemStatus(OBD2_S01PID03_FuelSystemStatus) *OBD2
	// SetS01PID04_CalcEngineLoad sets the physical value of the S01PID04_CalcEngineLoad signal.
	SetS01PID04_CalcEngineLoad(float64) *OBD2
	// SetRawS01PID04_CalcEngineLoad sets the raw (encoded) value of the S01PID04_CalcEngineLoad signal.
	SetRawS01PID04_CalcEngineLoad(uint8) *OBD2
	// SetS01PID05_EngineCoolantTemp sets the physical value of the S01PID05_EngineCoolantTemp signal.
	SetS01PID05_EngineCoolantTemp(float64) *OBD2
	// SetRawS01PID05_EngineCoolantTemp sets the raw (encoded) value of the S01PID05_EngineCoolantTemp signal.
	SetRawS01PID05_EngineCoolantTemp(uint8) *OBD2
	// SetS01PID06_ShortFuelTrimBank1 sets the physical value of the S01PID06_ShortFuelTrimBank1 signal.
	SetS01PID06_ShortFuelTrimBank1(float64) *OBD2
	// SetRawS01PID06_ShortFuelTrimBank1 sets the raw (encoded) value of the S01PID06_ShortFuelTrimBank1 signal.
	SetRawS01PID06_ShortFuelTrimBank1(uint8) *OBD2
	// SetS01PID07_LongFuelTrimBank1 sets the physical value of the S01PID07_LongFuelTrimBank1 signal.
	SetS01PID07_LongFuelTrimBank1(float64) *OBD2
	// SetRawS01PID07_LongFuelTrimBank1 sets the raw (encoded) value of the S01PID07_LongFuelTrimBank1 signal.
	SetRawS01PID07_LongFuelTrimBank1(uint8) *OBD2
	// SetS01PID08_ShortFuelTrimBank2 sets the physical value of the S01PID08_ShortFuelTrimBank2 signal.
	SetS01PID08_ShortFuelTrimBank2(float64) *OBD2
	// SetRawS01PID08_ShortFuelTrimBank2 sets the raw (encoded) value of the S01PID08_ShortFuelTrimBank2 signal.
	SetRawS01PID08_ShortFuelTrimBank2(uint8) *OBD2
	// SetS01PID09_LongFuelTrimBank2 sets the physical value of the S01PID09_LongFuelTrimBank2 signal.
	SetS01PID09_LongFuelTrimBank2(float64) *OBD2
	// SetRawS01PID09_LongFuelTrimBank2 sets the raw (encoded) value of the S01PID09_LongFuelTrimBank2 signal.
	SetRawS01PID09_LongFuelTrimBank2(uint8) *OBD2
	// SetS01PID0A_FuelPressure sets the physical value of the S01PID0A_FuelPressure signal.
	SetS01PID0A_FuelPressure(float64) *OBD2
	// SetRawS01PID0A_FuelPressure sets the raw (encoded) value of the S01PID0A_FuelPressure signal.
	SetRawS01PID0A_FuelPressure(uint8) *OBD2
	// SetS01PID0B_IntakeManiAbsPress sets the value of the S01PID0B_IntakeManiAbsPress signal.
	SetS01PID0B_IntakeManiAbsPress(uint8) *OBD2
	// SetS01PID0C_EngineRPM sets the physical value of the S01PID0C_EngineRPM signal.
	SetS01PID0C_EngineRPM(float64) *OBD2
	// SetRawS01PID0C_EngineRPM sets the raw (encoded) value of the S01PID0C_EngineRPM signal.
	SetRawS01PID0C_EngineRPM(uint16) *OBD2
	// SetS01PID0D_VehicleSpeed sets the value of the S01PID0D_VehicleSpeed signal.
	SetS01PID0D_VehicleSpeed(uint8) *OBD2
	// SetS01PID0E_TimingAdvance sets the physical value of the S01PID0E_TimingAdvance signal.
	SetS01PID0E_TimingAdvance(float64) *OBD2
	// SetRawS01PID0E_TimingAdvance sets the raw (encoded) value of the S01PID0E_TimingAdvance signal.
	SetRawS01PID0E_TimingAdvance(uint8) *OBD2
	// SetS01PID0F_IntakeAirTemperature sets the physical value of the S01PID0F_IntakeAirTemperature signal.
	SetS01PID0F_IntakeAirTemperature(float64) *OBD2
	// SetRawS01PID0F_IntakeAirTemperature sets the raw (encoded) value of the S01PID0F_IntakeAirTemperature signal.
	SetRawS01PID0F_IntakeAirTemperature(uint8) *OBD2
	// SetS01PID10_MAFAirFlowRate sets the physical value of the S01PID10_MAFAirFlowRate signal.
	SetS01PID10_MAFAirFlowRate(float64) *OBD2
	// SetRawS01PID10_MAFAirFlowRate sets the raw (encoded) value of the S01PID10_MAFAirFlowRate signal.
	SetRawS01PID10_MAFAirFlowRate(uint16) *OBD2
	// SetS01PID11_ThrottlePosition sets the physical value of the S01PID11_ThrottlePosition signal.
	SetS01PID11_ThrottlePosition(float64) *OBD2
	// SetRawS01PID11_ThrottlePosition sets the raw (encoded) value of the S01PID11_ThrottlePosition signal.
	SetRawS01PID11_ThrottlePosition(uint8) *OBD2
	// SetS01PID12_CmdSecAirStatus sets the value of the S01PID12_CmdSecAirStatus signal.
	SetS01PID12_CmdSecAirStatus(OBD2_S01PID12_CmdSecAirStatus) *OBD2
	// SetS01PID14_OxySensor1_Volt sets the physical value of the S01PID14_OxySensor1_Volt signal.
	SetS01PID14_OxySensor1_Volt(float64) *OBD2
	// SetRawS01PID14_OxySensor1_Volt sets the raw (encoded) value of the S01PID14_OxySensor1_Volt signal.
	SetRawS01PID14_OxySensor1_Volt(uint8) *OBD2
	// SetS01PID15_OxySensor2_Volt sets the physical value of the S01PID15_OxySensor2_Volt signal.
	SetS01PID15_OxySensor2_Volt(float64) *OBD2
	// SetRawS01PID15_OxySensor2_Volt sets the raw (encoded) value of the S01PID15_OxySensor2_Volt signal.
	SetRawS01PID15_OxySensor2_Volt(uint8) *OBD2
	// SetS01PID16_OxySensor3_Volt sets the physical value of the S01PID16_OxySensor3_Volt signal.
	SetS01PID16_OxySensor3_Volt(float64) *OBD2
	// SetRawS01PID16_OxySensor3_Volt sets the raw (encoded) value of the S01PID16_OxySensor3_Volt signal.
	SetRawS01PID16_OxySensor3_Volt(uint8) *OBD2
	// SetS01PID17_OxySensor4_Volt sets the physical value of the S01PID17_OxySensor4_Volt signal.
	SetS01PID17_OxySensor4_Volt(float64) *OBD2
	// SetRawS01PID17_OxySensor4_Volt sets the raw (encoded) value of the S01PID17_OxySensor4_Volt signal.
	SetRawS01PID17_OxySensor4_Volt(uint8) *OBD2
	// SetS01PID18_OxySensor5_Volt sets the physical value of the S01PID18_OxySensor5_Volt signal.
	SetS01PID18_OxySensor5_Volt(float64) *OBD2
	// SetRawS01PID18_OxySensor5_Volt sets the raw (encoded) value of the S01PID18_OxySensor5_Volt signal.
	SetRawS01PID18_OxySensor5_Volt(uint8) *OBD2
	// SetS01PID19_OxySensor6_Volt sets the physical value of the S01PID19_OxySensor6_Volt signal.
	SetS01PID19_OxySensor6_Volt(float64) *OBD2
	// SetRawS01PID19_OxySensor6_Volt sets the raw (encoded) value of the S01PID19_OxySensor6_Volt signal.
	SetRawS01PID19_OxySensor6_Volt(uint8) *OBD2
	// SetS01PID14_OxySensor1_STFT sets the physical value of the S01PID14_OxySensor1_STFT signal.
	SetS01PID14_OxySensor1_STFT(float64) *OBD2
	// SetRawS01PID14_OxySensor1_STFT sets the raw (encoded) value of the S01PID14_OxySensor1_STFT signal.
	SetRawS01PID14_OxySensor1_STFT(uint8) *OBD2
	// SetS01PID15_OxySensor2_STFT sets the physical value of the S01PID15_OxySensor2_STFT signal.
	SetS01PID15_OxySensor2_STFT(float64) *OBD2
	// SetRawS01PID15_OxySensor2_STFT sets the raw (encoded) value of the S01PID15_OxySensor2_STFT signal.
	SetRawS01PID15_OxySensor2_STFT(uint8) *OBD2
	// SetS01PID16_OxySensor3_STFT sets the physical value of the S01PID16_OxySensor3_STFT signal.
	SetS01PID16_OxySensor3_STFT(float64) *OBD2
	// SetRawS01PID16_OxySensor3_STFT sets the raw (encoded) value of the S01PID16_OxySensor3_STFT signal.
	SetRawS01PID16_OxySensor3_STFT(uint8) *OBD2
	// SetS01PID17_OxySensor4_STFT sets the physical value of the S01PID17_OxySensor4_STFT signal.
	SetS01PID17_OxySensor4_STFT(float64) *OBD2
	// SetRawS01PID17_OxySensor4_STFT sets the raw (encoded) value of the S01PID17_OxySensor4_STFT signal.
	SetRawS01PID17_OxySensor4_STFT(uint8) *OBD2
	// SetS01PID18_OxySensor5_STFT sets the physical value of the S01PID18_OxySensor5_STFT signal.
	SetS01PID18_OxySensor5_STFT(float64) *OBD2
	// SetRawS01PID18_OxySensor5_STFT sets the raw (encoded) value of the S01PID18_OxySensor5_STFT signal.
	SetRawS01PID18_OxySensor5_STFT(uint8) *OBD2
	// SetS01PID19_OxySensor6_STFT sets the physical value of the S01PID19_OxySensor6_STFT signal.
	SetS01PID19_OxySensor6_STFT(float64) *OBD2
	// SetRawS01PID19_OxySensor6_STFT sets the raw (encoded) value of the S01PID19_OxySensor6_STFT signal.
	SetRawS01PID19_OxySensor6_STFT(uint8) *OBD2
	// SetS01PID1A_OxySensor7_Volt sets the physical value of the S01PID1A_OxySensor7_Volt signal.
	SetS01PID1A_OxySensor7_Volt(float64) *OBD2
	// SetRawS01PID1A_OxySensor7_Volt sets the raw (encoded) value of the S01PID1A_OxySensor7_Volt signal.
	SetRawS01PID1A_OxySensor7_Volt(uint8) *OBD2
	// SetS01PID1A_OxySensor7_STFT sets the physical value of the S01PID1A_OxySensor7_STFT signal.
	SetS01PID1A_OxySensor7_STFT(float64) *OBD2
	// SetRawS01PID1A_OxySensor7_STFT sets the raw (encoded) value of the S01PID1A_OxySensor7_STFT signal.
	SetRawS01PID1A_OxySensor7_STFT(uint8) *OBD2
	// SetS01PID1B_OxySensor8_Volt sets the physical value of the S01PID1B_OxySensor8_Volt signal.
	SetS01PID1B_OxySensor8_Volt(float64) *OBD2
	// SetRawS01PID1B_OxySensor8_Volt sets the raw (encoded) value of the S01PID1B_OxySensor8_Volt signal.
	SetRawS01PID1B_OxySensor8_Volt(uint8) *OBD2
	// SetS01PID1B_OxySensor8_STFT sets the physical value of the S01PID1B_OxySensor8_STFT signal.
	SetS01PID1B_OxySensor8_STFT(float64) *OBD2
	// SetRawS01PID1B_OxySensor8_STFT sets the raw (encoded) value of the S01PID1B_OxySensor8_STFT signal.
	SetRawS01PID1B_OxySensor8_STFT(uint8) *OBD2
	// SetS01PID1C_OBDStandard sets the value of the S01PID1C_OBDStandard signal.
	SetS01PID1C_OBDStandard(OBD2_S01PID1C_OBDStandard) *OBD2
	// SetS01PID1F_TimeSinceEngStart sets the value of the S01PID1F_TimeSinceEngStart signal.
	SetS01PID1F_TimeSinceEngStart(uint16) *OBD2
	// SetS01PID20_PIDsSupported_21_40 sets the value of the S01PID20_PIDsSupported_21_40 signal.
	SetS01PID20_PIDsSupported_21_40(uint32) *OBD2
	// SetS01PID21_DistanceMILOn sets the value of the S01PID21_DistanceMILOn signal.
	SetS01PID21_DistanceMILOn(uint16) *OBD2
	// SetS01PID22_FuelRailPres sets the physical value of the S01PID22_FuelRailPres signal.
	SetS01PID22_FuelRailPres(float64) *OBD2
	// SetRawS01PID22_FuelRailPres sets the raw (encoded) value of the S01PID22_FuelRailPres signal.
	SetRawS01PID22_FuelRailPres(uint16) *OBD2
	// SetS01PID23_FuelRailGaug sets the physical value of the S01PID23_FuelRailGaug signal.
	SetS01PID23_FuelRailGaug(float64) *OBD2
	// SetRawS01PID23_FuelRailGaug sets the raw (encoded) value of the S01PID23_FuelRailGaug signal.
	SetRawS01PID23_FuelRailGaug(uint16) *OBD2
	// SetS01PID24_OxySensor1_FAER sets the physical value of the S01PID24_OxySensor1_FAER signal.
	SetS01PID24_OxySensor1_FAER(float64) *OBD2
	// SetRawS01PID24_OxySensor1_FAER sets the raw (encoded) value of the S01PID24_OxySensor1_FAER signal.
	SetRawS01PID24_OxySensor1_FAER(uint16) *OBD2
	// SetS01PID24_OxySensor1_Volt sets the physical value of the S01PID24_OxySensor1_Volt signal.
	SetS01PID24_OxySensor1_Volt(float64) *OBD2
	// SetRawS01PID24_OxySensor1_Volt sets the raw (encoded) value of the S01PID24_OxySensor1_Volt signal.
	SetRawS01PID24_OxySensor1_Volt(uint16) *OBD2
	// SetS01PID25_OxySensor2_FAER sets the physical value of the S01PID25_OxySensor2_FAER signal.
	SetS01PID25_OxySensor2_FAER(float64) *OBD2
	// SetRawS01PID25_OxySensor2_FAER sets the raw (encoded) value of the S01PID25_OxySensor2_FAER signal.
	SetRawS01PID25_OxySensor2_FAER(uint16) *OBD2
	// SetS01PID25_OxySensor2_Volt sets the physical value of the S01PID25_OxySensor2_Volt signal.
	SetS01PID25_OxySensor2_Volt(float64) *OBD2
	// SetRawS01PID25_OxySensor2_Volt sets the raw (encoded) value of the S01PID25_OxySensor2_Volt signal.
	SetRawS01PID25_OxySensor2_Volt(uint16) *OBD2
	// SetS01PID26_OxySensor3_FAER sets the physical value of the S01PID26_OxySensor3_FAER signal.
	SetS01PID26_OxySensor3_FAER(float64) *OBD2
	// SetRawS01PID26_OxySensor3_FAER sets the raw (encoded) value of the S01PID26_OxySensor3_FAER signal.
	SetRawS01PID26_OxySensor3_FAER(uint16) *OBD2
	// SetS01PID26_OxySensor3_Volt sets the physical value of the S01PID26_OxySensor3_Volt signal.
	SetS01PID26_OxySensor3_Volt(float64) *OBD2
	// SetRawS01PID26_OxySensor3_Volt sets the raw (encoded) value of the S01PID26_OxySensor3_Volt signal.
	SetRawS01PID26_OxySensor3_Volt(uint16) *OBD2
	// SetS01PID27_OxySensor4_FAER sets the physical value of the S01PID27_OxySensor4_FAER signal.
	SetS01PID27_OxySensor4_FAER(float64) *OBD2
	// SetRawS01PID27_OxySensor4_FAER sets the raw (encoded) value of the S01PID27_OxySensor4_FAER signal.
	SetRawS01PID27_OxySensor4_FAER(uint16) *OBD2
	// SetS01PID28_OxySensor5_FAER sets the physical value of the S01PID28_OxySensor5_FAER signal.
	SetS01PID28_OxySensor5_FAER(float64) *OBD2
	// SetRawS01PID28_OxySensor5_FAER sets the raw (encoded) value of the S01PID28_OxySensor5_FAER signal.
	SetRawS01PID28_OxySensor5_FAER(uint16) *OBD2
	// SetS01PID29_OxySensor6_FAER sets the physical value of the S01PID29_OxySensor6_FAER signal.
	SetS01PID29_OxySensor6_FAER(float64) *OBD2
	// SetRawS01PID29_OxySensor6_FAER sets the raw (encoded) value of the S01PID29_OxySensor6_FAER signal.
	SetRawS01PID29_OxySensor6_FAER(uint16) *OBD2
	// SetS01PID27_OxySensor4_Volt sets the physical value of the S01PID27_OxySensor4_Volt signal.
	SetS01PID27_OxySensor4_Volt(float64) *OBD2
	// SetRawS01PID27_OxySensor4_Volt sets the raw (encoded) value of the S01PID27_OxySensor4_Volt signal.
	SetRawS01PID27_OxySensor4_Volt(uint16) *OBD2
	// SetS01PID28_OxySensor5_Volt sets the physical value of the S01PID28_OxySensor5_Volt signal.
	SetS01PID28_OxySensor5_Volt(float64) *OBD2
	// SetRawS01PID28_OxySensor5_Volt sets the raw (encoded) value of the S01PID28_OxySensor5_Volt signal.
	SetRawS01PID28_OxySensor5_Volt(uint16) *OBD2
	// SetS01PID29_OxySensor6_Volt sets the physical value of the S01PID29_OxySensor6_Volt signal.
	SetS01PID29_OxySensor6_Volt(float64) *OBD2
	// SetRawS01PID29_OxySensor6_Volt sets the raw (encoded) value of the S01PID29_OxySensor6_Volt signal.
	SetRawS01PID29_OxySensor6_Volt(uint16) *OBD2
	// SetS01PID2A_OxySensor7_FAER sets the physical value of the S01PID2A_OxySensor7_FAER signal.
	SetS01PID2A_OxySensor7_FAER(float64) *OBD2
	// SetRawS01PID2A_OxySensor7_FAER sets the raw (encoded) value of the S01PID2A_OxySensor7_FAER signal.
	SetRawS01PID2A_OxySensor7_FAER(uint16) *OBD2
	// SetS01PID2A_OxySensor7_Volt sets the physical value of the S01PID2A_OxySensor7_Volt signal.
	SetS01PID2A_OxySensor7_Volt(float64) *OBD2
	// SetRawS01PID2A_OxySensor7_Volt sets the raw (encoded) value of the S01PID2A_OxySensor7_Volt signal.
	SetRawS01PID2A_OxySensor7_Volt(uint16) *OBD2
	// SetS01PID2B_OxySensor8_FAER sets the physical value of the S01PID2B_OxySensor8_FAER signal.
	SetS01PID2B_OxySensor8_FAER(float64) *OBD2
	// SetRawS01PID2B_OxySensor8_FAER sets the raw (encoded) value of the S01PID2B_OxySensor8_FAER signal.
	SetRawS01PID2B_OxySensor8_FAER(uint16) *OBD2
	// SetS01PID2B_OxySensor8_Volt sets the physical value of the S01PID2B_OxySensor8_Volt signal.
	SetS01PID2B_OxySensor8_Volt(float64) *OBD2
	// SetRawS01PID2B_OxySensor8_Volt sets the raw (encoded) value of the S01PID2B_OxySensor8_Volt signal.
	SetRawS01PID2B_OxySensor8_Volt(uint16) *OBD2
	// SetS01PID2C_CmdEGR sets the physical value of the S01PID2C_CmdEGR signal.
	SetS01PID2C_CmdEGR(float64) *OBD2
	// SetRawS01PID2C_CmdEGR sets the raw (encoded) value of the S01PID2C_CmdEGR signal.
	SetRawS01PID2C_CmdEGR(uint8) *OBD2
	// SetS01PID2D_EGRError sets the physical value of the S01PID2D_EGRError signal.
	SetS01PID2D_EGRError(float64) *OBD2
	// SetRawS01PID2D_EGRError sets the raw (encoded) value of the S01PID2D_EGRError signal.
	SetRawS01PID2D_EGRError(uint8) *OBD2
	// SetS01PID2E_CmdEvapPurge sets the physical value of the S01PID2E_CmdEvapPurge signal.
	SetS01PID2E_CmdEvapPurge(float64) *OBD2
	// SetRawS01PID2E_CmdEvapPurge sets the raw (encoded) value of the S01PID2E_CmdEvapPurge signal.
	SetRawS01PID2E_CmdEvapPurge(uint8) *OBD2
	// SetS01PID2F_FuelTankLevel sets the physical value of the S01PID2F_FuelTankLevel signal.
	SetS01PID2F_FuelTankLevel(float64) *OBD2
	// SetRawS01PID2F_FuelTankLevel sets the raw (encoded) value of the S01PID2F_FuelTankLevel signal.
	SetRawS01PID2F_FuelTankLevel(uint8) *OBD2
	// SetS01PID30_WarmUpsSinceCodeClear sets the value of the S01PID30_WarmUpsSinceCodeClear signal.
	SetS01PID30_WarmUpsSinceCodeClear(uint8) *OBD2
	// SetS01PID31_DistanceSinceCodeClear sets the value of the S01PID31_DistanceSinceCodeClear signal.
	SetS01PID31_DistanceSinceCodeClear(uint16) *OBD2
	// SetS01PID32_EvapSysVaporPres sets the physical value of the S01PID32_EvapSysVaporPres signal.
	SetS01PID32_EvapSysVaporPres(float64) *OBD2
	// SetRawS01PID32_EvapSysVaporPres sets the raw (encoded) value of the S01PID32_EvapSysVaporPres signal.
	SetRawS01PID32_EvapSysVaporPres(int16) *OBD2
	// SetS01PID33_AbsBaroPres sets the value of the S01PID33_AbsBaroPres signal.
	SetS01PID33_AbsBaroPres(uint8) *OBD2
	// SetS01PID34_OxySensor1_FAER sets the physical value of the S01PID34_OxySensor1_FAER signal.
	SetS01PID34_OxySensor1_FAER(float64) *OBD2
	// SetRawS01PID34_OxySensor1_FAER sets the raw (encoded) value of the S01PID34_OxySensor1_FAER signal.
	SetRawS01PID34_OxySensor1_FAER(uint16) *OBD2
	// SetS01PID34_OxySensor1_Crnt sets the physical value of the S01PID34_OxySensor1_Crnt signal.
	SetS01PID34_OxySensor1_Crnt(float64) *OBD2
	// SetRawS01PID34_OxySensor1_Crnt sets the raw (encoded) value of the S01PID34_OxySensor1_Crnt signal.
	SetRawS01PID34_OxySensor1_Crnt(uint16) *OBD2
	// SetS01PID35_OxySensor2_FAER sets the physical value of the S01PID35_OxySensor2_FAER signal.
	SetS01PID35_OxySensor2_FAER(float64) *OBD2
	// SetRawS01PID35_OxySensor2_FAER sets the raw (encoded) value of the S01PID35_OxySensor2_FAER signal.
	SetRawS01PID35_OxySensor2_FAER(uint16) *OBD2
	// SetS01PID35_OxySensor2_Crnt sets the physical value of the S01PID35_OxySensor2_Crnt signal.
	SetS01PID35_OxySensor2_Crnt(float64) *OBD2
	// SetRawS01PID35_OxySensor2_Crnt sets the raw (encoded) value of the S01PID35_OxySensor2_Crnt signal.
	SetRawS01PID35_OxySensor2_Crnt(uint16) *OBD2
	// SetS01PID36_OxySensor3_FAER sets the physical value of the S01PID36_OxySensor3_FAER signal.
	SetS01PID36_OxySensor3_FAER(float64) *OBD2
	// SetRawS01PID36_OxySensor3_FAER sets the raw (encoded) value of the S01PID36_OxySensor3_FAER signal.
	SetRawS01PID36_OxySensor3_FAER(uint16) *OBD2
	// SetS01PID36_OxySensor3_Crnt sets the physical value of the S01PID36_OxySensor3_Crnt signal.
	SetS01PID36_OxySensor3_Crnt(float64) *OBD2
	// SetRawS01PID36_OxySensor3_Crnt sets the raw (encoded) value of the S01PID36_OxySensor3_Crnt signal.
	SetRawS01PID36_OxySensor3_Crnt(uint16) *OBD2
	// SetS01PID37_OxySensor4_FAER sets the physical value of the S01PID37_OxySensor4_FAER signal.
	SetS01PID37_OxySensor4_FAER(float64) *OBD2
	// SetRawS01PID37_OxySensor4_FAER sets the raw (encoded) value of the S01PID37_OxySensor4_FAER signal.
	SetRawS01PID37_OxySensor4_FAER(uint16) *OBD2
	// SetS01PID38_OxySensor5_FAER sets the physical value of the S01PID38_OxySensor5_FAER signal.
	SetS01PID38_OxySensor5_FAER(float64) *OBD2
	// SetRawS01PID38_OxySensor5_FAER sets the raw (encoded) value of the S01PID38_OxySensor5_FAER signal.
	SetRawS01PID38_OxySensor5_FAER(uint16) *OBD2
	// SetS01PID39_OxySensor6_FAER sets the physical value of the S01PID39_OxySensor6_FAER signal.
	SetS01PID39_OxySensor6_FAER(float64) *OBD2
	// SetRawS01PID39_OxySensor6_FAER sets the raw (encoded) value of the S01PID39_OxySensor6_FAER signal.
	SetRawS01PID39_OxySensor6_FAER(uint16) *OBD2
	// SetS01PID37_OxySensor4_Crnt sets the physical value of the S01PID37_OxySensor4_Crnt signal.
	SetS01PID37_OxySensor4_Crnt(float64) *OBD2
	// SetRawS01PID37_OxySensor4_Crnt sets the raw (encoded) value of the S01PID37_OxySensor4_Crnt signal.
	SetRawS01PID37_OxySensor4_Crnt(uint16) *OBD2
	// SetS01PID38_OxySensor5_Crnt sets the physical value of the S01PID38_OxySensor5_Crnt signal.
	SetS01PID38_OxySensor5_Crnt(float64) *OBD2
	// SetRawS01PID38_OxySensor5_Crnt sets the raw (encoded) value of the S01PID38_OxySensor5_Crnt signal.
	SetRawS01PID38_OxySensor5_Crnt(uint16) *OBD2
	// SetS01PID39_OxySensor6_Crnt sets the physical value of the S01PID39_OxySensor6_Crnt signal.
	SetS01PID39_OxySensor6_Crnt(float64) *OBD2
	// SetRawS01PID39_OxySensor6_Crnt sets the raw (encoded) value of the S01PID39_OxySensor6_Crnt signal.
	SetRawS01PID39_OxySensor6_Crnt(uint16) *OBD2
	// SetS01PID3A_OxySensor7_FAER sets the physical value of the S01PID3A_OxySensor7_FAER signal.
	SetS01PID3A_OxySensor7_FAER(float64) *OBD2
	// SetRawS01PID3A_OxySensor7_FAER sets the raw (encoded) value of the S01PID3A_OxySensor7_FAER signal.
	SetRawS01PID3A_OxySensor7_FAER(uint16) *OBD2
	// SetS01PID3B_OxySensor8_FAER sets the physical value of the S01PID3B_OxySensor8_FAER signal.
	SetS01PID3B_OxySensor8_FAER(float64) *OBD2
	// SetRawS01PID3B_OxySensor8_FAER sets the raw (encoded) value of the S01PID3B_OxySensor8_FAER signal.
	SetRawS01PID3B_OxySensor8_FAER(uint16) *OBD2
	// SetS01PID3C_CatTempBank1Sens1 sets the physical value of the S01PID3C_CatTempBank1Sens1 signal.
	SetS01PID3C_CatTempBank1Sens1(float64) *OBD2
	// SetRawS01PID3C_CatTempBank1Sens1 sets the raw (encoded) value of the S01PID3C_CatTempBank1Sens1 signal.
	SetRawS01PID3C_CatTempBank1Sens1(uint16) *OBD2
	// SetS01PID3D_CatTempBank2Sens1 sets the physical value of the S01PID3D_CatTempBank2Sens1 signal.
	SetS01PID3D_CatTempBank2Sens1(float64) *OBD2
	// SetRawS01PID3D_CatTempBank2Sens1 sets the raw (encoded) value of the S01PID3D_CatTempBank2Sens1 signal.
	SetRawS01PID3D_CatTempBank2Sens1(uint16) *OBD2
	// SetS01PID3A_OxySensor7_Crnt sets the physical value of the S01PID3A_OxySensor7_Crnt signal.
	SetS01PID3A_OxySensor7_Crnt(float64) *OBD2
	// SetRawS01PID3A_OxySensor7_Crnt sets the raw (encoded) value of the S01PID3A_OxySensor7_Crnt signal.
	SetRawS01PID3A_OxySensor7_Crnt(uint16) *OBD2
	// SetS01PID3B_OxySensor8_Crnt sets the physical value of the S01PID3B_OxySensor8_Crnt signal.
	SetS01PID3B_OxySensor8_Crnt(float64) *OBD2
	// SetRawS01PID3B_OxySensor8_Crnt sets the raw (encoded) value of the S01PID3B_OxySensor8_Crnt signal.
	SetRawS01PID3B_OxySensor8_Crnt(uint16) *OBD2
	// SetS01PID3E_CatTempBank1Sens2 sets the physical value of the S01PID3E_CatTempBank1Sens2 signal.
	SetS01PID3E_CatTempBank1Sens2(float64) *OBD2
	// SetRawS01PID3E_CatTempBank1Sens2 sets the raw (encoded) value of the S01PID3E_CatTempBank1Sens2 signal.
	SetRawS01PID3E_CatTempBank1Sens2(uint16) *OBD2
	// SetS01PID3F_CatTempBank2Sens2 sets the physical value of the S01PID3F_CatTempBank2Sens2 signal.
	SetS01PID3F_CatTempBank2Sens2(float64) *OBD2
	// SetRawS01PID3F_CatTempBank2Sens2 sets the raw (encoded) value of the S01PID3F_CatTempBank2Sens2 signal.
	SetRawS01PID3F_CatTempBank2Sens2(uint16) *OBD2
	// SetS01PID40_PIDsSupported_41_60 sets the value of the S01PID40_PIDsSupported_41_60 signal.
	SetS01PID40_PIDsSupported_41_60(uint32) *OBD2
	// SetS01PID41_MonStatusDriveCycle sets the value of the S01PID41_MonStatusDriveCycle signal.
	SetS01PID41_MonStatusDriveCycle(uint32) *OBD2
	// SetS01PID42_ControlModuleVolt sets the physical value of the S01PID42_ControlModuleVolt signal.
	SetS01PID42_ControlModuleVolt(float64) *OBD2
	// SetRawS01PID42_ControlModuleVolt sets the raw (encoded) value of the S01PID42_ControlModuleVolt signal.
	SetRawS01PID42_ControlModuleVolt(uint16) *OBD2
	// SetS01PID43_AbsLoadValue sets the physical value of the S01PID43_AbsLoadValue signal.
	SetS01PID43_AbsLoadValue(float64) *OBD2
	// SetRawS01PID43_AbsLoadValue sets the raw (encoded) value of the S01PID43_AbsLoadValue signal.
	SetRawS01PID43_AbsLoadValue(uint16) *OBD2
	// SetS01PID44_FuelAirCmdEquiv sets the physical value of the S01PID44_FuelAirCmdEquiv signal.
	SetS01PID44_FuelAirCmdEquiv(float64) *OBD2
	// SetRawS01PID44_FuelAirCmdEquiv sets the raw (encoded) value of the S01PID44_FuelAirCmdEquiv signal.
	SetRawS01PID44_FuelAirCmdEquiv(uint16) *OBD2
	// SetS01PID45_RelThrottlePos sets the physical value of the S01PID45_RelThrottlePos signal.
	SetS01PID45_RelThrottlePos(float64) *OBD2
	// SetRawS01PID45_RelThrottlePos sets the raw (encoded) value of the S01PID45_RelThrottlePos signal.
	SetRawS01PID45_RelThrottlePos(uint8) *OBD2
	// SetS01PID46_AmbientAirTemp sets the physical value of the S01PID46_AmbientAirTemp signal.
	SetS01PID46_AmbientAirTemp(float64) *OBD2
	// SetRawS01PID46_AmbientAirTemp sets the raw (encoded) value of the S01PID46_AmbientAirTemp signal.
	SetRawS01PID46_AmbientAirTemp(uint8) *OBD2
	// SetS01PID47_AbsThrottlePosB sets the physical value of the S01PID47_AbsThrottlePosB signal.
	SetS01PID47_AbsThrottlePosB(float64) *OBD2
	// SetRawS01PID47_AbsThrottlePosB sets the raw (encoded) value of the S01PID47_AbsThrottlePosB signal.
	SetRawS01PID47_AbsThrottlePosB(uint8) *OBD2
	// SetS01PID48_AbsThrottlePosC sets the physical value of the S01PID48_AbsThrottlePosC signal.
	SetS01PID48_AbsThrottlePosC(float64) *OBD2
	// SetRawS01PID48_AbsThrottlePosC sets the raw (encoded) value of the S01PID48_AbsThrottlePosC signal.
	SetRawS01PID48_AbsThrottlePosC(uint8) *OBD2
	// SetS01PID49_AbsThrottlePosD sets the physical value of the S01PID49_AbsThrottlePosD signal.
	SetS01PID49_AbsThrottlePosD(float64) *OBD2
	// SetRawS01PID49_AbsThrottlePosD sets the raw (encoded) value of the S01PID49_AbsThrottlePosD signal.
	SetRawS01PID49_AbsThrottlePosD(uint8) *OBD2
	// SetS01PID4A_AbsThrottlePosE sets the physical value of the S01PID4A_AbsThrottlePosE signal.
	SetS01PID4A_AbsThrottlePosE(float64) *OBD2
	// SetRawS01PID4A_AbsThrottlePosE sets the raw (encoded) value of the S01PID4A_AbsThrottlePosE signal.
	SetRawS01PID4A_AbsThrottlePosE(uint8) *OBD2
	// SetS01PID4B_AbsThrottlePosF sets the physical value of the S01PID4B_AbsThrottlePosF signal.
	SetS01PID4B_AbsThrottlePosF(float64) *OBD2
	// SetRawS01PID4B_AbsThrottlePosF sets the raw (encoded) value of the S01PID4B_AbsThrottlePosF signal.
	SetRawS01PID4B_AbsThrottlePosF(uint8) *OBD2
	// SetS01PID4C_CmdThrottleAct sets the physical value of the S01PID4C_CmdThrottleAct signal.
	SetS01PID4C_CmdThrottleAct(float64) *OBD2
	// SetRawS01PID4C_CmdThrottleAct sets the raw (encoded) value of the S01PID4C_CmdThrottleAct signal.
	SetRawS01PID4C_CmdThrottleAct(uint8) *OBD2
	// SetS01PID4D_TimeRunMILOn sets the value of the S01PID4D_TimeRunMILOn signal.
	SetS01PID4D_TimeRunMILOn(uint16) *OBD2
	// SetS01PID4E_TimeSinceCodeClear sets the value of the S01PID4E_TimeSinceCodeClear signal.
	SetS01PID4E_TimeSinceCodeClear(uint16) *OBD2
	// SetS01PID4F_Max_FAER sets the value of the S01PID4F_Max_FAER signal.
	SetS01PID4F_Max_FAER(uint8) *OBD2
	// SetS01PID4F_Max_OxySensVol sets the value of the S01PID4F_Max_OxySensVol signal.
	SetS01PID4F_Max_OxySensVol(uint8) *OBD2
	// SetS01PID4F_Max_OxySensCrnt sets the value of the S01PID4F_Max_OxySensCrnt signal.
	SetS01PID4F_Max_OxySensCrnt(uint8) *OBD2
	// SetS01PID4F_Max_IntManiAbsPres sets the physical value of the S01PID4F_Max_IntManiAbsPres signal.
	SetS01PID4F_Max_IntManiAbsPres(float64) *OBD2
	// SetRawS01PID4F_Max_IntManiAbsPres sets the raw (encoded) value of the S01PID4F_Max_IntManiAbsPres signal.
	SetRawS01PID4F_Max_IntManiAbsPres(uint8) *OBD2
	// SetS01PID50_Max_AirFlowMAF sets the physical value of the S01PID50_Max_AirFlowMAF signal.
	SetS01PID50_Max_AirFlowMAF(float64) *OBD2
	// SetRawS01PID50_Max_AirFlowMAF sets the raw (encoded) value of the S01PID50_Max_AirFlowMAF signal.
	SetRawS01PID50_Max_AirFlowMAF(uint8) *OBD2
	// SetS01PID51_FuelType sets the value of the S01PID51_FuelType signal.
	SetS01PID51_FuelType(OBD2_S01PID51_FuelType) *OBD2
	// SetS01PID52_EthanolFuelPct sets the physical value of the S01PID52_EthanolFuelPct signal.
	SetS01PID52_EthanolFuelPct(float64) *OBD2
	// SetRawS01PID52_EthanolFuelPct sets the raw (encoded) value of the S01PID52_EthanolFuelPct signal.
	SetRawS01PID52_EthanolFuelPct(uint8) *OBD2
	// SetS01PID53_AbsEvapSysVapPres sets the physical value of the S01PID53_AbsEvapSysVapPres signal.
	SetS01PID53_AbsEvapSysVapPres(float64) *OBD2
	// SetRawS01PID53_AbsEvapSysVapPres sets the raw (encoded) value of the S01PID53_AbsEvapSysVapPres signal.
	SetRawS01PID53_AbsEvapSysVapPres(uint16) *OBD2
	// SetS01PID54_EvapSysVapPres sets the physical value of the S01PID54_EvapSysVapPres signal.
	SetS01PID54_EvapSysVapPres(float64) *OBD2
	// SetRawS01PID54_EvapSysVapPres sets the raw (encoded) value of the S01PID54_EvapSysVapPres signal.
	SetRawS01PID54_EvapSysVapPres(uint16) *OBD2
	// SetS01PID55_ShortSecOxyTrimBank1 sets the physical value of the S01PID55_ShortSecOxyTrimBank1 signal.
	SetS01PID55_ShortSecOxyTrimBank1(float64) *OBD2
	// SetRawS01PID55_ShortSecOxyTrimBank1 sets the raw (encoded) value of the S01PID55_ShortSecOxyTrimBank1 signal.
	SetRawS01PID55_ShortSecOxyTrimBank1(uint8) *OBD2
	// SetS01PID56_LongSecOxyTrimBank1 sets the physical value of the S01PID56_LongSecOxyTrimBank1 signal.
	SetS01PID56_LongSecOxyTrimBank1(float64) *OBD2
	// SetRawS01PID56_LongSecOxyTrimBank1 sets the raw (encoded) value of the S01PID56_LongSecOxyTrimBank1 signal.
	SetRawS01PID56_LongSecOxyTrimBank1(uint8) *OBD2
	// SetS01PID55_ShortSecOxyTrimBank3 sets the physical value of the S01PID55_ShortSecOxyTrimBank3 signal.
	SetS01PID55_ShortSecOxyTrimBank3(float64) *OBD2
	// SetRawS01PID55_ShortSecOxyTrimBank3 sets the raw (encoded) value of the S01PID55_ShortSecOxyTrimBank3 signal.
	SetRawS01PID55_ShortSecOxyTrimBank3(uint8) *OBD2
	// SetS01PID56_LongSecOxyTrimBank3 sets the physical value of the S01PID56_LongSecOxyTrimBank3 signal.
	SetS01PID56_LongSecOxyTrimBank3(float64) *OBD2
	// SetRawS01PID56_LongSecOxyTrimBank3 sets the raw (encoded) value of the S01PID56_LongSecOxyTrimBank3 signal.
	SetRawS01PID56_LongSecOxyTrimBank3(uint8) *OBD2
	// SetS01PID57_ShortSecOxyTrimBank2 sets the physical value of the S01PID57_ShortSecOxyTrimBank2 signal.
	SetS01PID57_ShortSecOxyTrimBank2(float64) *OBD2
	// SetRawS01PID57_ShortSecOxyTrimBank2 sets the raw (encoded) value of the S01PID57_ShortSecOxyTrimBank2 signal.
	SetRawS01PID57_ShortSecOxyTrimBank2(uint8) *OBD2
	// SetS01PID58_LongSecOxyTrimBank2 sets the physical value of the S01PID58_LongSecOxyTrimBank2 signal.
	SetS01PID58_LongSecOxyTrimBank2(float64) *OBD2
	// SetRawS01PID58_LongSecOxyTrimBank2 sets the raw (encoded) value of the S01PID58_LongSecOxyTrimBank2 signal.
	SetRawS01PID58_LongSecOxyTrimBank2(uint8) *OBD2
	// SetS01PID59_FuelRailAbsPres sets the physical value of the S01PID59_FuelRailAbsPres signal.
	SetS01PID59_FuelRailAbsPres(float64) *OBD2
	// SetRawS01PID59_FuelRailAbsPres sets the raw (encoded) value of the S01PID59_FuelRailAbsPres signal.
	SetRawS01PID59_FuelRailAbsPres(uint16) *OBD2
	// SetS01PID5A_RelAccelPedalPos sets the physical value of the S01PID5A_RelAccelPedalPos signal.
	SetS01PID5A_RelAccelPedalPos(float64) *OBD2
	// SetRawS01PID5A_RelAccelPedalPos sets the raw (encoded) value of the S01PID5A_RelAccelPedalPos signal.
	SetRawS01PID5A_RelAccelPedalPos(uint8) *OBD2
	// SetS01PID57_ShortSecOxyTrimBank4 sets the physical value of the S01PID57_ShortSecOxyTrimBank4 signal.
	SetS01PID57_ShortSecOxyTrimBank4(float64) *OBD2
	// SetRawS01PID57_ShortSecOxyTrimBank4 sets the raw (encoded) value of the S01PID57_ShortSecOxyTrimBank4 signal.
	SetRawS01PID57_ShortSecOxyTrimBank4(uint8) *OBD2
	// SetS01PID58_LongSecOxyTrimBank4 sets the physical value of the S01PID58_LongSecOxyTrimBank4 signal.
	SetS01PID58_LongSecOxyTrimBank4(float64) *OBD2
	// SetRawS01PID58_LongSecOxyTrimBank4 sets the raw (encoded) value of the S01PID58_LongSecOxyTrimBank4 signal.
	SetRawS01PID58_LongSecOxyTrimBank4(uint8) *OBD2
	// SetS01PID5B_HybrBatPackRemLife sets the physical value of the S01PID5B_HybrBatPackRemLife signal.
	SetS01PID5B_HybrBatPackRemLife(float64) *OBD2
	// SetRawS01PID5B_HybrBatPackRemLife sets the raw (encoded) value of the S01PID5B_HybrBatPackRemLife signal.
	SetRawS01PID5B_HybrBatPackRemLife(uint8) *OBD2
	// SetS01PID5C_EngineOilTemp sets the physical value of the S01PID5C_EngineOilTemp signal.
	SetS01PID5C_EngineOilTemp(float64) *OBD2
	// SetRawS01PID5C_EngineOilTemp sets the raw (encoded) value of the S01PID5C_EngineOilTemp signal.
	SetRawS01PID5C_EngineOilTemp(uint8) *OBD2
	// SetS01PID5D_FuelInjectionTiming sets the physical value of the S01PID5D_FuelInjectionTiming signal.
	SetS01PID5D_FuelInjectionTiming(float64) *OBD2
	// SetRawS01PID5D_FuelInjectionTiming sets the raw (encoded) value of the S01PID5D_FuelInjectionTiming signal.
	SetRawS01PID5D_FuelInjectionTiming(uint16) *OBD2
	// SetS01PID5E_EngineFuelRate sets the physical value of the S01PID5E_EngineFuelRate signal.
	SetS01PID5E_EngineFuelRate(float64) *OBD2
	// SetRawS01PID5E_EngineFuelRate sets the raw (encoded) value of the S01PID5E_EngineFuelRate signal.
	SetRawS01PID5E_EngineFuelRate(uint16) *OBD2
	// SetS01PID5F_EmissionReq sets the value of the S01PID5F_EmissionReq signal.
	SetS01PID5F_EmissionReq(uint8) *OBD2
	// SetS01PID60_PIDsSupported_61_80 sets the value of the S01PID60_PIDsSupported_61_80 signal.
	SetS01PID60_PIDsSupported_61_80(uint32) *OBD2
	// SetS01PID61_DemandEngTorqPct sets the physical value of the S01PID61_DemandEngTorqPct signal.
	SetS01PID61_DemandEngTorqPct(float64) *OBD2
	// SetRawS01PID61_DemandEngTorqPct sets the raw (encoded) value of the S01PID61_DemandEngTorqPct signal.
	SetRawS01PID61_DemandEngTorqPct(uint8) *OBD2
	// SetS01PID62_ActualEngTorqPct sets the physical value of the S01PID62_ActualEngTorqPct signal.
	SetS01PID62_ActualEngTorqPct(float64) *OBD2
	// SetRawS01PID62_ActualEngTorqPct sets the raw (encoded) value of the S01PID62_ActualEngTorqPct signal.
	SetRawS01PID62_ActualEngTorqPct(uint8) *OBD2
	// SetS01PID63_EngRefTorq sets the value of the S01PID63_EngRefTorq signal.
	SetS01PID63_EngRefTorq(uint16) *OBD2
	// SetS01PID64_EngPctTorq_Idle sets the physical value of the S01PID64_EngPctTorq_Idle signal.
	SetS01PID64_EngPctTorq_Idle(float64) *OBD2
	// SetRawS01PID64_EngPctTorq_Idle sets the raw (encoded) value of the S01PID64_EngPctTorq_Idle signal.
	SetRawS01PID64_EngPctTorq_Idle(uint8) *OBD2
	// SetS01PID64_EngPctTorq_EP1 sets the physical value of the S01PID64_EngPctTorq_EP1 signal.
	SetS01PID64_EngPctTorq_EP1(float64) *OBD2
	// SetRawS01PID64_EngPctTorq_EP1 sets the raw (encoded) value of the S01PID64_EngPctTorq_EP1 signal.
	SetRawS01PID64_EngPctTorq_EP1(uint8) *OBD2
	// SetS01PID64_EngPctTorq_EP2 sets the physical value of the S01PID64_EngPctTorq_EP2 signal.
	SetS01PID64_EngPctTorq_EP2(float64) *OBD2
	// SetRawS01PID64_EngPctTorq_EP2 sets the raw (encoded) value of the S01PID64_EngPctTorq_EP2 signal.
	SetRawS01PID64_EngPctTorq_EP2(uint8) *OBD2
	// SetS01PID64_EngPctTorq_EP3 sets the physical value of the S01PID64_EngPctTorq_EP3 signal.
	SetS01PID64_EngPctTorq_EP3(float64) *OBD2
	// SetRawS01PID64_EngPctTorq_EP3 sets the raw (encoded) value of the S01PID64_EngPctTorq_EP3 signal.
	SetRawS01PID64_EngPctTorq_EP3(uint8) *OBD2
	// SetS01PID65_AuxInputOutput sets the value of the S01PID65_AuxInputOutput signal.
	SetS01PID65_AuxInputOutput(uint8) *OBD2
	// SetS01PID66_MAFSensor sets the value of the S01PID66_MAFSensor signal.
	SetS01PID66_MAFSensor(uint8) *OBD2
	// SetS01PID64_EngPctTorq_EP4 sets the physical value of the S01PID64_EngPctTorq_EP4 signal.
	SetS01PID64_EngPctTorq_EP4(float64) *OBD2
	// SetRawS01PID64_EngPctTorq_EP4 sets the raw (encoded) value of the S01PID64_EngPctTorq_EP4 signal.
	SetRawS01PID64_EngPctTorq_EP4(uint8) *OBD2
	// SetS01PID67_EngineCoolantTemp sets the value of the S01PID67_EngineCoolantTemp signal.
	SetS01PID67_EngineCoolantTemp(uint8) *OBD2
	// SetS01PID68_IntakeAirTempSens sets the value of the S01PID68_IntakeAirTempSens signal.
	SetS01PID68_IntakeAirTempSens(uint8) *OBD2
	// SetS01PID69_CmdEGR_EGRError sets the value of the S01PID69_CmdEGR_EGRError signal.
	SetS01PID69_CmdEGR_EGRError(uint8) *OBD2
	// SetS01PID6A_CmdDieselIntAir sets the value of the S01PID6A_CmdDieselIntAir signal.
	SetS01PID6A_CmdDieselIntAir(uint8) *OBD2
	// SetS01PID6B_ExhaustGasTemp sets the value of the S01PID6B_ExhaustGasTemp signal.
	SetS01PID6B_ExhaustGasTemp(uint8) *OBD2
	// SetS01PID6C_CmdThrottleActRel sets the value of the S01PID6C_CmdThrottleActRel signal.
	SetS01PID6C_CmdThrottleActRel(uint8) *OBD2
	// SetS01PID6D_FuelPresContrSys sets the value of the S01PID6D_FuelPresContrSys signal.
	SetS01PID6D_FuelPresContrSys(uint8) *OBD2
	// SetS01PID6E_InjPresContrSys sets the value of the S01PID6E_InjPresContrSys signal.
	SetS01PID6E_InjPresContrSys(uint8) *OBD2
	// SetS01PID6F_TurboComprPres sets the value of the S01PID6F_TurboComprPres signal.
	SetS01PID6F_TurboComprPres(uint8) *OBD2
	// SetS01PID70_BoostPresCntrl sets the value of the S01PID70_BoostPresCntrl signal.
	SetS01PID70_BoostPresCntrl(uint8) *OBD2
	// SetS01PID80_PIDsSupported_81_A0 sets the value of the S01PID80_PIDsSupported_81_A0 signal.
	SetS01PID80_PIDsSupported_81_A0(uint32) *OBD2
	// SetS01PID8E_EngFrictionPctTorq sets the physical value of the S01PID8E_EngFrictionPctTorq signal.
	SetS01PID8E_EngFrictionPctTorq(float64) *OBD2
	// SetRawS01PID8E_EngFrictionPctTorq sets the raw (encoded) value of the S01PID8E_EngFrictionPctTorq signal.
	SetRawS01PID8E_EngFrictionPctTorq(uint8) *OBD2
	// SetS01PIDA0_PIDsSupported_A1_C0 sets the value of the S01PIDA0_PIDsSupported_A1_C0 signal.
	SetS01PIDA0_PIDsSupported_A1_C0(uint32) *OBD2
	// SetS01PIDC0_PIDsSupported_C1_E0 sets the value of the S01PIDC0_PIDsSupported_C1_E0 signal.
	SetS01PIDC0_PIDsSupported_C1_E0(uint32) *OBD2
}

type OBD2 struct {
	xxx_Length                          uint8
	xxx_Service                         OBD2_Service
	xxx_Response                        uint8
	xxx_S01PID                          OBD2_S01PID
	xxx_S02PID                          OBD2_S02PID
	xxx_S01PID00_PIDsSupported_01_20    uint32
	xxx_S01PID01_MonitorStatus          uint32
	xxx_S02PID02_FreezeDTC              uint16
	xxx_S01PID02_FreezeDTC              uint16
	xxx_S01PID03_FuelSystemStatus       OBD2_S01PID03_FuelSystemStatus
	xxx_S01PID04_CalcEngineLoad         uint8
	xxx_S01PID05_EngineCoolantTemp      uint8
	xxx_S01PID06_ShortFuelTrimBank1     uint8
	xxx_S01PID07_LongFuelTrimBank1      uint8
	xxx_S01PID08_ShortFuelTrimBank2     uint8
	xxx_S01PID09_LongFuelTrimBank2      uint8
	xxx_S01PID0A_FuelPressure           uint8
	xxx_S01PID0B_IntakeManiAbsPress     uint8
	xxx_S01PID0C_EngineRPM              uint16
	xxx_S01PID0D_VehicleSpeed           uint8
	xxx_S01PID0E_TimingAdvance          uint8
	xxx_S01PID0F_IntakeAirTemperature   uint8
	xxx_S01PID10_MAFAirFlowRate         uint16
	xxx_S01PID11_ThrottlePosition       uint8
	xxx_S01PID12_CmdSecAirStatus        OBD2_S01PID12_CmdSecAirStatus
	xxx_S01PID14_OxySensor1_Volt        uint8
	xxx_S01PID15_OxySensor2_Volt        uint8
	xxx_S01PID16_OxySensor3_Volt        uint8
	xxx_S01PID17_OxySensor4_Volt        uint8
	xxx_S01PID18_OxySensor5_Volt        uint8
	xxx_S01PID19_OxySensor6_Volt        uint8
	xxx_S01PID14_OxySensor1_STFT        uint8
	xxx_S01PID15_OxySensor2_STFT        uint8
	xxx_S01PID16_OxySensor3_STFT        uint8
	xxx_S01PID17_OxySensor4_STFT        uint8
	xxx_S01PID18_OxySensor5_STFT        uint8
	xxx_S01PID19_OxySensor6_STFT        uint8
	xxx_S01PID1A_OxySensor7_Volt        uint8
	xxx_S01PID1A_OxySensor7_STFT        uint8
	xxx_S01PID1B_OxySensor8_Volt        uint8
	xxx_S01PID1B_OxySensor8_STFT        uint8
	xxx_S01PID1C_OBDStandard            OBD2_S01PID1C_OBDStandard
	xxx_S01PID1F_TimeSinceEngStart      uint16
	xxx_S01PID20_PIDsSupported_21_40    uint32
	xxx_S01PID21_DistanceMILOn          uint16
	xxx_S01PID22_FuelRailPres           uint16
	xxx_S01PID23_FuelRailGaug           uint16
	xxx_S01PID24_OxySensor1_FAER        uint16
	xxx_S01PID24_OxySensor1_Volt        uint16
	xxx_S01PID25_OxySensor2_FAER        uint16
	xxx_S01PID25_OxySensor2_Volt        uint16
	xxx_S01PID26_OxySensor3_FAER        uint16
	xxx_S01PID26_OxySensor3_Volt        uint16
	xxx_S01PID27_OxySensor4_FAER        uint16
	xxx_S01PID28_OxySensor5_FAER        uint16
	xxx_S01PID29_OxySensor6_FAER        uint16
	xxx_S01PID27_OxySensor4_Volt        uint16
	xxx_S01PID28_OxySensor5_Volt        uint16
	xxx_S01PID29_OxySensor6_Volt        uint16
	xxx_S01PID2A_OxySensor7_FAER        uint16
	xxx_S01PID2A_OxySensor7_Volt        uint16
	xxx_S01PID2B_OxySensor8_FAER        uint16
	xxx_S01PID2B_OxySensor8_Volt        uint16
	xxx_S01PID2C_CmdEGR                 uint8
	xxx_S01PID2D_EGRError               uint8
	xxx_S01PID2E_CmdEvapPurge           uint8
	xxx_S01PID2F_FuelTankLevel          uint8
	xxx_S01PID30_WarmUpsSinceCodeClear  uint8
	xxx_S01PID31_DistanceSinceCodeClear uint16
	xxx_S01PID32_EvapSysVaporPres       int16
	xxx_S01PID33_AbsBaroPres            uint8
	xxx_S01PID34_OxySensor1_FAER        uint16
	xxx_S01PID34_OxySensor1_Crnt        uint16
	xxx_S01PID35_OxySensor2_FAER        uint16
	xxx_S01PID35_OxySensor2_Crnt        uint16
	xxx_S01PID36_OxySensor3_FAER        uint16
	xxx_S01PID36_OxySensor3_Crnt        uint16
	xxx_S01PID37_OxySensor4_FAER        uint16
	xxx_S01PID38_OxySensor5_FAER        uint16
	xxx_S01PID39_OxySensor6_FAER        uint16
	xxx_S01PID37_OxySensor4_Crnt        uint16
	xxx_S01PID38_OxySensor5_Crnt        uint16
	xxx_S01PID39_OxySensor6_Crnt        uint16
	xxx_S01PID3A_OxySensor7_FAER        uint16
	xxx_S01PID3B_OxySensor8_FAER        uint16
	xxx_S01PID3C_CatTempBank1Sens1      uint16
	xxx_S01PID3D_CatTempBank2Sens1      uint16
	xxx_S01PID3A_OxySensor7_Crnt        uint16
	xxx_S01PID3B_OxySensor8_Crnt        uint16
	xxx_S01PID3E_CatTempBank1Sens2      uint16
	xxx_S01PID3F_CatTempBank2Sens2      uint16
	xxx_S01PID40_PIDsSupported_41_60    uint32
	xxx_S01PID41_MonStatusDriveCycle    uint32
	xxx_S01PID42_ControlModuleVolt      uint16
	xxx_S01PID43_AbsLoadValue           uint16
	xxx_S01PID44_FuelAirCmdEquiv        uint16
	xxx_S01PID45_RelThrottlePos         uint8
	xxx_S01PID46_AmbientAirTemp         uint8
	xxx_S01PID47_AbsThrottlePosB        uint8
	xxx_S01PID48_AbsThrottlePosC        uint8
	xxx_S01PID49_AbsThrottlePosD        uint8
	xxx_S01PID4A_AbsThrottlePosE        uint8
	xxx_S01PID4B_AbsThrottlePosF        uint8
	xxx_S01PID4C_CmdThrottleAct         uint8
	xxx_S01PID4D_TimeRunMILOn           uint16
	xxx_S01PID4E_TimeSinceCodeClear     uint16
	xxx_S01PID4F_Max_FAER               uint8
	xxx_S01PID4F_Max_OxySensVol         uint8
	xxx_S01PID4F_Max_OxySensCrnt        uint8
	xxx_S01PID4F_Max_IntManiAbsPres     uint8
	xxx_S01PID50_Max_AirFlowMAF         uint8
	xxx_S01PID51_FuelType               OBD2_S01PID51_FuelType
	xxx_S01PID52_EthanolFuelPct         uint8
	xxx_S01PID53_AbsEvapSysVapPres      uint16
	xxx_S01PID54_EvapSysVapPres         uint16
	xxx_S01PID55_ShortSecOxyTrimBank1   uint8
	xxx_S01PID56_LongSecOxyTrimBank1    uint8
	xxx_S01PID55_ShortSecOxyTrimBank3   uint8
	xxx_S01PID56_LongSecOxyTrimBank3    uint8
	xxx_S01PID57_ShortSecOxyTrimBank2   uint8
	xxx_S01PID58_LongSecOxyTrimBank2    uint8
	xxx_S01PID59_FuelRailAbsPres        uint16
	xxx_S01PID5A_RelAccelPedalPos       uint8
	xxx_S01PID57_ShortSecOxyTrimBank4   uint8
	xxx_S01PID58_LongSecOxyTrimBank4    uint8
	xxx_S01PID5B_HybrBatPackRemLife     uint8
	xxx_S01PID5C_EngineOilTemp          uint8
	xxx_S01PID5D_FuelInjectionTiming    uint16
	xxx_S01PID5E_EngineFuelRate         uint16
	xxx_S01PID5F_EmissionReq            uint8
	xxx_S01PID60_PIDsSupported_61_80    uint32
	xxx_S01PID61_DemandEngTorqPct       uint8
	xxx_S01PID62_ActualEngTorqPct       uint8
	xxx_S01PID63_EngRefTorq             uint16
	xxx_S01PID64_EngPctTorq_Idle        uint8
	xxx_S01PID64_EngPctTorq_EP1         uint8
	xxx_S01PID64_EngPctTorq_EP2         uint8
	xxx_S01PID64_EngPctTorq_EP3         uint8
	xxx_S01PID65_AuxInputOutput         uint8
	xxx_S01PID66_MAFSensor              uint8
	xxx_S01PID64_EngPctTorq_EP4         uint8
	xxx_S01PID67_EngineCoolantTemp      uint8
	xxx_S01PID68_IntakeAirTempSens      uint8
	xxx_S01PID69_CmdEGR_EGRError        uint8
	xxx_S01PID6A_CmdDieselIntAir        uint8
	xxx_S01PID6B_ExhaustGasTemp         uint8
	xxx_S01PID6C_CmdThrottleActRel      uint8
	xxx_S01PID6D_FuelPresContrSys       uint8
	xxx_S01PID6E_InjPresContrSys        uint8
	xxx_S01PID6F_TurboComprPres         uint8
	xxx_S01PID70_BoostPresCntrl         uint8
	xxx_S01PID80_PIDsSupported_81_A0    uint32
	xxx_S01PID8E_EngFrictionPctTorq     uint8
	xxx_S01PIDA0_PIDsSupported_A1_C0    uint32
	xxx_S01PIDC0_PIDsSupported_C1_E0    uint32
}

func NewOBD2() *OBD2 {
	m := &OBD2{}
	m.Reset()
	return m
}

func (m *OBD2) Reset() {
	m.xxx_Length = 0
	m.xxx_Service = 0
	m.xxx_Response = 0
	m.xxx_S01PID = 0
	m.xxx_S02PID = 0
	m.xxx_S01PID00_PIDsSupported_01_20 = 0
	m.xxx_S01PID01_MonitorStatus = 0
	m.xxx_S02PID02_FreezeDTC = 0
	m.xxx_S01PID02_FreezeDTC = 0
	m.xxx_S01PID03_FuelSystemStatus = 0
	m.xxx_S01PID04_CalcEngineLoad = 0
	m.xxx_S01PID05_EngineCoolantTemp = 0
	m.xxx_S01PID06_ShortFuelTrimBank1 = 0
	m.xxx_S01PID07_LongFuelTrimBank1 = 0
	m.xxx_S01PID08_ShortFuelTrimBank2 = 0
	m.xxx_S01PID09_LongFuelTrimBank2 = 0
	m.xxx_S01PID0A_FuelPressure = 0
	m.xxx_S01PID0B_IntakeManiAbsPress = 0
	m.xxx_S01PID0C_EngineRPM = 0
	m.xxx_S01PID0D_VehicleSpeed = 0
	m.xxx_S01PID0E_TimingAdvance = 0
	m.xxx_S01PID0F_IntakeAirTemperature = 0
	m.xxx_S01PID10_MAFAirFlowRate = 0
	m.xxx_S01PID11_ThrottlePosition = 0
	m.xxx_S01PID12_CmdSecAirStatus = 0
	m.xxx_S01PID14_OxySensor1_Volt = 0
	m.xxx_S01PID15_OxySensor2_Volt = 0
	m.xxx_S01PID16_OxySensor3_Volt = 0
	m.xxx_S01PID17_OxySensor4_Volt = 0
	m.xxx_S01PID18_OxySensor5_Volt = 0
	m.xxx_S01PID19_OxySensor6_Volt = 0
	m.xxx_S01PID14_OxySensor1_STFT = 0
	m.xxx_S01PID15_OxySensor2_STFT = 0
	m.xxx_S01PID16_OxySensor3_STFT = 0
	m.xxx_S01PID17_OxySensor4_STFT = 0
	m.xxx_S01PID18_OxySensor5_STFT = 0
	m.xxx_S01PID19_OxySensor6_STFT = 0
	m.xxx_S01PID1A_OxySensor7_Volt = 0
	m.xxx_S01PID1A_OxySensor7_STFT = 0
	m.xxx_S01PID1B_OxySensor8_Volt = 0
	m.xxx_S01PID1B_OxySensor8_STFT = 0
	m.xxx_S01PID1C_OBDStandard = 0
	m.xxx_S01PID1F_TimeSinceEngStart = 0
	m.xxx_S01PID20_PIDsSupported_21_40 = 0
	m.xxx_S01PID21_DistanceMILOn = 0
	m.xxx_S01PID22_FuelRailPres = 0
	m.xxx_S01PID23_FuelRailGaug = 0
	m.xxx_S01PID24_OxySensor1_FAER = 0
	m.xxx_S01PID24_OxySensor1_Volt = 0
	m.xxx_S01PID25_OxySensor2_FAER = 0
	m.xxx_S01PID25_OxySensor2_Volt = 0
	m.xxx_S01PID26_OxySensor3_FAER = 0
	m.xxx_S01PID26_OxySensor3_Volt = 0
	m.xxx_S01PID27_OxySensor4_FAER = 0
	m.xxx_S01PID28_OxySensor5_FAER = 0
	m.xxx_S01PID29_OxySensor6_FAER = 0
	m.xxx_S01PID27_OxySensor4_Volt = 0
	m.xxx_S01PID28_OxySensor5_Volt = 0
	m.xxx_S01PID29_OxySensor6_Volt = 0
	m.xxx_S01PID2A_OxySensor7_FAER = 0
	m.xxx_S01PID2A_OxySensor7_Volt = 0
	m.xxx_S01PID2B_OxySensor8_FAER = 0
	m.xxx_S01PID2B_OxySensor8_Volt = 0
	m.xxx_S01PID2C_CmdEGR = 0
	m.xxx_S01PID2D_EGRError = 0
	m.xxx_S01PID2E_CmdEvapPurge = 0
	m.xxx_S01PID2F_FuelTankLevel = 0
	m.xxx_S01PID30_WarmUpsSinceCodeClear = 0
	m.xxx_S01PID31_DistanceSinceCodeClear = 0
	m.xxx_S01PID32_EvapSysVaporPres = 0
	m.xxx_S01PID33_AbsBaroPres = 0
	m.xxx_S01PID34_OxySensor1_FAER = 0
	m.xxx_S01PID34_OxySensor1_Crnt = 0
	m.xxx_S01PID35_OxySensor2_FAER = 0
	m.xxx_S01PID35_OxySensor2_Crnt = 0
	m.xxx_S01PID36_OxySensor3_FAER = 0
	m.xxx_S01PID36_OxySensor3_Crnt = 0
	m.xxx_S01PID37_OxySensor4_FAER = 0
	m.xxx_S01PID38_OxySensor5_FAER = 0
	m.xxx_S01PID39_OxySensor6_FAER = 0
	m.xxx_S01PID37_OxySensor4_Crnt = 0
	m.xxx_S01PID38_OxySensor5_Crnt = 0
	m.xxx_S01PID39_OxySensor6_Crnt = 0
	m.xxx_S01PID3A_OxySensor7_FAER = 0
	m.xxx_S01PID3B_OxySensor8_FAER = 0
	m.xxx_S01PID3C_CatTempBank1Sens1 = 0
	m.xxx_S01PID3D_CatTempBank2Sens1 = 0
	m.xxx_S01PID3A_OxySensor7_Crnt = 0
	m.xxx_S01PID3B_OxySensor8_Crnt = 0
	m.xxx_S01PID3E_CatTempBank1Sens2 = 0
	m.xxx_S01PID3F_CatTempBank2Sens2 = 0
	m.xxx_S01PID40_PIDsSupported_41_60 = 0
	m.xxx_S01PID41_MonStatusDriveCycle = 0
	m.xxx_S01PID42_ControlModuleVolt = 0
	m.xxx_S01PID43_AbsLoadValue = 0
	m.xxx_S01PID44_FuelAirCmdEquiv = 0
	m.xxx_S01PID45_RelThrottlePos = 0
	m.xxx_S01PID46_AmbientAirTemp = 0
	m.xxx_S01PID47_AbsThrottlePosB = 0
	m.xxx_S01PID48_AbsThrottlePosC = 0
	m.xxx_S01PID49_AbsThrottlePosD = 0
	m.xxx_S01PID4A_AbsThrottlePosE = 0
	m.xxx_S01PID4B_AbsThrottlePosF = 0
	m.xxx_S01PID4C_CmdThrottleAct = 0
	m.xxx_S01PID4D_TimeRunMILOn = 0
	m.xxx_S01PID4E_TimeSinceCodeClear = 0
	m.xxx_S01PID4F_Max_FAER = 0
	m.xxx_S01PID4F_Max_OxySensVol = 0
	m.xxx_S01PID4F_Max_OxySensCrnt = 0
	m.xxx_S01PID4F_Max_IntManiAbsPres = 0
	m.xxx_S01PID50_Max_AirFlowMAF = 0
	m.xxx_S01PID51_FuelType = 0
	m.xxx_S01PID52_EthanolFuelPct = 0
	m.xxx_S01PID53_AbsEvapSysVapPres = 0
	m.xxx_S01PID54_EvapSysVapPres = 0
	m.xxx_S01PID55_ShortSecOxyTrimBank1 = 0
	m.xxx_S01PID56_LongSecOxyTrimBank1 = 0
	m.xxx_S01PID55_ShortSecOxyTrimBank3 = 0
	m.xxx_S01PID56_LongSecOxyTrimBank3 = 0
	m.xxx_S01PID57_ShortSecOxyTrimBank2 = 0
	m.xxx_S01PID58_LongSecOxyTrimBank2 = 0
	m.xxx_S01PID59_FuelRailAbsPres = 0
	m.xxx_S01PID5A_RelAccelPedalPos = 0
	m.xxx_S01PID57_ShortSecOxyTrimBank4 = 0
	m.xxx_S01PID58_LongSecOxyTrimBank4 = 0
	m.xxx_S01PID5B_HybrBatPackRemLife = 0
	m.xxx_S01PID5C_EngineOilTemp = 0
	m.xxx_S01PID5D_FuelInjectionTiming = 0
	m.xxx_S01PID5E_EngineFuelRate = 0
	m.xxx_S01PID5F_EmissionReq = 0
	m.xxx_S01PID60_PIDsSupported_61_80 = 0
	m.xxx_S01PID61_DemandEngTorqPct = 0
	m.xxx_S01PID62_ActualEngTorqPct = 0
	m.xxx_S01PID63_EngRefTorq = 0
	m.xxx_S01PID64_EngPctTorq_Idle = 0
	m.xxx_S01PID64_EngPctTorq_EP1 = 0
	m.xxx_S01PID64_EngPctTorq_EP2 = 0
	m.xxx_S01PID64_EngPctTorq_EP3 = 0
	m.xxx_S01PID65_AuxInputOutput = 0
	m.xxx_S01PID66_MAFSensor = 0
	m.xxx_S01PID64_EngPctTorq_EP4 = 0
	m.xxx_S01PID67_EngineCoolantTemp = 0
	m.xxx_S01PID68_IntakeAirTempSens = 0
	m.xxx_S01PID69_CmdEGR_EGRError = 0
	m.xxx_S01PID6A_CmdDieselIntAir = 0
	m.xxx_S01PID6B_ExhaustGasTemp = 0
	m.xxx_S01PID6C_CmdThrottleActRel = 0
	m.xxx_S01PID6D_FuelPresContrSys = 0
	m.xxx_S01PID6E_InjPresContrSys = 0
	m.xxx_S01PID6F_TurboComprPres = 0
	m.xxx_S01PID70_BoostPresCntrl = 0
	m.xxx_S01PID80_PIDsSupported_81_A0 = 0
	m.xxx_S01PID8E_EngFrictionPctTorq = 0
	m.xxx_S01PIDA0_PIDsSupported_A1_C0 = 0
	m.xxx_S01PIDC0_PIDsSupported_C1_E0 = 0
}

func (m *OBD2) CopyFrom(o OBD2Reader) *OBD2 {
	f, _ := o.MarshalFrame()
	_ = m.UnmarshalFrame(f)
	return m
}

// Descriptor returns the OBD2 descriptor.
func (m *OBD2) Descriptor() *descriptor.Message {
	return Messages().OBD2.Message
}

// String returns a compact string representation of the message.
func (m *OBD2) String() string {
	return cantext.MessageString(m)
}

func (m *OBD2) Length() uint8 {
	return m.xxx_Length
}

func (m *OBD2) SetLength(v uint8) *OBD2 {
	m.xxx_Length = uint8(Messages().OBD2.Length.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) Service() OBD2_Service {
	return m.xxx_Service
}

func (m *OBD2) SetService(v OBD2_Service) *OBD2 {
	m.xxx_Service = OBD2_Service(Messages().OBD2.Service.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) Response() uint8 {
	return m.xxx_Response
}

func (m *OBD2) SetResponse(v uint8) *OBD2 {
	m.xxx_Response = uint8(Messages().OBD2.Response.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID() OBD2_S01PID {
	return m.xxx_S01PID
}

func (m *OBD2) SetS01PID(v OBD2_S01PID) *OBD2 {
	m.xxx_S01PID = OBD2_S01PID(Messages().OBD2.S01PID.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S02PID() OBD2_S02PID {
	return m.xxx_S02PID
}

func (m *OBD2) SetS02PID(v OBD2_S02PID) *OBD2 {
	m.xxx_S02PID = OBD2_S02PID(Messages().OBD2.S02PID.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID00_PIDsSupported_01_20() uint32 {
	return m.xxx_S01PID00_PIDsSupported_01_20
}

func (m *OBD2) SetS01PID00_PIDsSupported_01_20(v uint32) *OBD2 {
	m.xxx_S01PID00_PIDsSupported_01_20 = uint32(Messages().OBD2.S01PID00_PIDsSupported_01_20.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID01_MonitorStatus() uint32 {
	return m.xxx_S01PID01_MonitorStatus
}

func (m *OBD2) SetS01PID01_MonitorStatus(v uint32) *OBD2 {
	m.xxx_S01PID01_MonitorStatus = uint32(Messages().OBD2.S01PID01_MonitorStatus.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S02PID02_FreezeDTC() uint16 {
	return m.xxx_S02PID02_FreezeDTC
}

func (m *OBD2) SetS02PID02_FreezeDTC(v uint16) *OBD2 {
	m.xxx_S02PID02_FreezeDTC = uint16(Messages().OBD2.S02PID02_FreezeDTC.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID02_FreezeDTC() uint16 {
	return m.xxx_S01PID02_FreezeDTC
}

func (m *OBD2) SetS01PID02_FreezeDTC(v uint16) *OBD2 {
	m.xxx_S01PID02_FreezeDTC = uint16(Messages().OBD2.S01PID02_FreezeDTC.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID03_FuelSystemStatus() OBD2_S01PID03_FuelSystemStatus {
	return m.xxx_S01PID03_FuelSystemStatus
}

func (m *OBD2) SetS01PID03_FuelSystemStatus(v OBD2_S01PID03_FuelSystemStatus) *OBD2 {
	m.xxx_S01PID03_FuelSystemStatus = OBD2_S01PID03_FuelSystemStatus(Messages().OBD2.S01PID03_FuelSystemStatus.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID04_CalcEngineLoad() float64 {
	return Messages().OBD2.S01PID04_CalcEngineLoad.ToPhysical(float64(m.xxx_S01PID04_CalcEngineLoad))
}

func (m *OBD2) SetS01PID04_CalcEngineLoad(v float64) *OBD2 {
	m.xxx_S01PID04_CalcEngineLoad = uint8(Messages().OBD2.S01PID04_CalcEngineLoad.FromPhysical(v))
	return m
}

func (m *OBD2) RawS01PID04_CalcEngineLoad() uint8 {
	return m.xxx_S01PID04_CalcEngineLoad
}

func (m *OBD2) SetRawS01PID04_CalcEngineLoad(v uint8) *OBD2 {
	m.xxx_S01PID04_CalcEngineLoad = uint8(Messages().OBD2.S01PID04_CalcEngineLoad.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID05_EngineCoolantTemp() float64 {
	return Messages().OBD2.S01PID05_EngineCoolantTemp.ToPhysical(float64(m.xxx_S01PID05_EngineCoolantTemp))
}

func (m *OBD2) SetS01PID05_EngineCoolantTemp(v float64) *OBD2 {
	m.xxx_S01PID05_EngineCoolantTemp = uint8(Messages().OBD2.S01PID05_EngineCoolantTemp.FromPhysical(v))
	return m
}

func (m *OBD2) RawS01PID05_EngineCoolantTemp() uint8 {
	return m.xxx_S01PID05_EngineCoolantTemp
}

func (m *OBD2) SetRawS01PID05_EngineCoolantTemp(v uint8) *OBD2 {
	m.xxx_S01PID05_EngineCoolantTemp = uint8(Messages().OBD2.S01PID05_EngineCoolantTemp.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID06_ShortFuelTrimBank1() float64 {
	return Messages().OBD2.S01PID06_ShortFuelTrimBank1.ToPhysical(float64(m.xxx_S01PID06_ShortFuelTrimBank1))
}

func (m *OBD2) SetS01PID06_ShortFuelTrimBank1(v float64) *OBD2 {
	m.xxx_S01PID06_ShortFuelTrimBank1 = uint8(Messages().OBD2.S01PID06_ShortFuelTrimBank1.FromPhysical(v))
	return m
}

func (m *OBD2) RawS01PID06_ShortFuelTrimBank1() uint8 {
	return m.xxx_S01PID06_ShortFuelTrimBank1
}

func (m *OBD2) SetRawS01PID06_ShortFuelTrimBank1(v uint8) *OBD2 {
	m.xxx_S01PID06_ShortFuelTrimBank1 = uint8(Messages().OBD2.S01PID06_ShortFuelTrimBank1.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID07_LongFuelTrimBank1() float64 {
	return Messages().OBD2.S01PID07_LongFuelTrimBank1.ToPhysical(float64(m.xxx_S01PID07_LongFuelTrimBank1))
}

func (m *OBD2) SetS01PID07_LongFuelTrimBank1(v float64) *OBD2 {
	m.xxx_S01PID07_LongFuelTrimBank1 = uint8(Messages().OBD2.S01PID07_LongFuelTrimBank1.FromPhysical(v))
	return m
}

func (m *OBD2) RawS01PID07_LongFuelTrimBank1() uint8 {
	return m.xxx_S01PID07_LongFuelTrimBank1
}

func (m *OBD2) SetRawS01PID07_LongFuelTrimBank1(v uint8) *OBD2 {
	m.xxx_S01PID07_LongFuelTrimBank1 = uint8(Messages().OBD2.S01PID07_LongFuelTrimBank1.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID08_ShortFuelTrimBank2() float64 {
	return Messages().OBD2.S01PID08_ShortFuelTrimBank2.ToPhysical(float64(m.xxx_S01PID08_ShortFuelTrimBank2))
}

func (m *OBD2) SetS01PID08_ShortFuelTrimBank2(v float64) *OBD2 {
	m.xxx_S01PID08_ShortFuelTrimBank2 = uint8(Messages().OBD2.S01PID08_ShortFuelTrimBank2.FromPhysical(v))
	return m
}

func (m *OBD2) RawS01PID08_ShortFuelTrimBank2() uint8 {
	return m.xxx_S01PID08_ShortFuelTrimBank2
}

func (m *OBD2) SetRawS01PID08_ShortFuelTrimBank2(v uint8) *OBD2 {
	m.xxx_S01PID08_ShortFuelTrimBank2 = uint8(Messages().OBD2.S01PID08_ShortFuelTrimBank2.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID09_LongFuelTrimBank2() float64 {
	return Messages().OBD2.S01PID09_LongFuelTrimBank2.ToPhysical(float64(m.xxx_S01PID09_LongFuelTrimBank2))
}

func (m *OBD2) SetS01PID09_LongFuelTrimBank2(v float64) *OBD2 {
	m.xxx_S01PID09_LongFuelTrimBank2 = uint8(Messages().OBD2.S01PID09_LongFuelTrimBank2.FromPhysical(v))
	return m
}

func (m *OBD2) RawS01PID09_LongFuelTrimBank2() uint8 {
	return m.xxx_S01PID09_LongFuelTrimBank2
}

func (m *OBD2) SetRawS01PID09_LongFuelTrimBank2(v uint8) *OBD2 {
	m.xxx_S01PID09_LongFuelTrimBank2 = uint8(Messages().OBD2.S01PID09_LongFuelTrimBank2.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID0A_FuelPressure() float64 {
	return Messages().OBD2.S01PID0A_FuelPressure.ToPhysical(float64(m.xxx_S01PID0A_FuelPressure))
}

func (m *OBD2) SetS01PID0A_FuelPressure(v float64) *OBD2 {
	m.xxx_S01PID0A_FuelPressure = uint8(Messages().OBD2.S01PID0A_FuelPressure.FromPhysical(v))
	return m
}

func (m *OBD2) RawS01PID0A_FuelPressure() uint8 {
	return m.xxx_S01PID0A_FuelPressure
}

func (m *OBD2) SetRawS01PID0A_FuelPressure(v uint8) *OBD2 {
	m.xxx_S01PID0A_FuelPressure = uint8(Messages().OBD2.S01PID0A_FuelPressure.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID0B_IntakeManiAbsPress() uint8 {
	return m.xxx_S01PID0B_IntakeManiAbsPress
}

func (m *OBD2) SetS01PID0B_IntakeManiAbsPress(v uint8) *OBD2 {
	m.xxx_S01PID0B_IntakeManiAbsPress = uint8(Messages().OBD2.S01PID0B_IntakeManiAbsPress.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID0C_EngineRPM() float64 {
	return Messages().OBD2.S01PID0C_EngineRPM.ToPhysical(float64(m.xxx_S01PID0C_EngineRPM))
}

func (m *OBD2) SetS01PID0C_EngineRPM(v float64) *OBD2 {
	m.xxx_S01PID0C_EngineRPM = uint16(Messages().OBD2.S01PID0C_EngineRPM.FromPhysical(v))
	return m
}

func (m *OBD2) RawS01PID0C_EngineRPM() uint16 {
	return m.xxx_S01PID0C_EngineRPM
}

func (m *OBD2) SetRawS01PID0C_EngineRPM(v uint16) *OBD2 {
	m.xxx_S01PID0C_EngineRPM = uint16(Messages().OBD2.S01PID0C_EngineRPM.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID0D_VehicleSpeed() uint8 {
	return m.xxx_S01PID0D_VehicleSpeed
}

func (m *OBD2) SetS01PID0D_VehicleSpeed(v uint8) *OBD2 {
	m.xxx_S01PID0D_VehicleSpeed = uint8(Messages().OBD2.S01PID0D_VehicleSpeed.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID0E_TimingAdvance() float64 {
	return Messages().OBD2.S01PID0E_TimingAdvance.ToPhysical(float64(m.xxx_S01PID0E_TimingAdvance))
}

func (m *OBD2) SetS01PID0E_TimingAdvance(v float64) *OBD2 {
	m.xxx_S01PID0E_TimingAdvance = uint8(Messages().OBD2.S01PID0E_TimingAdvance.FromPhysical(v))
	return m
}

func (m *OBD2) RawS01PID0E_TimingAdvance() uint8 {
	return m.xxx_S01PID0E_TimingAdvance
}

func (m *OBD2) SetRawS01PID0E_TimingAdvance(v uint8) *OBD2 {
	m.xxx_S01PID0E_TimingAdvance = uint8(Messages().OBD2.S01PID0E_TimingAdvance.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID0F_IntakeAirTemperature() float64 {
	return Messages().OBD2.S01PID0F_IntakeAirTemperature.ToPhysical(float64(m.xxx_S01PID0F_IntakeAirTemperature))
}

func (m *OBD2) SetS01PID0F_IntakeAirTemperature(v float64) *OBD2 {
	m.xxx_S01PID0F_IntakeAirTemperature = uint8(Messages().OBD2.S01PID0F_IntakeAirTemperature.FromPhysical(v))
	return m
}

func (m *OBD2) RawS01PID0F_IntakeAirTemperature() uint8 {
	return m.xxx_S01PID0F_IntakeAirTemperature
}

func (m *OBD2) SetRawS01PID0F_IntakeAirTemperature(v uint8) *OBD2 {
	m.xxx_S01PID0F_IntakeAirTemperature = uint8(Messages().OBD2.S01PID0F_IntakeAirTemperature.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID10_MAFAirFlowRate() float64 {
	return Messages().OBD2.S01PID10_MAFAirFlowRate.ToPhysical(float64(m.xxx_S01PID10_MAFAirFlowRate))
}

func (m *OBD2) SetS01PID10_MAFAirFlowRate(v float64) *OBD2 {
	m.xxx_S01PID10_MAFAirFlowRate = uint16(Messages().OBD2.S01PID10_MAFAirFlowRate.FromPhysical(v))
	return m
}

func (m *OBD2) RawS01PID10_MAFAirFlowRate() uint16 {
	return m.xxx_S01PID10_MAFAirFlowRate
}

func (m *OBD2) SetRawS01PID10_MAFAirFlowRate(v uint16) *OBD2 {
	m.xxx_S01PID10_MAFAirFlowRate = uint16(Messages().OBD2.S01PID10_MAFAirFlowRate.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID11_ThrottlePosition() float64 {
	return Messages().OBD2.S01PID11_ThrottlePosition.ToPhysical(float64(m.xxx_S01PID11_ThrottlePosition))
}

func (m *OBD2) SetS01PID11_ThrottlePosition(v float64) *OBD2 {
	m.xxx_S01PID11_ThrottlePosition = uint8(Messages().OBD2.S01PID11_ThrottlePosition.FromPhysical(v))
	return m
}

func (m *OBD2) RawS01PID11_ThrottlePosition() uint8 {
	return m.xxx_S01PID11_ThrottlePosition
}

func (m *OBD2) SetRawS01PID11_ThrottlePosition(v uint8) *OBD2 {
	m.xxx_S01PID11_ThrottlePosition = uint8(Messages().OBD2.S01PID11_ThrottlePosition.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID12_CmdSecAirStatus() OBD2_S01PID12_CmdSecAirStatus {
	return m.xxx_S01PID12_CmdSecAirStatus
}

func (m *OBD2) SetS01PID12_CmdSecAirStatus(v OBD2_S01PID12_CmdSecAirStatus) *OBD2 {
	m.xxx_S01PID12_CmdSecAirStatus = OBD2_S01PID12_CmdSecAirStatus(Messages().OBD2.S01PID12_CmdSecAirStatus.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID14_OxySensor1_Volt() float64 {
	return Messages().OBD2.S01PID14_OxySensor1_Volt.ToPhysical(float64(m.xxx_S01PID14_OxySensor1_Volt))
}

func (m *OBD2) SetS01PID14_OxySensor1_Volt(v float64) *OBD2 {
	m.xxx_S01PID14_OxySensor1_Volt = uint8(Messages().OBD2.S01PID14_OxySensor1_Volt.FromPhysical(v))
	return m
}

func (m *OBD2) RawS01PID14_OxySensor1_Volt() uint8 {
	return m.xxx_S01PID14_OxySensor1_Volt
}

func (m *OBD2) SetRawS01PID14_OxySensor1_Volt(v uint8) *OBD2 {
	m.xxx_S01PID14_OxySensor1_Volt = uint8(Messages().OBD2.S01PID14_OxySensor1_Volt.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID15_OxySensor2_Volt() float64 {
	return Messages().OBD2.S01PID15_OxySensor2_Volt.ToPhysical(float64(m.xxx_S01PID15_OxySensor2_Volt))
}

func (m *OBD2) SetS01PID15_OxySensor2_Volt(v float64) *OBD2 {
	m.xxx_S01PID15_OxySensor2_Volt = uint8(Messages().OBD2.S01PID15_OxySensor2_Volt.FromPhysical(v))
	return m
}

func (m *OBD2) RawS01PID15_OxySensor2_Volt() uint8 {
	return m.xxx_S01PID15_OxySensor2_Volt
}

func (m *OBD2) SetRawS01PID15_OxySensor2_Volt(v uint8) *OBD2 {
	m.xxx_S01PID15_OxySensor2_Volt = uint8(Messages().OBD2.S01PID15_OxySensor2_Volt.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID16_OxySensor3_Volt() float64 {
	return Messages().OBD2.S01PID16_OxySensor3_Volt.ToPhysical(float64(m.xxx_S01PID16_OxySensor3_Volt))
}

func (m *OBD2) SetS01PID16_OxySensor3_Volt(v float64) *OBD2 {
	m.xxx_S01PID16_OxySensor3_Volt = uint8(Messages().OBD2.S01PID16_OxySensor3_Volt.FromPhysical(v))
	return m
}

func (m *OBD2) RawS01PID16_OxySensor3_Volt() uint8 {
	return m.xxx_S01PID16_OxySensor3_Volt
}

func (m *OBD2) SetRawS01PID16_OxySensor3_Volt(v uint8) *OBD2 {
	m.xxx_S01PID16_OxySensor3_Volt = uint8(Messages().OBD2.S01PID16_OxySensor3_Volt.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID17_OxySensor4_Volt() float64 {
	return Messages().OBD2.S01PID17_OxySensor4_Volt.ToPhysical(float64(m.xxx_S01PID17_OxySensor4_Volt))
}

func (m *OBD2) SetS01PID17_OxySensor4_Volt(v float64) *OBD2 {
	m.xxx_S01PID17_OxySensor4_Volt = uint8(Messages().OBD2.S01PID17_OxySensor4_Volt.FromPhysical(v))
	return m
}

func (m *OBD2) RawS01PID17_OxySensor4_Volt() uint8 {
	return m.xxx_S01PID17_OxySensor4_Volt
}

func (m *OBD2) SetRawS01PID17_OxySensor4_Volt(v uint8) *OBD2 {
	m.xxx_S01PID17_OxySensor4_Volt = uint8(Messages().OBD2.S01PID17_OxySensor4_Volt.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID18_OxySensor5_Volt() float64 {
	return Messages().OBD2.S01PID18_OxySensor5_Volt.ToPhysical(float64(m.xxx_S01PID18_OxySensor5_Volt))
}

func (m *OBD2) SetS01PID18_OxySensor5_Volt(v float64) *OBD2 {
	m.xxx_S01PID18_OxySensor5_Volt = uint8(Messages().OBD2.S01PID18_OxySensor5_Volt.FromPhysical(v))
	return m
}

func (m *OBD2) RawS01PID18_OxySensor5_Volt() uint8 {
	return m.xxx_S01PID18_OxySensor5_Volt
}

func (m *OBD2) SetRawS01PID18_OxySensor5_Volt(v uint8) *OBD2 {
	m.xxx_S01PID18_OxySensor5_Volt = uint8(Messages().OBD2.S01PID18_OxySensor5_Volt.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID19_OxySensor6_Volt() float64 {
	return Messages().OBD2.S01PID19_OxySensor6_Volt.ToPhysical(float64(m.xxx_S01PID19_OxySensor6_Volt))
}

func (m *OBD2) SetS01PID19_OxySensor6_Volt(v float64) *OBD2 {
	m.xxx_S01PID19_OxySensor6_Volt = uint8(Messages().OBD2.S01PID19_OxySensor6_Volt.FromPhysical(v))
	return m
}

func (m *OBD2) RawS01PID19_OxySensor6_Volt() uint8 {
	return m.xxx_S01PID19_OxySensor6_Volt
}

func (m *OBD2) SetRawS01PID19_OxySensor6_Volt(v uint8) *OBD2 {
	m.xxx_S01PID19_OxySensor6_Volt = uint8(Messages().OBD2.S01PID19_OxySensor6_Volt.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID14_OxySensor1_STFT() float64 {
	return Messages().OBD2.S01PID14_OxySensor1_STFT.ToPhysical(float64(m.xxx_S01PID14_OxySensor1_STFT))
}

func (m *OBD2) SetS01PID14_OxySensor1_STFT(v float64) *OBD2 {
	m.xxx_S01PID14_OxySensor1_STFT = uint8(Messages().OBD2.S01PID14_OxySensor1_STFT.FromPhysical(v))
	return m
}

func (m *OBD2) RawS01PID14_OxySensor1_STFT() uint8 {
	return m.xxx_S01PID14_OxySensor1_STFT
}

func (m *OBD2) SetRawS01PID14_OxySensor1_STFT(v uint8) *OBD2 {
	m.xxx_S01PID14_OxySensor1_STFT = uint8(Messages().OBD2.S01PID14_OxySensor1_STFT.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID15_OxySensor2_STFT() float64 {
	return Messages().OBD2.S01PID15_OxySensor2_STFT.ToPhysical(float64(m.xxx_S01PID15_OxySensor2_STFT))
}

func (m *OBD2) SetS01PID15_OxySensor2_STFT(v float64) *OBD2 {
	m.xxx_S01PID15_OxySensor2_STFT = uint8(Messages().OBD2.S01PID15_OxySensor2_STFT.FromPhysical(v))
	return m
}

func (m *OBD2) RawS01PID15_OxySensor2_STFT() uint8 {
	return m.xxx_S01PID15_OxySensor2_STFT
}

func (m *OBD2) SetRawS01PID15_OxySensor2_STFT(v uint8) *OBD2 {
	m.xxx_S01PID15_OxySensor2_STFT = uint8(Messages().OBD2.S01PID15_OxySensor2_STFT.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID16_OxySensor3_STFT() float64 {
	return Messages().OBD2.S01PID16_OxySensor3_STFT.ToPhysical(float64(m.xxx_S01PID16_OxySensor3_STFT))
}

func (m *OBD2) SetS01PID16_OxySensor3_STFT(v float64) *OBD2 {
	m.xxx_S01PID16_OxySensor3_STFT = uint8(Messages().OBD2.S01PID16_OxySensor3_STFT.FromPhysical(v))
	return m
}

func (m *OBD2) RawS01PID16_OxySensor3_STFT() uint8 {
	return m.xxx_S01PID16_OxySensor3_STFT
}

func (m *OBD2) SetRawS01PID16_OxySensor3_STFT(v uint8) *OBD2 {
	m.xxx_S01PID16_OxySensor3_STFT = uint8(Messages().OBD2.S01PID16_OxySensor3_STFT.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID17_OxySensor4_STFT() float64 {
	return Messages().OBD2.S01PID17_OxySensor4_STFT.ToPhysical(float64(m.xxx_S01PID17_OxySensor4_STFT))
}

func (m *OBD2) SetS01PID17_OxySensor4_STFT(v float64) *OBD2 {
	m.xxx_S01PID17_OxySensor4_STFT = uint8(Messages().OBD2.S01PID17_OxySensor4_STFT.FromPhysical(v))
	return m
}

func (m *OBD2) RawS01PID17_OxySensor4_STFT() uint8 {
	return m.xxx_S01PID17_OxySensor4_STFT
}

func (m *OBD2) SetRawS01PID17_OxySensor4_STFT(v uint8) *OBD2 {
	m.xxx_S01PID17_OxySensor4_STFT = uint8(Messages().OBD2.S01PID17_OxySensor4_STFT.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID18_OxySensor5_STFT() float64 {
	return Messages().OBD2.S01PID18_OxySensor5_STFT.ToPhysical(float64(m.xxx_S01PID18_OxySensor5_STFT))
}

func (m *OBD2) SetS01PID18_OxySensor5_STFT(v float64) *OBD2 {
	m.xxx_S01PID18_OxySensor5_STFT = uint8(Messages().OBD2.S01PID18_OxySensor5_STFT.FromPhysical(v))
	return m
}

func (m *OBD2) RawS01PID18_OxySensor5_STFT() uint8 {
	return m.xxx_S01PID18_OxySensor5_STFT
}

func (m *OBD2) SetRawS01PID18_OxySensor5_STFT(v uint8) *OBD2 {
	m.xxx_S01PID18_OxySensor5_STFT = uint8(Messages().OBD2.S01PID18_OxySensor5_STFT.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID19_OxySensor6_STFT() float64 {
	return Messages().OBD2.S01PID19_OxySensor6_STFT.ToPhysical(float64(m.xxx_S01PID19_OxySensor6_STFT))
}

func (m *OBD2) SetS01PID19_OxySensor6_STFT(v float64) *OBD2 {
	m.xxx_S01PID19_OxySensor6_STFT = uint8(Messages().OBD2.S01PID19_OxySensor6_STFT.FromPhysical(v))
	return m
}

func (m *OBD2) RawS01PID19_OxySensor6_STFT() uint8 {
	return m.xxx_S01PID19_OxySensor6_STFT
}

func (m *OBD2) SetRawS01PID19_OxySensor6_STFT(v uint8) *OBD2 {
	m.xxx_S01PID19_OxySensor6_STFT = uint8(Messages().OBD2.S01PID19_OxySensor6_STFT.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID1A_OxySensor7_Volt() float64 {
	return Messages().OBD2.S01PID1A_OxySensor7_Volt.ToPhysical(float64(m.xxx_S01PID1A_OxySensor7_Volt))
}

func (m *OBD2) SetS01PID1A_OxySensor7_Volt(v float64) *OBD2 {
	m.xxx_S01PID1A_OxySensor7_Volt = uint8(Messages().OBD2.S01PID1A_OxySensor7_Volt.FromPhysical(v))
	return m
}

func (m *OBD2) RawS01PID1A_OxySensor7_Volt() uint8 {
	return m.xxx_S01PID1A_OxySensor7_Volt
}

func (m *OBD2) SetRawS01PID1A_OxySensor7_Volt(v uint8) *OBD2 {
	m.xxx_S01PID1A_OxySensor7_Volt = uint8(Messages().OBD2.S01PID1A_OxySensor7_Volt.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID1A_OxySensor7_STFT() float64 {
	return Messages().OBD2.S01PID1A_OxySensor7_STFT.ToPhysical(float64(m.xxx_S01PID1A_OxySensor7_STFT))
}

func (m *OBD2) SetS01PID1A_OxySensor7_STFT(v float64) *OBD2 {
	m.xxx_S01PID1A_OxySensor7_STFT = uint8(Messages().OBD2.S01PID1A_OxySensor7_STFT.FromPhysical(v))
	return m
}

func (m *OBD2) RawS01PID1A_OxySensor7_STFT() uint8 {
	return m.xxx_S01PID1A_OxySensor7_STFT
}

func (m *OBD2) SetRawS01PID1A_OxySensor7_STFT(v uint8) *OBD2 {
	m.xxx_S01PID1A_OxySensor7_STFT = uint8(Messages().OBD2.S01PID1A_OxySensor7_STFT.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID1B_OxySensor8_Volt() float64 {
	return Messages().OBD2.S01PID1B_OxySensor8_Volt.ToPhysical(float64(m.xxx_S01PID1B_OxySensor8_Volt))
}

func (m *OBD2) SetS01PID1B_OxySensor8_Volt(v float64) *OBD2 {
	m.xxx_S01PID1B_OxySensor8_Volt = uint8(Messages().OBD2.S01PID1B_OxySensor8_Volt.FromPhysical(v))
	return m
}

func (m *OBD2) RawS01PID1B_OxySensor8_Volt() uint8 {
	return m.xxx_S01PID1B_OxySensor8_Volt
}

func (m *OBD2) SetRawS01PID1B_OxySensor8_Volt(v uint8) *OBD2 {
	m.xxx_S01PID1B_OxySensor8_Volt = uint8(Messages().OBD2.S01PID1B_OxySensor8_Volt.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID1B_OxySensor8_STFT() float64 {
	return Messages().OBD2.S01PID1B_OxySensor8_STFT.ToPhysical(float64(m.xxx_S01PID1B_OxySensor8_STFT))
}

func (m *OBD2) SetS01PID1B_OxySensor8_STFT(v float64) *OBD2 {
	m.xxx_S01PID1B_OxySensor8_STFT = uint8(Messages().OBD2.S01PID1B_OxySensor8_STFT.FromPhysical(v))
	return m
}

func (m *OBD2) RawS01PID1B_OxySensor8_STFT() uint8 {
	return m.xxx_S01PID1B_OxySensor8_STFT
}

func (m *OBD2) SetRawS01PID1B_OxySensor8_STFT(v uint8) *OBD2 {
	m.xxx_S01PID1B_OxySensor8_STFT = uint8(Messages().OBD2.S01PID1B_OxySensor8_STFT.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID1C_OBDStandard() OBD2_S01PID1C_OBDStandard {
	return m.xxx_S01PID1C_OBDStandard
}

func (m *OBD2) SetS01PID1C_OBDStandard(v OBD2_S01PID1C_OBDStandard) *OBD2 {
	m.xxx_S01PID1C_OBDStandard = OBD2_S01PID1C_OBDStandard(Messages().OBD2.S01PID1C_OBDStandard.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID1F_TimeSinceEngStart() uint16 {
	return m.xxx_S01PID1F_TimeSinceEngStart
}

func (m *OBD2) SetS01PID1F_TimeSinceEngStart(v uint16) *OBD2 {
	m.xxx_S01PID1F_TimeSinceEngStart = uint16(Messages().OBD2.S01PID1F_TimeSinceEngStart.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID20_PIDsSupported_21_40() uint32 {
	return m.xxx_S01PID20_PIDsSupported_21_40
}

func (m *OBD2) SetS01PID20_PIDsSupported_21_40(v uint32) *OBD2 {
	m.xxx_S01PID20_PIDsSupported_21_40 = uint32(Messages().OBD2.S01PID20_PIDsSupported_21_40.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID21_DistanceMILOn() uint16 {
	return m.xxx_S01PID21_DistanceMILOn
}

func (m *OBD2) SetS01PID21_DistanceMILOn(v uint16) *OBD2 {
	m.xxx_S01PID21_DistanceMILOn = uint16(Messages().OBD2.S01PID21_DistanceMILOn.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID22_FuelRailPres() float64 {
	return Messages().OBD2.S01PID22_FuelRailPres.ToPhysical(float64(m.xxx_S01PID22_FuelRailPres))
}

func (m *OBD2) SetS01PID22_FuelRailPres(v float64) *OBD2 {
	m.xxx_S01PID22_FuelRailPres = uint16(Messages().OBD2.S01PID22_FuelRailPres.FromPhysical(v))
	return m
}

func (m *OBD2) RawS01PID22_FuelRailPres() uint16 {
	return m.xxx_S01PID22_FuelRailPres
}

func (m *OBD2) SetRawS01PID22_FuelRailPres(v uint16) *OBD2 {
	m.xxx_S01PID22_FuelRailPres = uint16(Messages().OBD2.S01PID22_FuelRailPres.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID23_FuelRailGaug() float64 {
	return Messages().OBD2.S01PID23_FuelRailGaug.ToPhysical(float64(m.xxx_S01PID23_FuelRailGaug))
}

func (m *OBD2) SetS01PID23_FuelRailGaug(v float64) *OBD2 {
	m.xxx_S01PID23_FuelRailGaug = uint16(Messages().OBD2.S01PID23_FuelRailGaug.FromPhysical(v))
	return m
}

func (m *OBD2) RawS01PID23_FuelRailGaug() uint16 {
	return m.xxx_S01PID23_FuelRailGaug
}

func (m *OBD2) SetRawS01PID23_FuelRailGaug(v uint16) *OBD2 {
	m.xxx_S01PID23_FuelRailGaug = uint16(Messages().OBD2.S01PID23_FuelRailGaug.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID24_OxySensor1_FAER() float64 {
	return Messages().OBD2.S01PID24_OxySensor1_FAER.ToPhysical(float64(m.xxx_S01PID24_OxySensor1_FAER))
}

func (m *OBD2) SetS01PID24_OxySensor1_FAER(v float64) *OBD2 {
	m.xxx_S01PID24_OxySensor1_FAER = uint16(Messages().OBD2.S01PID24_OxySensor1_FAER.FromPhysical(v))
	return m
}

func (m *OBD2) RawS01PID24_OxySensor1_FAER() uint16 {
	return m.xxx_S01PID24_OxySensor1_FAER
}

func (m *OBD2) SetRawS01PID24_OxySensor1_FAER(v uint16) *OBD2 {
	m.xxx_S01PID24_OxySensor1_FAER = uint16(Messages().OBD2.S01PID24_OxySensor1_FAER.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID24_OxySensor1_Volt() float64 {
	return Messages().OBD2.S01PID24_OxySensor1_Volt.ToPhysical(float64(m.xxx_S01PID24_OxySensor1_Volt))
}

func (m *OBD2) SetS01PID24_OxySensor1_Volt(v float64) *OBD2 {
	m.xxx_S01PID24_OxySensor1_Volt = uint16(Messages().OBD2.S01PID24_OxySensor1_Volt.FromPhysical(v))
	return m
}

func (m *OBD2) RawS01PID24_OxySensor1_Volt() uint16 {
	return m.xxx_S01PID24_OxySensor1_Volt
}

func (m *OBD2) SetRawS01PID24_OxySensor1_Volt(v uint16) *OBD2 {
	m.xxx_S01PID24_OxySensor1_Volt = uint16(Messages().OBD2.S01PID24_OxySensor1_Volt.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID25_OxySensor2_FAER() float64 {
	return Messages().OBD2.S01PID25_OxySensor2_FAER.ToPhysical(float64(m.xxx_S01PID25_OxySensor2_FAER))
}

func (m *OBD2) SetS01PID25_OxySensor2_FAER(v float64) *OBD2 {
	m.xxx_S01PID25_OxySensor2_FAER = uint16(Messages().OBD2.S01PID25_OxySensor2_FAER.FromPhysical(v))
	return m
}

func (m *OBD2) RawS01PID25_OxySensor2_FAER() uint16 {
	return m.xxx_S01PID25_OxySensor2_FAER
}

func (m *OBD2) SetRawS01PID25_OxySensor2_FAER(v uint16) *OBD2 {
	m.xxx_S01PID25_OxySensor2_FAER = uint16(Messages().OBD2.S01PID25_OxySensor2_FAER.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID25_OxySensor2_Volt() float64 {
	return Messages().OBD2.S01PID25_OxySensor2_Volt.ToPhysical(float64(m.xxx_S01PID25_OxySensor2_Volt))
}

func (m *OBD2) SetS01PID25_OxySensor2_Volt(v float64) *OBD2 {
	m.xxx_S01PID25_OxySensor2_Volt = uint16(Messages().OBD2.S01PID25_OxySensor2_Volt.FromPhysical(v))
	return m
}

func (m *OBD2) RawS01PID25_OxySensor2_Volt() uint16 {
	return m.xxx_S01PID25_OxySensor2_Volt
}

func (m *OBD2) SetRawS01PID25_OxySensor2_Volt(v uint16) *OBD2 {
	m.xxx_S01PID25_OxySensor2_Volt = uint16(Messages().OBD2.S01PID25_OxySensor2_Volt.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID26_OxySensor3_FAER() float64 {
	return Messages().OBD2.S01PID26_OxySensor3_FAER.ToPhysical(float64(m.xxx_S01PID26_OxySensor3_FAER))
}

func (m *OBD2) SetS01PID26_OxySensor3_FAER(v float64) *OBD2 {
	m.xxx_S01PID26_OxySensor3_FAER = uint16(Messages().OBD2.S01PID26_OxySensor3_FAER.FromPhysical(v))
	return m
}

func (m *OBD2) RawS01PID26_OxySensor3_FAER() uint16 {
	return m.xxx_S01PID26_OxySensor3_FAER
}

func (m *OBD2) SetRawS01PID26_OxySensor3_FAER(v uint16) *OBD2 {
	m.xxx_S01PID26_OxySensor3_FAER = uint16(Messages().OBD2.S01PID26_OxySensor3_FAER.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID26_OxySensor3_Volt() float64 {
	return Messages().OBD2.S01PID26_OxySensor3_Volt.ToPhysical(float64(m.xxx_S01PID26_OxySensor3_Volt))
}

func (m *OBD2) SetS01PID26_OxySensor3_Volt(v float64) *OBD2 {
	m.xxx_S01PID26_OxySensor3_Volt = uint16(Messages().OBD2.S01PID26_OxySensor3_Volt.FromPhysical(v))
	return m
}

func (m *OBD2) RawS01PID26_OxySensor3_Volt() uint16 {
	return m.xxx_S01PID26_OxySensor3_Volt
}

func (m *OBD2) SetRawS01PID26_OxySensor3_Volt(v uint16) *OBD2 {
	m.xxx_S01PID26_OxySensor3_Volt = uint16(Messages().OBD2.S01PID26_OxySensor3_Volt.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID27_OxySensor4_FAER() float64 {
	return Messages().OBD2.S01PID27_OxySensor4_FAER.ToPhysical(float64(m.xxx_S01PID27_OxySensor4_FAER))
}

func (m *OBD2) SetS01PID27_OxySensor4_FAER(v float64) *OBD2 {
	m.xxx_S01PID27_OxySensor4_FAER = uint16(Messages().OBD2.S01PID27_OxySensor4_FAER.FromPhysical(v))
	return m
}

func (m *OBD2) RawS01PID27_OxySensor4_FAER() uint16 {
	return m.xxx_S01PID27_OxySensor4_FAER
}

func (m *OBD2) SetRawS01PID27_OxySensor4_FAER(v uint16) *OBD2 {
	m.xxx_S01PID27_OxySensor4_FAER = uint16(Messages().OBD2.S01PID27_OxySensor4_FAER.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID28_OxySensor5_FAER() float64 {
	return Messages().OBD2.S01PID28_OxySensor5_FAER.ToPhysical(float64(m.xxx_S01PID28_OxySensor5_FAER))
}

func (m *OBD2) SetS01PID28_OxySensor5_FAER(v float64) *OBD2 {
	m.xxx_S01PID28_OxySensor5_FAER = uint16(Messages().OBD2.S01PID28_OxySensor5_FAER.FromPhysical(v))
	return m
}

func (m *OBD2) RawS01PID28_OxySensor5_FAER() uint16 {
	return m.xxx_S01PID28_OxySensor5_FAER
}

func (m *OBD2) SetRawS01PID28_OxySensor5_FAER(v uint16) *OBD2 {
	m.xxx_S01PID28_OxySensor5_FAER = uint16(Messages().OBD2.S01PID28_OxySensor5_FAER.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID29_OxySensor6_FAER() float64 {
	return Messages().OBD2.S01PID29_OxySensor6_FAER.ToPhysical(float64(m.xxx_S01PID29_OxySensor6_FAER))
}

func (m *OBD2) SetS01PID29_OxySensor6_FAER(v float64) *OBD2 {
	m.xxx_S01PID29_OxySensor6_FAER = uint16(Messages().OBD2.S01PID29_OxySensor6_FAER.FromPhysical(v))
	return m
}

func (m *OBD2) RawS01PID29_OxySensor6_FAER() uint16 {
	return m.xxx_S01PID29_OxySensor6_FAER
}

func (m *OBD2) SetRawS01PID29_OxySensor6_FAER(v uint16) *OBD2 {
	m.xxx_S01PID29_OxySensor6_FAER = uint16(Messages().OBD2.S01PID29_OxySensor6_FAER.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID27_OxySensor4_Volt() float64 {
	return Messages().OBD2.S01PID27_OxySensor4_Volt.ToPhysical(float64(m.xxx_S01PID27_OxySensor4_Volt))
}

func (m *OBD2) SetS01PID27_OxySensor4_Volt(v float64) *OBD2 {
	m.xxx_S01PID27_OxySensor4_Volt = uint16(Messages().OBD2.S01PID27_OxySensor4_Volt.FromPhysical(v))
	return m
}

func (m *OBD2) RawS01PID27_OxySensor4_Volt() uint16 {
	return m.xxx_S01PID27_OxySensor4_Volt
}

func (m *OBD2) SetRawS01PID27_OxySensor4_Volt(v uint16) *OBD2 {
	m.xxx_S01PID27_OxySensor4_Volt = uint16(Messages().OBD2.S01PID27_OxySensor4_Volt.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID28_OxySensor5_Volt() float64 {
	return Messages().OBD2.S01PID28_OxySensor5_Volt.ToPhysical(float64(m.xxx_S01PID28_OxySensor5_Volt))
}

func (m *OBD2) SetS01PID28_OxySensor5_Volt(v float64) *OBD2 {
	m.xxx_S01PID28_OxySensor5_Volt = uint16(Messages().OBD2.S01PID28_OxySensor5_Volt.FromPhysical(v))
	return m
}

func (m *OBD2) RawS01PID28_OxySensor5_Volt() uint16 {
	return m.xxx_S01PID28_OxySensor5_Volt
}

func (m *OBD2) SetRawS01PID28_OxySensor5_Volt(v uint16) *OBD2 {
	m.xxx_S01PID28_OxySensor5_Volt = uint16(Messages().OBD2.S01PID28_OxySensor5_Volt.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID29_OxySensor6_Volt() float64 {
	return Messages().OBD2.S01PID29_OxySensor6_Volt.ToPhysical(float64(m.xxx_S01PID29_OxySensor6_Volt))
}

func (m *OBD2) SetS01PID29_OxySensor6_Volt(v float64) *OBD2 {
	m.xxx_S01PID29_OxySensor6_Volt = uint16(Messages().OBD2.S01PID29_OxySensor6_Volt.FromPhysical(v))
	return m
}

func (m *OBD2) RawS01PID29_OxySensor6_Volt() uint16 {
	return m.xxx_S01PID29_OxySensor6_Volt
}

func (m *OBD2) SetRawS01PID29_OxySensor6_Volt(v uint16) *OBD2 {
	m.xxx_S01PID29_OxySensor6_Volt = uint16(Messages().OBD2.S01PID29_OxySensor6_Volt.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID2A_OxySensor7_FAER() float64 {
	return Messages().OBD2.S01PID2A_OxySensor7_FAER.ToPhysical(float64(m.xxx_S01PID2A_OxySensor7_FAER))
}

func (m *OBD2) SetS01PID2A_OxySensor7_FAER(v float64) *OBD2 {
	m.xxx_S01PID2A_OxySensor7_FAER = uint16(Messages().OBD2.S01PID2A_OxySensor7_FAER.FromPhysical(v))
	return m
}

func (m *OBD2) RawS01PID2A_OxySensor7_FAER() uint16 {
	return m.xxx_S01PID2A_OxySensor7_FAER
}

func (m *OBD2) SetRawS01PID2A_OxySensor7_FAER(v uint16) *OBD2 {
	m.xxx_S01PID2A_OxySensor7_FAER = uint16(Messages().OBD2.S01PID2A_OxySensor7_FAER.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID2A_OxySensor7_Volt() float64 {
	return Messages().OBD2.S01PID2A_OxySensor7_Volt.ToPhysical(float64(m.xxx_S01PID2A_OxySensor7_Volt))
}

func (m *OBD2) SetS01PID2A_OxySensor7_Volt(v float64) *OBD2 {
	m.xxx_S01PID2A_OxySensor7_Volt = uint16(Messages().OBD2.S01PID2A_OxySensor7_Volt.FromPhysical(v))
	return m
}

func (m *OBD2) RawS01PID2A_OxySensor7_Volt() uint16 {
	return m.xxx_S01PID2A_OxySensor7_Volt
}

func (m *OBD2) SetRawS01PID2A_OxySensor7_Volt(v uint16) *OBD2 {
	m.xxx_S01PID2A_OxySensor7_Volt = uint16(Messages().OBD2.S01PID2A_OxySensor7_Volt.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID2B_OxySensor8_FAER() float64 {
	return Messages().OBD2.S01PID2B_OxySensor8_FAER.ToPhysical(float64(m.xxx_S01PID2B_OxySensor8_FAER))
}

func (m *OBD2) SetS01PID2B_OxySensor8_FAER(v float64) *OBD2 {
	m.xxx_S01PID2B_OxySensor8_FAER = uint16(Messages().OBD2.S01PID2B_OxySensor8_FAER.FromPhysical(v))
	return m
}

func (m *OBD2) RawS01PID2B_OxySensor8_FAER() uint16 {
	return m.xxx_S01PID2B_OxySensor8_FAER
}

func (m *OBD2) SetRawS01PID2B_OxySensor8_FAER(v uint16) *OBD2 {
	m.xxx_S01PID2B_OxySensor8_FAER = uint16(Messages().OBD2.S01PID2B_OxySensor8_FAER.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID2B_OxySensor8_Volt() float64 {
	return Messages().OBD2.S01PID2B_OxySensor8_Volt.ToPhysical(float64(m.xxx_S01PID2B_OxySensor8_Volt))
}

func (m *OBD2) SetS01PID2B_OxySensor8_Volt(v float64) *OBD2 {
	m.xxx_S01PID2B_OxySensor8_Volt = uint16(Messages().OBD2.S01PID2B_OxySensor8_Volt.FromPhysical(v))
	return m
}

func (m *OBD2) RawS01PID2B_OxySensor8_Volt() uint16 {
	return m.xxx_S01PID2B_OxySensor8_Volt
}

func (m *OBD2) SetRawS01PID2B_OxySensor8_Volt(v uint16) *OBD2 {
	m.xxx_S01PID2B_OxySensor8_Volt = uint16(Messages().OBD2.S01PID2B_OxySensor8_Volt.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID2C_CmdEGR() float64 {
	return Messages().OBD2.S01PID2C_CmdEGR.ToPhysical(float64(m.xxx_S01PID2C_CmdEGR))
}

func (m *OBD2) SetS01PID2C_CmdEGR(v float64) *OBD2 {
	m.xxx_S01PID2C_CmdEGR = uint8(Messages().OBD2.S01PID2C_CmdEGR.FromPhysical(v))
	return m
}

func (m *OBD2) RawS01PID2C_CmdEGR() uint8 {
	return m.xxx_S01PID2C_CmdEGR
}

func (m *OBD2) SetRawS01PID2C_CmdEGR(v uint8) *OBD2 {
	m.xxx_S01PID2C_CmdEGR = uint8(Messages().OBD2.S01PID2C_CmdEGR.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID2D_EGRError() float64 {
	return Messages().OBD2.S01PID2D_EGRError.ToPhysical(float64(m.xxx_S01PID2D_EGRError))
}

func (m *OBD2) SetS01PID2D_EGRError(v float64) *OBD2 {
	m.xxx_S01PID2D_EGRError = uint8(Messages().OBD2.S01PID2D_EGRError.FromPhysical(v))
	return m
}

func (m *OBD2) RawS01PID2D_EGRError() uint8 {
	return m.xxx_S01PID2D_EGRError
}

func (m *OBD2) SetRawS01PID2D_EGRError(v uint8) *OBD2 {
	m.xxx_S01PID2D_EGRError = uint8(Messages().OBD2.S01PID2D_EGRError.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID2E_CmdEvapPurge() float64 {
	return Messages().OBD2.S01PID2E_CmdEvapPurge.ToPhysical(float64(m.xxx_S01PID2E_CmdEvapPurge))
}

func (m *OBD2) SetS01PID2E_CmdEvapPurge(v float64) *OBD2 {
	m.xxx_S01PID2E_CmdEvapPurge = uint8(Messages().OBD2.S01PID2E_CmdEvapPurge.FromPhysical(v))
	return m
}

func (m *OBD2) RawS01PID2E_CmdEvapPurge() uint8 {
	return m.xxx_S01PID2E_CmdEvapPurge
}

func (m *OBD2) SetRawS01PID2E_CmdEvapPurge(v uint8) *OBD2 {
	m.xxx_S01PID2E_CmdEvapPurge = uint8(Messages().OBD2.S01PID2E_CmdEvapPurge.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID2F_FuelTankLevel() float64 {
	return Messages().OBD2.S01PID2F_FuelTankLevel.ToPhysical(float64(m.xxx_S01PID2F_FuelTankLevel))
}

func (m *OBD2) SetS01PID2F_FuelTankLevel(v float64) *OBD2 {
	m.xxx_S01PID2F_FuelTankLevel = uint8(Messages().OBD2.S01PID2F_FuelTankLevel.FromPhysical(v))
	return m
}

func (m *OBD2) RawS01PID2F_FuelTankLevel() uint8 {
	return m.xxx_S01PID2F_FuelTankLevel
}

func (m *OBD2) SetRawS01PID2F_FuelTankLevel(v uint8) *OBD2 {
	m.xxx_S01PID2F_FuelTankLevel = uint8(Messages().OBD2.S01PID2F_FuelTankLevel.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID30_WarmUpsSinceCodeClear() uint8 {
	return m.xxx_S01PID30_WarmUpsSinceCodeClear
}

func (m *OBD2) SetS01PID30_WarmUpsSinceCodeClear(v uint8) *OBD2 {
	m.xxx_S01PID30_WarmUpsSinceCodeClear = uint8(Messages().OBD2.S01PID30_WarmUpsSinceCodeClear.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID31_DistanceSinceCodeClear() uint16 {
	return m.xxx_S01PID31_DistanceSinceCodeClear
}

func (m *OBD2) SetS01PID31_DistanceSinceCodeClear(v uint16) *OBD2 {
	m.xxx_S01PID31_DistanceSinceCodeClear = uint16(Messages().OBD2.S01PID31_DistanceSinceCodeClear.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID32_EvapSysVaporPres() float64 {
	return Messages().OBD2.S01PID32_EvapSysVaporPres.ToPhysical(float64(m.xxx_S01PID32_EvapSysVaporPres))
}

func (m *OBD2) SetS01PID32_EvapSysVaporPres(v float64) *OBD2 {
	m.xxx_S01PID32_EvapSysVaporPres = int16(Messages().OBD2.S01PID32_EvapSysVaporPres.FromPhysical(v))
	return m
}

func (m *OBD2) RawS01PID32_EvapSysVaporPres() int16 {
	return m.xxx_S01PID32_EvapSysVaporPres
}

func (m *OBD2) SetRawS01PID32_EvapSysVaporPres(v int16) *OBD2 {
	m.xxx_S01PID32_EvapSysVaporPres = int16(Messages().OBD2.S01PID32_EvapSysVaporPres.SaturatedCastSigned(int64(v)))
	return m
}

func (m *OBD2) S01PID33_AbsBaroPres() uint8 {
	return m.xxx_S01PID33_AbsBaroPres
}

func (m *OBD2) SetS01PID33_AbsBaroPres(v uint8) *OBD2 {
	m.xxx_S01PID33_AbsBaroPres = uint8(Messages().OBD2.S01PID33_AbsBaroPres.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID34_OxySensor1_FAER() float64 {
	return Messages().OBD2.S01PID34_OxySensor1_FAER.ToPhysical(float64(m.xxx_S01PID34_OxySensor1_FAER))
}

func (m *OBD2) SetS01PID34_OxySensor1_FAER(v float64) *OBD2 {
	m.xxx_S01PID34_OxySensor1_FAER = uint16(Messages().OBD2.S01PID34_OxySensor1_FAER.FromPhysical(v))
	return m
}

func (m *OBD2) RawS01PID34_OxySensor1_FAER() uint16 {
	return m.xxx_S01PID34_OxySensor1_FAER
}

func (m *OBD2) SetRawS01PID34_OxySensor1_FAER(v uint16) *OBD2 {
	m.xxx_S01PID34_OxySensor1_FAER = uint16(Messages().OBD2.S01PID34_OxySensor1_FAER.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID34_OxySensor1_Crnt() float64 {
	return Messages().OBD2.S01PID34_OxySensor1_Crnt.ToPhysical(float64(m.xxx_S01PID34_OxySensor1_Crnt))
}

func (m *OBD2) SetS01PID34_OxySensor1_Crnt(v float64) *OBD2 {
	m.xxx_S01PID34_OxySensor1_Crnt = uint16(Messages().OBD2.S01PID34_OxySensor1_Crnt.FromPhysical(v))
	return m
}

func (m *OBD2) RawS01PID34_OxySensor1_Crnt() uint16 {
	return m.xxx_S01PID34_OxySensor1_Crnt
}

func (m *OBD2) SetRawS01PID34_OxySensor1_Crnt(v uint16) *OBD2 {
	m.xxx_S01PID34_OxySensor1_Crnt = uint16(Messages().OBD2.S01PID34_OxySensor1_Crnt.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID35_OxySensor2_FAER() float64 {
	return Messages().OBD2.S01PID35_OxySensor2_FAER.ToPhysical(float64(m.xxx_S01PID35_OxySensor2_FAER))
}

func (m *OBD2) SetS01PID35_OxySensor2_FAER(v float64) *OBD2 {
	m.xxx_S01PID35_OxySensor2_FAER = uint16(Messages().OBD2.S01PID35_OxySensor2_FAER.FromPhysical(v))
	return m
}

func (m *OBD2) RawS01PID35_OxySensor2_FAER() uint16 {
	return m.xxx_S01PID35_OxySensor2_FAER
}

func (m *OBD2) SetRawS01PID35_OxySensor2_FAER(v uint16) *OBD2 {
	m.xxx_S01PID35_OxySensor2_FAER = uint16(Messages().OBD2.S01PID35_OxySensor2_FAER.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID35_OxySensor2_Crnt() float64 {
	return Messages().OBD2.S01PID35_OxySensor2_Crnt.ToPhysical(float64(m.xxx_S01PID35_OxySensor2_Crnt))
}

func (m *OBD2) SetS01PID35_OxySensor2_Crnt(v float64) *OBD2 {
	m.xxx_S01PID35_OxySensor2_Crnt = uint16(Messages().OBD2.S01PID35_OxySensor2_Crnt.FromPhysical(v))
	return m
}

func (m *OBD2) RawS01PID35_OxySensor2_Crnt() uint16 {
	return m.xxx_S01PID35_OxySensor2_Crnt
}

func (m *OBD2) SetRawS01PID35_OxySensor2_Crnt(v uint16) *OBD2 {
	m.xxx_S01PID35_OxySensor2_Crnt = uint16(Messages().OBD2.S01PID35_OxySensor2_Crnt.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID36_OxySensor3_FAER() float64 {
	return Messages().OBD2.S01PID36_OxySensor3_FAER.ToPhysical(float64(m.xxx_S01PID36_OxySensor3_FAER))
}

func (m *OBD2) SetS01PID36_OxySensor3_FAER(v float64) *OBD2 {
	m.xxx_S01PID36_OxySensor3_FAER = uint16(Messages().OBD2.S01PID36_OxySensor3_FAER.FromPhysical(v))
	return m
}

func (m *OBD2) RawS01PID36_OxySensor3_FAER() uint16 {
	return m.xxx_S01PID36_OxySensor3_FAER
}

func (m *OBD2) SetRawS01PID36_OxySensor3_FAER(v uint16) *OBD2 {
	m.xxx_S01PID36_OxySensor3_FAER = uint16(Messages().OBD2.S01PID36_OxySensor3_FAER.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID36_OxySensor3_Crnt() float64 {
	return Messages().OBD2.S01PID36_OxySensor3_Crnt.ToPhysical(float64(m.xxx_S01PID36_OxySensor3_Crnt))
}

func (m *OBD2) SetS01PID36_OxySensor3_Crnt(v float64) *OBD2 {
	m.xxx_S01PID36_OxySensor3_Crnt = uint16(Messages().OBD2.S01PID36_OxySensor3_Crnt.FromPhysical(v))
	return m
}

func (m *OBD2) RawS01PID36_OxySensor3_Crnt() uint16 {
	return m.xxx_S01PID36_OxySensor3_Crnt
}

func (m *OBD2) SetRawS01PID36_OxySensor3_Crnt(v uint16) *OBD2 {
	m.xxx_S01PID36_OxySensor3_Crnt = uint16(Messages().OBD2.S01PID36_OxySensor3_Crnt.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID37_OxySensor4_FAER() float64 {
	return Messages().OBD2.S01PID37_OxySensor4_FAER.ToPhysical(float64(m.xxx_S01PID37_OxySensor4_FAER))
}

func (m *OBD2) SetS01PID37_OxySensor4_FAER(v float64) *OBD2 {
	m.xxx_S01PID37_OxySensor4_FAER = uint16(Messages().OBD2.S01PID37_OxySensor4_FAER.FromPhysical(v))
	return m
}

func (m *OBD2) RawS01PID37_OxySensor4_FAER() uint16 {
	return m.xxx_S01PID37_OxySensor4_FAER
}

func (m *OBD2) SetRawS01PID37_OxySensor4_FAER(v uint16) *OBD2 {
	m.xxx_S01PID37_OxySensor4_FAER = uint16(Messages().OBD2.S01PID37_OxySensor4_FAER.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID38_OxySensor5_FAER() float64 {
	return Messages().OBD2.S01PID38_OxySensor5_FAER.ToPhysical(float64(m.xxx_S01PID38_OxySensor5_FAER))
}

func (m *OBD2) SetS01PID38_OxySensor5_FAER(v float64) *OBD2 {
	m.xxx_S01PID38_OxySensor5_FAER = uint16(Messages().OBD2.S01PID38_OxySensor5_FAER.FromPhysical(v))
	return m
}

func (m *OBD2) RawS01PID38_OxySensor5_FAER() uint16 {
	return m.xxx_S01PID38_OxySensor5_FAER
}

func (m *OBD2) SetRawS01PID38_OxySensor5_FAER(v uint16) *OBD2 {
	m.xxx_S01PID38_OxySensor5_FAER = uint16(Messages().OBD2.S01PID38_OxySensor5_FAER.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID39_OxySensor6_FAER() float64 {
	return Messages().OBD2.S01PID39_OxySensor6_FAER.ToPhysical(float64(m.xxx_S01PID39_OxySensor6_FAER))
}

func (m *OBD2) SetS01PID39_OxySensor6_FAER(v float64) *OBD2 {
	m.xxx_S01PID39_OxySensor6_FAER = uint16(Messages().OBD2.S01PID39_OxySensor6_FAER.FromPhysical(v))
	return m
}

func (m *OBD2) RawS01PID39_OxySensor6_FAER() uint16 {
	return m.xxx_S01PID39_OxySensor6_FAER
}

func (m *OBD2) SetRawS01PID39_OxySensor6_FAER(v uint16) *OBD2 {
	m.xxx_S01PID39_OxySensor6_FAER = uint16(Messages().OBD2.S01PID39_OxySensor6_FAER.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID37_OxySensor4_Crnt() float64 {
	return Messages().OBD2.S01PID37_OxySensor4_Crnt.ToPhysical(float64(m.xxx_S01PID37_OxySensor4_Crnt))
}

func (m *OBD2) SetS01PID37_OxySensor4_Crnt(v float64) *OBD2 {
	m.xxx_S01PID37_OxySensor4_Crnt = uint16(Messages().OBD2.S01PID37_OxySensor4_Crnt.FromPhysical(v))
	return m
}

func (m *OBD2) RawS01PID37_OxySensor4_Crnt() uint16 {
	return m.xxx_S01PID37_OxySensor4_Crnt
}

func (m *OBD2) SetRawS01PID37_OxySensor4_Crnt(v uint16) *OBD2 {
	m.xxx_S01PID37_OxySensor4_Crnt = uint16(Messages().OBD2.S01PID37_OxySensor4_Crnt.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID38_OxySensor5_Crnt() float64 {
	return Messages().OBD2.S01PID38_OxySensor5_Crnt.ToPhysical(float64(m.xxx_S01PID38_OxySensor5_Crnt))
}

func (m *OBD2) SetS01PID38_OxySensor5_Crnt(v float64) *OBD2 {
	m.xxx_S01PID38_OxySensor5_Crnt = uint16(Messages().OBD2.S01PID38_OxySensor5_Crnt.FromPhysical(v))
	return m
}

func (m *OBD2) RawS01PID38_OxySensor5_Crnt() uint16 {
	return m.xxx_S01PID38_OxySensor5_Crnt
}

func (m *OBD2) SetRawS01PID38_OxySensor5_Crnt(v uint16) *OBD2 {
	m.xxx_S01PID38_OxySensor5_Crnt = uint16(Messages().OBD2.S01PID38_OxySensor5_Crnt.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID39_OxySensor6_Crnt() float64 {
	return Messages().OBD2.S01PID39_OxySensor6_Crnt.ToPhysical(float64(m.xxx_S01PID39_OxySensor6_Crnt))
}

func (m *OBD2) SetS01PID39_OxySensor6_Crnt(v float64) *OBD2 {
	m.xxx_S01PID39_OxySensor6_Crnt = uint16(Messages().OBD2.S01PID39_OxySensor6_Crnt.FromPhysical(v))
	return m
}

func (m *OBD2) RawS01PID39_OxySensor6_Crnt() uint16 {
	return m.xxx_S01PID39_OxySensor6_Crnt
}

func (m *OBD2) SetRawS01PID39_OxySensor6_Crnt(v uint16) *OBD2 {
	m.xxx_S01PID39_OxySensor6_Crnt = uint16(Messages().OBD2.S01PID39_OxySensor6_Crnt.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID3A_OxySensor7_FAER() float64 {
	return Messages().OBD2.S01PID3A_OxySensor7_FAER.ToPhysical(float64(m.xxx_S01PID3A_OxySensor7_FAER))
}

func (m *OBD2) SetS01PID3A_OxySensor7_FAER(v float64) *OBD2 {
	m.xxx_S01PID3A_OxySensor7_FAER = uint16(Messages().OBD2.S01PID3A_OxySensor7_FAER.FromPhysical(v))
	return m
}

func (m *OBD2) RawS01PID3A_OxySensor7_FAER() uint16 {
	return m.xxx_S01PID3A_OxySensor7_FAER
}

func (m *OBD2) SetRawS01PID3A_OxySensor7_FAER(v uint16) *OBD2 {
	m.xxx_S01PID3A_OxySensor7_FAER = uint16(Messages().OBD2.S01PID3A_OxySensor7_FAER.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID3B_OxySensor8_FAER() float64 {
	return Messages().OBD2.S01PID3B_OxySensor8_FAER.ToPhysical(float64(m.xxx_S01PID3B_OxySensor8_FAER))
}

func (m *OBD2) SetS01PID3B_OxySensor8_FAER(v float64) *OBD2 {
	m.xxx_S01PID3B_OxySensor8_FAER = uint16(Messages().OBD2.S01PID3B_OxySensor8_FAER.FromPhysical(v))
	return m
}

func (m *OBD2) RawS01PID3B_OxySensor8_FAER() uint16 {
	return m.xxx_S01PID3B_OxySensor8_FAER
}

func (m *OBD2) SetRawS01PID3B_OxySensor8_FAER(v uint16) *OBD2 {
	m.xxx_S01PID3B_OxySensor8_FAER = uint16(Messages().OBD2.S01PID3B_OxySensor8_FAER.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID3C_CatTempBank1Sens1() float64 {
	return Messages().OBD2.S01PID3C_CatTempBank1Sens1.ToPhysical(float64(m.xxx_S01PID3C_CatTempBank1Sens1))
}

func (m *OBD2) SetS01PID3C_CatTempBank1Sens1(v float64) *OBD2 {
	m.xxx_S01PID3C_CatTempBank1Sens1 = uint16(Messages().OBD2.S01PID3C_CatTempBank1Sens1.FromPhysical(v))
	return m
}

func (m *OBD2) RawS01PID3C_CatTempBank1Sens1() uint16 {
	return m.xxx_S01PID3C_CatTempBank1Sens1
}

func (m *OBD2) SetRawS01PID3C_CatTempBank1Sens1(v uint16) *OBD2 {
	m.xxx_S01PID3C_CatTempBank1Sens1 = uint16(Messages().OBD2.S01PID3C_CatTempBank1Sens1.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID3D_CatTempBank2Sens1() float64 {
	return Messages().OBD2.S01PID3D_CatTempBank2Sens1.ToPhysical(float64(m.xxx_S01PID3D_CatTempBank2Sens1))
}

func (m *OBD2) SetS01PID3D_CatTempBank2Sens1(v float64) *OBD2 {
	m.xxx_S01PID3D_CatTempBank2Sens1 = uint16(Messages().OBD2.S01PID3D_CatTempBank2Sens1.FromPhysical(v))
	return m
}

func (m *OBD2) RawS01PID3D_CatTempBank2Sens1() uint16 {
	return m.xxx_S01PID3D_CatTempBank2Sens1
}

func (m *OBD2) SetRawS01PID3D_CatTempBank2Sens1(v uint16) *OBD2 {
	m.xxx_S01PID3D_CatTempBank2Sens1 = uint16(Messages().OBD2.S01PID3D_CatTempBank2Sens1.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID3A_OxySensor7_Crnt() float64 {
	return Messages().OBD2.S01PID3A_OxySensor7_Crnt.ToPhysical(float64(m.xxx_S01PID3A_OxySensor7_Crnt))
}

func (m *OBD2) SetS01PID3A_OxySensor7_Crnt(v float64) *OBD2 {
	m.xxx_S01PID3A_OxySensor7_Crnt = uint16(Messages().OBD2.S01PID3A_OxySensor7_Crnt.FromPhysical(v))
	return m
}

func (m *OBD2) RawS01PID3A_OxySensor7_Crnt() uint16 {
	return m.xxx_S01PID3A_OxySensor7_Crnt
}

func (m *OBD2) SetRawS01PID3A_OxySensor7_Crnt(v uint16) *OBD2 {
	m.xxx_S01PID3A_OxySensor7_Crnt = uint16(Messages().OBD2.S01PID3A_OxySensor7_Crnt.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID3B_OxySensor8_Crnt() float64 {
	return Messages().OBD2.S01PID3B_OxySensor8_Crnt.ToPhysical(float64(m.xxx_S01PID3B_OxySensor8_Crnt))
}

func (m *OBD2) SetS01PID3B_OxySensor8_Crnt(v float64) *OBD2 {
	m.xxx_S01PID3B_OxySensor8_Crnt = uint16(Messages().OBD2.S01PID3B_OxySensor8_Crnt.FromPhysical(v))
	return m
}

func (m *OBD2) RawS01PID3B_OxySensor8_Crnt() uint16 {
	return m.xxx_S01PID3B_OxySensor8_Crnt
}

func (m *OBD2) SetRawS01PID3B_OxySensor8_Crnt(v uint16) *OBD2 {
	m.xxx_S01PID3B_OxySensor8_Crnt = uint16(Messages().OBD2.S01PID3B_OxySensor8_Crnt.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID3E_CatTempBank1Sens2() float64 {
	return Messages().OBD2.S01PID3E_CatTempBank1Sens2.ToPhysical(float64(m.xxx_S01PID3E_CatTempBank1Sens2))
}

func (m *OBD2) SetS01PID3E_CatTempBank1Sens2(v float64) *OBD2 {
	m.xxx_S01PID3E_CatTempBank1Sens2 = uint16(Messages().OBD2.S01PID3E_CatTempBank1Sens2.FromPhysical(v))
	return m
}

func (m *OBD2) RawS01PID3E_CatTempBank1Sens2() uint16 {
	return m.xxx_S01PID3E_CatTempBank1Sens2
}

func (m *OBD2) SetRawS01PID3E_CatTempBank1Sens2(v uint16) *OBD2 {
	m.xxx_S01PID3E_CatTempBank1Sens2 = uint16(Messages().OBD2.S01PID3E_CatTempBank1Sens2.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID3F_CatTempBank2Sens2() float64 {
	return Messages().OBD2.S01PID3F_CatTempBank2Sens2.ToPhysical(float64(m.xxx_S01PID3F_CatTempBank2Sens2))
}

func (m *OBD2) SetS01PID3F_CatTempBank2Sens2(v float64) *OBD2 {
	m.xxx_S01PID3F_CatTempBank2Sens2 = uint16(Messages().OBD2.S01PID3F_CatTempBank2Sens2.FromPhysical(v))
	return m
}

func (m *OBD2) RawS01PID3F_CatTempBank2Sens2() uint16 {
	return m.xxx_S01PID3F_CatTempBank2Sens2
}

func (m *OBD2) SetRawS01PID3F_CatTempBank2Sens2(v uint16) *OBD2 {
	m.xxx_S01PID3F_CatTempBank2Sens2 = uint16(Messages().OBD2.S01PID3F_CatTempBank2Sens2.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID40_PIDsSupported_41_60() uint32 {
	return m.xxx_S01PID40_PIDsSupported_41_60
}

func (m *OBD2) SetS01PID40_PIDsSupported_41_60(v uint32) *OBD2 {
	m.xxx_S01PID40_PIDsSupported_41_60 = uint32(Messages().OBD2.S01PID40_PIDsSupported_41_60.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID41_MonStatusDriveCycle() uint32 {
	return m.xxx_S01PID41_MonStatusDriveCycle
}

func (m *OBD2) SetS01PID41_MonStatusDriveCycle(v uint32) *OBD2 {
	m.xxx_S01PID41_MonStatusDriveCycle = uint32(Messages().OBD2.S01PID41_MonStatusDriveCycle.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID42_ControlModuleVolt() float64 {
	return Messages().OBD2.S01PID42_ControlModuleVolt.ToPhysical(float64(m.xxx_S01PID42_ControlModuleVolt))
}

func (m *OBD2) SetS01PID42_ControlModuleVolt(v float64) *OBD2 {
	m.xxx_S01PID42_ControlModuleVolt = uint16(Messages().OBD2.S01PID42_ControlModuleVolt.FromPhysical(v))
	return m
}

func (m *OBD2) RawS01PID42_ControlModuleVolt() uint16 {
	return m.xxx_S01PID42_ControlModuleVolt
}

func (m *OBD2) SetRawS01PID42_ControlModuleVolt(v uint16) *OBD2 {
	m.xxx_S01PID42_ControlModuleVolt = uint16(Messages().OBD2.S01PID42_ControlModuleVolt.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID43_AbsLoadValue() float64 {
	return Messages().OBD2.S01PID43_AbsLoadValue.ToPhysical(float64(m.xxx_S01PID43_AbsLoadValue))
}

func (m *OBD2) SetS01PID43_AbsLoadValue(v float64) *OBD2 {
	m.xxx_S01PID43_AbsLoadValue = uint16(Messages().OBD2.S01PID43_AbsLoadValue.FromPhysical(v))
	return m
}

func (m *OBD2) RawS01PID43_AbsLoadValue() uint16 {
	return m.xxx_S01PID43_AbsLoadValue
}

func (m *OBD2) SetRawS01PID43_AbsLoadValue(v uint16) *OBD2 {
	m.xxx_S01PID43_AbsLoadValue = uint16(Messages().OBD2.S01PID43_AbsLoadValue.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID44_FuelAirCmdEquiv() float64 {
	return Messages().OBD2.S01PID44_FuelAirCmdEquiv.ToPhysical(float64(m.xxx_S01PID44_FuelAirCmdEquiv))
}

func (m *OBD2) SetS01PID44_FuelAirCmdEquiv(v float64) *OBD2 {
	m.xxx_S01PID44_FuelAirCmdEquiv = uint16(Messages().OBD2.S01PID44_FuelAirCmdEquiv.FromPhysical(v))
	return m
}

func (m *OBD2) RawS01PID44_FuelAirCmdEquiv() uint16 {
	return m.xxx_S01PID44_FuelAirCmdEquiv
}

func (m *OBD2) SetRawS01PID44_FuelAirCmdEquiv(v uint16) *OBD2 {
	m.xxx_S01PID44_FuelAirCmdEquiv = uint16(Messages().OBD2.S01PID44_FuelAirCmdEquiv.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID45_RelThrottlePos() float64 {
	return Messages().OBD2.S01PID45_RelThrottlePos.ToPhysical(float64(m.xxx_S01PID45_RelThrottlePos))
}

func (m *OBD2) SetS01PID45_RelThrottlePos(v float64) *OBD2 {
	m.xxx_S01PID45_RelThrottlePos = uint8(Messages().OBD2.S01PID45_RelThrottlePos.FromPhysical(v))
	return m
}

func (m *OBD2) RawS01PID45_RelThrottlePos() uint8 {
	return m.xxx_S01PID45_RelThrottlePos
}

func (m *OBD2) SetRawS01PID45_RelThrottlePos(v uint8) *OBD2 {
	m.xxx_S01PID45_RelThrottlePos = uint8(Messages().OBD2.S01PID45_RelThrottlePos.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID46_AmbientAirTemp() float64 {
	return Messages().OBD2.S01PID46_AmbientAirTemp.ToPhysical(float64(m.xxx_S01PID46_AmbientAirTemp))
}

func (m *OBD2) SetS01PID46_AmbientAirTemp(v float64) *OBD2 {
	m.xxx_S01PID46_AmbientAirTemp = uint8(Messages().OBD2.S01PID46_AmbientAirTemp.FromPhysical(v))
	return m
}

func (m *OBD2) RawS01PID46_AmbientAirTemp() uint8 {
	return m.xxx_S01PID46_AmbientAirTemp
}

func (m *OBD2) SetRawS01PID46_AmbientAirTemp(v uint8) *OBD2 {
	m.xxx_S01PID46_AmbientAirTemp = uint8(Messages().OBD2.S01PID46_AmbientAirTemp.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID47_AbsThrottlePosB() float64 {
	return Messages().OBD2.S01PID47_AbsThrottlePosB.ToPhysical(float64(m.xxx_S01PID47_AbsThrottlePosB))
}

func (m *OBD2) SetS01PID47_AbsThrottlePosB(v float64) *OBD2 {
	m.xxx_S01PID47_AbsThrottlePosB = uint8(Messages().OBD2.S01PID47_AbsThrottlePosB.FromPhysical(v))
	return m
}

func (m *OBD2) RawS01PID47_AbsThrottlePosB() uint8 {
	return m.xxx_S01PID47_AbsThrottlePosB
}

func (m *OBD2) SetRawS01PID47_AbsThrottlePosB(v uint8) *OBD2 {
	m.xxx_S01PID47_AbsThrottlePosB = uint8(Messages().OBD2.S01PID47_AbsThrottlePosB.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID48_AbsThrottlePosC() float64 {
	return Messages().OBD2.S01PID48_AbsThrottlePosC.ToPhysical(float64(m.xxx_S01PID48_AbsThrottlePosC))
}

func (m *OBD2) SetS01PID48_AbsThrottlePosC(v float64) *OBD2 {
	m.xxx_S01PID48_AbsThrottlePosC = uint8(Messages().OBD2.S01PID48_AbsThrottlePosC.FromPhysical(v))
	return m
}

func (m *OBD2) RawS01PID48_AbsThrottlePosC() uint8 {
	return m.xxx_S01PID48_AbsThrottlePosC
}

func (m *OBD2) SetRawS01PID48_AbsThrottlePosC(v uint8) *OBD2 {
	m.xxx_S01PID48_AbsThrottlePosC = uint8(Messages().OBD2.S01PID48_AbsThrottlePosC.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID49_AbsThrottlePosD() float64 {
	return Messages().OBD2.S01PID49_AbsThrottlePosD.ToPhysical(float64(m.xxx_S01PID49_AbsThrottlePosD))
}

func (m *OBD2) SetS01PID49_AbsThrottlePosD(v float64) *OBD2 {
	m.xxx_S01PID49_AbsThrottlePosD = uint8(Messages().OBD2.S01PID49_AbsThrottlePosD.FromPhysical(v))
	return m
}

func (m *OBD2) RawS01PID49_AbsThrottlePosD() uint8 {
	return m.xxx_S01PID49_AbsThrottlePosD
}

func (m *OBD2) SetRawS01PID49_AbsThrottlePosD(v uint8) *OBD2 {
	m.xxx_S01PID49_AbsThrottlePosD = uint8(Messages().OBD2.S01PID49_AbsThrottlePosD.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID4A_AbsThrottlePosE() float64 {
	return Messages().OBD2.S01PID4A_AbsThrottlePosE.ToPhysical(float64(m.xxx_S01PID4A_AbsThrottlePosE))
}

func (m *OBD2) SetS01PID4A_AbsThrottlePosE(v float64) *OBD2 {
	m.xxx_S01PID4A_AbsThrottlePosE = uint8(Messages().OBD2.S01PID4A_AbsThrottlePosE.FromPhysical(v))
	return m
}

func (m *OBD2) RawS01PID4A_AbsThrottlePosE() uint8 {
	return m.xxx_S01PID4A_AbsThrottlePosE
}

func (m *OBD2) SetRawS01PID4A_AbsThrottlePosE(v uint8) *OBD2 {
	m.xxx_S01PID4A_AbsThrottlePosE = uint8(Messages().OBD2.S01PID4A_AbsThrottlePosE.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID4B_AbsThrottlePosF() float64 {
	return Messages().OBD2.S01PID4B_AbsThrottlePosF.ToPhysical(float64(m.xxx_S01PID4B_AbsThrottlePosF))
}

func (m *OBD2) SetS01PID4B_AbsThrottlePosF(v float64) *OBD2 {
	m.xxx_S01PID4B_AbsThrottlePosF = uint8(Messages().OBD2.S01PID4B_AbsThrottlePosF.FromPhysical(v))
	return m
}

func (m *OBD2) RawS01PID4B_AbsThrottlePosF() uint8 {
	return m.xxx_S01PID4B_AbsThrottlePosF
}

func (m *OBD2) SetRawS01PID4B_AbsThrottlePosF(v uint8) *OBD2 {
	m.xxx_S01PID4B_AbsThrottlePosF = uint8(Messages().OBD2.S01PID4B_AbsThrottlePosF.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID4C_CmdThrottleAct() float64 {
	return Messages().OBD2.S01PID4C_CmdThrottleAct.ToPhysical(float64(m.xxx_S01PID4C_CmdThrottleAct))
}

func (m *OBD2) SetS01PID4C_CmdThrottleAct(v float64) *OBD2 {
	m.xxx_S01PID4C_CmdThrottleAct = uint8(Messages().OBD2.S01PID4C_CmdThrottleAct.FromPhysical(v))
	return m
}

func (m *OBD2) RawS01PID4C_CmdThrottleAct() uint8 {
	return m.xxx_S01PID4C_CmdThrottleAct
}

func (m *OBD2) SetRawS01PID4C_CmdThrottleAct(v uint8) *OBD2 {
	m.xxx_S01PID4C_CmdThrottleAct = uint8(Messages().OBD2.S01PID4C_CmdThrottleAct.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID4D_TimeRunMILOn() uint16 {
	return m.xxx_S01PID4D_TimeRunMILOn
}

func (m *OBD2) SetS01PID4D_TimeRunMILOn(v uint16) *OBD2 {
	m.xxx_S01PID4D_TimeRunMILOn = uint16(Messages().OBD2.S01PID4D_TimeRunMILOn.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID4E_TimeSinceCodeClear() uint16 {
	return m.xxx_S01PID4E_TimeSinceCodeClear
}

func (m *OBD2) SetS01PID4E_TimeSinceCodeClear(v uint16) *OBD2 {
	m.xxx_S01PID4E_TimeSinceCodeClear = uint16(Messages().OBD2.S01PID4E_TimeSinceCodeClear.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID4F_Max_FAER() uint8 {
	return m.xxx_S01PID4F_Max_FAER
}

func (m *OBD2) SetS01PID4F_Max_FAER(v uint8) *OBD2 {
	m.xxx_S01PID4F_Max_FAER = uint8(Messages().OBD2.S01PID4F_Max_FAER.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID4F_Max_OxySensVol() uint8 {
	return m.xxx_S01PID4F_Max_OxySensVol
}

func (m *OBD2) SetS01PID4F_Max_OxySensVol(v uint8) *OBD2 {
	m.xxx_S01PID4F_Max_OxySensVol = uint8(Messages().OBD2.S01PID4F_Max_OxySensVol.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID4F_Max_OxySensCrnt() uint8 {
	return m.xxx_S01PID4F_Max_OxySensCrnt
}

func (m *OBD2) SetS01PID4F_Max_OxySensCrnt(v uint8) *OBD2 {
	m.xxx_S01PID4F_Max_OxySensCrnt = uint8(Messages().OBD2.S01PID4F_Max_OxySensCrnt.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID4F_Max_IntManiAbsPres() float64 {
	return Messages().OBD2.S01PID4F_Max_IntManiAbsPres.ToPhysical(float64(m.xxx_S01PID4F_Max_IntManiAbsPres))
}

func (m *OBD2) SetS01PID4F_Max_IntManiAbsPres(v float64) *OBD2 {
	m.xxx_S01PID4F_Max_IntManiAbsPres = uint8(Messages().OBD2.S01PID4F_Max_IntManiAbsPres.FromPhysical(v))
	return m
}

func (m *OBD2) RawS01PID4F_Max_IntManiAbsPres() uint8 {
	return m.xxx_S01PID4F_Max_IntManiAbsPres
}

func (m *OBD2) SetRawS01PID4F_Max_IntManiAbsPres(v uint8) *OBD2 {
	m.xxx_S01PID4F_Max_IntManiAbsPres = uint8(Messages().OBD2.S01PID4F_Max_IntManiAbsPres.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID50_Max_AirFlowMAF() float64 {
	return Messages().OBD2.S01PID50_Max_AirFlowMAF.ToPhysical(float64(m.xxx_S01PID50_Max_AirFlowMAF))
}

func (m *OBD2) SetS01PID50_Max_AirFlowMAF(v float64) *OBD2 {
	m.xxx_S01PID50_Max_AirFlowMAF = uint8(Messages().OBD2.S01PID50_Max_AirFlowMAF.FromPhysical(v))
	return m
}

func (m *OBD2) RawS01PID50_Max_AirFlowMAF() uint8 {
	return m.xxx_S01PID50_Max_AirFlowMAF
}

func (m *OBD2) SetRawS01PID50_Max_AirFlowMAF(v uint8) *OBD2 {
	m.xxx_S01PID50_Max_AirFlowMAF = uint8(Messages().OBD2.S01PID50_Max_AirFlowMAF.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID51_FuelType() OBD2_S01PID51_FuelType {
	return m.xxx_S01PID51_FuelType
}

func (m *OBD2) SetS01PID51_FuelType(v OBD2_S01PID51_FuelType) *OBD2 {
	m.xxx_S01PID51_FuelType = OBD2_S01PID51_FuelType(Messages().OBD2.S01PID51_FuelType.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID52_EthanolFuelPct() float64 {
	return Messages().OBD2.S01PID52_EthanolFuelPct.ToPhysical(float64(m.xxx_S01PID52_EthanolFuelPct))
}

func (m *OBD2) SetS01PID52_EthanolFuelPct(v float64) *OBD2 {
	m.xxx_S01PID52_EthanolFuelPct = uint8(Messages().OBD2.S01PID52_EthanolFuelPct.FromPhysical(v))
	return m
}

func (m *OBD2) RawS01PID52_EthanolFuelPct() uint8 {
	return m.xxx_S01PID52_EthanolFuelPct
}

func (m *OBD2) SetRawS01PID52_EthanolFuelPct(v uint8) *OBD2 {
	m.xxx_S01PID52_EthanolFuelPct = uint8(Messages().OBD2.S01PID52_EthanolFuelPct.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID53_AbsEvapSysVapPres() float64 {
	return Messages().OBD2.S01PID53_AbsEvapSysVapPres.ToPhysical(float64(m.xxx_S01PID53_AbsEvapSysVapPres))
}

func (m *OBD2) SetS01PID53_AbsEvapSysVapPres(v float64) *OBD2 {
	m.xxx_S01PID53_AbsEvapSysVapPres = uint16(Messages().OBD2.S01PID53_AbsEvapSysVapPres.FromPhysical(v))
	return m
}

func (m *OBD2) RawS01PID53_AbsEvapSysVapPres() uint16 {
	return m.xxx_S01PID53_AbsEvapSysVapPres
}

func (m *OBD2) SetRawS01PID53_AbsEvapSysVapPres(v uint16) *OBD2 {
	m.xxx_S01PID53_AbsEvapSysVapPres = uint16(Messages().OBD2.S01PID53_AbsEvapSysVapPres.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID54_EvapSysVapPres() float64 {
	return Messages().OBD2.S01PID54_EvapSysVapPres.ToPhysical(float64(m.xxx_S01PID54_EvapSysVapPres))
}

func (m *OBD2) SetS01PID54_EvapSysVapPres(v float64) *OBD2 {
	m.xxx_S01PID54_EvapSysVapPres = uint16(Messages().OBD2.S01PID54_EvapSysVapPres.FromPhysical(v))
	return m
}

func (m *OBD2) RawS01PID54_EvapSysVapPres() uint16 {
	return m.xxx_S01PID54_EvapSysVapPres
}

func (m *OBD2) SetRawS01PID54_EvapSysVapPres(v uint16) *OBD2 {
	m.xxx_S01PID54_EvapSysVapPres = uint16(Messages().OBD2.S01PID54_EvapSysVapPres.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID55_ShortSecOxyTrimBank1() float64 {
	return Messages().OBD2.S01PID55_ShortSecOxyTrimBank1.ToPhysical(float64(m.xxx_S01PID55_ShortSecOxyTrimBank1))
}

func (m *OBD2) SetS01PID55_ShortSecOxyTrimBank1(v float64) *OBD2 {
	m.xxx_S01PID55_ShortSecOxyTrimBank1 = uint8(Messages().OBD2.S01PID55_ShortSecOxyTrimBank1.FromPhysical(v))
	return m
}

func (m *OBD2) RawS01PID55_ShortSecOxyTrimBank1() uint8 {
	return m.xxx_S01PID55_ShortSecOxyTrimBank1
}

func (m *OBD2) SetRawS01PID55_ShortSecOxyTrimBank1(v uint8) *OBD2 {
	m.xxx_S01PID55_ShortSecOxyTrimBank1 = uint8(Messages().OBD2.S01PID55_ShortSecOxyTrimBank1.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID56_LongSecOxyTrimBank1() float64 {
	return Messages().OBD2.S01PID56_LongSecOxyTrimBank1.ToPhysical(float64(m.xxx_S01PID56_LongSecOxyTrimBank1))
}

func (m *OBD2) SetS01PID56_LongSecOxyTrimBank1(v float64) *OBD2 {
	m.xxx_S01PID56_LongSecOxyTrimBank1 = uint8(Messages().OBD2.S01PID56_LongSecOxyTrimBank1.FromPhysical(v))
	return m
}

func (m *OBD2) RawS01PID56_LongSecOxyTrimBank1() uint8 {
	return m.xxx_S01PID56_LongSecOxyTrimBank1
}

func (m *OBD2) SetRawS01PID56_LongSecOxyTrimBank1(v uint8) *OBD2 {
	m.xxx_S01PID56_LongSecOxyTrimBank1 = uint8(Messages().OBD2.S01PID56_LongSecOxyTrimBank1.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID55_ShortSecOxyTrimBank3() float64 {
	return Messages().OBD2.S01PID55_ShortSecOxyTrimBank3.ToPhysical(float64(m.xxx_S01PID55_ShortSecOxyTrimBank3))
}

func (m *OBD2) SetS01PID55_ShortSecOxyTrimBank3(v float64) *OBD2 {
	m.xxx_S01PID55_ShortSecOxyTrimBank3 = uint8(Messages().OBD2.S01PID55_ShortSecOxyTrimBank3.FromPhysical(v))
	return m
}

func (m *OBD2) RawS01PID55_ShortSecOxyTrimBank3() uint8 {
	return m.xxx_S01PID55_ShortSecOxyTrimBank3
}

func (m *OBD2) SetRawS01PID55_ShortSecOxyTrimBank3(v uint8) *OBD2 {
	m.xxx_S01PID55_ShortSecOxyTrimBank3 = uint8(Messages().OBD2.S01PID55_ShortSecOxyTrimBank3.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID56_LongSecOxyTrimBank3() float64 {
	return Messages().OBD2.S01PID56_LongSecOxyTrimBank3.ToPhysical(float64(m.xxx_S01PID56_LongSecOxyTrimBank3))
}

func (m *OBD2) SetS01PID56_LongSecOxyTrimBank3(v float64) *OBD2 {
	m.xxx_S01PID56_LongSecOxyTrimBank3 = uint8(Messages().OBD2.S01PID56_LongSecOxyTrimBank3.FromPhysical(v))
	return m
}

func (m *OBD2) RawS01PID56_LongSecOxyTrimBank3() uint8 {
	return m.xxx_S01PID56_LongSecOxyTrimBank3
}

func (m *OBD2) SetRawS01PID56_LongSecOxyTrimBank3(v uint8) *OBD2 {
	m.xxx_S01PID56_LongSecOxyTrimBank3 = uint8(Messages().OBD2.S01PID56_LongSecOxyTrimBank3.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID57_ShortSecOxyTrimBank2() float64 {
	return Messages().OBD2.S01PID57_ShortSecOxyTrimBank2.ToPhysical(float64(m.xxx_S01PID57_ShortSecOxyTrimBank2))
}

func (m *OBD2) SetS01PID57_ShortSecOxyTrimBank2(v float64) *OBD2 {
	m.xxx_S01PID57_ShortSecOxyTrimBank2 = uint8(Messages().OBD2.S01PID57_ShortSecOxyTrimBank2.FromPhysical(v))
	return m
}

func (m *OBD2) RawS01PID57_ShortSecOxyTrimBank2() uint8 {
	return m.xxx_S01PID57_ShortSecOxyTrimBank2
}

func (m *OBD2) SetRawS01PID57_ShortSecOxyTrimBank2(v uint8) *OBD2 {
	m.xxx_S01PID57_ShortSecOxyTrimBank2 = uint8(Messages().OBD2.S01PID57_ShortSecOxyTrimBank2.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID58_LongSecOxyTrimBank2() float64 {
	return Messages().OBD2.S01PID58_LongSecOxyTrimBank2.ToPhysical(float64(m.xxx_S01PID58_LongSecOxyTrimBank2))
}

func (m *OBD2) SetS01PID58_LongSecOxyTrimBank2(v float64) *OBD2 {
	m.xxx_S01PID58_LongSecOxyTrimBank2 = uint8(Messages().OBD2.S01PID58_LongSecOxyTrimBank2.FromPhysical(v))
	return m
}

func (m *OBD2) RawS01PID58_LongSecOxyTrimBank2() uint8 {
	return m.xxx_S01PID58_LongSecOxyTrimBank2
}

func (m *OBD2) SetRawS01PID58_LongSecOxyTrimBank2(v uint8) *OBD2 {
	m.xxx_S01PID58_LongSecOxyTrimBank2 = uint8(Messages().OBD2.S01PID58_LongSecOxyTrimBank2.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID59_FuelRailAbsPres() float64 {
	return Messages().OBD2.S01PID59_FuelRailAbsPres.ToPhysical(float64(m.xxx_S01PID59_FuelRailAbsPres))
}

func (m *OBD2) SetS01PID59_FuelRailAbsPres(v float64) *OBD2 {
	m.xxx_S01PID59_FuelRailAbsPres = uint16(Messages().OBD2.S01PID59_FuelRailAbsPres.FromPhysical(v))
	return m
}

func (m *OBD2) RawS01PID59_FuelRailAbsPres() uint16 {
	return m.xxx_S01PID59_FuelRailAbsPres
}

func (m *OBD2) SetRawS01PID59_FuelRailAbsPres(v uint16) *OBD2 {
	m.xxx_S01PID59_FuelRailAbsPres = uint16(Messages().OBD2.S01PID59_FuelRailAbsPres.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID5A_RelAccelPedalPos() float64 {
	return Messages().OBD2.S01PID5A_RelAccelPedalPos.ToPhysical(float64(m.xxx_S01PID5A_RelAccelPedalPos))
}

func (m *OBD2) SetS01PID5A_RelAccelPedalPos(v float64) *OBD2 {
	m.xxx_S01PID5A_RelAccelPedalPos = uint8(Messages().OBD2.S01PID5A_RelAccelPedalPos.FromPhysical(v))
	return m
}

func (m *OBD2) RawS01PID5A_RelAccelPedalPos() uint8 {
	return m.xxx_S01PID5A_RelAccelPedalPos
}

func (m *OBD2) SetRawS01PID5A_RelAccelPedalPos(v uint8) *OBD2 {
	m.xxx_S01PID5A_RelAccelPedalPos = uint8(Messages().OBD2.S01PID5A_RelAccelPedalPos.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID57_ShortSecOxyTrimBank4() float64 {
	return Messages().OBD2.S01PID57_ShortSecOxyTrimBank4.ToPhysical(float64(m.xxx_S01PID57_ShortSecOxyTrimBank4))
}

func (m *OBD2) SetS01PID57_ShortSecOxyTrimBank4(v float64) *OBD2 {
	m.xxx_S01PID57_ShortSecOxyTrimBank4 = uint8(Messages().OBD2.S01PID57_ShortSecOxyTrimBank4.FromPhysical(v))
	return m
}

func (m *OBD2) RawS01PID57_ShortSecOxyTrimBank4() uint8 {
	return m.xxx_S01PID57_ShortSecOxyTrimBank4
}

func (m *OBD2) SetRawS01PID57_ShortSecOxyTrimBank4(v uint8) *OBD2 {
	m.xxx_S01PID57_ShortSecOxyTrimBank4 = uint8(Messages().OBD2.S01PID57_ShortSecOxyTrimBank4.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID58_LongSecOxyTrimBank4() float64 {
	return Messages().OBD2.S01PID58_LongSecOxyTrimBank4.ToPhysical(float64(m.xxx_S01PID58_LongSecOxyTrimBank4))
}

func (m *OBD2) SetS01PID58_LongSecOxyTrimBank4(v float64) *OBD2 {
	m.xxx_S01PID58_LongSecOxyTrimBank4 = uint8(Messages().OBD2.S01PID58_LongSecOxyTrimBank4.FromPhysical(v))
	return m
}

func (m *OBD2) RawS01PID58_LongSecOxyTrimBank4() uint8 {
	return m.xxx_S01PID58_LongSecOxyTrimBank4
}

func (m *OBD2) SetRawS01PID58_LongSecOxyTrimBank4(v uint8) *OBD2 {
	m.xxx_S01PID58_LongSecOxyTrimBank4 = uint8(Messages().OBD2.S01PID58_LongSecOxyTrimBank4.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID5B_HybrBatPackRemLife() float64 {
	return Messages().OBD2.S01PID5B_HybrBatPackRemLife.ToPhysical(float64(m.xxx_S01PID5B_HybrBatPackRemLife))
}

func (m *OBD2) SetS01PID5B_HybrBatPackRemLife(v float64) *OBD2 {
	m.xxx_S01PID5B_HybrBatPackRemLife = uint8(Messages().OBD2.S01PID5B_HybrBatPackRemLife.FromPhysical(v))
	return m
}

func (m *OBD2) RawS01PID5B_HybrBatPackRemLife() uint8 {
	return m.xxx_S01PID5B_HybrBatPackRemLife
}

func (m *OBD2) SetRawS01PID5B_HybrBatPackRemLife(v uint8) *OBD2 {
	m.xxx_S01PID5B_HybrBatPackRemLife = uint8(Messages().OBD2.S01PID5B_HybrBatPackRemLife.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID5C_EngineOilTemp() float64 {
	return Messages().OBD2.S01PID5C_EngineOilTemp.ToPhysical(float64(m.xxx_S01PID5C_EngineOilTemp))
}

func (m *OBD2) SetS01PID5C_EngineOilTemp(v float64) *OBD2 {
	m.xxx_S01PID5C_EngineOilTemp = uint8(Messages().OBD2.S01PID5C_EngineOilTemp.FromPhysical(v))
	return m
}

func (m *OBD2) RawS01PID5C_EngineOilTemp() uint8 {
	return m.xxx_S01PID5C_EngineOilTemp
}

func (m *OBD2) SetRawS01PID5C_EngineOilTemp(v uint8) *OBD2 {
	m.xxx_S01PID5C_EngineOilTemp = uint8(Messages().OBD2.S01PID5C_EngineOilTemp.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID5D_FuelInjectionTiming() float64 {
	return Messages().OBD2.S01PID5D_FuelInjectionTiming.ToPhysical(float64(m.xxx_S01PID5D_FuelInjectionTiming))
}

func (m *OBD2) SetS01PID5D_FuelInjectionTiming(v float64) *OBD2 {
	m.xxx_S01PID5D_FuelInjectionTiming = uint16(Messages().OBD2.S01PID5D_FuelInjectionTiming.FromPhysical(v))
	return m
}

func (m *OBD2) RawS01PID5D_FuelInjectionTiming() uint16 {
	return m.xxx_S01PID5D_FuelInjectionTiming
}

func (m *OBD2) SetRawS01PID5D_FuelInjectionTiming(v uint16) *OBD2 {
	m.xxx_S01PID5D_FuelInjectionTiming = uint16(Messages().OBD2.S01PID5D_FuelInjectionTiming.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID5E_EngineFuelRate() float64 {
	return Messages().OBD2.S01PID5E_EngineFuelRate.ToPhysical(float64(m.xxx_S01PID5E_EngineFuelRate))
}

func (m *OBD2) SetS01PID5E_EngineFuelRate(v float64) *OBD2 {
	m.xxx_S01PID5E_EngineFuelRate = uint16(Messages().OBD2.S01PID5E_EngineFuelRate.FromPhysical(v))
	return m
}

func (m *OBD2) RawS01PID5E_EngineFuelRate() uint16 {
	return m.xxx_S01PID5E_EngineFuelRate
}

func (m *OBD2) SetRawS01PID5E_EngineFuelRate(v uint16) *OBD2 {
	m.xxx_S01PID5E_EngineFuelRate = uint16(Messages().OBD2.S01PID5E_EngineFuelRate.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID5F_EmissionReq() uint8 {
	return m.xxx_S01PID5F_EmissionReq
}

func (m *OBD2) SetS01PID5F_EmissionReq(v uint8) *OBD2 {
	m.xxx_S01PID5F_EmissionReq = uint8(Messages().OBD2.S01PID5F_EmissionReq.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID60_PIDsSupported_61_80() uint32 {
	return m.xxx_S01PID60_PIDsSupported_61_80
}

func (m *OBD2) SetS01PID60_PIDsSupported_61_80(v uint32) *OBD2 {
	m.xxx_S01PID60_PIDsSupported_61_80 = uint32(Messages().OBD2.S01PID60_PIDsSupported_61_80.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID61_DemandEngTorqPct() float64 {
	return Messages().OBD2.S01PID61_DemandEngTorqPct.ToPhysical(float64(m.xxx_S01PID61_DemandEngTorqPct))
}

func (m *OBD2) SetS01PID61_DemandEngTorqPct(v float64) *OBD2 {
	m.xxx_S01PID61_DemandEngTorqPct = uint8(Messages().OBD2.S01PID61_DemandEngTorqPct.FromPhysical(v))
	return m
}

func (m *OBD2) RawS01PID61_DemandEngTorqPct() uint8 {
	return m.xxx_S01PID61_DemandEngTorqPct
}

func (m *OBD2) SetRawS01PID61_DemandEngTorqPct(v uint8) *OBD2 {
	m.xxx_S01PID61_DemandEngTorqPct = uint8(Messages().OBD2.S01PID61_DemandEngTorqPct.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID62_ActualEngTorqPct() float64 {
	return Messages().OBD2.S01PID62_ActualEngTorqPct.ToPhysical(float64(m.xxx_S01PID62_ActualEngTorqPct))
}

func (m *OBD2) SetS01PID62_ActualEngTorqPct(v float64) *OBD2 {
	m.xxx_S01PID62_ActualEngTorqPct = uint8(Messages().OBD2.S01PID62_ActualEngTorqPct.FromPhysical(v))
	return m
}

func (m *OBD2) RawS01PID62_ActualEngTorqPct() uint8 {
	return m.xxx_S01PID62_ActualEngTorqPct
}

func (m *OBD2) SetRawS01PID62_ActualEngTorqPct(v uint8) *OBD2 {
	m.xxx_S01PID62_ActualEngTorqPct = uint8(Messages().OBD2.S01PID62_ActualEngTorqPct.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID63_EngRefTorq() uint16 {
	return m.xxx_S01PID63_EngRefTorq
}

func (m *OBD2) SetS01PID63_EngRefTorq(v uint16) *OBD2 {
	m.xxx_S01PID63_EngRefTorq = uint16(Messages().OBD2.S01PID63_EngRefTorq.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID64_EngPctTorq_Idle() float64 {
	return Messages().OBD2.S01PID64_EngPctTorq_Idle.ToPhysical(float64(m.xxx_S01PID64_EngPctTorq_Idle))
}

func (m *OBD2) SetS01PID64_EngPctTorq_Idle(v float64) *OBD2 {
	m.xxx_S01PID64_EngPctTorq_Idle = uint8(Messages().OBD2.S01PID64_EngPctTorq_Idle.FromPhysical(v))
	return m
}

func (m *OBD2) RawS01PID64_EngPctTorq_Idle() uint8 {
	return m.xxx_S01PID64_EngPctTorq_Idle
}

func (m *OBD2) SetRawS01PID64_EngPctTorq_Idle(v uint8) *OBD2 {
	m.xxx_S01PID64_EngPctTorq_Idle = uint8(Messages().OBD2.S01PID64_EngPctTorq_Idle.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID64_EngPctTorq_EP1() float64 {
	return Messages().OBD2.S01PID64_EngPctTorq_EP1.ToPhysical(float64(m.xxx_S01PID64_EngPctTorq_EP1))
}

func (m *OBD2) SetS01PID64_EngPctTorq_EP1(v float64) *OBD2 {
	m.xxx_S01PID64_EngPctTorq_EP1 = uint8(Messages().OBD2.S01PID64_EngPctTorq_EP1.FromPhysical(v))
	return m
}

func (m *OBD2) RawS01PID64_EngPctTorq_EP1() uint8 {
	return m.xxx_S01PID64_EngPctTorq_EP1
}

func (m *OBD2) SetRawS01PID64_EngPctTorq_EP1(v uint8) *OBD2 {
	m.xxx_S01PID64_EngPctTorq_EP1 = uint8(Messages().OBD2.S01PID64_EngPctTorq_EP1.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID64_EngPctTorq_EP2() float64 {
	return Messages().OBD2.S01PID64_EngPctTorq_EP2.ToPhysical(float64(m.xxx_S01PID64_EngPctTorq_EP2))
}

func (m *OBD2) SetS01PID64_EngPctTorq_EP2(v float64) *OBD2 {
	m.xxx_S01PID64_EngPctTorq_EP2 = uint8(Messages().OBD2.S01PID64_EngPctTorq_EP2.FromPhysical(v))
	return m
}

func (m *OBD2) RawS01PID64_EngPctTorq_EP2() uint8 {
	return m.xxx_S01PID64_EngPctTorq_EP2
}

func (m *OBD2) SetRawS01PID64_EngPctTorq_EP2(v uint8) *OBD2 {
	m.xxx_S01PID64_EngPctTorq_EP2 = uint8(Messages().OBD2.S01PID64_EngPctTorq_EP2.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID64_EngPctTorq_EP3() float64 {
	return Messages().OBD2.S01PID64_EngPctTorq_EP3.ToPhysical(float64(m.xxx_S01PID64_EngPctTorq_EP3))
}

func (m *OBD2) SetS01PID64_EngPctTorq_EP3(v float64) *OBD2 {
	m.xxx_S01PID64_EngPctTorq_EP3 = uint8(Messages().OBD2.S01PID64_EngPctTorq_EP3.FromPhysical(v))
	return m
}

func (m *OBD2) RawS01PID64_EngPctTorq_EP3() uint8 {
	return m.xxx_S01PID64_EngPctTorq_EP3
}

func (m *OBD2) SetRawS01PID64_EngPctTorq_EP3(v uint8) *OBD2 {
	m.xxx_S01PID64_EngPctTorq_EP3 = uint8(Messages().OBD2.S01PID64_EngPctTorq_EP3.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID65_AuxInputOutput() uint8 {
	return m.xxx_S01PID65_AuxInputOutput
}

func (m *OBD2) SetS01PID65_AuxInputOutput(v uint8) *OBD2 {
	m.xxx_S01PID65_AuxInputOutput = uint8(Messages().OBD2.S01PID65_AuxInputOutput.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID66_MAFSensor() uint8 {
	return m.xxx_S01PID66_MAFSensor
}

func (m *OBD2) SetS01PID66_MAFSensor(v uint8) *OBD2 {
	m.xxx_S01PID66_MAFSensor = uint8(Messages().OBD2.S01PID66_MAFSensor.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID64_EngPctTorq_EP4() float64 {
	return Messages().OBD2.S01PID64_EngPctTorq_EP4.ToPhysical(float64(m.xxx_S01PID64_EngPctTorq_EP4))
}

func (m *OBD2) SetS01PID64_EngPctTorq_EP4(v float64) *OBD2 {
	m.xxx_S01PID64_EngPctTorq_EP4 = uint8(Messages().OBD2.S01PID64_EngPctTorq_EP4.FromPhysical(v))
	return m
}

func (m *OBD2) RawS01PID64_EngPctTorq_EP4() uint8 {
	return m.xxx_S01PID64_EngPctTorq_EP4
}

func (m *OBD2) SetRawS01PID64_EngPctTorq_EP4(v uint8) *OBD2 {
	m.xxx_S01PID64_EngPctTorq_EP4 = uint8(Messages().OBD2.S01PID64_EngPctTorq_EP4.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID67_EngineCoolantTemp() uint8 {
	return m.xxx_S01PID67_EngineCoolantTemp
}

func (m *OBD2) SetS01PID67_EngineCoolantTemp(v uint8) *OBD2 {
	m.xxx_S01PID67_EngineCoolantTemp = uint8(Messages().OBD2.S01PID67_EngineCoolantTemp.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID68_IntakeAirTempSens() uint8 {
	return m.xxx_S01PID68_IntakeAirTempSens
}

func (m *OBD2) SetS01PID68_IntakeAirTempSens(v uint8) *OBD2 {
	m.xxx_S01PID68_IntakeAirTempSens = uint8(Messages().OBD2.S01PID68_IntakeAirTempSens.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID69_CmdEGR_EGRError() uint8 {
	return m.xxx_S01PID69_CmdEGR_EGRError
}

func (m *OBD2) SetS01PID69_CmdEGR_EGRError(v uint8) *OBD2 {
	m.xxx_S01PID69_CmdEGR_EGRError = uint8(Messages().OBD2.S01PID69_CmdEGR_EGRError.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID6A_CmdDieselIntAir() uint8 {
	return m.xxx_S01PID6A_CmdDieselIntAir
}

func (m *OBD2) SetS01PID6A_CmdDieselIntAir(v uint8) *OBD2 {
	m.xxx_S01PID6A_CmdDieselIntAir = uint8(Messages().OBD2.S01PID6A_CmdDieselIntAir.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID6B_ExhaustGasTemp() uint8 {
	return m.xxx_S01PID6B_ExhaustGasTemp
}

func (m *OBD2) SetS01PID6B_ExhaustGasTemp(v uint8) *OBD2 {
	m.xxx_S01PID6B_ExhaustGasTemp = uint8(Messages().OBD2.S01PID6B_ExhaustGasTemp.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID6C_CmdThrottleActRel() uint8 {
	return m.xxx_S01PID6C_CmdThrottleActRel
}

func (m *OBD2) SetS01PID6C_CmdThrottleActRel(v uint8) *OBD2 {
	m.xxx_S01PID6C_CmdThrottleActRel = uint8(Messages().OBD2.S01PID6C_CmdThrottleActRel.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID6D_FuelPresContrSys() uint8 {
	return m.xxx_S01PID6D_FuelPresContrSys
}

func (m *OBD2) SetS01PID6D_FuelPresContrSys(v uint8) *OBD2 {
	m.xxx_S01PID6D_FuelPresContrSys = uint8(Messages().OBD2.S01PID6D_FuelPresContrSys.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID6E_InjPresContrSys() uint8 {
	return m.xxx_S01PID6E_InjPresContrSys
}

func (m *OBD2) SetS01PID6E_InjPresContrSys(v uint8) *OBD2 {
	m.xxx_S01PID6E_InjPresContrSys = uint8(Messages().OBD2.S01PID6E_InjPresContrSys.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID6F_TurboComprPres() uint8 {
	return m.xxx_S01PID6F_TurboComprPres
}

func (m *OBD2) SetS01PID6F_TurboComprPres(v uint8) *OBD2 {
	m.xxx_S01PID6F_TurboComprPres = uint8(Messages().OBD2.S01PID6F_TurboComprPres.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID70_BoostPresCntrl() uint8 {
	return m.xxx_S01PID70_BoostPresCntrl
}

func (m *OBD2) SetS01PID70_BoostPresCntrl(v uint8) *OBD2 {
	m.xxx_S01PID70_BoostPresCntrl = uint8(Messages().OBD2.S01PID70_BoostPresCntrl.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID80_PIDsSupported_81_A0() uint32 {
	return m.xxx_S01PID80_PIDsSupported_81_A0
}

func (m *OBD2) SetS01PID80_PIDsSupported_81_A0(v uint32) *OBD2 {
	m.xxx_S01PID80_PIDsSupported_81_A0 = uint32(Messages().OBD2.S01PID80_PIDsSupported_81_A0.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PID8E_EngFrictionPctTorq() float64 {
	return Messages().OBD2.S01PID8E_EngFrictionPctTorq.ToPhysical(float64(m.xxx_S01PID8E_EngFrictionPctTorq))
}

func (m *OBD2) SetS01PID8E_EngFrictionPctTorq(v float64) *OBD2 {
	m.xxx_S01PID8E_EngFrictionPctTorq = uint8(Messages().OBD2.S01PID8E_EngFrictionPctTorq.FromPhysical(v))
	return m
}

func (m *OBD2) RawS01PID8E_EngFrictionPctTorq() uint8 {
	return m.xxx_S01PID8E_EngFrictionPctTorq
}

func (m *OBD2) SetRawS01PID8E_EngFrictionPctTorq(v uint8) *OBD2 {
	m.xxx_S01PID8E_EngFrictionPctTorq = uint8(Messages().OBD2.S01PID8E_EngFrictionPctTorq.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PIDA0_PIDsSupported_A1_C0() uint32 {
	return m.xxx_S01PIDA0_PIDsSupported_A1_C0
}

func (m *OBD2) SetS01PIDA0_PIDsSupported_A1_C0(v uint32) *OBD2 {
	m.xxx_S01PIDA0_PIDsSupported_A1_C0 = uint32(Messages().OBD2.S01PIDA0_PIDsSupported_A1_C0.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *OBD2) S01PIDC0_PIDsSupported_C1_E0() uint32 {
	return m.xxx_S01PIDC0_PIDsSupported_C1_E0
}

func (m *OBD2) SetS01PIDC0_PIDsSupported_C1_E0(v uint32) *OBD2 {
	m.xxx_S01PIDC0_PIDsSupported_C1_E0 = uint32(Messages().OBD2.S01PIDC0_PIDsSupported_C1_E0.SaturatedCastUnsigned(uint64(v)))
	return m
}

// OBD2_Service models the Service signal of the OBD2 message.
type OBD2_Service uint8

// Value descriptions for the Service signal of the OBD2 message.
const (
	OBD2_Service_Showcurrentdata           OBD2_Service = 1
	OBD2_Service_Showfreezeframedata       OBD2_Service = 2
	OBD2_Service_ShowstoredDTCs            OBD2_Service = 3
	OBD2_Service_ClearDTCsandstoredvalues  OBD2_Service = 4
	OBD2_Service_Oxygensensormonitoring    OBD2_Service = 5
	OBD2_Service_Othersystemmonitoring     OBD2_Service = 6
	OBD2_Service_ShowpendingDTCs           OBD2_Service = 7
	OBD2_Service_Controlonboardsystem      OBD2_Service = 8
	OBD2_Service_Requestvehicleinformation OBD2_Service = 9
	OBD2_Service_PermanentDTCsClearedDTCs  OBD2_Service = 10
)

func (v OBD2_Service) String() string {
	switch v {
	case 1:
		return "Show current data "
	case 2:
		return "Show freeze frame data "
	case 3:
		return "Show stored DTCs "
	case 4:
		return "Clear DTCs and stored values"
	case 5:
		return "Oxygen sensor monitoring "
	case 6:
		return "Other system monitoring "
	case 7:
		return "Show pending DTCs "
	case 8:
		return "Control on-board system "
	case 9:
		return "Request vehicle information "
	case 10:
		return "Permanent DTCs (Cleared DTCs) "
	default:
		return fmt.Sprintf("OBD2_Service(%d)", v)
	}
}

// OBD2_S01PID models the S01PID signal of the OBD2 message.
type OBD2_S01PID uint8

// Value descriptions for the S01PID signal of the OBD2 message.
const (
	OBD2_S01PID_S01PID00PIDsSupported0120      OBD2_S01PID = 0
	OBD2_S01PID_S01PID01MonitorStatus          OBD2_S01PID = 1
	OBD2_S01PID_S01PID02FreezeDTC              OBD2_S01PID = 2
	OBD2_S01PID_S01PID03FuelSystemStatus       OBD2_S01PID = 3
	OBD2_S01PID_S01PID04CalcEngineLoad         OBD2_S01PID = 4
	OBD2_S01PID_S01PID05EngineCoolantTemp      OBD2_S01PID = 5
	OBD2_S01PID_S01PID06ShortFuelTrimBank1     OBD2_S01PID = 6
	OBD2_S01PID_S01PID07LongFuelTrimBank1      OBD2_S01PID = 7
	OBD2_S01PID_S01PID08ShortFuelTrimBank2     OBD2_S01PID = 8
	OBD2_S01PID_S01PID09LongFuelTrimBank2      OBD2_S01PID = 9
	OBD2_S01PID_S01PID0AFuelPressure           OBD2_S01PID = 10
	OBD2_S01PID_S01PID0BIntakeManiAbsPress     OBD2_S01PID = 11
	OBD2_S01PID_S01PID0CEngineRPM              OBD2_S01PID = 12
	OBD2_S01PID_S01PID0DVehicleSpeed           OBD2_S01PID = 13
	OBD2_S01PID_S01PID0ETimingAdvance          OBD2_S01PID = 14
	OBD2_S01PID_S01PID0FIntakeAirTemperature   OBD2_S01PID = 15
	OBD2_S01PID_S01PID10MAFAirFlowRate         OBD2_S01PID = 16
	OBD2_S01PID_S01PID11ThrottlePosition       OBD2_S01PID = 17
	OBD2_S01PID_S01PID12CmdSecAirStatus        OBD2_S01PID = 18
	OBD2_S01PID_S01PID14OxySensor1             OBD2_S01PID = 20
	OBD2_S01PID_S01PID15OxySensor2             OBD2_S01PID = 21
	OBD2_S01PID_S01PID16OxySensor3             OBD2_S01PID = 22
	OBD2_S01PID_S01PID17OxySensor4             OBD2_S01PID = 23
	OBD2_S01PID_S01PID18OxySensor5             OBD2_S01PID = 24
	OBD2_S01PID_S01PID19OxySensor6             OBD2_S01PID = 25
	OBD2_S01PID_S01PID1AOxySensor7             OBD2_S01PID = 26
	OBD2_S01PID_S01PID1BOxySensor8             OBD2_S01PID = 27
	OBD2_S01PID_S01PID1COBDStandard            OBD2_S01PID = 28
	OBD2_S01PID_S01PID1FTimeSinceEngStart      OBD2_S01PID = 31
	OBD2_S01PID_S01PID20PIDsSupported2140      OBD2_S01PID = 32
	OBD2_S01PID_S01PID21DistanceMILOn          OBD2_S01PID = 33
	OBD2_S01PID_S01PID22FuelRailPres           OBD2_S01PID = 34
	OBD2_S01PID_S01PID23FuelRailGaug           OBD2_S01PID = 35
	OBD2_S01PID_S01PID24OxySensor1             OBD2_S01PID = 36
	OBD2_S01PID_S01PID25OxySensor2             OBD2_S01PID = 37
	OBD2_S01PID_S01PID26OxySensor3             OBD2_S01PID = 38
	OBD2_S01PID_S01PID27OxySensor4             OBD2_S01PID = 39
	OBD2_S01PID_S01PID28OxySensor5             OBD2_S01PID = 40
	OBD2_S01PID_S01PID29OxySensor6             OBD2_S01PID = 41
	OBD2_S01PID_S01PID2AOxySensor7             OBD2_S01PID = 42
	OBD2_S01PID_S01PID2BOxySensor8             OBD2_S01PID = 43
	OBD2_S01PID_S01PID2CCmdEGR                 OBD2_S01PID = 44
	OBD2_S01PID_S01PID2DEGRError               OBD2_S01PID = 45
	OBD2_S01PID_S01PID2ECmdEvapPurge           OBD2_S01PID = 46
	OBD2_S01PID_S01PID2FFuelTankLevel          OBD2_S01PID = 47
	OBD2_S01PID_S01PID30WarmUpsSinceCodeClear  OBD2_S01PID = 48
	OBD2_S01PID_S01PID31DistanceSinceCodeClear OBD2_S01PID = 49
	OBD2_S01PID_S01PID32EvapSysVaporPres       OBD2_S01PID = 50
	OBD2_S01PID_S01PID33AbsBaroPres            OBD2_S01PID = 51
	OBD2_S01PID_S01PID34OxySensor1             OBD2_S01PID = 52
	OBD2_S01PID_S01PID35OxySensor2             OBD2_S01PID = 53
	OBD2_S01PID_S01PID36OxySensor3             OBD2_S01PID = 54
	OBD2_S01PID_S01PID37OxySensor4             OBD2_S01PID = 55
	OBD2_S01PID_S01PID38OxySensor5             OBD2_S01PID = 56
	OBD2_S01PID_S01PID39OxySensor6             OBD2_S01PID = 57
	OBD2_S01PID_S01PID3AOxySensor7             OBD2_S01PID = 58
	OBD2_S01PID_S01PID3BOxySensor8             OBD2_S01PID = 59
	OBD2_S01PID_S01PID3CCatTempBank1Sens1      OBD2_S01PID = 60
	OBD2_S01PID_S01PID3DCatTempBank2Sens1      OBD2_S01PID = 61
	OBD2_S01PID_S01PID3ECatTempBank1Sens2      OBD2_S01PID = 62
	OBD2_S01PID_S01PID3FCatTempBank2Sens2      OBD2_S01PID = 63
	OBD2_S01PID_S01PID40PIDsSupported4160      OBD2_S01PID = 64
	OBD2_S01PID_S01PID41MonStatusDriveCycle    OBD2_S01PID = 65
	OBD2_S01PID_S01PID42ControlModuleVolt      OBD2_S01PID = 66
	OBD2_S01PID_S01PID43AbsLoadValue           OBD2_S01PID = 67
	OBD2_S01PID_S01PID44FuelAirCmdEquiv        OBD2_S01PID = 68
	OBD2_S01PID_S01PID45RelThrottlePos         OBD2_S01PID = 69
	OBD2_S01PID_S01PID46AmbientAirTemp         OBD2_S01PID = 70
	OBD2_S01PID_S01PID47AbsThrottlePosB        OBD2_S01PID = 71
	OBD2_S01PID_S01PID48AbsThrottlePosC        OBD2_S01PID = 72
	OBD2_S01PID_S01PID49AbsThrottlePosD        OBD2_S01PID = 73
	OBD2_S01PID_S01PID4AAbsThrottlePosE        OBD2_S01PID = 74
	OBD2_S01PID_S01PID4BAbsThrottlePosF        OBD2_S01PID = 75
	OBD2_S01PID_S01PID4CCmdThrottleAct         OBD2_S01PID = 76
	OBD2_S01PID_S01PID4DTimeRunMILOn           OBD2_S01PID = 77
	OBD2_S01PID_S01PID4ETimeSinceCodeClear     OBD2_S01PID = 78
	OBD2_S01PID_S01PID4FMaxMultiple            OBD2_S01PID = 79
	OBD2_S01PID_S01PID50MaxAirFlowMAF          OBD2_S01PID = 80
	OBD2_S01PID_S01PID51FuelType               OBD2_S01PID = 81
	OBD2_S01PID_S01PID52EthanolFuelPct         OBD2_S01PID = 82
	OBD2_S01PID_S01PID53AbsEvapSysVapPres      OBD2_S01PID = 83
	OBD2_S01PID_S01PID54EvapSysVapPres         OBD2_S01PID = 84
	OBD2_S01PID_S01PID55ShortSecOxyTrimBankX   OBD2_S01PID = 85
	OBD2_S01PID_S01PID56LongSecOxyTrimBankX    OBD2_S01PID = 86
	OBD2_S01PID_S01PID57ShortSecOxyTrimBankX   OBD2_S01PID = 87
	OBD2_S01PID_S01PID58LongSecOxyTrimBankX    OBD2_S01PID = 88
	OBD2_S01PID_S01PID59FuelRailAbsPres        OBD2_S01PID = 89
	OBD2_S01PID_S01PID5ARelAccelPedalPos       OBD2_S01PID = 90
	OBD2_S01PID_S01PID5BHybrBatPackRemLife     OBD2_S01PID = 91
	OBD2_S01PID_S01PID5CEngineOilTemp          OBD2_S01PID = 92
	OBD2_S01PID_S01PID5DFuelInjectionTiming    OBD2_S01PID = 93
	OBD2_S01PID_S01PID5EEngineFuelRate         OBD2_S01PID = 94
	OBD2_S01PID_S01PID5FEmissionReq            OBD2_S01PID = 95
	OBD2_S01PID_S01PID60PIDsSupported6180      OBD2_S01PID = 96
	OBD2_S01PID_S01PID61DemandEngTorqPct       OBD2_S01PID = 97
	OBD2_S01PID_S01PID62ActualEngTorqPct       OBD2_S01PID = 98
	OBD2_S01PID_S01PID63EngRefTorq             OBD2_S01PID = 99
	OBD2_S01PID_S01PID64EngPctTorq             OBD2_S01PID = 100
	OBD2_S01PID_S01PID65AuxInputOutput         OBD2_S01PID = 101
	OBD2_S01PID_S01PID66MAFSensor              OBD2_S01PID = 102
	OBD2_S01PID_S01PID67EngineCoolantTemp      OBD2_S01PID = 103
	OBD2_S01PID_S01PID68IntakeAirTempSens      OBD2_S01PID = 104
	OBD2_S01PID_S01PID69CmdEGREGRError         OBD2_S01PID = 105
	OBD2_S01PID_S01PID6ACmdDieselIntAir        OBD2_S01PID = 106
	OBD2_S01PID_S01PID6BExhaustGasTemp         OBD2_S01PID = 107
	OBD2_S01PID_S01PID6CCmdThrottleActRel      OBD2_S01PID = 108
	OBD2_S01PID_S01PID6DFuelPresContrSys       OBD2_S01PID = 109
	OBD2_S01PID_S01PID6EInjPresContrSys        OBD2_S01PID = 110
	OBD2_S01PID_S01PID6FTurboComprPres         OBD2_S01PID = 111
	OBD2_S01PID_S01PID70BoostPresCntrl         OBD2_S01PID = 112
	OBD2_S01PID_S01PID80PIDsSupported81A0      OBD2_S01PID = 128
	OBD2_S01PID_S01PID8EEngFrictionPctTorq     OBD2_S01PID = 142
	OBD2_S01PID_S01PIDA0PIDsSupportedA1C0      OBD2_S01PID = 160
	OBD2_S01PID_S01PIDC0PIDsSupportedC1E0      OBD2_S01PID = 192
)

func (v OBD2_S01PID) String() string {
	switch v {
	case 0:
		return "S01PID00_PIDsSupported_01_20"
	case 1:
		return "S01PID01_MonitorStatus"
	case 2:
		return "S01PID02_FreezeDTC"
	case 3:
		return "S01PID03_FuelSystemStatus"
	case 4:
		return "S01PID04_CalcEngineLoad"
	case 5:
		return "S01PID05_EngineCoolantTemp"
	case 6:
		return "S01PID06_ShortFuelTrimBank1"
	case 7:
		return "S01PID07_LongFuelTrimBank1"
	case 8:
		return "S01PID08_ShortFuelTrimBank2"
	case 9:
		return "S01PID09_LongFuelTrimBank2"
	case 10:
		return "S01PID0A_FuelPressure"
	case 11:
		return "S01PID0B_IntakeManiAbsPress"
	case 12:
		return "S01PID0C_EngineRPM"
	case 13:
		return "S01PID0D_VehicleSpeed"
	case 14:
		return "S01PID0E_TimingAdvance"
	case 15:
		return "S01PID0F_IntakeAirTemperature"
	case 16:
		return "S01PID10_MAFAirFlowRate"
	case 17:
		return "S01PID11_ThrottlePosition"
	case 18:
		return "S01PID12_CmdSecAirStatus"
	case 20:
		return "S01PID14_OxySensor1"
	case 21:
		return "S01PID15_OxySensor2"
	case 22:
		return "S01PID16_OxySensor3"
	case 23:
		return "S01PID17_OxySensor4"
	case 24:
		return "S01PID18_OxySensor5"
	case 25:
		return "S01PID19_OxySensor6"
	case 26:
		return "S01PID1A_OxySensor7"
	case 27:
		return "S01PID1B_OxySensor8"
	case 28:
		return "S01PID1C_OBDStandard"
	case 31:
		return "S01PID1F_TimeSinceEngStart"
	case 32:
		return "S01PID20_PIDsSupported_21_40"
	case 33:
		return "S01PID21_DistanceMILOn"
	case 34:
		return "S01PID22_FuelRailPres"
	case 35:
		return "S01PID23_FuelRailGaug"
	case 36:
		return "S01PID24_OxySensor1"
	case 37:
		return "S01PID25_OxySensor2"
	case 38:
		return "S01PID26_OxySensor3"
	case 39:
		return "S01PID27_OxySensor4"
	case 40:
		return "S01PID28_OxySensor5"
	case 41:
		return "S01PID29_OxySensor6"
	case 42:
		return "S01PID2A_OxySensor7"
	case 43:
		return "S01PID2B_OxySensor8"
	case 44:
		return "S01PID2C_CmdEGR"
	case 45:
		return "S01PID2D_EGRError"
	case 46:
		return "S01PID2E_CmdEvapPurge"
	case 47:
		return "S01PID2F_FuelTankLevel"
	case 48:
		return "S01PID30_WarmUpsSinceCodeClear"
	case 49:
		return "S01PID31_DistanceSinceCodeClear"
	case 50:
		return "S01PID32_EvapSysVaporPres"
	case 51:
		return "S01PID33_AbsBaroPres"
	case 52:
		return "S01PID34_OxySensor1"
	case 53:
		return "S01PID35_OxySensor2"
	case 54:
		return "S01PID36_OxySensor3"
	case 55:
		return "S01PID37_OxySensor4"
	case 56:
		return "S01PID38_OxySensor5"
	case 57:
		return "S01PID39_OxySensor6"
	case 58:
		return "S01PID3A_OxySensor7"
	case 59:
		return "S01PID3B_OxySensor8"
	case 60:
		return "S01PID3C_CatTempBank1Sens1"
	case 61:
		return "S01PID3D_CatTempBank2Sens1"
	case 62:
		return "S01PID3E_CatTempBank1Sens2"
	case 63:
		return "S01PID3F_CatTempBank2Sens2"
	case 64:
		return "S01PID40_PIDsSupported_41_60"
	case 65:
		return "S01PID41_MonStatusDriveCycle"
	case 66:
		return "S01PID42_ControlModuleVolt"
	case 67:
		return "S01PID43_AbsLoadValue"
	case 68:
		return "S01PID44_FuelAirCmdEquiv"
	case 69:
		return "S01PID45_RelThrottlePos"
	case 70:
		return "S01PID46_AmbientAirTemp"
	case 71:
		return "S01PID47_AbsThrottlePosB"
	case 72:
		return "S01PID48_AbsThrottlePosC"
	case 73:
		return "S01PID49_AbsThrottlePosD"
	case 74:
		return "S01PID4A_AbsThrottlePosE"
	case 75:
		return "S01PID4B_AbsThrottlePosF"
	case 76:
		return "S01PID4C_CmdThrottleAct"
	case 77:
		return "S01PID4D_TimeRunMILOn"
	case 78:
		return "S01PID4E_TimeSinceCodeClear"
	case 79:
		return "S01PID4F_MaxMultiple"
	case 80:
		return "S01PID50_Max_AirFlowMAF"
	case 81:
		return "S01PID51_FuelType"
	case 82:
		return "S01PID52_EthanolFuelPct"
	case 83:
		return "S01PID53_AbsEvapSysVapPres"
	case 84:
		return "S01PID54_EvapSysVapPres"
	case 85:
		return "S01PID55_ShortSecOxyTrimBankX"
	case 86:
		return "S01PID56_LongSecOxyTrimBankX"
	case 87:
		return "S01PID57_ShortSecOxyTrimBankX"
	case 88:
		return "S01PID58_LongSecOxyTrimBankX"
	case 89:
		return "S01PID59_FuelRailAbsPres"
	case 90:
		return "S01PID5A_RelAccelPedalPos"
	case 91:
		return "S01PID5B_HybrBatPackRemLife"
	case 92:
		return "S01PID5C_EngineOilTemp"
	case 93:
		return "S01PID5D_FuelInjectionTiming"
	case 94:
		return "S01PID5E_EngineFuelRate"
	case 95:
		return "S01PID5F_EmissionReq"
	case 96:
		return "S01PID60_PIDsSupported_61_80"
	case 97:
		return "S01PID61_DemandEngTorqPct"
	case 98:
		return "S01PID62_ActualEngTorqPct"
	case 99:
		return "S01PID63_EngRefTorq"
	case 100:
		return "S01PID64_EngPctTorq"
	case 101:
		return "S01PID65_AuxInputOutput"
	case 102:
		return "S01PID66_MAFSensor"
	case 103:
		return "S01PID67_EngineCoolantTemp"
	case 104:
		return "S01PID68_IntakeAirTempSens"
	case 105:
		return "S01PID69_CmdEGR_EGRError"
	case 106:
		return "S01PID6A_CmdDieselIntAir"
	case 107:
		return "S01PID6B_ExhaustGasTemp"
	case 108:
		return "S01PID6C_CmdThrottleActRel"
	case 109:
		return "S01PID6D_FuelPresContrSys"
	case 110:
		return "S01PID6E_InjPresContrSys"
	case 111:
		return "S01PID6F_TurboComprPres"
	case 112:
		return "S01PID70_BoostPresCntrl"
	case 128:
		return "S01PID80_PIDsSupported_81_A0"
	case 142:
		return "S01PID8E_EngFrictionPctTorq"
	case 160:
		return "S01PIDA0_PIDsSupported_A1_C0"
	case 192:
		return "S01PIDC0_PIDsSupported_C1_E0"
	default:
		return fmt.Sprintf("OBD2_S01PID(%d)", v)
	}
}

// OBD2_S02PID models the S02PID signal of the OBD2 message.
type OBD2_S02PID uint8

// Value descriptions for the S02PID signal of the OBD2 message.
const (
	OBD2_S02PID_S02PID02FreezeDTC OBD2_S02PID = 2
)

func (v OBD2_S02PID) String() string {
	switch v {
	case 2:
		return "S02PID02_FreezeDTC"
	default:
		return fmt.Sprintf("OBD2_S02PID(%d)", v)
	}
}

// OBD2_S01PID03_FuelSystemStatus models the S01PID03_FuelSystemStatus signal of the OBD2 message.
type OBD2_S01PID03_FuelSystemStatus uint16

// Value descriptions for the S01PID03_FuelSystemStatus signal of the OBD2 message.
const (
	OBD2_S01PID03_FuelSystemStatus_Openloopinsuffengtemp   OBD2_S01PID03_FuelSystemStatus = 1
	OBD2_S01PID03_FuelSystemStatus_Closedloopoxysens       OBD2_S01PID03_FuelSystemStatus = 2
	OBD2_S01PID03_FuelSystemStatus_Openloopengloadfuelcut  OBD2_S01PID03_FuelSystemStatus = 4
	OBD2_S01PID03_FuelSystemStatus_Openloopsystemfailure   OBD2_S01PID03_FuelSystemStatus = 8
	OBD2_S01PID03_FuelSystemStatus_Closedloopfeedbackissue OBD2_S01PID03_FuelSystemStatus = 16
)

func (v OBD2_S01PID03_FuelSystemStatus) String() string {
	switch v {
	case 1:
		return "Open loop (insuff. eng. temp.)"
	case 2:
		return "Closed loop (oxy sens)"
	case 4:
		return "Open loop (eng. load, fuel cut)"
	case 8:
		return "Open loop (system failure)"
	case 16:
		return "Closed loop (feedback issue)"
	default:
		return fmt.Sprintf("OBD2_S01PID03_FuelSystemStatus(%d)", v)
	}
}

// OBD2_S01PID12_CmdSecAirStatus models the S01PID12_CmdSecAirStatus signal of the OBD2 message.
type OBD2_S01PID12_CmdSecAirStatus uint8

// Value descriptions for the S01PID12_CmdSecAirStatus signal of the OBD2 message.
const (
	OBD2_S01PID12_CmdSecAirStatus_Upstream                 OBD2_S01PID12_CmdSecAirStatus = 1
	OBD2_S01PID12_CmdSecAirStatus_Downstreamcatalyticconv  OBD2_S01PID12_CmdSecAirStatus = 2
	OBD2_S01PID12_CmdSecAirStatus_Fromoutsideatmosphereoff OBD2_S01PID12_CmdSecAirStatus = 4
	OBD2_S01PID12_CmdSecAirStatus_Pumpcmdonfordiagn        OBD2_S01PID12_CmdSecAirStatus = 8
)

func (v OBD2_S01PID12_CmdSecAirStatus) String() string {
	switch v {
	case 1:
		return "Upstream"
	case 2:
		return "Downstream catalytic conv"
	case 4:
		return "From outside atmosphere/off"
	case 8:
		return "Pump cmd on for diagn."
	default:
		return fmt.Sprintf("OBD2_S01PID12_CmdSecAirStatus(%d)", v)
	}
}

// OBD2_S01PID1C_OBDStandard models the S01PID1C_OBDStandard signal of the OBD2 message.
type OBD2_S01PID1C_OBDStandard uint8

// Value descriptions for the S01PID1C_OBDStandard signal of the OBD2 message.
const (
	OBD2_S01PID1C_OBDStandard_OBDIIasdefinedbytheCARB OBD2_S01PID1C_OBDStandard = 1
	OBD2_S01PID1C_OBDStandard_OBDasdefinedbytheEPA    OBD2_S01PID1C_OBDStandard = 2
	OBD2_S01PID1C_OBDStandard_OBDandOBDII             OBD2_S01PID1C_OBDStandard = 3
	OBD2_S01PID1C_OBDStandard_OBDI                    OBD2_S01PID1C_OBDStandard = 4
	OBD2_S01PID1C_OBDStandard_NotOBDcompliant         OBD2_S01PID1C_OBDStandard = 5
	OBD2_S01PID1C_OBDStandard_EOBDEurope              OBD2_S01PID1C_OBDStandard = 6
	OBD2_S01PID1C_OBDStandard_EOBDandOBDII            OBD2_S01PID1C_OBDStandard = 7
	OBD2_S01PID1C_OBDStandard_EOBDandOBD              OBD2_S01PID1C_OBDStandard = 8
	OBD2_S01PID1C_OBDStandard_EOBDOBDandOBDII         OBD2_S01PID1C_OBDStandard = 9
	OBD2_S01PID1C_OBDStandard_JOBDJapan               OBD2_S01PID1C_OBDStandard = 10
	OBD2_S01PID1C_OBDStandard_JOBDandOBDII            OBD2_S01PID1C_OBDStandard = 11
	OBD2_S01PID1C_OBDStandard_JOBDandEOBD             OBD2_S01PID1C_OBDStandard = 12
	OBD2_S01PID1C_OBDStandard_JOBDEOBDandOBDII        OBD2_S01PID1C_OBDStandard = 13
	OBD2_S01PID1C_OBDStandard_Reserved1               OBD2_S01PID1C_OBDStandard = 14
	OBD2_S01PID1C_OBDStandard_Reserved2               OBD2_S01PID1C_OBDStandard = 15
	OBD2_S01PID1C_OBDStandard_Reserved3               OBD2_S01PID1C_OBDStandard = 16
	OBD2_S01PID1C_OBDStandard_EngManuDiagEMD          OBD2_S01PID1C_OBDStandard = 17
	OBD2_S01PID1C_OBDStandard_EMDEnhancedEMD          OBD2_S01PID1C_OBDStandard = 18
	OBD2_S01PID1C_OBDStandard_HDOBDC                  OBD2_S01PID1C_OBDStandard = 19
	OBD2_S01PID1C_OBDStandard_HDOBD                   OBD2_S01PID1C_OBDStandard = 20
	OBD2_S01PID1C_OBDStandard_WWHOBD                  OBD2_S01PID1C_OBDStandard = 21
	OBD2_S01PID1C_OBDStandard_Reserved4               OBD2_S01PID1C_OBDStandard = 22
	OBD2_S01PID1C_OBDStandard_HDEOBDI                 OBD2_S01PID1C_OBDStandard = 23
	OBD2_S01PID1C_OBDStandard_HDEOBDIN                OBD2_S01PID1C_OBDStandard = 24
	OBD2_S01PID1C_OBDStandard_HDEOBDII                OBD2_S01PID1C_OBDStandard = 25
	OBD2_S01PID1C_OBDStandard_HDEOBDIIN               OBD2_S01PID1C_OBDStandard = 26
	OBD2_S01PID1C_OBDStandard_Reserved5               OBD2_S01PID1C_OBDStandard = 27
	OBD2_S01PID1C_OBDStandard_OBDBr1                  OBD2_S01PID1C_OBDStandard = 28
	OBD2_S01PID1C_OBDStandard_OBDBr2                  OBD2_S01PID1C_OBDStandard = 29
	OBD2_S01PID1C_OBDStandard_KOBD                    OBD2_S01PID1C_OBDStandard = 30
	OBD2_S01PID1C_OBDStandard_IOBDI                   OBD2_S01PID1C_OBDStandard = 31
	OBD2_S01PID1C_OBDStandard_IOBDII                  OBD2_S01PID1C_OBDStandard = 32
	OBD2_S01PID1C_OBDStandard_HDEOBDIV                OBD2_S01PID1C_OBDStandard = 33
	OBD2_S01PID1C_OBDStandard_Reserved6               OBD2_S01PID1C_OBDStandard = 34
	OBD2_S01PID1C_OBDStandard_Reserved7               OBD2_S01PID1C_OBDStandard = 35
)

func (v OBD2_S01PID1C_OBDStandard) String() string {
	switch v {
	case 1:
		return "OBD-II as defined by the CARB"
	case 2:
		return "OBD as defined by the EPA"
	case 3:
		return "OBD and OBD-II"
	case 4:
		return "OBD-I"
	case 5:
		return "Not OBD compliant"
	case 6:
		return "EOBD (Europe)"
	case 7:
		return "EOBD and OBD-II"
	case 8:
		return "EOBD and OBD"
	case 9:
		return "EOBD, OBD and OBD II"
	case 10:
		return "JOBD (Japan)"
	case 11:
		return "JOBD and OBD II"
	case 12:
		return "JOBD and EOBD"
	case 13:
		return "JOBD, EOBD, and OBD II"
	case 14:
		return "Reserved"
	case 15:
		return "Reserved"
	case 16:
		return "Reserved"
	case 17:
		return "Eng. Manu. Diag. (EMD)"
	case 18:
		return "EMD Enhanced (EMD+)"
	case 19:
		return "HD OBD-C"
	case 20:
		return "HD OBD"
	case 21:
		return "WWH OBD"
	case 22:
		return "Reserved"
	case 23:
		return "HD EOBD-I"
	case 24:
		return "HD EOBD-I N"
	case 25:
		return "HD EOBD-II"
	case 26:
		return "HD EOBD-II N"
	case 27:
		return "Reserved"
	case 28:
		return "OBDBr-1"
	case 29:
		return "OBDBr-2"
	case 30:
		return "KOBD"
	case 31:
		return "IOBD I"
	case 32:
		return "IOBD II"
	case 33:
		return "HD EOBD-IV"
	case 34:
		return "Reserved"
	case 35:
		return "Reserved"
	default:
		return fmt.Sprintf("OBD2_S01PID1C_OBDStandard(%d)", v)
	}
}

// OBD2_S01PID51_FuelType models the S01PID51_FuelType signal of the OBD2 message.
type OBD2_S01PID51_FuelType uint8

// Value descriptions for the S01PID51_FuelType signal of the OBD2 message.
const (
	OBD2_S01PID51_FuelType_Notavailable              OBD2_S01PID51_FuelType = 0
	OBD2_S01PID51_FuelType_Gasoline                  OBD2_S01PID51_FuelType = 1
	OBD2_S01PID51_FuelType_Methanol                  OBD2_S01PID51_FuelType = 2
	OBD2_S01PID51_FuelType_Ethanol                   OBD2_S01PID51_FuelType = 3
	OBD2_S01PID51_FuelType_Diesel                    OBD2_S01PID51_FuelType = 4
	OBD2_S01PID51_FuelType_LPG                       OBD2_S01PID51_FuelType = 5
	OBD2_S01PID51_FuelType_CNG                       OBD2_S01PID51_FuelType = 6
	OBD2_S01PID51_FuelType_Propane                   OBD2_S01PID51_FuelType = 7
	OBD2_S01PID51_FuelType_Electric                  OBD2_S01PID51_FuelType = 8
	OBD2_S01PID51_FuelType_BifuelrunningGasoline     OBD2_S01PID51_FuelType = 9
	OBD2_S01PID51_FuelType_BifuelrunningMethanol     OBD2_S01PID51_FuelType = 10
	OBD2_S01PID51_FuelType_BifuelrunningEthanol      OBD2_S01PID51_FuelType = 11
	OBD2_S01PID51_FuelType_BifuelrunningLPG          OBD2_S01PID51_FuelType = 12
	OBD2_S01PID51_FuelType_BifuelrunningCNG          OBD2_S01PID51_FuelType = 13
	OBD2_S01PID51_FuelType_BifuelrunningPropane      OBD2_S01PID51_FuelType = 14
	OBD2_S01PID51_FuelType_BifuelrunningElectricity  OBD2_S01PID51_FuelType = 15
	OBD2_S01PID51_FuelType_Bifuelelectriccombeng     OBD2_S01PID51_FuelType = 16
	OBD2_S01PID51_FuelType_Hybridgasoline            OBD2_S01PID51_FuelType = 17
	OBD2_S01PID51_FuelType_HybridEthanol             OBD2_S01PID51_FuelType = 18
	OBD2_S01PID51_FuelType_HybridDiesel              OBD2_S01PID51_FuelType = 19
	OBD2_S01PID51_FuelType_HybridElectric            OBD2_S01PID51_FuelType = 20
	OBD2_S01PID51_FuelType_Hybridrunningelectriccomb OBD2_S01PID51_FuelType = 21
	OBD2_S01PID51_FuelType_HybridRegenerative        OBD2_S01PID51_FuelType = 22
	OBD2_S01PID51_FuelType_Bifuelrunningdiesel       OBD2_S01PID51_FuelType = 23
)

func (v OBD2_S01PID51_FuelType) String() string {
	switch v {
	case 0:
		return "Not available"
	case 1:
		return "Gasoline"
	case 2:
		return "Methanol"
	case 3:
		return "Ethanol"
	case 4:
		return "Diesel"
	case 5:
		return "LPG"
	case 6:
		return "CNG"
	case 7:
		return "Propane"
	case 8:
		return "Electric"
	case 9:
		return "Bifuel running Gasoline"
	case 10:
		return "Bifuel running Methanol"
	case 11:
		return "Bifuel running Ethanol"
	case 12:
		return "Bifuel running LPG"
	case 13:
		return "Bifuel running CNG"
	case 14:
		return "Bifuel running Propane"
	case 15:
		return "Bifuel running Electricity"
	case 16:
		return "Bifuel electric/comb. eng."
	case 17:
		return "Hybrid gasoline"
	case 18:
		return "Hybrid Ethanol"
	case 19:
		return "Hybrid Diesel"
	case 20:
		return "Hybrid Electric"
	case 21:
		return "Hybrid running electric/comb."
	case 22:
		return "Hybrid Regenerative"
	case 23:
		return "Bifuel running diesel"
	default:
		return fmt.Sprintf("OBD2_S01PID51_FuelType(%d)", v)
	}
}

// Frame returns a CAN frame representing the message.
func (m *OBD2) Frame() can.Frame {
	md := Messages().OBD2
	f := can.Frame{ID: md.ID, IsExtended: md.IsExtended, Length: md.Length.Length}
	md.Length.MarshalUnsigned(&f.Data, uint64(m.xxx_Length))
	md.Service.MarshalUnsigned(&f.Data, uint64(m.xxx_Service))
	md.Response.MarshalUnsigned(&f.Data, uint64(m.xxx_Response))
	if m.xxx_Service == 1 {
		md.S01PID.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID))
	}
	if m.xxx_Service == 2 {
		md.S02PID.MarshalUnsigned(&f.Data, uint64(m.xxx_S02PID))
	}
	if m.xxx_Service == 0 {
		md.S01PID00_PIDsSupported_01_20.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID00_PIDsSupported_01_20))
	}
	if m.xxx_Service == 1 {
		md.S01PID01_MonitorStatus.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID01_MonitorStatus))
	}
	if m.xxx_Service == 2 {
		md.S02PID02_FreezeDTC.MarshalUnsigned(&f.Data, uint64(m.xxx_S02PID02_FreezeDTC))
	}
	if m.xxx_Service == 2 {
		md.S01PID02_FreezeDTC.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID02_FreezeDTC))
	}
	if m.xxx_Service == 3 {
		md.S01PID03_FuelSystemStatus.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID03_FuelSystemStatus))
	}
	if m.xxx_Service == 4 {
		md.S01PID04_CalcEngineLoad.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID04_CalcEngineLoad))
	}
	if m.xxx_Service == 5 {
		md.S01PID05_EngineCoolantTemp.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID05_EngineCoolantTemp))
	}
	if m.xxx_Service == 6 {
		md.S01PID06_ShortFuelTrimBank1.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID06_ShortFuelTrimBank1))
	}
	if m.xxx_Service == 7 {
		md.S01PID07_LongFuelTrimBank1.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID07_LongFuelTrimBank1))
	}
	if m.xxx_Service == 8 {
		md.S01PID08_ShortFuelTrimBank2.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID08_ShortFuelTrimBank2))
	}
	if m.xxx_Service == 9 {
		md.S01PID09_LongFuelTrimBank2.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID09_LongFuelTrimBank2))
	}
	if m.xxx_Service == 10 {
		md.S01PID0A_FuelPressure.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID0A_FuelPressure))
	}
	if m.xxx_Service == 11 {
		md.S01PID0B_IntakeManiAbsPress.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID0B_IntakeManiAbsPress))
	}
	if m.xxx_Service == 12 {
		md.S01PID0C_EngineRPM.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID0C_EngineRPM))
	}
	if m.xxx_Service == 13 {
		md.S01PID0D_VehicleSpeed.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID0D_VehicleSpeed))
	}
	if m.xxx_Service == 14 {
		md.S01PID0E_TimingAdvance.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID0E_TimingAdvance))
	}
	if m.xxx_Service == 15 {
		md.S01PID0F_IntakeAirTemperature.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID0F_IntakeAirTemperature))
	}
	if m.xxx_Service == 16 {
		md.S01PID10_MAFAirFlowRate.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID10_MAFAirFlowRate))
	}
	if m.xxx_Service == 17 {
		md.S01PID11_ThrottlePosition.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID11_ThrottlePosition))
	}
	if m.xxx_Service == 18 {
		md.S01PID12_CmdSecAirStatus.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID12_CmdSecAirStatus))
	}
	if m.xxx_Service == 20 {
		md.S01PID14_OxySensor1_Volt.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID14_OxySensor1_Volt))
	}
	if m.xxx_Service == 21 {
		md.S01PID15_OxySensor2_Volt.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID15_OxySensor2_Volt))
	}
	if m.xxx_Service == 22 {
		md.S01PID16_OxySensor3_Volt.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID16_OxySensor3_Volt))
	}
	if m.xxx_Service == 23 {
		md.S01PID17_OxySensor4_Volt.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID17_OxySensor4_Volt))
	}
	if m.xxx_Service == 24 {
		md.S01PID18_OxySensor5_Volt.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID18_OxySensor5_Volt))
	}
	if m.xxx_Service == 25 {
		md.S01PID19_OxySensor6_Volt.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID19_OxySensor6_Volt))
	}
	if m.xxx_Service == 20 {
		md.S01PID14_OxySensor1_STFT.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID14_OxySensor1_STFT))
	}
	if m.xxx_Service == 21 {
		md.S01PID15_OxySensor2_STFT.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID15_OxySensor2_STFT))
	}
	if m.xxx_Service == 22 {
		md.S01PID16_OxySensor3_STFT.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID16_OxySensor3_STFT))
	}
	if m.xxx_Service == 23 {
		md.S01PID17_OxySensor4_STFT.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID17_OxySensor4_STFT))
	}
	if m.xxx_Service == 24 {
		md.S01PID18_OxySensor5_STFT.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID18_OxySensor5_STFT))
	}
	if m.xxx_Service == 25 {
		md.S01PID19_OxySensor6_STFT.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID19_OxySensor6_STFT))
	}
	if m.xxx_Service == 26 {
		md.S01PID1A_OxySensor7_Volt.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID1A_OxySensor7_Volt))
	}
	if m.xxx_Service == 26 {
		md.S01PID1A_OxySensor7_STFT.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID1A_OxySensor7_STFT))
	}
	if m.xxx_Service == 27 {
		md.S01PID1B_OxySensor8_Volt.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID1B_OxySensor8_Volt))
	}
	if m.xxx_Service == 27 {
		md.S01PID1B_OxySensor8_STFT.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID1B_OxySensor8_STFT))
	}
	if m.xxx_Service == 28 {
		md.S01PID1C_OBDStandard.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID1C_OBDStandard))
	}
	if m.xxx_Service == 31 {
		md.S01PID1F_TimeSinceEngStart.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID1F_TimeSinceEngStart))
	}
	if m.xxx_Service == 32 {
		md.S01PID20_PIDsSupported_21_40.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID20_PIDsSupported_21_40))
	}
	if m.xxx_Service == 33 {
		md.S01PID21_DistanceMILOn.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID21_DistanceMILOn))
	}
	if m.xxx_Service == 34 {
		md.S01PID22_FuelRailPres.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID22_FuelRailPres))
	}
	if m.xxx_Service == 35 {
		md.S01PID23_FuelRailGaug.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID23_FuelRailGaug))
	}
	if m.xxx_Service == 36 {
		md.S01PID24_OxySensor1_FAER.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID24_OxySensor1_FAER))
	}
	if m.xxx_Service == 36 {
		md.S01PID24_OxySensor1_Volt.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID24_OxySensor1_Volt))
	}
	if m.xxx_Service == 37 {
		md.S01PID25_OxySensor2_FAER.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID25_OxySensor2_FAER))
	}
	if m.xxx_Service == 37 {
		md.S01PID25_OxySensor2_Volt.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID25_OxySensor2_Volt))
	}
	if m.xxx_Service == 38 {
		md.S01PID26_OxySensor3_FAER.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID26_OxySensor3_FAER))
	}
	if m.xxx_Service == 38 {
		md.S01PID26_OxySensor3_Volt.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID26_OxySensor3_Volt))
	}
	if m.xxx_Service == 39 {
		md.S01PID27_OxySensor4_FAER.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID27_OxySensor4_FAER))
	}
	if m.xxx_Service == 40 {
		md.S01PID28_OxySensor5_FAER.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID28_OxySensor5_FAER))
	}
	if m.xxx_Service == 41 {
		md.S01PID29_OxySensor6_FAER.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID29_OxySensor6_FAER))
	}
	if m.xxx_Service == 39 {
		md.S01PID27_OxySensor4_Volt.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID27_OxySensor4_Volt))
	}
	if m.xxx_Service == 40 {
		md.S01PID28_OxySensor5_Volt.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID28_OxySensor5_Volt))
	}
	if m.xxx_Service == 41 {
		md.S01PID29_OxySensor6_Volt.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID29_OxySensor6_Volt))
	}
	if m.xxx_Service == 42 {
		md.S01PID2A_OxySensor7_FAER.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID2A_OxySensor7_FAER))
	}
	if m.xxx_Service == 42 {
		md.S01PID2A_OxySensor7_Volt.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID2A_OxySensor7_Volt))
	}
	if m.xxx_Service == 43 {
		md.S01PID2B_OxySensor8_FAER.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID2B_OxySensor8_FAER))
	}
	if m.xxx_Service == 43 {
		md.S01PID2B_OxySensor8_Volt.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID2B_OxySensor8_Volt))
	}
	if m.xxx_Service == 44 {
		md.S01PID2C_CmdEGR.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID2C_CmdEGR))
	}
	if m.xxx_Service == 45 {
		md.S01PID2D_EGRError.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID2D_EGRError))
	}
	if m.xxx_Service == 46 {
		md.S01PID2E_CmdEvapPurge.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID2E_CmdEvapPurge))
	}
	if m.xxx_Service == 47 {
		md.S01PID2F_FuelTankLevel.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID2F_FuelTankLevel))
	}
	if m.xxx_Service == 48 {
		md.S01PID30_WarmUpsSinceCodeClear.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID30_WarmUpsSinceCodeClear))
	}
	if m.xxx_Service == 49 {
		md.S01PID31_DistanceSinceCodeClear.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID31_DistanceSinceCodeClear))
	}
	if m.xxx_Service == 50 {
		md.S01PID32_EvapSysVaporPres.MarshalSigned(&f.Data, int64(m.xxx_S01PID32_EvapSysVaporPres))
	}
	if m.xxx_Service == 51 {
		md.S01PID33_AbsBaroPres.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID33_AbsBaroPres))
	}
	if m.xxx_Service == 52 {
		md.S01PID34_OxySensor1_FAER.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID34_OxySensor1_FAER))
	}
	if m.xxx_Service == 52 {
		md.S01PID34_OxySensor1_Crnt.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID34_OxySensor1_Crnt))
	}
	if m.xxx_Service == 53 {
		md.S01PID35_OxySensor2_FAER.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID35_OxySensor2_FAER))
	}
	if m.xxx_Service == 53 {
		md.S01PID35_OxySensor2_Crnt.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID35_OxySensor2_Crnt))
	}
	if m.xxx_Service == 54 {
		md.S01PID36_OxySensor3_FAER.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID36_OxySensor3_FAER))
	}
	if m.xxx_Service == 54 {
		md.S01PID36_OxySensor3_Crnt.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID36_OxySensor3_Crnt))
	}
	if m.xxx_Service == 55 {
		md.S01PID37_OxySensor4_FAER.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID37_OxySensor4_FAER))
	}
	if m.xxx_Service == 56 {
		md.S01PID38_OxySensor5_FAER.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID38_OxySensor5_FAER))
	}
	if m.xxx_Service == 57 {
		md.S01PID39_OxySensor6_FAER.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID39_OxySensor6_FAER))
	}
	if m.xxx_Service == 55 {
		md.S01PID37_OxySensor4_Crnt.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID37_OxySensor4_Crnt))
	}
	if m.xxx_Service == 56 {
		md.S01PID38_OxySensor5_Crnt.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID38_OxySensor5_Crnt))
	}
	if m.xxx_Service == 57 {
		md.S01PID39_OxySensor6_Crnt.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID39_OxySensor6_Crnt))
	}
	if m.xxx_Service == 58 {
		md.S01PID3A_OxySensor7_FAER.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID3A_OxySensor7_FAER))
	}
	if m.xxx_Service == 59 {
		md.S01PID3B_OxySensor8_FAER.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID3B_OxySensor8_FAER))
	}
	if m.xxx_Service == 60 {
		md.S01PID3C_CatTempBank1Sens1.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID3C_CatTempBank1Sens1))
	}
	if m.xxx_Service == 61 {
		md.S01PID3D_CatTempBank2Sens1.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID3D_CatTempBank2Sens1))
	}
	if m.xxx_Service == 58 {
		md.S01PID3A_OxySensor7_Crnt.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID3A_OxySensor7_Crnt))
	}
	if m.xxx_Service == 59 {
		md.S01PID3B_OxySensor8_Crnt.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID3B_OxySensor8_Crnt))
	}
	if m.xxx_Service == 62 {
		md.S01PID3E_CatTempBank1Sens2.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID3E_CatTempBank1Sens2))
	}
	if m.xxx_Service == 63 {
		md.S01PID3F_CatTempBank2Sens2.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID3F_CatTempBank2Sens2))
	}
	if m.xxx_Service == 64 {
		md.S01PID40_PIDsSupported_41_60.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID40_PIDsSupported_41_60))
	}
	if m.xxx_Service == 65 {
		md.S01PID41_MonStatusDriveCycle.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID41_MonStatusDriveCycle))
	}
	if m.xxx_Service == 66 {
		md.S01PID42_ControlModuleVolt.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID42_ControlModuleVolt))
	}
	if m.xxx_Service == 67 {
		md.S01PID43_AbsLoadValue.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID43_AbsLoadValue))
	}
	if m.xxx_Service == 68 {
		md.S01PID44_FuelAirCmdEquiv.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID44_FuelAirCmdEquiv))
	}
	if m.xxx_Service == 69 {
		md.S01PID45_RelThrottlePos.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID45_RelThrottlePos))
	}
	if m.xxx_Service == 70 {
		md.S01PID46_AmbientAirTemp.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID46_AmbientAirTemp))
	}
	if m.xxx_Service == 71 {
		md.S01PID47_AbsThrottlePosB.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID47_AbsThrottlePosB))
	}
	if m.xxx_Service == 72 {
		md.S01PID48_AbsThrottlePosC.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID48_AbsThrottlePosC))
	}
	if m.xxx_Service == 73 {
		md.S01PID49_AbsThrottlePosD.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID49_AbsThrottlePosD))
	}
	if m.xxx_Service == 74 {
		md.S01PID4A_AbsThrottlePosE.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID4A_AbsThrottlePosE))
	}
	if m.xxx_Service == 75 {
		md.S01PID4B_AbsThrottlePosF.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID4B_AbsThrottlePosF))
	}
	if m.xxx_Service == 76 {
		md.S01PID4C_CmdThrottleAct.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID4C_CmdThrottleAct))
	}
	if m.xxx_Service == 77 {
		md.S01PID4D_TimeRunMILOn.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID4D_TimeRunMILOn))
	}
	if m.xxx_Service == 78 {
		md.S01PID4E_TimeSinceCodeClear.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID4E_TimeSinceCodeClear))
	}
	if m.xxx_Service == 79 {
		md.S01PID4F_Max_FAER.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID4F_Max_FAER))
	}
	if m.xxx_Service == 79 {
		md.S01PID4F_Max_OxySensVol.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID4F_Max_OxySensVol))
	}
	if m.xxx_Service == 79 {
		md.S01PID4F_Max_OxySensCrnt.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID4F_Max_OxySensCrnt))
	}
	if m.xxx_Service == 79 {
		md.S01PID4F_Max_IntManiAbsPres.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID4F_Max_IntManiAbsPres))
	}
	if m.xxx_Service == 80 {
		md.S01PID50_Max_AirFlowMAF.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID50_Max_AirFlowMAF))
	}
	if m.xxx_Service == 81 {
		md.S01PID51_FuelType.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID51_FuelType))
	}
	if m.xxx_Service == 82 {
		md.S01PID52_EthanolFuelPct.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID52_EthanolFuelPct))
	}
	if m.xxx_Service == 83 {
		md.S01PID53_AbsEvapSysVapPres.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID53_AbsEvapSysVapPres))
	}
	if m.xxx_Service == 84 {
		md.S01PID54_EvapSysVapPres.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID54_EvapSysVapPres))
	}
	if m.xxx_Service == 85 {
		md.S01PID55_ShortSecOxyTrimBank1.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID55_ShortSecOxyTrimBank1))
	}
	if m.xxx_Service == 86 {
		md.S01PID56_LongSecOxyTrimBank1.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID56_LongSecOxyTrimBank1))
	}
	if m.xxx_Service == 85 {
		md.S01PID55_ShortSecOxyTrimBank3.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID55_ShortSecOxyTrimBank3))
	}
	if m.xxx_Service == 86 {
		md.S01PID56_LongSecOxyTrimBank3.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID56_LongSecOxyTrimBank3))
	}
	if m.xxx_Service == 87 {
		md.S01PID57_ShortSecOxyTrimBank2.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID57_ShortSecOxyTrimBank2))
	}
	if m.xxx_Service == 88 {
		md.S01PID58_LongSecOxyTrimBank2.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID58_LongSecOxyTrimBank2))
	}
	if m.xxx_Service == 89 {
		md.S01PID59_FuelRailAbsPres.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID59_FuelRailAbsPres))
	}
	if m.xxx_Service == 90 {
		md.S01PID5A_RelAccelPedalPos.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID5A_RelAccelPedalPos))
	}
	if m.xxx_Service == 87 {
		md.S01PID57_ShortSecOxyTrimBank4.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID57_ShortSecOxyTrimBank4))
	}
	if m.xxx_Service == 88 {
		md.S01PID58_LongSecOxyTrimBank4.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID58_LongSecOxyTrimBank4))
	}
	if m.xxx_Service == 91 {
		md.S01PID5B_HybrBatPackRemLife.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID5B_HybrBatPackRemLife))
	}
	if m.xxx_Service == 92 {
		md.S01PID5C_EngineOilTemp.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID5C_EngineOilTemp))
	}
	if m.xxx_Service == 93 {
		md.S01PID5D_FuelInjectionTiming.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID5D_FuelInjectionTiming))
	}
	if m.xxx_Service == 94 {
		md.S01PID5E_EngineFuelRate.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID5E_EngineFuelRate))
	}
	if m.xxx_Service == 95 {
		md.S01PID5F_EmissionReq.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID5F_EmissionReq))
	}
	if m.xxx_Service == 96 {
		md.S01PID60_PIDsSupported_61_80.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID60_PIDsSupported_61_80))
	}
	if m.xxx_Service == 97 {
		md.S01PID61_DemandEngTorqPct.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID61_DemandEngTorqPct))
	}
	if m.xxx_Service == 98 {
		md.S01PID62_ActualEngTorqPct.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID62_ActualEngTorqPct))
	}
	if m.xxx_Service == 99 {
		md.S01PID63_EngRefTorq.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID63_EngRefTorq))
	}
	if m.xxx_Service == 100 {
		md.S01PID64_EngPctTorq_Idle.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID64_EngPctTorq_Idle))
	}
	if m.xxx_Service == 100 {
		md.S01PID64_EngPctTorq_EP1.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID64_EngPctTorq_EP1))
	}
	if m.xxx_Service == 100 {
		md.S01PID64_EngPctTorq_EP2.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID64_EngPctTorq_EP2))
	}
	if m.xxx_Service == 100 {
		md.S01PID64_EngPctTorq_EP3.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID64_EngPctTorq_EP3))
	}
	if m.xxx_Service == 101 {
		md.S01PID65_AuxInputOutput.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID65_AuxInputOutput))
	}
	if m.xxx_Service == 102 {
		md.S01PID66_MAFSensor.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID66_MAFSensor))
	}
	if m.xxx_Service == 100 {
		md.S01PID64_EngPctTorq_EP4.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID64_EngPctTorq_EP4))
	}
	if m.xxx_Service == 103 {
		md.S01PID67_EngineCoolantTemp.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID67_EngineCoolantTemp))
	}
	if m.xxx_Service == 104 {
		md.S01PID68_IntakeAirTempSens.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID68_IntakeAirTempSens))
	}
	if m.xxx_Service == 105 {
		md.S01PID69_CmdEGR_EGRError.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID69_CmdEGR_EGRError))
	}
	if m.xxx_Service == 106 {
		md.S01PID6A_CmdDieselIntAir.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID6A_CmdDieselIntAir))
	}
	if m.xxx_Service == 107 {
		md.S01PID6B_ExhaustGasTemp.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID6B_ExhaustGasTemp))
	}
	if m.xxx_Service == 108 {
		md.S01PID6C_CmdThrottleActRel.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID6C_CmdThrottleActRel))
	}
	if m.xxx_Service == 109 {
		md.S01PID6D_FuelPresContrSys.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID6D_FuelPresContrSys))
	}
	if m.xxx_Service == 110 {
		md.S01PID6E_InjPresContrSys.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID6E_InjPresContrSys))
	}
	if m.xxx_Service == 111 {
		md.S01PID6F_TurboComprPres.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID6F_TurboComprPres))
	}
	if m.xxx_Service == 112 {
		md.S01PID70_BoostPresCntrl.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID70_BoostPresCntrl))
	}
	if m.xxx_Service == 128 {
		md.S01PID80_PIDsSupported_81_A0.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID80_PIDsSupported_81_A0))
	}
	if m.xxx_Service == 142 {
		md.S01PID8E_EngFrictionPctTorq.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PID8E_EngFrictionPctTorq))
	}
	if m.xxx_Service == 160 {
		md.S01PIDA0_PIDsSupported_A1_C0.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PIDA0_PIDsSupported_A1_C0))
	}
	if m.xxx_Service == 192 {
		md.S01PIDC0_PIDsSupported_C1_E0.MarshalUnsigned(&f.Data, uint64(m.xxx_S01PIDC0_PIDsSupported_C1_E0))
	}
	return f
}

// MarshalFrame encodes the message as a CAN frame.
func (m *OBD2) MarshalFrame() (can.Frame, error) {
	return m.Frame(), nil
}

// UnmarshalFrame decodes the message from a CAN frame.
func (m *OBD2) UnmarshalFrame(f can.Frame) error {
	md := Messages().OBD2
	switch {
	case f.ID != md.ID:
		return fmt.Errorf(
			"unmarshal OBD2: expects ID 417001749 (got %s with ID %d)", f.String(), f.ID,
		)
	case f.Length != md.Length.Length:
		return fmt.Errorf(
			"unmarshal OBD2: expects length 8 (got %s with length %d)", f.String(), f.Length,
		)
	case f.IsRemote:
		return fmt.Errorf(
			"unmarshal OBD2: expects non-remote frame (got remote frame %s)", f.String(),
		)
	case f.IsExtended != md.IsExtended:
		return fmt.Errorf(
			"unmarshal OBD2: expects extended ID (got %s with standard ID)", f.String(),
		)
	}
	m.xxx_Length = uint8(md.Length.UnmarshalUnsigned(f.Data))
	m.xxx_Service = OBD2_Service(md.Service.UnmarshalUnsigned(f.Data))
	m.xxx_Response = uint8(md.Response.UnmarshalUnsigned(f.Data))
	if m.xxx_Service == 1 {
		m.xxx_S01PID = OBD2_S01PID(md.S01PID.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 2 {
		m.xxx_S02PID = OBD2_S02PID(md.S02PID.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 0 {
		m.xxx_S01PID00_PIDsSupported_01_20 = uint32(md.S01PID00_PIDsSupported_01_20.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 1 {
		m.xxx_S01PID01_MonitorStatus = uint32(md.S01PID01_MonitorStatus.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 2 {
		m.xxx_S02PID02_FreezeDTC = uint16(md.S02PID02_FreezeDTC.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 2 {
		m.xxx_S01PID02_FreezeDTC = uint16(md.S01PID02_FreezeDTC.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 3 {
		m.xxx_S01PID03_FuelSystemStatus = OBD2_S01PID03_FuelSystemStatus(md.S01PID03_FuelSystemStatus.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 4 {
		m.xxx_S01PID04_CalcEngineLoad = uint8(md.S01PID04_CalcEngineLoad.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 5 {
		m.xxx_S01PID05_EngineCoolantTemp = uint8(md.S01PID05_EngineCoolantTemp.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 6 {
		m.xxx_S01PID06_ShortFuelTrimBank1 = uint8(md.S01PID06_ShortFuelTrimBank1.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 7 {
		m.xxx_S01PID07_LongFuelTrimBank1 = uint8(md.S01PID07_LongFuelTrimBank1.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 8 {
		m.xxx_S01PID08_ShortFuelTrimBank2 = uint8(md.S01PID08_ShortFuelTrimBank2.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 9 {
		m.xxx_S01PID09_LongFuelTrimBank2 = uint8(md.S01PID09_LongFuelTrimBank2.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 10 {
		m.xxx_S01PID0A_FuelPressure = uint8(md.S01PID0A_FuelPressure.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 11 {
		m.xxx_S01PID0B_IntakeManiAbsPress = uint8(md.S01PID0B_IntakeManiAbsPress.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 12 {
		m.xxx_S01PID0C_EngineRPM = uint16(md.S01PID0C_EngineRPM.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 13 {
		m.xxx_S01PID0D_VehicleSpeed = uint8(md.S01PID0D_VehicleSpeed.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 14 {
		m.xxx_S01PID0E_TimingAdvance = uint8(md.S01PID0E_TimingAdvance.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 15 {
		m.xxx_S01PID0F_IntakeAirTemperature = uint8(md.S01PID0F_IntakeAirTemperature.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 16 {
		m.xxx_S01PID10_MAFAirFlowRate = uint16(md.S01PID10_MAFAirFlowRate.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 17 {
		m.xxx_S01PID11_ThrottlePosition = uint8(md.S01PID11_ThrottlePosition.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 18 {
		m.xxx_S01PID12_CmdSecAirStatus = OBD2_S01PID12_CmdSecAirStatus(md.S01PID12_CmdSecAirStatus.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 20 {
		m.xxx_S01PID14_OxySensor1_Volt = uint8(md.S01PID14_OxySensor1_Volt.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 21 {
		m.xxx_S01PID15_OxySensor2_Volt = uint8(md.S01PID15_OxySensor2_Volt.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 22 {
		m.xxx_S01PID16_OxySensor3_Volt = uint8(md.S01PID16_OxySensor3_Volt.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 23 {
		m.xxx_S01PID17_OxySensor4_Volt = uint8(md.S01PID17_OxySensor4_Volt.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 24 {
		m.xxx_S01PID18_OxySensor5_Volt = uint8(md.S01PID18_OxySensor5_Volt.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 25 {
		m.xxx_S01PID19_OxySensor6_Volt = uint8(md.S01PID19_OxySensor6_Volt.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 20 {
		m.xxx_S01PID14_OxySensor1_STFT = uint8(md.S01PID14_OxySensor1_STFT.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 21 {
		m.xxx_S01PID15_OxySensor2_STFT = uint8(md.S01PID15_OxySensor2_STFT.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 22 {
		m.xxx_S01PID16_OxySensor3_STFT = uint8(md.S01PID16_OxySensor3_STFT.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 23 {
		m.xxx_S01PID17_OxySensor4_STFT = uint8(md.S01PID17_OxySensor4_STFT.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 24 {
		m.xxx_S01PID18_OxySensor5_STFT = uint8(md.S01PID18_OxySensor5_STFT.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 25 {
		m.xxx_S01PID19_OxySensor6_STFT = uint8(md.S01PID19_OxySensor6_STFT.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 26 {
		m.xxx_S01PID1A_OxySensor7_Volt = uint8(md.S01PID1A_OxySensor7_Volt.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 26 {
		m.xxx_S01PID1A_OxySensor7_STFT = uint8(md.S01PID1A_OxySensor7_STFT.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 27 {
		m.xxx_S01PID1B_OxySensor8_Volt = uint8(md.S01PID1B_OxySensor8_Volt.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 27 {
		m.xxx_S01PID1B_OxySensor8_STFT = uint8(md.S01PID1B_OxySensor8_STFT.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 28 {
		m.xxx_S01PID1C_OBDStandard = OBD2_S01PID1C_OBDStandard(md.S01PID1C_OBDStandard.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 31 {
		m.xxx_S01PID1F_TimeSinceEngStart = uint16(md.S01PID1F_TimeSinceEngStart.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 32 {
		m.xxx_S01PID20_PIDsSupported_21_40 = uint32(md.S01PID20_PIDsSupported_21_40.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 33 {
		m.xxx_S01PID21_DistanceMILOn = uint16(md.S01PID21_DistanceMILOn.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 34 {
		m.xxx_S01PID22_FuelRailPres = uint16(md.S01PID22_FuelRailPres.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 35 {
		m.xxx_S01PID23_FuelRailGaug = uint16(md.S01PID23_FuelRailGaug.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 36 {
		m.xxx_S01PID24_OxySensor1_FAER = uint16(md.S01PID24_OxySensor1_FAER.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 36 {
		m.xxx_S01PID24_OxySensor1_Volt = uint16(md.S01PID24_OxySensor1_Volt.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 37 {
		m.xxx_S01PID25_OxySensor2_FAER = uint16(md.S01PID25_OxySensor2_FAER.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 37 {
		m.xxx_S01PID25_OxySensor2_Volt = uint16(md.S01PID25_OxySensor2_Volt.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 38 {
		m.xxx_S01PID26_OxySensor3_FAER = uint16(md.S01PID26_OxySensor3_FAER.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 38 {
		m.xxx_S01PID26_OxySensor3_Volt = uint16(md.S01PID26_OxySensor3_Volt.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 39 {
		m.xxx_S01PID27_OxySensor4_FAER = uint16(md.S01PID27_OxySensor4_FAER.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 40 {
		m.xxx_S01PID28_OxySensor5_FAER = uint16(md.S01PID28_OxySensor5_FAER.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 41 {
		m.xxx_S01PID29_OxySensor6_FAER = uint16(md.S01PID29_OxySensor6_FAER.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 39 {
		m.xxx_S01PID27_OxySensor4_Volt = uint16(md.S01PID27_OxySensor4_Volt.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 40 {
		m.xxx_S01PID28_OxySensor5_Volt = uint16(md.S01PID28_OxySensor5_Volt.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 41 {
		m.xxx_S01PID29_OxySensor6_Volt = uint16(md.S01PID29_OxySensor6_Volt.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 42 {
		m.xxx_S01PID2A_OxySensor7_FAER = uint16(md.S01PID2A_OxySensor7_FAER.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 42 {
		m.xxx_S01PID2A_OxySensor7_Volt = uint16(md.S01PID2A_OxySensor7_Volt.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 43 {
		m.xxx_S01PID2B_OxySensor8_FAER = uint16(md.S01PID2B_OxySensor8_FAER.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 43 {
		m.xxx_S01PID2B_OxySensor8_Volt = uint16(md.S01PID2B_OxySensor8_Volt.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 44 {
		m.xxx_S01PID2C_CmdEGR = uint8(md.S01PID2C_CmdEGR.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 45 {
		m.xxx_S01PID2D_EGRError = uint8(md.S01PID2D_EGRError.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 46 {
		m.xxx_S01PID2E_CmdEvapPurge = uint8(md.S01PID2E_CmdEvapPurge.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 47 {
		m.xxx_S01PID2F_FuelTankLevel = uint8(md.S01PID2F_FuelTankLevel.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 48 {
		m.xxx_S01PID30_WarmUpsSinceCodeClear = uint8(md.S01PID30_WarmUpsSinceCodeClear.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 49 {
		m.xxx_S01PID31_DistanceSinceCodeClear = uint16(md.S01PID31_DistanceSinceCodeClear.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 50 {
		m.xxx_S01PID32_EvapSysVaporPres = int16(md.S01PID32_EvapSysVaporPres.UnmarshalSigned(f.Data))
	}
	if m.xxx_Service == 51 {
		m.xxx_S01PID33_AbsBaroPres = uint8(md.S01PID33_AbsBaroPres.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 52 {
		m.xxx_S01PID34_OxySensor1_FAER = uint16(md.S01PID34_OxySensor1_FAER.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 52 {
		m.xxx_S01PID34_OxySensor1_Crnt = uint16(md.S01PID34_OxySensor1_Crnt.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 53 {
		m.xxx_S01PID35_OxySensor2_FAER = uint16(md.S01PID35_OxySensor2_FAER.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 53 {
		m.xxx_S01PID35_OxySensor2_Crnt = uint16(md.S01PID35_OxySensor2_Crnt.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 54 {
		m.xxx_S01PID36_OxySensor3_FAER = uint16(md.S01PID36_OxySensor3_FAER.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 54 {
		m.xxx_S01PID36_OxySensor3_Crnt = uint16(md.S01PID36_OxySensor3_Crnt.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 55 {
		m.xxx_S01PID37_OxySensor4_FAER = uint16(md.S01PID37_OxySensor4_FAER.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 56 {
		m.xxx_S01PID38_OxySensor5_FAER = uint16(md.S01PID38_OxySensor5_FAER.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 57 {
		m.xxx_S01PID39_OxySensor6_FAER = uint16(md.S01PID39_OxySensor6_FAER.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 55 {
		m.xxx_S01PID37_OxySensor4_Crnt = uint16(md.S01PID37_OxySensor4_Crnt.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 56 {
		m.xxx_S01PID38_OxySensor5_Crnt = uint16(md.S01PID38_OxySensor5_Crnt.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 57 {
		m.xxx_S01PID39_OxySensor6_Crnt = uint16(md.S01PID39_OxySensor6_Crnt.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 58 {
		m.xxx_S01PID3A_OxySensor7_FAER = uint16(md.S01PID3A_OxySensor7_FAER.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 59 {
		m.xxx_S01PID3B_OxySensor8_FAER = uint16(md.S01PID3B_OxySensor8_FAER.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 60 {
		m.xxx_S01PID3C_CatTempBank1Sens1 = uint16(md.S01PID3C_CatTempBank1Sens1.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 61 {
		m.xxx_S01PID3D_CatTempBank2Sens1 = uint16(md.S01PID3D_CatTempBank2Sens1.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 58 {
		m.xxx_S01PID3A_OxySensor7_Crnt = uint16(md.S01PID3A_OxySensor7_Crnt.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 59 {
		m.xxx_S01PID3B_OxySensor8_Crnt = uint16(md.S01PID3B_OxySensor8_Crnt.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 62 {
		m.xxx_S01PID3E_CatTempBank1Sens2 = uint16(md.S01PID3E_CatTempBank1Sens2.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 63 {
		m.xxx_S01PID3F_CatTempBank2Sens2 = uint16(md.S01PID3F_CatTempBank2Sens2.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 64 {
		m.xxx_S01PID40_PIDsSupported_41_60 = uint32(md.S01PID40_PIDsSupported_41_60.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 65 {
		m.xxx_S01PID41_MonStatusDriveCycle = uint32(md.S01PID41_MonStatusDriveCycle.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 66 {
		m.xxx_S01PID42_ControlModuleVolt = uint16(md.S01PID42_ControlModuleVolt.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 67 {
		m.xxx_S01PID43_AbsLoadValue = uint16(md.S01PID43_AbsLoadValue.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 68 {
		m.xxx_S01PID44_FuelAirCmdEquiv = uint16(md.S01PID44_FuelAirCmdEquiv.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 69 {
		m.xxx_S01PID45_RelThrottlePos = uint8(md.S01PID45_RelThrottlePos.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 70 {
		m.xxx_S01PID46_AmbientAirTemp = uint8(md.S01PID46_AmbientAirTemp.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 71 {
		m.xxx_S01PID47_AbsThrottlePosB = uint8(md.S01PID47_AbsThrottlePosB.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 72 {
		m.xxx_S01PID48_AbsThrottlePosC = uint8(md.S01PID48_AbsThrottlePosC.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 73 {
		m.xxx_S01PID49_AbsThrottlePosD = uint8(md.S01PID49_AbsThrottlePosD.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 74 {
		m.xxx_S01PID4A_AbsThrottlePosE = uint8(md.S01PID4A_AbsThrottlePosE.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 75 {
		m.xxx_S01PID4B_AbsThrottlePosF = uint8(md.S01PID4B_AbsThrottlePosF.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 76 {
		m.xxx_S01PID4C_CmdThrottleAct = uint8(md.S01PID4C_CmdThrottleAct.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 77 {
		m.xxx_S01PID4D_TimeRunMILOn = uint16(md.S01PID4D_TimeRunMILOn.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 78 {
		m.xxx_S01PID4E_TimeSinceCodeClear = uint16(md.S01PID4E_TimeSinceCodeClear.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 79 {
		m.xxx_S01PID4F_Max_FAER = uint8(md.S01PID4F_Max_FAER.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 79 {
		m.xxx_S01PID4F_Max_OxySensVol = uint8(md.S01PID4F_Max_OxySensVol.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 79 {
		m.xxx_S01PID4F_Max_OxySensCrnt = uint8(md.S01PID4F_Max_OxySensCrnt.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 79 {
		m.xxx_S01PID4F_Max_IntManiAbsPres = uint8(md.S01PID4F_Max_IntManiAbsPres.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 80 {
		m.xxx_S01PID50_Max_AirFlowMAF = uint8(md.S01PID50_Max_AirFlowMAF.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 81 {
		m.xxx_S01PID51_FuelType = OBD2_S01PID51_FuelType(md.S01PID51_FuelType.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 82 {
		m.xxx_S01PID52_EthanolFuelPct = uint8(md.S01PID52_EthanolFuelPct.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 83 {
		m.xxx_S01PID53_AbsEvapSysVapPres = uint16(md.S01PID53_AbsEvapSysVapPres.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 84 {
		m.xxx_S01PID54_EvapSysVapPres = uint16(md.S01PID54_EvapSysVapPres.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 85 {
		m.xxx_S01PID55_ShortSecOxyTrimBank1 = uint8(md.S01PID55_ShortSecOxyTrimBank1.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 86 {
		m.xxx_S01PID56_LongSecOxyTrimBank1 = uint8(md.S01PID56_LongSecOxyTrimBank1.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 85 {
		m.xxx_S01PID55_ShortSecOxyTrimBank3 = uint8(md.S01PID55_ShortSecOxyTrimBank3.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 86 {
		m.xxx_S01PID56_LongSecOxyTrimBank3 = uint8(md.S01PID56_LongSecOxyTrimBank3.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 87 {
		m.xxx_S01PID57_ShortSecOxyTrimBank2 = uint8(md.S01PID57_ShortSecOxyTrimBank2.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 88 {
		m.xxx_S01PID58_LongSecOxyTrimBank2 = uint8(md.S01PID58_LongSecOxyTrimBank2.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 89 {
		m.xxx_S01PID59_FuelRailAbsPres = uint16(md.S01PID59_FuelRailAbsPres.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 90 {
		m.xxx_S01PID5A_RelAccelPedalPos = uint8(md.S01PID5A_RelAccelPedalPos.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 87 {
		m.xxx_S01PID57_ShortSecOxyTrimBank4 = uint8(md.S01PID57_ShortSecOxyTrimBank4.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 88 {
		m.xxx_S01PID58_LongSecOxyTrimBank4 = uint8(md.S01PID58_LongSecOxyTrimBank4.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 91 {
		m.xxx_S01PID5B_HybrBatPackRemLife = uint8(md.S01PID5B_HybrBatPackRemLife.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 92 {
		m.xxx_S01PID5C_EngineOilTemp = uint8(md.S01PID5C_EngineOilTemp.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 93 {
		m.xxx_S01PID5D_FuelInjectionTiming = uint16(md.S01PID5D_FuelInjectionTiming.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 94 {
		m.xxx_S01PID5E_EngineFuelRate = uint16(md.S01PID5E_EngineFuelRate.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 95 {
		m.xxx_S01PID5F_EmissionReq = uint8(md.S01PID5F_EmissionReq.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 96 {
		m.xxx_S01PID60_PIDsSupported_61_80 = uint32(md.S01PID60_PIDsSupported_61_80.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 97 {
		m.xxx_S01PID61_DemandEngTorqPct = uint8(md.S01PID61_DemandEngTorqPct.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 98 {
		m.xxx_S01PID62_ActualEngTorqPct = uint8(md.S01PID62_ActualEngTorqPct.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 99 {
		m.xxx_S01PID63_EngRefTorq = uint16(md.S01PID63_EngRefTorq.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 100 {
		m.xxx_S01PID64_EngPctTorq_Idle = uint8(md.S01PID64_EngPctTorq_Idle.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 100 {
		m.xxx_S01PID64_EngPctTorq_EP1 = uint8(md.S01PID64_EngPctTorq_EP1.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 100 {
		m.xxx_S01PID64_EngPctTorq_EP2 = uint8(md.S01PID64_EngPctTorq_EP2.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 100 {
		m.xxx_S01PID64_EngPctTorq_EP3 = uint8(md.S01PID64_EngPctTorq_EP3.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 101 {
		m.xxx_S01PID65_AuxInputOutput = uint8(md.S01PID65_AuxInputOutput.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 102 {
		m.xxx_S01PID66_MAFSensor = uint8(md.S01PID66_MAFSensor.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 100 {
		m.xxx_S01PID64_EngPctTorq_EP4 = uint8(md.S01PID64_EngPctTorq_EP4.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 103 {
		m.xxx_S01PID67_EngineCoolantTemp = uint8(md.S01PID67_EngineCoolantTemp.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 104 {
		m.xxx_S01PID68_IntakeAirTempSens = uint8(md.S01PID68_IntakeAirTempSens.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 105 {
		m.xxx_S01PID69_CmdEGR_EGRError = uint8(md.S01PID69_CmdEGR_EGRError.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 106 {
		m.xxx_S01PID6A_CmdDieselIntAir = uint8(md.S01PID6A_CmdDieselIntAir.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 107 {
		m.xxx_S01PID6B_ExhaustGasTemp = uint8(md.S01PID6B_ExhaustGasTemp.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 108 {
		m.xxx_S01PID6C_CmdThrottleActRel = uint8(md.S01PID6C_CmdThrottleActRel.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 109 {
		m.xxx_S01PID6D_FuelPresContrSys = uint8(md.S01PID6D_FuelPresContrSys.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 110 {
		m.xxx_S01PID6E_InjPresContrSys = uint8(md.S01PID6E_InjPresContrSys.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 111 {
		m.xxx_S01PID6F_TurboComprPres = uint8(md.S01PID6F_TurboComprPres.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 112 {
		m.xxx_S01PID70_BoostPresCntrl = uint8(md.S01PID70_BoostPresCntrl.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 128 {
		m.xxx_S01PID80_PIDsSupported_81_A0 = uint32(md.S01PID80_PIDsSupported_81_A0.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 142 {
		m.xxx_S01PID8E_EngFrictionPctTorq = uint8(md.S01PID8E_EngFrictionPctTorq.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 160 {
		m.xxx_S01PIDA0_PIDsSupported_A1_C0 = uint32(md.S01PIDA0_PIDsSupported_A1_C0.UnmarshalUnsigned(f.Data))
	}
	if m.xxx_Service == 192 {
		m.xxx_S01PIDC0_PIDsSupported_C1_E0 = uint32(md.S01PIDC0_PIDsSupported_C1_E0.UnmarshalUnsigned(f.Data))
	}
	return nil
}

// Nodes returns the bit29 node descriptors.
func Nodes() *NodesDescriptor {
	return nd
}

// NodesDescriptor contains all bit29 node descriptors.
type NodesDescriptor struct {
}

// Messages returns the bit29 message descriptors.
func Messages() *MessagesDescriptor {
	return md
}

// MessagesDescriptor contains all bit29 message descriptors.
type MessagesDescriptor struct {
	OBD2 *OBD2Descriptor
}

// UnmarshalFrame unmarshals the provided bit29 CAN frame.
func (md *MessagesDescriptor) UnmarshalFrame(f can.Frame) (generated.Message, error) {
	switch f.ID {
	case md.OBD2.ID:
		var msg OBD2
		if err := msg.UnmarshalFrame(f); err != nil {
			return nil, fmt.Errorf("unmarshal bit29 frame: %w", err)
		}
		return &msg, nil
	default:
		return nil, fmt.Errorf("unmarshal bit29 frame: ID not in database: %d", f.ID)
	}
}

type OBD2Descriptor struct {
	*descriptor.Message
	Length                          *descriptor.Signal
	Service                         *descriptor.Signal
	Response                        *descriptor.Signal
	S01PID                          *descriptor.Signal
	S02PID                          *descriptor.Signal
	S01PID00_PIDsSupported_01_20    *descriptor.Signal
	S01PID01_MonitorStatus          *descriptor.Signal
	S02PID02_FreezeDTC              *descriptor.Signal
	S01PID02_FreezeDTC              *descriptor.Signal
	S01PID03_FuelSystemStatus       *descriptor.Signal
	S01PID04_CalcEngineLoad         *descriptor.Signal
	S01PID05_EngineCoolantTemp      *descriptor.Signal
	S01PID06_ShortFuelTrimBank1     *descriptor.Signal
	S01PID07_LongFuelTrimBank1      *descriptor.Signal
	S01PID08_ShortFuelTrimBank2     *descriptor.Signal
	S01PID09_LongFuelTrimBank2      *descriptor.Signal
	S01PID0A_FuelPressure           *descriptor.Signal
	S01PID0B_IntakeManiAbsPress     *descriptor.Signal
	S01PID0C_EngineRPM              *descriptor.Signal
	S01PID0D_VehicleSpeed           *descriptor.Signal
	S01PID0E_TimingAdvance          *descriptor.Signal
	S01PID0F_IntakeAirTemperature   *descriptor.Signal
	S01PID10_MAFAirFlowRate         *descriptor.Signal
	S01PID11_ThrottlePosition       *descriptor.Signal
	S01PID12_CmdSecAirStatus        *descriptor.Signal
	S01PID14_OxySensor1_Volt        *descriptor.Signal
	S01PID15_OxySensor2_Volt        *descriptor.Signal
	S01PID16_OxySensor3_Volt        *descriptor.Signal
	S01PID17_OxySensor4_Volt        *descriptor.Signal
	S01PID18_OxySensor5_Volt        *descriptor.Signal
	S01PID19_OxySensor6_Volt        *descriptor.Signal
	S01PID14_OxySensor1_STFT        *descriptor.Signal
	S01PID15_OxySensor2_STFT        *descriptor.Signal
	S01PID16_OxySensor3_STFT        *descriptor.Signal
	S01PID17_OxySensor4_STFT        *descriptor.Signal
	S01PID18_OxySensor5_STFT        *descriptor.Signal
	S01PID19_OxySensor6_STFT        *descriptor.Signal
	S01PID1A_OxySensor7_Volt        *descriptor.Signal
	S01PID1A_OxySensor7_STFT        *descriptor.Signal
	S01PID1B_OxySensor8_Volt        *descriptor.Signal
	S01PID1B_OxySensor8_STFT        *descriptor.Signal
	S01PID1C_OBDStandard            *descriptor.Signal
	S01PID1F_TimeSinceEngStart      *descriptor.Signal
	S01PID20_PIDsSupported_21_40    *descriptor.Signal
	S01PID21_DistanceMILOn          *descriptor.Signal
	S01PID22_FuelRailPres           *descriptor.Signal
	S01PID23_FuelRailGaug           *descriptor.Signal
	S01PID24_OxySensor1_FAER        *descriptor.Signal
	S01PID24_OxySensor1_Volt        *descriptor.Signal
	S01PID25_OxySensor2_FAER        *descriptor.Signal
	S01PID25_OxySensor2_Volt        *descriptor.Signal
	S01PID26_OxySensor3_FAER        *descriptor.Signal
	S01PID26_OxySensor3_Volt        *descriptor.Signal
	S01PID27_OxySensor4_FAER        *descriptor.Signal
	S01PID28_OxySensor5_FAER        *descriptor.Signal
	S01PID29_OxySensor6_FAER        *descriptor.Signal
	S01PID27_OxySensor4_Volt        *descriptor.Signal
	S01PID28_OxySensor5_Volt        *descriptor.Signal
	S01PID29_OxySensor6_Volt        *descriptor.Signal
	S01PID2A_OxySensor7_FAER        *descriptor.Signal
	S01PID2A_OxySensor7_Volt        *descriptor.Signal
	S01PID2B_OxySensor8_FAER        *descriptor.Signal
	S01PID2B_OxySensor8_Volt        *descriptor.Signal
	S01PID2C_CmdEGR                 *descriptor.Signal
	S01PID2D_EGRError               *descriptor.Signal
	S01PID2E_CmdEvapPurge           *descriptor.Signal
	S01PID2F_FuelTankLevel          *descriptor.Signal
	S01PID30_WarmUpsSinceCodeClear  *descriptor.Signal
	S01PID31_DistanceSinceCodeClear *descriptor.Signal
	S01PID32_EvapSysVaporPres       *descriptor.Signal
	S01PID33_AbsBaroPres            *descriptor.Signal
	S01PID34_OxySensor1_FAER        *descriptor.Signal
	S01PID34_OxySensor1_Crnt        *descriptor.Signal
	S01PID35_OxySensor2_FAER        *descriptor.Signal
	S01PID35_OxySensor2_Crnt        *descriptor.Signal
	S01PID36_OxySensor3_FAER        *descriptor.Signal
	S01PID36_OxySensor3_Crnt        *descriptor.Signal
	S01PID37_OxySensor4_FAER        *descriptor.Signal
	S01PID38_OxySensor5_FAER        *descriptor.Signal
	S01PID39_OxySensor6_FAER        *descriptor.Signal
	S01PID37_OxySensor4_Crnt        *descriptor.Signal
	S01PID38_OxySensor5_Crnt        *descriptor.Signal
	S01PID39_OxySensor6_Crnt        *descriptor.Signal
	S01PID3A_OxySensor7_FAER        *descriptor.Signal
	S01PID3B_OxySensor8_FAER        *descriptor.Signal
	S01PID3C_CatTempBank1Sens1      *descriptor.Signal
	S01PID3D_CatTempBank2Sens1      *descriptor.Signal
	S01PID3A_OxySensor7_Crnt        *descriptor.Signal
	S01PID3B_OxySensor8_Crnt        *descriptor.Signal
	S01PID3E_CatTempBank1Sens2      *descriptor.Signal
	S01PID3F_CatTempBank2Sens2      *descriptor.Signal
	S01PID40_PIDsSupported_41_60    *descriptor.Signal
	S01PID41_MonStatusDriveCycle    *descriptor.Signal
	S01PID42_ControlModuleVolt      *descriptor.Signal
	S01PID43_AbsLoadValue           *descriptor.Signal
	S01PID44_FuelAirCmdEquiv        *descriptor.Signal
	S01PID45_RelThrottlePos         *descriptor.Signal
	S01PID46_AmbientAirTemp         *descriptor.Signal
	S01PID47_AbsThrottlePosB        *descriptor.Signal
	S01PID48_AbsThrottlePosC        *descriptor.Signal
	S01PID49_AbsThrottlePosD        *descriptor.Signal
	S01PID4A_AbsThrottlePosE        *descriptor.Signal
	S01PID4B_AbsThrottlePosF        *descriptor.Signal
	S01PID4C_CmdThrottleAct         *descriptor.Signal
	S01PID4D_TimeRunMILOn           *descriptor.Signal
	S01PID4E_TimeSinceCodeClear     *descriptor.Signal
	S01PID4F_Max_FAER               *descriptor.Signal
	S01PID4F_Max_OxySensVol         *descriptor.Signal
	S01PID4F_Max_OxySensCrnt        *descriptor.Signal
	S01PID4F_Max_IntManiAbsPres     *descriptor.Signal
	S01PID50_Max_AirFlowMAF         *descriptor.Signal
	S01PID51_FuelType               *descriptor.Signal
	S01PID52_EthanolFuelPct         *descriptor.Signal
	S01PID53_AbsEvapSysVapPres      *descriptor.Signal
	S01PID54_EvapSysVapPres         *descriptor.Signal
	S01PID55_ShortSecOxyTrimBank1   *descriptor.Signal
	S01PID56_LongSecOxyTrimBank1    *descriptor.Signal
	S01PID55_ShortSecOxyTrimBank3   *descriptor.Signal
	S01PID56_LongSecOxyTrimBank3    *descriptor.Signal
	S01PID57_ShortSecOxyTrimBank2   *descriptor.Signal
	S01PID58_LongSecOxyTrimBank2    *descriptor.Signal
	S01PID59_FuelRailAbsPres        *descriptor.Signal
	S01PID5A_RelAccelPedalPos       *descriptor.Signal
	S01PID57_ShortSecOxyTrimBank4   *descriptor.Signal
	S01PID58_LongSecOxyTrimBank4    *descriptor.Signal
	S01PID5B_HybrBatPackRemLife     *descriptor.Signal
	S01PID5C_EngineOilTemp          *descriptor.Signal
	S01PID5D_FuelInjectionTiming    *descriptor.Signal
	S01PID5E_EngineFuelRate         *descriptor.Signal
	S01PID5F_EmissionReq            *descriptor.Signal
	S01PID60_PIDsSupported_61_80    *descriptor.Signal
	S01PID61_DemandEngTorqPct       *descriptor.Signal
	S01PID62_ActualEngTorqPct       *descriptor.Signal
	S01PID63_EngRefTorq             *descriptor.Signal
	S01PID64_EngPctTorq_Idle        *descriptor.Signal
	S01PID64_EngPctTorq_EP1         *descriptor.Signal
	S01PID64_EngPctTorq_EP2         *descriptor.Signal
	S01PID64_EngPctTorq_EP3         *descriptor.Signal
	S01PID65_AuxInputOutput         *descriptor.Signal
	S01PID66_MAFSensor              *descriptor.Signal
	S01PID64_EngPctTorq_EP4         *descriptor.Signal
	S01PID67_EngineCoolantTemp      *descriptor.Signal
	S01PID68_IntakeAirTempSens      *descriptor.Signal
	S01PID69_CmdEGR_EGRError        *descriptor.Signal
	S01PID6A_CmdDieselIntAir        *descriptor.Signal
	S01PID6B_ExhaustGasTemp         *descriptor.Signal
	S01PID6C_CmdThrottleActRel      *descriptor.Signal
	S01PID6D_FuelPresContrSys       *descriptor.Signal
	S01PID6E_InjPresContrSys        *descriptor.Signal
	S01PID6F_TurboComprPres         *descriptor.Signal
	S01PID70_BoostPresCntrl         *descriptor.Signal
	S01PID80_PIDsSupported_81_A0    *descriptor.Signal
	S01PID8E_EngFrictionPctTorq     *descriptor.Signal
	S01PIDA0_PIDsSupported_A1_C0    *descriptor.Signal
	S01PIDC0_PIDsSupported_C1_E0    *descriptor.Signal
}

// Database returns the bit29 database descriptor.
func (md *MessagesDescriptor) Database() *descriptor.Database {
	return d
}

var nd = &NodesDescriptor{}

var md = &MessagesDescriptor{
	OBD2: &OBD2Descriptor{
		Message:                         d.Messages[0],
		Length:                          d.Messages[0].Signals[0],
		Service:                         d.Messages[0].Signals[1],
		Response:                        d.Messages[0].Signals[2],
		S01PID:                          d.Messages[0].Signals[3],
		S02PID:                          d.Messages[0].Signals[4],
		S01PID00_PIDsSupported_01_20:    d.Messages[0].Signals[5],
		S01PID01_MonitorStatus:          d.Messages[0].Signals[6],
		S02PID02_FreezeDTC:              d.Messages[0].Signals[7],
		S01PID02_FreezeDTC:              d.Messages[0].Signals[8],
		S01PID03_FuelSystemStatus:       d.Messages[0].Signals[9],
		S01PID04_CalcEngineLoad:         d.Messages[0].Signals[10],
		S01PID05_EngineCoolantTemp:      d.Messages[0].Signals[11],
		S01PID06_ShortFuelTrimBank1:     d.Messages[0].Signals[12],
		S01PID07_LongFuelTrimBank1:      d.Messages[0].Signals[13],
		S01PID08_ShortFuelTrimBank2:     d.Messages[0].Signals[14],
		S01PID09_LongFuelTrimBank2:      d.Messages[0].Signals[15],
		S01PID0A_FuelPressure:           d.Messages[0].Signals[16],
		S01PID0B_IntakeManiAbsPress:     d.Messages[0].Signals[17],
		S01PID0C_EngineRPM:              d.Messages[0].Signals[18],
		S01PID0D_VehicleSpeed:           d.Messages[0].Signals[19],
		S01PID0E_TimingAdvance:          d.Messages[0].Signals[20],
		S01PID0F_IntakeAirTemperature:   d.Messages[0].Signals[21],
		S01PID10_MAFAirFlowRate:         d.Messages[0].Signals[22],
		S01PID11_ThrottlePosition:       d.Messages[0].Signals[23],
		S01PID12_CmdSecAirStatus:        d.Messages[0].Signals[24],
		S01PID14_OxySensor1_Volt:        d.Messages[0].Signals[25],
		S01PID15_OxySensor2_Volt:        d.Messages[0].Signals[26],
		S01PID16_OxySensor3_Volt:        d.Messages[0].Signals[27],
		S01PID17_OxySensor4_Volt:        d.Messages[0].Signals[28],
		S01PID18_OxySensor5_Volt:        d.Messages[0].Signals[29],
		S01PID19_OxySensor6_Volt:        d.Messages[0].Signals[30],
		S01PID14_OxySensor1_STFT:        d.Messages[0].Signals[31],
		S01PID15_OxySensor2_STFT:        d.Messages[0].Signals[32],
		S01PID16_OxySensor3_STFT:        d.Messages[0].Signals[33],
		S01PID17_OxySensor4_STFT:        d.Messages[0].Signals[34],
		S01PID18_OxySensor5_STFT:        d.Messages[0].Signals[35],
		S01PID19_OxySensor6_STFT:        d.Messages[0].Signals[36],
		S01PID1A_OxySensor7_Volt:        d.Messages[0].Signals[37],
		S01PID1A_OxySensor7_STFT:        d.Messages[0].Signals[38],
		S01PID1B_OxySensor8_Volt:        d.Messages[0].Signals[39],
		S01PID1B_OxySensor8_STFT:        d.Messages[0].Signals[40],
		S01PID1C_OBDStandard:            d.Messages[0].Signals[41],
		S01PID1F_TimeSinceEngStart:      d.Messages[0].Signals[42],
		S01PID20_PIDsSupported_21_40:    d.Messages[0].Signals[43],
		S01PID21_DistanceMILOn:          d.Messages[0].Signals[44],
		S01PID22_FuelRailPres:           d.Messages[0].Signals[45],
		S01PID23_FuelRailGaug:           d.Messages[0].Signals[46],
		S01PID24_OxySensor1_FAER:        d.Messages[0].Signals[47],
		S01PID24_OxySensor1_Volt:        d.Messages[0].Signals[48],
		S01PID25_OxySensor2_FAER:        d.Messages[0].Signals[49],
		S01PID25_OxySensor2_Volt:        d.Messages[0].Signals[50],
		S01PID26_OxySensor3_FAER:        d.Messages[0].Signals[51],
		S01PID26_OxySensor3_Volt:        d.Messages[0].Signals[52],
		S01PID27_OxySensor4_FAER:        d.Messages[0].Signals[53],
		S01PID28_OxySensor5_FAER:        d.Messages[0].Signals[54],
		S01PID29_OxySensor6_FAER:        d.Messages[0].Signals[55],
		S01PID27_OxySensor4_Volt:        d.Messages[0].Signals[56],
		S01PID28_OxySensor5_Volt:        d.Messages[0].Signals[57],
		S01PID29_OxySensor6_Volt:        d.Messages[0].Signals[58],
		S01PID2A_OxySensor7_FAER:        d.Messages[0].Signals[59],
		S01PID2A_OxySensor7_Volt:        d.Messages[0].Signals[60],
		S01PID2B_OxySensor8_FAER:        d.Messages[0].Signals[61],
		S01PID2B_OxySensor8_Volt:        d.Messages[0].Signals[62],
		S01PID2C_CmdEGR:                 d.Messages[0].Signals[63],
		S01PID2D_EGRError:               d.Messages[0].Signals[64],
		S01PID2E_CmdEvapPurge:           d.Messages[0].Signals[65],
		S01PID2F_FuelTankLevel:          d.Messages[0].Signals[66],
		S01PID30_WarmUpsSinceCodeClear:  d.Messages[0].Signals[67],
		S01PID31_DistanceSinceCodeClear: d.Messages[0].Signals[68],
		S01PID32_EvapSysVaporPres:       d.Messages[0].Signals[69],
		S01PID33_AbsBaroPres:            d.Messages[0].Signals[70],
		S01PID34_OxySensor1_FAER:        d.Messages[0].Signals[71],
		S01PID34_OxySensor1_Crnt:        d.Messages[0].Signals[72],
		S01PID35_OxySensor2_FAER:        d.Messages[0].Signals[73],
		S01PID35_OxySensor2_Crnt:        d.Messages[0].Signals[74],
		S01PID36_OxySensor3_FAER:        d.Messages[0].Signals[75],
		S01PID36_OxySensor3_Crnt:        d.Messages[0].Signals[76],
		S01PID37_OxySensor4_FAER:        d.Messages[0].Signals[77],
		S01PID38_OxySensor5_FAER:        d.Messages[0].Signals[78],
		S01PID39_OxySensor6_FAER:        d.Messages[0].Signals[79],
		S01PID37_OxySensor4_Crnt:        d.Messages[0].Signals[80],
		S01PID38_OxySensor5_Crnt:        d.Messages[0].Signals[81],
		S01PID39_OxySensor6_Crnt:        d.Messages[0].Signals[82],
		S01PID3A_OxySensor7_FAER:        d.Messages[0].Signals[83],
		S01PID3B_OxySensor8_FAER:        d.Messages[0].Signals[84],
		S01PID3C_CatTempBank1Sens1:      d.Messages[0].Signals[85],
		S01PID3D_CatTempBank2Sens1:      d.Messages[0].Signals[86],
		S01PID3A_OxySensor7_Crnt:        d.Messages[0].Signals[87],
		S01PID3B_OxySensor8_Crnt:        d.Messages[0].Signals[88],
		S01PID3E_CatTempBank1Sens2:      d.Messages[0].Signals[89],
		S01PID3F_CatTempBank2Sens2:      d.Messages[0].Signals[90],
		S01PID40_PIDsSupported_41_60:    d.Messages[0].Signals[91],
		S01PID41_MonStatusDriveCycle:    d.Messages[0].Signals[92],
		S01PID42_ControlModuleVolt:      d.Messages[0].Signals[93],
		S01PID43_AbsLoadValue:           d.Messages[0].Signals[94],
		S01PID44_FuelAirCmdEquiv:        d.Messages[0].Signals[95],
		S01PID45_RelThrottlePos:         d.Messages[0].Signals[96],
		S01PID46_AmbientAirTemp:         d.Messages[0].Signals[97],
		S01PID47_AbsThrottlePosB:        d.Messages[0].Signals[98],
		S01PID48_AbsThrottlePosC:        d.Messages[0].Signals[99],
		S01PID49_AbsThrottlePosD:        d.Messages[0].Signals[100],
		S01PID4A_AbsThrottlePosE:        d.Messages[0].Signals[101],
		S01PID4B_AbsThrottlePosF:        d.Messages[0].Signals[102],
		S01PID4C_CmdThrottleAct:         d.Messages[0].Signals[103],
		S01PID4D_TimeRunMILOn:           d.Messages[0].Signals[104],
		S01PID4E_TimeSinceCodeClear:     d.Messages[0].Signals[105],
		S01PID4F_Max_FAER:               d.Messages[0].Signals[106],
		S01PID4F_Max_OxySensVol:         d.Messages[0].Signals[107],
		S01PID4F_Max_OxySensCrnt:        d.Messages[0].Signals[108],
		S01PID4F_Max_IntManiAbsPres:     d.Messages[0].Signals[109],
		S01PID50_Max_AirFlowMAF:         d.Messages[0].Signals[110],
		S01PID51_FuelType:               d.Messages[0].Signals[111],
		S01PID52_EthanolFuelPct:         d.Messages[0].Signals[112],
		S01PID53_AbsEvapSysVapPres:      d.Messages[0].Signals[113],
		S01PID54_EvapSysVapPres:         d.Messages[0].Signals[114],
		S01PID55_ShortSecOxyTrimBank1:   d.Messages[0].Signals[115],
		S01PID56_LongSecOxyTrimBank1:    d.Messages[0].Signals[116],
		S01PID55_ShortSecOxyTrimBank3:   d.Messages[0].Signals[117],
		S01PID56_LongSecOxyTrimBank3:    d.Messages[0].Signals[118],
		S01PID57_ShortSecOxyTrimBank2:   d.Messages[0].Signals[119],
		S01PID58_LongSecOxyTrimBank2:    d.Messages[0].Signals[120],
		S01PID59_FuelRailAbsPres:        d.Messages[0].Signals[121],
		S01PID5A_RelAccelPedalPos:       d.Messages[0].Signals[122],
		S01PID57_ShortSecOxyTrimBank4:   d.Messages[0].Signals[123],
		S01PID58_LongSecOxyTrimBank4:    d.Messages[0].Signals[124],
		S01PID5B_HybrBatPackRemLife:     d.Messages[0].Signals[125],
		S01PID5C_EngineOilTemp:          d.Messages[0].Signals[126],
		S01PID5D_FuelInjectionTiming:    d.Messages[0].Signals[127],
		S01PID5E_EngineFuelRate:         d.Messages[0].Signals[128],
		S01PID5F_EmissionReq:            d.Messages[0].Signals[129],
		S01PID60_PIDsSupported_61_80:    d.Messages[0].Signals[130],
		S01PID61_DemandEngTorqPct:       d.Messages[0].Signals[131],
		S01PID62_ActualEngTorqPct:       d.Messages[0].Signals[132],
		S01PID63_EngRefTorq:             d.Messages[0].Signals[133],
		S01PID64_EngPctTorq_Idle:        d.Messages[0].Signals[134],
		S01PID64_EngPctTorq_EP1:         d.Messages[0].Signals[135],
		S01PID64_EngPctTorq_EP2:         d.Messages[0].Signals[136],
		S01PID64_EngPctTorq_EP3:         d.Messages[0].Signals[137],
		S01PID65_AuxInputOutput:         d.Messages[0].Signals[138],
		S01PID66_MAFSensor:              d.Messages[0].Signals[139],
		S01PID64_EngPctTorq_EP4:         d.Messages[0].Signals[140],
		S01PID67_EngineCoolantTemp:      d.Messages[0].Signals[141],
		S01PID68_IntakeAirTempSens:      d.Messages[0].Signals[142],
		S01PID69_CmdEGR_EGRError:        d.Messages[0].Signals[143],
		S01PID6A_CmdDieselIntAir:        d.Messages[0].Signals[144],
		S01PID6B_ExhaustGasTemp:         d.Messages[0].Signals[145],
		S01PID6C_CmdThrottleActRel:      d.Messages[0].Signals[146],
		S01PID6D_FuelPresContrSys:       d.Messages[0].Signals[147],
		S01PID6E_InjPresContrSys:        d.Messages[0].Signals[148],
		S01PID6F_TurboComprPres:         d.Messages[0].Signals[149],
		S01PID70_BoostPresCntrl:         d.Messages[0].Signals[150],
		S01PID80_PIDsSupported_81_A0:    d.Messages[0].Signals[151],
		S01PID8E_EngFrictionPctTorq:     d.Messages[0].Signals[152],
		S01PIDA0_PIDsSupported_A1_C0:    d.Messages[0].Signals[153],
		S01PIDC0_PIDsSupported_C1_E0:    d.Messages[0].Signals[154],
	},
}

var d = (*descriptor.Database)(&descriptor.Database{
	SourceFile: (string)("dummy/obd2/bit29.dbc"),
	Version:    (string)(""),
	Messages: ([]*descriptor.Message)([]*descriptor.Message{
		(*descriptor.Message)(&descriptor.Message{
			Name:        (string)("OBD2"),
			ID:          (uint32)(417001749),
			IsExtended:  (bool)(true),
			Length:      (uint8)(8),
			SendType:    (descriptor.SendType)(0),
			Description: (string)("OBD2 0xDA00"),
			Signals: ([]*descriptor.Signal)([]*descriptor.Signal{
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("Length"),
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
					Name:             (string)("Service"),
					Start:            (uint8)(11),
					Length:           (uint8)(4),
					IsBigEndian:      (bool)(true),
					IsSigned:         (bool)(false),
					IsFloat:          (bool)(false),
					IsMultiplexer:    (bool)(true),
					IsMultiplexed:    (bool)(false),
					MultiplexerValue: (uint)(0),
					Offset:           (float64)(0),
					Scale:            (float64)(1),
					Min:              (float64)(0),
					Max:              (float64)(15),
					Unit:             (string)(""),
					Description:      (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)([]*descriptor.ValueDescription{
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(1),
							Description: (string)("Show current data "),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(2),
							Description: (string)("Show freeze frame data "),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(3),
							Description: (string)("Show stored DTCs "),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(4),
							Description: (string)("Clear DTCs and stored values"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(5),
							Description: (string)("Oxygen sensor monitoring "),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(6),
							Description: (string)("Other system monitoring "),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(7),
							Description: (string)("Show pending DTCs "),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(8),
							Description: (string)("Control on-board system "),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(9),
							Description: (string)("Request vehicle information "),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(10),
							Description: (string)("Permanent DTCs (Cleared DTCs) "),
						}),
					}),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("Response"),
					Start:             (uint8)(15),
					Length:            (uint8)(4),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(15),
					Unit:              (string)(""),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:             (string)("S01PID"),
					Start:            (uint8)(23),
					Length:           (uint8)(8),
					IsBigEndian:      (bool)(true),
					IsSigned:         (bool)(false),
					IsFloat:          (bool)(false),
					IsMultiplexer:    (bool)(true),
					IsMultiplexed:    (bool)(true),
					MultiplexerValue: (uint)(1),
					Offset:           (float64)(0),
					Scale:            (float64)(1),
					Min:              (float64)(0),
					Max:              (float64)(255),
					Unit:             (string)(""),
					Description:      (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)([]*descriptor.ValueDescription{
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(0),
							Description: (string)("S01PID00_PIDsSupported_01_20"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(1),
							Description: (string)("S01PID01_MonitorStatus"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(2),
							Description: (string)("S01PID02_FreezeDTC"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(3),
							Description: (string)("S01PID03_FuelSystemStatus"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(4),
							Description: (string)("S01PID04_CalcEngineLoad"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(5),
							Description: (string)("S01PID05_EngineCoolantTemp"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(6),
							Description: (string)("S01PID06_ShortFuelTrimBank1"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(7),
							Description: (string)("S01PID07_LongFuelTrimBank1"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(8),
							Description: (string)("S01PID08_ShortFuelTrimBank2"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(9),
							Description: (string)("S01PID09_LongFuelTrimBank2"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(10),
							Description: (string)("S01PID0A_FuelPressure"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(11),
							Description: (string)("S01PID0B_IntakeManiAbsPress"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(12),
							Description: (string)("S01PID0C_EngineRPM"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(13),
							Description: (string)("S01PID0D_VehicleSpeed"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(14),
							Description: (string)("S01PID0E_TimingAdvance"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(15),
							Description: (string)("S01PID0F_IntakeAirTemperature"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(16),
							Description: (string)("S01PID10_MAFAirFlowRate"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(17),
							Description: (string)("S01PID11_ThrottlePosition"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(18),
							Description: (string)("S01PID12_CmdSecAirStatus"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(20),
							Description: (string)("S01PID14_OxySensor1"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(21),
							Description: (string)("S01PID15_OxySensor2"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(22),
							Description: (string)("S01PID16_OxySensor3"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(23),
							Description: (string)("S01PID17_OxySensor4"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(24),
							Description: (string)("S01PID18_OxySensor5"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(25),
							Description: (string)("S01PID19_OxySensor6"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(26),
							Description: (string)("S01PID1A_OxySensor7"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(27),
							Description: (string)("S01PID1B_OxySensor8"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(28),
							Description: (string)("S01PID1C_OBDStandard"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(31),
							Description: (string)("S01PID1F_TimeSinceEngStart"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(32),
							Description: (string)("S01PID20_PIDsSupported_21_40"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(33),
							Description: (string)("S01PID21_DistanceMILOn"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(34),
							Description: (string)("S01PID22_FuelRailPres"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(35),
							Description: (string)("S01PID23_FuelRailGaug"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(36),
							Description: (string)("S01PID24_OxySensor1"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(37),
							Description: (string)("S01PID25_OxySensor2"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(38),
							Description: (string)("S01PID26_OxySensor3"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(39),
							Description: (string)("S01PID27_OxySensor4"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(40),
							Description: (string)("S01PID28_OxySensor5"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(41),
							Description: (string)("S01PID29_OxySensor6"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(42),
							Description: (string)("S01PID2A_OxySensor7"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(43),
							Description: (string)("S01PID2B_OxySensor8"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(44),
							Description: (string)("S01PID2C_CmdEGR"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(45),
							Description: (string)("S01PID2D_EGRError"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(46),
							Description: (string)("S01PID2E_CmdEvapPurge"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(47),
							Description: (string)("S01PID2F_FuelTankLevel"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(48),
							Description: (string)("S01PID30_WarmUpsSinceCodeClear"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(49),
							Description: (string)("S01PID31_DistanceSinceCodeClear"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(50),
							Description: (string)("S01PID32_EvapSysVaporPres"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(51),
							Description: (string)("S01PID33_AbsBaroPres"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(52),
							Description: (string)("S01PID34_OxySensor1"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(53),
							Description: (string)("S01PID35_OxySensor2"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(54),
							Description: (string)("S01PID36_OxySensor3"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(55),
							Description: (string)("S01PID37_OxySensor4"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(56),
							Description: (string)("S01PID38_OxySensor5"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(57),
							Description: (string)("S01PID39_OxySensor6"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(58),
							Description: (string)("S01PID3A_OxySensor7"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(59),
							Description: (string)("S01PID3B_OxySensor8"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(60),
							Description: (string)("S01PID3C_CatTempBank1Sens1"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(61),
							Description: (string)("S01PID3D_CatTempBank2Sens1"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(62),
							Description: (string)("S01PID3E_CatTempBank1Sens2"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(63),
							Description: (string)("S01PID3F_CatTempBank2Sens2"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(64),
							Description: (string)("S01PID40_PIDsSupported_41_60"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(65),
							Description: (string)("S01PID41_MonStatusDriveCycle"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(66),
							Description: (string)("S01PID42_ControlModuleVolt"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(67),
							Description: (string)("S01PID43_AbsLoadValue"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(68),
							Description: (string)("S01PID44_FuelAirCmdEquiv"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(69),
							Description: (string)("S01PID45_RelThrottlePos"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(70),
							Description: (string)("S01PID46_AmbientAirTemp"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(71),
							Description: (string)("S01PID47_AbsThrottlePosB"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(72),
							Description: (string)("S01PID48_AbsThrottlePosC"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(73),
							Description: (string)("S01PID49_AbsThrottlePosD"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(74),
							Description: (string)("S01PID4A_AbsThrottlePosE"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(75),
							Description: (string)("S01PID4B_AbsThrottlePosF"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(76),
							Description: (string)("S01PID4C_CmdThrottleAct"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(77),
							Description: (string)("S01PID4D_TimeRunMILOn"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(78),
							Description: (string)("S01PID4E_TimeSinceCodeClear"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(79),
							Description: (string)("S01PID4F_MaxMultiple"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(80),
							Description: (string)("S01PID50_Max_AirFlowMAF"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(81),
							Description: (string)("S01PID51_FuelType"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(82),
							Description: (string)("S01PID52_EthanolFuelPct"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(83),
							Description: (string)("S01PID53_AbsEvapSysVapPres"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(84),
							Description: (string)("S01PID54_EvapSysVapPres"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(85),
							Description: (string)("S01PID55_ShortSecOxyTrimBankX"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(86),
							Description: (string)("S01PID56_LongSecOxyTrimBankX"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(87),
							Description: (string)("S01PID57_ShortSecOxyTrimBankX"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(88),
							Description: (string)("S01PID58_LongSecOxyTrimBankX"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(89),
							Description: (string)("S01PID59_FuelRailAbsPres"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(90),
							Description: (string)("S01PID5A_RelAccelPedalPos"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(91),
							Description: (string)("S01PID5B_HybrBatPackRemLife"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(92),
							Description: (string)("S01PID5C_EngineOilTemp"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(93),
							Description: (string)("S01PID5D_FuelInjectionTiming"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(94),
							Description: (string)("S01PID5E_EngineFuelRate"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(95),
							Description: (string)("S01PID5F_EmissionReq"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(96),
							Description: (string)("S01PID60_PIDsSupported_61_80"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(97),
							Description: (string)("S01PID61_DemandEngTorqPct"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(98),
							Description: (string)("S01PID62_ActualEngTorqPct"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(99),
							Description: (string)("S01PID63_EngRefTorq"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(100),
							Description: (string)("S01PID64_EngPctTorq"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(101),
							Description: (string)("S01PID65_AuxInputOutput"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(102),
							Description: (string)("S01PID66_MAFSensor"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(103),
							Description: (string)("S01PID67_EngineCoolantTemp"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(104),
							Description: (string)("S01PID68_IntakeAirTempSens"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(105),
							Description: (string)("S01PID69_CmdEGR_EGRError"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(106),
							Description: (string)("S01PID6A_CmdDieselIntAir"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(107),
							Description: (string)("S01PID6B_ExhaustGasTemp"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(108),
							Description: (string)("S01PID6C_CmdThrottleActRel"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(109),
							Description: (string)("S01PID6D_FuelPresContrSys"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(110),
							Description: (string)("S01PID6E_InjPresContrSys"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(111),
							Description: (string)("S01PID6F_TurboComprPres"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(112),
							Description: (string)("S01PID70_BoostPresCntrl"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(128),
							Description: (string)("S01PID80_PIDsSupported_81_A0"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(142),
							Description: (string)("S01PID8E_EngFrictionPctTorq"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(160),
							Description: (string)("S01PIDA0_PIDsSupported_A1_C0"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(192),
							Description: (string)("S01PIDC0_PIDsSupported_C1_E0"),
						}),
					}),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:             (string)("S02PID"),
					Start:            (uint8)(23),
					Length:           (uint8)(8),
					IsBigEndian:      (bool)(true),
					IsSigned:         (bool)(false),
					IsFloat:          (bool)(false),
					IsMultiplexer:    (bool)(true),
					IsMultiplexed:    (bool)(true),
					MultiplexerValue: (uint)(2),
					Offset:           (float64)(0),
					Scale:            (float64)(1),
					Min:              (float64)(0),
					Max:              (float64)(255),
					Unit:             (string)(""),
					Description:      (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)([]*descriptor.ValueDescription{
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(2),
							Description: (string)("S02PID02_FreezeDTC"),
						}),
					}),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S01PID00_PIDsSupported_01_20"),
					Start:             (uint8)(31),
					Length:            (uint8)(32),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(4.294967295e+09),
					Unit:              (string)(""),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S01PID01_MonitorStatus"),
					Start:             (uint8)(31),
					Length:            (uint8)(32),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(1),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(4.294967295e+09),
					Unit:              (string)(""),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S02PID02_FreezeDTC"),
					Start:             (uint8)(31),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(2),
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
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S01PID02_FreezeDTC"),
					Start:             (uint8)(31),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(2),
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
				(*descriptor.Signal)(&descriptor.Signal{
					Name:             (string)("S01PID03_FuelSystemStatus"),
					Start:            (uint8)(31),
					Length:           (uint8)(16),
					IsBigEndian:      (bool)(true),
					IsSigned:         (bool)(false),
					IsFloat:          (bool)(false),
					IsMultiplexer:    (bool)(false),
					IsMultiplexed:    (bool)(true),
					MultiplexerValue: (uint)(3),
					Offset:           (float64)(0),
					Scale:            (float64)(1),
					Min:              (float64)(0),
					Max:              (float64)(65535),
					Unit:             (string)(""),
					Description:      (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)([]*descriptor.ValueDescription{
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(1),
							Description: (string)("Open loop (insuff. eng. temp.)"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(2),
							Description: (string)("Closed loop (oxy sens)"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(4),
							Description: (string)("Open loop (eng. load, fuel cut)"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(8),
							Description: (string)("Open loop (system failure)"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(16),
							Description: (string)("Closed loop (feedback issue)"),
						}),
					}),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S01PID04_CalcEngineLoad"),
					Start:             (uint8)(31),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(4),
					Offset:            (float64)(0),
					Scale:             (float64)(0.39216),
					Min:               (float64)(0),
					Max:               (float64)(100),
					Unit:              (string)("%"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S01PID05_EngineCoolantTemp"),
					Start:             (uint8)(31),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(5),
					Offset:            (float64)(-40),
					Scale:             (float64)(1),
					Min:               (float64)(-40),
					Max:               (float64)(215),
					Unit:              (string)("degC"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S01PID06_ShortFuelTrimBank1"),
					Start:             (uint8)(31),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(6),
					Offset:            (float64)(-100),
					Scale:             (float64)(0.78125),
					Min:               (float64)(-100),
					Max:               (float64)(99.21875),
					Unit:              (string)("%"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S01PID07_LongFuelTrimBank1"),
					Start:             (uint8)(31),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(7),
					Offset:            (float64)(-100),
					Scale:             (float64)(0.78125),
					Min:               (float64)(-100),
					Max:               (float64)(99.21875),
					Unit:              (string)("%"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S01PID08_ShortFuelTrimBank2"),
					Start:             (uint8)(31),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(8),
					Offset:            (float64)(-100),
					Scale:             (float64)(0.78125),
					Min:               (float64)(-100),
					Max:               (float64)(99.21875),
					Unit:              (string)("%"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S01PID09_LongFuelTrimBank2"),
					Start:             (uint8)(31),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(9),
					Offset:            (float64)(-100),
					Scale:             (float64)(0.78125),
					Min:               (float64)(-100),
					Max:               (float64)(99.21875),
					Unit:              (string)("%"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S01PID0A_FuelPressure"),
					Start:             (uint8)(31),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(10),
					Offset:            (float64)(0),
					Scale:             (float64)(3),
					Min:               (float64)(0),
					Max:               (float64)(765),
					Unit:              (string)("kPa"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S01PID0B_IntakeManiAbsPress"),
					Start:             (uint8)(31),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(11),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(255),
					Unit:              (string)("kPa"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S01PID0C_EngineRPM"),
					Start:             (uint8)(31),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(12),
					Offset:            (float64)(0),
					Scale:             (float64)(0.25),
					Min:               (float64)(0),
					Max:               (float64)(16383.75),
					Unit:              (string)("rpm"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S01PID0D_VehicleSpeed"),
					Start:             (uint8)(31),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(13),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(255),
					Unit:              (string)("km/h"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S01PID0E_TimingAdvance"),
					Start:             (uint8)(31),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(14),
					Offset:            (float64)(-64),
					Scale:             (float64)(0.5),
					Min:               (float64)(-64),
					Max:               (float64)(63.5),
					Unit:              (string)("deg before TDC"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S01PID0F_IntakeAirTemperature"),
					Start:             (uint8)(31),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(15),
					Offset:            (float64)(-40),
					Scale:             (float64)(1),
					Min:               (float64)(-40),
					Max:               (float64)(215),
					Unit:              (string)("degC"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S01PID10_MAFAirFlowRate"),
					Start:             (uint8)(31),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(16),
					Offset:            (float64)(0),
					Scale:             (float64)(0.01),
					Min:               (float64)(0),
					Max:               (float64)(655.35),
					Unit:              (string)("grams/sec"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S01PID11_ThrottlePosition"),
					Start:             (uint8)(31),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(17),
					Offset:            (float64)(0),
					Scale:             (float64)(0.39216),
					Min:               (float64)(0),
					Max:               (float64)(100),
					Unit:              (string)("%"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:             (string)("S01PID12_CmdSecAirStatus"),
					Start:            (uint8)(31),
					Length:           (uint8)(8),
					IsBigEndian:      (bool)(true),
					IsSigned:         (bool)(false),
					IsFloat:          (bool)(false),
					IsMultiplexer:    (bool)(false),
					IsMultiplexed:    (bool)(true),
					MultiplexerValue: (uint)(18),
					Offset:           (float64)(0),
					Scale:            (float64)(1),
					Min:              (float64)(0),
					Max:              (float64)(255),
					Unit:             (string)(""),
					Description:      (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)([]*descriptor.ValueDescription{
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(1),
							Description: (string)("Upstream"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(2),
							Description: (string)("Downstream catalytic conv"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(4),
							Description: (string)("From outside atmosphere/off"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(8),
							Description: (string)("Pump cmd on for diagn."),
						}),
					}),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S01PID14_OxySensor1_Volt"),
					Start:             (uint8)(31),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(20),
					Offset:            (float64)(0),
					Scale:             (float64)(0.005),
					Min:               (float64)(0),
					Max:               (float64)(1.275),
					Unit:              (string)("volts"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S01PID15_OxySensor2_Volt"),
					Start:             (uint8)(31),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(21),
					Offset:            (float64)(0),
					Scale:             (float64)(0.005),
					Min:               (float64)(0),
					Max:               (float64)(1.275),
					Unit:              (string)("volts"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S01PID16_OxySensor3_Volt"),
					Start:             (uint8)(31),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(22),
					Offset:            (float64)(0),
					Scale:             (float64)(0.005),
					Min:               (float64)(0),
					Max:               (float64)(1.275),
					Unit:              (string)("volts"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S01PID17_OxySensor4_Volt"),
					Start:             (uint8)(31),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(23),
					Offset:            (float64)(0),
					Scale:             (float64)(0.005),
					Min:               (float64)(0),
					Max:               (float64)(1.275),
					Unit:              (string)("volts"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S01PID18_OxySensor5_Volt"),
					Start:             (uint8)(31),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(24),
					Offset:            (float64)(0),
					Scale:             (float64)(0.005),
					Min:               (float64)(0),
					Max:               (float64)(1.275),
					Unit:              (string)("volts"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S01PID19_OxySensor6_Volt"),
					Start:             (uint8)(31),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(25),
					Offset:            (float64)(0),
					Scale:             (float64)(0.005),
					Min:               (float64)(0),
					Max:               (float64)(1.275),
					Unit:              (string)("volts"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S01PID14_OxySensor1_STFT"),
					Start:             (uint8)(39),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(20),
					Offset:            (float64)(-100),
					Scale:             (float64)(0.78125),
					Min:               (float64)(-100),
					Max:               (float64)(99.21875),
					Unit:              (string)("%"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S01PID15_OxySensor2_STFT"),
					Start:             (uint8)(39),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(21),
					Offset:            (float64)(-100),
					Scale:             (float64)(0.78125),
					Min:               (float64)(-100),
					Max:               (float64)(99.21875),
					Unit:              (string)("%"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S01PID16_OxySensor3_STFT"),
					Start:             (uint8)(39),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(22),
					Offset:            (float64)(-100),
					Scale:             (float64)(0.78125),
					Min:               (float64)(-100),
					Max:               (float64)(99.21875),
					Unit:              (string)("%"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S01PID17_OxySensor4_STFT"),
					Start:             (uint8)(39),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(23),
					Offset:            (float64)(-100),
					Scale:             (float64)(0.78125),
					Min:               (float64)(-100),
					Max:               (float64)(99.21875),
					Unit:              (string)("%"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S01PID18_OxySensor5_STFT"),
					Start:             (uint8)(39),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(24),
					Offset:            (float64)(-100),
					Scale:             (float64)(0.78125),
					Min:               (float64)(-100),
					Max:               (float64)(99.21875),
					Unit:              (string)("%"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S01PID19_OxySensor6_STFT"),
					Start:             (uint8)(39),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(25),
					Offset:            (float64)(-100),
					Scale:             (float64)(0.78125),
					Min:               (float64)(-100),
					Max:               (float64)(99.21875),
					Unit:              (string)("%"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S01PID1A_OxySensor7_Volt"),
					Start:             (uint8)(31),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(26),
					Offset:            (float64)(0),
					Scale:             (float64)(0.005),
					Min:               (float64)(0),
					Max:               (float64)(1.275),
					Unit:              (string)("volts"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S01PID1A_OxySensor7_STFT"),
					Start:             (uint8)(39),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(26),
					Offset:            (float64)(-100),
					Scale:             (float64)(0.78125),
					Min:               (float64)(-100),
					Max:               (float64)(99.21875),
					Unit:              (string)("%"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S01PID1B_OxySensor8_Volt"),
					Start:             (uint8)(31),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(27),
					Offset:            (float64)(0),
					Scale:             (float64)(0.005),
					Min:               (float64)(0),
					Max:               (float64)(1.275),
					Unit:              (string)("volts"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S01PID1B_OxySensor8_STFT"),
					Start:             (uint8)(39),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(27),
					Offset:            (float64)(-100),
					Scale:             (float64)(0.78125),
					Min:               (float64)(-100),
					Max:               (float64)(99.21875),
					Unit:              (string)("%"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:             (string)("S01PID1C_OBDStandard"),
					Start:            (uint8)(31),
					Length:           (uint8)(8),
					IsBigEndian:      (bool)(true),
					IsSigned:         (bool)(false),
					IsFloat:          (bool)(false),
					IsMultiplexer:    (bool)(false),
					IsMultiplexed:    (bool)(true),
					MultiplexerValue: (uint)(28),
					Offset:           (float64)(0),
					Scale:            (float64)(1),
					Min:              (float64)(0),
					Max:              (float64)(255),
					Unit:             (string)(""),
					Description:      (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)([]*descriptor.ValueDescription{
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(1),
							Description: (string)("OBD-II as defined by the CARB"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(2),
							Description: (string)("OBD as defined by the EPA"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(3),
							Description: (string)("OBD and OBD-II"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(4),
							Description: (string)("OBD-I"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(5),
							Description: (string)("Not OBD compliant"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(6),
							Description: (string)("EOBD (Europe)"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(7),
							Description: (string)("EOBD and OBD-II"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(8),
							Description: (string)("EOBD and OBD"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(9),
							Description: (string)("EOBD, OBD and OBD II"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(10),
							Description: (string)("JOBD (Japan)"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(11),
							Description: (string)("JOBD and OBD II"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(12),
							Description: (string)("JOBD and EOBD"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(13),
							Description: (string)("JOBD, EOBD, and OBD II"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(14),
							Description: (string)("Reserved"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(15),
							Description: (string)("Reserved"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(16),
							Description: (string)("Reserved"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(17),
							Description: (string)("Eng. Manu. Diag. (EMD)"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(18),
							Description: (string)("EMD Enhanced (EMD+)"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(19),
							Description: (string)("HD OBD-C"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(20),
							Description: (string)("HD OBD"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(21),
							Description: (string)("WWH OBD"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(22),
							Description: (string)("Reserved"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(23),
							Description: (string)("HD EOBD-I"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(24),
							Description: (string)("HD EOBD-I N"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(25),
							Description: (string)("HD EOBD-II"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(26),
							Description: (string)("HD EOBD-II N"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(27),
							Description: (string)("Reserved"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(28),
							Description: (string)("OBDBr-1"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(29),
							Description: (string)("OBDBr-2"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(30),
							Description: (string)("KOBD"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(31),
							Description: (string)("IOBD I"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(32),
							Description: (string)("IOBD II"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(33),
							Description: (string)("HD EOBD-IV"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(34),
							Description: (string)("Reserved"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(35),
							Description: (string)("Reserved"),
						}),
					}),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S01PID1F_TimeSinceEngStart"),
					Start:             (uint8)(31),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(31),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(65535),
					Unit:              (string)("seconds"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S01PID20_PIDsSupported_21_40"),
					Start:             (uint8)(31),
					Length:            (uint8)(32),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(32),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(4.294967295e+09),
					Unit:              (string)(""),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S01PID21_DistanceMILOn"),
					Start:             (uint8)(31),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(33),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(65535),
					Unit:              (string)("km"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S01PID22_FuelRailPres"),
					Start:             (uint8)(31),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(34),
					Offset:            (float64)(0),
					Scale:             (float64)(0.079),
					Min:               (float64)(0),
					Max:               (float64)(5177.265),
					Unit:              (string)("kPa"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S01PID23_FuelRailGaug"),
					Start:             (uint8)(31),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(35),
					Offset:            (float64)(0),
					Scale:             (float64)(10),
					Min:               (float64)(0),
					Max:               (float64)(655350),
					Unit:              (string)("kPa"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S01PID24_OxySensor1_FAER"),
					Start:             (uint8)(31),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(36),
					Offset:            (float64)(0),
					Scale:             (float64)(3.0517578125e-05),
					Min:               (float64)(0),
					Max:               (float64)(2),
					Unit:              (string)("ratio"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S01PID24_OxySensor1_Volt"),
					Start:             (uint8)(47),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(36),
					Offset:            (float64)(0),
					Scale:             (float64)(0.0001220703125),
					Min:               (float64)(0),
					Max:               (float64)(2),
					Unit:              (string)("volts"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S01PID25_OxySensor2_FAER"),
					Start:             (uint8)(31),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(37),
					Offset:            (float64)(0),
					Scale:             (float64)(3.0517578125e-05),
					Min:               (float64)(0),
					Max:               (float64)(2),
					Unit:              (string)("ratio"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S01PID25_OxySensor2_Volt"),
					Start:             (uint8)(47),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(37),
					Offset:            (float64)(0),
					Scale:             (float64)(0.0001220703125),
					Min:               (float64)(0),
					Max:               (float64)(8),
					Unit:              (string)("volts"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S01PID26_OxySensor3_FAER"),
					Start:             (uint8)(31),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(38),
					Offset:            (float64)(0),
					Scale:             (float64)(3.0517578125e-05),
					Min:               (float64)(0),
					Max:               (float64)(2),
					Unit:              (string)("ratio"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S01PID26_OxySensor3_Volt"),
					Start:             (uint8)(47),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(38),
					Offset:            (float64)(0),
					Scale:             (float64)(0.0001220703125),
					Min:               (float64)(0),
					Max:               (float64)(8),
					Unit:              (string)("volts"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S01PID27_OxySensor4_FAER"),
					Start:             (uint8)(31),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(39),
					Offset:            (float64)(0),
					Scale:             (float64)(3.0517578125e-05),
					Min:               (float64)(0),
					Max:               (float64)(2),
					Unit:              (string)("ratio"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S01PID28_OxySensor5_FAER"),
					Start:             (uint8)(31),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(40),
					Offset:            (float64)(0),
					Scale:             (float64)(3.0517578125e-05),
					Min:               (float64)(0),
					Max:               (float64)(2),
					Unit:              (string)("ratio"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S01PID29_OxySensor6_FAER"),
					Start:             (uint8)(31),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(41),
					Offset:            (float64)(0),
					Scale:             (float64)(3.0517578125e-05),
					Min:               (float64)(0),
					Max:               (float64)(2),
					Unit:              (string)("ratio"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S01PID27_OxySensor4_Volt"),
					Start:             (uint8)(47),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(39),
					Offset:            (float64)(0),
					Scale:             (float64)(0.0001220703125),
					Min:               (float64)(0),
					Max:               (float64)(8),
					Unit:              (string)("volts"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S01PID28_OxySensor5_Volt"),
					Start:             (uint8)(47),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(40),
					Offset:            (float64)(0),
					Scale:             (float64)(0.0001220703125),
					Min:               (float64)(0),
					Max:               (float64)(8),
					Unit:              (string)("volts"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S01PID29_OxySensor6_Volt"),
					Start:             (uint8)(47),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(41),
					Offset:            (float64)(0),
					Scale:             (float64)(0.0001220703125),
					Min:               (float64)(0),
					Max:               (float64)(8),
					Unit:              (string)("volts"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S01PID2A_OxySensor7_FAER"),
					Start:             (uint8)(31),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(42),
					Offset:            (float64)(0),
					Scale:             (float64)(3.0517578125e-05),
					Min:               (float64)(0),
					Max:               (float64)(2),
					Unit:              (string)("ratio"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S01PID2A_OxySensor7_Volt"),
					Start:             (uint8)(47),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(42),
					Offset:            (float64)(0),
					Scale:             (float64)(0.0001220703125),
					Min:               (float64)(0),
					Max:               (float64)(8),
					Unit:              (string)("volts"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S01PID2B_OxySensor8_FAER"),
					Start:             (uint8)(31),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(43),
					Offset:            (float64)(0),
					Scale:             (float64)(3.0517578125e-05),
					Min:               (float64)(0),
					Max:               (float64)(2),
					Unit:              (string)("ratio"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S01PID2B_OxySensor8_Volt"),
					Start:             (uint8)(47),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(43),
					Offset:            (float64)(0),
					Scale:             (float64)(0.0001220703125),
					Min:               (float64)(0),
					Max:               (float64)(8),
					Unit:              (string)("volts"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S01PID2C_CmdEGR"),
					Start:             (uint8)(31),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(44),
					Offset:            (float64)(0),
					Scale:             (float64)(0.392156862745098),
					Min:               (float64)(0),
					Max:               (float64)(100),
					Unit:              (string)("%"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S01PID2D_EGRError"),
					Start:             (uint8)(31),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(45),
					Offset:            (float64)(-100),
					Scale:             (float64)(0.78125),
					Min:               (float64)(-100),
					Max:               (float64)(99.21875),
					Unit:              (string)("%"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S01PID2E_CmdEvapPurge"),
					Start:             (uint8)(31),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(46),
					Offset:            (float64)(0),
					Scale:             (float64)(0.392156862745098),
					Min:               (float64)(0),
					Max:               (float64)(100),
					Unit:              (string)("%"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S01PID2F_FuelTankLevel"),
					Start:             (uint8)(31),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(47),
					Offset:            (float64)(0),
					Scale:             (float64)(0.392156862745098),
					Min:               (float64)(0),
					Max:               (float64)(100),
					Unit:              (string)("%"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S01PID30_WarmUpsSinceCodeClear"),
					Start:             (uint8)(31),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(48),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(255),
					Unit:              (string)("count"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S01PID31_DistanceSinceCodeClear"),
					Start:             (uint8)(31),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(49),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(65535),
					Unit:              (string)("km"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S01PID32_EvapSysVaporPres"),
					Start:             (uint8)(31),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(50),
					Offset:            (float64)(0),
					Scale:             (float64)(0.25),
					Min:               (float64)(-8192),
					Max:               (float64)(8191.75),
					Unit:              (string)("Pa"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S01PID33_AbsBaroPres"),
					Start:             (uint8)(31),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(51),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(255),
					Unit:              (string)("kPa"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S01PID34_OxySensor1_FAER"),
					Start:             (uint8)(31),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(52),
					Offset:            (float64)(0),
					Scale:             (float64)(3.0517578125e-05),
					Min:               (float64)(0),
					Max:               (float64)(2),
					Unit:              (string)("ratio"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S01PID34_OxySensor1_Crnt"),
					Start:             (uint8)(47),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(52),
					Offset:            (float64)(-128),
					Scale:             (float64)(0.00390625),
					Min:               (float64)(-128),
					Max:               (float64)(128),
					Unit:              (string)("mA"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S01PID35_OxySensor2_FAER"),
					Start:             (uint8)(31),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(53),
					Offset:            (float64)(0),
					Scale:             (float64)(3.0517578125e-05),
					Min:               (float64)(0),
					Max:               (float64)(2),
					Unit:              (string)("ratio"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S01PID35_OxySensor2_Crnt"),
					Start:             (uint8)(47),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(53),
					Offset:            (float64)(-128),
					Scale:             (float64)(0.00390625),
					Min:               (float64)(-128),
					Max:               (float64)(128),
					Unit:              (string)("mA"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S01PID36_OxySensor3_FAER"),
					Start:             (uint8)(31),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(54),
					Offset:            (float64)(0),
					Scale:             (float64)(3.0517578125e-05),
					Min:               (float64)(0),
					Max:               (float64)(2),
					Unit:              (string)("ratio"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S01PID36_OxySensor3_Crnt"),
					Start:             (uint8)(47),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(54),
					Offset:            (float64)(-128),
					Scale:             (float64)(0.00390625),
					Min:               (float64)(-128),
					Max:               (float64)(128),
					Unit:              (string)("mA"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S01PID37_OxySensor4_FAER"),
					Start:             (uint8)(31),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(55),
					Offset:            (float64)(0),
					Scale:             (float64)(3.0517578125e-05),
					Min:               (float64)(0),
					Max:               (float64)(2),
					Unit:              (string)("ratio"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S01PID38_OxySensor5_FAER"),
					Start:             (uint8)(31),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(56),
					Offset:            (float64)(0),
					Scale:             (float64)(3.0517578125e-05),
					Min:               (float64)(0),
					Max:               (float64)(2),
					Unit:              (string)("ratio"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S01PID39_OxySensor6_FAER"),
					Start:             (uint8)(31),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(57),
					Offset:            (float64)(0),
					Scale:             (float64)(3.0517578125e-05),
					Min:               (float64)(0),
					Max:               (float64)(2),
					Unit:              (string)("ratio"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S01PID37_OxySensor4_Crnt"),
					Start:             (uint8)(47),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(55),
					Offset:            (float64)(-128),
					Scale:             (float64)(0.00390625),
					Min:               (float64)(-128),
					Max:               (float64)(128),
					Unit:              (string)("mA"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S01PID38_OxySensor5_Crnt"),
					Start:             (uint8)(47),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(56),
					Offset:            (float64)(-128),
					Scale:             (float64)(0.00390625),
					Min:               (float64)(-128),
					Max:               (float64)(128),
					Unit:              (string)("mA"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S01PID39_OxySensor6_Crnt"),
					Start:             (uint8)(47),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(57),
					Offset:            (float64)(-128),
					Scale:             (float64)(0.00390625),
					Min:               (float64)(-128),
					Max:               (float64)(128),
					Unit:              (string)("mA"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S01PID3A_OxySensor7_FAER"),
					Start:             (uint8)(31),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(58),
					Offset:            (float64)(0),
					Scale:             (float64)(3.0517578125e-05),
					Min:               (float64)(0),
					Max:               (float64)(2),
					Unit:              (string)("ratio"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S01PID3B_OxySensor8_FAER"),
					Start:             (uint8)(31),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(59),
					Offset:            (float64)(0),
					Scale:             (float64)(3.0517578125e-05),
					Min:               (float64)(0),
					Max:               (float64)(2),
					Unit:              (string)("ratio"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S01PID3C_CatTempBank1Sens1"),
					Start:             (uint8)(31),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(60),
					Offset:            (float64)(-40),
					Scale:             (float64)(0.1),
					Min:               (float64)(-40),
					Max:               (float64)(6513.5),
					Unit:              (string)("degC"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S01PID3D_CatTempBank2Sens1"),
					Start:             (uint8)(31),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(61),
					Offset:            (float64)(-40),
					Scale:             (float64)(0.1),
					Min:               (float64)(-40),
					Max:               (float64)(6513.5),
					Unit:              (string)("degC"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S01PID3A_OxySensor7_Crnt"),
					Start:             (uint8)(47),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(58),
					Offset:            (float64)(-128),
					Scale:             (float64)(0.00390625),
					Min:               (float64)(-128),
					Max:               (float64)(128),
					Unit:              (string)("mA"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S01PID3B_OxySensor8_Crnt"),
					Start:             (uint8)(47),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(59),
					Offset:            (float64)(-128),
					Scale:             (float64)(0.00390625),
					Min:               (float64)(-128),
					Max:               (float64)(128),
					Unit:              (string)("mA"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S01PID3E_CatTempBank1Sens2"),
					Start:             (uint8)(31),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(62),
					Offset:            (float64)(-40),
					Scale:             (float64)(0.1),
					Min:               (float64)(-40),
					Max:               (float64)(6513.5),
					Unit:              (string)("degC"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S01PID3F_CatTempBank2Sens2"),
					Start:             (uint8)(31),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(63),
					Offset:            (float64)(-40),
					Scale:             (float64)(0.1),
					Min:               (float64)(-40),
					Max:               (float64)(6513.5),
					Unit:              (string)("degC"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S01PID40_PIDsSupported_41_60"),
					Start:             (uint8)(31),
					Length:            (uint8)(32),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(64),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(4.294967295e+09),
					Unit:              (string)(""),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S01PID41_MonStatusDriveCycle"),
					Start:             (uint8)(31),
					Length:            (uint8)(32),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(65),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(4.294967295e+09),
					Unit:              (string)(""),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S01PID42_ControlModuleVolt"),
					Start:             (uint8)(31),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(66),
					Offset:            (float64)(0),
					Scale:             (float64)(0.001),
					Min:               (float64)(0),
					Max:               (float64)(65.535),
					Unit:              (string)("V"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S01PID43_AbsLoadValue"),
					Start:             (uint8)(31),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(67),
					Offset:            (float64)(0),
					Scale:             (float64)(0.392156862745098),
					Min:               (float64)(0),
					Max:               (float64)(25700),
					Unit:              (string)("%"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S01PID44_FuelAirCmdEquiv"),
					Start:             (uint8)(31),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(68),
					Offset:            (float64)(0),
					Scale:             (float64)(3.0517578125e-05),
					Min:               (float64)(0),
					Max:               (float64)(2),
					Unit:              (string)("ratio"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S01PID45_RelThrottlePos"),
					Start:             (uint8)(31),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(69),
					Offset:            (float64)(0),
					Scale:             (float64)(0.392156862745098),
					Min:               (float64)(0),
					Max:               (float64)(100),
					Unit:              (string)("%"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S01PID46_AmbientAirTemp"),
					Start:             (uint8)(31),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(70),
					Offset:            (float64)(-40),
					Scale:             (float64)(1),
					Min:               (float64)(-40),
					Max:               (float64)(215),
					Unit:              (string)("degC"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S01PID47_AbsThrottlePosB"),
					Start:             (uint8)(31),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(71),
					Offset:            (float64)(0),
					Scale:             (float64)(0.392156862745098),
					Min:               (float64)(0),
					Max:               (float64)(100),
					Unit:              (string)("%"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S01PID48_AbsThrottlePosC"),
					Start:             (uint8)(31),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(72),
					Offset:            (float64)(0),
					Scale:             (float64)(0.392156862745098),
					Min:               (float64)(0),
					Max:               (float64)(100),
					Unit:              (string)("%"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S01PID49_AbsThrottlePosD"),
					Start:             (uint8)(31),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(73),
					Offset:            (float64)(0),
					Scale:             (float64)(0.392156862745098),
					Min:               (float64)(0),
					Max:               (float64)(100),
					Unit:              (string)("%"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S01PID4A_AbsThrottlePosE"),
					Start:             (uint8)(31),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(74),
					Offset:            (float64)(0),
					Scale:             (float64)(0.392156862745098),
					Min:               (float64)(0),
					Max:               (float64)(100),
					Unit:              (string)("%"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S01PID4B_AbsThrottlePosF"),
					Start:             (uint8)(31),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(75),
					Offset:            (float64)(0),
					Scale:             (float64)(0.392156862745098),
					Min:               (float64)(0),
					Max:               (float64)(100),
					Unit:              (string)("%"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S01PID4C_CmdThrottleAct"),
					Start:             (uint8)(31),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(76),
					Offset:            (float64)(0),
					Scale:             (float64)(0.392156862745098),
					Min:               (float64)(0),
					Max:               (float64)(100),
					Unit:              (string)("%"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S01PID4D_TimeRunMILOn"),
					Start:             (uint8)(31),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(77),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(65535),
					Unit:              (string)("minutes"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S01PID4E_TimeSinceCodeClear"),
					Start:             (uint8)(31),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(78),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(65535),
					Unit:              (string)("minutes"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S01PID4F_Max_FAER"),
					Start:             (uint8)(31),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(79),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(255),
					Unit:              (string)("ratio"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S01PID4F_Max_OxySensVol"),
					Start:             (uint8)(39),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(79),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(255),
					Unit:              (string)("V"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S01PID4F_Max_OxySensCrnt"),
					Start:             (uint8)(47),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(79),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(255),
					Unit:              (string)("mA"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S01PID4F_Max_IntManiAbsPres"),
					Start:             (uint8)(55),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(79),
					Offset:            (float64)(0),
					Scale:             (float64)(10),
					Min:               (float64)(0),
					Max:               (float64)(2550),
					Unit:              (string)("kPa"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S01PID50_Max_AirFlowMAF"),
					Start:             (uint8)(31),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(80),
					Offset:            (float64)(0),
					Scale:             (float64)(10),
					Min:               (float64)(0),
					Max:               (float64)(2550),
					Unit:              (string)("g/s"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:             (string)("S01PID51_FuelType"),
					Start:            (uint8)(31),
					Length:           (uint8)(8),
					IsBigEndian:      (bool)(true),
					IsSigned:         (bool)(false),
					IsFloat:          (bool)(false),
					IsMultiplexer:    (bool)(false),
					IsMultiplexed:    (bool)(true),
					MultiplexerValue: (uint)(81),
					Offset:           (float64)(0),
					Scale:            (float64)(1),
					Min:              (float64)(0),
					Max:              (float64)(255),
					Unit:             (string)(""),
					Description:      (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)([]*descriptor.ValueDescription{
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(0),
							Description: (string)("Not available"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(1),
							Description: (string)("Gasoline"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(2),
							Description: (string)("Methanol"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(3),
							Description: (string)("Ethanol"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(4),
							Description: (string)("Diesel"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(5),
							Description: (string)("LPG"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(6),
							Description: (string)("CNG"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(7),
							Description: (string)("Propane"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(8),
							Description: (string)("Electric"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(9),
							Description: (string)("Bifuel running Gasoline"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(10),
							Description: (string)("Bifuel running Methanol"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(11),
							Description: (string)("Bifuel running Ethanol"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(12),
							Description: (string)("Bifuel running LPG"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(13),
							Description: (string)("Bifuel running CNG"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(14),
							Description: (string)("Bifuel running Propane"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(15),
							Description: (string)("Bifuel running Electricity"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(16),
							Description: (string)("Bifuel electric/comb. eng."),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(17),
							Description: (string)("Hybrid gasoline"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(18),
							Description: (string)("Hybrid Ethanol"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(19),
							Description: (string)("Hybrid Diesel"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(20),
							Description: (string)("Hybrid Electric"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(21),
							Description: (string)("Hybrid running electric/comb."),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(22),
							Description: (string)("Hybrid Regenerative"),
						}),
						(*descriptor.ValueDescription)(&descriptor.ValueDescription{
							Value:       (int64)(23),
							Description: (string)("Bifuel running diesel"),
						}),
					}),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S01PID52_EthanolFuelPct"),
					Start:             (uint8)(31),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(82),
					Offset:            (float64)(0),
					Scale:             (float64)(0.392156862745098),
					Min:               (float64)(0),
					Max:               (float64)(100),
					Unit:              (string)("%"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S01PID53_AbsEvapSysVapPres"),
					Start:             (uint8)(31),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(83),
					Offset:            (float64)(0),
					Scale:             (float64)(0.005),
					Min:               (float64)(0),
					Max:               (float64)(327.675),
					Unit:              (string)("kPa"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S01PID54_EvapSysVapPres"),
					Start:             (uint8)(31),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(84),
					Offset:            (float64)(-32767),
					Scale:             (float64)(1),
					Min:               (float64)(-32767),
					Max:               (float64)(32768),
					Unit:              (string)("Pa"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S01PID55_ShortSecOxyTrimBank1"),
					Start:             (uint8)(31),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(85),
					Offset:            (float64)(-100),
					Scale:             (float64)(0.78125),
					Min:               (float64)(-100),
					Max:               (float64)(99.21875),
					Unit:              (string)("%"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S01PID56_LongSecOxyTrimBank1"),
					Start:             (uint8)(31),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(86),
					Offset:            (float64)(-100),
					Scale:             (float64)(0.78125),
					Min:               (float64)(-100),
					Max:               (float64)(99.21875),
					Unit:              (string)("%"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S01PID55_ShortSecOxyTrimBank3"),
					Start:             (uint8)(39),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(85),
					Offset:            (float64)(-100),
					Scale:             (float64)(0.78125),
					Min:               (float64)(-100),
					Max:               (float64)(99.21875),
					Unit:              (string)("%"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S01PID56_LongSecOxyTrimBank3"),
					Start:             (uint8)(39),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(86),
					Offset:            (float64)(-100),
					Scale:             (float64)(0.78125),
					Min:               (float64)(-100),
					Max:               (float64)(99.21875),
					Unit:              (string)("%"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S01PID57_ShortSecOxyTrimBank2"),
					Start:             (uint8)(31),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(87),
					Offset:            (float64)(-100),
					Scale:             (float64)(0.78125),
					Min:               (float64)(-100),
					Max:               (float64)(99.21875),
					Unit:              (string)("%"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S01PID58_LongSecOxyTrimBank2"),
					Start:             (uint8)(31),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(88),
					Offset:            (float64)(-100),
					Scale:             (float64)(0.78125),
					Min:               (float64)(-100),
					Max:               (float64)(99.21875),
					Unit:              (string)("%"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S01PID59_FuelRailAbsPres"),
					Start:             (uint8)(31),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(89),
					Offset:            (float64)(0),
					Scale:             (float64)(10),
					Min:               (float64)(0),
					Max:               (float64)(655350),
					Unit:              (string)("kPa"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S01PID5A_RelAccelPedalPos"),
					Start:             (uint8)(31),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(90),
					Offset:            (float64)(0),
					Scale:             (float64)(0.392156862745098),
					Min:               (float64)(0),
					Max:               (float64)(100),
					Unit:              (string)("%"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S01PID57_ShortSecOxyTrimBank4"),
					Start:             (uint8)(39),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(87),
					Offset:            (float64)(-100),
					Scale:             (float64)(0.78125),
					Min:               (float64)(-100),
					Max:               (float64)(99.21875),
					Unit:              (string)("%"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S01PID58_LongSecOxyTrimBank4"),
					Start:             (uint8)(39),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(88),
					Offset:            (float64)(-100),
					Scale:             (float64)(0.78125),
					Min:               (float64)(-100),
					Max:               (float64)(99.21875),
					Unit:              (string)("%"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S01PID5B_HybrBatPackRemLife"),
					Start:             (uint8)(31),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(91),
					Offset:            (float64)(0),
					Scale:             (float64)(0.392156862745098),
					Min:               (float64)(0),
					Max:               (float64)(100),
					Unit:              (string)("%"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S01PID5C_EngineOilTemp"),
					Start:             (uint8)(31),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(92),
					Offset:            (float64)(-40),
					Scale:             (float64)(1),
					Min:               (float64)(-40),
					Max:               (float64)(215),
					Unit:              (string)("degC"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S01PID5D_FuelInjectionTiming"),
					Start:             (uint8)(31),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(93),
					Offset:            (float64)(-210),
					Scale:             (float64)(0.0078125),
					Min:               (float64)(-210),
					Max:               (float64)(301.9921875),
					Unit:              (string)("deg"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S01PID5E_EngineFuelRate"),
					Start:             (uint8)(31),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(94),
					Offset:            (float64)(0),
					Scale:             (float64)(0.05),
					Min:               (float64)(0),
					Max:               (float64)(3276.75),
					Unit:              (string)("L/h"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S01PID5F_EmissionReq"),
					Start:             (uint8)(31),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(95),
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
					Name:              (string)("S01PID60_PIDsSupported_61_80"),
					Start:             (uint8)(31),
					Length:            (uint8)(32),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(96),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(4.294967295e+09),
					Unit:              (string)(""),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S01PID61_DemandEngTorqPct"),
					Start:             (uint8)(31),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(97),
					Offset:            (float64)(-125),
					Scale:             (float64)(1),
					Min:               (float64)(-125),
					Max:               (float64)(130),
					Unit:              (string)("%"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S01PID62_ActualEngTorqPct"),
					Start:             (uint8)(31),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(98),
					Offset:            (float64)(-125),
					Scale:             (float64)(1),
					Min:               (float64)(-125),
					Max:               (float64)(130),
					Unit:              (string)("%"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S01PID63_EngRefTorq"),
					Start:             (uint8)(31),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(99),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(65535),
					Unit:              (string)("Nm"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S01PID64_EngPctTorq_Idle"),
					Start:             (uint8)(31),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(100),
					Offset:            (float64)(-125),
					Scale:             (float64)(1),
					Min:               (float64)(-125),
					Max:               (float64)(130),
					Unit:              (string)("%"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S01PID64_EngPctTorq_EP1"),
					Start:             (uint8)(39),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(100),
					Offset:            (float64)(-125),
					Scale:             (float64)(1),
					Min:               (float64)(-125),
					Max:               (float64)(130),
					Unit:              (string)("%"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S01PID64_EngPctTorq_EP2"),
					Start:             (uint8)(47),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(100),
					Offset:            (float64)(-125),
					Scale:             (float64)(1),
					Min:               (float64)(-125),
					Max:               (float64)(130),
					Unit:              (string)("%"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S01PID64_EngPctTorq_EP3"),
					Start:             (uint8)(55),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(100),
					Offset:            (float64)(-125),
					Scale:             (float64)(1),
					Min:               (float64)(-125),
					Max:               (float64)(130),
					Unit:              (string)("%"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S01PID65_AuxInputOutput"),
					Start:             (uint8)(31),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(101),
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
					Name:              (string)("S01PID66_MAFSensor"),
					Start:             (uint8)(31),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(102),
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
					Name:              (string)("S01PID64_EngPctTorq_EP4"),
					Start:             (uint8)(63),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(100),
					Offset:            (float64)(-125),
					Scale:             (float64)(1),
					Min:               (float64)(-125),
					Max:               (float64)(130),
					Unit:              (string)("%"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S01PID67_EngineCoolantTemp"),
					Start:             (uint8)(31),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(103),
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
					Name:              (string)("S01PID68_IntakeAirTempSens"),
					Start:             (uint8)(31),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(104),
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
					Name:              (string)("S01PID69_CmdEGR_EGRError"),
					Start:             (uint8)(31),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(105),
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
					Name:              (string)("S01PID6A_CmdDieselIntAir"),
					Start:             (uint8)(31),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(106),
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
					Name:              (string)("S01PID6B_ExhaustGasTemp"),
					Start:             (uint8)(31),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(107),
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
					Name:              (string)("S01PID6C_CmdThrottleActRel"),
					Start:             (uint8)(31),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(108),
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
					Name:              (string)("S01PID6D_FuelPresContrSys"),
					Start:             (uint8)(31),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(109),
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
					Name:              (string)("S01PID6E_InjPresContrSys"),
					Start:             (uint8)(31),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(110),
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
					Name:              (string)("S01PID6F_TurboComprPres"),
					Start:             (uint8)(31),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(111),
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
					Name:              (string)("S01PID70_BoostPresCntrl"),
					Start:             (uint8)(31),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(112),
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
					Name:              (string)("S01PID80_PIDsSupported_81_A0"),
					Start:             (uint8)(31),
					Length:            (uint8)(32),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(128),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(4.294967295e+09),
					Unit:              (string)(""),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S01PID8E_EngFrictionPctTorq"),
					Start:             (uint8)(31),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(142),
					Offset:            (float64)(-125),
					Scale:             (float64)(1),
					Min:               (float64)(-125),
					Max:               (float64)(130),
					Unit:              (string)("%"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S01PIDA0_PIDsSupported_A1_C0"),
					Start:             (uint8)(31),
					Length:            (uint8)(32),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(160),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(4.294967295e+09),
					Unit:              (string)(""),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("S01PIDC0_PIDsSupported_C1_E0"),
					Start:             (uint8)(31),
					Length:            (uint8)(32),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsFloat:           (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(true),
					MultiplexerValue:  (uint)(192),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(4.294967295e+09),
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
	Nodes: ([]*descriptor.Node)(nil),
})
