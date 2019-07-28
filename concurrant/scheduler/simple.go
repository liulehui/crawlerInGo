package scheduler

import "github.com/liulehui/crawler/concurrant/engine"

type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (s *SimpleScheduler) WorkerChan() chan engine.Request {
	return s.workerChan
}

func (s *SimpleScheduler) WorkerReady(chan engine.Request) {
}

func (s *SimpleScheduler) Run() {
	s.workerChan = make(chan engine.Request)
}

func (s *SimpleScheduler) Submit(
	r engine.Request) {
		// goroutine chan 防止 循环等待
	go func() { s.workerChan <- r }()
}
