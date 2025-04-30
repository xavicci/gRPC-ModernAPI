package restfibonacciserver

import "sync"

type AsyncStore struct {
	mu            sync.Mutex
	current       int
	requestdRange int
	numbers       []int
}
