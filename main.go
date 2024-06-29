package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)



type Movie struct {
	Id string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
}

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(movies); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index , item := range movies {
		if item.Id == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _ , item := range movies {
		if item.Id == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var movie Movie 
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.Id = strconv.Itoa(rand.Intn(100000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}
	

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)

	for index , item :=range movies {
		if item.Id == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie 
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.Id = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}

var movies []Movie

func main (){
	r := mux.NewRouter()
	

	
	movies = append(movies, Movie{
		Id: "1", 
		Isbn: "4598", 
		Title: "The Lord of the Rings", 
		Director: &Director{FirstName: "Jack", LastName: "Jhones"},
	})
	
	movies = append(movies, Movie{
		Id: "2", 
		Isbn: "1234", 
		Title: "Inception", 
		Director: &Director{FirstName: "Christopher", LastName: "Nolan"},
	})
	
	movies = append(movies, Movie{
		Id: "3", 
		Isbn: "5678", 
		Title: "The Matrix", 
		Director: &Director{FirstName: "Lana", LastName: "Wachowski"},
	})
	
	movies = append(movies, Movie{
		Id: "4", 
		Isbn: "91011", 
		Title: "Interstellar", 
		Director: &Director{FirstName: "Christopher", LastName: "Nolan"},
	})
	
	movies = append(movies, Movie{
		Id: "5", 
		Isbn: "121314", 
		Title: "The Godfather", 
		Director: &Director{FirstName: "Francis", LastName: "Coppola"},
	})
	
	r.HandleFunc("/movies", getMovies ).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies",createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")


	fmt.Println("starting server at port 3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}