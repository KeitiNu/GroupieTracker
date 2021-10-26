package main

import (
	"fmt"
	"groupie-tracker/create"
	"html/template"
	"net/http"
)

func main() {
	http.Handle("/templates/", http.StripPrefix("/templates/", http.FileServer(http.Dir("./templates/"))))
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./css/"))))

	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/artist/", ArtistHandler)

	fmt.Println("Server listening on port 3000")
	http.ListenAndServe(":3000", nil)

}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {

		if r.URL.Path != "/" {
			t, _ := template.ParseFiles("templates/error.html")
			t.Execute(w, http.StatusNotFound)	
			return
		}

		t, err := template.ParseFiles("templates/index.html")
		if err != nil {
			http.Error(w, "500: internal server error", http.StatusInternalServerError)
			return
		}
		artist, err := create.GetArtists()
		if err != nil {
			http.Error(w, "500: internal server error", http.StatusInternalServerError)
			return
		}
		t.Execute(w, artist)
	} else {
		http.Error(w, "400: Bad Request", http.StatusBadRequest)
		return
	}
}

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {

		t, err := template.ParseFiles("templates/profile.html")
		if err != nil {
			http.Error(w, "500: internal server error", http.StatusInternalServerError)
			return
		}
		artist, err := create.GetArtists()
		if err != nil {
			http.Error(w, "500: internal server error", http.StatusInternalServerError)
			return
		}

		urlString := string(r.URL.Path)[8:]

		for i, v := range artist {
			if v.Name == urlString {
				t.Execute(w, artist[i])
				return
			}
		}

		http.Error(w, "404 page not found.", http.StatusNotFound)
		return
	} else {
		http.Error(w, "400: bad request.", http.StatusBadRequest)
	}

}
