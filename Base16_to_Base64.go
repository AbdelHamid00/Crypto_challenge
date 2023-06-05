package main

import (
	"fmt"
	"strings"
	"errors"
)
// start reading from here ^-^ .
var (
	// this two tables are the runes of the  base16 and base64 stored in a string . we gonna use after in the mapping and the replacement .
	table16 string = "0123456789ABCDEF"
	table64 string = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
	// The Base 16 mapping , rune to the binary representation of it's value stored in a string .
	// 0-9: 0-9 
	// A-F: 10-15 
	map16 = make(map[rune]string)
	// The Base 64 mapping , String of it's binary representation to it's rune 
	// 0-25: A-Z
	// 26-51: a-z
	// 52-61: 0-9
	// 62: +
	// 63: /
	map64 = make(map[string]rune)
)

func main() {
	var base16 string = "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	var err error
	// make the base16 string all Uppercase
	base16 = strings.ToUpper(base16)
	// checking if the base16 stirng  contains [0-9 || A-F]
	base16, err = check_base16(base16)
	if err != nil {
		panic(err)
	}
	// every 3 character of the base16 gives as 2 character in base64
	// we gonna do some Padding to adjust that . 
	for length := 3 - len(base16) % 3; length > 0; length-- {
		base16 = "0" + base16
	}
	// fmt.Printf("the base16 string : %s\n", base16)
	// initialize the maps  (rune -> binary representation) 
	//for the map16 , rune -> 4 runes
	// and the opposite for the map64 , 6runes -> rune 
	initializer()
	// creating a binaryString that reprensent the base16 string in a specific way, 
	// every rune from the base16 string are transformed to 4rune , the 4rune are the binary representation of the value in 4bits .
	binaryString := hx_to_binary(base16)
	// from the binaryString we gonna take every 6rune which will be mapped to it's rune in the map64 . 
	base64 := hex_to_base64(binaryString)
	// Now lets Print the base64 string .
	fmt.Println(base64)
	// Thanks for Reading . 
}

func initializer() {
	i := byte(0)
	var binaryString string

	for _, char := range table16 {
		binaryString = fmt.Sprintf("%04b", i)
		map16[char] = binaryString
		i++
	}
	i = 0
	for _, char := range table64 {
		binaryString = fmt.Sprintf("%06b", i)
		map64[binaryString]= char
		i++
	}
}

func check_base16(base16 string) (string, error) {
	if len(base16) >= 2 && strings.Compare(base16[0:2], "0X") == 0 {
		base16 = base16[2:len(base16)]
	}
	for _,char := range base16 {
		if (byte(char) < 48 || byte(char) > 57) && (byte(char) < 65 || byte(char) > 70){
			return base16, errors.New("Invalid input !")
		}
	}
	return base16, nil
}

func hx_to_binary(base16 string) string {
	var binaryString string

	for  _, char := range base16 {
		binaryString += map16[char]
	}
	// fmt.Printf("the binary representation : %s\n", binaryString)
	return binaryString
}

func hex_to_base64(binaryString string) string {
	var base64 string
	for i := 0; i < len(binaryString); i += 6 {
		base64 += string(map64[binaryString[i:i+6]])
	}
	// You can Add a little "for loop" here to remove the Zeros in the left of the base64 string
	// the zeros in this case are 'A'

	// I prefere not .
	return base64
}