package awsProfileHandler

import "testing"

func TestSetAWSProfile(t *testing.T) {

	err := SetAWSProfile("default")
	if err != nil {
		t.Fatal(err)
	}

	got := GetAWSProfile()
	if got != "default" {
		t.Fatalf("expected AWS_PROFILE to 'default', but received '%v'", got)
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

func TestGetNimbusBucketName(t *testing.T) {
	bucketName, err := GetNimbusBucketName()
	if err != nil {
		t.Fatal(err)
	}

	if len(bucketName) <= 0 {
		t.Fatal("got bucket name without any data")
	}
	t.Log("got bucket name: ", bucketName)
}

func TestGetTestBucketName(t *testing.T) {
	bucketName, err := GetTestBucketName()
	if err != nil {
		t.Fatal(err)
	}

	if len(bucketName) <= 0 {
		t.Fatal("got bucket name without any data")
	}
	t.Log("got bucket name: ", bucketName)
}
