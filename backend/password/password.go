package password

import (
	"crypto/rand"
	"encoding/hex"
	"golang.org/x/crypto/argon2"
	"math/big"
)

func getSalt(nBytes uint32) ([]byte, error) {
	salt := make([]byte, nBytes)
	for i := 0; i < int(nBytes); i += 1 {
		t, err := rand.Int(rand.Reader, big.NewInt(256))
		if err != nil {
			return nil, err
		}
		salt[i] = byte(t.Int64())
	}
	return salt, nil
}

func HashPassword(password string) (string, error) {
	salt, err := getSalt(16)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(argon2.IDKey([]byte(password), salt, 2, 15*1024, 1, 16)), nil
}
