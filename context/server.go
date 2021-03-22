package context

import (
	"context"
	"fmt"
	"net/http"
)

type Store interface {
	Fetch(ctx context.Context) (string, error)
	//Cancel()
}

func Server(store Store) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		//ctx := request.Context()
		//
		//dataChan := make(chan string, 1)
		//
		//go func() {
		//	dataChan <- store.Fetch()
		//}()
		//
		//select {
		//case data := <-dataChan:
		//	fmt.Fprint(writer, data)
		//case <-ctx.Done():
		//	store.Cancel()
		//}

		data, err := store.Fetch(request.Context())

		if err != nil {
			return
		}

		fmt.Fprint(writer, data)
	}
}
