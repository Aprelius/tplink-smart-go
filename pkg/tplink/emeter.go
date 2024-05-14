package tplink

import (
    "encoding/json"

    "github.com/Aralocke/tplink-smart-go/v1/pkg/devices"
)

type EMeter struct {
    device devices.Addressable
    mgr    *DeviceManager
}

func newElectricityMeter(d devices.Addressable, mgr *DeviceManager) *EMeter {
    return &EMeter{device: d, mgr: mgr}
}

func (e *EMeter) Realtime() (*RealTimeEnergy, error) {
    var emeterInfo ElectricityMeterInfo

    res, err := e.mgr.Marshal(e.device, emeterInfo)
    if err != nil {
        return nil, err
    }

    if err = json.Unmarshal(res, &emeterInfo); err != nil {
        return nil, err
    }

    return emeterInfo.Realtime(), nil
}
