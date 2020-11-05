package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

type omdbHandler struct {
	client *omdbClient
}

func (h omdbHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	log.Println("Getting the movies")

	io.WriteString(res, "hello")
}

func main() {
	omdb := InitOmdbClient(os.Getenv("OMDB_BASEURL"), os.Getenv("OMDB_APIKEY"))
	oh := omdbHandler{client: omdb}

	log.Println("Starting web service... ")
	mux := http.NewServeMux()
	mux.Handle("/", oh)

	http.ListenAndServe(":8080", mux)
}
