package restfibonacciserver

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type App struct {
	asyncStores map[string]*AsyncStore
}

func NewApp() *App {
	return &App{
		asyncStores: make(map[string]*AsyncStore),
	}
}

func (a *App) Start() {
	router := mux.NewRouter()
	router.Use(generateRequestIDMiddleware)
	router.HandleFunc("fibonacci/sync/{number:[0-9]+}", a.)

}

func generateRequestIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("request-id") == "" {
			requestID := uuid.New().String()
			r.Header.Add("request-id", requestID)
		}
		next.ServeHTTP(w, r)
	})
}
