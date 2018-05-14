package main

import (
	"io/ioutil"
	"fmt"
	"os"
)

func read_file(filename string) []uint8 {
	data, err := ioutil.ReadFile(filename)		// uint8で配列が返ってくる
	if err != nil {
		fmt.Errorf("cannot read %s\n", filename)
		os.Exit(-1)
	}
	fmt.Printf("read %s\n", filename)
	return data
}

func output_file(filename string, data []uint8) {
	file, err := os.Create(filename)
	defer file.Close()
	if err != nil{
		fmt.Errorf("cannot create output file")
	}else{
		file.Write(data)
	}
}


func main(){

	inputName := "./text/bin.txt"
	outputName := "./text/bin_decoded.txt"
	data := read_file(inputName)

	comp := &RunlengthFixed{}

	fmt.Println("before : ")
	fmt.Printf("before : %d bytes\n", len(data))
	fmt.Println(data)

	compressed := comp.Encode(data)
	fmt.Println("after")
	fmt.Printf("after : %d bytes\n", len(compressed))
	fmt.Println(compressed)
	output_file(outputName, compressed)

	decoded := comp.Decode(compressed)
	fmt.Printf("decode : %d bytes\n", len(decoded))
	fmt.Println(decoded)

}