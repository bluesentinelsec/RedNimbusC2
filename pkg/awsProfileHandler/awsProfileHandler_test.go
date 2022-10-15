package awsProfileHandler

import "testing"

func TestSetAWSProfile(t *testing.T) {

	err := SetAWSProfile("default")
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetAWSRegion(t *testing.T) {
	region, err := GetAWSRegion()
	if err != nil {
		t.Fatal(err)
	}
	if len(region) <= 0 {
		t.Fatal("got region without any data")
	}
	t.Log("got region: ", region)
}

func TestGetAWSAccountID(t *testing.T) {
	accountID, err := GetAWSAccountID()
	if err != nil {
		t.Fatal(err)
	}
	if len(accountID) <= 0 {
		t.Fatal("got account ID without any data")
	}
	t.Log("got account ID: ", accountID)
}
