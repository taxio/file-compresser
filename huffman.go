package main

import (
	"sort"
)

type Huffman struct {}
func (h *Huffman)Encode(data []uint8) []uint8 {
	sorted := make([]uint8, len(data))
	copy(sorted, data)
	sort.Slice(sorted, func(i, j int) bool {return sorted[i] < sorted[j]})

	// create bit data

	// create encoded data

	return data
}
func (h *Huffman)Decode(data []uint8) []uint8 {
	// parse bit data table
	// create decoded data
	return data
}
