package bloomapi

import "sync"

var (
	workers   = make(map[string]*FilterWorker)
	workersMu sync.RWMutex
)
