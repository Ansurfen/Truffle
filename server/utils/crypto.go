package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

const (
	FrontKey  = "It's front  key."
	BackedKey = "It's backed key."
)

func HashPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	AssertWithPanic(err)
	return string(hash)
}

func EqualHashPassword(raw, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(raw))
	return err
}

func MD5(str string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}

func EqualMD5(raw, hash string) bool {
	return MD5(raw) == hash
}

func SHA256(str string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(str)))
}

func EqualSHA256(raw, hash string) bool {
	return SHA256(raw) == hash
}

func EncodeAESWithKey(key, str string) string {
	hash := md5.New()
	hash.Write([]byte(key))
	keyData := hash.Sum(nil)
	block, err := aes.NewCipher(keyData)
	AssertWithPanic(err)
	iv := []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}
	enc := cipher.NewCBCEncrypter(block, iv)
	content := _PKCS5Padding([]byte(str), block.BlockSize())
	crypted := make([]byte, len(content))
	enc.CryptBlocks(crypted, content)
	return base64.StdEncoding.EncodeToString(crypted)
}

func DecodeAESWithKey(key, str string) string {
	hash := md5.New()
	hash.Write([]byte(key))
	keyData := hash.Sum(nil)
	block, err := aes.NewCipher([]byte(keyData))
	AssertWithPanic(err)
	iv := []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}
	messageData, _ := base64.StdEncoding.DecodeString(str)
	dec := cipher.NewCBCDecrypter(block, iv)
	decrypted := make([]byte, len(messageData))
	dec.CryptBlocks(decrypted, messageData)
	return string(_PKCS5Unpadding(decrypted))
}

func EqualAES(key, raw, hash string) bool {
	return EncodeAESWithKey(key, raw) == hash
}

func _PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func _PKCS5Unpadding(encrypt []byte) []byte {
	padding := encrypt[len(encrypt)-1]
	return encrypt[:len(encrypt)-int(padding)]
}
