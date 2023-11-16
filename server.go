package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"
)

type Page struct {
	Title string
	Body []byte
}

func (page *Page) save() error {
	filename:= page.Title + ".txt"
	return os.WriteFile(filename, page.Body, 0600)
}

func load(title string) (*Page, error) {
	filename:= title + ".txt"
	body,err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}


func renderTemplate(w http.ResponseWriter, name string, page *Page) {
	temp, _ := template.ParseFiles(name + ".html")
	temp.Execute(w, page)
}

func editHandler(w http.ResponseWriter, r *http.Request){
	title := r.URL.Path[len("/edit/"):]
	page, err:= load(title)
	if err != nil {
		page:= &Page{Title: title}
		fmt.Print(page.Title)
	}
	renderTemplate(w, "views/edit", page)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	page, _ := load(title)
	renderTemplate(w, "views/index", page)
}


func main(){

	p1 := &Page{Title:"TestingPage", Body: []byte("Testing sample body")}
	p1.save()

	// http.HandleFunc("/", handler)
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	
	fmt.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
	
}
