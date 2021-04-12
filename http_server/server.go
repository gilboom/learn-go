package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

type PlayerStore interface {
	GetPlayerScore(string) int
	RecordWin(string)
	GetLeague() League
}

type PlayerServer struct {
	store PlayerStore
	mu    sync.Mutex
	http.Handler
}

func NewPlayerServer(store PlayerStore) *PlayerServer {
	s := &PlayerServer{store: store}
	router := http.NewServeMux()

	router.Handle("/league", http.HandlerFunc(s.leagueHandler))
	router.Handle("/players/", http.HandlerFunc(s.playersHandler))

	s.Handler = router
	return s
}

func (s *PlayerServer) showScore(w http.ResponseWriter, player string) {
	score := s.store.GetPlayerScore(player)
	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, score)
}

func (s *PlayerServer) processWins(w http.ResponseWriter, player string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
}

type Player struct {
	Name string
	Wins int
}

const jsonContentType = "application/json"

func (s *PlayerServer) leagueHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("content-type", jsonContentType)
	json.NewEncoder(writer).Encode(s.getLeagueTable())
	//writer.WriteHeader(http.StatusOK)
}

func (s *PlayerServer) getLeagueTable() []Player {
	return s.store.GetLeague()
}

func (s *PlayerServer) playersHandler(writer http.ResponseWriter, request *http.Request) {
	player := request.URL.Path[len("/players/"):]
	switch request.Method {
	case http.MethodPost:
		s.processWins(writer, player)
	case http.MethodGet:
		s.showScore(writer, player)
	}
}

//func GetPlayerScore(name string) string {
//	if name == "gilboom" {
//		return "20"
//	}
//
//	if name == "phantom" {
//		return "10"
//	}
//
//	return ""
//}
