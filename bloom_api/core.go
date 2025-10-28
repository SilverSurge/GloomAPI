package bloomapi

import (
	"fmt"
	"sync/atomic"
)

func (fw *FilterWorker) run() {
	for task := range fw.Queue {
		atomic.AddUint64(&processedTasks, 1)
		switch task.Action {
		case AddElements:

			req, ok := task.Args.(AddElementsRequest)
			if !ok {
				task.Resp <- fmt.Errorf("invalid args for AddElements")
				close(task.Resp)
				continue
			}

			for _, e := range req.Elements {
				fw.Filter.Add(e)
			}

			response := map[string]interface{}{
				"message": fmt.Sprintf("%d elements added to filter '%s'", len(req.Elements), fw.ID),
			}
			task.Resp <- response
			close(task.Resp)

		case CheckElements:
			req, ok := task.Args.(CheckElementsRequest)
			if !ok {
				task.Resp <- fmt.Errorf("invalid args for CheckElements")
				close(task.Resp)
				continue
			}

			results := make([]bool, len(req.Elements))
			for idx, e := range req.Elements {
				results[idx] = fw.Filter.Check(e)
			}

			response := map[string]interface{}{
				"results": results,
			}
			task.Resp <- response
			close(task.Resp)
		}
	}
}
