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

var DISCGOLFDATABASE *sql.DB

func main() {
	router := mux.NewRouter()

	var err error
	DISCGOLFDATABASE, err = sql.Open("sqlite3", "../database/discgolf.db")
	if err != nil {
		log.Fatal(err)
	}
	defer DISCGOLFDATABASE.Close()

	router.HandleFunc("/", hello).Methods(http.MethodPost)
	router.HandleFunc("/getdiscs", getDiscsHandler).Methods(http.MethodPost)
	router.HandleFunc("/getdiscbyname", getDiscByNameHandler).Methods(http.MethodPost)
	// getInnova
	// getDiscraft
	// getDynamicDiscs
	// getDiscByName
	// create a new database that includes a spot for an image, description, and a price + everything we already have

	http.ListenAndServe(":8000", router)
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hey there!"))
}

func getDiscsHandler(w http.ResponseWriter, r *http.Request) {
	discs := discgolfdb.GetAllDiscs(DISCGOLFDATABASE)

	discsJson, err := json.Marshal(discs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(discsJson)
}

func getDiscByNameHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	if name == "" {
		http.Error(w, "Name parameter is required", http.StatusBadRequest)
		return
	}

	disc, err := discgolfdb.GetDiscByName(DISCGOLFDATABASE, name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	discJson, err := json.Marshal(disc)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(discJson)
}
