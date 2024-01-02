package main

import "fmt"

func main() {
	// 後から要素数を変更することができない
	// 変更する場合はスライスなどを用いる⇒javaのList的な感じ？

	// 明示的定義
	var arr1 [3]int
	fmt.Println(arr1)
	fmt.Printf("%T\n", arr1)

	var arr2 [3]string = [3]string{"A", "B"}
	fmt.Println(arr2)
	fmt.Printf("%T\n", arr2)

	// 暗黙的定義
	arr3 := [3]int{1, 2, 3}
	fmt.Println(arr3)
	fmt.Printf("%T\n", arr3)

	// 要素数の省略宣言
	arr4 := [...]string{"C", "D"}
	fmt.Println(arr4)
	fmt.Printf("%T\n", arr4)

	// 配列操作（値の取り出し）
	fmt.Println(arr2[0])
	fmt.Println(arr2[1])
	fmt.Println(arr2[2])
	fmt.Println(arr2[2-1]) // 最後の要素をとる場合は 配列[<要素数>-1]

	arr2[2] = "C"
	fmt.Println(arr2)

	var arr5 [4]int
	// arr5 = arr1 // 要素数が異なるため、同じint型の配列でも代入はできない
	fmt.Println(arr5)

	// 配列の要素数を調べる
	fmt.Println(len(arr1))
}
