package main

import "fmt"

// panic recover・・・例外処理にあたる
// 強力な処理にあたるため、あまり使わないことが推奨されている。
// 基本はエラーハンドリングを用いた処理を行うことが推奨されている。

func main() {

	// panicの内容を変数xに代入し、出力する処理
	defer func() {
		if x := recover(); x != nil {
			fmt.Println(x)
		}
	}()

	panic("runtime error")
	fmt.Println("Start")
}
