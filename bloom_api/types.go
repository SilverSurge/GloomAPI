package bloomapi

import "github.com/SilverSurge/Gloom/bloom"

type FilterAction int

const (
	CreateFilter FilterAction = iota
	ListFilters
	AddElements
	CheckElements
	UnionFilters
	SaveFilters
	LoadFilters
	DeleteFilters
)

type FilterTask struct {
	Action FilterAction
	Args   interface{}
	Resp   chan interface{}
}

type FilterWorker struct {
	ID     string
	Filter *bloom.Bloom
	Queue  chan FilterTask
}

/*
FilterTask Action should be one of:
*/

type CreateFilterRequest struct {
	ID                string  `json:"id"`
	NAdd              uint64  `json:"n_add"`
	FalsePositiveProb float64 `json:"false_positive_prob"`
}
