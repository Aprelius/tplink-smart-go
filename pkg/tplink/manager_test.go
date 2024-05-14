package tplink

import (
	"bytes"
	"fmt"
	"net"
	"strconv"
	"strings"
	"testing"
)

type MockTpLinkDevice struct {
	address string
	test    *testing.T
	sock    net.Listener
	System  SystemInfo
	port    uint16
}

func NewMockTpLinkDevice(t *testing.T, address string, port uint16) *MockTpLinkDevice {
	device := new(MockTpLinkDevice)
	device.address = address
	device.test = t
	device.port = port
	return device
}

func (d *MockTpLinkDevice) Address() string {
	return d.address
}

func (d *MockTpLinkDevice) clientHandler(client net.Conn) {
	defer func() { _ = client.Close() }()

	buf := make([]byte, 4096)
	length, err := client.Read(buf)
	if err != nil {
		d.Test().Errorf("failed to read from client socket '%s': %s\n",
			client.RemoteAddr(), err)
		return
	}
	if length == 0 {
		fmt.Printf("read zero bytes from '%s'", client.RemoteAddr())
		return
	}

	var data *bytes.Buffer
	var ok bool

	if data, ok = Decrypt(buf); !ok {
		fmt.Printf("failed to decrypt incomming data")
		return
	}

	_ = data
}

func (d *MockTpLinkDevice) Listen() {
	var err error

	addr := fmt.Sprintf("%s:%d", d.address, d.port)
	if d.sock, err = net.Listen("tcp", addr); err != nil {
		d.Test().Fatalf("failed to listen on '%s:%d': %s",
			d.address, d.port, err)
	}

	if d.port == 0 {
		addr := d.sock.Addr().String()
		in := strings.Split(addr, ":")
		if len(in) <= 1 {
			d.Test().Fatalf("invalid local address '%s'", addr)
		}

		var port int
		if port, err = strconv.Atoi(in[len(in)-1]); err != nil {
			d.Test().Fatalf("failed to convert local port: %s", err)
		}

		d.port = uint16(port)
	}
	go d.listenHandler()
}

func (d *MockTpLinkDevice) listenHandler() {
	for {
		sock, err := d.sock.Accept()
		if err != nil {
			fmt.Printf("failed to accept socket: %s", err)
			return
		}

		d.clientHandler(sock)
	}
}

func (d *MockTpLinkDevice) Port() uint16 {
	return d.port
}

func (d *MockTpLinkDevice) Stop() {
	_ = d.sock.Close()
}

func (d *MockTpLinkDevice) Test() *testing.T {
	return d.test
}

func TestManagerStartStop(t *testing.T) {
	/*
	   mock := NewMockTpLinkDevice(t, "12.0.0.1", 0)
	   mock.Listen()

	   manager := devices.NewDeviceManager()
	   api := NewDeviceManager(manager)

	   config := PlugConfig(mock.Address(), devices.WithPort(mock.Port()))
	   device, err := api.LoadDevice(&config)
	   if err != nil {
	       t.Fatalf("failed to load mock device: %s", err)
	   }

	   _ = device
	*/
}
