package main

import "fmt"

func cipherEncrypt(offset int, input string) string {
	var result string
	var initialCharacter int = int('a') - 1
	var lastCharacter int = int('z')
	offset = offset % initialCharacter

	for _, character := range input {
		var encryptedCharacterAscii int = int(character) + offset

		if encryptedCharacterAscii > lastCharacter {
			result += string(rune(((int(character) + offset) - lastCharacter) + initialCharacter))
		} else {
			result += string(rune((int(character) + offset)))
		}
	}

	return result
}

func main() {
	fmt.Println(cipherEncrypt(10, "alterra"))
}
