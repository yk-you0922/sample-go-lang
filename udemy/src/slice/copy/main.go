package main

import "fmt"

// スライス
// コピー

func main() {
	sl := []int{100, 200}
	sl2 := sl

	sl2[0] = 1000 // slの一番目の要素も同様に更新されてしまう。
	fmt.Println(sl)
	fmt.Println(sl2)

	sl3 := []int{1, 2, 3, 4, 5}
	sl4 := make([]int, 5, 10)
	fmt.Println(sl4)
	n := copy(sl4, sl3) // n・・・コピーに成功した数を戻り値としてcopy関数は実行される

	fmt.Println(n, sl4)
}
