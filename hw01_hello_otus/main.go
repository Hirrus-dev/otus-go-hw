package main

import (
	"fmt"

	stringUtils "github.com/agrison/go-commons-lang/stringUtils"
)

func main() {
	stringIn := "Hello, OTUS!"
	stringOut := stringUtils.Reverse(stringIn)
	fmt.Println(stringOut)
}
