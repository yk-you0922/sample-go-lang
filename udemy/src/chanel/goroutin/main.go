package main

import (
	"fmt"
	"time"
)

// チャネルとゴルーチン

/*
* チャネルから整数を受け取り、それを表示する関数
 */
func reciever(c chan int) {
	for {
		i := <- c
		fmt.Println(i)
	}
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	// fmt.Println(<-ch1)

	// チャネルからデータを受信し、表示する。
	go reciever(ch1)
	go reciever(ch2)

	// チャネルにデータを送る
	i := 0
	for i < 100 {
		ch1 <- i
		ch2 <- i
		time.Sleep(50 * time.Millisecond)
		i++
	}

}