package sqs

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"

	"queued/driver"
)

type Head struct {
	Queue     string
	PollTimer int64

	client *sqs.SQS
	region string
}

func (h *Head) setClient(region string) {
	h.client = sqs.New(
		session.Must(session.NewSession()),
		aws.NewConfig().WithRegion(region),
	)
}

// Poll retrieve the job from the queue
func (h Head) Poll() ([]driver.Job, error) {
	output, err := h.client.ReceiveMessage(&sqs.ReceiveMessageInput{
		QueueUrl:        &h.Queue,
		WaitTimeSeconds: &h.PollTimer,
	})

	if err != nil {
		// TODO: try again
	}

	jobs := make([]driver.Job, len(output.Messages))
	for _, msg := range output.Messages {
		jobs = append(jobs, Job{msg})
	}

	return jobs, nil
}

// Post put the job back to the queue
func (h Head) Post(job driver.Job) (string, error) {
	fmt.Println("sqs")
	return "", nil
}

// Delete remove the job from the queue for good
func (h Head) Delete(job driver.Job) error {
	return nil
}
