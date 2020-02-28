package main

import(
	"fmt"
	"os"
)



func main(){
	for i:=1;i<=10;i++{
		os.MkdirAll(fmt.Sprintf("/tmp/go-dirs/%02d",i),os.ModePerm)
	}

}
