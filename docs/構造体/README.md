# 構造体

## 構造体とは
Goにはオブジェクト指向における`Class`というものが存在しないので代用品として構造体という考え方がある。

- オブジェクト指向プログラミング言語のクラスのような存在
- 複数の型の値をひとまとめにしたようなもの
- Javaのエンティティクラス的な感じ

## 基礎文法
ユーザ型を定義
```go
type User struct {
	Name string
	Age int
	X, Y int
}
```

main関数での利用
```go
package main

import "fmt"

// ユーザ型の構造体の定義
type User struct {
	Name string
	Age int
}


func main() {
	var user1 User
	fmt.Println(user1)
}
```

## メソッド
任意の方に関連付けられた関数のこと。

下記のような構造になる。
```go
func (構造体引数 構造体型名) 関数名(引数) {
    // TODO: 実装
}
```

```go
package main

import "fmt"

type User struct {
	Name string
	Age int
}

// メソッド
func (u User) SayName() {
	fmt.Println(u.Name)
}

// 参照渡し（ポインタ）していることに注意。⇒ポインタ型にしないと更新されない。
func (u *User) SetName(name string) {
	u.Name = name
}

func main() {
	user1 := User{Name: "user1", Age: 10}
	user1.SayName() // user1 と出力

    user1.SetName("A")
	user1.SayName() // A と出力
}
```

## コンストラクタ
Goにはコンストラクタ関数は存在しない。  
存在しないが、パターンとしてコンストラクタ関数を使うことがよくある。

```go
package main

import "fmt"

type User struct {
	Name string
	Age int
}

// 返り値がポインタ型（参照型）となっていることに注意。
func NewUser(name string, age int) *User {
	return &User{Name: name, Age: age} // アドレスを返却するので & をつける
}

func main() {
	user1 := NewUser("user1", 10)
	fmt.Println(user1)
}
```

## 構造体とスライス
JavaでいうところのList等に類するものっぽい

スライスの宣言
```go
type User struct {
	Name string
	Age int
}

type Users []*User
```

スライスの宣言（非推奨）  
こちらでも記述できるが、あまり使われない。
```go
type Users struct {
	Users []*Users
}
```

スライスの利用
```go
func main() {
	user1 := User{Name: "user1", Age: 10}
	user2 := User{Name: "user2", Age: 20}
	user3 := User{Name: "user3", Age: 30}
	user4 := User{Name: "user4", Age: 40}

	users := Users{}
	users = append(users, &user1)
	users = append(users, &user2)
	users = append(users, &user3, &user4)

	fmt.Println(users)

	for _, u := range users {
		fmt.Println(*u)
	}

	// make関数でもスライスを生成できる
	users2 := make([]*User, 0)
	users2 = append(users2, &user1, &user2)

	for _, u := range users2 {
		fmt.Println(*u)
	}
}
```

## 独自型
int型のMyInt型を定義  
元はint型だが、int型との計算ができないなどもある。
```go
type MyInt int
```
