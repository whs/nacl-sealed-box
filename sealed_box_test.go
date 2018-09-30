package box_test

import (
	"bytes"
	"crypto/rand"
	"testing"

	sealedBox "github.com/whs/nacl-sealed-box"
	"golang.org/x/crypto/nacl/box"
)

func TestSealOpen(t *testing.T) {
	input := []byte("test message")

	pk, sk, err := box.GenerateKey(rand.Reader)
	if err != nil {
		t.Fatal(err)
	}

	boxed, err := sealedBox.Seal(input, pk)
	t.Log(boxed)
	if err != nil {
		t.Fatal(err)
	}

	opened, err := sealedBox.Open(boxed, pk, sk)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(opened, input) {
		t.Errorf("incorrect decryption. expected %v got %v", input, opened)
	}
}
