# 標準パッケージ

## 目次
- [標準パッケージ](#標準パッケージ)
	- [目次](#目次)
	- [公式ドキュメントリンク集](#公式ドキュメントリンク集)
	- [os](#os)
		- [os.Exit](#osexit)
		- [注意点](#注意点)
		- [os.Open](#osopen)
		- [os.Args](#osargs)
		- [ファイル操作](#ファイル操作)
		- [os.Open](#osopen-1)
		- [os.Create](#oscreate)
		- [ファイルの読み込み](#ファイルの読み込み)
		- [オープンファイル](#オープンファイル)
	- [time](#time)
		- [基本](#基本)
		- [文字列から時間を生成](#文字列から時間を生成)
		- [X時間後を出力](#x時間後を出力)
		- [時間の比較](#時間の比較)
		- [sleep](#sleep)
	- [math](#math)
	- [rand](#rand)
	- [flag](#flag)
	- [fmt](#fmt)
		- [デフォルトフォーマット](#デフォルトフォーマット)
		- [%+vフラグ付きデフォルトフォーマット](#vフラグ付きデフォルトフォーマット)
		- [%#v Go言語の構文で表現](#v-go言語の構文で表現)
		- [%T 値の型](#t-値の型)
		- [%t 単語、trueまたはfalse](#t-単語trueまたはfalse)
		- [%s 文字列、スライス](#s-文字列スライス)
		- [数値系](#数値系)
		- [ポインタ](#ポインタ)
		- [出力先を指定する](#出力先を指定する)
	- [log](#log)
		- [通常のログ（`log.Print()`）](#通常のログlogprint)
		- [Exitコードを伴ったログ（`log.Fetal()`）](#exitコードを伴ったログlogfetal)
		- [パニックを伴ったログ出力(`log.Panic()`)](#パニックを伴ったログ出力logpanic)
		- [ファイルに出力](#ファイルに出力)
		- [ログのフォーマットを指定](#ログのフォーマットを指定)
		- [ロガーの生成](#ロガーの生成)
	- [strconv](#strconv)
	- [strings](#strings)
	- [bufio](#bufio)
	- [ioutil](#ioutil)
	- [regexp](#regexp)
		- [正規表現のフラグ](#正規表現のフラグ)
		- [幅を持たない正規表現のパターン](#幅を持たない正規表現のパターン)
		- [繰り返しを表す正規表現](#繰り返しを表す正規表現)
		- [正規表現の文字クラス](#正規表現の文字クラス)
		- [正規表現文字グループ](#正規表現文字グループ)
		- [正規表現にマッチした文字列の取得](#正規表現にマッチした文字列の取得)
		- [正規表現のグループによるサブマッチ](#正規表現のグループによるサブマッチ)
		- [正規表現による文字列の置換](#正規表現による文字列の置換)
		- [正規表現による文字列の分割](#正規表現による文字列の分割)
	- [sync](#sync)
		- [非同期処理の排他制御(sync.Mutex)](#非同期処理の排他制御syncmutex)
		- [goroutineの終了を待ち受ける(sync.WaitGroup)](#goroutineの終了を待ち受けるsyncwaitgroup)
		- [json形式への変換時の補足](#json形式への変換時の補足)
		- [omitempty](#omitempty)
		- [jsonからの変換（デコード）](#jsonからの変換デコード)
		- [マーシャルのカスタム](#マーシャルのカスタム)
		- [アンマーシャルのカスタム](#アンマーシャルのカスタム)
	- [sort](#sort)
		- [Ints / Strings](#ints--strings)
		- [Slice / SliceStable](#slice--slicestable)
		- [カスタムソート](#カスタムソート)
	- [context](#context)
	- [net/url](#neturl)
	- [net/http client](#nethttp-client)
	- [net/http server](#nethttp-server)
		- [サーバーを起動する](#サーバーを起動する)
		- [ハンドラーを定義して画面に文字列を表示](#ハンドラーを定義して画面に文字列を表示)
		- [マルチプレクサの設定](#マルチプレクサの設定)


## 公式ドキュメントリンク集
| No  | パッケージ名    | 説明                                                                                                                                                                      | リンク                       |
| --- | --------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------- |
| 1   | os              | オペレーティング システム機能へのプラットフォームに依存しないインターフェイスを提供                                                                                       | https://pkg.go.dev/os        |
| 2   | time            | パッケージ時間には、時間を測定および表示する機能                                                                                                                          | https://pkg.go.dev/time      |
| 3   | math            | 基本的な定数と数学関数を提供                                                                                                                                              | https://pkg.go.dev/math      |
| 4   | rand            | シミュレーションなどのタスクに適した擬似乱数生成器を実装。セキュリティが重要な作業には使用しない                                                                          | https://pkg.go.dev/math/rand |
| 5   | flag            | コマンドライン フラグ解析を実装                                                                                                                                           | https://pkg.go.dev/flag      |
| 6   | fmt             | フォーマットされた I/O を実装                                                                                                                                             | https://pkg.go.dev/fmt       |
| 7   | log             | 単純なログ パッケージを実装                                                                                                                                               | https://pkg.go.dev/log       |
| 8   | strconv         | 基本データ型の文字列表現との変換を実装                                                                                                                                    | https://pkg.go.dev/strconv   |
| 9   | strings         | UTF-8 でエンコードされた文字列を操作するための単純な関数を実装                                                                                                            | https://pkg.go.dev/strings   |
| 10  | bufio           | バッファリングされた I/O を実装                                                                                                                                           | https://pkg.go.dev/bufio     |
| 11  | loutil          | I/O ユーティリティ関数を実装。非推奨: Go 1.16 では、同じ機能がパッケージioまたはパッケージosによって提供されている。                                                      | https://pkg.go.dev/io/ioutil |
| 12  | regexp          | 正規表現検索を実装                                                                                                                                                        | https://pkg.go.dev/regexp    |
| 13  | sync            | 相互排他ロックなどの基本的な同期プリミティブを提供。このパッケージで定義されている型を含む値はコピーしない。                                                              | https://pkg.go.dev/sync      |
| 14  | crypto          | 共通の暗号定数を収集                                                                                                                                                      | https://pkg.go.dev/crypto    |
| 15  | json            | RFC 7159で定義されている JSON のエンコードとデコードを実装                                                                                                                | https://pkg.go.dev/json      |
| 16  | sort            | スライスとユーザー定義のコレクションをソートするためのプリミティブを提供                                                                                                  | https://pkg.go.dev/sort      |
| 17  | context         | 期限、キャンセル シグナル、および API 境界を越えてプロセス間でリクエストをスコープとするその他の値を伝達する Context タイプを定義<br>主にタイムアウトを設定する際に用いる | https://pkg.go.dev/context   |
| 18  | net/url         | HTTP クライアントとサーバーの実装を提供                                                                                                                                   | https://pkg.go.dev/net       |
| 19  | net/http client | HTTP クライアントとサーバーの実装を提供                                                                                                                                   | https://pkg.go.dev/net       |
| 20  | net/http server | HTTP クライアントとサーバーの実装を提供                                                                                                                                   | https://pkg.go.dev/net       |


## os
### os.Exit
osパッケージを利用すると任意のタイミングで処理を終了し、Exitステータスを指定できる。

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	// Exit
	os.Exit(1)
	fmt.Println("Start")
}
```

### 注意点
defer文（遅延処理）も破棄されて呼び出されないため注意
```go
package main

import (
	"fmt"
	"os"
)

func main() {
	defer func() {
		fmt.Println("defer") // 処理が終わっているため呼び出されない
	}()
	os.Exit(0)
}
```

### os.Open
ファイル操作時に利用する

```go
package main

import (
	"log"
	"os"
)

func main() {
	_, err := os.Open("A.txt") // 存在しないA.txt
	if err != nil {
		log.Fatalln(err)
	}
}
```

実行結果
```bash
go run main.go

2024/01/05 13:05:18 open A.txt: no such file or directory
exit status 1
```

### os.Args
実行ファイル実行時に指定した引数を受け取り、処理をする。

下記のファイルを用意する。

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	// Args
	fmt.Println(os.Args[0])
	fmt.Println(os.Args[1])
	fmt.Println(os.Args[2])
	fmt.Println(os.Args[3])

	// os.Argsの要素数を表示
	fmt.Printf("length=%d\n", len(os.Args))

	// os.Argsの中身を全て表示
	for i, v := range os.Args {
		fmt.Println(i, v)
	}
}
```

ビルドする。（main.goをgetargsという名前の実行ファイルにする）
```bash
go build -o getargs main.go
```

実行する
```bash
./getargs 123 456 789

./getargs
123
456
789
length=4
0 ./getargs
1 123
2 456
3 789
```

※index[0]は実行ファイル名が入る。
※index[1]-index[3]に指定した引数が入る。

### ファイル操作
### os.Open
ファイルを開く
```go
package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("test.txt")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(f.Name()) // ファイル名を出力 test.txt

	defer f.Close()
}
```

### os.Create
ファイルの作成

```go
package main

import (
	"os"
)

func main() {
	f, _ := os.Create("foo.txt") // ファイル作成

	f.Write([]byte("Hello\n")) // Helloをバイトに変換し、書き込み（改行コード付き）

	f.WriteAt([]byte("Golang"), 6) // Golangをバイトに変換し、書き込み

	f.Seek(0, os.SEEK_END) // Seekでファイルの末尾に移動

	f.WriteString("Yaah") // 末尾に移動後、Yaahを書き込み
}
```

実行後、作成されたファイルfoo.txt
```text
Hello
GolangYaah
```

### ファイルの読み込み
```go
package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("foo.txt")
	if err != nil {
		log.Fatalln(err)
	}

	defer f.Close()

	bs := make([]byte, 128) // バイトのスライスを作成

	n, err := f.Read(bs) // 作成したバイトのスライスを読み込み
	fmt.Println(n)
	fmt.Println(string(bs)) // バイトのスライスを文字列に変換して出力

	bs2 := make([]byte, 128)

	nn, err := f.ReadAt(bs2, 10)
	fmt.Println(nn)
	fmt.Println(string(bs2))
}
```

実行結果
```bash
go run main.go

16
Hello
GolangYaah
6
ngYaah
```

### オープンファイル
読み込んだファイルの操作

オープンファイルオプション
- O_RDONLY
  - 読み込み専用
- O_WRONLY
  - 書き込み専用
- O_RDWR
  - 読み書き可能
- O_APPEND
  - ファイル末尾に追記
- O_CREATE
  - ファイルが無ければ作成
- O_TRUNC
  - 可能であればファイルの内容をオープン時に空にする

オプションは`os.O_RDWR|os.O_CREATE`のように複数付与することも可能。  
例）読み書き可能でなければ新しく作成する。

```go
package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
    // 引数:ファイル名、オプション、権限
	f, err := os.OpenFile("foo.txt", os.O_RDONLY, 0666)
	if err != nil {
		log.Fatalln(err)
	}

	defer f.Close()

	bs := make([]byte, 128)
	n, err := f.Read(bs)
	fmt.Println(n)
	fmt.Println(string(bs))
}

```


実行結果
```
go run main.go

16
Hello
GolangYaah
```

## time
日時に関するパッケージ

### 基本
```go
package main

import (
	"fmt"
	"time"
)

func main() {

	t := time.Now() // 現在時刻の取得
	fmt.Println(t)

	t2 := time.Date(2020, 6, 10, 10, 10, 10, 0, time.Local) // 2020年6月10日 10:10:10
	fmt.Println(t2)
	fmt.Println(t2.Year())
	fmt.Println(t2.YearDay())
	fmt.Println(t2.Month())
	fmt.Println(t2.Weekday())
	fmt.Println(t2.Day())
	fmt.Println(t2.Hour())
	fmt.Println(t2.Minute())
	fmt.Println(t2.Second())
	fmt.Println(t2.Nanosecond())
	fmt.Println(t2.Zone())
}
```

### 文字列から時間を生成
```go
package main

import (
	"fmt"
	"time"
)

func main() {
	d, _ := time.ParseDuration("2h30m")
	fmt.Println(d)
}
```

### X時間後を出力
```go
package main

import (
	"fmt"
	"time"
)

func main() {
	t3 := time.Now() // 現在時刻
	t3 = t3.Add(2*time.Minute + 15*time.Second) // 現在時刻の2時間30分後を表すtime.Time型
	fmt.Println(t3)
}
```

### 時間の比較
```go
package main

import (
	"fmt"
	"time"
)

func main() {
	t5 := time.Date(2021, 7, 24, 0, 0, 0, 0, time.Local)
	t6 := time.Now()
	fmt.Println(t6)

	d2 := t5.Sub(t6) // 時刻の差分を取得
	fmt.Println(d2)

	fmt.Println(t6.Before(t5)) // false
	fmt.Println(t6.After(t5)) // true
	fmt.Println(t5.Before(t6)) // true
	fmt.Println(t5.After(t6)) // false
}
```

### sleep
```go
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("スタート")
	time.Sleep(5 * time.Second)
	fmt.Println("5秒間停止後")
}
```

## math
数学に関連するパッケージ

円周率・平方根・数値型に関するような内容が含まれている。
```go
package main

import (
	"fmt"
	"math"
)

func main() {
	// 円周率
	fmt.Println(math.Pi)

	// 2の平方根
	fmt.Println(math.Sqrt2)

	// 立方根
	fmt.Println(math.Cbrt(8))

	// 数値型に関する定数
	fmt.Println(math.MaxFloat32) // float32で表現可能な最大値
	fmt.Println(math.SmallestNonzeroFloat32) // float32で表現可能な最小値（0ではない）
	fmt.Println(math.MaxFloat64) // float64で表現可能な最大値
	fmt.Println(math.SmallestNonzeroFloat64) // float64で表現可能な最小値（0ではない）
	fmt.Println(math.MaxInt64) // Int64で表現可能な最大値
	fmt.Println(math.MinInt64) // Int64で表現可能な最小値（0ではない）

	// 絶対値
	fmt.Println(math.Abs(100))
	fmt.Println(math.Abs(-100))

	// 累乗
	fmt.Println(math.Pow(0, 2))
	fmt.Println(math.Pow(2, 2))

	// 最大値、最小値
	fmt.Println(math.Max(1, 2))
	fmt.Println(math.Min(1, 2))

	// 小数点以下の切り捨て、切り上げ

	// 数値の正負を問わず、小数点以下を単純に切り捨てる
	fmt.Println(math.Trunc(3.5))
	fmt.Println(math.Trunc(-3.5))
	
	// 引数の数値より小さい最大の整数を返す
	fmt.Println(math.Floor(3.5))
	fmt.Println(math.Floor(-3.5))

	// 引数の数値より大きい最小の整数を返す
	fmt.Println(math.Ceil(3.5))
	fmt.Println(math.Ceil(-3.5))
}
```

## rand
疑似乱数を生成するパッケージ

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(256)

	fmt.Println(rand.Float64())
	fmt.Println(rand.Float64())
	fmt.Println(rand.Float64())

	// 現在時刻をシードに使った疑似乱数生成
	fmt.Println(time.Now().UnixNano())
	rand.Seed(time.Now().UnixNano())

	// 0-99の疑似乱数
	fmt.Println(rand.Intn(100))

	// int型の疑似乱数
	fmt.Println(rand.Int())
	// int32型の疑似乱数
	fmt.Println(rand.Int31())
	// int64型の疑似乱数
	fmt.Println(rand.Int63())
	// uint32型の疑似乱数
	fmt.Println(rand.Uint32())
}
```
現在時刻を用いた乱数生成が一番手軽。

## flag
コマンドラインからプログラムに与えられた引数やオプションを効率的に処理する際に用いるパッケージ  
シンプルなコマンドを作成する際に利用されるパッケージ

```go
package main

import (
	"flag"
	"fmt"
)

/**
* コマンドラインを処理するサンプル
* go run main.go -n 20 -m message -x
 */
func main() {

	var (
		max int
		msg string
		x bool
	)

	// IntVar 整数のオプション
	flag.IntVar(&max, "n", 32, "処理数の最大値")

	// StringVar 文字列のオプション
	flag.StringVar(&msg, "m", "", "処理メッセージ")

	// BoolVar bool型のオプション コマンドラインに与えられたらtrue なければfalse
	flag.BoolVar(&x, "x", false, "拡張オプション")

	// コマンドラインをパース
	flag.Parse()

	fmt.Println("処理数の最大値 =", max)
	fmt.Println("処理メッセージ =", msg)
	fmt.Println("拡張オプション =", x)
}
```

処理結果
```bash
go run main.go -n 20 -m message -x

処理数の最大値 = 20
処理メッセージ = message
拡張オプション = true
```

引数全くなし処理結果
```bash
go run main.go

処理数の最大値 = 32
処理メッセージ = 
拡張オプション = false
```
デフォルトで付与している処理結果が出力される。

存在しないオプションを渡した場合はヘルプメッセージが出力される仕組みとなっている。
```bash
go run main.go -z

flag provided but not defined: -z
Usage of /tmp/go-build1517213371/b001/exe/main:
  -m string
        処理メッセージ
  -n int
        処理数の最大値 (default 32)
  -x    拡張オプション
exit status 2
```

## fmt
標準出力に用いるパッケージ

Goでは書式指定子 %... のことを verb と表記している。
- 論理値(bool)：%t
- 符号付き整数(int, int8など)：%d
- 符号なし整数(uint, uint8など) ：%d
- 浮動小数点数(float64など) ：%g
- 複素数(complex128など)：%g
- 文字列(string)：%s
- チャネル(chan)：%p
- ポインタ(pointer)：%p

### デフォルトフォーマット
```go
package main

import (
    "fmt"
)

type Person struct {
    Name     string
    Favorite string
}

func main() {
    p := new(Person)
    p.Name = "sekky"
    p.Favorite = "Programming"
    fmt.Printf("構造体 = %v", p) // 構造体 = &{sekky Programming}
}
```

### %+vフラグ付きデフォルトフォーマット
フラグ（%v）を付け加えるとフィールド名が表示される
```go
package main

import (
    "fmt"
)

type Person struct {
    Name     string
    Favorite string
}

func main() {
    p := new(Person)
    p.Name = "sekky"
    p.Favorite = "Programming"
    fmt.Printf("構造体 = %+v", p) // 構造体 = &{Name:sekky Favorite:Programming}
}
```

### %#v Go言語の構文で表現

```go
package main

import (
    "fmt"
)

type Person struct {
    Name     string
    Favorite string
}

func main() {
    p := new(Person)
    p.Name = "sekky"
    p.Favorite = "Programming"
    fmt.Printf("構造体 = %#v\n", p)

    a := [...]string{"Go", "Java", "Python", "Ruby", "Rust"} // 構造体 = &main.Person{Name:"sekky", Favorite:"Programming"}
    fmt.Printf("配列 = %#v", a) // 配列 = [5]string{"Go", "Java", "Python", "Ruby", "Rust"}
}
```
### %T 値の型
`%T`を指定することにより、型を出力できるようになる。

```go
package main

import (
    "fmt"
)

type Person struct {
    Name     string
    Favorite string
}

func main() {
    p := new(Person)
    p.Name = "sekky"
    p.Favorite = "Programming"
    fmt.Printf("構造体 = %T\n", p) // 構造体 = *main.Person

    a := [...]string{"Go", "Java", "Python", "Ruby", "Rust"}
    fmt.Printf("配列 = %T\n", a) // 配列 = [5]string

    s := []int{0, 1, 2}
    fmt.Printf("スライス = %T\n", s) // スライス = []int

    n := 7
    fmt.Printf("int = %T\n", n) // int = int

    str := "sekky"
    fmt.Printf("string = %T\n", str) // string = string

    b := true
    fmt.Printf("bool = %T\n", b) // bool = bool
}
```

### %t 単語、trueまたはfalse
```go
package main

import (
    "fmt"
)

type Person struct {
    Name     string
    Favorite string
}

func main() {

    b := true
    fmt.Printf("%t\n", b) // true

    b = false
    fmt.Printf("%t\n", b) // false

    str := "Go"
    fmt.Printf("%t\n", str) // %!t(string=Go)

    n := 1
    fmt.Printf("%t\n", n) // %!t(int=1)

    p := new(Person)
    p.Name = "sekky"
    fmt.Printf("%t\n", p) // &{%!t(string=sekky) %!t(string=)}
}
```

### %s 文字列、スライス
```go
package main

import (
    "fmt"
)

func main() {

    str := "Go"
    fmt.Printf("文字列 = %s\n", str) // 文字列 = Go

    s := []string{"Go", "Java"}
    fmt.Printf("スライス = %s", s) // スライス = [Go Java]
}
```

### 数値系
```go
package main

import (
    "fmt"
)

func main() {

    // %b 基数2
    b := 5
    fmt.Printf("基数2 = %b\n", b) // 基数2 = 101

    // %c 対応するUnicodeコードポイントによって表される文字
    c := '\101'
    fmt.Printf("Unicodeコードポイント = %c\n", c) // Unicodeコードポイント = A

    // %d 基数10
    d := 5
    fmt.Printf("基数10 = %d\n", d) // 基数10 = 5

    // %o 基数8
    o := 20
    fmt.Printf("基数8 = %o\n", o) // 基数8 = 24

    // %x 基数16、10以上の数には小文字(a-f)を使用
    x := 2059
    fmt.Printf("基数16、10以上の数には小文字(a-f)を使用 = %x\n", x) // 基数16、10以上の数には小文字(a-f)を使用 = 80b

    // %X 基数16、10以上の数には大文字(A-F)を使用
    X := 2059
    fmt.Printf("基数16、10以上の数には大文字(A-F)を使用 = %X\n", X) // 基数16、10以上の数には大文字(A-F)を使用 = 80B

    // %U ユニコードフォーマット: U+1234; "U+%x"と同じ。デフォルトは、4桁
    U := '\101'
    fmt.Printf("ユニコードフォーマット = %U\n", U) // ユニコードフォーマット = U+0041
}
```

### ポインタ
```go
package main

import (
    "fmt"
)

type Person struct {
    Name     string
    Favorite string
}

func main() {
    p := new(Person)
    p.Name = "sekky"
    p.Favorite = "Programming"
    fmt.Printf("構造体 = %p", p) // 構造体 = 0x1040a130
}
```

### 出力先を指定する
```go
package main

import (
	"fmt"
	"os"
)

func main() {
	// 標準出力
	fmt.Fprint(os.Stdout, "Hello")
	fmt.Fprint(os.Stdout, "Hello")
	fmt.Fprint(os.Stdout, "Hello")

	f, _ := os.Create("test.txt") // test.txtを作成
	defer f.Close() // 遅延処理、最後にファイルをクローズ処理

	fmt.Fprint(f, "Fprint") // test.txtに"Fprint"という文字列を書き込む
}
```

## log
シンプルなログを出力するためのパッケージ  
日付と時刻とメッセージ内容が出力される(形式：`YYYY/MM/DD HH:mm:dd <message>`)

### 通常のログ（`log.Print()`）
```go
package main

import (
	"log"
	"os"
)

func main() {
	// ログの出力先を変更
	log.SetOutput(os.Stdout)

	log.Print("Log\n")
	log.Println("Log2")
	log.Printf("Log%d\n", 3)
}
```

結果
```bash
go run main.go

2024/01/05 18:37:07 Log
2024/01/05 18:37:07 Log2
2024/01/05 18:37:07 Log3
```

### Exitコードを伴ったログ（`log.Fetal()`）
```go
package main

import (
	"log"
	"os"
)

func main() {
	// ログの出力先を変更
	log.SetOutput(os.Stdout)

	// log.Print("Log\n")
	// log.Print("Log2")
	// log.Printf("Log%d\n", 3)

	log.Fatal("Log\n")
	log.Fatalln("Log2")
	log.Fatalf("Log%d\n", 3)
}
```

結果
```bash
go run main.go

2024/01/05 18:38:33 Log
exit status 1
```
１回目のFetalログを出力し、exitコードが出力され、他のログは出ない。

### パニックを伴ったログ出力(`log.Panic()`)
```go
package main

import (
	"log"
	"os"
)

func main() {
	// ログの出力先を変更
	log.SetOutput(os.Stdout)

	log.Panic("Log\n")
	log.Panicln("Log2")
	log.Panicf("Log%d\n", 3)
}
```

結果
```bash
go run main.go

2024/01/05 18:39:59 Log
panic: Log


goroutine 1 [running]:
log.Panic({0xc000104f60?, 0x4c3678?, 0xc000012018?})
        /usr/local/go/src/log/log.go:384 +0x65
main.main()
        /home/yk-you0922/go/src/sample-go-lang/udemy/src/standardLibraries/log/main.go:20 +0x65
exit status 2
```
１回目のログだけを出力し、パニックとしてプログラムが終了する。

### ファイルに出力
```go
package main

import (
	"log"
	"os"
)

func main() {
	f, err := os.Create("test.log") // ログ出力ファイルの作成
	if err != nil {
		return // ファイルが作成できなければ早期リターン
	}

	log.SetOutput(f) // 出力先をファイルに設定
	log.Println("ファイルに書き込む") // ファイルにログを書き込む
}
```

結果(test.log)
```log
2024/01/05 18:44:28 ファイルに書き込む
```

### ログのフォーマットを指定
```go
package main

import (
	"log"
	"os"
)

func main() {
	// ログの出力先を変更
	log.SetOutput(os.Stdout)

	// 標準のログフォーマット
	log.SetFlags(log.LstdFlags)
	log.Println("A") // 2024/01/05 18:48:36 A

	// マイクロ秒を追加
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)
	log.Println("B") // 2024/01/05 18:48:36.848490 B

	// 時刻とファイルの行番号（短縮系）
	log.SetFlags(log.Ltime | log.Lshortfile)
	log.Println("C") // 18:48:36 main.go:22: C

	log.SetFlags(log.LstdFlags)

	// ログのプリフィックスを設定
	log.SetPrefix("[LOG]")
	log.Println("E") // [LOG]2024/01/05 18:48:36 E
}
```

### ロガーの生成
Goのlogパッケージはデフォルトで設定されている1つの設定を全体に適応しているため、小回りが利かないようになっているため  
ロガーを生成し、個別に設定を付与することができる。

```go
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
		logger.Fatalln("Exit, err")
	}
}

```

出力結果
```bash
go run main.go

2024/01/05 19:00:36 main.go:11: message
2024/01/05 19:00:36 message
2024/01/05 19:00:36 main.go:18: Exit, err
exit status 1
```
ロガーでファイルの行番号まで指定しているので、行番号まで出力された。

## strconv
基本データ型の文字列表現との変換を実装するパッケージ  
参考コード：[strconv-sample](../../../udemy/src/standardLibraries/strconv/main.go)

## strings
UTF-8 でエンコードされた文字列を操作するための単純な関数を実装  
参考コード：[strconv-sample](../../../udemy/src/standardLibraries/strings/main.go)

- 文字列結合
  - strings.Join()
- 文字列を繰り返して結合
  - strings.Repeat()
- 文字列検索
  - strings.Index()
  - strings.LastIndex()
  - strings.IndexAny()
  - strings.LastIndexAny()
- 文字列の置換
  - strings.Replace()
  - strings.ReplaceAll()
- 文字列を分割
  - strings.Split()
  - strings.SplitAfter()
  - strings.SplitN()
  - strings.SplitAfterN()
- 大文字、小文字の変換
  - strings.ToLower()
  - strings.ToUpper()
- 文字列から空白を取り除く
  - strings.TrimSpace()
- 文字列からスペース区切りのフィールドを取り出す
  - strings.Fields()

## bufio
バッファリングされた I/O を実装。  
io.Reader または io.Writer オブジェクトをラップし、同様にインターフェイスを実装する別のオブジェクト  
 (Reader または Writer) を作成するが、テキスト I/O のバッファリングといくつかのヘルプを提供

```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 標準入力を行単位で読み込む
	// スキャナの生成
	scanner := bufio.NewScanner(os.Stdin)

	// 入力のスキャンが成功する限り繰り返すループ
	for scanner.Scan() {
		// スキャン内容を文字列で出力
		fmt.Println(scanner.Text())
	}

	// スキャンにエラーが発生した場合の処理
	if err := scanner.Err(); err != nil {
		fmt.Println(os.Stderr, "読み込みエラー", err)
	}

}
```

## ioutil
I/Oユーティリティ関数を提供するパッケージ  
ただし、基本的にはosパッケージで代用できるので、あまり利用しない。

サンプルコード
```go
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	// 入力全体を読み込む
	f, _ := os.Open("foo.txt")
	bs, _ := ioutil.ReadAll(f)
	fmt.Println(string(bs))

	// ファイルに書き込む
	if err := ioutil.WriteFile("bar.txt", bs, 0666); err != nil {
		log.Fatalln(err)
	}
}
```

## regexp
正規表現についての機能が実装されたパッケージ

参考コード：[strconv-sample](../../../udemy/src/standardLibraries/regexp/main.go)

- regexp.MatchString()
  - 第一引数に第二引数が含まれることを確認する。
  - 簡易的な正規表現であることに注意。基本的には次のregexp.Compile()を用いてからマッチするかを確認する。
- regexp.Compile()
  - 引数の文字列をコンパイルし、正規表現としてマッチするかを判定するための準備を行う。
  - コンパイル後にマッチするかを確認するのが基本。
- regexp.MustCompile()
  - Compile()と異なり、第二返り値にエラー型を返すのではなく、直接ランタイムエラーとする点が異なる。
  - 後のバグを防ぐために利用されるため、推奨される。
  - 正規表現が動的に利用される場合にはCompile()を利用すること。

### 正規表現のフラグ
- i : 大文字・小文字を区別しない
- m : マルチラインモード（`^`と`&`が文頭、文末に加えて行頭、行末にマッチする）
- s : `.`が`\n`にマッチ
- U : 最小マッチへの変換（`x*`は`x*?`へ、`x+`は`x+?`へ）

### 幅を持たない正規表現のパターン
- ^ : 文頭（`m`フラグが有効な場合は行頭にも）
- $ : 文末（`m`フラグが有効な場合には行末にも）
- \A : 文頭
- \z : 文末
- \b : ASCIIによるワード協会
- \B : 非ASCIIによるワード協会

### 繰り返しを表す正規表現
- x* : 1回以上繰り返すx(最大マッチ)
- x+ : 1回以上繰り返すx(最大マッチ)
- x? : 0回以上1回以下繰り返すx
- x{n, m} : n回以上m回以下繰り返すx(最大マッチ)
- x{n, } : n回以上繰り返すx(最大マッチ)
- x{x} : n回繰り返すx(最大マッチ)
- x*? : 0回以上繰り返すx(最小マッチ)
- x+? : 0回以上1回以下繰り返すx(最小マッチ)
- x?? : 0回以上1回以下繰り返すx(0回優先)
- x{n, m}? : n回以上m回以下繰り返す(最小マッチ)
- x{n, }? : n回以上繰り返すx(最小マッチ)
- x{n}? : n回繰り返すx(最小マッチ)

### 正規表現の文字クラス
例）
- [XYZ] : 任意の文字クラスを定義（XYZのどれか一つでも当てはまるか）
- ^[0-9A-Za-z_]{3}$ : 英数字とアンダースコアを含んだ3文字であるか
- [^0-9A-Za-z_] : 英数字とアンダースコア以外の文字でマッチするか

### 正規表現文字グループ
- (正規表現)グループ（順序によるキャプチャ）
- (?:正規表現)グループ（キャプチャされない）
- (?:P<Name>正規表現)名前付きグループ

### 正規表現にマッチした文字列の取得
参考コード：[strconv-sample](../../../udemy/src/standardLibraries/regexp2/main.go)

- regexp.FindString()
  - はじめにマッチした部分をstring型で返却する関数
  - 第一引数：検査対象文字列
- regexp.FindAllString()
  - マッチした部分を全て取得できる
  - 第一引数：検査対象文字列
  - 第二引数：取得する文字列の最大数（-1指定で文字列のスライス型で全て返却する）

### 正規表現のグループによるサブマッチ
- regexp.FindAllStringSubMatch()
  - グループにマッチした部分をサブマッチとして取得できる関数
  - 第一引数：検査対象文字列
  - 第二引数：取得する要素数の最大数（-1指定で文字列のスライス型で全て返却する）

### 正規表現による文字列の置換
参考コード：[strconv-sample](../../../udemy/src/standardLibraries/regexp3/main.go)

- regexp.ReplaceAllString()
  - 正規表現にマッチした文字列を第二引数の指定文字列に置き換える
  - 第一引数：置き換え対象文字列
  - 第二引数：置き換える文字列

### 正規表現による文字列の分割
- regexp.Split()
  - 正規表現にマッチした文字列を置き換える
  - 第一引数：置き換え対象文字列
  - 第二引数：マッチした箇所を何カ所置き換えるか指定する（-1指定の場合は全て置き換える）

## sync
goの非同期処理や排他処理を支援する機能がまとめられたパッケージ

非同期処理にはgoroutineとチャネルが提供されているが、あらゆる場面で対応できるわけではない。  
そのため、syncパッケージを利用し、対応できる場面を増やしていくことが必要。

### 非同期処理の排他制御(sync.Mutex)

下記の例は非同期処理におけるレースコンディション（競合）についての解決サンプルコードとなる。

まずは競合が起こっている状態のコード

```go
package main

import (
	"fmt"
	"time"
)


var st struct { A, B, C int}

// 引数の値をstのA、B、Cにそれぞれ代入する。
func UpdateAndPrint(n int) {
	st.A = n
	time.Sleep(time.Microsecond)
	st.B = n
	time.Sleep(time.Microsecond)
	st.C = n
	time.Sleep(time.Millisecond)
	fmt.Println(st)
}

func main() {
	for i := 0; i < 5; i++ {
		go func() {
			for i := 0; i < 1000; i++ {
				UpdateAndPrint(i)
			}
		}()
	}

	// 無限ループ。UpdateAndPrint()が終わる前にmain関数自体が処理を終えるため設定している。
	for {

	}

}
```

実行結果
```go
・
・
・
{997 997 998} // UpdateAndPrint()の処理上、全部同じ値が入るはずなのに、入っていない。
{999 997 998} // UpdateAndPrint()の処理上、全部同じ値が入るはずなのに、入っていない。
{997 999 998} // UpdateAndPrint()の処理上、全部同じ値が入るはずなのに、入っていない。
{997 999 999} // UpdateAndPrint()の処理上、全部同じ値が入るはずなのに、入っていない。
{994 997 999} // UpdateAndPrint()の処理上、全部同じ値が入るはずなのに、入っていない。
{998 998 998}
{999 999 998} // UpdateAndPrint()の処理上、全部同じ値が入るはずなのに、入っていない。
{999 999 999}
{998 999 999} // UpdateAndPrint()の処理上、全部同じ値が入るはずなのに、入っていない。
{995 998 999} // UpdateAndPrint()の処理上、全部同じ値が入るはずなのに、入っていない。
{998 998 998}
{996 996 998} // UpdateAndPrint()の処理上、全部同じ値が入るはずなのに、入っていない。
{996 996 996}
{996 996 998} // UpdateAndPrint()の処理上、全部同じ値が入るはずなのに、入っていない。
{999 999 999}
{997 999 999} // UpdateAndPrint()の処理上、全部同じ値が入るはずなのに、入っていない。
{997 997 999} // UpdateAndPrint()の処理上、全部同じ値が入るはずなのに、入っていない。
{997 997 997}
{998 998 998}
{999 999 999}
```

上記のように、全て同じ値が入るはずの関数に対して、異なる値が入っているケースがあり、これが競合状態である。

これを解決するためにsyncパッケージを利用する。（ミューテックスを利用）

```go
package main

import (
	"fmt"
	"sync"
	"time"
)


var st struct { A, B, C int}

// ミューテックスを保持するパッケージ変数（参照型）
var mutex *sync.Mutex

// 引数の値をstのA、B、Cにそれぞれ代入する。
func UpdateAndPrint(n int) {
	// ロック⇒アンロックするまでは一つのgoroutineしか処理ができないようにする。
	mutex.Lock()
	
	st.A = n
	time.Sleep(time.Microsecond)
	st.B = n
	time.Sleep(time.Microsecond)
	st.C = n
	time.Sleep(time.Millisecond)
	fmt.Println(st)

	// アンロック⇒またロックするまでは複数のgoroutineが処理できる。
	mutex.Unlock()
}

func main() {
	mutex = new(sync.Mutex)

	for i := 0; i < 5; i++ {
		go func() {
			for i := 0; i < 1000; i++ {
				UpdateAndPrint(i)
			}
		}()
	}

	// 無限ループ。UpdateAndPrint()が終わる前にmain関数自体が処理を終えるため設定している。
	for {

	}

}
```

syncのミューテックス型には一つのgoroutineのみがロックを取得できる性質があり、既にロックされているミューテックスに対して別のgoroutineが処理できないようにできる。

処理結果
```go
・
・
・
{998 998 998}
{997 997 997}
{997 997 997}
{997 997 997}
{997 997 997}
{999 999 999}
{998 998 998}
{998 998 998}
{998 998 998}
{998 998 998}
{999 999 999}
{999 999 999}
{999 999 999}
{999 999 999}
```

全ての値が一致することを確認。

注意）せっかくの非同期処理の意味がなくなるのでロックする範囲や単位は慎重に検討する必要がある。

### goroutineの終了を待ち受ける(sync.WaitGroup)
下記は3つの非同期処理によるforループを実装している。  
しかし、sync.WaitGroupを利用しない場合、各標準出力処理が終わるよりも前にmain関数が終了するため、  
何一つ処理が実施されずに終了する。

```go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type A struct{}

type User struct {
	ID      int       `json:"id"` // フィールド  型  jsonに変換した際のフィールド名（指定なしも可能）
	Name    string    `json:"name"`
	Email   string    `json:"email"`
	Created time.Time `json:"created"`
	A       A         `json:"A"`
}

func main() {

	u := new(User)
	u.ID = 1
	u.Name = "test"
	u.Email = "test@test.com"
	u.Created = time.Now()

	// Marshal JSONに変換 ⇒ 返り値：byteのスライス
	bs, err := json.Marshal(u)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(bs))
}
```

処理結果
```go
{"id":1,"name":"test","email":"test@test.com","created":"2024-01-06T22:30:25.760522681+09:00","A":{}}
```

### json形式への変換時の補足
idのint型をjson上、文字列として扱いたい場合は下記のように設定する。
```go
type User struct {
	ID      int       `json:"id,string"`
	/** 以下省略 */
}
```

処理結果
```go
{"id":"1","name":"test","email":"test@test.com","created":"2024-01-06T22:30:25.760522681+09:00","A":{}}
```

idのint型をjson上、表示したくない場合は下記のように設定する。
```go
type User struct {
	ID      int       `json:"-"`
	/** 以下省略 */
}
```

処理結果
```go
{"name":"test","email":"test@test.com","created":"2024-01-06T22:30:25.760522681+09:00","A":{}}
```

### omitempty
各型の初期値（intなら0、stringなら""）が構造体で設定された場合、json上表示したくない場合がある。  
その際の設定方法として`omitempty`という設定方法がある。

```go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type A struct{}

type User struct {
	ID      int       `json:"id,omitempty"` // 初期値の場合、jsonに表示しない
	Name    string    `json:"name,omitempty"` // 初期値の場合、jsonに表示しない
	Email   string    `json:"email"`
	Created time.Time `json:"created"`
	A       *A        `json:"A,omitempty"` // 空の構造体の場合、ポインタ型にしてomitempty指定するとjsonに表示されなくなる。
}

func main() {

	u := new(User)
	u.ID = 0 // 初期値が設定
	u.Name = ""// 初期値が設定
	u.Email = "test@test.com"
	u.Created = time.Now()

	// Marshal JSONに変換 ⇒ 返り値：byteのスライス
	bs, err := json.Marshal(u)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(bs))
}
```

処理結果
```go
{"email":"test@test.com","created":"2024-01-06T22:30:25.760522681+09:00"}
```

### jsonからの変換（デコード）


```go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type A struct{}

type User struct {
	ID      int       `json:"id"`
	Name    string    `json:"name"`
	Email   string    `json:"email"`
	Created time.Time `json:"created"`
	A       *A        `json:"A"`
}

func main() {

	u := new(User)
	u.ID = 1
	u.Name = "test"
	u.Email = "test@test.com"
	u.Created = time.Now()

	// Marshal JSONに変換 ⇒ 返り値：byteのスライス
	bs, err := json.Marshal(u)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(bs))

	// -----------------------------------

	fmt.Printf("%T\n", bs) // bsの型を調べる⇒バイトのスライス

	u2 := new(User)

	// Unmarshal JSONをデータに変換
	if err := json.Unmarshal(bs, u2); err != nil {
		fmt.Println(err)
	}

	fmt.Println(u2)
}
```

### マーシャルのカスタム
構造体からJSONに変更する際に値をカスタムしたい際は、ユーザのメソッドとしてMarshalJSONというメソッドを用意する必要がある。

```go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

// マーシャルのカスタム

type A struct{}

type User struct {
	ID      int       `json:"id"`
	Name    string    `json:"name"`
	Email   string    `json:"email"`
	Created time.Time `json:"created"`
	A       *A        `json:"A"`
}

// Json⇒構造体への変換時に値をカスタムする
func (u User) MarshalJSON() ([]byte, error) {
	v, err := json.Marshal(&struct {
		Name string
	}{
		Name: "Mr" + u.Name,
	})
	return v, err
}

func main() {

	u := new(User)
	u.ID = 1
	u.Name = "test"
	u.Email = "test@test.com"
	u.Created = time.Now()

	// Marshal JSONに変換 ⇒ 返り値：byteのスライス
	bs, err := json.Marshal(u) // この際にMarshalJSON()メソッドが暗黙的に呼び出されている。
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(bs))

	// -----------------------------------

	fmt.Printf("%T\n", bs) // bsの型を調べる⇒バイトのスライス

	u2 := new(User)

	// Unmarshal JSONをデータに変換
	if err := json.Unmarshal(bs, u2); err != nil {
		fmt.Println(err)
	}

	fmt.Println(u2)
}
```

出力結果
```go
{"Name":"Mrtest"} // "test"ではなく"Mrtest"とカスタムした値になっている。
[]uint8
&{0 Mrtest  0001-01-01 00:00:00 +0000 UTC <nil>} // "test"ではなく"Mrtest"とカスタムした値になっている。
```

**どのような場面で利用するのか。**
構造体に別の構造体が含まれており、それが空の場合、`omitempty`による非表示が効かない（ポインタ型にすればできるが）場合がある。  
その際に不都合なため、非表示するために利用するケースなどがある。

参考：[MarshalJSONを実装する](https://qiita.com/taroshin/items/be00bea3371ade705a2d)

### アンマーシャルのカスタム
```go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

// アンマーシャルのカスタム

type A struct{}

type User struct {
	ID      int       `json:"id"`
	Name    string    `json:"name"`
	Email   string    `json:"email"`
	Created time.Time `json:"created"`
	A       A         `json:"A"`
}

// 構造体⇒JSONへの変換時に値をカスタムする
func (u User) MarshalJSON() ([]byte, error) {
	v, err := json.Marshal(&struct {
		Name string
	}{
		Name: "Mr" + u.Name,
	})
	return v, err
}

// JSON⇒構造体への変換時に値をカスタムする
func (u *User) UnmarshalJSON(b []byte) error {
	// 仮のユーザ型としてUser2型を作成する
	type User2 struct {
		Name string
	}
	var u2 User2
	err := json.Unmarshal(b, &u2)
	if err != nil {
		fmt.Println(err)
	}
	u.Name = u2.Name + "!"
	return err
}

func main() {

	u := new(User)
	u.ID = 1
	u.Name = "test"
	u.Email = "test@test.com"
	u.Created = time.Now()

	// Marshal JSONに変換 ⇒ 返り値：byteのスライス
	bs, err := json.Marshal(u) // この際にMarshalJSON()メソッドが暗黙的に呼び出されている。
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(bs))

	// -----------------------------------

	fmt.Printf("%T\n", bs) // bsの型を調べる⇒バイトのスライス

	u2 := new(User)

	// Unmarshal JSONをデータに変換
	if err := json.Unmarshal(bs, u2); err != nil {
		fmt.Println(err)
	}

	fmt.Println(u2)
}
```

出力結果
```go
{"Name":"Mrtest"} // カスタムマーシャルにより、Mrtestに変換
[]uint8
&{0 Mrtest!  0001-01-01 00:00:00 +0000 UTC {}} // カスタムアンマーシャルにより、Mrtest⇒Mrtest!に変換
```

## sort
スライスなどをソートする機能が用意されているパッケージ

### Ints / Strings
- Ints : 数値を並び替える
- Strings : 文字列を並び替える
```go
package main

import (
	"fmt"
	"sort"
)

func main() {
	i := []int{5, 3, 2, 4, 5, 6, 4, 8, 9, 8, 7, 10}
	s := []string{"a", "2", "j"}

	fmt.Println(i, s)

	sort.Ints(i)

	sort.Strings(s)

	fmt.Println(i, s)

}
```

出力結果
```
[5 3 2 4 5 6 4 8 9 8 7 10] [a 2 j]
[2 3 4 4 5 5 6 7 8 8 9 10] [2 a j]
```

### Slice / SliceStable
sort.Sliceとsort.SliceStableの違い
- sort.Slice : 安定ソートを保証しない
- sort.SliceStable : 安定ソートを保証する

> NOTE:  
> 安定ソートとは  
> 安定ソート（あんていソート、stable sort）とは、ソート（並び替え）のアルゴリズムのうち、  
> 同等なデータのソート前の順序が、ソート後も保存されるものをいう。つまり、ソート途中の各状態において、  
> 常に順位の位置関係を保っていることをいう。  
> たとえば、学生番号順に整列済みの学生データを、テストの点数順で安定ソートを用いて並べ替えたとき、  
> ソート後のデータにおいて、同じ点数の学生は学生番号順で並ぶようになっている。

Sliceサンプルコード
```go
package main

import (
	"fmt"
	"sort"
)

type Entry struct {
	Name string
	Value int
}

func main() {
	el := []Entry{
		{"A", 20},
		{"F", 40},
		{"i", 30},
		{"b", 10},
		{"t", 15},
		{"y", 30},
		{"c", 30},
		{"w", 30},
	}
	fmt.Println(el)

	// 文字列の順で並び替える
	// 第一引数：スライス
	// 第二引数：条件式（無名関数・今回は名前を昇順でソート）
	sort.Slice(el, func(i, j int) bool {return el[i].Name < el[j].Name})

	fmt.Println("--------------------------------------------------------")
	fmt.Println(el)
	fmt.Println("--------------------------------------------------------")
}
```

出力結果
```
[{A 20} {F 40} {i 30} {b 10} {t 15} {y 30} {c 30} {w 30}]
--------------------------------------------------------
[{A 20} {F 40} {b 10} {c 30} {i 30} {t 15} {w 30} {y 30}]
--------------------------------------------------------
```

SliceStableのサンプルコード
```go
package main

import (
	"fmt"
	"sort"
)

type Entry struct {
	Name string
	Value int
}

func main() {
	el := []Entry{
		{"A", 20},
		{"F", 40},
		{"i", 30},
		{"b", 10},
		{"t", 15},
		{"y", 30},
		{"c", 30},
		{"w", 30},
	}
	fmt.Println(el)

	// 文字列の順で並び替える
	sort.SliceStable(el, func(i, j int) bool {return el[i].Name < el[j].Name})

	fmt.Println("--------------------------------------------------------")
	fmt.Println(el)
	fmt.Println("--------------------------------------------------------")
}
```

出力結果
```
[{A 20} {F 40} {i 30} {b 10} {t 15} {y 30} {c 30} {w 30}]
--------------------------------------------------------
[{A 20} {F 40} {b 10} {c 30} {i 30} {t 15} {w 30} {y 30}]
--------------------------------------------------------
```

### カスタムソート
- sort.Sort()メソッドを実装する
- Len(), Swap(), Less()を実装する。※実装しないとコンパイルエラー
- Less()にてカスタムソートの内容を実装する。

サンプルコード
```go
package main

import (
	"fmt"
	"sort"
)

type Entry struct {
	Name  string
	Value int
}

type List []Entry

func (l List) Len() int {
	return len(l)
}

func (l List) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

// ここでカスタムする。
func (l List) Less(i, j int) bool {
	if l[i].Value == l[j].Value {
		return (l[i].Name < l[j].Name)
	} else {
		return (l[i].Value < l[j].Value)
	}
}

func main() {
	m := map[string]int{"S": 1, "J": 4, "A": 3, "N": 3}

	lt := List{}
	for k, v := range m {
		e := Entry{k, v}
		lt = append(lt, e)
	}

	// Sort ⇒ Len(), Swap(), Less()の3つの関数を実装することで利用可能となる
	// Less()にカスタムソートを実装する。
	// ※実装しなければコンパイルエラー
	sort.Sort(lt)
	fmt.Println(lt)

	// Reverse
	sort.Sort(sort.Reverse(lt))
	fmt.Println(lt)
}
```

出力結果
```
[{S 1} {A 3} {N 3} {J 4}] // 昇順
[{J 4} {N 3} {A 3} {S 1}] // 降順
```

## context
APIのサーバーやクライアントを利用する際にリクエストに対してタイムアウトやキャンセル処理ができるようにする機能を提供する。

下記はサンプルコード
```go
package main

import (
	"context"
	"fmt"
	"time"
)

func longProcess(ctx context.Context, ch chan string) { // <1>
	fmt.Println("開始")
	time.Sleep(2 * time.Second)
	fmt.Println("終了")
	ch <- "実行結果"
}

func main() {
	ch := make(chan string)

	ctx := context.Background() // <2>

	ctx, cancel := context.WithTimeout(ctx, 1*time.Second) // <3>

	defer cancel() // <4>

	go longProcess(ctx, ch) // <5>

L:
	for { // <6>
		select { // <7>
		case <-ctx.Done():
			fmt.Println("##########Error##########")
			fmt.Println(ctx.Err())
			break L
		case s := <-ch:
			fmt.Println(s)
			fmt.Println("success")
			break L
		}
	}

	fmt.Println("ループ終了")
}
```
解説
1. longProcess()関数は2秒スリープし、最後に引数のチャネルに文字列"実行結果"を送信する。
2. あ
3. contextの引数にcontextとタイムアウト（１秒）を渡し、タイムアウト値を設定する。
4. 1秒立った場合、遅延処理としてキャンセルをする。
5. goroutineの非同期処理にてlongProcess()関数を実行する
6. 無限ループにて、チャネルに送信されたデータを処理する
7. selectにて、チャネルに文字列が挿入されている場合は"success"を表示し、処理終了。contextが終了している場合はエラーを出力し処理終了。

出力結果
```
開始
##########Error##########
context deadline exceeded
ループ終了
```

エラーとなる。（タイムアウトしている）

タイムアウトをさせないようにするには下記の箇所を設定しなおす必要がある。

```go
package main

import (
	"context"
	"fmt"
	"time"
)

func longProcess(ctx context.Context, ch chan string) {
	fmt.Println("開始")
	time.Sleep(2 * time.Second) // もしくは処理遅延時間をタイムアウト値よりも短くする。
	fmt.Println("終了")
	ch <- "実行結果"
}

func main() {
	ch := make(chan string)

	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second) // タイムアウトを1秒⇒3秒に変更

	defer cancel()

	go longProcess(ctx, ch)

L:
	for {
		select {
		case <-ctx.Done():
			fmt.Println("##########Error##########")
			fmt.Println(ctx.Err())
			break L
		case s := <-ch:
			fmt.Println(s)
			fmt.Println("success")
			break L
		}
	}

	fmt.Println("ループ終了")
}
```

処理結果
```
開始
終了
実行結果
success
ループ終了
```

## net/url
URL文字列を処理する機能を保持したパッケージ

参考コード: [サンプルコード](../../../udemy/src/standardLibraries/net/url/main.go)

## net/http client
HTTPメソッドを用いたGETやPOST処理を提供するパッケージ

参考コード：[](../../../udemy/src/standardLibraries/net/http%20client/main.go)

## net/http server
FWの根幹となる機能を提供するパッケージ

参考コード：[](../../../udemy/src/standardLibraries/net/http%20server//main.go)

- http.ListenAndServer()
  - 第一引数：サーバーのポート。デフォルト値は80番ポート。
  - 第二引数：ハンドラーを指定する。デフォルトはマルチプレクサ。

### サーバーを起動する

```go
package main

import "net/http"

func main() {
	http.ListenAndServe(":8080", nil)
}
```

上記で`go run main.go`を実行するとサーバーが起動する。⇒しかし、404 Page Not Foundになる。

これはハンドラーが指定されていないためである。  
サーバーに何かしらの処理をさせたい場合は、ハンドラーを指定しなければならない。

次にMyHandlerを設定して画面に文字列を表示するサンプルコードに修正する。

### ハンドラーを定義して画面に文字列を表示

```go
package main

import (
	"fmt"
	"net/http"
)

type MyHandler struct {} // <1>

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) { // <2>
	fmt.Fprintf(w, "Hello World")
}

func main() {
	http.ListenAndServe(":8080", &MyHandler{}) // <3>
}
```

1. ハンドラーの構造体を設定
2. ハンドラーのメソッドとして`ServeHTTP`を定義し、画面に`Hello World`と表示するように定義
3. http.ListenAndServe()にてハンドラーをポインタ型で参照するように定義

上記で画面に `Hello World` と表示されるようになった。

しかし、現状はマルチプレクサでの指定をしていないので、どのような操作をしても画面に `Hello World` が表示されるだけの状態になっている。

次にURL毎に異なる処理を実施するようにしたい。

### マルチプレクサの設定
http.ListenAndServe()の第二引数をnilにしてマルチプレクサを利用する。

```go
package main

import (
	"html/template"
	"log"
	"net/http"
)

func top(w http.ResponseWriter, r *http.Request) { // <1>
	t, err := template.ParseFiles("tmpl.html") // テンプレートの構造体を生成する
	if err != nil {
		log.Panicln(err)
	}
	t.Execute(w, "Hello World Top Page") // テンプレートの実行。第一引数はレスポンスライター、第二引数はデータを設定。
}

func about(w http.ResponseWriter, r *http.Request) { // <2>
	t, err := template.ParseFiles("about.html")
	if err != nil {
		log.Panicln(err)
	}
	t.Execute(w, "Hello World About Page")
}

func main() {
	http.HandleFunc("/top", top) // <3>
	http.HandleFunc("/about", about) // <4>
	http.ListenAndServe(":8080", nil) // <5>
}
```

1. topページの処理を行う関数を用意する。
2. aboutページの処理を行う関数を用意する。
3. http.HandleFunc()にて、第一引数にURLパス、第二引数に関数を設定する。(/topに対してtop関数)
4. http.HandleFunc()にて、第一引数にURLパス、第二引数に関数を設定する。(/aboutに対してabout関数)
5. マルチプレクサを利用するため、第二引数をnilにする。

上記で `http://localhost:8080/top` と `http://localhost:8080/about` に遷移すると、それぞれ行いたい表示処理ができている。

各種HTMLの補足  
tmpl.html
```html
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Go Web Programming</title>
</head>
<body>
  {{ . }} <!-- この記述により、t.Execute(w, "Hello World Top Page")の文字列が表示される。 -->
</body>
</html>
```