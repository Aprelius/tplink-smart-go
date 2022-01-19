package tplink

import (
    "encoding/json"
    "fmt"
    "math"
    "testing"
)

func floatCompare(left float64, right float64) bool {
	const Threshold = 1e-9
	return math.Abs(left-right) <= Threshold
}

func TestNetworkSettings(t *testing.T) {
	var n NetworkSettings

	network := "abc123"
	password := "def456"
	data := "{\"netif\":{\"set_stainfo\":{\"ssid\":\"%s\",\"password\":\"%s\",\"key_type\":3}}}"

	if err := json.Unmarshal(
		[]byte(fmt.Sprintf(data, network, password)),
		&n); err != nil {
		t.Fatalf("failed to unmarshal system alias: %s", err)
	}
	if n.GetNetwork() != network {
		t.Fatalf("unexpected network '%s'", n.GetNetwork())
	}
	if n.GetPassword() != password {
		t.Fatalf("unexpected password '%s'", n.GetPassword())
	}

	network = "123abc"
	password = "456def"
	n.SetSettings(network, password)
	s, err := json.Marshal(n)
	expected := fmt.Sprintf(data, network, password)

	if err != nil {
		t.Fatalf("failed to marshal network settings")
	}

	generated := string(s)
	if generated != expected {
		t.Fatalf("failed to generate '%s' expected string '%s'",
			generated, expected)
	}
}

func TestSystemAlias(t *testing.T) {
	var a SystemAlias

	alias := "alias"
	data := "{\"system\":{\"set_dev_alias\":{\"alias\":\"%s\"}}}"

	if err := json.Unmarshal([]byte(fmt.Sprintf(data, alias)), &a); err != nil {
		t.Fatalf("failed to unmarshal system alias: %s", err)
	}
	if a.GetAlias() != alias {
		t.Fatalf("unexpected relay state '%s'", a.GetAlias())
	}

	alias = "abc123"
	a.SetAlias(alias)
	s, err := json.Marshal(a)
	expected := fmt.Sprintf(data, alias)

	if err != nil {
		t.Fatalf("failed to marshal relay state")
	}

	generated := string(s)
	if generated != expected {
		t.Fatalf("failed to generate '%s' expected string '%s'",
			generated, expected)
	}
}

func TestSystemDeviceId(t *testing.T) {
	var h SystemDeviceId

	devId := "abc123"
	data := "{\"system\":{\"set_device_id\":{\"deviceId\":\"%s\"}}}"

	if err := json.Unmarshal([]byte(fmt.Sprintf(data, devId)), &h); err != nil {
		t.Fatalf("failed to unmarshal system device id: %s", err)
	}
	if h.GetDeviceId() != devId {
		t.Fatalf("unexpected device id '%s'", h.GetDeviceId())
	}

	devId = "def456"
	h.SetDeviceId(devId)
	s, err := json.Marshal(h)
	expected := fmt.Sprintf(data, devId)

	if err != nil {
		t.Fatalf("failed to marshal device id")
	}

	generated := string(s)
	if generated != expected {
		t.Fatalf("failed to generate '%s' expected string '%s'",
			generated, expected)
	}
}

func TestSystemHardwareId(t *testing.T) {
	var h SystemHardwareId

	hwId := "abc123"
	data := "{\"system\":{\"set_hw_id\":{\"hwId\":\"%s\"}}}"

	if err := json.Unmarshal([]byte(fmt.Sprintf(data, hwId)), &h); err != nil {
		t.Fatalf("failed to unmarshal system hardare id: %s", err)
	}
	if h.GetHardwareId() != hwId {
		t.Fatalf("unexpected hardware id '%s'", h.GetHardwareId())
	}

	hwId = "def456"
	h.SetHardwareId(hwId)
	s, err := json.Marshal(h)
	expected := fmt.Sprintf(data, hwId)

	if err != nil {
		t.Fatalf("failed to marshal hardare id")
	}

	generated := string(s)
	if generated != expected {
		t.Fatalf("failed to generate '%s' expected string '%s'",
			generated, expected)
	}
}

func TestSystemLedState(t *testing.T) {
	var r SystemLedState
	data := "{\"system\":{\"set_led_off\":{\"off\":%d}}}"

	if err := json.Unmarshal([]byte(fmt.Sprintf(data, 1)), &r); err != nil {
		t.Fatalf("failed to unmarshal system led state: %s", err)
	}
	if !r.GetState() {
		t.Fatalf("unexpected led state '%t'", r.GetState())
	}

	r.SetState(false)
	s, err := json.Marshal(r)
	expected := fmt.Sprintf(data, 0)

	if err != nil {
		t.Fatalf("failed to marshal led state")
	}

	generated := string(s)
	if generated != expected {
		t.Fatalf("failed to generate '%s' expected string '%s'",
			generated, expected)
	}
}

func TestSystemLocation(t *testing.T) {
	var loc SystemLocation
	data := "{\"system\":{\"set_dev_location\":{\"latitude\":%.3f,\"longitude\":%.3f}}}"

	var lat = 456.123
	var lon = 123.456

	if err := json.Unmarshal([]byte(fmt.Sprintf(data, lat, lon)), &loc); err != nil {
		t.Fatalf("failed to unmarshal system location: %s", err)
	}
	if !floatCompare(loc.GetLatitude(), lat) {
		t.Fatalf("unexpected latitude value '%.3f'", loc.GetLatitude())
	}
	if !floatCompare(loc.GetLongitude(), lon) {
		t.Fatalf("unexpected longitude value '%.3f'", loc.GetLongitude())
	}

	lat = 987.654
	lon = 654.987

	loc.SetLocation(lat, lon)
	s, err := json.Marshal(loc)
	expected := fmt.Sprintf(data, lat, lon)

	if err != nil {
		t.Fatalf("failed to marshal location values")
	}

	generated := string(s)
	if generated != expected {
		t.Fatalf("failed to generate '%s' expected string '%s'",
			generated, expected)
	}
}

func TestSystemRelayState(t *testing.T) {
	var r SystemRelayState
	data := "{\"system\":{\"set_relay_state\":{\"state\":%d}}}"

	if err := json.Unmarshal([]byte(fmt.Sprintf(data, 1)), &r); err != nil {
		t.Fatalf("failed to unmarshal system relay state: %s", err)
	}
	if !r.GetRelayState() {
		t.Fatalf("unexpected relay state '%t'", r.GetRelayState())
	}

	r.SetRelayState(false)
	s, err := json.Marshal(r)
	expected := fmt.Sprintf(data, 0)

	if err != nil {
		t.Fatalf("failed to marshal relay state")
	}

	generated := string(s)
	if generated != expected {
		t.Fatalf("failed to generate '%s' expected string '%s'",
			generated, expected)
	}
}

func TestSystemReboot(t *testing.T) {
	var r SystemReboot
	data := "{\"system\":{\"reboot\":{\"delay\":%d}}}"

	if err := json.Unmarshal([]byte(fmt.Sprintf(data, 60)), &r); err != nil {
		t.Fatalf("failed to unmarshal system reboot: %s", err)
	}

	if r.GetDelay() != 60 {
		t.Fatalf("unexpected duration '%d'", r.GetDelay())
	}

	r.SetDelay(55)
	s, err := json.Marshal(r)
	expected := fmt.Sprintf(data, 55)

	if err != nil {
		t.Fatalf("failed to marshal reboot delay")
	}

	generated := string(s)
	if generated != expected {
		t.Fatalf("failed to generate '%s' expected string '%s'",
			generated, expected)
	}
}

func TestSystemReset(t *testing.T) {
	var r SystemReset
	data := "{\"system\":{\"reset\":{\"delay\":%d}}}"

	if err := json.Unmarshal([]byte(fmt.Sprintf(data, 60)), &r); err != nil {
		t.Fatalf("failed to unmarshal system reset: %s", err)
	}

	if r.GetDelay() != 60 {
		t.Fatalf("unexpected duration '%d'", r.GetDelay())
	}

	r.SetDelay(55)
	s, err := json.Marshal(r)
	expected := fmt.Sprintf(data, 55)

	if err != nil {
		t.Fatalf("failed to marshal reset delay")
	}

	generated := string(s)
	if generated != expected {
		t.Fatalf("failed to generate '%s' expected string '%s'",
			generated, expected)
	}
}

func TestMarshalSystemInfo(t *testing.T) {
	var d DeviceInfo

	data := "{\"system\": {\"get_sysinfo\": {\"err_code\": 0, \"sw_ver\": \"1.1.1 Build 160725 Rel.164033\", \"hw_ver\": \"1.0\", \"type\": \"IOT.SMARTPLUGSWITCH\", \"model\": \"HS110(US)\", \"mac\": \"50:C7:BF:84:2F:03\", \"deviceId\": \"80069BA4A24AA2A34AF01633A5F4A02A1886345F\", \"hwId\": \"60FF6B258734EA6880E186F8C96DDC61\", \"fwId\": \"060BFEA28A8CD1E67146EB5B2B599CC8\", \"oemId\": \"FFF22CFF774A0B89F7624BFC6F50D5DE\", \"alias\": \"Core HS110\", \"dev_name\": \"Wi-Fi Smart Plug With Energy Monitoring\", \"icon_hash\": \"\", \"relay_state\": 1, \"on_time\": 10726298, \"active_mode\": \"none\", \"feature\": \"TIM:ENE\", \"updating\": 0, \"rssi\": -61, \"led_off\": 0, \"latitude\": 0, \"longitude\": 0}}}"
	if err := json.Unmarshal([]byte(data), &d); err != nil {
		t.Fatalf("failed to unmarshal system info: %s", err)
	}
	info := d.SystemInfo()
	if info.ErrorCode != 0 {
		t.Fatalf("unexpected error code: %d", info.ErrorCode)
	}
	if info.HardwareVersion != "1.0" {
		t.Fatalf("unexpected hardware version code: %s", info.HardwareVersion)
	}
}

func TestUnmarshalSystemInfo(t *testing.T) {
	var d DeviceInfo
	data, err := json.Marshal(d)
	if err != nil {
		t.Fatalf("failed to marshal empty device info: %s", err)
	}

	generated := string(data)
	expected := "{\"system\":{\"get_sysinfo\":{}}}"

	if generated != expected {
		t.Fatalf("failed to generate '%s' expected string '%s'",
			generated, expected)
	}
}
