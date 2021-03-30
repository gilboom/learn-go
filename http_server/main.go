package main

import (
	"log"
	"net/http"
)

func main() {
	store := NewInMemoryPlayerStore()
	server := &PlayerServer{store: store}
	log.Fatal(http.ListenAndServe(":3001", server))
}
