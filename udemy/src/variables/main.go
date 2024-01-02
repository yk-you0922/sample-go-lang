package main

import "fmt"

// 関数より外のスコープで暗黙的型定義はできない。
// i5 := 500

// 明示的な変数定義はできる
var i5 int = 500

func main() {
	// 明示的な変数定義
	// var 変数名 型 = 値
	var i int = 100

	fmt.Println(i)

	var s string = "Hello Go"
	fmt.Println(s)

	// 同じ型であれば、1行で複数の変数定義もできる。
	var t, f bool = true, false
	fmt.Println(t, f)

	// ()を用いた一括定義
	var (
		i2 int    = 200
		s2 string = "Golang"
	)
	fmt.Println(i2, s2)

	// 変数の型のみを定義（値は定義しない）※各型の初期値が設定される。（intは0, stringは空文字 ""が設定）
	var i3 int
	var s3 string
	fmt.Println(i3, s3)

	i3 = 300
	s3 = "Go"
	fmt.Println(i3, s3)

	// 暗黙的な定義
	// 変数名 := 値
	i4 := 400
	fmt.Println(i4)

	i4 = 450
	fmt.Println(i4)

	i4 = 500
	fmt.Println(i4)

	// i4 = "Hello" // ※型が異なるため、コンパイルエラーとなる。
	// fmt.Println(i4)

	fmt.Println(i5)

	// 関数呼び出し
	outer()
	// fmt.Println(s4) // outer()関数内での変数はmain関数内で呼び出せない
}

func outer() {
	var s4 string = "outer"
	fmt.Println(s4)
}
