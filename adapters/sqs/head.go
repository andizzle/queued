package sqs

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"

	"queued/adapters"
)

// Head listens to SQS
type Head struct {
	Queue       string
	PollTimer   int64
	MaxMessages int64

	client *sqs.SQS
	region string
}

// SetClient set sqs client to head
func (h *Head) SetClient(region string) {
	h.client = sqs.New(
		session.Must(session.NewSession()),
		aws.NewConfig().WithRegion(region),
	)
}

// RefreshToken reauth the session
func (h *Head) RefreshToken() {

}

// Poll retrieve the job from the queue
func (h Head) Poll() ([]adapters.AJob, error) {
	output, err := h.client.ReceiveMessage(&sqs.ReceiveMessageInput{
		QueueUrl:            &h.Queue,
		WaitTimeSeconds:     &h.PollTimer,
		MaxNumberOfMessages: &h.MaxMessages,
	})

	if err != nil {
		// TODO: try again
		log.Fatal(err)
	}

	jobs := make([]adapters.AJob, len(output.Messages))
	for id, msg := range output.Messages {
		job := Job{msg}
		jobs[id] = job
	}

	return jobs, nil
}

// Post put the job back to the queue
func (h Head) Post(job adapters.AJob) (string, error) {
	fmt.Println("sqs")
	return "", nil
}

// Delete remove the job from the queue for good
func (h Head) Delete(job adapters.AJob) error {
	return nil
}
