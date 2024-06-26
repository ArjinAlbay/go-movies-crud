package main

import (
"fmt"
"log"
"encoding/json"
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
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
}

func getMovies (w http.ResponseWriter, r *http.Request) {
w.Header().Set("Content-Type", "application/json")
json.NewEncoder(w).Encode(movies)
}

func deleteMovie (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params
	}
	

var movies []Movie

func main (){


	r := mux.NewRouter()

	movies = append(movies , Movie{Id: "1" , Isbn :"4598", Title:"the lord of the rings", Director : &Director(FirstName : "Jack" , LastName : "Jhones")})
	r.HandleFunc("/movies", getMovies ).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies",createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")


	fmt.println("starting server at port 3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}