package main

/*
#include <stdio.h>
void review(char* s){
	printf("[いちご100%%] %s", s);
}
*/
import "C"

func main() {
	s := C.CString("ふつう\n")
	C.review(s)
}
