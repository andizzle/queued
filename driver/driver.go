package driver

type Job interface {
	GetJobID() *string
	GetBody() *string
}
