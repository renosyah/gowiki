package main
     	
  import (

	"html/template"
	"io/ioutil"
	"net/http"
)
    	
    	type Page struct {
   	Title string
    	Body  []byte
    	}
    
    	func (p *Page) save() error {
    		filename := p.Title + "test.txt"
    		return ioutil.WriteFile(filename, p.Body, 0600)
    	}
    	
    	func loadPage(title string) (*Page, error) {
    		filename := title + "test.txt"
    		body, err := ioutil.ReadFile(filename)
    		if err != nil {
    			return nil, err
    		}
    		return &Page{Title: title, Body: body}, nil
    	}

func editHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/edit/"):]
    p, err := loadPage(title)
    if err != nil {
        p = &Page{Title: title}
    }
    t, _ := template.ParseFiles("edit.html")
    t.Execute(w, p)
}


func viewHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/view/"):]
    p, _ := loadPage(title)
    t, _ := template.ParseFiles("view.html")
    t.Execute(w, p)
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/save/"):]
    body := r.FormValue("body")
    p := &Page{Title: title, Body: []byte(body)}
    p.save()
    http.Redirect(w, r, "/view/"+title, http.StatusFound)
}
    	
    	func main() {
    	http.HandleFunc("/view/", viewHandler)
    http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)
    		http.ListenAndServe(":7070", nil)
    	}