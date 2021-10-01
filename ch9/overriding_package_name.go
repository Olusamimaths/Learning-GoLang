package main

import (
	crand "crypto/rand"
	"encoding/binary"
	"math/rand"
)

func seedRand() *rand.Rand {
	var b [8]byte
	_, err := crand.Read(b[:])
	if err != nil {
		panic("cannot seed with crytographic random number generator")
	}
	r := rand.New(rand.NewSource(int64(binary.LittleEndian.Uint64(b[:]))))
	return r;
}

func main() {
	
}