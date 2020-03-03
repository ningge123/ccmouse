package scheduler

import "ccmouse/crawler/engine"

type SimpleScheduler struct {
	WorkChan chan engine.Request
}

func (simpleScheduler *SimpleScheduler) Submit(request engine.Request)  {
	go func() {
		simpleScheduler.WorkChan <- request
	}()
}

func (simpleScheduler *SimpleScheduler) WorkerReady(work chan engine.Request) {
	//
}

func (simpleScheduler *SimpleScheduler) WorkerChan() chan engine.Request {
	return simpleScheduler.WorkChan
}

func (simpleScheduler *SimpleScheduler) Run() {
	simpleScheduler.WorkChan = make(chan engine.Request)
}