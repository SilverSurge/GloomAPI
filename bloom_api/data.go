package bloomapi

import "sync"

var (
	workers   = make(map[string]*FilterWorker)
	workersMu sync.RWMutex
)

// for stats
var (
	droppedTasks   uint64
	processedTasks uint64
)
