package main

import (
	"fmt"
	"strings"
)

func main() {
	// 文字列を結合する
	s1 := strings.Join([]string{"A", "B", "C"}, ",") // 結合対象スライス, 分割文字列
	s2 := strings.Join([]string{"A", "B", "C"}, "")
	fmt.Println(s1, s2)

	// 文字列に含まれる部分文字列を検索する
	i1 := strings.Index("ABCDE", "A")
	i2 := strings.Index("ABCDE", "D")
	fmt.Println(i1, i2)

	i3 := strings.LastIndex("ABCDABCD", "BC")
	fmt.Println(i3)

	i4 := strings.IndexAny("ABC", "ABC")
	
	i5 := strings.LastIndexAny("ABC", "ABC")
	fmt.Println(i4, i5)

	// 文字列を繰り返して結合する。
	s3 := strings.Repeat("ABC", 4)
	s4 := strings.Repeat("ABC", 0)
	fmt.Println(s3, s4)

	// 文字列の置換
	s5 := strings.Replace("AAAAA", "A", "B", 1)
	s6 := strings.Replace("AAAAA", "A", "B", -1)
	fmt.Println(s5, s6)

	// 文字列を分割する
	s7 := strings.Split("A,B,C,D,E", ",")
	s8 := strings.SplitAfter("A,B,C,D,E", ",")
	s9 := strings.SplitN("A,B,C,D,E", ",", 2)
	s10 := strings.SplitAfterN("A,B,C,D,E", ",", 4)
	fmt.Println(s7)
	fmt.Println(s8)
	fmt.Println(s9)
	fmt.Println(s10)

	// 大文字⇒小文字への変換
	s11 := strings.ToLower("ABC")
	s12 := strings.ToLower("E")

	// 小文字⇒大文字への変換
	s13 := strings.ToUpper("abc")
	s14 := strings.ToUpper("e")
	fmt.Println(s11, s12, s13, s14)

	// 文字列から空白を取り除く
	s15 := strings.TrimSpace("   -    ABC   -     ") // 全角スペース
	s16 := strings.TrimSpace("　　　ABC　　　　") // 半角スペース
	fmt.Println(s15, s16)
	
	// 文字列からスペースで区切られたフィールドを取り出す
	s18 := strings.Fields("a b c")
	fmt.Println(s18)
	fmt.Println(s18[1])
}
