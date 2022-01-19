package devices

import (
    "go.uber.org/zap"
    "time"
)

type DeviceManagerOption func(*DeviceManagerOptions)

type DeviceManagerOptions struct {
    ConnectTimeout time.Duration
    Logger         *zap.Logger
    RecvTimeout    time.Duration
}

func WithLogger(logger *zap.Logger) DeviceManagerOption {
    return func(dm *DeviceManagerOptions) {
        dm.Logger = logger
    }
}

func DefaultDeviceManagerOptions() *DeviceManagerOptions {
    return &DeviceManagerOptions{
        ConnectTimeout: 60 * time.Second,
        RecvTimeout:    60 * time.Second,
    }
}

type DeviceManager struct {
    connectTimeout time.Duration
    devices        []*Device
    logger         *zap.Logger
    recvTimeout    time.Duration
}

func NewDeviceManager(opts ...DeviceManagerOption) *DeviceManager {
    options := DefaultDeviceManagerOptions()
    for _, option := range opts {
        option(options)
    }

    if options.Logger == nil {
        options.Logger = zap.NewNop()
    }

    mgr := new(DeviceManager)
    mgr.connectTimeout = options.ConnectTimeout
    mgr.logger = options.Logger
    mgr.recvTimeout = options.RecvTimeout

    return mgr
}

func (dm *DeviceManager) Logger() *zap.Logger {
    return dm.logger
}

func (dm *DeviceManager) NewDevice(cfg *DeviceConfig, options ...DeviceOption) *Device {
    device := NewDevice(cfg, options...)
    dm.devices = append(dm.devices, device)

    return device
}
