package logs

import (
	"context"
	"log"
	"net/http"

	"cloud.google.com/go/errorreporting"
)

var errorClient *errorreporting.Client

func Config(ctx context.Context, projectId, serviceName string) error {
	var err error
	errorClient, err = errorreporting.NewClient(ctx, projectId, errorreporting.Config{
		ServiceName: serviceName,
		OnError: func(err error) {
			log.Printf("Could not log error: %v", err)
		},
	})
	return err
}

func Close() {
	if errorClient != nil {
		errorClient.Close()
	}
}

func Error(stack string, err error) {
	if errorClient != nil {
		errorClient.Report(errorreporting.Entry{
			Error: err,
			Stack: []byte(stack),
		})
	}
}

func ErrorHttp(stack string, err error, req *http.Request) {
	if errorClient != nil {
		errorClient.Report(errorreporting.Entry{
			Req:   req,
			Error: err,
			Stack: []byte(stack),
		})
	}
}
