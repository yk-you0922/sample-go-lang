package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s string = "100"

	i, err := strconv.Atoi(s)

	// エラーハンドリング
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("i = %T\n", i)
}
