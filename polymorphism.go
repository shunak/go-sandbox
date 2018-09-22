package main

import "fmt"

type Calc interface { // define interface
	Calc(a int, b int) int
}

// Add
type Add struct{
}

func (x Add) Calc(a int, b int) int {
	return(a+b)
}

// Sub
type Sub struct{
}


func (x Sub) Calc(a int, b int) int {
	return(a-b)
}


func main(){
	var add Add
	var sub Sub
	var cal Calc // interface

	cal = add
	fmt.Printf("足し算　=%d\n", cal.Calc(2,1))

	cal = sub
	fmt.Printf("引き算　=%d\n", cal.Calc(2,1))

}














