package main

import (
	"fmt"
	"os"
)

func main() {
	// 標準出力
	fmt.Fprint(os.Stdout, "Hello")
	fmt.Fprint(os.Stdout, "Hello")
	fmt.Fprint(os.Stdout, "Hello")

	f, _ := os.Create("test.txt")
	defer f.Close()

	fmt.Fprint(f, "Fprint")
}
