//package file_compresser
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

func encode_runlength(data []uint8) []uint8 {
	var compressed []uint8
	var p, cnt uint8 = data[0], 0
	for i, d := range data {
		cnt++
		if i == 0{
			p = d
			cnt = 0
			continue
		}

		if p != d {
			compressed = append(compressed, p)
			compressed = append(compressed, cnt)
			cnt = 0
		}
		p = d
	}
	if cnt == 0{
		compressed = append(compressed, p)
		compressed = append(compressed, 1)
	}
	return compressed
}

func decode_runlength(data []uint8) []uint8{
	var decoded []uint8
	if len(data) % 2 == 1{
		fmt.Errorf("this lngrs file is incorrect\n")
		os.Exit(-1)
	}

	for i:=0; i<len(data); i+=2{
		for j:=0; j<int(data[i+1]); j++{
			decoded = append(decoded, data[i])
		}
	}
	return decoded
}

func main(){

	filename := "./img/taxio.png"
	data := read_file(filename)

	fmt.Println("encode_runlength compress")

	fmt.Println("before")
	fmt.Print(data[0:10])
	fmt.Print("...")
	fmt.Print(data[len(data)-10:])
	fmt.Printf("\n%d bytes\n", len(data))

	compressed := encode_runlength(data)
	fmt.Println("after")
	fmt.Print(compressed[0:10])
	fmt.Print("...")
	fmt.Print(compressed[len(compressed)-10:])
	fmt.Printf("\n%d bytes\n", len(compressed))
	//output_file("./img/taxio.lngrs", compressed)

	decoded := decode_runlength(compressed)
	fmt.Println("decode")
	fmt.Print(decoded[0:10])
	fmt.Print("...")
	fmt.Print(decoded[len(decoded)-10:])
	fmt.Printf("\n%d bytes\n", len(decoded))

}