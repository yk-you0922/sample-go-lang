package main

import (
	"fmt"
)

// 型 byte

func main() {
	byteA := []byte{72, 73}
	fmt.Println(byteA)

	// 文字列への変換
	fmt.Println(string(byteA))

	c := []byte("HI")
	fmt.Println(c)

	fmt.Println(string(c))
}
