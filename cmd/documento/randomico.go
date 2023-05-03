package documento

import (
	"hash/maphash"
	"math/rand"
)

func GetRandomico() *rand.Rand {
	return rand.New(rand.NewSource(int64(new(maphash.Hash).Sum64())))
}
