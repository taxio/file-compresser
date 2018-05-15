package main

import (
	"fmt"
	"os"
	"errors"
)

// 通常のRun Length法
type RunlengthFixed struct {
}
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

var bitBuff uint32 = 0
var lenBitBuff uint8 = 0
var outputData []uint8
func initBitBuff(){
	bitBuff = 0
	lenBitBuff = 0
	outputData = make([]uint8, 0)
}

// 入力されたデータのビットを逆順にして返す
func reverseBits(data uint32, lenData uint8) uint32 {
	var r uint32 = 0
	for i:=int(lenData-1); i>=0; i--{
		r |= (data>>uint8(i) & 1) << (lenData-uint8(i)-1)
	}
	return r
}

// TODO: 多分bitが逆になってる
func addOutputBuff(data uint32, lenData uint8) error{
	if lenBitBuff+lenData > 32{
		return errors.New("over output buffer")
	}
	// dataの並びを反転
	data = reverseBits(data, lenData)

	// バッファに追加
	bitBuff |= data << lenBitBuff
	lenBitBuff += lenData

	// バッファに8bit以上溜まったら出力
	for lenBitBuff/8 > 0{
		o := bitBuff & 0xff
		o = reverseBits(o, 8)
		outputData = append(outputData, uint8(o))
		bitBuff >>= 8
		lenBitBuff -= 8
	}
	return nil
}

func displayBits(bits uint32, lenBits uint8){
	for i:=int(lenBits)-1; i>=0; i--{
		fmt.Printf("%d", 1 & (bits >> uint8(i)))
	}
}

// Wyle符号でのRun Length法
type RunlengthWyle struct {}
func (r *RunlengthWyle)Encode(data []uint8) []uint8 {
	initBitBuff()
	var p uint8 = data[0]
	var cnt uint32 = 1
	for _, d := range data[1:]{
		if p != d {
			wyle, lenWyle := convertWyle(cnt)
			addOutputBuff(wyle, lenWyle)
			displayBits(wyle, lenWyle)
			addOutputBuff(uint32(p), 8)
			displayBits(uint32(p), 8)
			cnt = 0
		}
		p = d
		cnt++
	}
	wyle, lenWyle := convertWyle(cnt)
	addOutputBuff(wyle, lenWyle)
	displayBits(wyle, lenWyle)
	addOutputBuff(uint32(p), 8)
	addOutputBuff(0, 8-lenBitBuff)
	displayBits(uint32(p), 8)
	fmt.Println("")
	return outputData
}

//func (r *RunlengthWyle)Decode(data []uint8) []uint8{
//	for _, d := range data {
//		fmt.Printf("%b", d)
//	}
//	fmt.Println("")
//	return data
//}

