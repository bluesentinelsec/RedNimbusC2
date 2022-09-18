package tasker

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"

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
var lambdaPayloadFile *os.File

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
	taskJSON := convertToJSON(taskObj)
	writeLambdaPayload(taskJSON)

	// write task object to disk as temporary file
	// invoke an AWS CLI command, passing the task file
	// as a payload to set task lambda

	// read environment variable and encrypt
	// task plus arguments if needed
	log.Debug(secretKey)
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

func (taskObj *TaskObject) SetKeyEnv(envVarName string) {
	secretKey = os.Getenv(envVarName)
}

func writeLambdaPayload(payloadJSON []byte) {

	log.Debug("tasker.writeLambdaPayload")

	// create a temporary file
	lambdaPayloadFile, err := ioutil.TempFile("", "nimbusC2")
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
}

func convertToJSON(task *TaskObject) []byte {

	log.Debug("tasker.convertToJSON")
	taskJSON, err := json.Marshal(task)
	if err != nil {
		log.Fatal("tasker.convertToJSON ", err)
	}

	return taskJSON
}
