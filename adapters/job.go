package adapters

// AJob is a job from the queue
type AJob interface {
	GetJobID() *string
	GetBody() *string
}
