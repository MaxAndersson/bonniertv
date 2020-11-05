package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sort"
)

type omdbClient struct {
	baseUrl string
	apiKey  string
}
type MovieResult struct { // borrowed from https://github.com/eefret/gomdb/blob/master/gomdb.go
	Title             string
	Year              string
	Rated             string
	Released          string
	Runtime           string
	Genre             string
	Director          string
	Writer            string
	Actors            string
	Plot              string
	Language          string
	Country           string
	Awards            string
	Poster            string
	Metascore         string
	ImdbRating        string
	ImdbVotes         string
	ImdbID            string
	Type              string
	TomatoMeter       string
	TomatoImage       string
	TomatoRating      string
	TomatoReviews     string
	TomatoFresh       string
	TomatoRotten      string
	TomatoConsensus   string
	TomatoUserMeter   string
	TomatoUserRating  string
	TomatoUserReviews string
	TomatoURL         string
	DVD               string
	BoxOffice         string
	Production        string
	Website           string
	Response          string
	Error             string
}

type byTitle []MovieResult

func (s byTitle) Len() int {
	return len(s)
}
func (s byTitle) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byTitle) Less(i, j int) bool {
	return s[i].Title < s[j].Title
}

func InitOmdbClient(baseUrl, apiKey string) *omdbClient {
	return &omdbClient{baseUrl: baseUrl, apiKey: apiKey}
}
func (c *omdbClient) Get(params string) (*MovieResult, error) {
	resp, err := http.Get(c.baseUrl + "/?apiKey=" + c.apiKey + "&" + params)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	message := new(MovieResult)
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	_ = decoder.Decode(&message)

	return message, nil
}
func (c *omdbClient) GetById(id string) (*MovieResult, error) {
	params := "i=" + id
	result, err := c.Get(params)
	if err != nil {
		return nil, err
	}
	return result, nil
}
func (c *omdbClient) GetByIds(ids []string) ([]MovieResult, error) {
	result := []MovieResult{}
	for _, id := range ids {
		resp, err := c.GetById(id)
		if err != nil {
			return nil, err
		}
		result = append(result, *resp)

	}
	sort.Sort(byTitle(result))
	return result, nil
}
