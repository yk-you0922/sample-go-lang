package main

import "fmt"

type User struct {
	Name string
	Age int
}

func (u User) SayName() {
	fmt.Println(u.Name)
}

func (u *User) SetName(name string) {
	u.Name = name
}

func main() {
	user1 := User{Name: "user1", Age: 10}
	user1.SayName()

	user1.SetName("A")
	user1.SayName()
}
