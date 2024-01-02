package main

// チャネル-select

func main() {
	ch1 := make(chan int, 2)
	ch2 := make(chan string, 2)

	ch2 <-"A"
}