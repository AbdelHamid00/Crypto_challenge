package main

import(
	"fmt"
	"encoding/hex"
)

func	XorTwoStrings(in1 []byte, in2 []byte) []byte {
	if (len(in1) != len(in2)){
		panic("The strings should be equal to this function to work properly !")
	}
	var result []byte

	for i := 0; i < len(in1) && i < len(in2); i++{
		xor := in1[i] ^ in2[i]
		result = append(result, xor)
	}
	return result
}

func main(){
	// two equal-length buffers in1 and in2
	var in1 string = "1c0111001f010100061a024b53535009181c"
	var in2 string = "686974207468652062756c6c277320657965"

	in1Decoded, _ := hex.DecodeString(in1)
	in2Decoded, _ := hex.DecodeString(in2)
	result := XorTwoStrings([]byte(in1Decoded), []byte(in2Decoded))
	fmt.Println(string(hex.EncodeToString(result)))

	// resultDecoded, _ := hex.DecodeString(string(result))
	// fmt.Println(string(resultDecoded))
}