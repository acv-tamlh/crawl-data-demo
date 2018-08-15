package main

import (
    "encoding/json"
    "log"
    "net/http"
    "github.com/gorilla/mux"
)

type Person struct {
	ID        string   `json:"id,omitempty`
	Firstname string   `json:"firstname,omitempty`
	Lastname  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}
type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

var people []Person

func main()  {
  people = append(people, Person{ID: "1", Firstname: "John", Lastname: "Doe", Address: &Address{City: "City X", State: "State X"}})
  people = append(people, Person{ID: "2", Firstname: "Koko", Lastname: "Doe", Address: &Address{City: "City Z", State: "State Y"}})
  people = append(people, Person{ID: "3", Firstname: "Francis", Lastname: "Sunday"})

  router := mux.NewRouter()
  router.HandleFunc("/people", GetPeople).Methods("GET")
  router.HandleFunc("/people/{id}", GetPerson).Methods("GET")
  router.HandleFunc("/people/{id}", CreatePerson).Methods("POST")
  router.HandleFunc("/people/{id}", DeletePerson).Methods("DELETE")
  log.Fatal(http.ListenAndServe(":8000", router))
}

func GetPeople(w http.ResponseWriter, r *http.Request)  {
  params := mux.Vars(r)
  for _, item := range people {
    if item.ID == params["id"] {
        json.NewEncoder(w).Encode(item)
        return
    }
  }
  json.NewEncoder(w).Encode(&Person{})
}
func GetPerson(w http.ResponseWriter, r *http.Request)  {
}
func CreatePerson(w http.ResponseWriter, r *http.Request)  {
}
func DeletePerson(w http.ResponseWriter, r *http.Request)  {
}

// func GetPeople(w http.ResponseWriter, r *http.Request) {}
// func GetPerson(w http.ResponseWriter, r *http.Request) {}
// func CreatePerson(w http.ResponseWriter, r *http.Request) {}
// func DeletePerson(w http.ResponseWriter, r *http.Request) {}
