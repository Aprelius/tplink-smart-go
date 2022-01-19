package devices

type DeviceType int

const (
	UnknownDevice DeviceType = iota
	BulbDevice
	PlugDevice
)

var deviceMap = map[DeviceType]string{
	UnknownDevice: "Unknown",
	BulbDevice:    "Bulb",
	PlugDevice:    "Plug",
}

func (t DeviceType) String() string {
	if v, ok := deviceMap[t]; ok {
		return v
	}
	return deviceMap[UnknownDevice]
}

const (
    DefaultPort = 9999
)

type DeviceConfig struct {
    Address string
    Type    DeviceType
    Port    uint16
}

type DeviceConfigOption func(*DeviceConfig)

func NewDeviceConfig(
    address string,
    options ...DeviceConfigOption) DeviceConfig {

    cfg := DeviceConfig{}
    cfg.Address = address
    cfg.Type = UnknownDevice

    for _, opt := range options {
        opt(&cfg)
    }

    if cfg.Port == 0 {
        cfg.Port = DefaultPort
    }

    return cfg
}

func WithDeviceType(deviceType DeviceType) DeviceConfigOption {
    return func(c *DeviceConfig) {
        c.Type = deviceType
    }
}

func WithPort(port uint16) DeviceConfigOption {
    return func(c *DeviceConfig) {
        c.Port = port
    }
}

func NewDeviceConfigGroup(configs ...DeviceConfig) []DeviceConfig {
    return configs
}

type DeviceOption func(*Device)

type Addressable interface {
    Address() string
    Port() uint16
}

type Device struct {
    address     string
    deviceId    string
    deviceName  string
	deviceType  DeviceType
    features    []string
    firmwareId  string
    hardwareId  string
    hardwareVer string
    modelVer    string
    oemId       string
    port        uint16
    softwareVer string
}

var (
    _ Addressable = (*Device)(nil)
)

func WithDeviceId(deviceId string) DeviceOption {
    return func(d *Device) {
        d.deviceId = deviceId
    }
}

func WithDeviceName(deviceName string) DeviceOption {
    return func(d *Device) {
        d.deviceName = deviceName
    }
}

func WithFeatures(features []string) DeviceOption {
    return func(d *Device) {
        d.features = make([]string, len(features))
        copy(d.features, features)
    }
}

func WithFirmwareId(firmwareId string) DeviceOption {
    return func(d *Device) {
        d.firmwareId = firmwareId
    }
}

func WithHardwareId(hardwareId string) DeviceOption {
    return func(d *Device) {
        d.hardwareId = hardwareId
    }
}

func WithHardwareVersion(hardwareVer string) DeviceOption {
    return func(d *Device) {
        d.hardwareVer = hardwareVer
    }
}

func WithModelVersion(modelVer string) DeviceOption {
    return func(d *Device) {
        d.modelVer = modelVer
    }
}

func WithManufacturerId(oemId string) DeviceOption {
    return func(d *Device) {
        d.oemId = oemId
    }
}

func WithSoftwareVersion(softwareVer string) DeviceOption {
    return func(d *Device) {
        d.softwareVer = softwareVer
    }
}

func NewDevice(cfg *DeviceConfig, options ...DeviceOption) *Device {
	device := &Device{
        address: cfg.Address,
        deviceType: cfg.Type,
        port: cfg.Port}

    for _, option := range options {
        option(device)
    }

    return device
}

func (d *Device) Address() string {
    return d.address
}

func (d *Device) DeviceId() string {
	return d.deviceId
}

func (d *Device) DeviceName() string {
	return d.deviceName
}

func (d *Device) DeviceType() DeviceType {
	return d.deviceType
}

func (d *Device) Features() []string {
    return d.features
}

func (d *Device) FirmwareId() string {
	return d.firmwareId
}

func (d *Device) HardwareId() string {
	return d.hardwareId
}

func (d *Device) HardwareVersion() string {
	return d.hardwareVer
}

func (d *Device) ManufacturerId() string {
	return d.oemId
}

func (d *Device) Model() string {
	return d.modelVer
}

func (d *Device) Port() uint16 {
    return d.port
}

func (d *Device) SoftwareVersion() string {
	return d.softwareVer
}
