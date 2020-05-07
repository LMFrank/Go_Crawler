package schedular

import "go_crawler/engine"

type QueueSchedular struct {
	requestChan chan engine.Request
	workerChan  chan chan engine.Request
}

func (s *QueueSchedular) Workchan() chan engine.Request {
	return make(chan engine.Request)
}

func (s *QueueSchedular) Submit(r engine.Request) {
	s.requestChan <- r
}

func (s *QueueSchedular) WorkReady(w chan engine.Request) {
	s.workerChan <- w
}

func (s *QueueSchedular) Run() {
	s.workerChan = make(chan chan engine.Request)
	s.requestChan = make(chan engine.Request)
	go func() {
		var requestQ []engine.Request
		var workQ []chan engine.Request
		for {
			var activeRequest engine.Request
			var activeWork chan engine.Request

			if len(requestQ) > 0 && len(workQ) > 0 {
				activeRequest = requestQ[0]
				activeWork = workQ[0]
			}
			select {
			case r := <-s.requestChan:
				requestQ = append(requestQ, r)
			case w := <-s.workerChan:
				workQ = append(workQ, w)
			case activeWork <- activeRequest:
				workQ = workQ[1:]
				requestQ = requestQ[1:]
			}
		}
	}()
}
