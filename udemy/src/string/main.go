package main

import "fmt"

func main() {
	var s string = "Hello Golang"
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	var si string = "300"
	fmt.Println(si)
	fmt.Printf("%T\n", si)

	// 複数行にわたる文字列の出力
	fmt.Println(`test
	test
		test`)

	// ダブルクォーテーションを出力するとき
	fmt.Println("\"")
	fmt.Println(`"`)

	// 一文字目を切り出し(バイト型で出力される)
	fmt.Println("文字列切り抜き（バイト型で出力される）")
	fmt.Println(s[0])
	fmt.Println(s[1])
	fmt.Println(s[2])
	fmt.Println("上記のバイトで切り取られたものを文字列に変換して出力")
	fmt.Println(string(s[0]))
	fmt.Println(string(s[1]))
	fmt.Println(string(s[2]))
}
