package main

import "fmt"

// マップ

func main() {
	// 明示的宣言
	var m = map[string]int{"A": 100, "B": 200}
	fmt.Println(m)

	// 暗黙的宣言
	m2 := map[string]int{"A": 100, "B": 200}
	fmt.Println(m2)

	m3 := map[int]string{
		1: "A",
		2: "B",
	}
	fmt.Println(m3)

	// make関数による宣言
	m4 := make(map[int]string)
	fmt.Println(m4)
	m4[1] = "JAPAN"
	m4[2] = "USA"
	fmt.Println(m4)

	// 値の取り出し map変数名[<keyとなる値・文字列>]
	fmt.Println(m["A"])
	fmt.Println(m4[2])
	fmt.Println(m4[3]) // 登録されていないキーは型の初期値を返却してしまうことに注意。

	// 値取り出し時のエラーハンドリング
	s, ok := m4[1]
	fmt.Println(s)
	fmt.Println(ok) // 取り出しに成功時=true, 失敗時=false

	// エラー挙動確認
	s2, ok2 := m4[3]
	if !ok2 {
		fmt.Println("error")
	}
	fmt.Println(s2, ok2)

	// 削除 delete関数
	m4[3] = "CHINA"
	fmt.Println(m4)
	delete(m4, 3)
	fmt.Println(m4)

	// len関数による要素数の確認
	fmt.Println(len(m4))
}