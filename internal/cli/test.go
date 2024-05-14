package cli

import (
	"fmt"
	"os"
	"time"

	"github.com/Aralocke/tplink-smart-go/v1/pkg/devices"
	"github.com/Aralocke/tplink-smart-go/v1/pkg/tplink"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type testCommandArgs struct {
}

var (
	testCmd = &cobra.Command{
		Long:  "",
		Short: "",
		Use:   "",
		RunE: func(cmd *cobra.Command, args []string) error {
			var cmdArgs testCommandArgs
			return runTestCmd(&cmdArgs)
		},
	}
)

func runTestCmd(args *testCommandArgs) error {
	_ = args

	log := rootLogger.WithOptions(
		zap.WrapCore(func(_ zapcore.Core) zapcore.Core {
			return rootLogger.Core()
		}),
	)

	manager := devices.NewDeviceManager(
		devices.WithLogger(log))
	api := tplink.NewDeviceManager(manager)

	config := tplink.PlugConfig("10.5.0.9")

	var device *devices.Device
	var err error

	if device, err = api.LoadDevice(&config); err != nil {
		fmt.Printf("Failed to load devices: %s", err)
		os.Exit(1)
	}

	info, err := api.SystemInfo(device)
	if err != nil {
		fmt.Printf("Failed to load SystemInfo for '%s:%d': %s",
			device.Address(), device.Port(), err)
		os.Exit(1)
	}

	tplink.DumpDeviceInfo(device, info)
	fmt.Println("---------------------")

	if err = api.SetRelayState(device, true); err != nil {
		fmt.Printf("failed to turn device on: %s", err)
		os.Exit(1)
	}

	time.Sleep(5 * time.Second)
	if err = api.SetRelayState(device, false); err != nil {
		fmt.Printf("failed to turn device off: %s", err)
		os.Exit(1)
	}

	if err = api.Reboot(device, 1); err != nil {
		fmt.Printf("failed to turn device off: %s", err)
		os.Exit(1)
	}

	return nil
}
