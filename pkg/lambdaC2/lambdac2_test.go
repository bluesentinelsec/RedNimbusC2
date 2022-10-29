package lambdac2

import (
	"testing"

	"github.com/bluesentinelsec/rednimbusc2/pkg/tasker"
)

var taskIDforTesting string = "827f9655-1815-41ac-89b5-0941ce32024c"

func createTestTask() (*tasker.TaskObject, error) {
	// create a new task object
	task := tasker.NewTask()

	// hard code task ID for testing
	task.TaskID = taskIDforTesting

	// give command to agent
	task.SetAgentTask("get-pid")

	// pass to task Handler
	_, err := HandleSetLambdaTask(task)
	if err != nil {
		return nil, err
	}

	return task, err
}

func TestHandleSetLambdaTask(t *testing.T) {

	_, err := createTestTask()
	if err != nil {
		t.Fatal(err)
	}
}

func TestHandleGetLambdaTask(t *testing.T) {

	// create test task
	task, err := createTestTask()
	if err != nil {
		t.Fatal(err)
	}

	// confirm we can get the task we just created
	_, err = HandleGetLambdaTask(task)
	if err != nil {
		t.Fatal(err)
	}
}

func TestHandleRemoveLambdaTask(t *testing.T) {

	// create test task
	task, err := createTestTask()
	if err != nil {
		t.Fatal(err)
	}

	// remove the task
	_, err = HandleRemoveLambdaTask(task)
	if err != nil {
		t.Fatal(err)
	}

	// should return an error - fail if it doesn't
	_, err = HandleGetLambdaTask(task)
	if err == nil {
		t.Fatal("task not deleted: ", task.TaskID)
	}
}
