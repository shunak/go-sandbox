package main

import (
	"fmt"
)

func main() {

	var p *int
	n := 10

	p = &n

	fmt.Println(&n)
	fmt.Println(p)
	fmt.Println(*p)

}
