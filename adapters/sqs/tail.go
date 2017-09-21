package sqs

import (
	"log"
	"strings"

	"net/http"

	"queued/driver"
)

type Tail struct {
	URL         string
	Port        int
	Path        string
	MaxTry      int
	MimeType    string
	KeepAlive   bool
	WaitTime    int
	MaxWaitTime int
	ContentType string

	client *http.Client
}

func (t *Tail) setClient() {
	t.client = &http.Client{}
}

func (t Tail) Post(job driver.Job) error {
	rep, err := t.client.Post(t.URL, t.ContentType, strings.NewReader(*job.GetBody()))

	if err != nil {
		return err
	}

	if rep.StatusCode > 400 {
		// error
		content := []byte{}
		rep.Body.Read(content)

		log.Fatal(content)
	}

	return nil
}
