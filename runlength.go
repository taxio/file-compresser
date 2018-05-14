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
			compressed = append(compressed, p)
			compressed = append(compressed, cnt)
			cnt = 0
		}
		p = d
		cnt++
	}
	compressed = append(compressed, p)
	compressed = append(compressed, cnt)
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

