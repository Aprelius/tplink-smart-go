package tplink

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"strings"
	"time"

	"github.com/Aralocke/tplink-smart-go/v1/pkg/utils"
	"go.uber.org/zap"

	"github.com/Aralocke/tplink-smart-go/v1/pkg/devices"
)

var ErrProtocolOperationFailed = errors.New("operation on device returned an error")
var ErrUnsupportedFeature = errors.New("feature is not supported by the device")

type Feature int

const (
	FeatureElecMeter Feature = iota
)

var featureMap = map[Feature]string{
	FeatureElecMeter: "ENE",
}

type DeviceManager struct {
	dvManager *devices.DeviceManager
	devices   []*devices.Device
	logger    *zap.Logger
}

func NewDeviceManager(dm *devices.DeviceManager) *DeviceManager {
	dvManager := new(DeviceManager)
	dvManager.dvManager = dm
	dvManager.logger = dm.Logger()
	return dvManager
}

func (m *DeviceManager) Devices() []*devices.Device {
	return m.devices
}

func (m *DeviceManager) ElectricityMeter(d *devices.Device) (*EMeter, error) {
	if !m.Supports(d, FeatureElecMeter) {
		return nil, ErrUnsupportedFeature
	}

	return newElectricityMeter(d, m), nil
}

func (m *DeviceManager) LoadDevice(cfg *devices.DeviceConfig) (*devices.Device, error) {
	device := devices.NewDevice(cfg)

	info, err := m.SystemInfo(device)
	if err != nil {
		return nil, err
	}

	device = devices.NewDevice(cfg,
		devices.WithDeviceId(info.DeviceId),
		devices.WithDeviceName(info.DeviceName),
		devices.WithFeatures(strings.Split(info.Features, ":")),
		devices.WithFirmwareId(info.FirmwareId),
		devices.WithHardwareId(info.HardwareId),
		devices.WithHardwareVersion(info.HardwareVersion),
		devices.WithManufacturerId(info.ManufacturerId),
		devices.WithModelVersion(info.Model),
		devices.WithSoftwareVersion(info.SoftwareVersion),
	)

	m.devices = append(m.devices, device)
	return device, nil
}

func (m *DeviceManager) LoadDevices(devices []devices.DeviceConfig) error {
	for _, device := range devices {
		if _, err := m.LoadDevice(&device); err != nil {
			return err
		}
	}
	return nil
}

func (m *DeviceManager) Logger() *zap.Logger {
	return m.logger
}

func (m *DeviceManager) Marshal(d devices.Addressable, in interface{}) ([]byte, error) {
	var err error
	var s, res []byte

	s, err = json.Marshal(in)
	if err != nil {
		return []byte{}, err
	}

	m.Logger().Debug("marshal message to device",
		zap.String("device",
			fmt.Sprintf("%s:%d", d.Address(), d.Port())),
		zap.String("message", string(s)))

	sender := devices.NewSyncSender(d,
		devices.WithEncoding(Decrypt, Encrypt),
		devices.WithTimeout(30*time.Second))

	for i := 0; i < 5; i++ {
		res, err = sender.Send(s)
		if e, ok := err.(net.Error); ok {
			if i == 5 || (!e.Timeout() && !e.Temporary()) {
				return []byte{}, nil
			}

			m.Logger().Info("retrying device message",
				zap.Int("retry", i),
				zap.String("address",
					fmt.Sprintf("%s:%d", d.Address(), d.Port())),
				zap.Error(err))
			continue
		}
		break
	}

	m.Logger().Debug("unmarshal message from device",
		zap.String("device",
			fmt.Sprintf("%s:%d", d.Address(), d.Port())),
		zap.String("message", string(res)))

	return res, nil
}

func (m *DeviceManager) Off(d *devices.Device) error {
	return m.SetRelayState(d, false)
}

func (m *DeviceManager) On(d *devices.Device) error {
	return m.SetRelayState(d, true)
}

func (m *DeviceManager) Reboot(d *devices.Device, delay int) error {
	var r SystemReboot
	r.SetDelay(delay)

	res, err := m.Marshal(d, r)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(res, &r); err != nil {
		return err
	}

	if r.ErrorCode() != 0 {
		return ErrProtocolOperationFailed
	}

	return nil
}

func (m *DeviceManager) Reset(d *devices.Device, delay int) error {
	var r SystemReset
	r.SetDelay(delay)

	res, err := m.Marshal(d, r)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(res, &r); err != nil {
		return err
	}

	if r.ErrorCode() != 0 {
		return ErrProtocolOperationFailed
	}

	return nil
}

func (m *DeviceManager) SetAlias(d *devices.Device, alias string) error {
	var a SystemAlias
	a.SetAlias(alias)

	res, err := m.Marshal(d, a)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(res, &a); err != nil {
		return err
	}

	if a.ErrorCode() != 0 {
		return ErrProtocolOperationFailed
	}

	return nil
}

func (m *DeviceManager) SetRelayState(d *devices.Device, st bool) error {
	var r SystemRelayState
	r.System.RelayState.ErrorCode = 1
	r.SetRelayState(st)

	res, err := m.Marshal(d, r)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(res, &r); err != nil {
		return err
	}

	if r.ErrorCode() != 0 {
		return ErrProtocolOperationFailed
	}

	return nil
}

func (m *DeviceManager) Supports(d *devices.Device, feat Feature) bool {
	if feature, ok := featureMap[feat]; ok {
		return utils.Contains(d.Features(), feature)
	}

	return false
}

func (m *DeviceManager) SystemInfo(d devices.Addressable) (*SystemInfo, error) {
	var deviceInfo DeviceInfo

	res, err := m.Marshal(d, deviceInfo)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(res, &deviceInfo); err != nil {
		return nil, err
	}

	info := deviceInfo.SystemInfo()
	if info.ErrorCode != 0 {
		return nil, ErrProtocolOperationFailed
	}

	return info, nil
}
