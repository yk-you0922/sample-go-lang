package main

import (
	"fmt"
	"time"
)

// チャネル-クローズ

func reciever(name string, ch chan int) {
	for {
		i, ok := <- ch
		if !ok {
			break
		}
		fmt.Println(name, i)
	}
	fmt.Println(name + "END")
}

func main() {
	ch1 := make(chan int, 2)

	// close(ch1) // クローズ。クローズ後にデータを送るとruntime error

	// ch1 <- 1 // runtime error

	// fmt.Println(<-ch1) // クローズ後でも受信自体はできる。

	// 1つ目：チャネルから受け取った値。 2つ目：true・・・チャネルオープン, false・・・チャネルクローズ
	// false補足: チャネルのバッファ内が空である、かつ、クローズされているという状態。
	// i, ok := <-ch1 
	// fmt.Println(i, ok)

	go reciever("1. goroutin", ch1)
	go reciever("2. goroutin", ch1)
	go reciever("3. goroutin", ch1)

	i := 0
	for i < 100 {
		ch1 <- i
		i++
	}
	close(ch1)
	time.Sleep(3 * time.Second) // goroutinの完了を待つために簡易的にsleep
	
}