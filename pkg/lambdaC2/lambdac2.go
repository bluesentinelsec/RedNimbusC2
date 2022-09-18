// package lambdac2 handels all C2-related
// requests, such as setting and getting tasks
package lambdac2

import (
	"encoding/json"
	"io/ioutil"

	"github.com/bluesentinelsec/rednimbusc2/pkg/s3wrapper"
	"github.com/bluesentinelsec/rednimbusc2/pkg/tasker"
	log "github.com/sirupsen/logrus"
)

var bucket string = "nimbusc2"
var tasksKey string = "tasks/"
var tmp string = "/tmp/"

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

func HandleUpdateLambdaTask(taskObj *tasker.TaskObject) {
	log.Fatal("implement me!")
}

func HandleGetLambdaTask(taskObj *tasker.TaskObject) ([]byte, error) {

	log.Info("getting task: ", taskObj.TaskID)

	key := tasksKey + taskObj.TaskID
	outFile := tmp + taskObj.TaskID
	log.Debugf("downloading file s3://%v/%v", bucket, key)
	err := s3wrapper.GetFile(bucket, key, outFile)
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
