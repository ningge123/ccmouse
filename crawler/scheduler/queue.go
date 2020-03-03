package scheduler

import "ccmouse/crawler/engine"

type QueueScheduler struct {
	requestChan chan engine.Request
	workerChan chan chan engine.Request
}

func (queueScheduler *QueueScheduler) Submit(request engine.Request)  {
	queueScheduler.requestChan <- request
}

func (queueScheduler *QueueScheduler) WorkerChan() chan engine.Request {
	return make(chan engine.Request)
}

func (queueScheduler *QueueScheduler) WorkerReady(work chan engine.Request)  {
	queueScheduler.workerChan <- work
}

func (queueScheduler *QueueScheduler) Run() {
	queueScheduler.workerChan = make(chan chan engine.Request)
	queueScheduler.requestChan = make(chan engine.Request)
	go func() {
		var requestQueue []engine.Request
		var workerQueue []chan engine.Request
		for {
			var activeRequest engine.Request
			var activeWorker chan engine.Request

			if len(requestQueue) > 0 && len(workerQueue) > 0 {
				activeRequest = requestQueue[0]
				activeWorker  = workerQueue[0]
			}
			select {
				case requestChan := <- queueScheduler.requestChan:
					requestQueue = append(requestQueue, requestChan)
				case workerChan := <- queueScheduler.workerChan:
					workerQueue = append(workerQueue, workerChan)
				case activeWorker <- activeRequest:
					requestQueue = requestQueue[1:]
					workerQueue  = workerQueue[1:]
			}
		}
	}()
}