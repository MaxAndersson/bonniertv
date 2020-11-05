package main

import (
	"log"
	"os"
	"testing"
)

func TestClientGet(t *testing.T) {
	client := InitOmdbClient(os.Getenv("OMDB_BASEURL"), os.Getenv("OMDB_APIKEY"))
	resp := client.Get("i=tt0133093")
	log.Println(resp)
}

func TestClientGetbyId(t *testing.T) {
	client := InitOmdbClient(os.Getenv("OMDB_BASEURL"), os.Getenv("OMDB_APIKEY"))
	resp := client.GetById("tt0133093")
	log.Println(resp)
}
func TestClientGetbyIds(t *testing.T) {
	client := InitOmdbClient(os.Getenv("OMDB_BASEURL"), os.Getenv("OMDB_APIKEY"))
	ids := []string{"tt0133093", "tt0133093", "tt0133093"}
	resp := client.GetByIds(ids)
	log.Println(resp)
}
