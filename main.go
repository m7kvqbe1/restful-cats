package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Cat struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var cats = []Cat{
	{ID: "1", Name: "Whiskers", Age: 3},
	{ID: "2", Name: "Shadow", Age: 5},
}

func main() {
	http.HandleFunc("/cats", handleCats)
	http.HandleFunc("/cats/", handleCat)
	log.Println("Listening on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleCats(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getCats(w, r)
	case http.MethodPost:
		createCat(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Method Not Allowed")
	}
}

func handleCat(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/cats/"):]

	switch r.Method {
	case http.MethodGet:
		getCat(w, r, id)
	case http.MethodPut:
		updateCat(w, r, id)
	// case http.MethodPatch:
	// 	updateCatPartial(w, r, id)
	case http.MethodDelete:
		deleteCat(w, r, id)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Method Not Allowed")
	}
}

func getCats(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cats)
}

func createCat(w http.ResponseWriter, r *http.Request) {
	var newCat Cat

	if err := json.NewDecoder(r.Body).Decode(&newCat); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Bad Request")
		return
	}

	cats = append(cats, newCat)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newCat)
}

func getCat(w http.ResponseWriter, _ *http.Request, id string) {
	for _, cat := range cats {
		if cat.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(cat)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "Not Found")
}

func updateCat(w http.ResponseWriter, r *http.Request, id string) {
	var updatedCat Cat

	if err := json.NewDecoder(r.Body).Decode(&updatedCat); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Bad Request")
		return
	}

	for i, cat := range cats {
		if cat.ID == id {
			cats[i] = updatedCat
			json.NewEncoder(w).Encode(updatedCat)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "Not Found")
}

// func updateCatPartial(w http.ResponseWriter, r *http.Request, id string) {
// 	w.WriteHeader(http.StatusNotImplemented)
// 	fmt.Fprintf(w, "Not Implemented")
// }

func deleteCat(w http.ResponseWriter, _ *http.Request, id string) {
	for i, cat := range cats {
		if cat.ID == id {
			cats = append(cats[:i], cats[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "Not Found")
}
