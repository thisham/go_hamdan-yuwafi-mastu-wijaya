package main

import (
	"fmt"
	"os"
)

type Student struct {
	name   string
	offset int
}

type Cipher interface {
	encode() string
	decode() string
}

const (
	initialCharacter int = int('a') - 1
	lastCharacter    int = int('z')
	alphabeticOffset     = lastCharacter - initialCharacter
	offset               = 10
)

var checkIsAlphabet = func(encodedAscii int) bool {
	return encodedAscii > initialCharacter && encodedAscii <= lastCharacter
}

func (data Student) encode() string {
	var encoded string

	for _, character := range data.name {
		// skip encoding non-alphabetic character
		if !checkIsAlphabet(int(character)) {
			encoded += string(character)
			continue
		}

		encodedAscii := int(character) + data.offset
		if encodedAscii > lastCharacter {
			encoded += string(rune(encodedAscii - alphabeticOffset))
			continue
		}
		encoded += string(rune(encodedAscii))
	}

	return encoded
}

func (data Student) decode() string {
	var decoded string

	for _, character := range data.name {
		// skip encoding non-alphabetic character
		if !checkIsAlphabet(int(character)) {
			decoded += string(character)
			continue
		}

		decodedAscii := int(character) - data.offset
		if decodedAscii < initialCharacter {
			decoded += string(rune(decodedAscii + alphabeticOffset))
			continue
		}
		decoded += string(rune(decodedAscii))
	}

	return decoded
}

func main() {
	var student = Student{}
	var studentCipher Cipher = &student

	if len(os.Args) < 3 {
		fmt.Println("incomplete command!")
		fmt.Println("complete commands are: starting with command followed by 'encrypt' or 'decrypt' subcommand, then place your strings to encrypt or decrypt.")
		return
	}

	// getting number three params as will be encoded string
	student.name = os.Args[2]
	student.offset = offset

	switch os.Args[1] {
	case "encrypt":
		fmt.Println("encoded student name:", studentCipher.encode())
		break
	case "decrypt":
		fmt.Println("decrypted encrypted student name:", studentCipher.decode())
		break
	default:
		fmt.Println("subcommand not found!")
	}
}
