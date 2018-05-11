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

	runLength := &Runlength{}
	var comp Base = runLength
	fmt.Println("encode_runlength compress")

	fmt.Println("before")
	fmt.Print(data[0:10])
	fmt.Print("...")
	fmt.Print(data[len(data)-10:])
	fmt.Printf("\n%d bytes\n", len(data))

	compressed := comp.Encode(data)
	fmt.Println("after")
	fmt.Print(compressed[0:10])
	fmt.Print("...")
	fmt.Print(compressed[len(compressed)-10:])
	fmt.Printf("\n%d bytes\n", len(compressed))
	//output_file("./img/taxio.lngrs", compressed)

	decoded := comp.Decode(compressed)
	fmt.Println("decode")
	fmt.Print(decoded[0:10])
	fmt.Print("...")
	fmt.Print(decoded[len(decoded)-10:])
	fmt.Printf("\n%d bytes\n", len(decoded))

}