package main

import "fmt"

func main() {
	// 無名関数
	f := func(x, y int) int {
		return x + y
	}
	i := f(1, 2)
	fmt.Println(i)

	// 無名関数を直接実行し、変数i2に値を代入するパターン
	i2 := func(x, y int) int {
		return x + y
	}(1, 2)
	fmt.Println(i2)
}
