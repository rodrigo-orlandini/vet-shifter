package utils

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

type argonParams struct {
	memory      uint32
	iterations  uint32
	parallelism uint8
	saltLen     uint32
	keyLen      uint32
}

func Argon2Hash(raw string) string {
	p := argonParams{
		memory:      uint32(64 * 1024),
		iterations:  uint32(3),
		parallelism: uint8(2),
		saltLen:     uint32(16),
		keyLen:      uint32(32),
	}

	salt := make([]byte, p.saltLen)
	rand.Read(salt)

	hash := argon2.IDKey([]byte(raw), salt, p.iterations, p.memory, p.parallelism, p.keyLen)

	encodedSalt := base64.RawStdEncoding.EncodeToString(salt)
	encodedHash := base64.RawStdEncoding.EncodeToString(hash)

	encoded := fmt.Sprintf(
		"$argon2$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s",
		argon2.Version,
		p.memory,
		p.iterations,
		p.parallelism,
		encodedSalt,
		encodedHash,
	)

	return encoded
}

func Argon2Compare(raw string, hash string) bool {
	parts := strings.Split(hash, "$")

	var memory, iterations uint32
	var parallelism uint8

	fmt.Sscanf(parts[4], "m=%d,t=%d,p=%d", &memory, &iterations, &parallelism)

	decodedSalt, err := base64.RawStdEncoding.DecodeString(parts[5])
	if err != nil {
		return false
	}

	decodedHash, err := base64.RawStdEncoding.DecodeString(parts[6])
	if err != nil {
		return false
	}

	keyLen := uint32(len(decodedHash))
	comparisonHash := argon2.IDKey([]byte(raw), decodedSalt, iterations, memory, parallelism, keyLen)

	return subtle.ConstantTimeCompare(decodedHash, comparisonHash) == 1
}
