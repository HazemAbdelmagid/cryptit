package main

import (
	"fmt"
	"github.com/HazemAbdelmagid/cryptit/encrypt"
	"github.com/HazemAbdelmagid/cryptit/decrypt"
)

func main() {
	encryptedStr := encrypt.Nimbus("KodeKloud")
	fmt.Println(encryptedStr)
	fmt.Println(decrypt.Nimbus(encryptedStr))
}
