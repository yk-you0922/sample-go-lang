package main

import "fmt"

// ユーザ型の構造体の定義
type User struct {
	Name string
	Age int
}

func UpdateUser(user User) {
	user.Name = "A"
	user.Age = 1000
}

func UpdateUser2(user *User) {
	user.Name = "A"
	user.Age = 1000
}

func main() {
	var user1 User
	fmt.Println(user1)
	// 値の更新
	user1.Name = "user1"
	user1.Age = 10
	fmt.Println(user1) // {user1 10}

	// 暗黙的なユーザ定義
	user2 := User{}
	fmt.Println(user2)
	// 値の更新
	user2.Name = "user2"
	user2.Age = 20
	fmt.Println(user2) // {user2 20}

	// 初期値を持たせた宣言
	user3 := User{Name: "user3", Age: 30}
	fmt.Println(user3) // {user3 30}

	// フィールドを指定せずに初期値を持たせた宣言
	// 注意：構造体の順番に沿った宣言が必要。順番を誤るとエラーになる
	user4 := User{"user4", 40}
	fmt.Println(user4) // {user4 40}

	// newを利用した宣言⇒ポインタを利用した宣言になる
	user5 := new(User)
	fmt.Println(user5) // &{ 0}

	// アドレス演算子を利用した宣言⇒newと同じ（アドレス演算子を利用した書き方のほうが多い）
	user6 := &User{}
	fmt.Println(user6) // &{ 0}

	UpdateUser(user1)
	UpdateUser2(user6) // ポインタを引数として取る関数

	fmt.Println(user1) // {user1 10}⇒更新されていない
	fmt.Println(user6) // &{A 1000}⇒更新された
}
