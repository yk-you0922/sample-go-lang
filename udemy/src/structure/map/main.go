package main

import "fmt"

type User struct {
	Name string
	Age int
}

func main() {
	m := map[int]User {
		1: {Name: "user1", Age: 10},
		2: {Name: "user2", Age: 20},
		3: {Name: "user3", Age: 30},
	}
	fmt.Println(m)

	m2 := map[User]string {
		{Name: "user1", Age: 10}: "Tokyo",
		{Name: "user2", Age: 20}: "Osaka",
		{Name: "user3", Age: 30}: "Nagoya",
	}
	fmt.Println(m2)

	m3 := make(map[int]User)
	fmt.Println(m3)
	m3[1] = User{Name: "user3", Age: 30}
	fmt.Println(m3)

	for _, v := range m {
		fmt.Println(v)
	}

	for _, v := range m2 {
		fmt.Println(v)
	}

	for _, v := range m3 {
		fmt.Println(v)
	}
}
