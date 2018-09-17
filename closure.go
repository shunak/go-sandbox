package main

import "fmt"

func main() {
	var kansu = fn()
	fmt.Println(kansu())
	fmt.Println(kansu())
	fmt.Println(kansu())
}

func fn() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}
