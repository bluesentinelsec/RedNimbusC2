package s3wrapper

import (
	"os"
	"testing"

	"github.com/google/uuid"
)

var bucket string = "nimbusc2-testing"
var key string = "test_file.txt"
var newBucket string = ""

func TestPutFile(t *testing.T) {

	err := PutFile("test_files/test_file.txt", "nimbusc2-testing", "test_file.txt")
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetFile(t *testing.T) {

	dstFile := "./test_file.txt"

	err := GetFile("nimbusc2-testing", "test_file.txt", dstFile)
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
	filelist, err := ListFiles(bucket)
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
		t.Fatalf("unable to find %v in %v", key, bucket)
	}
}

func TestRemoveFile(t *testing.T) {
	err := RemoveFile(bucket, key)
	if err != nil {
		t.Fatal(err)
	}
}

func TestCreateBucket(t *testing.T) {

	randomName, err := uuid.NewUUID()
	if err != nil {
		t.Fatal(err)
	}

	newBucket = randomName.String()
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
