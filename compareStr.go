package main

import (
	"fmt"
)

func main() {

	str1 := "Hello"
	str2 := "world"

	if str1 == str2 {
		fmt.Println(str1, str2)
	} else {
		fmt.Println("文字列はことなります")
	}
}
