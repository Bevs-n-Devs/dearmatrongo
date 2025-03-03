package encrypt

import (
	"crypto/sha256"
	"encoding/hex"
)

/*
HashData takes a string and returns a SHA-256 hash of the string.

The resulting hash is a fixed-size 256-bit string, represented as
a 64-character hexadecimal string.
*/
func HashData(data string) string {
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}

/*
VerifyHash takes an input string and a stored hash string and
returns true if the two hashes match, or false if they do not.

This is a simple equality check and does not provide any
additional security features. It is the responsibility of the
caller to ensure that the input string and stored hash are valid
and have been secured appropriately.
*/
func VerifyHash(input string, storedHash string) bool {
	return input == storedHash // Compare with the stored hash
}
