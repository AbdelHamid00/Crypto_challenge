package Xor

import (
	"Cryptopals/clc"
	// "Cryptopals/Xor"
	// "math"
	"fmt"
	"sort"
)

func BitsDefference(b1 byte, b2 byte) int {
	def := b1 ^ b2
	counter := int(0)
	for def != 0 {
		counter += int(def & 1)
		def = def >> 1
	}
	return counter
}

func HammigDestance(str1 string, str2 string) int{
	counter := int(0)
	l1 := len(str1)
	l2 := len(str2)
	for i := 0; i < clc.Min(l1, l2); i++ {
		if str1[i] != str2[i] {
			counter += BitsDefference(str1[i], str2[i])
		}
	}
	counter += clc.Abs(l1 - l2) * 8
	return counter
}

func GetTheScore(input []byte, KeySize int) float64 {
	var Blocks [8]string
	sum := float64(0)

	Blocks[0] = string(input[0:KeySize])
	Blocks[1] = string(input[KeySize:KeySize*2])
	// Blocks[2] = string(input[KeySize*2:KeySize*3])
	// Blocks[3] = string(input[KeySize*3:KeySize*4])
	// Blocks[4] = string(input[KeySize*4:KeySize*5])
	// Blocks[5] = string(input[KeySize*5:KeySize*6])
	// Blocks[6] = string(input[KeySize*6:KeySize*7])
	// Blocks[7] = string(input[KeySize*7:KeySize*8])
	k := 0
	for i := 0; i < 2; i++ {
		for j := i + 1; j < 2; j++ {
			k++
			sum += float64(HammigDestance(Blocks[i], Blocks[j]))
		}
	}
	return sum / float64(k) 
}

type KeyValue struct {
	Key   int
	Value float64
}

func DivideTheinputByKeySize(originalBytes []byte, chunkSize []KeyValue) [][]byte {
	l := len(originalBytes)
	totalChunks := l / chunkSize[0].Key
	if l % chunkSize[0].Key != 0 {
		totalChunks += 1
	}
	chunks := make([][]byte, totalChunks)
	for i := 0; i < totalChunks; i++ {
		start := i*chunkSize[0].Key
		end := (i+1)*chunkSize[0].Key
		if end > l {
			end = l
		}
		chunks[i] = originalBytes[start : end]
	}
	return chunks
}

func GetTheKeySize(input []byte)  []KeyValue{
	Arr := make([]KeyValue, 0, 39)

	for KeySize := 2; KeySize <= 40; KeySize++ {
		actual_score := GetTheScore(input, KeySize) / float64(KeySize)
		Arr = append(Arr, KeyValue{KeySize, actual_score})
	}
	sort.Slice(Arr, func(i, j int) bool {
		return Arr[i].Value < Arr[j].Value
	})
	return Arr
}

func BreakRepeatingKeyXOR(input []byte){
	Arr := GetTheKeySize(input)
	// fmt.Println(Arr)
	blocks := DivideTheinputByKeySize(input, Arr)
	samebytes :=make([][]byte, Arr[0].Key)
	var KeyBytes []byte
	for j := 0; j < Arr[0].Key; j++ {
		for i := 0; i < len(blocks); i++ {
			if (j < len(blocks[i])) {
				samebytes[j] = append(samebytes[j], blocks[i][j])
			}
		}
		_, _, Key := GetTheBestScoreForOne(string(samebytes[j]))
		KeyBytes = append(KeyBytes, Key)
	}
	result := RepeatingKeyXOR(input, KeyBytes)
	fmt.Println(string(result))
}