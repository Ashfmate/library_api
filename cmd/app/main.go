package main

import (
	"log"
	"net/http"

	"github.com/Ashfmate/library_api/api"
)

func main() {
	http.HandleFunc("/getbooks", api.GetData)
	http.HandleFunc("/postbooks", api.PostData)

	log.Fatal(http.ListenAndServe(":8080", nil))
}