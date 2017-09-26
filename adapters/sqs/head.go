package sqs

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"

	"queued/adapters"
)

const AZURL = "http://169.254.169.254/latest/meta-data/placement/availability-zone"

// Head listens to SQS
type Head struct {
	Queue       string
	PollTimer   int64
	MaxMessages int64
	MaxTry      int64

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
	All := "All"

	output, err := h.client.ReceiveMessage(&sqs.ReceiveMessageInput{
		AttributeNames:      []*string{&All},
		QueueUrl:            &h.Queue,
		WaitTimeSeconds:     &h.PollTimer,
		MaxNumberOfMessages: &h.MaxMessages,
	})

	if err != nil {
		// TODO: try again
		log.Fatal(err)
	}

	jobs := []adapters.AJob{}
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
func (h Head) Delete(job Job) error {
	_, err := h.client.DeleteMessage(&sqs.DeleteMessageInput{
		QueueUrl:      &h.Queue,
		ReceiptHandle: job.GetReceiptHandle(),
	})

	if err != nil {
		return err
	}

	return nil
}
