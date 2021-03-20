package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
)

func hexStringToBase64String(hexString string) string {
	bytes, err := hex.DecodeString(hexString)
	if err != nil {
		log.Fatal(err)
	}
	base64String := base64.StdEncoding.EncodeToString(bytes)
	return base64String
}

func main() {
	const hs string = "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	b64s := hexStringToBase64String(hs)
	fmt.Println(b64s)
}
