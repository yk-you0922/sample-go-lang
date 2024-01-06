package main

import (
	"fmt"
	"regexp"
)

func main() {
	// 正規表現による文字列の置換
	re1 := regexp.MustCompile(`\s+`) // スペースを対象とした正規表現
	fmt.Println(re1.ReplaceAllString("AAA BBB CCC", ",")) // スペースをカンマに置き換え
	re2 := regexp.MustCompile(`、|。`)
	fmt.Println(re2.ReplaceAllString("私は、Golangを使用する、プログラマーです。", ""))

	// 正規表現による文字列の分割
	re3 := regexp.MustCompile(`(abc|ABC)(xyz|XYZ)`)
	fmt.Println(re3.Split("ASHVJV<HABCXYZKNJBJVKABCXYZ", -1))

	re4 := regexp.MustCompile(`\s+`)
	fmt.Println(re4.Split("aaaaaaaaa      bbbbbbb   cccccc", -1))
}
