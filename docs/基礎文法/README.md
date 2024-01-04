# Go言語基礎文法

## 変数
### 明示的な変数定義
```go
// var 変数名 型 = 値
var i int = 100

// ()を用いた一括定義
var (
    i2 int    = 200
    s2 string = "Golang"
)

// 変数の型のみを定義（値は定義しない）※各型の初期値が設定される。（intは0, stringは空文字 ""が設定）
var i3 int
var s3 string
```

### 暗黙的な変数定義
```go
// 変数名 := 値
i4 := 400
i4 = 450 // 再代入時には「:」は不要
i4 = 500
```

## 配列
### append/make/len/cap
* append・・・配列の最後の要素に追加する
* make	・・・各型の初期値で配列を生成する
* len	・・・配列の持つ要素数を調べる
* cap	・・・配列の持つ容量（メモリ）を調べる ⇒メモリ使用量を気にする開発時によく利用する。（性能面）

```go
	sl := []int{100, 200}
	fmt.Println(sl)

	// append
	fmt.Println("appendで300を追加")
	sl = append(sl, 300)
	fmt.Println(sl) // [100, 200, 300]
	fmt.Println("appendで400, 500を追加")
	sl = append(sl, 400, 500)
	fmt.Println(sl) // [100, 200, 300, 400, 500]

	// make
	fmt.Println("makeでint型の配列を生成")
	sl2 := make([]int, 5)
	fmt.Println(sl2) // [0, 0, 0, 0, 0]

	// len
	fmt.Println("lenで要素数を検索")
	fmt.Println(len(sl2)) // 5

	// cap
	fmt.Println("capで配列の容量を調べる")
	fmt.Println(cap(sl2)) // 5

	sl3 := make([]int, 5, 10) // 要素数5、キャパシティ10の初期値配列を生成
	fmt.Println(len(sl3)) // 5
	fmt.Println(cap(sl3)) // 10

	sl3 = append(sl3, 1, 2, 3, 4, 5, 6, 7)
	fmt.Println(len(sl3)) // 12
	fmt.Println(cap(sl3)) // 20 
    // 注意)7個要素が追加されたので17ではなく、20になっている※補足参照
```

### キャパシティ（容量）の補足
容量以上の要素が追加されると、メモリの消費量が倍になってしまう。  
パフォーマンス（実行速度）が低下するので、注意すること。

## チャネル
* チャネル・・・複数のgoroutin間でのデータ受け渡し
* データの送受信を行うデータ構造
* 送信専用・受信専用と明示的な宣言ができる ※明示的な宣言を行わない場合は送受信どちらもできるものとして扱われる
* 受信専用 var <変数名> <-chan <型>
* 送信専用 var <変数名 chan<- <型>
* データはキューのような構造をしており、先入れ・先出しの順序保障がされていることが特徴

```go
	var ch1 chan int
	// チャネルの生成と初期化を行い、データの読み書きができる状態になる。
	ch1 = make(chan int)

	ch2 := make(chan int) // 直接make関数を使って宣言もできる

	fmt.Println(cap(ch1)) // キャパシティ0
	fmt.Println(cap(ch2)) // キャパシティ0

	// バッファサイズを指定して生成
	ch3 := make(chan int, 5)
	fmt.Println(cap(ch3)) // キャパシティ5

	// データをチャネルに送信
	ch3 <- 1
	fmt.Println(len(ch3)) // 要素数1
	ch3 <- 2
	ch3 <- 3
	fmt.Println("len", len(ch3)) // 要素数3

	// チャネルからデータを受信 ※データを送る度にチャネルからデータがなくなる
	i := <-ch3
	fmt.Println(i) // 1 ※チャネルの先頭の1を取得し変数格納
	fmt.Println("len", len(ch3)) // 要素数2 ※3つから1つなくなったため。
	i2 := <-ch3
	fmt.Println(i2) // 2 ※チャネルの先頭の2を取得し変数格納
	fmt.Println("len", len(ch3)) // 要素数1 ※2つから1つなくなったため。
	fmt.Println(<-ch3) // 3 ※現在のチャネルに残ったもの
	fmt.Println("len", len(ch3)) // 1 ※これまでチャネルから2つとったので

	// バッファ量を超えたデータ量の場合
	ch3 <- 1
	fmt.Println(<-ch3) // 途中でデータを取り出すと、デッドロックにはならない。
	ch3 <- 2
	ch3 <- 3
	ch3 <- 4
	ch3 <- 5
	ch3 <- 6 // デッドロックになる。※実行時エラー
```

### チャネルの利用例（ChatGPT）
- **非同期処理の同期化**
  - チャネルを使うと、一つのゴルーチンが別のゴルーチンの結果を待つことができます。例えば、複数のWebリクエストを並行して行い、それぞれの応答を合成する際に役立ちます。
- **データストリームの処理**
  - リアルタイムで生成されるデータ（例えば、ログファイルやストリーミングデータ）を処理する際に、一つのゴルーチンがデータを読み取り、別のゴルーチンがそれを処理するといった形で使用されます。
- **タスクの分散**
  - 複数のタスクを複数のゴルーチンに分散して処理する際に、タスクの結果を集約するためにチャネルが使われます。例えば、大量のデータを処理する際に、データを小さなチャンクに分割し、各チャンクを異なるゴルーチンで処理した後、その結果をチャネルを通じて集約します。
- **イベント駆動型プログラミング**
  - GUIアプリケーションやゲーム開発などで、イベント（例えば、ボタンクリックやキーボードイベント）を監視するゴルーチンがあり、そのイベントに基づいて特定のアクションを起こすときに使用されます。
- **リソースの制御**
  - チャネルはリソースへのアクセスを制御するためにも使用されます。例えば、データベース接続プールで、限られた数の接続を複数のゴルーチン間で共有する際に役立ちます。

### ゴルーチン
平行処理を行うための記述

**基本構文**
```go
func main() { 
    go 関数名(引数) 
}
```

[ゴルーチンのわかりやすい記事](https://zenn.dev/farstep/articles/f712e05bd6ff9d)
```go
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
```

### close
チャネルはクローズという状態を持つことができ、明示的に停止させることができる。

チャネルがクローズされた後にデータを送信するとRunTimeErrorになる。

※クローズ後も受信はできる。

```go
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
```

### Sync.Wait
非同期処理をする際に、タスクAが終わらないとタスクBができないような処理の場合には、処理を待たせることも可能。
詳細は[こちら](https://kirohi.com/%E3%80%90go%E5%85%A5%E9%96%80%E3%80%91%E3%82%B4%E3%83%AB%E3%83%BC%E3%83%81%E3%83%B3%E3%81%AB%E3%81%A4%E3%81%84%E3%81%A6%E3%81%AE%E7%B0%A1%E5%8D%98%E3%81%AA%E8%A7%A3%E8%AA%AC)

### チャネル-select
複数チャネルを操作する場合に利用する文法。  
select内でどのcase文に処理が移行するのかはランダムとなっている。

```go
	ch1 := make(chan int, 2)
	ch2 := make(chan string, 2)

	ch1 <- 1
	ch2 <- "A"
	ch1 <- 2
	ch2 <- "B"

	/**
	* どのチャネルを実施するかを選択する
	 */
	// どのケース式が実行されるかはランダム
	// 何かしらの値があればdefaultには入らない
	select {
	case v1 := <-ch1:
		fmt.Println(v1 + 1000)
	case v2 := <-ch2:
		fmt.Println(v2 + "!?")
	default:
		fmt.Println("どちらでもない")
	}
```

エラーとなる例
```go
	ch1 := make(chan int, 2)
	ch2 := make(chan string, 2)

	ch1 <- 1
	
	v1 := <-ch1
	fmt.Println(v1 + 1000)
	v2 := <-ch2
	fmt.Println(v2 + "!?") // ch2に値がないのでエラー

	/**
	* どのチャネルを実施するかを選択する
	 */
	// どのケース式が実行されるかはランダム
	// 何かしらの値があればdefaultには入らない
	select {
	case v1 := <-ch1:
		fmt.Println(v1 + 1000)
	case v2 := <-ch2:
		fmt.Println(v2 + "!?")
	default:
		fmt.Println("どちらでもない")
	}
```

## ポインタ
変数の一種で、メモリアドレスを格納する変数のこと。  
変数のメモリ位置に直接アクセスしてその値を変更する方法を提供する。

関数間で値を受け渡し、メモリを動的に管理し、連結リストや木構造のような複雑なデータ構造を作成するために使用する。

**ポインタ型の宣言**
ポインタ型は型の前に `*` をつけて宣言する。
```go
var p *int
```

**ポインタ型の取得**
変数からポインタを取得する場合は変数の型の前に `&` をつけて宣言する。
```go
package main

import "fmt"

func main() {
    a := 42
    b := &a
    fmt.Println(a, b)
    // 42 0xc00001c030
}
```

その他の解説は[こちら](https://recursionist.io/learn/languages/go/complex/pointer)を参照のこと。

特に構造体のフィールドをセットする際などは意識しないと更新されない状態になるので意識すること。