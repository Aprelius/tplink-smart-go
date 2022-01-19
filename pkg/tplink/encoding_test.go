package tplink

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"testing"
)

var testData = map[string][]byte{
	"test123": {0x00, 0x00, 0x00, 0x07, 0xDF, 0xBA, 0xC9, 0xBD, 0x8C, 0xBE, 0x8D},
	"abcdefhijklmnopqrstuvwxyz": {0x0, 0x0, 0x0, 0x19, 0xca, 0xa8, 0xcb, 0xaf, 0xca, 0xac,
		0xc4, 0xad, 0xc7, 0xac, 0xc0, 0xad, 0xc3, 0xac, 0xdc, 0xad, 0xdf, 0xac, 0xd8, 0xad,
		0xdb, 0xac, 0xd4, 0xad, 0xd7},
	"{\"test123\": {\"abc\": 456}}": {0x0, 0x0, 0x0, 0x19, 0xd0, 0xf2, 0x86, 0xe3, 0x90, 0xe4,
		0xd5, 0xe7, 0xd4, 0xf6, 0xcc, 0xec, 0x97, 0xb5, 0xd4, 0xb6, 0xd5, 0xf7, 0xcd, 0xed,
		0xd9, 0xec, 0xda, 0xa7, 0xda},
}

func CompareDecrypted(t *testing.T, in string, out []byte) {
	fmt.Printf("msg = '%s'\n", in)
	fmt.Printf("input(%d) = %x\n", len(in), md5.Sum([]byte(in)))
	fmt.Printf("output(%d) = %x\n", len(out), md5.Sum(out))

	if string(out) != in {
		t.Fatalf("Decrypted message '%s' does not match '%s'", out, in)
	}
}

func CompareEncrypted(t *testing.T, msg string, in []byte, out []byte) {
	fmt.Printf("msg = '%s'\n", msg)
	fmt.Printf("input(%d) = %x\n", len(in), md5.Sum(in))
	fmt.Printf("output(%d) = %x\n", len(out), md5.Sum(out))

	if res := bytes.Compare(in, out); res != 0 {
		t.Fatalf("Encrypting '%s' did not generate expected result", msg)
	}
}

func DecryptMessage(t *testing.T, msg string, in []byte) *bytes.Buffer {
	if buf, ok := Decrypt(in); ok {
		return buf
	}
    t.Fatalf("Decrypting '%s' failed", msg)
    return nil
}

func EncryptMessage(t *testing.T, msg []byte) *bytes.Buffer {
	if buf, ok := Encrypt(msg); ok {
		return buf
	}
    t.Fatalf("Encrypting '%s' failed", string(msg))
	return nil
}

func TestTpLinkEncrypt(t *testing.T) {
	for msg, expected := range testData {
		buf := EncryptMessage(t, []byte(msg))
		CompareEncrypted(t, msg, expected, buf.Bytes())
	}
}

func TestTpLinkDecrypt(t *testing.T) {
	for msg, in := range testData {
		output := DecryptMessage(t, msg, in)
		CompareDecrypted(t, msg, output.Bytes())
	}
}
