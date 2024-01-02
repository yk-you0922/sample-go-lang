package main

import "fmt"

// 定数は関数外のグローバルに定義することが多い
// 頭文字を大文字とすることで他のパッケージからも呼び出せるようになる
// 頭文字小文字の定数宣言は他のパッケージで使われたくないプライベートという扱いになる
const Pi = 3.14

const (
	URL      = "http://xxx.co.jp"
	SiteName = "test"
)

// 値を省略して定義できる
const (
	A = 1
	B // 定義していないが、Aに続いて1が入る。
	C // 定義していないが、Aに続いて1が入る。
	D = "A"
	E // 定義していないが、Dに続いて"A"が入る。
	F // 定義していないが、Dに続いて"A"が入る。
)

// 定数は最大値を持っていないため、巨大な値をいれれる。
// var Big int = 9223372036854775807 + 1 // 最大値オーバー
const Big = 9223372036854775807 + 1

const (
	c0 = iota // iota : 連続する整数の連番
	c1
	c2
)

func main() {
	fmt.Println(Pi)

	// Pi = 3 // 定数のため、再代入できない
	// fmt.Println(Pi)

	fmt.Println(URL, SiteName)

	fmt.Println(A, B, C, D, E, F)

	fmt.Println(Big - 1)
	
	fmt.Println(c0, c1, c2)
}
