package main

import "fmt"

func main() {
	var i int = 100

	/*
	* int型の種類色々
	* int8
	* int16
	* int32
	* int64
	* それぞれ格納できる最小値・最大値が異なる
	* 環境依存（PCが64bitなら int64 を使う）
	 */

	var i2 int64 = 200

	fmt.Println(i + 50)
	// fmt.Println(i + i2) // 型（bit数)が異なるため、計算できない⇒型変換すればできる。

	// 型を調べる
	fmt.Printf("%T\n", i2) // int64が出力される

	// 型変換
	fmt.Printf("%T\n", int32(i2)) // int32が出力される。
}
