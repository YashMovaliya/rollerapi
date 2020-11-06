package main

import "net/http"

type Coaster struct {
	Name			string `json:"name"`
	Manufacturer	string `json:"manufacturer"`
	ID				string `json:"id"`
	InPark			string `json:"inPark"`
	Height			string `json:"height"`
}

type coasterHandlers struct {
	store map[string] Coaster 
}

func (h *coasterHandlers) get(w http.ResponseWriter, r *http.Request){

}

func newCoasterHandlers() *coasterHandlers {
	return &coasterHandlers{
		store : map[string] Coaster{
			"id1": Coaster{
				Name: "Fury 325",
				Height: 99,
				ID: "id1",
				InPark: "Carowimds",
				Manufacturer: "B+M"
			},
		},
	}
}

func main() {
	coasterHandlers := newCoasterHandlers()
	http.HandleFunc("/coasters", coasterHandlers.get)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

