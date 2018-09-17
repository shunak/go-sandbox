package main

import "fmt"


func main(){
	yobidasi()
}


func yobidasi(){
	defer fmt.Println("終了")
	fmt.Println("先に表示")
}