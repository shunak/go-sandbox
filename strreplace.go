package main

// import "curses"
import "fmt"
import "strings"

func main() {

	s := strings.Replace("郷に入っては郷に従え", "郷", "Go", 1)

	fmt.Println(s)

}
