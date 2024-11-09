package main

import (
	"html/template"
	"fmt"
	"os"
	"log"
	"net/http"
)

type Page struct{
	Title string
	Body []byte
}

// * Function to save pages

func (p *Page) save() error {
	filename := p.Title + ".txt"
	fmt.Println(filename)
	return os.WriteFile(filename, p.Body, 0600)
}

// * Function to load pages
func loadPage(title string) (*Page, error){
	filename := title + ".txt"
	body, err := os.ReadFile((filename))

	if err != nil{
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

// * Let's write a function to render templates
func templateRenderer(w http.ResponseWriter, tmpl string ,p *Page){
	t, _ := template.ParseFiles(tmpl + ".html")
	t.Execute(w, p)
}

// * Note we use _to ignore the error return value 
func viewHandler(w http.ResponseWriter, r *http.Request){
	title := r.URL.Path[len("/view/"):]
	p, _ := loadPage(title)
	templateRenderer(w, "view", p)
}

// * Edit Handler for making edits possible
func editHandler(w http.ResponseWriter, r *http.Request){
	title := r.URL.Path[len("/edit/"):]
	p, err := loadPage(title)

	if err == nil{
		p = &Page{Title: title}
	}

	templateRenderer(w, "edit", p)
}

func main() {
    http.HandleFunc("/view/", viewHandler)
    http.HandleFunc("/edit/", editHandler)
    // http.HandleFunc("/save/", saveHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}