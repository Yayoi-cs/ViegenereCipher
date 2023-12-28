package main

import (
	"fmt"
)

func decryptVigenere(cipherText, key string) string {
	var decrypted []rune
	for i, char := range cipherText {
		keyIndex := i % len(key)
		keyChar := rune(key[keyIndex])
		if char >= 'A' && char <= 'Z' {
			if keyChar >= 'a' && keyChar <= 'z' {
				keyChar = keyChar - 'a' + 'A'
			}
			decryptedChar := 'A' + (char-keyChar+26)%26
			decrypted = append(decrypted, decryptedChar)
		} else if char >= 'a' && char <= 'z' {
			if keyChar >= 'A' && keyChar <= 'Z' {
				keyChar = keyChar - 'A' + 'a'
			}
			decryptedChar := 'a' + (char-keyChar+26)%26
			decrypted = append(decrypted, decryptedChar)
		} else {
			decrypted = append(decrypted, char)
		}
	}
	return string(decrypted)
}
func main() {
	key := "CYLAB"
	cipherText := "rgnoDVD{O0NU_WQ3_G1G3O3T3_A1AH3S_f85729e7}"
	plainText := decryptVigenere(cipherText, key)
	fmt.Println("Result :" + plainText)
}
