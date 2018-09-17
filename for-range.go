package main

import (
	"fmt"
)

func main() {

	for i, t := range []int{1, 2, 3, 4, 5} {
		fmt.Printf("i=%d, t=%d\n", i, t)
	}

}
