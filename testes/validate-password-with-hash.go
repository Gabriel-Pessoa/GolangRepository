package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"

	"golang.org/x/crypto/pbkdf2"
)

func main() {
	password := "lose size pods pies oven daub"
	passwordHash := "1000:c8e1331ec7765cfb408dcf9698f24609:d6f8debdebe090d1d90d72457d36170b77ed7a07c4658dd7a09069aabe9a09ab1501470db0d7ec6123fbdf5952d5e5d4f148bf6cf598fc237b78e12150810bc0"
	x, err := comparePasswordAndHash(password, passwordHash)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(x)
}

func comparePasswordAndHash(password, passwordHash string) (bool, error) {
	parts := strings.Split(passwordHash, ":")

	iteration, err := strconv.Atoi(parts[0])
	if err != nil {
		return false, err
	}

	salt, err := fromHexToByte(parts[1])
	if err != nil {
		return false, err
	}

	hash, err := fromHexToByte(parts[2])
	if err != nil {
		return false, err
	}

	testHash := pbkdf2.Key([]byte(password), salt, iteration, 64, sha1.New) // pattern SHA-1

	diff := byte(len(hash) ^ len(testHash)) // XOR operator

	for i := 0; i < len(hash) && i < len(testHash); i++ {
		diff ^= hash[i] ^ testHash[i] // bitwise XOR assignment
	}

	return diff == 0, nil
}

func fromHexToByte(s string) ([]byte, error) {
	bytes, err := hex.DecodeString(s) // decoding to ASCII hex
	if err != nil {
		return nil, err
	}
	return bytes, nil
}
