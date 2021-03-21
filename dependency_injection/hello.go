package main

import (
	"fmt"
	"io"
	"net/http"
)

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func MyGreetHandler(writer http.ResponseWriter, request *http.Request) {
	Greet(writer, "world")
}

func main() {
	http.ListenAndServe(":3001", http.HandlerFunc(MyGreetHandler))
}
