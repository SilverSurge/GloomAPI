package bloomapi

import (
	"net/http"
	"strings"

	"github.com/SilverSurge/Gloom/bloom"
	"github.com/gin-gonic/gin"
)

func pingHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"ping": "pong",
	})
}

func createFilterHandler(c *gin.Context) {
	var req CreateFilterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	if len(strings.TrimSpace(req.ID)) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	workersMu.Lock()
	w, exists := workers[req.ID]
	workersMu.Unlock()

	if exists {
		c.JSON(http.StatusFound, gin.H{
			"message": "Bloom filter already exists",
			"id":      w.ID,
			"n_bits":  w.Filter.State.NBits,
			"n_hash":  w.Filter.State.NHash,
		})
	} else {

		nBits, nHash := bloom.GetOptimalParameters(req.NAdd, req.FalsePositiveProb)
		newW := &FilterWorker{
			ID:     req.ID,
			Filter: bloom.NewBloomDefault(req.ID, nBits, nHash), // assume NewDefault() returns a usable Bloom
			Queue:  make(chan FilterTask, 128),
		}

		workersMu.Lock()
		workers[req.ID] = newW
		workersMu.Unlock()
		c.JSON(http.StatusCreated, gin.H{
			"message": "Bloom filter created successfully",
			"id":      req.ID,
			"n_bits":  nBits,
			"n_hash":  nHash,
		})
	}
}

func listFiltersHandler(c *gin.Context) {
	workersMu.RLock()
	defer workersMu.RUnlock()

	if len(workers) == 0 {
		c.JSON(http.StatusOK, gin.H{"filters": []gin.H{}})
		return
	}

	list := make([]gin.H, 0, len(workers))
	for id, w := range workers {
		list = append(list, gin.H{
			"id":       id,
			"n_bits":   w.Filter.State.NBits,
			"n_hashes": w.Filter.State.NHash,
			"queue":    len(w.Queue),
		})
	}

	c.JSON(http.StatusOK, gin.H{"filters": list})
}
