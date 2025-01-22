package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

// GenerateRandomString generates a random string of the specified length
// using characters from the given charset.
func GenerateRandomString(length int, charset string) (string, error) {
	if length <= 0 {
		return "", fmt.Errorf("invalid length: %d", length)
	}

	var result string
	charsetLength := big.NewInt(int64(len(charset)))

	for i := 0; i < length; i++ {
		index, err := rand.Int(rand.Reader, charsetLength)
		if err != nil {
			return "", err
		}
		result += string(charset[index.Int64()])
	}

	return result, nil
}

func main() {
	// Define the character set for the password
	charset := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!@#$%^&*()-_=+[]{}<>"
	
	// Desired password length
	passwordLength := 13

	// Generate the password
	password, err := GenerateRandomString(passwordLength, charset)
	if err != nil {
		fmt.Println("Error generating password:", err)
		return
	}

	fmt.Println("Generated Password:", password)
}
