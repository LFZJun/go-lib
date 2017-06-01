package hashl

import (
	"unsafe"

	"github.com/dgryski/go-farm"
)

func HashIndex32(key string, size uint32) uint32 {
	return farm.Hash32(str2bytes(key)) % size
}

func str2bytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}
