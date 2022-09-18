// package s3wrapper contains functions
// for downloading and uploading files
// to/from S3
package s3wrapper

import (
	"bytes"
	"context"
	"io/ioutil"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	log "github.com/sirupsen/logrus"
)

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

// GetFile downloads the file specified by s3://bucket/key
// and writes it to disk as dstFile
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

	// read file data from S3 response
	buf := new(bytes.Buffer)
	buf.ReadFrom(s3Response.Body)

	// write S3 file to disk
	err = ioutil.WriteFile(dstFile, buf.Bytes(), 0600)
	if err != nil {
		return err
	}

	return nil
}

// RemoveFile deletes from S3 the file specified by
// s3://bucket/key
func RemoveFile(bucket string, key string) error {

	log.Infof("removing file s3://%v/%v", bucket, key)

	// create an S3 client from AWS_PROFILE
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return err
	}
	client := s3.NewFromConfig(cfg)

	_, err = client.DeleteObject(context.TODO(), &s3.DeleteObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})

	if err != nil {
		return err
	}
	return nil
}

// S3FileData contains metadata
// used when invoking ListFiles(bucket)
type S3FileData struct {
	Name         string    // name of the S3 file
	Size         int64     // size of S3 file
	LastModified time.Time // last modified time
	Etag         string    // a unique identifier
}

// ListFiles lists files within an S3 bucket
func ListFiles(bucket string) ([]S3FileData, error) {

	log.Infof("listing files in s3://%v", bucket)

	var filesInBucket []S3FileData

	// create an S3 client from AWS_PROFILE
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return nil, err
	}
	client := s3.NewFromConfig(cfg)

	s3Response, err := client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket: aws.String(bucket),
	})
	if err != nil {
		return nil, err
	}

	for _, eachContent := range s3Response.Contents {
		file := S3FileData{
			Name:         *eachContent.Key,
			Size:         eachContent.Size,
			LastModified: *eachContent.LastModified,
			Etag:         *eachContent.ETag,
		}

		filesInBucket = append(filesInBucket, file)
	}

	return filesInBucket, nil
}

// CreateBucket creates an S3 bucket
func CreateBucket(bucket string) error {

	log.Infof("creating bucket s3://%v", bucket)

	// create an S3 client from AWS_PROFILE
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return err
	}
	client := s3.NewFromConfig(cfg)

	_, err = client.CreateBucket(context.TODO(), &s3.CreateBucketInput{
		Bucket: aws.String(bucket),
	})
	if err != nil {
		return err
	}

	return nil
}

// RemoveBucket deletes the S3 bucket
func RemoveBucket(bucket string) error {

	log.Infof("deleting bucket s3://%v", bucket)

	// create an S3 client from AWS_PROFILE
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return err
	}
	client := s3.NewFromConfig(cfg)
	_, err = client.DeleteBucket(context.TODO(), &s3.DeleteBucketInput{
		Bucket: aws.String(bucket),
	})

	if err != nil {
		return err
	}

	return nil
}
