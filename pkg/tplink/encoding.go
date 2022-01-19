package tplink

import (
    "bytes"
    "encoding/binary"
    "io"
)

const (
    EncryptionKey byte = 0xAB
)

func Decrypt(data []byte) (*bytes.Buffer, bool) {
    key := EncryptionKey

    if len(data) < 5 {
        return nil, false
    }

    buf := bytes.NewBuffer(data)

    var msgLen uint32
    if err := binary.Read(buf, binary.BigEndian, &msgLen); err != nil {
        return nil, false
    }

    msg := new(bytes.Buffer)
    msg.Grow(int(msgLen))

ForLoop:
    for {
        b, err := buf.ReadByte()
        if err != nil {
            switch err {
            case io.EOF:
                break ForLoop
            default:
                return nil, false
            }
        }

        plain := key ^ b
        key = b
        msg.WriteByte(plain)

        if uint32(msg.Len()) == msgLen {
            break
        }
    }

    return msg, msg.Len() != 0
}

func Encrypt(in []byte) (*bytes.Buffer, bool) {
    buf := new(bytes.Buffer)
    key := EncryptionKey

    if len(in) > 0 {
        buf.Grow(len(in) + 4)
        _ = binary.Write(buf, binary.BigEndian, uint32(len(in)))

        for _, c := range in {
            cipher := c ^ key
            key = cipher
            _ = binary.Write(buf, binary.BigEndian, cipher)
        }
    }

    return buf, buf.Len() != 0
}
