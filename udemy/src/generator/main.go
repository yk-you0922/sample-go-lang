package main

import "fmt"

// ジェネレーター

/*
* Goにはジェネレーターはない。
* クロージャーを利用することで簡易的なジェネレーターができる。
*/

func integers() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	ints := integers()
	fmt.Println(ints())// 1
	fmt.Println(ints())// 2
	fmt.Println(ints())// 3
	fmt.Println(ints())// 4
	fmt.Println(ints())// 5
	fmt.Println(ints())// 6

	fmt.Println("--------------------------")

	// 変数intsとは値が共有されないため、注意
	otherInts := integers();
	fmt.Println(otherInts())
	fmt.Println(otherInts())
	fmt.Println(otherInts())
	fmt.Println(otherInts())
	fmt.Println(otherInts())
	fmt.Println(otherInts())
}
