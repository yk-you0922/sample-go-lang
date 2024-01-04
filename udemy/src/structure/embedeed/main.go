package main

import "fmt"

// Tの構造体にUser型を埋め込む
type T struct {
	// User User
	User // 型名の省略も可能
}

type User struct {
	Name string
	Age int
}

func (u *User) SetName() {
	u.Name = "A"
}

func main() {
	t := T{User: User{Name: "user1", Age: 10}}
	fmt.Println(t)
	fmt.Println(t.User)
	fmt.Println(t.User.Name)
	fmt.Println(t.User.Age)

	// 型名を省略した場合はtから直接Name等にアクセス可能
	fmt.Println(t.Name)

	// User型のメソッド実行
	t.User.SetName()
	fmt.Println(t)
}
