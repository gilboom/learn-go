package main

import (
	"log"
	"net/http"
	"os"
)

const dbFileName = "game.db.json"

func main() {
	file, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("Open file failed %v", err)
	}

	store, err := NewFileSystemStore(file)
	if err != nil {
		log.Fatalf("problem creating file system store %s, %v", file.Name(), err)
	}

	server := NewPlayerServer(store)

	if err := http.ListenAndServe(":3001", server); err != nil {
		log.Fatalf("could not listen on port 3001 %v", err)
	}
}
