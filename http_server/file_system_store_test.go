package main

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestFileSystemStore(t *testing.T) {

	t.Run("works with empty file", func(t *testing.T) {
		file, cleanDatabase := createTempFile(t, "")
		defer cleanDatabase()

		_, err := NewFileSystemStore(file)
		assertNotError(t, err)
	})

	t.Run("get sorted league", func(t *testing.T) {
		//database := strings.NewReader(`[
		//	{"Name": "GilBoom", "Wins": 10},
		//	{"Name": "Phantom", "Wins": 11}
		//]`)

		file, removeFile := createTempFile(t, `[
			{"Name": "GilBoom", "Wins": 10},
			{"Name": "Phantom", "Wins": 11}
		]`)
		defer removeFile()

		store, err := NewFileSystemStore(file)
		assertNotError(t, err)

		got := store.GetLeague()

		want := []Player{
			{"Phantom", 11},
			{"GilBoom", 10},
		}

		got = store.GetLeague()
		assertLeague(t, got, want)
	})

	t.Run("get player score", func(t *testing.T) {
		//database := strings.NewReader(`[
		//	{"Name": "GilBoom", "Wins": 10},
		//	{"Name": "Phantom", "Wins": 11}
		//]`)

		file, removeFile := createTempFile(t, `[
			{"Name": "GilBoom", "Wins": 10},
			{"Name": "Phantom", "Wins": 11}
		]`)
		defer removeFile()

		store, err := NewFileSystemStore(file)
		assertNotError(t, err)

		got := store.GetPlayerScore("GilBoom")
		want := 10
		assertScoreEquals(t, got, want)
	})

	t.Run("store wins for existing players", func(t *testing.T) {
		file, removeFile := createTempFile(t, `[
			{"Name": "GilBoom", "Wins": 10},
			{"Name": "Phantom", "Wins": 11}
		]`)
		defer removeFile()

		store, err := NewFileSystemStore(file)
		assertNotError(t, err)

		store.RecordWin("GilBoom")

		got := store.GetPlayerScore("GilBoom")
		want := 11

		assertScoreEquals(t, got, want)
	})

	t.Run("store wins for new players", func(t *testing.T) {
		file, removeFile := createTempFile(t, `[
			{"Name": "GilBoom", "Wins": 10},
			{"Name": "Phantom", "Wins": 11}
		]`)
		defer removeFile()

		store, err := NewFileSystemStore(file)
		assertNotError(t, err)

		store.RecordWin("Bob")

		got := store.GetPlayerScore("Bob")
		want := 1

		assertScoreEquals(t, got, want)
	})
}

func assertScoreEquals(t testing.TB, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("got score %d but want score %d", got, want)
	}
}

func assertNotError(t testing.TB, err error) {
	t.Helper()

	if err != nil {
		t.Fatalf("didn't expect an error but got one, %v", err)
	}
}

func createTempFile(t testing.TB, initialData string) (*os.File, func()) {
	t.Helper()

	file, err := ioutil.TempFile("", "db")
	if err != nil {
		t.Fatalf("could not create temp file %v", err)
	}

	file.Write([]byte(initialData))

	removeFile := func() {
		file.Close()
		os.Remove(file.Name())
	}

	return file, removeFile
}
