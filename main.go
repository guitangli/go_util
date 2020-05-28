package main

import (
	"crypto/rand"
	"fmt"
	"github.com/guitangli/go_util/secure"
)

func main()  {
	key := "12345678901234567890123456789012"
	err := secure.InitAESCBC([]byte(key))
	if err != nil {
		panic(err)
	}
	blockSize := secure.AESCBCBlockSize()
	iv := make([]byte, blockSize)
	_, err = rand.Read(iv)
	if err != nil {
		panic(err)
	}
	//secure.AESCBCEncryptSetIV(iv)
	//secure.AESCBCDecryptSetIV(iv)
	fmt.Printf("iv=%v\n", iv)

	plaintext := "1234567890123456"

	ciphertext, err := secure.AESCBCEncryptNoPadding([]byte(plaintext))
	if err != nil {
		panic(err)
	}
	fmt.Printf("ciphertext=%v", ciphertext)

	compare, err := secure.AESCBCDecryptNoPadding(ciphertext)
	if err != nil {
		panic(err)
	}
	fmt.Printf("compare=%v", string(compare))
	if string(compare) != plaintext {
		panic("diff")
	}
}
