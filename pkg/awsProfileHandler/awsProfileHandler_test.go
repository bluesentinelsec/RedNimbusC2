package awsProfileHandler

import "testing"

func TestSetAWSProfile(t *testing.T) {

	err := SetAWSProfile("default")
	if err != nil {
		t.Fatal(err)
	}
}
