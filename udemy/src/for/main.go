package main

import "fmt"

const Sep = "--------------------------------------------------"

func main() {
	fmt.Println("簡易なfor文")
	i := 0
	for {
		i++
		if i == 3 {
			break
		}
		fmt.Println("Loop")
	}

	fmt.Println(Sep)

	fmt.Println("条件付きfor文")
	point := 0
	for point < 10 {
		fmt.Println(point)
		point++
	}

	fmt.Println(Sep)

	fmt.Println("古典的なfor文")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Println(Sep)

	fmt.Println("continueを用いたfor文")
	for i := 0; i < 10; i++ {
		if i % 2 == 0 {
			continue
		}
		fmt.Println(i)
	}

	fmt.Println(Sep)

	fmt.Println("配列を用いたfor文")
	arr := [3]int{1, 2, 3}
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	fmt.Println(Sep)

	fmt.Println("配列を用いたfor文")
}
