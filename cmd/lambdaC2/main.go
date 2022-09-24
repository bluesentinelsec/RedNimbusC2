package main

import (
	"context"
	"fmt"
	"reflect"

	"github.com/aws/aws-lambda-go/lambda"
	lambdac2 "github.com/bluesentinelsec/rednimbusc2/pkg/lambdaC2"
	"github.com/bluesentinelsec/rednimbusc2/pkg/tasker"
	log "github.com/sirupsen/logrus"
)

func HandleRequest(ctx context.Context, taskObj tasker.TaskObject) (lambdac2.LambdaReturnObject, error) {

	log.Info("received nimbusC2 request")

	// display the task object so we can see
	// if we submitted a well formed task
	displayTask(taskObj)

	// pass task object to the task handler for processing
	returnObj, err := lambdac2.RouteTaskToHandler(&taskObj)
	if err != nil {
		log.Fatal(err)
	}
	log.Info("request succeeded")
	log.Info(returnObj, err)
	return returnObj, err
}

func main() {
	lambda.Start(HandleRequest)
}

func displayTask(taskObj tasker.TaskObject) {
	taskReflect := reflect.ValueOf(taskObj)
	typeOfS := taskReflect.Type()
	for i := 0; i < taskReflect.Len(); i++ {
		fmt.Printf("%v: %v\n", typeOfS.Field(i).Name, taskReflect.Field(i).Interface())
	}
}
