package crypto

import (
	"fmt"
	"syscall"
	"unsafe"
)

var (
	modcrypt32      = syscall.NewLazyDLL("crypt32.dll")
	procDecryptData = modcrypt32.NewProc("CryptUnprotectData")
)

type DATA_BLOB struct {
	cbData uint32
	pData  *byte
}

func newBlob(data []byte) *DATA_BLOB {
	if len(data) == 0 {
		return &DATA_BLOB{}
	}
	return &DATA_BLOB{
		cbData: uint32(len(data)),
		pData:  &data[0],
	}
}

func (b *DATA_BLOB) ToByteArray() []byte {
	return unsafe.Slice(b.pData, b.cbData)
}

func dpapi(data []byte) ([]byte, error) {
	din := newBlob(data)
	dout := newBlob(nil)
	r, _, err := procDecryptData.Call(uintptr(unsafe.Pointer(din)), 0, 0, 0, 0, 1, uintptr(unsafe.Pointer(dout)))
	if r == 0 {
		return nil, fmt.Errorf("CryptUnprotectData failed: %w", err.(syscall.Errno))
	}
	return dout.ToByteArray(), nil
}
