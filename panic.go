package main

import "fmt"

func main(){
	fmt.Println("helloをコール")
	hello()
}

func hello()int{
	fmt.Println("パニックを呼び出します")
	panic("パニックを呼び出しました")
}