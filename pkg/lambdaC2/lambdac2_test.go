package lambdac2

import (
	"testing"

	"github.com/bluesentinelsec/rednimbusc2/pkg/tasker"
)

var taskIDforTesting string = "827f9655-1815-41ac-89b5-0941ce32024c"

func TestHandleSetLambdaTask(t *testing.T) {

	task := tasker.NewTask()
	task.TaskID = taskIDforTesting
	task.SetImplantTask("get-pid")
	_, err := HandleSetLambdaTask(task)
	if err != nil {
		t.Fatal(err)
	}
}

func TestHandleGetLambdaTask(t *testing.T) {
	task := tasker.NewTask()
	task.TaskID = taskIDforTesting
	taskData, err := HandleGetLambdaTask(task)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(taskData)
}

func TestHandleRemoveLambdaTask(t *testing.T) {

	task := tasker.NewTask()
	task.TaskID = taskIDforTesting
	_, err := HandleRemoveLambdaTask(task)
	if err != nil {
		t.Fatal(err)
	}

	// should return an error - fail if it doesn't
	_, err = HandleGetLambdaTask(task)
	if err == nil {
		t.Fatal("task not deleted: ", task.TaskID)
	}
}
