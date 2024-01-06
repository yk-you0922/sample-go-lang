package main

import (
	"fmt"
	"regexp"
)

func main() {
	re := regexp.MustCompile((`(abc|ABC)(xyz|XYZ)`))

	// 正規表現にマッチした文字列の取得
	// FindString
	fs1 := re.FindString("AAAAABCXYZZZZZZ")
	fmt.Println(fs1) // ABCXYZ
	fs2 := re.FindAllString("ABCXYZABCXYZ", -1)
	fmt.Println(fs2) // [ABCXYZ ABCXYZ]

	// 正規表現のグループによるサブマッチ
	// FindAllStringSubMatch
	re1 := regexp.MustCompile(`(\d+)-(\d+)-(\d+)`) // 一回以上繰り返す整数値-一回以上繰り返す整数値-一回以上繰り返す整数値 を取得する正規表現
	s := `
			0123-456-7889
			111-222-333
			556-787-899
		`
	ms := re1.FindAllStringSubmatch(s, -1)
	fmt.Println(ms)

	for _, v := range ms {
		fmt.Println(v)
	}
}
