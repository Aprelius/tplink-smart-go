package tplink

import (
	"fmt"
	"strings"

	"github.com/Aralocke/tplink-smart-go/v1/pkg/devices"
	"github.com/Aralocke/tplink-smart-go/v1/pkg/utils"
)

func DumpDeviceInfo(device *devices.Device, info *SystemInfo) {
	fmt.Println("Device Information")
	fmt.Printf("\tAddress: %s\n", device.Address())

	fmt.Printf("\tAlias: %s\n", info.Alias)

	fmt.Printf("\tDevice Model: %s\n", device.Model())
	fmt.Printf("\tDevice Type: %s\n", device.DeviceType())
	fmt.Printf("\tDevice Identifier: %s\n", device.DeviceId())
	fmt.Printf("\tDevice Name: %s\n", device.DeviceName())
	fmt.Printf("\tFeatures: %s\n", strings.Join(device.Features(), ", "))
	fmt.Println("Device State")

	fmt.Printf("\tUptime: %s\n", utils.PrettyDuration(info.UpTime))

	if info.RelayState != 0 {
		fmt.Println("\tState: On")
	} else {
		fmt.Println("\tState: Off")
	}

	fmt.Println("Version Information")
	fmt.Printf("\tSoftware Version: %s\n", device.SoftwareVersion())
	fmt.Printf("\tHardware Version: %s\n", device.HardwareVersion())
	fmt.Println("Network Status")

	fmt.Printf("\tMAC Address: %s\n", info.MacAddress)
	fmt.Printf("\tSignal Strength: %s (%d)\n",
		utils.SignalStrength(info.SignalStrength),
		info.SignalStrength)
}
