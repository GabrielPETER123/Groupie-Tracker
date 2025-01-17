package main

import (
	"Groupie-Tracker/API"
	"fmt"
	"net/http"
	"html/template"

)

var (
	Artists = []API.Artists{}
	Dates = API.IndexDates{}
	Locations = API.IndexLocations{}
	Relation = API.IndexRelation{}
	ListDates []string
	ListGroup []string
)

func GetAllApi() {
	Artists = API.GetArtists()
	Dates = API.GetDates()
	Locations = API.GetLocations()
	Relation = API.GetRelation()
}


//? List des Groupes
func listGroup() {
	for i := 0; i < len(Artists); i++ {
		ListGroup = append(ListGroup, Artists[i].Name)
	}
	fmt.Println(ListGroup)
}

//? List des Dates
func listDates() {
	for i := 0; i < len(Dates.Index); i++ {
		ListDates = append(ListDates, Dates.Index[i].Dates...)
	}
	fmt.Println(ListDates)
}

func main() {
	GetAllApi()
	//** Listes pour l'affichage */
	listGroup()
	listDates()

	//** Handlers */
	http.HandleFunc("/", homePageHandler)
    http.HandleFunc("/Artists", artistsPageHandler)
	http.HandleFunc("/Group", groupPageHandler)

	//** Static files */
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
    fmt.Print("(http://localhost:8080/)")
    http.ListenAndServe(":8080", nil)
}


func homePageHandler(w http.ResponseWriter, r *http.Request) {
    tmpl := template.Must(template.ParseFiles("index.html"))
    tmpl.Execute(w, nil)
}

func artistsPageHandler(w http.ResponseWriter, r *http.Request) {    
    tmpl := template.Must(template.ParseFiles("index.html"))
	tmpl.Execute(w, nil)
}

func groupPageHandler(w http.ResponseWriter, r *http.Request) {    
	tmpl := template.Must(template.ParseFiles("index.html"))
	tmpl.Execute(w, nil)
}
