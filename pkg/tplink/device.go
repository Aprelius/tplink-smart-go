package tplink

import (
	"github.com/Aprelius/tplink-smart-go/v1/pkg/devices"
)

func BulbConfig(address string, options... devices.DeviceConfigOption) devices.DeviceConfig {
    return devices.NewDeviceConfig(address,
        append(options,
            devices.WithDeviceType(devices.BulbDevice))...)
}

func PlugConfig(address string, options... devices.DeviceConfigOption) devices.DeviceConfig {
    return devices.NewDeviceConfig(address,
        append(options,
            devices.WithDeviceType(devices.BulbDevice))...)
}
