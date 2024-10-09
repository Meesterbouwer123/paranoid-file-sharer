package fileencryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

var aes_gcm cipher.AEAD

func Init() {
	key := []byte(os.Getenv("ENCRYPTION_KEY"))
	encryptor, err := aes.NewCipher(key)
	if err != nil {
		log.Fatal(err)
	}

	aes_gcm, err = cipher.NewGCM(encryptor)
	if err != nil {
		log.Fatal(err)
	}
}

// Generate a random filename so the uploaded filename isn't visible on disk
func GetFilename() string {
	// 12 random bytes conterted to hex characters will probalby be enough
	bytes := make([]byte, 12)
	rand.Read(bytes)
	return hex.EncodeToString(bytes)
}

func IsValidFilename(filename string) bool {
	bytes, err := hex.DecodeString(filename)
	return err == nil && len(bytes) == 12
}

func EncryptFile(filename string, mime_type string, contents []byte) []byte {
	// pack the data into a single array, and encrypt that one
	packed := packFile(filename, mime_type, contents)

	// encrypt the packed data
	nonce := make([]byte, aes_gcm.NonceSize())
	_, err := rand.Read(nonce)
	if err != nil {
		log.Fatal(err)
	}
	cipthertext := aes_gcm.Seal(nonce, nonce, packed, nil)

	return cipthertext
}

func DecryptFile(data []byte) (filename, mime_type string, contents []byte, err error) {
	// decrypt the data and unpack to the wanted values
	nonce_size := aes_gcm.NonceSize()
	nonce, ciphertext := data[:nonce_size], data[nonce_size:]
	plaintext, err := aes_gcm.Open(nonce, nonce, ciphertext, nil)
	if err != nil {
		return "", "", make([]byte, 0), err
	}

	filename, mime_type, contents, err = unpackFile(plaintext[nonce_size:])
	if err != nil {
		return "", "", make([]byte, 0), err
	}

	return filename, mime_type, contents, nil
}

func packFile(filename, mime_type string, contents []byte) []byte {
	metadata := fmt.Sprintf("%s%%%s|", filename, mime_type)
	result := append([]byte(metadata), contents...)

	return result
}

func unpackFile(data []byte) (filename, mime_type string, contents []byte, err error) {
	seperator_index := slices.Index(data, '|')
	if seperator_index == -1 {
		return "", "", make([]byte, 0), fmt.Errorf("data doesn't contain metadata")
	}
	metadata, data := data[:seperator_index], data[seperator_index+1:]
	metadata_str := string(metadata)

	split := strings.Split(metadata_str, "%")
	if len(split) != 2 {
		return "", "", make([]byte, 0), fmt.Errorf("data doesn't contain metadata")
	}

	return split[0], split[1], data, nil
}
