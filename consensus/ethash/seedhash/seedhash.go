package seedhash

import (
	"hash"
	"golang.org/x/crypto/sha3"
)
/*
	This bs file is needed to resolve cyclic dependency ethapi on ethash, 
	thus enabling us to work with ethapi in ethhash 
*/

const EpochLength = 30000
func SeedHash(block uint64) []byte {
	seed := make([]byte, 32)
	if block < EpochLength {
		return seed
	}
	keccak256 := MakeHasher(sha3.NewLegacyKeccak256())
	for i := 0; i < int(block/EpochLength); i++ {
		keccak256(seed, seed)
	}
	return seed
}

// hasher is a repetitive hasher allowing the same hash data structures to be
// reused between hash runs instead of requiring new ones to be created.
type Hasher func(dest []byte, data []byte)

// makeHasher creates a repetitive hasher, allowing the same hash data structures to
// be reused between hash runs instead of requiring new ones to be created. The returned
// function is not thread safe!
func MakeHasher(h hash.Hash) Hasher {
	// sha3.state supports Read to get the sum, use it to avoid the overhead of Sum.
	// Read alters the state but we reset the hash before every operation.
	type readerHash interface {
		hash.Hash
		Read([]byte) (int, error)
	}
	rh, ok := h.(readerHash)
	if !ok {
		panic("can't find Read method on hash")
	}
	outputLen := rh.Size()
	return func(dest []byte, data []byte) {
		rh.Reset()
		rh.Write(data)
		rh.Read(dest[:outputLen])
	}
}
