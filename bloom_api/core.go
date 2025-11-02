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
			fw.addElements(&task)

		case CheckElements:
			fw.checkElements(&task)
		}
	}
}

func (fw *FilterWorker) addElements(task *FilterTask) {
	req, ok := task.Args.(AddElementsRequest)
	if !ok {
		task.Resp <- fmt.Errorf("invalid args for AddElements")
		return
	}

	for _, e := range req.Elements {
		fw.Filter.Add(e)
	}

	response := map[string]interface{}{
		"message": fmt.Sprintf("%d elements added to filter '%s'", len(req.Elements), fw.ID),
	}
	task.Resp <- response
}

func (fw *FilterWorker) checkElements(task *FilterTask) {
	req, ok := task.Args.(CheckElementsRequest)
	if !ok {
		task.Resp <- fmt.Errorf("invalid args for CheckElements")
		return
	}

	results := make([]bool, len(req.Elements))
	for idx, e := range req.Elements {
		results[idx] = fw.Filter.Check(e)
	}

	response := map[string]interface{}{
		"results": results,
	}
	task.Resp <- response
}
