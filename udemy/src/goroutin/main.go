package main

import (
	"fmt"
	"time"
)

// 平行
// go goroutin

/*
* 「go」を関数の前につけるだけで並行処理を実装できる。
* 並行処理・・・スレッドよりも小さい単位で処理を行う処理
* 並行処理・並列処理・同期処理・非同期処理などの分類があるため、注意。
*/

// 並行処理関数
func sub() {
	for {
		fmt.Println("Sub Loop")
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	go sub()
	go sub()

	for {
		fmt.Println("Main Loop")
		time.Sleep(200 * time.Millisecond)
	}
}