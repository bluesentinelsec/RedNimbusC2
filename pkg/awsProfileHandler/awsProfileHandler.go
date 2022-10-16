package awsProfileHandler

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sts"
)

// SetAWSProfile specifies the AWS profile to use.
// By default, the AWS SDK checks the AWS_PROFILE environment variable
// to determine which profile to use. If the AWS_PROFILE variable
// is not set in your environment, the SDK uses the credentials
// for the [default] profile. To use one of the alternate profiles,
// set or change the value of the AWS_PROFILE environment variable.
func SetAWSProfile(profile string) error {
	err := os.Setenv("AWS_PROFILE", profile)
	return err
}

// returns the current AWS profile based on environment variable
func GetAWSProfile() string {
	profile := os.Getenv("AWS_PROFILE")
	if len(profile) <= 0 {
		return "default"
	}
	return profile
}

// GetAWSRegion returns the default region
// for the current AWS profile
func GetAWSRegion() (string, error) {
	ctx := context.TODO()
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return "", err
	}
	return cfg.Region, err
}

// GetAWSRegion returns the account ID
// for the current AWS profile
func GetAWSAccountID() (string, error) {
	ctx := context.TODO()
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return "", err
	}
	client := sts.NewFromConfig(cfg)
	identity, err := client.GetCallerIdentity(ctx, &sts.GetCallerIdentityInput{})
	if err != nil {
		return "", err
	}
	return aws.ToString(identity.Account), err
}

func GetNimbusBucketName() (string, error) {
	accountID, err := GetAWSAccountID()
	if err != nil {
		return "", err
	}
	region, err := GetAWSRegion()
	if err != nil {
		return "", err
	}
	bucketName := fmt.Sprintf("red-nimbus-c2-%v-%v", region, accountID)
	return bucketName, err
}

func GetTestBucketName() (string, error) {
	accountID, err := GetAWSAccountID()
	if err != nil {
		return "", err
	}
	region, err := GetAWSRegion()
	if err != nil {
		return "", err
	}
	bucketName := fmt.Sprintf("red-nimbus-c2-testing-%v-%v", region, accountID)
	return bucketName, err
}
