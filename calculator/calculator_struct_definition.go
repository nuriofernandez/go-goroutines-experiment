package calculator

import (
	"sync"
)

type Calculator struct {
	cache map[int]int
	lock  sync.RWMutex
}

func New() *Calculator {
	return &Calculator{
		cache: make(map[int]int),
	}
}
