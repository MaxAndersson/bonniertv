package main

import (
  "log"
  "net/http"
  "io"
)
type omdbHandler int
func(h omdbHandler) ServeHTTP(res http.ResponseWriter, req *http.Request){
  io.WriteString(res, "hello")
} 

func main(){
  var oh omdbHandler
  log.Println("Starting web service...")
  mux := http.NewServeMux()
  mux.Handle("/",oh)
  
  http.ListenAndServe(":8080", mux)
}
