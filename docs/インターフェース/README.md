# インターフェース
同じメソッド（機能）を持つ複数の型をひとくくりにして扱うための仕組み

interfaceを利用することによって同じメソッド名を持つ型を同様に扱うことができる。

## 使い方
interface定義
1. `Person`型と`Car`型の2つの型を用意する。
2. それぞれの型で`ToString`メソッドを用意する。
3. interfaceとしてStringfyを用意し、その中で`ToString`メソッドを定義する。

呼び出し元での利用
1. `Stringfy`インターフェースのスライスを用意する。
2. for文でループさせ、呼び出す。

```go
package main

import "fmt"

type Stringfy interface {
	ToString() string
}

type Person struct {
	Name string
	Age int
}

/**
* 文字列を返却する関数
*/
func (p *Person) ToString() string {
	return fmt.Sprintf("Name=%v, Age=%v", p.Name, p.Age)
}

type Car struct {
	Number string
	Model string
}

/**
* 文字列を返却する関数
*/
func (c *Car) ToString() string {
	return fmt.Sprintf("Number=%v, Model=%v", c.Number, c.Model)
}

func main() {

	vs := []Stringfy{
		&Person{Name: "Taro", Age: 21},
		&Car{Number: "123-456", Model: "AB-1234"},
	}

	for _, v := range vs {
		fmt.Println(v.ToString())
	}
}
```

### 注意点
interface型の変数を利用する際、メソッドしか使えなくなり、そこに当てはまる構造体のフィールドは利用できなくなる。

上記の例であれば、Car.Modelを出力したくないとしても、Car.ToString()に定義されている内容しか利用できない。

## カスタムエラー
Goの組み込み型のエラーはインターフェースとして定義されている。
```go
type error interface {
	Error() string
}
```

エラーハンドリングなどで頻繁に利用されるエラー型

### カスタムエラーの定義
カスタムエラーは下記のように定義する。
```go
type MyError struct {
	Message string
	ErrCode int
}

func (e *MyError) Error() string {
	return e.Message
}
```

### カスタムエラーの利用
カスタムエラーは下記のように利用する。
```go
func main() {
	err := RaiseError()
	fmt.Println(err.Error())
}
```

ただし、上記の場合、変数`err`に格納したフィールド（例えばエラーコード）にはアクセスできないため、  
下記のように型アサーションを利用し、復元してアクセスするようにする。

```go
	err := RaiseError()
	fmt.Println(err.Error())

    // ok=trueの場合、MyError型であることを判定した扱いとなる。
	e, ok := err.(*MyError)
	if ok {
		fmt.Println(e.ErrCode)
	}
```

## Stringer
fmtパッケージに付与されているインターフェース  
`fmt.Stringer`として呼び出す。

文字列を返すstringのみが定義されている。

下記のように定義されている。
```go
type Stringer interface {
	String() string
}
```

上記をカスタマイズして、任意の型の文字列表現をカスタマイズすることができる。

### 使い方
下記の実装を行う。

```go
package main

import "fmt"

type Point struct {
	A int
	B string
}

func main() {
	p:= &Point{100, "ABC"}
	fmt.Println(p)
}
```
この際の出力は以下のようになる。
```bash
&{100 ABC}
```

ここに`Stringer`を実装する。

```go
package main

import "fmt"

type Point struct {
	A int
	B string
}

// <<A, B>>というように値の出力を変換する。
func (p* Point) String() string {
	return fmt.Sprintf("<<%v, %v>>", p.A, p.B)
}

func main() {
	p:= &Point{100, "ABC"}
	fmt.Println(p)
}
```

すると下記のように出力フォーマットを変換できる。

```bash
<<100, ABC>>
```