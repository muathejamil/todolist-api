package main

import (
    "net/http"
)


func main() {
    mux := http.ServeMux()
	mux.Handle("/todos/", &TodosHandler)
	http.ListenAndServe("localhost:8000", mux)
}