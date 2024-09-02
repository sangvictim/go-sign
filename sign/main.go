package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

func signData(privateKey *rsa.PrivateKey, data []byte) ([]byte, error) {
	hashed := sha256.Sum256(data)

	// menandatangani has dengan kunci private
	signature, err := rsa.SignPSS(rand.Reader, privateKey, crypto.SHA256, hashed[:], nil)
	if err != nil {
		return nil, err
	}

	return signature, nil
}

func loadPrivateKey(fileName string) (*rsa.PrivateKey, error) {
	keyFile, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	defer keyFile.Close()

	fileInfo, _ := keyFile.Stat()
	size := fileInfo.Size()
	keyPem := make([]byte, size)
	keyFile.Read(keyPem)

	block, _ := pem.Decode(keyPem)
	privateKey, _ := x509.ParsePKCS1PrivateKey(block.Bytes)

	return privateKey, nil
}

func main() {
	data := []byte("Ini adalah data yang akan ditandatangani.")

	// Memuat kunci privat dari file
	privateKey, err := loadPrivateKey("private_key.pem")
	if err != nil {
		fmt.Println("Gagal memuat kunci privat:", err)
		os.Exit(1)
	}

	// Menandatangani data
	signature, err := signData(privateKey, data)
	if err != nil {
		fmt.Println("Gagal menandatangani data:", err)
		os.Exit(1)
	}

	// Menyimpan tanda tangan ke file
	err = os.WriteFile("signature.sign", signature, 0644)
	if err != nil {
		fmt.Println("Gagal menyimpan tanda tangan:", err)
		os.Exit(1)
	}

	fmt.Println("Data berhasil ditandatangani.")
}
