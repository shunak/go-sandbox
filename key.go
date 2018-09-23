package main

import "curses"
import "fmt"

func main(){
 	curses.Kinit()
	for{
		if curses.Kbhit() == 1{
			moji :=curses.Getch()
			fmt.Println( string(moji) )
			break
		}
	}
	curses.Kclose()
}

