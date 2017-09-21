package sqs

import "github.com/aws/aws-sdk-go/service/sqs"

// Job that holds sqs message detail
type Job struct {
	*sqs.Message
}

func (j Job) GetJobID() *string {
	return j.MessageId
}

func (j Job) GetBody() *string {
	return j.Body
}
