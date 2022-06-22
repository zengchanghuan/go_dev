package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type IndexData struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var indexData IndexData
	indexData.Title = "go learning"
	indexData.Desc = "go code"
	jsonStr, _ := json.Marshal(indexData)
	w.Write(jsonStr)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/", index)
	error := server.ListenAndServe()
	if error != nil {
		log.Println(error)
	}

}
