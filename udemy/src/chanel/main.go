package main

import "fmt"

/*
* チャネル・・・複数のgoroutin間でのデータ受け渡し
* データの送受信を行うデータ構造
* 送信専用・受信専用と明示的な宣言ができる ※明示的な宣言を行わない場合は送受信どちらもできるものとして扱われる
* 受信専用 var <変数名> <-chan <型>
* 送信専用 var <変数名 chan<- <型>
* データはキューのような構造をしており、先入れ・先出しの順序保障がされていることが特徴
 */
func main() {
	// var ch <-chan int // 受信専用
	// var ch chan<- int // 送信専用

	var ch1 chan int
	// チャネルの生成と初期化を行い、データの読み書きができる状態になる。
	ch1 = make(chan int)

	ch2 := make(chan int) // 直接make関数を使って宣言もできる

	fmt.Println(cap(ch1))
	fmt.Println(cap(ch2))

	// バッファサイズを指定して生成
	ch3 := make(chan int, 5)
	fmt.Println(cap(ch3))

	// データをチャネルに送信
	ch3 <- 1
	fmt.Println(len(ch3))
	ch3 <- 2
	ch3 <- 3
	fmt.Println("len", len(ch3))

	// チャネルからデータを受信 ※データを送る度にチャネルからデータがなくなる
	i := <-ch3
	fmt.Println(i)
	fmt.Println("len", len(ch3))
	i2 := <-ch3
	fmt.Println(i2)
	fmt.Println("len", len(ch3))
	fmt.Println(<-ch3)
	fmt.Println("len", len(ch3))

	// バッファ量を超えたデータ量の場合
	ch3 <- 1
	fmt.Println(<-ch3) // 途中でデータを取り出すと、デッドロックにはならない。
	ch3 <- 2
	ch3 <- 3
	ch3 <- 4
	ch3 <- 5
	ch3 <- 6 // デッドロックになる。※実行時エラー

}
