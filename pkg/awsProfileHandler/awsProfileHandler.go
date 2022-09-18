package awsProfileHandler

import "os"

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
