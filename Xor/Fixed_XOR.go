package Xor

import (
	"Cryptopals/clc"
)

func	XorTwoStrings(in1 []byte, in2 []byte) []byte {
	var result []byte
	Length := clc.Min(len(in1), len(in2))
	for i := 0; i < Length; i++{
		xor := in1[i] ^ in2[i]
		result = append(result, xor)
	}
	return result
}
