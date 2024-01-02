package main

import "fmt"

// 関数を返す関数
func ReturnFunc() func() {
	return func() {
		fmt.Println("Im a function")
	}
}

func main() {
	f := ReturnFunc() // ReturnFunc()を変数fに代入
	f() // 変数fに代入したReturnFunc()を実行
}
