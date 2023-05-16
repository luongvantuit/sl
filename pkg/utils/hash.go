package utils

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
)

// Hash Id to Key
func HashId(id string) (*string, error) {
	var rawId []byte
	copy(rawId, []byte(id))
	sha := sha256.New()
	sha.Write(rawId)
	// Hash Id
	var hashId string = fmt.Sprintf("%x", sha.Sum(nil))
	return &hashId, nil
}

func HashSysId(id string) (*string, error) {
	var rawId []byte
	copy(rawId, []byte(id))
	sha := sha512.New()
	sha.Write(rawId)
	var hashSysId string = fmt.Sprintf("%x", sha.Sum(nil))
	return &hashSysId, nil
}
