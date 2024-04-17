package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"discgolf/discgolfdb"
	// "discgolf/models"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

var DISCGOLFDATBASE *sql.DB

func main() {
	router := mux.NewRouter()

	var err error
	DISCGOLFDATBASE, err = sql.Open("sqlite3", "../database/discgolf.db")
	if err != nil {
		log.Fatal(err)
	}
	defer DISCGOLFDATBASE.Close()

	router.HandleFunc("/", hello).Methods(http.MethodPost)
	router.HandleFunc("/getdiscs", getDiscsHandler).Methods(http.MethodPost)

	http.ListenAndServe(":8000", router)
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hey there!"))
}

func getDiscsHandler(w http.ResponseWriter, r *http.Request) {
	discs := discgolfdb.GetAllDiscs(DISCGOLFDATBASE)

	discsJson, err := json.Marshal(discs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(discsJson)
}
