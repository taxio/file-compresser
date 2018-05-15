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

func output_file(filename string, data []uint8) error {

	// create output file
	file, err := os.Create(filename)
	defer file.Close()
	if err != nil{
		return err
	}
	file.Write(data)

	return nil
}


func main(){

	inputName := "./text/bin.txt"
	compressedName := "./output/compressed.txt"
	decodedName := "./output/decoded.txt"
	data := read_file(inputName)

	comp := &RunlengthWyle{}
	fmt.Printf("before : %d bytes\n", len(data))

	compressed := comp.Encode(data)
	fmt.Printf("after : %d bytes\n", len(compressed))
	err := output_file(compressedName, compressed)
	if err != nil{
		fmt.Errorf("%v\n", err)
	}
	fmt.Printf("Raito : %f%%\n", float32(len(compressed))/float32(len(data))*100)

	decoded := comp.Decode(compressed)
	fmt.Printf("decode : %d bytes\n", len(decoded))
	err = output_file(decodedName, decoded)
	if err != nil{
		fmt.Errorf("%v\n", err)
	}

}