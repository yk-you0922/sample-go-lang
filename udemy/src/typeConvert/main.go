package main

import (
	"fmt"
	"strconv"
)

func main() {
	/*
	* 数値　⇔　浮動小数点の変換
	* 浮動小数点 ⇒ 数値への変換時、小数点以下は切り捨てになるため注意
	 */
	var i int = 1
	fl64 := float64(i)
	fmt.Println(fl64)
	fmt.Printf("i= %T\n", i)
	fmt.Printf("fl64 = %T\n", fl64)

	i2 := int(fl64)
	fmt.Println(fl64)
	fmt.Printf("i2= %T\n", i2)
	fmt.Printf("fl64 = %T\n", fl64)

	fl := 10.5
	i3 := int(fl)
	fmt.Println(i3)
	fmt.Printf("i3= %T\n", i3)
	fmt.Printf("fl = %T\n", fl)

	/*
	* 文字列　⇔　数値の変換
	* strconv()を用いる
	* 文字列 ⇒ 数値 strconv.Atoi()
	* 数値 ⇒ 文字列 strconv.Itoa()
	* strconv()は戻り値二つ（1つ目：変換結果、2つ目：エラー発生時のエラー内容
	 */
	var s string = "100"
	// var s string = "a" // 数値にできない文字列を与えた場合は以下でエラーハンドリングできる
	fmt.Printf("s = %T\n", s)

	// strconv()では二つの戻り値があるが、アンダーバーを定義することで使わない（破棄する）ことができる
	// i4, _ := strconv.Atoi(s)
	i4, err := strconv.Atoi(s) // エラーハンドリング
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(i4)
	fmt.Printf("i4 = %T\n", i4)

	var i5 int = 200
	s2 := strconv.Itoa(i5)
	fmt.Println(s2)
	fmt.Printf("s2 = %T\n", s2)

	/*
	* 文字列　⇔　バイト配列の変換
	 */
	 // 文字列 ⇒ バイト配列
	var h string = "Hello World"
	b := []byte(h)
	fmt.Println(b)

	// バイト配列 ⇒ 文字列
	h2 := string(b)
	fmt.Println(h2)
}
