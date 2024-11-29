package auth

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
	"log"
)

func decryptAES256(key []byte, iv []byte, encrypted string) (string, error) {
	enc, err := base64.StdEncoding.DecodeString(encrypted)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	if len(enc)%aes.BlockSize != 0 {
		return "", errors.New("ciphertext is not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(enc, enc)

	// 去除填充字节
	enc = pkcs7Unpad(enc)
	if enc == nil {
		return "", errors.New("padding is invalid")
	}

	return string(enc), nil
}

// pkcs7Unpad 去除PKCS#7填充
func pkcs7Unpad(data []byte) []byte {
	length := len(data)
	if length == 0 {
		return nil
	}
	paddingLen := int(data[length-1])
	if paddingLen > length || paddingLen > aes.BlockSize {
		return nil
	}
	return data[:length-paddingLen]
}

func DecryptSK(systemSk, userSk string) (string, error) {
	if systemSk == userSk {
		return systemSk, nil
	}
	key := []byte(systemSk)          // 密钥替换为各产品对应的系统级sk
	iv := []byte("1234567890123456") // 16字节IV  保持不变
	encrypted := userSk              // 加密内容替换为接口返回的加密的用户级SK

	decrypted, err := decryptAES256(key, iv, encrypted)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	log.Println("Decrypted:", decrypted)
	return decrypted, nil
}
