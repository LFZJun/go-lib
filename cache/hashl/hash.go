package hashl

import (
	"github.com/dgryski/go-farm"
)

func HashIndex32(key string, size uint32) uint32 {
	return farm.Hash32([]byte(key)) % size
}
