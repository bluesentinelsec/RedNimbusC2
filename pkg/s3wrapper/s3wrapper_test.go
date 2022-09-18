package s3wrapper

import (
	"os"
	"testing"
)

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
