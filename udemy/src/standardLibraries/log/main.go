package main

import (
	"log"
	"os"
)

func main() {
	// ロガーの生成
	logger := log.New(os.Stdout, "", log.Ldate | log.Ltime | log.Lshortfile)
	logger.Println("message")
	log.Println("message")

	// 条件分岐。エラーで終了させる。
	_, err := os.Open("fadafdsafa")
	if err != nil {
		// ログ出力
		// log.Fatalln("Exit", err)
		logger.Fatalln("Exit, err")
	}
}
