package main

import "fmt"

func main() {
	var i int
	i = Plus(1, 2)
	fmt.Println(i)

	var i2, i3 int
	i2, i3 = Div(9, 4)
	// i2, _ = Div(9, 4) // 使いたくない戻り値はアンダーバーとして破棄できる
	fmt.Println(i2, i3)

	var i4 int
	i4 = Double(200)
	fmt.Println(i4)

	NoReturn();
}

// 同じ型の引数は纏めることができる
// func Plus(x int, y int) int {
func Plus(x, y int) int {
	return x + y
}

// 複数の戻り値の関数
func Div(x, y int) (int, int) {
	var q int = x / y
	var r int = x % y
	return q, r
}

// 戻り値の変数を指定した関数
// 引数と戻り値が明確なのでわかりやすい
func Double(price int) (result int) {
	result = price * 2
	return // 戻り値の変数を指定しているため、returnの変数名は省略できる
}

func NoReturn () {
	fmt.Println("No Return")
	return
}