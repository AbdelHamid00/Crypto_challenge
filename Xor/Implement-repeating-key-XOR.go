package Xor

import (
	// "fmt"
)

func RepeatingKeyXOR(str []byte, key []byte) []byte{
	var result []byte
	j := 0
	Strlength := len(str)
	Keylength := len(key)

	for i := 0; i <= Keylength; i++ {
		i %= Keylength
		if j == Strlength {
			break 
		}
		xor := str[j] ^ key[i]
		result = append(result, xor)
		j++
	}
	return result
}