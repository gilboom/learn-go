package main

import (
	"fmt"
	"net/http"
	"strings"
	"sync"
)

type PlayerStore interface {
	GetPlayerScore(string) int
	RecordWin(string)
}

type PlayerServer struct {
	store PlayerStore
	mu    sync.Mutex
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

func (s *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	switch r.Method {
	case http.MethodPost:
		s.processWins(w, player)
	case http.MethodGet:
		s.showScore(w, player)
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
