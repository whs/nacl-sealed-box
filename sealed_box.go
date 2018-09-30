package box

import (
	"crypto/rand"
	"errors"

	"golang.org/x/crypto/nacl/box"
)

const publicKeyLength = 32
const privateKeyLength = 32

// Seal encrypt and sign a message with the given pk and ephemeral sk
func Seal(message []byte, pk *[publicKeyLength]byte) ([]byte, error) {
	out := make([]byte, publicKeyLength, len(message)+box.Overhead+publicKeyLength)

	ePk, eSk, err := box.GenerateKey(rand.Reader)

	if err != nil {
		return nil, err
	}

	copy(out, ePk[:])
	nonce, err := generateNonce(ePk, pk)

	if err != nil {
		return nil, err
	}

	out = box.Seal(out, message, nonce, pk, eSk)

	// XXX: ePk, eSk, nonce are not zeroed

	return out, nil
}

// Open decrypt and verify a sealed message
func Open(ciphertext []byte, pk *[publicKeyLength]byte, sk *[privateKeyLength]byte) ([]byte, error) {
	var ePk [publicKeyLength]byte
	copy(ePk[:], ciphertext)

	nonce, err := generateNonce(&ePk, pk)

	if err != nil {
		return nil, err
	}

	out := make([]byte, 0, len(ciphertext)-box.Overhead-publicKeyLength)

	out, opened := box.Open(out, ciphertext[publicKeyLength:], nonce, &ePk, sk)

	if !opened {
		return nil, errors.New("sealed box opening fail")
	}

	return out, nil
}
