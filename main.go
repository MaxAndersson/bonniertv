package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type omdbHandler struct {
	client *omdbClient
}

func (h omdbHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	log.Println("Getting the movies")
	ids := []string{"tt0133093", "tt0816692", "tt1375666", "tt0172495", "tt0137523"}
	resp := h.client.GetByIds(ids)
	data, err := json.Marshal(resp)
	if err != nil {
		log.Println(err)
	}
	res.Write(data)
}

func main() {
	omdb := InitOmdbClient(os.Getenv("OMDB_BASEURL"), os.Getenv("OMDB_APIKEY"))
	oh := omdbHandler{client: omdb}

	log.Println("Starting web service... ")
	mux := http.NewServeMux()
	mux.Handle("/", oh)

	http.ListenAndServe(":8080", mux)
}
