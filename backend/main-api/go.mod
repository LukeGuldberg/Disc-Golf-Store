module discgolf/main

go 1.22.0

require (
	discgolf/discgolfdb v0.0.0-00010101000000-000000000000
	discgolf/models v0.0.0-00010101000000-000000000000
	github.com/gorilla/mux v1.8.1
	github.com/mattn/go-sqlite3 v1.14.22
)

replace discgolf/discgolfdb => ../database

replace discgolf/models => ../models
