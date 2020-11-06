package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"sort"
	"time"

	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
)

type omdbHandler struct {
	client *omdbClient
	ring   *redis.Ring
	cache  *cache.Cache
	ctx    context.Context
}

type MovieResponse struct {
	data []MovieResult
}

func (h omdbHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	log.Println("Getting the movies")
	ids := []string{"tt0133093", "tt0816692", "tt1375666", "tt0172495", "tt0137523"}
	misses := []string{}
	results := []MovieResult{}

	for _, id := range ids {
		var movie MovieResult
		err := h.cache.Get(h.ctx, id, &movie)
		if err == nil {
			results = append(results, movie)
			log.Println("Cache hit : ", movie)
		} else {
			misses = append(misses, id)
			log.Println("Cache miss : ", id)
		}
	}
	resp, err := h.client.GetByIds(misses)
	for _, m := range resp {
		if err := h.cache.Set(&cache.Item{
			Ctx:   h.ctx,
			Key:   m.ImdbID,
			Value: m,
			TTL:   time.Minute,
		}); err != nil {
			panic(err)
		}

		results = append(results, m)

	}
	sort.Sort(byTitle(results))

	data, err := json.Marshal(results)

	if err != nil {
		log.Println(err)
	}
	res.Write(data)
}

func enforceSecret(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		secret := r.Header.Get("X-Secret")

		if secret != "1234" {
			return

		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	ring := redis.NewRing(&redis.RingOptions{
		Addrs: map[string]string{
			"server1": ":6379",
		},
	})
	cache := cache.New(&cache.Options{
		Redis:      ring,
		LocalCache: cache.NewTinyLFU(1000, time.Minute),
	})
	ctx := context.TODO()

	omdb := InitOmdbClient(os.Getenv("OMDB_BASEURL"), os.Getenv("OMDB_APIKEY"))
	oh := omdbHandler{client: omdb, ring: ring, cache: cache, ctx: ctx}

	log.Println("Starting web service... ")
	mux := http.NewServeMux()
	mux.Handle("/", enforceSecret(oh))

	http.ListenAndServe(":8080", mux)
}
