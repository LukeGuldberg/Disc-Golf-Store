package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"discgolf/database"
	"discgolf/models"

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
	

	http.ListenAndServe(":8000", router)
}
func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hey there!"))
}
