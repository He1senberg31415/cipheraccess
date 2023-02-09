package main

import "fmt"

func main() {
	fmt.Println(genKey())
	fmt.Println()
	kp := genKeyPair()
	fmt.Println(kp)
	fmt.Println()
	encryptedData := encryptData([]byte("test"), kp.PublicKey)
	fmt.Println(encryptedData)
	fmt.Println(string(decryptData(encryptedData, kp)))
}

