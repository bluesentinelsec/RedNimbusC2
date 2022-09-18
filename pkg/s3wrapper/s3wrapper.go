package s3wrapper

import (
	"bytes"
	"context"
	"io/ioutil"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	log "github.com/sirupsen/logrus"
)

/*
GetFile

ListBucket

ListBucketRecursive
*/

// PutFile uploads srcFile to s3 as s3://bucket/key
func PutFile(srcFile string, bucket string, key string) error {

	log.Infof("uploading %v to s3://%v/%v", srcFile, bucket, key)

	// get a pointer to src file
	fileP, err := os.Open(srcFile)
	if err != nil {
		return err
	}

	// create an S3 client from AWS_PROFILE
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return err
	}
	client := s3.NewFromConfig(cfg)

	// upload the file
	_, err = client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   fileP,
	})
	if err != nil {
		return err
	}

	return nil
}

func GetFile(bucket string, key string, dstFile string) error {

	log.Infof("downloading s3://%v/%v to %v", bucket, key, dstFile)

	// create an S3 client from AWS_PROFILE
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return err
	}
	client := s3.NewFromConfig(cfg)

	// download the file
	s3Response, err := client.GetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return err
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(s3Response.Body)

	err = ioutil.WriteFile(dstFile, buf.Bytes(), 0600)
	if err != nil {
		return err
	}

	return nil
}
