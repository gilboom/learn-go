package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sort"
)

type FileSystemStore struct {
	database io.Writer
	league   League
}

func NewFileSystemStore(file *os.File) (*FileSystemStore, error) {
	err := initializePlayerDBFile(file)
	if err != nil {
		return nil, fmt.Errorf("problem initializing player db file")
	}

	league, err := NewLeague(file)
	if err != nil {
		return nil, fmt.Errorf("problem loading player store from file %s, %v", file.Name(), err)
	}
	tape := &Tape{file: file}
	store := &FileSystemStore{database: tape, league: league}
	return store, nil
}

func initializePlayerDBFile(file *os.File) error {
	file.Seek(0, 0)

	stat, err := file.Stat()
	if err != nil {
		return fmt.Errorf("problem getting file info from file %s, %v", file.Name(), err)
	}

	if stat.Size() == 0 {
		file.Write([]byte("[]"))
		file.Seek(0, 0)
	}

	return nil
}

func (s *FileSystemStore) GetLeague() League {
	sort.Slice(s.league, func(i, j int) bool {
		return s.league[i].Wins > s.league[j].Wins
	})
	return s.league
}

func (s *FileSystemStore) GetPlayerScore(name string) int {
	player := s.GetLeague().Find(name)
	if player == nil {
		return 0
	}
	return player.Wins
}

func (s *FileSystemStore) RecordWin(name string) {
	player := s.league.Find(name)

	if player != nil {
		player.Wins++
	} else {
		s.league = append(s.league, Player{name, 1})
	}

	json.NewEncoder(s.database).Encode(s.league)
}
