package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

func generateKeyPair(bitSize int) (*rsa.PrivateKey, *rsa.PublicKey) {
	// membuat private key
	privateKey, err := rsa.GenerateKey(rand.Reader, bitSize)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	//menghasilkan public key dari private key
	publicKey := &privateKey.PublicKey

	return privateKey, publicKey

}

func savePEMKey(fileName string, key *rsa.PrivateKey) {
	outFile, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error creating file:", err)
		os.Exit(1)
	}

	defer outFile.Close()

	// Menyimpan kunci privat dalam format PEM
	var privateKey = &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(key),
	}

	pem.Encode(outFile, privateKey)
}

func savePublicPEMKey(fileName string, key *rsa.PublicKey) {
	outFile, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error creating file:", err)
		os.Exit(1)
	}

	defer outFile.Close()

	// Menyimpan kunci publik dalam format PEM
	var publicKey = &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: x509.MarshalPKCS1PublicKey(key),
	}

	pem.Encode(outFile, publicKey)
}
func main() {
	bitSize := 2048

	privateKey, publicKey := generateKeyPair(bitSize)

	// Menyimpan kunci privat dan publik ke file
	savePEMKey("private_key.pem", privateKey)
	savePublicPEMKey("public_key.pem", publicKey)

	fmt.Println("Kunci RSA berhasil dibuat dan disimpan.")
}
