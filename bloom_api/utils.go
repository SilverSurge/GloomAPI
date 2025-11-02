package bloomapi

func addChansToPool() {
	chanPool = make(chan chan interface{}, poolSize)
	for i := 0; i < poolSize; i++ {
		chanPool <- make(chan interface{}, 1) // each channel can hold one result
	}
}

func getChan() chan interface{} {
	return <-chanPool
}

func putChan(ch chan interface{}) {
	select {
	case <-ch:
	default:
	}
	chanPool <- ch
}
