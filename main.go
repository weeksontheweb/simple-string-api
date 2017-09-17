package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
)

type ReverseString struct {
	ReversedString string `json:"reversedstring"`
}

type NextInAsciiString struct {
	NextInAsciiDoneString string `json:"nextinasciistring"`
}

type PreviousInAsciiString struct {
	PreviousInAsciiDoneString string `json:"previousinasciistring"`
}

//Reverses the input string.
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

//Converts each character in the input string to the next ASCII value.
func NextInAsciiEndPoint(w http.ResponseWriter, req *http.Request) {

	var upscaledString string

	params := mux.Vars(req)

	//Convert the word passed in the URL to an ASCII array
	byteArray := []byte(params["toUpscale"])

	//Traverse through string, character by character.
	for i, _ := range byteArray {

		upscaledString = upscaledString + string(byteArray[i] + 1)
	}

	//Create an instance of the UpscaleString struct and load it.
	nextInAsciiString := NextInAsciiString{NextInAsciiDoneString: upscaledString}

	json.NewEncoder(w).Encode(nextInAsciiString)
}

//Converts each character in the input string to the next ASCII value.
func PreviousInAsciiEndPoint(w http.ResponseWriter, req *http.Request) {

	var downscaledString string

	params := mux.Vars(req)

	//Convert the word passed in the URL to an ASCII array
	byteArray := []byte(params["toDownscale"])

	//Traverse through string, character by character.
	for i, _ := range byteArray {

		downscaledString = downscaledString + string(byteArray[i] - 1)
	}

	//Create an instance of the UpscaleString struct and load it.
	previousInAsciiString := PreviousInAsciiString{PreviousInAsciiDoneString: downscaledString}

	json.NewEncoder(w).Encode(previousInAsciiString)
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/reverse-string/{toReverse}", ToReverseEndPoint).Methods("GET")
	router.HandleFunc("/next-in-ascii/{toUpscale}", NextInAsciiEndPoint).Methods("GET")
	router.HandleFunc("/previous-in-ascii/{toDownscale}", PreviousInAsciiEndPoint).Methods("GET")

	log.Fatal(http.ListenAndServe(":12345", router))
}