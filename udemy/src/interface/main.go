package main

import "fmt"

func main() {
	/*
	* 他のstring型やint型と互換性のある型定義
	* ただし、他の型特有の演算などはできない⇒あくまでinterface型という扱いになるため。
	 */
	var x interface{}
	fmt.Println(x) // <nil>が出力される

	x = 1
	fmt.Println(x)
	// fmt.Println(x + 3) // int型特有の計算（足し算）などはできない

	x = "interface"
	fmt.Println(x)
	// fmt.Println(x[0]) // string型特有の文字列の切り出しなどはできない

	x = []int{1, 2, 3}
	fmt.Println(x)

}
