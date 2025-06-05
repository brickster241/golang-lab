package learnlab

import (
	"crypto/rand"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"io"
)

func hashFunctions() {
	pwd := "password123"
	salt, err := generateSalt()
	if err != nil {
		fmt.Println(err)
		return
	}
	hashSaltedPwd := hashPassword(pwd, salt) 
	saltEncoded := base64.StdEncoding.EncodeToString(salt)

	// Store the encoded salt and hash in database
	hash256 := sha256.Sum256([]byte(hashSaltedPwd))
	hash512 := sha512.Sum512([]byte(hashSaltedPwd))
	fmt.Println("Pwd  :", pwd)
	fmt.Println("Salt :", salt)
	fmt.Println("Encoded Salt :", saltEncoded)
	fmt.Println("Hashed Password (with Salt) :", hashSaltedPwd)
	fmt.Println("\nHash 256 :", hash256)
	fmt.Printf("SHA-256 Hash Hex val: %x\n\n", hash256)
	fmt.Println("\nHash 512 :", hash512)
	fmt.Printf("SHA-512 Hash Hex val: %x\n\n", hash512)
}

func generateSalt() ([]byte, error) {
	salt := make([]byte, 16)
	_, err := io.ReadFull(rand.Reader, salt)
	if err != nil {
		fmt.Println("Error Generating salt :", err)
		return nil, err
	}
	return salt, nil
}

func hashPassword(pwd string, salt []byte) string {
	saltedPwd := append(salt, []byte(pwd)...)
	hashedSaltedPwd := sha256.Sum256(saltedPwd)
	return base64.StdEncoding.EncodeToString(hashedSaltedPwd[:])
}
