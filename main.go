package main

import (
	"gopkg.in/yaml.v2"

	"queued/adapters/sqs"
)

func boot() {
	// read the yaml config
	// construct the pipe with tubes99
	yaml.Marshal("")

	tube := Tube{head: sqs.Head{}, tail: sqs.Tail{}}
	tube.Watch()
}

func main() {
	boot()
}
