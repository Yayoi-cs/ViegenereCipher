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
	key := "agfl"
	cipherText := "pohzCZK{m311a50_0x_a1rn3x3_h1ah3x6kp60egf}"
	plainText := decryptVigenere(cipherText, key)
	fmt.Println("Result :" + plainText)
}
