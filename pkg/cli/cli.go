package cli

import (
	"fmt"
	"os"

	"github.com/akamensky/argparse"
	"github.com/bluesentinelsec/rednimbusc2/pkg/awsProfileHandler"
	"github.com/bluesentinelsec/rednimbusc2/pkg/tasker"
	log "github.com/sirupsen/logrus"
)

// CLI variables
// give these file-scope so
// we don't need to pass 8+ args
// to functions
var parser *argparse.Parser
var sessionIDFlag *string
var groupFlag *string
var taskIDFlag *string
var cmdFlag *string
var argsFlag *string
var execTimeFlag *string
var keyEnvFlag *string
var awsProfileFlag *string

func InvokeCLI(args []string) {

	parser = argparse.NewParser(args[0], "A client for interacting with Red Nimbus C2 services")

	// top level commands - red operators will invoke
	// these commands to control implant behavior
	setTaskCmd := parser.NewCommand("set-task", "Issue a task to an agent")
	getTaskCmd := parser.NewCommand("get-task", "Read a pending agent task")
	updateTaskCmd := parser.NewCommand("update-task", "Update a pending agent task")
	removeTaskCmd := parser.NewCommand("remove-task", "Delete a pending agent task")
	getSessionCmd := parser.NewCommand("get-session", "Get detailed info about the session")
	listSessionsCmd := parser.NewCommand("list-sessions", "Get summarized info about all existing sessions")
	removeSessionCmd := parser.NewCommand("remove-session", "Terminate session")
	terminateSessionsCmd := parser.NewCommand("terminate-sessions", "Terminate all existing sessions")

	// arguments that modify command behavior
	sessionIDFlag = parser.String("s", "session-id", &argparse.Options{Required: false, Help: "The agent session you wish to task"})
	groupFlag = parser.String("g", "session-group", &argparse.Options{Required: false, Help: "The agent group you wish to task"})
	taskIDFlag = parser.String("i", "task-id", &argparse.Options{Required: false, Help: "The task ID you wish to read/update/delete"})
	cmdFlag = parser.String("c", "cmd", &argparse.Options{Required: false, Help: "The command the agent(s) will execute"})
	argsFlag = parser.String("a", "args", &argparse.Options{Required: false, Help: "Command arguments"})
	execTimeFlag = parser.String("t", "time", &argparse.Options{Required: false, Help: "The Unix epoch time which agents should execute the task (based on the agent's time zone)"})
	keyEnvFlag = parser.String("k", "key-env", &argparse.Options{Required: false, Help: "Use the provided environment variable as a password to encrypt sensitive task details"})
	awsProfileFlag = parser.String("p", "aws-profile", &argparse.Options{Required: false, Help: "The AWS profile to utilize for interacting with AWS services"})
	enableVerboseCmd := parser.Flag("v", "verbose", &argparse.Options{Required: false, Help: "Enable verbose console logging"})

	// Parse input
	err := parser.Parse(os.Args)
	if err != nil {
		// In case of error print error and print usage
		// This can also be done by passing -h or --help flags
		fmt.Print(parser.Usage(err))
	}

	// setup console logging
	if *enableVerboseCmd {
		log.SetLevel(log.DebugLevel)
	}

	// set AWS profile - uses default if left blank
	if *awsProfileFlag != "" {
		awsProfileHandler.SetAWSProfile(*awsProfileFlag)
	}

	if setTaskCmd.Happened() {
		log.Debug("invoke set-task")
		// ToDo - confirm task is a valid session
		invokeSetLambdaTask()
		return
	}

	if getTaskCmd.Happened() {
		log.Debug("invoke get-task")
		invokeGetLambdaTask()
		return
	}

	if updateTaskCmd.Happened() {
		log.Debug("invoke update-task")
		invokeUpdateLambdaTask()
		return
	}

	if removeTaskCmd.Happened() {
		log.Debug("invoke remove-task")
		invokeRemoveLambdaTask()
		return
	}

	if getSessionCmd.Happened() {
		log.Debug("invoke remove-task")
		invokeRemoveLambdaTask()
		return
	}

	if listSessionsCmd.Happened() {
		log.Fatal("sorry, not implemented")
	}

	if removeSessionCmd.Happened() {
		log.Fatal("sorry, not implemented")
	}

	if terminateSessionsCmd.Happened() {
		log.Fatal("sorry, not implemented")
	}

}

func configureTaskObject(lambdaHandler string) *tasker.TaskObject {
	taskObj := tasker.NewTask()
	taskObj.SetLambdaHandler(lambdaHandler)
	taskObj.SetKeyEnv(*keyEnvFlag)
	taskObj.SetSessionID(*sessionIDFlag)
	taskObj.SetGroupName(*groupFlag)
	taskObj.SetExecTime(*execTimeFlag)
	taskObj.SetImplantTask(*cmdFlag)
	taskObj.SetArguments(*argsFlag)
	return taskObj
}

// invokeSetLambdaTask creates a task
// object based on the CLI arguments
// and then invokes the
// Set Task Lambda function.
func invokeSetLambdaTask() {

	taskObj := configureTaskObject("HandleSetLambdaTask")
	tasker.SetLambdaTask(taskObj)
}

func invokeUpdateLambdaTask() {

	taskObj := configureTaskObject("invokeUpdateLambdaTask")
	tasker.UpdateLambdaTask(taskObj)
}

func invokeRemoveLambdaTask() {

	taskObj := configureTaskObject("HandleGetLambdaTask")
	if *taskIDFlag != "" {
		tasker.RemoveLambdaTaskWithID(taskObj)

	} else if *groupFlag != "" {
		tasker.RemoveLambdaTaskWithGroup(taskObj)

	} else {
		tasker.RemoveLambdaTaskAll(taskObj)
	}
}

func invokeGetLambdaTask() {

	log.Debug("invokeGetLambdaTask")

	if *taskIDFlag != "" {
		taskObj := configureTaskObject("HandleGetLambdaTaskFromID")
		// set the ID to the task-id we want to read
		taskObj.SetTaskID(*taskIDFlag)
		tasker.GetLambdaTaskFromID(taskObj)

	} else if *groupFlag != "" {
		taskObj := configureTaskObject("HandleGetLambdaTaskFromGroup")
		taskObj.SetGroupName(*groupFlag)
		tasker.GetLambdaTaskFromGroup(taskObj)

	} else {
		taskObj := configureTaskObject("HandleGetLambdaTaskAll")
		taskObj.SetTaskID("all")
		tasker.GetLambdaTaskAll(taskObj)
	}
}
