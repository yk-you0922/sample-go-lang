package main

import "fmt"

func main() {
	/*
	* float・・・float32とfloat64がある。
	* 暗黙型定義の場合はfloat32に自動的になる。
	* float32はあまり使わない。
	*/
	var fl64 float64 = 2.4
	fmt.Println(fl64)

	fl := 3.2 // 暗黙的型定義のため、float32になる
	fmt.Println(fl)

	fmt.Printf("%T, %T\n", fl64, fl)

	var fl32 float32 = 1.2
	fmt.Printf("%T\n", fl32)

	// 型変換
	fmt.Printf("%T\n", float64(fl32))

	zero := 0.0
	pinf := 1.0 / zero
	fmt.Println(pinf)

	ninf := -1.0 / zero
	fmt.Println(ninf)

	nan := zero /zero
	fmt.Println(nan)
}
