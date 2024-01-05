package main

import (
	"fmt"
	"sample-go-lang/udemy/src/package/scope/foo"
)

func appName() string {
	const AppName = "GoApp"
	var Version string = "1.0"
	return AppName + " " + Version
}

func main() {
	fmt.Println(foo.Max)
	fmt.Println(foo.ReturnMin())

	fmt.Println(appName())
}
