package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 標準入力を行単位で読み込む
	// スキャナの生成
	scanner := bufio.NewScanner(os.Stdin)

	// 入力のスキャンが成功する限り繰り返すループ
	for scanner.Scan() {
		// スキャン内容を文字列で出力
		fmt.Println(scanner.Text())
	}

	// スキャンにエラーが発生した場合の処理
	if err := scanner.Err(); err != nil {
		fmt.Println(os.Stderr, "読み込みエラー", err)
	}

}
