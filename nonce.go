package box

import (
	"golang.org/x/crypto/blake2b"
)

const nonceLength = 24

func generateNonce(pk1 *[publicKeyLength]byte, pk2 *[publicKeyLength]byte) (*[nonceLength]byte, error) {
	hasher, err := blake2b.New(nonceLength, nil)

	if err != nil {
		return nil, err
	}

	hasher.Write(pk1[:])
	hasher.Write(pk2[:])

	hash := hasher.Sum(nil)
	var out [nonceLength]byte
	copy(out[:], hash)

	return &out, nil
}
