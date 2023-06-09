package main

import (
	"Cryptopals/Xor"

	"encoding/base64"
	"io/ioutil"
	// "os"
	"log"
	"fmt"
	// "bufio"
)

func main() {
	filePath := "./InputFiles/input6.txt"
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("File contents:\n%s\n", string(content))
	input, _:= base64.StdEncoding.DecodeString(string(content))
	Xor.BreakRepeatingKeyXOR(input)





	// filePath := "./InputFiles/input4.txt"
	// fd, err := os.Open(filePath)
    // if err != nil {
    //     log.Fatalf("unable to read file: %v", err)
    // }
    // defer fd.Close()
	// lines := make([]string, 327)
	// scanner := bufio.NewScanner(fd)
	// for i := 0; scanner.Scan() && i < 327; i++ {
	// 	lines[i] = scanner.Text()
	// 	fmt.Println(lines[i])
	// }

	// // checking if an error happend during the file scanning 
    // if error := scanner.Err(); error != nil {
    //     log.Fatal(error)
    // }
	 
	// str, key, err := Xor.DetectSingleCharXOR(lines)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("The result :")
	// fmt.Println(key)
	// fmt.Println(str)
}