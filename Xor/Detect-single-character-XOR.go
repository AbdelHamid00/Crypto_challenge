package Xor

import (
	// "encoding/hex"
	"math"
	"strings"
	// "bytes"
)


func scorePlaintext(plaintext []byte) float64 {
	frequency := map[byte]float64{
		'a': 0.08167, 'b': 0.01492, 'c': 0.02782, 'd': 0.04253, 'e': 0.12702, 'f': 0.02228, 'g': 0.02015, 'h': 0.06094, 'i': 0.06966, 'j': 0.00253, 'k': 0.01772, 'l': 0.04025, 'm': 0.02406, 'n': 0.06749, 'o': 0.07507, 'p': 0.01929, 'q': 0.00095, 'r': 0.05987, 's': 0.06327, 't': 0.09056, 'u': 0.02758, 'v': 0.00978, 'w': 0.02360, 'x': 0.00250, 'y': 0.01974, 'z': 0.00074,
	}
	score := float64(0.0)
	l := len(plaintext)
	// plaintext = bytes.ToLower(plaintext) // just for optimisation but still not good enough
	for letter, frequency := range frequency {
		actual := float64(strings.Count(string(plaintext), string(letter))) / float64(l)
		score += math.Pow(math.Abs(actual - frequency), 2)
	}
	// lower the score is better .
	return score
}

func singleByteXORCipher(plaintext []byte, key byte) []byte {
	ciphertext := make([]byte, len(plaintext))
	for i := 0; i < len(plaintext); i++ {
		ciphertext[i] = plaintext[i] ^ key
	}
	return ciphertext
}

func GetTheBestScoreForOne(hexstring string) (minScore float64, plaintext string, Key byte) {
	minScore = float64(math.Inf(1))
	for key := 0; key <= 255; key++ {
		ciphertext := singleByteXORCipher([]byte(hexstring), byte(key))
		score := scorePlaintext(ciphertext)
		if score < minScore {
			minScore = score
			Key = byte(key)
			plaintext = string(ciphertext)
		}
	}
	return minScore, plaintext, Key
}

func DetectSingleCharXOR(lines []string) (plaintext string, Key byte, err error) {
	minScore := float64(math.Inf(1))
	for i := 0; i < len(lines); i++ {
		// plaintextBytes , err := hex.DecodeString(lines[i])
		// if err != nil {
		// 	return plaintext, Key, err
		// }

		score, Replaintext, ReKey := GetTheBestScoreForOne(lines[i])
		if score < minScore {
			minScore = score
			plaintext = Replaintext
			Key = ReKey
		}
	}
	return plaintext, Key, nil
}
