package main

import (
	"fmt"
	"os"
)

// defer

/*
* defer ・・・ 関数終了時に実行する内容を定義する。
* リソースの解放処理に用いることが多そう。
 */
func TestDefer() {
	defer fmt.Println("END")
	fmt.Println("START")
}

// deferは後から登録された順に評価される。　※注意点
func RunDefer() {
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
}

func main() {
	TestDefer()

	// 無名関数を用いることで、複数の処理を纏めて登録できる。
	// defer func() {
	// 	fmt.Println("1")
	// 	fmt.Println("2")
	// 	fmt.Println("3")
	// }()
	RunDefer()

	/*
	* リソースの解放処理例
	 */
	file, err := os.Create("test.text")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	file.Write([]byte("Hello"))
}
