package main

import (
	"crypto/rand"
	"fmt"
)

// GenerateAESKey 生成指定长度的AES密钥
func GenerateAESKey(length int) ([]byte, error) {
	key := make([]byte, length/8) // AES密钥长度可以是128, 192, 或256位，对应16, 24, 或32字节
	_, err := rand.Read(key)
	if err != nil {
		return nil, fmt.Errorf("failed to generate AES key: %w", err)
	}
	return key, nil
}

// GenerateNonce 生成GCM所需的nonce
func GenerateNonce(size int) ([]byte, error) {
	nonce := make([]byte, size) // GCM nonce的标准长度为12字节
	_, err := rand.Read(nonce)
	if err != nil {
		return nil, fmt.Errorf("failed to generate nonce: %w", err)
	}
	return nonce, nil
}

func main() {
	key, err := GenerateAESKey(128) // 生成256位的AES密钥
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Generated AES Key: %x\n", key)

	nonce, err := GenerateNonce(6) // 生成12字节的GCM nonce
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Generated Nonce: %x\n", nonce)
}
