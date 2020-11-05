package main

import (
	"encoding/json"
	"log"
	"os"
	"testing"
)

func TestClientGet(t *testing.T) {
	client := InitOmdbClient(os.Getenv("OMDB_BASEURL"), os.Getenv("OMDB_APIKEY"))
	resp := client.Get("i=tt0133093")
	data, err := json.Marshal(resp)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(data))
}

func TestClientGetbyId(t *testing.T) {
	client := InitOmdbClient(os.Getenv("OMDB_BASEURL"), os.Getenv("OMDB_APIKEY"))
	resp := client.GetById("tt0133093")
	data, err := json.Marshal(resp)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(data))
}
func TestClientGetbyIds(t *testing.T) {
	client := InitOmdbClient(os.Getenv("OMDB_BASEURL"), os.Getenv("OMDB_APIKEY"))
	ids := []string{"tt0133093", "tt0816692", "tt1375666", "tt0172495", "tt0137523"}
	resp := client.GetByIds(ids)
	data, err := json.Marshal(resp)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(data))
}
