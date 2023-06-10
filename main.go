package main

import (
	"Cryptopals/Xor"
	// "Cryptopals/AES"
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
	// key := string("YELLOW SUBMARINE")
	// fmt.Printf("File contents:\n%s\n", string(content))
	// AES.Aes128Encryption(content, key)
	input := string("ecPEJPcAsWTfb9SBuvIQxUK2vo1FsiGsDyVR82jlaFJ+N2tLrrunDyUzRDP2sP2QQ/VSgm6o7260zD04FyhpTNeY9VEciVVY+uK5tviDI6MgIv2aCTp9zDs7rHmgAQdKqZFtp2JxHzx3ftZ3hER9BdHEc+bF3Y0YeXAoz0m43YOwG6+jTX+x6ItKe7CncMQDq3IOtyr0FvIqzdFhieSDDQ==")
	Plaintext, err := base64.StdEncoding.DecodeString(string(input))
	if err != nil {
		panic(err)
	}
	fmt.Println(Plaintext)
	result := Xor.RepeatingKeyXOR(Plaintext, []byte("s"))
	fmt.Println(string(result))
}