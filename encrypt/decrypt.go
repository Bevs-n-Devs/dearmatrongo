package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"

	"github.com/Bevs-n-Devs/dearmatrongo/logs"
)

/*
Decrypt decrypts the given encrypted data using AES-GCM with the master key.
It expects the data to contain the nonce followed by the ciphertext.

Parameters:

	data ([]byte): The encrypted data containing the nonce and ciphertext.

Returns:

	([]byte): The decrypted plaintext if successful.
	(error): An error if the decryption process fails, such as an invalid key or corrupted data.
*/
func Decrypt(data []byte) ([]byte, error) {
	// Create a new AES cipher block using the same master key
	block, err := aes.NewCipher(masterKey)
	if err != nil {
		logs.Logs(logError, fmt.Sprintf("Error creating AES cipher block: %s", err.Error()))
		return nil, err // Return error if key is invalid
	}

	// Create a GCM cipher from the AES block
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		logs.Logs(logError, fmt.Sprintf("Error creating GCM cipher: %s", err.Error()))
		return nil, err // Return error if GCM initialization fails
	}

	// Extract the nonce from the start of the encrypted data
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]

	// Decrypt the ciphertext using AES-GCM
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		logs.Logs(logError, fmt.Sprintf("Error decrypting data: %s", err.Error()))
		return nil, err // Return error if decryption fails
	}

	// Return the decrypted plaintext
	logs.Logs(logInfo, "Data decrypted successfully")
	return plaintext, nil
}
