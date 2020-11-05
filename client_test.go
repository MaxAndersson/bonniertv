package main

import (
	"encoding/json"
	"log"
	"os"
	"sort"
	"testing"
)

func TestClientGet(t *testing.T) {
	client := InitOmdbClient(os.Getenv("OMDB_BASEURL"), os.Getenv("OMDB_APIKEY"))
	resp, err := client.Get("i=tt0133093")
	if resp.Title != "The Matrix" {
		t.Error("The matrix has you")
	}
	data, err := json.Marshal(resp)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(data))
}

func TestClientGetbyId(t *testing.T) {
	client := InitOmdbClient(os.Getenv("OMDB_BASEURL"), os.Getenv("OMDB_APIKEY"))
	resp, err := client.GetById("tt0133093")
	if resp.Title != "The Matrix" {
		t.Error("The matrix has you")
	}

	//data, err := json.Marshal(resp)
	if err != nil {
		log.Fatal(err)
	}

	//log.Println(string(data))
}
func TestClientGetbyIdsAndSorted(t *testing.T) {
	client := InitOmdbClient(os.Getenv("OMDB_BASEURL"), os.Getenv("OMDB_APIKEY"))
	ids := []string{"tt0133093", "tt0816692", "tt1375666", "tt0172495", "tt0137523"}
	resp, err := client.GetByIds(ids)
	isSorted := sort.IsSorted(byTitle(resp))
	if !isSorted {
		t.Error("not sorted")
	}
	//	data, err := json.Marshal(resp)
	if err != nil {
		log.Fatal(err)
	}

	//	log.Println(string(data))
}
