package main

import "fmt"

// スライス
// 宣言、操作

func main() {
	// 初期宣言
	var sl []int
	fmt.Println(sl)
	fmt.Printf("%T\n", sl)

	// 明示的な宣言
	var sl2 []int = []int{100, 200}
	fmt.Println(sl2)
	fmt.Printf("%T\n", sl2)

	// 暗黙的な宣言
	sl3 := []string{"A", "B"}
	fmt.Println(sl3)
	fmt.Printf("%T\n", sl3)

	// make関数によるスライス生成
	sl4 := make([]int, 5)
	fmt.Println(sl4)
	fmt.Printf("%T\n", sl4)

	// 値の代入
	sl2[0] = 1000
	fmt.Println(sl2)

	// 値の取り出し
	sl5 := []int{1, 2, 3, 4, 5}
	fmt.Println(sl5)

	fmt.Println(sl5[0])
	fmt.Println(sl5[2:4])            // 2番目-4番目の3,4を取り出す
	fmt.Println(sl5[:2])             // 2番目の手前の1, 2を取り出す
	fmt.Println(sl5[2:])             // 2番目の以降の3,4,5を取り出す
	fmt.Println(sl5[:])              // 全ての要素を取り出す
	fmt.Println(sl5[len(sl5)-1])     // 最後の要素を取り出す
	fmt.Println(sl5[1 : len(sl5)-1]) // 2, 3, 4を取り出す
}
