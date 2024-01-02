package main

import "fmt"

// スライス
// append, make, len, cap

/*
* append・・・配列の最後の要素に追加する
* make	・・・各型の初期値で配列を生成する
* len	・・・配列の持つ要素数を調べる
* cap	・・・配列の持つ容量（メモリ）を調べる
*	⇒メモリ使用量を気にする開発時によく利用する。（性能面）
 */

func main() {
	sl := []int{100, 200}
	fmt.Println("初期宣言時の配列")
	fmt.Println(sl)

	// append
	fmt.Println("appendで300を追加")
	sl = append(sl, 300)
	fmt.Println(sl)
	fmt.Println("appendで400, 500を追加")
	sl = append(sl, 400, 500)
	fmt.Println(sl)

	// make
	fmt.Println("makeでint型の配列を生成")
	sl2 := make([]int, 5)
	fmt.Println(sl2)

	// len
	fmt.Println("lenで要素数を検索")
	fmt.Println(len(sl2))

	// cap
	fmt.Println("capで配列の容量を調べる")
	fmt.Println(cap(sl2))

	sl3 := make([]int, 5, 10)
	fmt.Println(len(sl3))
	fmt.Println(cap(sl3))

	sl3 = append(sl3, 1, 2, 3, 4, 5, 6, 7)
	fmt.Println(len(sl3))
	fmt.Println(cap(sl3))
}
