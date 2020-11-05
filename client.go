package main

import (
	"log"
	"net/http"
)

type omdbClient struct {
	baseUrl string
	apiKey  string
}

func InitOmdbClient(baseUrl, apiKey string) *omdbClient {
	return &omdbClient{baseUrl: baseUrl, apiKey: apiKey}
}
func (c *omdbClient) Get(params string) *http.Response {
	resp, err := http.Get(c.baseUrl + "/?apiKey=" + c.apiKey + "&" + params)
	if err != nil {
		log.Println(err.Error())
	}
	defer resp.Body.Close()
	return resp
}
func (c *omdbClient) GetById(id string) *http.Response {
	params := "i=" + id
	return c.Get(params)
}
func (c *omdbClient) GetByIds(ids []string) string {
	for _, id := range ids {
		c.GetById(id)
	}
	return ""
}
