package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", getPokemon)
	err := http.ListenAndServe(":8080", nil)

	fmt.Println(err)
}

func getPokemon(resp http.ResponseWriter, req *http.Request) {
	request := GetPokemonRequest{}
	json.NewDecoder(req.Body).Decode(&request)

	response := Pokemon{
		Name: request.Name,
		Type: "feu",
	}

	json.NewEncoder(resp).Encode(response)
}

type Pokemon struct {
	Name  string `json:"name"`
	Type  string `json:"type"`
	Level int    `json:"level,omitempty"`
}

type GetPokemonRequest struct {
	Name string `json:"name"`
}
