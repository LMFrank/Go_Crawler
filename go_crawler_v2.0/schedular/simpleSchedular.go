package schedular

import "go_crawler/engine"

type SimpleSchedular struct {
	workerChan chan engine.Request
}

func (s *SimpleSchedular) Run() {
	s.workerChan = make(chan engine.Request)
}

func (s *SimpleSchedular) WorkReady(requests chan engine.Request) {
	return
}

func (s *SimpleSchedular) Workchan() chan engine.Request {
	return s.workerChan
}

func (s *SimpleSchedular) Submit(r engine.Request) {
	go func() {
		s.workerChan <- r
	}()
}
