package main

import "fmt"

func main() {
	var fnc func(int, int) int
	fnc = add
	fmt.Println(fnc(2, 1))
}

func add(a int, b int) int {
	return (a + b)
}
