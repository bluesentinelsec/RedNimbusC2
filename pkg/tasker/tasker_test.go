package tasker

import "testing"

func TestSetLambdaTask(t *testing.T) {
	taskObj := NewTask()
	taskObj.SetLambdaHandler("HandleSetLambdaTask")
	taskObj.SetKeyEnv("test")
	taskObj.SetSessionID("test")
	taskObj.SetGroupName("test")
	taskObj.SetExecTime("test")
	taskObj.SetImplantTask("get-process")
	taskObj.SetArguments("nothing")
	err := SetLambdaTask(taskObj)
	if err != nil {
		t.Fatal(err)
	}
}
