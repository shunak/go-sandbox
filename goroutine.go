package main

import (
	"fmt"
	"time"
)

var n = 100

func main(){
	fmt.Println("メイン関数を実行します")
	fmt.Printf("nの中身は %dです\n",n)
	go sub()
	time.Sleep(time.Millisecond *2)
}

func sub(){
	fmt.Println("goルーチンを実行中")
	fmt.Printf("nの中身は %dです\n",n)

	time.Sleep(time.Millisecond *10)
	fmt.Println("メイン関数が先に終了されｔら、ここは表示されません")
}
