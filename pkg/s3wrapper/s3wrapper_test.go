package s3wrapper

import (
	"fmt"
	"os"
	"testing"

	"github.com/bluesentinelsec/rednimbusc2/pkg/awsProfileHandler"
	"github.com/google/uuid"
)

var key string = "test_file.txt"
var newBucket string = ""

func getTestBucketName() (string, error) {
	accountID, err := awsProfileHandler.GetAWSAccountID()
	if err != nil {
		return "", err
	}
	region, err := awsProfileHandler.GetAWSRegion()
	if err != nil {
		return "", err
	}
	bucketName := fmt.Sprintf("red-nimbus-c2-testing-%v-%v", region, accountID)
	return bucketName, err
}

func TestPutFile(t *testing.T) {
	bucketName, err := getTestBucketName()
	if err != nil {
		t.Fatal(err)
	}
	err = PutFile("test_files/test_file.txt", bucketName, "test_file.txt")
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetFile(t *testing.T) {

	dstFile := "./test_file.txt"

	bucketName, err := getTestBucketName()
	if err != nil {
		t.Fatal(err)
	}

	err = GetFile(bucketName, "test_file.txt", dstFile)
	if err != nil {
		t.Fatal(err)
	}

	// confirm s3 file was downloaded
	_, err = os.Stat(dstFile)
	if err != nil {
		t.Fatal(err)
	}

	// delete downloaded test file
	defer os.Remove(dstFile)
}

func TestListFiles(t *testing.T) {

	bucketName, err := getTestBucketName()
	if err != nil {
		t.Fatal(err)
	}

	filelist, err := ListFiles(bucketName)
	if err != nil {
		t.Fatal(err)
	}

	passed := false
	for _, eachFile := range filelist {
		if eachFile.Name == key {
			passed = true
		}
	}

	if !passed {
		t.Fatalf("unable to find %v in %v", key, bucketName)
	}
}

func TestRemoveFile(t *testing.T) {
	bucketName, err := getTestBucketName()
	if err != nil {
		t.Fatal(err)
	}

	err = RemoveFile(bucketName, key)
	if err != nil {
		t.Fatal(err)
	}
}

func TestCreateBucket(t *testing.T) {

	randomName, err := uuid.NewUUID()
	if err != nil {
		t.Fatal(err)
	}

	region, err := awsProfileHandler.GetAWSRegion()
	if err != nil {
		t.Fatal(err)
	}

	accountID, err := awsProfileHandler.GetAWSAccountID()
	if err != nil {
		t.Fatal(err)
	}

	newBucket = fmt.Sprintf("%v-%v-%v", randomName.String(), region, accountID)
	err = CreateBucket(newBucket)
	if err != nil {
		t.Fatal(err)
	}
}

func TestRemoveBucket(t *testing.T) {
	err := RemoveBucket(newBucket)
	if err != nil {
		t.Fatal(err)
	}
}
