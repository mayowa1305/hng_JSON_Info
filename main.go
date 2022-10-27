package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type internInfo struct {
	SlackUsername string `json:"slackUsername"`
	Backend       bool   `json:"backend"`
	Age           int    `json:"age"`
	Bio           string `json:"bio"`
}

var boluwatife []internInfo

func main() {
	boluwatife = append(boluwatife, internInfo{SlackUsername: "bolu_adx", Backend: true, Age: 22, Bio: "my name is adeyinka boluwatife and i'm a golang backend developer, i'm also a student of electrical and electronics engineering in federal university of agriculture abeokuta. i'm ready to give my best for this hng internship"})

	r := mux.NewRouter()

	r.HandleFunc("/internInfo", internInfo1).Methods("GET")

	fmt.Printf("starting server at port :8080\n")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func internInfo1(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(boluwatife)
}
