package main

import (
	"encoding/json"
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
func (c *omdbClient) Get(params string) *map[string]interface{} {
	resp, err := http.Get(c.baseUrl + "/?apiKey=" + c.apiKey + "&" + params)
	if err != nil {
		log.Println(err.Error())
	}
	message := make(map[string]interface{})
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	_ = decoder.Decode(&message)

	return &message
}
func (c *omdbClient) GetById(id string) *map[string]interface{} {
	params := "i=" + id
	return c.Get(params)
}
func (c *omdbClient) GetByIds(ids []string) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, id := range ids {
		resp := c.GetById(id)
		result = append(result, *resp)

	}
	return result
}
