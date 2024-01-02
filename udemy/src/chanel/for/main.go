package main

import "fmt"

// チャネル-for

/*
* チャネルのfor文はclose()して使うのが基本⇒デッドロックがかかるため。
*/

func main() {
	ch1 := make(chan int, 3)
	ch1 <- 1
	ch1 <- 2
	ch1 <- 3
	close(ch1)
	// 簡易for文
	for i := range ch1 {
		fmt.Println(i)
	}
}