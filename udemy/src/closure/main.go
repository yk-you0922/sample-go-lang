package main

import "fmt"

// クロージャー
func Later() func(string) string {
	var store string // クロージャーが参照される限り破棄されない。

	// 前回のstoreに入っている値をsに代入し、返却する。
	return func(next string) string {
		s := store
		store = next
		return s
	}
}

func main() {
	f := Later()
	fmt.Println(f("Hello"))  // 初回で変数storeが空文字のため、何も出力されない。
	fmt.Println(f("My"))     // 前回の"Hello"が出力される
	fmt.Println(f("Name"))   // 前回の"My"が出力される
	fmt.Println(f("Is"))     // 前回の"Name"が出力される
	fmt.Println(f("Golang")) // 前回の"Is"が出力される
}
