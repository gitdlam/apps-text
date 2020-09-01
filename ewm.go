package text

import (
	"encoding/base64"
	"encoding/hex"
	"math/big"
	"strings"
)

const (
	encodeEWM = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz{}"
)

var (
	EWMEncoding *base64.Encoding
)

func init() {
	EWMEncoding = base64.NewEncoding(encodeEWM).WithPadding(base64.NoPadding)
}

// UUID, or GUID, is a 128 bit number used as an ID number. The same number can be represented in different forms.
//
// EWMEncodeUUID() converts hexadecimal to a non-standard base64 used by the EWM system.
//
// EWMDecodeUUID() converts the the non-standard base64 to hexadecimal.
func EWMEncodeUUID(hx string) string {
	b, err := hex.DecodeString(hx)
	if err != nil {
		return ""
	}
	padding := make([]byte, 16-len(b))
	return EWMEncoding.EncodeToString(append(padding, b...))
}

// UUID, or GUID, is a 128 bit number used as an ID number. The same number can be represented in different forms.
//
// EWMEncodeUUID() converts hexadecimal to a non-standard base64 used by the EWM system.
//
// EWMDecodeUUID() converts the the non-standard base64 to hexadecimal.
func EWMDecodeUUID(c22 string) string {
	b, _ := EWMEncoding.DecodeString(c22)
	s := hex.EncodeToString(b)
	return strings.Repeat("0", 32-len(s)) + strings.ToUpper(s)
}

// slightly slower
func EWMEncodeUUIDSlower(hex string) string {
	bigInt := new(big.Int)
	bigInt.SetString(hex, 16)
	padding := make([]byte, 16-len(bigInt.Bytes()))
	return EWMEncoding.EncodeToString(append(padding, bigInt.Bytes()...))
}

// not tested for speed but expect to be slightly slower
func EWMDecodeUUIDSlower(c22 string) string {
	b, _ := EWMEncoding.DecodeString(c22)
	bigInt := new(big.Int)
	bigInt.SetBytes(b)
	s := strings.ToUpper(bigInt.Text(16))
	return strings.Repeat("0", 32-len(s)) + s
}
