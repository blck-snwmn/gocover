package main

import (
	"encoding/hex"
	"fmt"

	"github.com/blck-snwmn/gocover/crypto"
	"github.com/blck-snwmn/gocover/decrypt"
)

type P struct {
	crypto crypto.P
	// coverage 100% if blck-snwmn/gocover/decrypt doesn't exist
	decrypt decrypt.P
}

func (p P) Crypto(b []byte) []byte {
	return p.crypto.Crypto(b)
}

func (p P) Decrypt(b []byte) []byte {
	return p.decrypt.Decrypt(b)
}

func main() {
	p := P{}
	c := p.Crypto([]byte("64787530818579737b3072827f877e30767f88307a857d8083307f8675823071307c718a8930747f773e"))
	fmt.Println(hex.EncodeToString(c))
}
