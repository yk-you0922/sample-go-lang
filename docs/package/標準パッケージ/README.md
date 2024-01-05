# 標準パッケージ

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

## flag

## fmt

## log

## strconv

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