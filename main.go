package main

import (
	// "Cryptopals/Xor"
	"Cryptopals/AES"
	// "encoding/base64"
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

	// text := []byte("abc")
	// key := []byte("YELLOW SUBMARINE")
	// content := base64.StdEncoding.EncodeToString(text)

	// Encrypted , err := AES.Aes128Encryption(string(content), key)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return 
	// }
	// fmt.Println(string(base64.StdEncoding.EncodeToString(Encrypted)))
	tab := [][]byte {
		{0xd4, 0xe0, 0xb8, 0x1e},
		{0xbf, 0xb4, 0x41, 0x27},
		{0x5d, 0x52, 0x11, 0x98},
		{0x30, 0xae, 0xf1, 0xe5},
	}
	newtab := AES.MixColumns(tab)
	for i := 0; i < 4; i++ {
		for j:=0; j < 4; j++ {
			fmt.Printf("%x ,", newtab[i][j])
		}
		fmt.Printf("\n")
	}
}