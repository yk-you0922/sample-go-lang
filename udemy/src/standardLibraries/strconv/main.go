package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 真偽値を文字列に変換
	bt := true
	fmt.Printf("%T\n", strconv.FormatBool(bt))

	// 値と型の文字列
	var t string = "%v, %T\n"

	// 整数を文字列に変換
	i := strconv.FormatInt(-100, 10)
	fmt.Printf(t, i, i)
	// 簡易的に変換可能
	i2 := strconv.Itoa(100)
	fmt.Printf(t, i2, i2)

	// 浮動小数点型に変換
	fmt.Println(strconv.FormatFloat(123.456, 'E', -1, 64))
	// 指数表現による文字列化（小数点以下２桁まで）
	fmt.Println(strconv.FormatFloat(123.456, 'e', 2, 64))
	// 実数表現による文字列化
	fmt.Println(strconv.FormatFloat(123.456, 'f', -1, 64))
	// 実数表現による文字列化（小数点以下２桁まで）
	fmt.Println(strconv.FormatFloat(123.456, 'f', 2, 64))
	// 指数部の大きさで変動する表現による文字列化
	fmt.Println(strconv.FormatFloat(123.456, 'g', -1, 64))
	fmt.Println(strconv.FormatFloat(123456789.123, 'g', -1, 64))
	// 指数部の大きさで変動する表現による文字列化（仮数部全体が4桁まで）
	fmt.Println(strconv.FormatFloat(123.456, 'g', 4, 64))
	// 指数部の大きさで変動する表現による文字列化（仮数部全体が8桁まで）
	fmt.Println(strconv.FormatFloat(123456789.123, 'g', 8, 64))

	// 文字列を真偽値に変換（True）第二引数はエラー型
	bt1, _ := strconv.ParseBool("true")
	fmt.Printf(t, bt1, bt1)
	bt2, _ := strconv.ParseBool("1")
	bt3, _ := strconv.ParseBool("t")
	bt4, _ := strconv.ParseBool("T")
	bt5, _ := strconv.ParseBool("TRUE")
	bt6, _ := strconv.ParseBool("True")
	fmt.Println(bt2, bt3, bt4, bt5, bt6)

	// 文字列を真偽値に変換（false）第二引数はエラー型
	bf1, _ := strconv.ParseBool("false")
	fmt.Printf(t, bf1, bf1)
	bf2, _ := strconv.ParseBool("0")
	bf3, _ := strconv.ParseBool("f")
	bf4, _ := strconv.ParseBool("F")
	bf5, _ := strconv.ParseBool("FALSE")
	bf6, _ := strconv.ParseBool("False")
	fmt.Println(bf2, bf3, bf4, bf5, bf6)

	// 文字列を整数に変換する。第二引数はエラー型
	i3, _ := strconv.ParseInt("12345", 10, 0) // 変換対象文字列, X進数, 精度
	fmt.Printf(t, i3, i3)
	i4, _ := strconv.ParseInt("-1", 10, 0)
	fmt.Printf(t, i4, i4)
	// 簡易的な変換も可能
	i10, _ := strconv.Atoi("123")
	fmt.Println(i10)

	// 文字列を浮動小数点型に変換する。第二引数はエラー型
	fl1, _ := strconv.ParseFloat("3.14", 64) // 変換対象文字列, 精度
	fl2, _ := strconv.ParseFloat(".2", 64)
	fl3, _ := strconv.ParseFloat("-2", 64)
	fl4, _ := strconv.ParseFloat("1.2345e8", 64)
	fl5, _ := strconv.ParseFloat("1.2345E8", 64)
	fmt.Println(fl1, fl2, fl3, fl4, fl5)
}
