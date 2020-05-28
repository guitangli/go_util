package secure

import (
	"crypto/aes"
	"crypto/cipher"
	"errors"
	"fmt"
)

var (
	encryptCipher cipher.BlockMode
	decryptCipher cipher.BlockMode
	blockSize int

	blockSizeErr = errors.New("plaintext not divisible")
)

type AESCBCInterface interface {
	cipher.BlockMode
	SetIV([]byte)
}

func InitAESCBC(key []byte) error {
	block1, err := aes.NewCipher(key)
	if err != nil {
		return err
	}
	block2, err := aes.NewCipher(key)
	if err != nil {
		return err
	}
	blockSize = block1.BlockSize()
	fmt.Printf("InitAESCBC BlockSize=%v\n", blockSize)
	// init empty iv
	iv := make([]byte, blockSize)
	decryptCipher = cipher.NewCBCDecrypter(block1, iv)
	encryptCipher = cipher.NewCBCEncrypter(block2, iv)
	return nil
}

// 内部不会增加padding字段。
func AESCBCEncryptNoPadding(plaintext []byte) ([]byte, error) {
	if len(plaintext) % blockSize != 0 {
		return nil, blockSizeErr
	}
	dst := make([]byte, len(plaintext))
	encryptCipher.CryptBlocks(dst, plaintext)
	return dst, nil
}

// 不是整的blocksize时内部会增加padding
func AESCBCDecryptNoPadding(ciphertext []byte) ([]byte, error) {
	if len(ciphertext) % blockSize != 0 {
		return nil, blockSizeErr
	}
	dst := make([]byte, len(ciphertext))
	decryptCipher.CryptBlocks(dst, ciphertext)
	return dst, nil
}

func AESCBCEncryptSetIV(iv []byte)  {
	cipher, ok := encryptCipher.(AESCBCInterface)
	if !ok {
		panic("interface")
	}
	cipher.SetIV(iv)
}

func AESCBCDecryptSetIV(iv []byte)  {
	cipher, ok := decryptCipher.(AESCBCInterface)
	if !ok {
		panic("interface")
	}
	cipher.SetIV(iv)
}

func AESCBCBlockSize() int {
	return blockSize
}
