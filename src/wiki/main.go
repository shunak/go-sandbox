package main

import (
	"fmt"
	"io/ioutil"
)

// data structure of wiki
type Page struct {
	Title string // title
	Body  []byte // inner title
}

// save method of txt file
func (p *Page) save() error {
	// make text file by Title name and save the file.
	filename := p.Title + ".txt"
	// 0600 is permission settings, 0600 is permission which your own is permitted.
	return ioutil.WriteFile(filename, p.Body, 0600)
}

// load file name from texttitle and return new Page pointer
func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	// if get err, make body value nil and return as error
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func main() {
	p1 := &Page{Title: "TestPage", Body: []byte("This is a sample page.")}
	p1.save()
	p2, _ := loadPage("TestPage")
	fmt.Println(string(p2.Body))
}
