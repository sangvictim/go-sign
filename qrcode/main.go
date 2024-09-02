package main

import (
	"fmt"

	"github.com/skip2/go-qrcode"
)

func main() {

	png, err := qrcode.New("Hello, World!", qrcode.Medium)

	if err != nil {
		panic(err)
	}

	png.WriteFile(256, "qrcode.png")
	fmt.Println("QR code created successfully!")
}
