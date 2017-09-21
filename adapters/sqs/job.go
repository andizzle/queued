package sqs

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/aws/aws-sdk-go/service/sqs"
)

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

func (j Job) CheckSum() bool {
	sum := md5.Sum([]byte(*j.Body))

	return hex.EncodeToString(sum[:]) == *j.MD5OfBody
}
