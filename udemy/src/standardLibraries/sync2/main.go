package main

import (
	"fmt"
	"sync"
)

func main() {
	// sync.WaitGroupを生成
	wg := new(sync.WaitGroup)

	// 待ち受けるgoroutineの数は3
	wg.Add(3)

	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("1st Goroutine")
		}
		wg.Done() // 完了
	}()

	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("2nd Goroutine")
		}
		wg.Done() // 完了
	}()

	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("3rd Goroutine")
		}
		wg.Done() // 完了
	}()

	// goroutineの終了を待ち受ける。wg,Done()が3つ完了するまで待つ。
	wg.Wait()
}
