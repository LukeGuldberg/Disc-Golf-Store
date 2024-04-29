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

	router.HandleFunc("/", hello).Methods("GET")
	router.HandleFunc("/getdiscs", getDiscsHandler).Methods("GET")
	router.HandleFunc("/getdiscbyname", getDiscByNameHandler).Methods("GET")
	router.HandleFunc("/getInnova", getInnovaHandler).Methods("GET")
	router.HandleFunc("/getDiscraft", getDiscraftHandler).Methods("GET")
	router.HandleFunc("/getDynamicDiscs", getDynamicDiscsHandler).Methods("GET")

	http.ListenAndServe(":8000", router)
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Content-Type", "application/json") 
}

func hello(w http.ResponseWriter, r *http.Request) {
	enableCors(&w) 
	w.Write([]byte("Hey there!"))
}

func getDiscsHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w) 
	discs := discgolfdb.GetAllDiscs(DISCGOLFDATABASE)
	discsJson, err := json.Marshal(discs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(discsJson)
}

func getDiscByNameHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	name := r.URL.Query().Get("name") 
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
	w.Write(discJson)
}

func getInnovaHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w) 
	discs, err := discgolfdb.GetDiscsByManufacturer(DISCGOLFDATABASE, "Innova")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	discsJson, err := json.Marshal(discs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(discsJson)
}

func getDiscraftHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	discs, err := discgolfdb.GetDiscsByManufacturer(DISCGOLFDATABASE, "Discraft")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	discsJson, err := json.Marshal(discs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(discsJson)
}

func getDynamicDiscsHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w) 
	discs, err := discgolfdb.GetDiscsByManufacturer(DISCGOLFDATABASE, "Dynamic Discs")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	discsJson, err := json.Marshal(discs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(discsJson)
}