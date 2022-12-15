package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}
type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie
func getMovies(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	json NewEncoder(w).Encode(movies)
}
func deleteMovie(w http.ResponseWriter, r *http.Request ) {
	w.Header().set("Content-Type","application/json")
	params := mux.Vars(r)
	for index,item := range movies {
		if item.ID == params["id"]{
			movies = append(movies[:index],movies[index+1]... )
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
	
}
func getMovie(w http.ResponseWriter,r *http.Request){
w.Header().Set("Content-type","application/json")
params := mux.Vars(r)
fot _, item := range movies{
	if item.ID == params["id"]{
		json.NewEncoder(w).Encode(item)
		return
	}
}
}

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn: "45454", Title: "movie", Director: &Director{Firstname: "john", Lastname: "doe"}})
	movies = append(movies, Movie{ID: "1", Isbn: "45454", Title: "movie", Director: &Director{Firstname: "john", Lastname: "doe"}})
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("starting server ar port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}
