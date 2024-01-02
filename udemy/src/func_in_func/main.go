package main

import "fmt"

// 関数を引数に取る関数
func CallFunction(f func()) {
	f() // 引数の関数を実行する。
}

func main() {
	CallFunction(func() {
		fmt.Println("Im a Function")
	})
}