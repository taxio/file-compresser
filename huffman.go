package main

import (
	"sort"
	"fmt"
)

type HuffmanNode struct {
	data uint8
	cnt uint8

	parent *HuffmanNode
	left *HuffmanNode
	right *HuffmanNode
}
func NewHuffmanNode(data uint8) HuffmanNode{
	node := HuffmanNode{}
	node.data = data
	node.cnt = 0
	node.parent = nil
	node.left = nil
	node.right = nil
	return node
}

func countContains(data []uint8, target uint8) uint8{
	var cnt uint8 = 0
	for _, d := range data {
		if d == target{
			cnt++
		}
	}
	return cnt
}

type Huffman struct {}
func (h *Huffman)Encode(data []uint8) []uint8 {
	nodes := make([]HuffmanNode, 0, 256)
	for i:=0; i<256; i++{
		cnt := countContains(data, uint8(i))
		if cnt > 0{
			node := NewHuffmanNode(uint8(i))
			node.cnt = cnt
			nodes = append(nodes, node)
		}
	}
	fmt.Printf("node len : %d\n", len(nodes))
	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].cnt < nodes[j].cnt
	})
	fmt.Println(nodes)

	// create bit data


	// create encoded data

	return data
}
func (h *Huffman)Decode(data []uint8) []uint8 {
	// parse bit data table
	// create decoded data
	return data
}
