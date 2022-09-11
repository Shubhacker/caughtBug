package main

import (
	db "caughtBug/caughtBug/DB"
	"fmt"
	"log"
	"net/http"
)

func main() {
	db.DBSetup()
	HandleRequest()
}

func HandleRequest() {
	http.HandleFunc("/", TestRestApi)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func TestRestApi(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "TestRestAPI")
}
