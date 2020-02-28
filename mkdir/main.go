package main

import(
	"os"
	"syscall"
)

func main(){
	defaultUmask:=syscall.Umask(0)
	os.Mkdir("/tmp/go-dir2",0777)
	syscall.Umask(defaultUmask)
	//print(defaultUmask)
}
