package main

import (
	"crypto/aes"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
)

func main() {
	h := md5.New()
	io.WriteString(h, "Hello, world!")
	fmt.Printf("%x\n", h.Sum(nil))
	io.WriteString(h, "Hi, today is sunday.")
	fmt.Printf("%x\n", h.Sum(nil))

	sum := sha256.Sum256([]byte("hello world\n"))
	fmt.Printf("%x\n", sum)

	// cipher key
	key := "thisis32bitlongpassphraseimusini"
	// plaintext
	pt := "This is a jimmy "
	c := EncryptAES([]byte(key), pt)
	// plaintext
	fmt.Println(pt)
	// ciphertext
	fmt.Println(c)
	// decrypt
	DecryptAES([]byte(key), c)
}

func EncryptAES(key []byte, plaintext string) string {
	c, err := aes.NewCipher(key)
	CheckError(err)
	out := make([]byte, len(plaintext))
	c.Encrypt(out, []byte(plaintext))
	return hex.EncodeToString(out)
}

func DecryptAES(key []byte, ct string) {
	ciphertext, _ := hex.DecodeString(ct)
	c, err := aes.NewCipher(key)
	CheckError(err)
	pt := make([]byte, len(ciphertext))
	c.Decrypt(pt, ciphertext)
	s := string(pt[:])
	fmt.Println("DECRYPTED:", s)
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}