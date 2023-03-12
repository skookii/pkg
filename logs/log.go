package logs

import (
	"context"
	"log"

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

func Error(err error) {
	if errorClient != nil {
		errorClient.Report(errorreporting.Entry{
			Error: err,
		})
	}
}
