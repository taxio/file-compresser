package main

type Base interface {
	Encode([]uint8) []uint8
	Decode([]uint8) []uint8
}

