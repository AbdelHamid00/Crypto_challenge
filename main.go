package main

import (
	// "Cryptopals/Xor"
	"Cryptopals/AES"
	"encoding/base64"
	// "io/ioutil"
	// // "os"
	// "log"
	"fmt"
	// // "bufio"
)

func main() {
	// filePath := "./InputFiles/input7.txt"
	// content, err := ioutil.ReadFile(filePath)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// The Keys should be encoded in base64

	text := []byte("abc")
	key := []byte("YELLOW SUBMARINE")
	content := base64.StdEncoding.EncodeToString(text)

	Encrypted , err := AES.Aes128Encryption(string(content), key)
	if err != nil {
		fmt.Println(err)
		return 
	}
	fmt.Println(string(base64.StdEncoding.EncodeToString(Encrypted)))
	// fmt.Printf("File contents:\n%s\n", string(content))
	// AES.Aes128Encryption(content, key)
	// input := string("TslNiU11yhtK0ki2nZ2TkzUuCwrzQ8uEfcVlDdvmNrvru0qwqCsLBNBwbn9gezAy4FQ/j16eeFbWJDxE5wgT3PQ7h0TheN5uBtyN+inxXlNqChF+AhNAfTcRtxe74wpKpn3dOlXhQMerq/swrtsa66m7vWwezOkf+Dy7Ezs0vexhPUdjw9PKcK4F3+hUQb0G0P91vx0tUiF6itmYAGLjiw==")
	// key := string("A4+TGuUSoMqGhyxPgx4fJg==")
	// Plaintext, err := base64.StdEncoding.DecodeString(input)
	// if err != nil {
	// 	panic(err)
	// }
	// PlainKey, errr := base64.StdEncoding.DecodeString(key)
	// if errr != nil {
	// 	panic(err)
	// }
	
	// fmt.Println(string(Plaintext))
	// result := Xor.RepeatingKeyXOR(Plaintext, PlainKey)
	// fmt.Println(string(result))
}