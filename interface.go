package main

import "fmt"

// unfinished

type Hito struct {
	name string
	age int
}

type Human interface{
	aisatsu()
	jikoshoukai()
}


func (ps Hito) aisatsu(){
	fmt.Printf("私の名前は%sです\n",ps.name)
}

func (ps Hito) jikoshoukai(){
	fmt.Printf("私は%s %d歳です。\n",ps.name,ps.age)
}

func main(){
	m:=Hito{"キコ",17}
	var h Human = m

	h.aisatsu()
	h.jikoshoukai()
}



