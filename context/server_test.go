package context

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type SpyStore struct {
	response string
	//cancelled bool
	t *testing.T
}

func (s *SpyStore) Fetch(ctx context.Context) (string, error) {
	//time.Sleep(time.Millisecond * 100)
	//return s.response

	data := make(chan string, 1)

	go func() {
		var result string
		for _, c := range s.response {
			select {
			case <-ctx.Done():
				s.t.Log("spy store got cancelled")
				return
			default:
				time.Sleep(10 * time.Millisecond)
				result += string(c)
			}
		}
		data <- result
	}()

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case res := <-data:
		return res, nil
	}
}

//func (s *SpyStore) Cancel() {
//	s.cancelled = true
//}

//func (s *SpyStore) assertWasCancelled() {
//	s.t.Helper()
//	if !s.cancelled {
//		s.t.Errorf("store was not told to cancel")
//	}
//}

//func (s *SpyStore) assertWasNotCancelled() {
//	s.t.Helper()
//	if s.cancelled {
//		s.t.Error("it should not have cancelled the store")
//	}
//}

type SpyResponseWriter struct {
	written bool
}

func (s *SpyResponseWriter) Header() http.Header {
	s.written = true
	return nil
}

func (s *SpyResponseWriter) Write([]byte) (int, error) {
	s.written = true
	return 0, nil
}

func (s *SpyResponseWriter) WriteHeader(statusCode int) {
	s.written = true
}

func TestServer(t *testing.T) {
	t.Run("returns data from store", func(t *testing.T) {
		data := "hello, world"
		//store := &SpyStore{data, false, t}
		store := &SpyStore{data, t}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		if response.Body.String() != data {
			t.Errorf("got %q but want %q", response.Body.String(), data)
		}

		//store.assertWasNotCancelled()
	})

	t.Run("tells store to cancel work if request is cancelled", func(t *testing.T) {
		data := "hello, world"
		//store := &SpyStore{data, false, t}
		store := &SpyStore{data, t}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		request = request.WithContext(cancellingCtx)

		response := &SpyResponseWriter{written: false}

		svr.ServeHTTP(response, request)

		//store.assertWasCancelled()

		if response.written {
			t.Errorf("a response should not have been written")
		}
	})
}
