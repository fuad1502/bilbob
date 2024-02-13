package passwords

import (
	"crypto/rand"
	"encoding/hex"
	"golang.org/x/crypto/argon2"
	"math/big"
)

const SaltSize = 16
const KeySize = 16

func GenerateSalt() ([]byte, error) {
	salt := make([]byte, SaltSize)
	for i := 0; i < int(SaltSize); i += 1 {
		t, err := rand.Int(rand.Reader, big.NewInt(256))
		if err != nil {
			return nil, err
		}
		salt[i] = byte(t.Int64())
	}
	return salt, nil
}

func HashPassword(password string, salt []byte) string {
	hashString := hex.EncodeToString(argon2.IDKey([]byte(password), salt, 2, 15*1024, 1, KeySize))
	return hashString
}
