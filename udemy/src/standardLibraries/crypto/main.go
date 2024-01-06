package main

import (
	"crypto/md5"
	"fmt"
	"io"
)

func main() {
	// MD5ハッシュ値を生成
	// 任意の文字列からMD5ハッシュ値を生成する処理例
	h := md5.New()

	io.WriteString(h, "ABCDE") // 文字列"ABCDE"をハッシュ値に変換

	fmt.Println(h.Sum(nil))

	s := fmt.Sprintf("%x", h.Sum(nil)) // ハッシュ値から16進数の文字列を生成する。
	fmt.Println(s)
}
