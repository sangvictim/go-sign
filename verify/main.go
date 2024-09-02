package main

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

func verifySignature(publicKey *rsa.PublicKey, data, signature []byte) error {
	// Membuat hash dari data
	hashed := sha256.Sum256(data)

	// Verifikasi tanda tangan dengan kunci publik
	err := rsa.VerifyPSS(publicKey, crypto.SHA256, hashed[:], signature, nil)
	if err != nil {
		return err
	}

	return nil
}

func loadPublicKey(fileName string) (*rsa.PublicKey, error) {
	keyFile, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer keyFile.Close()

	fileInfo, _ := keyFile.Stat()
	size := fileInfo.Size()
	keyPEM := make([]byte, size)
	keyFile.Read(keyPEM)

	block, _ := pem.Decode(keyPEM)
	pubKey, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return pubKey, nil
}

func main() {
	data := []byte("Ini adalah data yang akan ditandatangani.")

	// Memuat kunci publik dari file
	publicKey, err := loadPublicKey("public_key.pem")
	if err != nil {
		fmt.Println("Gagal memuat kunci publik:", err)
		os.Exit(1)
	}

	// Memuat tanda tangan dari file
	signature, err := os.ReadFile("signature.dat")
	if err != nil {
		fmt.Println("Gagal memuat tanda tangan:", err)
		os.Exit(1)
	}

	// Verifikasi tanda tangan
	err = verifySignature(publicKey, data, signature)
	if err != nil {
		fmt.Println("Tanda tangan tidak valid:", err)
	} else {
		fmt.Println("Tanda tangan valid.")
	}
}
