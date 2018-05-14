package main

import (
	"fmt"
	"os"
)

// 通常のRun Length法
type RunlengthFixed struct {}
func (r *RunlengthFixed)Encode(data []uint8) []uint8 {
	var compressed []uint8
	var p, cnt uint8 = data[0], 1
	for _, d := range data[1:] {
		if p != d || cnt == 255{
			compressed = append(compressed, p, cnt)
			cnt = 0
		}
		p = d
		cnt++
	}
	compressed = append(compressed, p, cnt)
	return compressed
}

func (r *RunlengthFixed)Decode(data []uint8) []uint8{
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

var bitBuff uint32 = 0
var lenBitBuff uint8 = 0
func initBitBuff(){
	bitBuff = 0
	lenBitBuff = 0
}
func convertWyle(data uint32) (uint32, uint8){

	// 先頭部分の1の長さを計算
	var lenData uint8 = 0
	tmp := data-1
	tmp /= 4
	for tmp > 0{
		tmp /= 2
		lenData++
	}

	// 先頭部分の1を追加
	var ret uint32 = 0
	if lenData > 0{
		for i:=0; i<int(lenData); i++{
			ret |= 1<<(uint8(i)+ lenData +3)
		}
	}

	// dataを追加
	ret |= data

	// 区切り文字の0とdataの長さを加算
	lenData += 1+ lenData +2

	return ret, lenData
}
//func appendBitBuff(outputData []uint8, bit uint8, lenBit uint8){
//	bitBuff = bitBuff | bit << lenBitBuff
//	lenBitBuff++
//
//	if lenBitBuff == 8 {
//	}
//}

// Wyle符号でのRun Length法
type RunlengthWyle struct {}
func (r *RunlengthWyle)Encode(data []uint8) []uint8 {
	//var compressed []uint8
	var p uint8 = data[0]
	var cnt uint32 = 1
	for _, d := range data[1:]{
		if p != d {
			cnt = 0
		}
		cnt++
	}
	return data
}

