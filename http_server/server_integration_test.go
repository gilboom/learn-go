package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecordWinsAndRetrievingThem(t *testing.T) {
	file, removeFile := createTempFile(t, "[]")
	defer removeFile()

	store, err := NewFileSystemStore(file)
	assertNotError(t, err)

	server := NewPlayerServer(store)
	player := "GilBoom"

	server.ServeHTTP(httptest.NewRecorder(), newRecordWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newRecordWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newRecordWinRequest(player))

	t.Run("get score", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newGetScoreRequest(player))

		assertResponseStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "3")
	})

	t.Run("get league", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newLeagueRequest())
		assertResponseStatus(t, response.Code, http.StatusOK)

		got := getLeagueFromResponse(t, response.Body)
		want := []Player{
			{player, 3},
		}
		assertLeague(t, got, want)
	})
}
