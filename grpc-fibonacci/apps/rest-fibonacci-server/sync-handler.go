package restfibonacciserver

import (
	"net/http"

	"github.com/gorilla/mux"
)

type SyncResponse struct {
	TimeTaken        string `json:"timeTaken"`
	FibonacciNumbers []int  `json:"fibonacciNumbers"`
}

func (a *App) fibonacciSyncHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	number
}
