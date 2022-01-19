package tplink

import "math"

// Basic Instructions

type aliasValue struct {
	errorCode
	Value string `json:"alias"`
}

type delayTime struct {
	errorCode
	Delay int `json:"delay"`
}

type deviceIdValue struct {
	errorCode
	Value string `json:"deviceId"`
}

type errorCode struct {
	ErrorCode int `json:"err_code,omitempty"`
}

type hardwareIdValue struct {
	errorCode
	Value string `json:"hwId"`
}

type locationValues struct {
	errorCode
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type networkValues struct {
	errorCode
	NetworkName string `json:"ssid"`
	Password    string `json:"password"`
	KeyType     int    `json:"key_type"`
}

type onOffValue struct {
	errorCode
	Value int `json:"off"`
}

type stateValue struct {
	errorCode
	State int `json:"state"`
}

// EnergyMeter Types

type RealTimeEnergy struct {
	errorCode
	Current float32 `json:"current,omitempty"`
	Voltage float32 `json:"voltage,omitempty"`
	Power   float32 `json:"power,omitempty"`
	Total   float32 `json:"total,omitempty"`
}

type GetRealTimeEnergy struct {
	Energy RealTimeEnergy `json:"get_realtime"`
}

type ElectricityMeterInfo struct {
	EMeter GetRealTimeEnergy `json:"emeter"`
}

func (e *ElectricityMeterInfo) Realtime() *RealTimeEnergy {
	return &e.EMeter.Energy
}

// Daily Average
// SEND {"emeter": {"get_daystat": {"month": 11, "year": 2021}}}
// RECV {"emeter":{"get_daystat":{"day_list":[{"year":2021,"month":11,"day":1,"energy":1.172000},{"year":2021,"month":11,"day":2,"energy":1.170000},{"year":2021,"month":11,"day":3,"energy":1.128000},{"year":2021,"month":11,"day":4,"energy":1.148000},{"year":2021,"month":11,"day":5,"energy":1.171000},{"year":2021,"month":11,"day":6,"energy":1.169000},{"year":2021,"month":11,"day":7,"energy":1.166000},{"year":2021,"month":11,"day":8,"energy":1.163000},{"year":2021,"month":11,"day":9,"energy":1.159000},{"year":2021,"month":11,"day":10,"energy":1.126000},{"year":2021,"month":11,"day":11,"energy":1.130000},{"year":2021,"month":11,"day":12,"energy":1.133000},{"year":2021,"month":11,"day":13,"energy":1.131000},{"year":2021,"month":11,"day":14,"energy":0.899000}],"err_code":0}}}
// 1.133214285714286

// Daily Usage
// SEND {"emeter": {"get_daystat": {"month": 11, "year": 2021}}}
// RECV {"emeter":{"get_daystat":{"day_list":[{"year":2021,"month":11,"day":1,"energy":1.172000},{"year":2021,"month":11,"day":2,"energy":1.170000},{"year":2021,"month":11,"day":3,"energy":1.128000},{"year":2021,"month":11,"day":4,"energy":1.148000},{"year":2021,"month":11,"day":5,"energy":1.171000},{"year":2021,"month":11,"day":6,"energy":1.169000},{"year":2021,"month":11,"day":7,"energy":1.166000},{"year":2021,"month":11,"day":8,"energy":1.163000},{"year":2021,"month":11,"day":9,"energy":1.159000},{"year":2021,"month":11,"day":10,"energy":1.126000},{"year":2021,"month":11,"day":11,"energy":1.130000},{"year":2021,"month":11,"day":12,"energy":1.133000},{"year":2021,"month":11,"day":13,"energy":1.131000},{"year":2021,"month":11,"day":14,"energy":0.900000}],"err_code":0}}}
// [{'year': 2021, 'month': 11, 'day': 1, 'energy': 1.172}, {'year': 2021, 'month': 11, 'day': 2, 'energy': 1.17}, {'year': 2021, 'month': 11, 'day': 3, 'energy': 1.128}, {'year': 2021, 'month': 11, 'day': 4, 'energy': 1.148}, {'year': 2021, 'month': 11, 'day': 5, 'energy': 1.171}, {'year': 2021, 'month': 11, 'day': 6, 'energy': 1.169}, {'year': 2021, 'month': 11, 'day': 7, 'energy': 1.166}, {'year': 2021, 'month': 11, 'day': 8, 'energy': 1.163}, {'year': 2021, 'month': 11, 'day': 9, 'energy': 1.159}, {'year': 2021, 'month': 11, 'day': 10, 'energy': 1.126}, {'year': 2021, 'month': 11, 'day': 11, 'energy': 1.13}, {'year': 2021, 'month': 11, 'day': 12, 'energy': 1.133}, {'year': 2021, 'month': 11, 'day': 13, 'energy': 1.131}, {'year': 2021, 'month': 11, 'day': 14, 'energy': 0.9}]

type DayStat struct {
	Day    int     `json:"day,omitempty"`
	Energy float32 `json:"energy,omitempty"`
	Month  int     `json:"month,omitempty"`
	Year   int     `json:"year,omitempty"`
}

type DailyStats struct {
	DayList []DayStat `json:"day_list,omitempty"`
	errorCode
	Month int `json:"month,omitempty"`
	Year  int `json:"year,omitempty"`
}

type getDailyStat struct {
	Stats DailyStats `json:"get_daystat"`
}

type EMeterDailyStats struct {
	EMeter getDailyStat `json:"emeter"`
}

func (e *EMeterDailyStats) DailyStats() *DailyStats {
	return &e.EMeter.Stats
}

// Monthly Average
// SEND {"emeter": {"get_daystat": {"month": 11, "year": 2021}}}
// RECV {"emeter":{"get_daystat":{"day_list":[{"year":2021,"month":11,"day":1,"energy":1.172000},{"year":2021,"month":11,"day":2,"energy":1.170000},{"year":2021,"month":11,"day":3,"energy":1.128000},{"year":2021,"month":11,"day":4,"energy":1.148000},{"year":2021,"month":11,"day":5,"energy":1.171000},{"year":2021,"month":11,"day":6,"energy":1.169000},{"year":2021,"month":11,"day":7,"energy":1.166000},{"year":2021,"month":11,"day":8,"energy":1.163000},{"year":2021,"month":11,"day":9,"energy":1.159000},{"year":2021,"month":11,"day":10,"energy":1.126000},{"year":2021,"month":11,"day":11,"energy":1.130000},{"year":2021,"month":11,"day":12,"energy":1.133000},{"year":2021,"month":11,"day":13,"energy":1.131000},{"year":2021,"month":11,"day":14,"energy":0.900000}],"err_code":0}}}
// 1.1332857142857145

// Monthly Usage
// SEND {"emeter": {"get_monthstat": {"year": 2021}}}
// RECV {"emeter":{"get_monthstat":{"month_list":[{"year":2021,"month":1,"energy":23.111000},{"year":2021,"month":2,"energy":168.620000},{"year":2021,"month":3,"energy":12.785000},{"year":2021,"month":4,"energy":31.806000},{"year":2021,"month":5,"energy":33.877000},{"year":2021,"month":6,"energy":32.959000},{"year":2021,"month":7,"energy":34.774000},{"year":2021,"month":8,"energy":35.270000},{"year":2021,"month":9,"energy":35.105000},{"year":2021,"month":10,"energy":36.254000},{"year":2021,"month":11,"energy":15.867000}],"err_code":0}}}
// [{'year': 2021, 'month': 1, 'energy': 23.111}, {'year': 2021, 'month': 2, 'energy': 168.62}, {'year': 2021, 'month': 3, 'energy': 12.785}, {'year': 2021, 'month': 4, 'energy': 31.806}, {'year': 2021, 'month': 5, 'energy': 33.877}, {'year': 2021, 'month': 6, 'energy': 32.959}, {'year': 2021, 'month': 7, 'energy': 34.774}, {'year': 2021, 'month': 8, 'energy': 35.27}, {'year': 2021, 'month': 9, 'energy': 35.105}, {'year': 2021, 'month': 10, 'energy': 36.254}, {'year': 2021, 'month': 11, 'energy': 15.867}]

type MonthlyStats struct {
	MonthList []DayStat `json:"month_list,omitempty"`
	errorCode
	Year int `json:"year,omitempty"`
}

type getMonthStat struct {
	Stats MonthlyStats `json:"get_monthstat"`
}

type EMeterMonthlyStats struct {
	EMeter getMonthStat `json:"emeter"`
}

func (e *EMeterMonthlyStats) MonthlyStats() *MonthlyStats {
	return &e.EMeter.Stats
}

// Clear Usage Stats
// SEND {"emeter": {"erase_emeter_stat": {}}}
// RECV {"emeter":{"erase_emeter_stat":{"err_code":0}}}
// {'emeter': {'erase_emeter_stat': {'err_code': 0}}}

// SystemInfo Types

type SystemInfo struct {
	errorCode
	SoftwareVersion string  `json:"sw_ver,omitempty"`
	HardwareVersion string  `json:"hw_ver,omitempty"`
	Type            string  `json:"type,omitempty"`
	Model           string  `json:"model,omitempty"`
	MacAddress      string  `json:"mac,omitempty"`
	DeviceId        string  `json:"deviceId,omitempty"`
	HardwareId      string  `json:"hwId,omitempty"`
	FirmwareId      string  `json:"fwId,omitempty"`
	ManufacturerId  string  `json:"oemId,omitempty"`
	Alias           string  `json:"alias,omitempty"`
	DeviceName      string  `json:"dev_name,omitempty"`
	IconHash        string  `json:"icon_hash,omitempty"`
	RelayState      int     `json:"relay_state,omitempty"`
	UpTime          int     `json:"on_time,omitempty"`
	ActiveMode      string  `json:"active_mode,omitempty"`
	Features        string  `json:"feature,omitempty"`
	Updating        int     `json:"updating,omitempty"`
	SignalStrength  int     `json:"rssi,omitempty"`
	LedStatus       int     `json:"led_off,omitempty"`
	Latitude        float32 `json:"latitude,omitempty"`
	Longitude       float32 `json:"longitude,omitempty"`
}

type GetSystemInfo struct {
	Info SystemInfo `json:"get_sysinfo"`
}

type DeviceInfo struct {
	System GetSystemInfo `json:"system"`
}

func (d *DeviceInfo) SystemInfo() *SystemInfo {
	return &d.System.Info
}

// Network Configuration

type networkSettings struct {
	Settings networkValues `json:"set_stainfo"`
}

type NetworkSettings struct {
	Interface networkSettings `json:"netif"`
}

func (n *NetworkSettings) ErrorCode() int {
    return n.Interface.Settings.ErrorCode
}

func (n *NetworkSettings) GetNetwork() string {
	return n.Interface.Settings.NetworkName
}

func (n *NetworkSettings) GetPassword() string {
	return n.Interface.Settings.Password
}

func (n *NetworkSettings) SetSettings(network string, password string) {
	n.Interface.Settings.KeyType = 3 // WPA2
	n.Interface.Settings.NetworkName = network
	n.Interface.Settings.Password = password
}

// System Alias

type systemAlias struct {
	Alias aliasValue `json:"set_dev_alias"`
}

type SystemAlias struct {
	System systemAlias `json:"system"`
}

func (a *SystemAlias) ErrorCode() int {
    return a.System.Alias.ErrorCode
}

func (a *SystemAlias) GetAlias() string {
	return a.System.Alias.Value
}

func (a *SystemAlias) SetAlias(s string) {
	a.System.Alias.Value = s
}

// System Device ID

type deviceId struct {
	DeviceId deviceIdValue `json:"set_device_id"`
}

type SystemDeviceId struct {
	System deviceId `json:"system"`
}

func (d *SystemDeviceId) ErrorCode() int {
    return d.System.DeviceId.ErrorCode
}

func (d *SystemDeviceId) GetDeviceId() string {
	return d.System.DeviceId.Value
}

func (d *SystemDeviceId) SetDeviceId(id string) {
	d.System.DeviceId.Value = id
}

// System Hardware ID

type hardwareId struct {
	HardwareId hardwareIdValue `json:"set_hw_id"`
}

type SystemHardwareId struct {
	System hardwareId `json:"system"`
}

func (d *SystemHardwareId) ErrorCode() int {
    return d.System.HardwareId.ErrorCode
}

func (d *SystemHardwareId) GetHardwareId() string {
	return d.System.HardwareId.Value
}

func (d *SystemHardwareId) SetHardwareId(id string) {
	d.System.HardwareId.Value = id
}

// LED Status

type ledState struct {
	LedState onOffValue `json:"set_led_off"`
}

type SystemLedState struct {
	System ledState `json:"system"`
}

func (s *SystemLedState) ErrorCode() int {
    return s.System.LedState.ErrorCode
}

func (s *SystemLedState) GetState() bool {
	return s.System.LedState.Value != 0
}

func (s *SystemLedState) SetState(st bool) {
	if st {
		s.System.LedState.Value = 1
	} else {
		s.System.LedState.Value = 0
	}
}

// System Location

type locationValue struct {
	Location locationValues `json:"set_dev_location"`
}

type SystemLocation struct {
	System locationValue `json:"system"`
}

func (loc *SystemLocation) ErrorCode() int {
    return loc.System.Location.ErrorCode
}

func (loc *SystemLocation) GetLatitude() float64 {
	lat := loc.System.Location.Latitude
	return math.Floor(lat*1000) / 1000
}

func (loc *SystemLocation) GetLongitude() float64 {
	lon := loc.System.Location.Longitude
	return math.Floor(lon*1000) / 1000
}

func (loc *SystemLocation) SetLocation(lat float64, lon float64) {
	loc.System.Location.Latitude = lat
	loc.System.Location.Longitude = lon
}

// Reboot Command

type reboot struct {
	Reboot delayTime `json:"reboot"`
}

type SystemReboot struct {
	System reboot `json:"system"`
}

func (r *SystemReboot) ErrorCode() int {
	return r.System.Reboot.ErrorCode
}

func (r *SystemReboot) GetDelay() int {
	return r.System.Reboot.Delay
}

func (r *SystemReboot) SetDelay(delay int) {
	r.System.Reboot.Delay = delay
}

// Relay State

type relayState struct {
	RelayState stateValue `json:"set_relay_state"`
}

type SystemRelayState struct {
	System relayState `json:"system"`
}

func (r *SystemRelayState) ErrorCode() int {
	return r.System.RelayState.ErrorCode
}

func (r *SystemRelayState) GetRelayState() bool {
	return r.System.RelayState.State != 0
}

func (r *SystemRelayState) SetRelayState(st bool) {
	if st {
		r.System.RelayState.State = 1
	} else {
		r.System.RelayState.State = 0
	}
}

// Reset Command

type reset struct {
	Reset delayTime `json:"reset"`
}

type SystemReset struct {
	System reset `json:"system"`
}

func (r *SystemReset) ErrorCode() int {
	return r.System.Reset.ErrorCode
}

func (r *SystemReset) GetDelay() int {
	return r.System.Reset.Delay
}

func (r *SystemReset) SetDelay(d int) {
	r.System.Reset.Delay = d
}
