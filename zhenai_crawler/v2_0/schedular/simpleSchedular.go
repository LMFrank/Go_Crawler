package schedular

import "cralwer_v2.0/engine"

type SimpleSchedular struct {
	workerChan chan engine.Request
}

func (s *SimpleSchedular) ConfigureMasterWorkerChan(c chan engine.Request) {
	s.workerChan = c
}

func (s *SimpleSchedular) Submit(r engine.Request) {
	go func() {
		s.workerChan <- r
	}()
}
