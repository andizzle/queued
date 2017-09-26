package sqs

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"

	"github.com/aws/aws-sdk-go/service/sqs"
)

// Job that holds sqs message detail
type Job struct {
	*sqs.Message
}

func (j Job) GetAttribute(s string) int64 {
	attr, _ := strconv.ParseInt(*j.Attributes[s], 10, 0)

	return attr
}

func (j Job) GetJobID() *string {
	return j.MessageId
}

func (j Job) GetBody() *string {
	return j.Body
}

func (j Job) GetReceiptHandle() *string {
	return j.ReceiptHandle
}

func (j Job) CheckSum() bool {
	sum := md5.Sum([]byte(*j.Body))

	return hex.EncodeToString(sum[:]) == *j.MD5OfBody
}
