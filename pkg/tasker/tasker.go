package tasker

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/bluesentinelsec/rednimbusc2/pkg/shellexec"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

// TaskObject is used to give
// target implants instructions,
// such as downloading a file,
// running a command, or terminating the
// C2 session.
type TaskObject struct {
	// taskID uniquely identifies the task.
	// This is generated automatically for
	// every new task.
	TaskID string `json:"taskID"`

	// sessionID specifies the exact implant session
	// that should receive and execute the task.
	SessionID string `json:"sessionID"`

	// LambdaHandler determines which Lambda C2 function/
	// processes the task request
	LambdaHandler string `json:"lambdaHandler"`

	// groupName is a label for implant sessions;
	// sessions with the matching group name will
	// receive and execute the task.
	GroupName string `json:"groupName"`

	// execTime specifies the time at which
	// an implant session should execute the task.
	// The time should be in ISO 8601 format, for
	// example: 2018-12-10T13:45:00.000Z.
	ExecTime string `json:"execTime"`

	// Task specifies a supported task to run.
	// For a list of supported tasks, see TBD.
	Task string `json:"task"`

	// Arguments holds settings needed for
	// different taskings. For example,
	// the 'get-file' task requires an
	// argument indicating which file to
	// retrieve.
	Arguments []string `json:"arguments"`
}

var secretKey string

//var lambdaPayloadFile *os.File

// NewTask generates a new task object
// with a unique task ID. This function
// should be called before calling SetLambdaTask()
// or UpdateLambdaTask().
func NewTask() *TaskObject {
	taskObj := TaskObject{}
	taskObj.createTaskID() // generate a UUID to uniquely identify the task
	return &taskObj
}

func SetLambdaTask(taskObj *TaskObject) {

	log.Debug("tasker.SetLambdaTask")

	// read environment variable and encrypt
	// task plus arguments if needed
	log.Debug(secretKey)

	// convert the task object to JSON for use by Lambda
	taskJSON := convertToJSON(taskObj)

	// write task JSON to disk as temporary file
	fileWritten := writeLambdaPayload(taskJSON)

	// cleanup task file when finished
	defer os.Remove(fileWritten)

	// invoke an AWS CLI command, passing the task file
	// as a payload to set task lambda
	outFile := "/tmp/outfile"
	cmd := fmt.Sprintf("aws lambda invoke --function-name nimbusC2Handler --invocation-type Event --payload file://%v --cli-binary-format raw-in-base64-out %v", fileWritten, outFile)
	err := shellexec.ExecShellCmd(cmd)
	if err != nil {
		log.Fatal(err)
	}
}

func UpdateLambdaTask(taskObj *TaskObject) {
	// write task to disk as temporary file
	// invoke via AWS CLI command
	// send updated task to update lambda

	// read environment variable and encrypt
	// task plus arguments if needed
	log.Debug(secretKey)
}

func GetLambdaTaskFromID(id string)       {}
func GetLambdaTaskFromGroup(group string) {}
func GetLambdaTaskAll()                   {}

func RemoveLambdaTaskWithID(id string)       {}
func RemoveLambdaTaskWithGroup(group string) {}
func RemoveLambdaTaskAll()                   {}

//-----------------------------------
// TaskObject setters and getters.
// Use these functions to get/set
// the TaskObject member variables.
//-----------------------------------

func (taskObj *TaskObject) createTaskID() {
	taskObj.TaskID = uuid.NewString()
}

func (taskObj *TaskObject) getTaskID() string {
	return taskObj.TaskID
}

func (taskObj *TaskObject) SetSessionID(id string) {
	taskObj.SessionID = id
}

func (taskObj *TaskObject) GetSessionID() string {
	return taskObj.SessionID
}

func (taskObj *TaskObject) SetGroupName(group string) {
	taskObj.GroupName = group
}
func (taskObj *TaskObject) GetGroupName() string {
	return taskObj.GroupName
}

func (taskObj *TaskObject) SetExecTime(timeISO8601 string) {
	taskObj.ExecTime = timeISO8601
}

func (taskObj *TaskObject) GetExecTime() string {
	return taskObj.ExecTime
}

func (taskObj *TaskObject) SetImplantTask(task string) {
	taskObj.Task = task
}

func (taskObj *TaskObject) GetImplantTask() string {
	return taskObj.Task
}

func (taskObj *TaskObject) SetArguments(arguments string) {
	argList := strings.Split(arguments, ",")
	taskObj.Arguments = append(taskObj.Arguments, argList...)
}

func (taskObj *TaskObject) GetArguments() []string {
	return taskObj.Arguments
}

func (taskObj *TaskObject) SetLambdaHandler(handler string) {
	taskObj.LambdaHandler = handler
}

func (taskObj *TaskObject) GetLambdaHandler() string {
	return taskObj.LambdaHandler
}

func (taskObj *TaskObject) SetKeyEnv(envVarName string) {
	secretKey = os.Getenv(envVarName)
}

func writeLambdaPayload(payloadJSON []byte) string {

	log.Debug("tasker.writeLambdaPayload")

	// create a temporary file
	localTmp := os.TempDir()
	lambdaPayloadFile, err := ioutil.TempFile(localTmp, "nimbusC2")
	if err != nil {
		log.Fatal("tasker.writeLambdaPayload ", err)
	}

	log.Debug("writing Lambda payload file to: ", lambdaPayloadFile.Name())

	// write lambda payload JSON to temp file
	_, err = lambdaPayloadFile.Write(payloadJSON)
	if err != nil {
		log.Fatal("tasker.writeLambdaPayload ", err)
	}

	// caller will delete the temp file when finished
	return lambdaPayloadFile.Name()
}

func convertToJSON(task *TaskObject) []byte {

	log.Debug("tasker.convertToJSON")
	taskJSON, err := json.Marshal(task)
	if err != nil {
		log.Fatal("tasker.convertToJSON ", err)
	}

	return taskJSON
}
