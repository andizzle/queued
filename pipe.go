package main

import "queued/driver"

// Head interacts with the queue
type Head interface {
	Poll() ([]driver.Job, error)
	Post(driver.Job) (string, error)
	Delete(driver.Job) error
}

// Tail interacts with the worker
type Tail interface {
	Post(driver.Job) error
}

// Tube has a Head and a Tail
type Tube struct {
	head Head
	tail Tail
	// TODO: add in logger
}

// Watch the activities on the Head
// If there's a message, send it to Tail
// Watch is a long lasting job
func (t *Tube) Watch() {

	for {
		jobs, err := t.head.Poll()
		if err != nil {
			// TODO: recover
		}

		for _, job := range jobs {
			error := t.tail.Post(job)
			if error != nil {
				// push the job back to the queue
				t.head.Post(job)
			} else {
				// job is done, delete from the queue
				t.head.Delete(job)
			}
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
