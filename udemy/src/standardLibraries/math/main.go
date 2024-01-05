package main

import (
	"fmt"
	"math"
)

func main() {
	// 円周率
	fmt.Println(math.Pi)

	// 2の平方根
	fmt.Println(math.Sqrt2)

	// 立方根
	fmt.Println(math.Cbrt(8))

	// 数値型に関する定数
	fmt.Println(math.MaxFloat32) // float32で表現可能な最大値
	fmt.Println(math.SmallestNonzeroFloat32) // float32で表現可能な最小値（0ではない）
	fmt.Println(math.MaxFloat64) // float64で表現可能な最大値
	fmt.Println(math.SmallestNonzeroFloat64) // float64で表現可能な最小値（0ではない）
	fmt.Println(math.MaxInt64) // Int64で表現可能な最大値
	fmt.Println(math.MinInt64) // Int64で表現可能な最小値（0ではない）

	// 絶対値
	fmt.Println(math.Abs(100))
	fmt.Println(math.Abs(-100))

	// 累乗
	fmt.Println(math.Pow(0, 2))
	fmt.Println(math.Pow(2, 2))

	// 最大値、最小値
	fmt.Println(math.Max(1, 2))
	fmt.Println(math.Min(1, 2))

	// 小数点以下の切り捨て、切り上げ

	// 数値の正負を問わず、小数点以下を単純に切り捨てる
	fmt.Println(math.Trunc(3.5))
	fmt.Println(math.Trunc(-3.5))
	
	// 引数の数値より小さい最大の整数を返す
	fmt.Println(math.Floor(3.5))
	fmt.Println(math.Floor(-3.5))

	// 引数の数値より大きい最小の整数を返す
	fmt.Println(math.Ceil(3.5))
	fmt.Println(math.Ceil(-3.5))
}
