package main

import "fmt"

// スライス
// for

func main() {
	sl := []string{"A", "B", "C"}
	fmt.Println(sl)

	// i・・・index		もし使用しない場合はアンダーバーで宣言することで破棄できる
	// v・・・value
	for i, v := range sl {
		fmt.Println(i, v)
	}

	// 古典的forによる記述
	for i := 0; i < len(sl); i++ {
		fmt.Println(sl[i])
	}
}