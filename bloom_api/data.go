package bloomapi

import "sync"

// constants
const (
	poolSize = 256
)

// for core logic
var (
	workers   = make(map[string]*FilterWorker)
	workersMu sync.RWMutex
	chanPool  chan chan interface{}
)

// for stats
var (
	droppedTasks   uint64
	processedTasks uint64
)
