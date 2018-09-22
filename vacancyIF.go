package main

import "fmt"

type Hello interface{} // vacancy interface

// vacancy interface is first grade object so, 
// substitutable for any other values and types.

func main(){

	var hh Hello = "my name is Yaako"
	fmt.Println(hh)

	hh=12334
	fmt.Println(hh)
}