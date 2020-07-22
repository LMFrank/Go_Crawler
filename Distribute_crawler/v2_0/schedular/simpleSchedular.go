package schedular

import "crawler_v2.0/engine"

type SimpleSchedular struct {
	workerChan chan engine.Request
}

func (s *SimpleSchedular) WorkerChan() chan engine.Request {
	return s.workerChan
}

func (s *SimpleSchedular) WorkerReady(requests chan engine.Request) {
}

func (s *SimpleSchedular) Run() {
	s.workerChan = make(chan engine.Request)
}

func (s *SimpleSchedular) Submit(r engine.Request) {
	go func() {
		s.workerChan <- r
	}()
}
