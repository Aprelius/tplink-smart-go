package devices

import (
    "bytes"
    "context"
    "errors"
    "fmt"
    "net"
    "time"
)

type DecodeFn func([]byte) (*bytes.Buffer, bool)
type EncodeFn func([]byte) (*bytes.Buffer, bool)

type SenderOption func(*SyncSender)

func WithEncoding(decodeFn DecodeFn, encodeFn EncodeFn) SenderOption{
    return func(s *SyncSender) {
        s.decodeFn = decodeFn
        s.encodeFn = encodeFn
    }
}

func WithTimeout(timeout time.Duration) SenderOption {
    return func(s *SyncSender) {
        s.timeout = timeout
    }
}

type SyncSender struct {
    ctx      context.Context
    device   Addressable
    decodeFn DecodeFn
    encodeFn EncodeFn
    timeout  time.Duration
}

func NewSyncSender(d Addressable, options ...SenderOption) *SyncSender {
    sender := &SyncSender{device: d}
    sender.ctx = context.Background()

    for _, opt := range options {
        opt(sender)
    }

    return sender
}

func (s *SyncSender) Send(data []byte) ([]byte, error) {
    var encoded *bytes.Buffer
    if s.encodeFn != nil {
        var ok bool
        if encoded, ok = s.encodeFn(data); !ok {
            return []byte{}, errors.New("")
        }
    }

    var ctx context.Context
    var cancelFn context.CancelFunc

    if s.timeout > 0 {
        ctx, cancelFn = context.WithTimeout(s.ctx, s.timeout)
    } else {
        ctx, cancelFn = context.WithCancel(s.ctx)
    }
    defer cancelFn()

    var dialer net.Dialer
    address := fmt.Sprintf(
        "%s:%d",
        s.device.Address(),
        s.device.Port())

    sock, err := dialer.DialContext(ctx, "tcp", address)
    if err != nil {
        return []byte{}, err
    }

    defer func() { _ = sock.Close() }()

    _, err = sock.Write(encoded.Bytes())
    if err != nil {
        return []byte{}, err
    }

    buffer := make([]byte, 4096)
    _, err = sock.Read(buffer)

    if err != nil {
        return []byte{}, err
    }

    if s.decodeFn != nil {
        if buf, ok := s.decodeFn(buffer); ok {
            return buf.Bytes(), nil
        }
        return []byte{}, errors.New("")
    }

    return buffer, nil
}
