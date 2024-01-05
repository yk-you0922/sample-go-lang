# テスト

## 基本
- テスト対象パッケージと同じ階層に作る
- <テスト対象パッケージ名> + _test.goで作成
  - 例）main.go⇒main_test.go
- テスト用パッケージ `testing`
- テスト用パッケージ名はテスト対象パッケージ名
  - 例）main.go⇒ `package main`

サンプルコード

テスト対象パッケージ(main.go)
```go
package main

import "fmt"

func IsOne(i int) bool {
	if i == 1 {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println(IsOne(1))
	fmt.Println(IsOne(0))
}
```

テストコード(main_test.go)
```go
package main // <1>

import "testing" // <2>

var Debug bool = false

func TestIsOne(t *testing.T) { // <3> <4>
	i := 1
	if Debug { // <5>
		t.Skip("スキップさせる")
	}
	v := IsOne(i) // <6>

	if !v { // <7>
		t.Errorf("%v != %v", i, 1)
	}
}
```
1. パッケージ名はテスト対象パッケージ名。
2. テスト用パッケージをimport。
3. 命名は `Test + テスト対象メソッド名` とする。
4. テスト用パッケージを引数として指定する。
5. テストさせるかどうかを引数でコントロールする⇒今回はfalseなのでテストを実施する。
6. テスト対象関数の実行。
7. テスト結果のチェック。

テストコマンド
```bash
cd テストコードの階層
go test
または
go test -v
```
※ `-v`オプションを付与することによりテストの詳細が出力される。

テスト成功時
```bash
go test
PASS
ok      sample-go-lang/udemy/src/test   0.001s
```

テスト失敗時
```bash
go test
--- FAIL: TestIsOne (0.00s)
    main_test.go:15: 0 != 1
FAIL
exit status 1
FAIL    sample-go-lang/udemy/src/test   0.001s
```

## パッケージ内のテスト実行コマンド
```bash
cd ルートディレクトリ
go test ./...
```
ルートディレクトリで `go test ./...` と入力する。

## カバレッジ
```bash
go test -cover

ok      sample-go-lang/udemy/src/test   0.002s  coverage: 28.6% of statements
ok      sample-go-lang/udemy/src/test/alib      0.002s  coverage: 100.0% of statements
```