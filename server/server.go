package server

import (
	bs "blogging/service"
	"sync"
)

type server struct {
	strategy bs.BlogStrategy
}

var (
	instance *server
	once     sync.Once
)

// GetInstance returns a singleton instance of the server
func GetInstance(strategy bs.BlogStrategy) *server {
	once.Do(func() {
		instance = &server{
			strategy: strategy,
		}
	})
	return instance
}
