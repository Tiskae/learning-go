package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Planet struct {
	Name       string `json:"name"`
	Population string `json:"population"`
	Terrain    string `json:"terrain"`
}

type Person struct {
	Name         string `json:"name"`
	HomeworldURL string `json:"homeworld"`
	Homewrorld   Planet
}

type AllPeople struct {
	People []Person `json:"results"`
}

func getHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home page")
}

const BaseURL = "https://swapi.co/api/"

func getPeople(w http.ResponseWriter, r *http.Request) {
	res, err := http.Get(BaseURL + "people")

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Print("Failed to request Star Wars people")
	}

	fmt.Println(res)

	bytes, err := io.ReadAll(res.Body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	var people AllPeople

	fmt.Println(string(bytes))
	if err := json.Unmarshal(bytes, &people); err != nil {
		fmt.Println("Error parsing json", err)
	}

	fmt.Println(people)
}

func main() {
	http.HandleFunc("/", getHome)
	http.HandleFunc("/people/", getPeople)
	fmt.Println("Serving on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
