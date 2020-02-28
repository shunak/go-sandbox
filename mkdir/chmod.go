package main

import(
	"os"
)

func main(){
	directory:="/tmp/go-dir3"
	os.Mkdir(directory,os.ModePerm)
	os.Chmod(directory,0777)
}
