package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
	"os"
)

func genKeyPair() *rsa.PrivateKey {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return privateKey
}

func genKey() []byte {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return key
}

func encryptWithRSA(data []byte, publicKey *rsa.PublicKey) []byte {
	encryptedData, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, publicKey, data, nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return encryptedData
}

func decryptWithRSA(encryptedData []byte, privateKey *rsa.PrivateKey) []byte {
	decryptedData, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, privateKey, encryptedData, nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return decryptedData
}

func pad(src []byte) []byte {
	padding := aes.BlockSize - len(src)%aes.BlockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padtext...)
}

func unpad(src []byte) ([]byte, error) {
	length := len(src)
	unpadding := int(src[length-1])

	if unpadding > length {
		return nil, fmt.Errorf("unpad error. This could happen when incorrect encryption key is used")
	}

	return src[:(length - unpadding)], nil
}

func encryptWithKey(data []byte, key []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	data = pad(data)
	encryptedData := make([]byte, aes.BlockSize+len(data))
	iv := encryptedData[:aes.BlockSize]
	_, err = rand.Read(iv)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(encryptedData[aes.BlockSize:], data)

	return encryptedData
}

func decryptWithKey(encryptedData []byte, key []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	iv := encryptedData[:aes.BlockSize]
	encryptedData = encryptedData[aes.BlockSize:]
	decryptedData := make([]byte, len(encryptedData))
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(decryptedData, encryptedData)
	decryptedData, err = unpad(decryptedData)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return decryptedData
}

func encryptData(data []byte, publicKey rsa.PublicKey) [][]byte {
	key := genKey()
	encryptedData := encryptWithKey(data, key)
	encryptedKey := encryptWithRSA(key, &publicKey)
	return [][]byte{encryptedData, encryptedKey}
}

func decryptData(encryptedData [][]byte, privateKey *rsa.PrivateKey) []byte {
	key := decryptWithRSA(encryptedData[1], privateKey)
	decryptedData := decryptWithKey(encryptedData[0], key)
	return decryptedData
}

func hash(data []byte) string {
	return fmt.Sprintf("%x", sha256.Sum256(data))
}
