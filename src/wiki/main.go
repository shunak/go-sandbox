package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// data structure of wiki
type Page struct {
	Title string // title
	Body  []byte // inner title
}

// set path address and get string length
const lenPath = len("/view/")

// response handler
func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[lenPath:]
	p, _ := loadPage(title)
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[lenPath:]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	fmt.Fprintf(w, "<h1>Editing %s</h1>"+
		"<form action=\"/save/%s\" method=\"POST\">"+
		"<textarea name=\"body\">%s</textarea><br>"+
		"<input type=\"submit\" value=\"Save\">"+
		"</form>", p.Title, p.Title, p.Body)
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
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler) // make edit page
	http.ListenAndServe(":8080", nil)
}
