package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/bluesentinelsec/rednimbusc2/pkg/tasker"
	log "github.com/sirupsen/logrus"
)

func HandleRequest(ctx context.Context, task tasker.TaskObject) (string, error) {

	log.Info("running updateTask: ", task)
	log.Info("get task ID")
	log.Info("look for S3 file: <tasks/task_id>")
	log.Info("overwrite <tasks/task_id> with new task content")
	log.Info("returning success")
	return "success", nil
}

func main() {
	lambda.Start(HandleRequest)
}

/*
S3 design
given a task ID, write task to S3 as <tasks/task_id>
*/
