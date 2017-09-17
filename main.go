package main

import (
	//"fmt"
	//"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
)

type ReverseString struct {
	ReversedString string `json:"reversedstring"`
}

/*
func CreatePersonEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var person Person
	_ = json.NewDecoder(req.Body).Decode(&person)
	person.ID = params["id"]
	people = append(people, person)
	json.NewEncoder(w).Encode(people)
}

func DeletePersonEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for index, item := range people {
		if item.ID == params["id"] {
			people = append(people[:index], people[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(people)
}
*/
func ToReverseEndPoint(w http.ResponseWriter, req *http.Request) {

	var originalString string
	var reversedString string

	params := mux.Vars(req)

	//Load the original string with the word passed in the URL.
	originalString = params["toReverse"]

	//Traverse through string, character by character.
	//In this loop we only need the position of each character, and not the actual character.
	//Could have used a loop in this case, but wanted to get into the habit of using the 'range' keyword.
	for i, _ := range originalString {

		//The next line reverses the string by using slices building up reversedString.
		//For a 6 character string:
		//	Iteration 1 of the loop = originalString[5:6]
		//	Iteration 2 of the loop = originalString[4:5]
		//	Iteration 3 of the loop = originalString[3:4]
		//	Iteration 4 of the loop = originalString[2:3]
		//	Iteration 5 of the loop = originalString[1:2]
		//	Iteration 6 of the loop = originalString[0:1]
		reversedString = reversedString + originalString[len(originalString) -1 - i:len(originalString) - i]
	}

	//Create an instance of the ReverseString struct and load it.
	reverseString := ReverseString{ReversedString: reversedString}

	//De-marshall it.
	json.NewEncoder(w).Encode(reverseString)
}


func main() {
	router := mux.NewRouter()

	router.HandleFunc("/reverse-string/{toReverse}", ToReverseEndPoint).Methods("GET")
	//router.HandleFunc("/people/{id}", GetPersonEndpoint).Methods("GET")
	//router.HandleFunc("/people/{id}", CreatePersonEndpoint).Methods("POST")
	//router.HandleFunc("/people/{id}", DeletePersonEndpoint).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":12345", router))
}