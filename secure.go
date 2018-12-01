package main

import (
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"strings"
)

// HashingAlgo type of hashing algorithm
type HashingAlgo string

const (
	// SHA256 hashing algorithm
	SHA256 HashingAlgo = "SHA256"
	// SHA512 hashing algorithm
	SHA512 HashingAlgo = "SHA512"
)

// HashWithSha512 hash the input string with sha256
func hashWithSha512(input string) (string, error) {
	hasher := sha512.New()
	hasher.Write([]byte(input))
	return hex.EncodeToString(hasher.Sum(nil)), nil
}

// Hash the input string with the given algorithm
func Hash(input string, algo HashingAlgo) (string, error) {
	if len(strings.TrimSpace(input)) == 0 {
		return "", fmt.Errorf("can not hash empty string")
	}
	switch algo {
	case SHA512:
		return hashWithSha512(input)
	default:
		return "", fmt.Errorf("hashing algorith[%s] not supported/valid", algo)
	}
}

// MatchHash compare hashed  value of given input string with the given hashed value
func MatchHash(input, hashedVal string, algo HashingAlgo) bool {
	if hashedInput, err := Hash(input, algo); err == nil && hashedInput == hashedVal {
		return true
	}
	return false
}
