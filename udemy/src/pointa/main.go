package main

import "fmt"

// ポインタ

func Double(i int) {
	i = i * 2
}

func DoubleV2(i *int) {
	*i = *i * 2
}

func DoubleV3(s []int) {
	for i, v := range s {
		s[i] = v * 2
	}
}

func main() {
	var n int = 100
	fmt.Println(n) // 100

	// アドレス表示
	fmt.Println(&n) // 0xc000018080

	Double(n)
	fmt.Println(n) // 100のまま

	var p *int = &n // アドレスを渡している
	fmt.Println(p) // 0xc000018080
	fmt.Println(*p)

	// *p = 300
	// fmt.Println(n) // 300になる。⇒アドレスを変更しているため。

	DoubleV2(&n)
	fmt.Println(n)

	DoubleV2(p)
	fmt.Println(*p)

	var sl []int = []int{1, 2, 3}
	DoubleV3(sl)
	fmt.Println(sl)
}
