package main

import (
	"fmt"
)

func main() {

	set := map[int]bool{}

	set[175] = true
	set[119] = true
	set[450] = true

	fmt.Println(set)
	// iteration1
	for key, _ := range set {
		fmt.Print(key)
		fmt.Print(" ")

	}
	fmt.Println()

	// iteration2

	for key, _ := range set {
		fmt.Print(key)
		fmt.Print(" ")
	}
	fmt.Println()

}
