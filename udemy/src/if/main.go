package main

import (
	"fmt"
	"math/rand"
)

// 条件分岐

func main() {
	var result int = rand.Intn(3) // ランダムな整数を生成する。
	fmt.Println(result)

	if result == 2 {
		fmt.Println("two")
	} else if result == 1 {
		fmt.Println("one")
	} else {
		fmt.Println("I don't know")
	}

	// 簡易分付きIf文
	if b := 100; b == 100 {
		fmt.Println(b)
		fmt.Println("one hundred")
	}
	// fmt.Println(b) // if文内部でしか変数bは参照できない。
}