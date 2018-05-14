package main

import (
	"sort"
	"fmt"
)

type HuffmanNode struct {
	data uint8
	cnt uint32
	code *[]bool

	parent *HuffmanNode
	left *HuffmanNode
	right *HuffmanNode
}
func NewHuffmanNode(data uint8) HuffmanNode{
	node := HuffmanNode{}
	node.data = data
	node.cnt = 0
	code := make([]bool, 0, 256)
	node.code = &code
	node.parent = nil
	node.left = nil
	node.right = nil
	return node
}

func countContains(data []uint8, target uint8) uint32{
	var cnt uint32 = 0
	for _, d := range data {
		if d == target{
			cnt++
		}
	}
	return cnt
}

func addNode(head HuffmanNode, left HuffmanNode, right HuffmanNode) HuffmanNode{
	head.left = &left
	left.parent = &head
	head.right = &right
	right.parent = &head
	head.cnt = head.left.cnt + head.right.cnt
	return head
}

func create_tree(nodes []HuffmanNode) HuffmanNode{
	head := NewHuffmanNode(0)
	if len(nodes) == 1 {
		head.left = &nodes[0]
		head.cnt = head.left.cnt
		return head
	}
	head = addNode(head, nodes[0], nodes[1])
	for _, node := range nodes[2:] {
		tmp := NewHuffmanNode(0)
		if node.cnt < head.cnt {
			tmp = addNode(tmp, node, head)
		}else{
			tmp = addNode(tmp, head, node)
		}
		head = tmp
	}
	return head
}

func attachCode(node HuffmanNode, nextBit bool){
	parentCode := *node.parent.code
	*node.code = append(*node.code, parentCode...)
	*node.code = append(*node.code, nextBit)
	if node.right == nil && node.left == nil{
		return
	}
	if node.right != nil{
		attachCode(*node.right, false)
	}
	if node.left != nil{
		attachCode(*node.left, true)
	}
}

type Huffman struct {}
func (h *Huffman)Encode(data []uint8) []uint8 {
	// calc freq
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

	// create tree
	tree_head := create_tree(nodes)
	fmt.Println(tree_head)

	// create bit data
	attachCode(*tree_head.right, false)
	attachCode(*tree_head.left, true)

	// create encoded data

	return data
}
func (h *Huffman)Decode(data []uint8) []uint8 {
	// parse bit data table
	// create decoded data
	return data
}
