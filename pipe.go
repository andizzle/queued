package main

import (
	"queued/adapters"
)

// Head interacts with the queue
type Head interface {
	Poll() ([]adapters.AJob, error)
	Post(adapters.AJob) (string, error)
	Delete(adapters.AJob) error
}

// Tail interacts with the worker
type Tail interface {
	Post(adapters.AJob) error
}

// Tube has a Head and a Tail
type Tube struct {
	Head Head
	Tail Tail
	// TODO: add in logger
}

// Watch the activities on the Head
// If there's a message, send it to Tail
// Watch is a long lasting job
func (t *Tube) Watch() {

	for {
		jobs, err := t.Head.Poll()
		if err != nil {
			// TODO: recover
		}

		for _, job := range jobs {
			go func(job adapters.AJob) {
				error := t.Tail.Post(job)
				if error != nil {
					// push the job back to the queue
					t.Head.Post(job)
				} else {
					// job is done, delete from the queue
					t.Head.Delete(job)
				}
			}(job)
		}
	}

}

// Pipe is a collection of Tubes
type Pipe struct {
	tubes map[string]Tube
}

// Start the pipe, which contains multiple tubes
func (p *Pipe) Start() {
	for _, tube := range p.tubes {
		// start tube watch in a channel
		go tube.Watch()
	}
}
