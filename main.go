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

	filename := "./img/taxio.png"
	data := read_file(filename)

	compAlg := &Huffman{}
	var comp Base = compAlg
	fmt.Println("encode_runlength compress")

	fmt.Println("before")
	fmt.Printf("%v...%v\n", data[:10], data[len(data)-10:])
	fmt.Printf("%d bytes\n", len(data))

	compressed := comp.Encode(data)
	fmt.Println("after")
	fmt.Printf("%v...%v\n", compressed[:10], compressed[len(compressed)-10:])
	fmt.Printf("%d bytes\n", len(compressed))
	output_file("./img/encode.lngrs", compressed)

	decoded := comp.Decode(compressed)
	fmt.Println("decode")
	fmt.Printf("%v...%v\n", decoded[:10], decoded[len(decoded)-10:])
	fmt.Printf("%d bytes\n", len(decoded))
	output_file("./img/decoded.png", decoded)

}