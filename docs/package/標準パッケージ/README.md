# 標準パッケージ

## 公式ドキュメントリンク集
| No  | パッケージ名    | 説明                                                                                                                              | リンク                       |
| --- | --------------- | --------------------------------------------------------------------------------------------------------------------------------- | ---------------------------- |
| 1   | os              | オペレーティング システム機能へのプラットフォームに依存しないインターフェイスを提供                                               | https://pkg.go.dev/os        |
| 2   | time            | パッケージ時間には、時間を測定および表示する機能                                                                                  | https://pkg.go.dev/time      |
| 3   | math            | 基本的な定数と数学関数を提供                                                                                                      | https://pkg.go.dev/math      |
| 4   | rand            | シミュレーションなどのタスクに適した擬似乱数生成器を実装。セキュリティが重要な作業には使用しない                                  | https://pkg.go.dev/math/rand |
| 5   | flag            | コマンドライン フラグ解析を実装                                                                                                   | https://pkg.go.dev/flag      |
| 6   | fmt             | フォーマットされた I/O を実装                                                                                                     | https://pkg.go.dev/fmt       |
| 7   | log             | 単純なログ パッケージを実装                                                                                                       | https://pkg.go.dev/log       |
| 8   | strconv         | 基本データ型の文字列表現との変換を実装                                                                                            | https://pkg.go.dev/strconv   |
| 9   | strings         | UTF-8 でエンコードされた文字列を操作するための単純な関数を実装                                                                    | https://pkg.go.dev/strings   |
| 10  | bufio           | バッファリングされた I/O を実装                                                                                                   | https://pkg.go.dev/bufio     |
| 11  | loutil          | I/O ユーティリティ関数を実装。非推奨: Go 1.16 では、同じ機能がパッケージioまたはパッケージosによって提供されている。              | https://pkg.go.dev/io/ioutil |
| 12  | regexp          | 正規表現検索を実装                                                                                                                | https://pkg.go.dev/regexp    |
| 13  | sync            | 相互排他ロックなどの基本的な同期プリミティブを提供。このパッケージで定義されている型を含む値はコピーしない。                      | https://pkg.go.dev/sync      |
| 14  | crypto          | 共通の暗号定数を収集                                                                                                              | https://pkg.go.dev/crypto    |
| 15  | json            | RFC 7159で定義されている JSON のエンコードとデコードを実装                                                                        | https://pkg.go.dev/json      |
| 16  | sort            | スライスとユーザー定義のコレクションをソートするためのプリミティブを提供                                                          | https://pkg.go.dev/sort      |
| 17  | context         | 期限、キャンセル シグナル、および API 境界を越えてプロセス間でリクエストをスコープとするその他の値を伝達する Context タイプを定義 | https://pkg.go.dev/context   |
| 18  | net/url         | HTTP クライアントとサーバーの実装を提供                                                                                           | https://pkg.go.dev/net       |
| 19  | net/http client | HTTP クライアントとサーバーの実装を提供                                                                                           | https://pkg.go.dev/net       |
| 20  | net/http server | HTTP クライアントとサーバーの実装を提供                                                                                           | https://pkg.go.dev/net       |


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


## bufio

## ioutil

## regexp

## sync

## crypto

## json

## sort

## context

## net/url

## net/http client

## net/http server