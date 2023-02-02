package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"flag"
	"fmt"
	"io"
	"os"
)

var helpText = `
Example:
  ./ezAES                     # interactively
  ./ezAES -t hello            # encrypt
  ./ezAES -d -t [long string] # decrypt
  ./ezAES -t hello -k 123     # not safe
`

func main() {
	helpFlag := flag.Bool("h", false, "Show help")
	typeFlag := flag.Bool("d", false, "Decrypt")
	key := flag.String("k", "", "Pass key from commoand line")
	txt := flag.String("t", "", "Pass txt from commoand line")
	flag.Parse()
	if *helpFlag {
		flag.Usage()
		fmt.Println(helpText)
		os.Exit(0)
	}
	if *txt == "" {
		fmt.Println("Input your txt:")
		fmt.Scanln(txt)
	}
	if *key == "" {
		fmt.Println("Input your key:")
		fmt.Scanln(key)
	}

	var ans string
	var err error
	if *typeFlag {
		ans, err = DecryptAES(*key, *txt)
	} else {
		ans, err = EncryptAES(*key, *txt)
	}
	if err != nil {
		panic(err)
	}

	fmt.Println(ans)
}

func DecryptAES(password string, crypt64 string) (string, error) {
	if crypt64 == "" {
		return "", nil
	}

	key := make([]byte, 32)
	copy(key, []byte(password))

	crypt, err := base64.StdEncoding.DecodeString(crypt64)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	iv := crypt[:aes.BlockSize]
	crypt = crypt[aes.BlockSize:]
	decrypted := make([]byte, len(crypt))
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(decrypted, crypt)

	return string(decrypted[:len(decrypted)-int(decrypted[len(decrypted)-1])]), nil
}

func EncryptAES(password string, plaintext string) (string, error) {
	if plaintext == "" {
		return "", nil
	}

	key := make([]byte, 32)
	copy(key, []byte(password))
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	content := []byte(plaintext)
	blockSize := block.BlockSize()
	padding := blockSize - len(content)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	content = append(content, padtext...)

	ciphertext := make([]byte, aes.BlockSize+len(content))

	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], content)

	return base64.StdEncoding.EncodeToString(ciphertext), nil
}
