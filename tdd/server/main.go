package main

import (
	"log"
	"net/http"
)

func main() {
	store := NewInMemoryPlayerStore()
	server := NewPlayerServer(store)
	log.Println("Server started!")
	log.Fatal(http.ListenAndServe(":5000", server))
}
