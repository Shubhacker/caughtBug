package main

import (
	db "caughtBug/caughtBug/DB"
	"caughtBug/caughtBug/fields"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var database *sql.DB
var err error

func fetchBugInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "TestRestAPI")
	w.Header().Set("Content-Type", "application/json")
	bugInfo, err := db.FetchBugInformation(database)
	if err != nil {
		log.Println("Error from Database : ", bugInfo)
	}
	json.NewEncoder(w).Encode(bugInfo)
}

func fetchBugInfoWithFilter(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)
	var bugFilter fields.BugFilter

	Id := param["Id"]
	name := param["app"]
	desc := param["desc"]
	posted := param["posted"]
	present := param["present"]

	bugFilter.UniqueId = &Id
	bugFilter.ApplicationName = &name
	bugFilter.BugDescription = &desc
	bugFilter.PostedBy = &posted
	bugFilter.StillPresent = &present

	bugInfo, err := db.FetchBugInformationWithFilter(database, bugFilter)
	if err != nil {
		log.Println("Error from DB : ", bugInfo)
	}

	json.NewEncoder(w).Encode(bugInfo)
}

func main() {
	database, err = db.DBSetup()
	if err != nil {
		log.Println("Error : ", err.Error())
	}
	HandleRequest()
}

func HandleRequest() {
	// Init Router
	r := mux.NewRouter()

	// Endpoints
	r.HandleFunc("/api/FetchBugInfo", fetchBugInfo).Methods("GET")
	r.HandleFunc("/api/FetchBugInfoWithFilter/{Id}/{app}", fetchBugInfoWithFilter).Methods("GET")

	// Server logic
	log.Println("Server running at : 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
