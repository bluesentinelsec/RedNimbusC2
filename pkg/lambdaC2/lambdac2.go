// package lambdac2 handels all C2-related
// requests, such as setting and getting tasks
package lambdac2

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/bluesentinelsec/rednimbusc2/pkg/awsProfileHandler"
	"github.com/bluesentinelsec/rednimbusc2/pkg/s3wrapper"
	"github.com/bluesentinelsec/rednimbusc2/pkg/tasker"
	log "github.com/sirupsen/logrus"
)

// LambdaReturnObject is a generic object
// used for returning information from the
// Lambda function back to the operator
// or implant
type LambdaReturnObject struct {
	ReturnType string `json:"returnType"`
	Length     int    `json:"length"`
	Value      []byte `json:"value"`
}

var tasksKey string = "tasks/"
var tmp string = "/tmp/"

func RouteTaskToHandler(taskObj *tasker.TaskObject) (LambdaReturnObject, error) {
	log.Debug("RouteTaskToHandler")
	var err error

	log.Debug("initializing return object")
	returnObject := LambdaReturnObject{}

	log.Debug("getting correct handler for task")
	handler := taskObj.GetLambdaHandler()
	log.Debug("handler is: ", handler)

	switch handler {
	case "HandleSetLambdaTask":
		log.Debug("HandleSetLambdaTask")
		returnObject, err = HandleSetLambdaTask(taskObj)

	case "HandleUpdateLambdaTask":
		log.Debug("HandleUpdateLambdaTask")
		err = HandleUpdateLambdaTask(taskObj)

	case "HandleGetLambdaTaskFromID":
		log.Debug("HandleGetLambdaTask")
		returnObject, err = HandleGetLambdaTask(taskObj)

	case "HandleGetLambdaTaskFromGroup":
		log.Debug("HandleGetLambdaTask")
		returnObject, err = HandleGetLambdaTaskFromGroup(taskObj)

	case "HandleGetLambdaTaskAll":
		log.Debug("HandleGetLambdaTask")
		returnObject, err = HandleGetLambdaTask(taskObj)

	case "HandleRemoveLambdaTask":
		log.Debug("HandleRemoveLambdaTask")
		returnObject, err = HandleRemoveLambdaTask(taskObj)
	// ToDo: add routes access by target implant

	default:
		e := fmt.Sprintf("received invalid Lambda handler in task object: %v\n", taskObj.GetLambdaHandler())
		return returnObject, errors.New(e)
	}

	return returnObject, err
}

// HandleSetLambdaTask writes the provided task object
// to s3://nimbusc2/tasks/{taskID}
func HandleSetLambdaTask(taskObj *tasker.TaskObject) (LambdaReturnObject, error) {

	log.Info("setting new task")
	retObj := LambdaReturnObject{}

	log.Debug("converting task to JSON")
	taskJSON, err := json.Marshal(taskObj)
	if err != nil {
		return retObj, err
	}

	taskFile := tmp + taskObj.TaskID
	log.Debug("writing task file to disk as: ", taskFile)
	err = ioutil.WriteFile(taskFile, taskJSON, 0600)
	if err != nil {
		return retObj, err
	}

	key := tasksKey + taskObj.TaskID
	bucketName, err := awsProfileHandler.GetNimbusBucketName()
	if err != nil {
		return retObj, err
	}
	log.Debugf("writing task to s3://%v/%v", bucketName, key)
	err = s3wrapper.PutFile(taskFile, bucketName, key)
	if err != nil {
		return retObj, err
	}

	log.Info("successfully set task: ", taskObj.TaskID)

	retObj = LambdaReturnObject{
		ReturnType: "set-task",
		Length:     len(taskJSON),
		Value:      taskJSON,
	}

	return retObj, nil
}

func HandleUpdateLambdaTask(taskObj *tasker.TaskObject) error {
	return errors.New("sorry, HandleUpdateLambdaTask is not implemented")
}

func HandleGetLambdaTask(taskObj *tasker.TaskObject) (LambdaReturnObject, error) {

	log.Info("getting task: ", taskObj.TaskID)
	retObj := LambdaReturnObject{}

	key := tasksKey + taskObj.TaskID
	log.Debug("set S3 key: ", key)

	outFile := tmp + taskObj.TaskID
	log.Debug("set out file: ", outFile)

	bucketName, err := awsProfileHandler.GetNimbusBucketName()
	log.Debug("got S3 bucket name: ", bucketName)
	if err != nil {
		return retObj, err
	}
	log.Debugf("downloading file s3://%v/%v", bucketName, key)
	err = s3wrapper.GetFile(bucketName, key, outFile)
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

func HandleGetLambdaTaskFromGroup(taskObj *tasker.TaskObject) (LambdaReturnObject, error) {
	retObj := LambdaReturnObject{}
	err := errors.New("sorry, this function is not implemented")
	return retObj, err
}

func HandleGetLambdaTaskAll(taskObj *tasker.TaskObject) (LambdaReturnObject, error) {
	retObj := LambdaReturnObject{}
	err := errors.New("sorry, this function is not implemented")
	return retObj, err
}

func HandleRemoveLambdaTask(taskObj *tasker.TaskObject) (LambdaReturnObject, error) {

	log.Info("deleting task: ", taskObj.TaskID)
	retObj := LambdaReturnObject{}

	bucketName, err := awsProfileHandler.GetNimbusBucketName()
	if err != nil {
		return retObj, err
	}
	key := tasksKey + taskObj.TaskID
	err = s3wrapper.RemoveFile(bucketName, key)
	if err != nil {
		return retObj, err
	}

	s := fmt.Sprintf("successfully deleted task: %v", taskObj.TaskID)
	log.Info(s)

	retObj.ReturnType = "remove-task"
	retObj.Length = len(s)
	retObj.Value = []byte(s)

	return retObj, nil
}
