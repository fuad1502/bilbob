package passwords

import (
	"crypto/rand"
	"encoding/hex"
	"golang.org/x/crypto/argon2"
	"math/big"
)

const saltSize = 16
const keySize = 16

func generateSalt() ([]byte, error) {
	salt := make([]byte, saltSize)
	for i := 0; i < int(saltSize); i += 1 {
		t, err := rand.Int(rand.Reader, big.NewInt(256))
		if err != nil {
			return nil, err
		}
		salt[i] = byte(t.Int64())
	}
	return salt, nil
}

func hashPassword(password string, salt []byte) string {
	hashString := hex.EncodeToString(argon2.IDKey([]byte(password), salt, 2, 15*1024, 1, keySize))
	return hashString
}

func GenerateSaltAndHash(password string) (string, error) {
	saltBytes, err := generateSalt()
	if err != nil {
		return "", err
	}
	hash := hashPassword(password, saltBytes)
	salt := hex.EncodeToString(saltBytes)
	return salt + hash, nil
}

func VerifyPassword(password string, saltAndHash string) (bool, error) {
	salt := saltAndHash[:saltSize*2]
	hash := saltAndHash[saltSize*2:]
	saltBytes, err := hex.DecodeString(salt)
	if err != nil {
		return false, err
	}
	computedHash := hashPassword(password, saltBytes)
	if computedHash == hash {
		return true, nil
	} else {
		return false, nil
	}
}
