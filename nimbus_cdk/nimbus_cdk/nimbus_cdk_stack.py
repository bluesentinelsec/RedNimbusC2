import aws_cdk as cdk
from aws_cdk import (
    # Duration,
    Stack,
    aws_lambda as lambda_,
    aws_s3 as s3_,
    aws_iam as iam_
)
from constructs import Construct


class NimbusCdkStack(Stack):

    def __init__(self, scope: Construct, construct_id: str, **kwargs) -> None:
        super().__init__(scope, construct_id, **kwargs)

        # define Lambda function that handles C2
        nimbusC2Handler = lambda_.Function(self,
                                           "nimbusC2-operator-lambda-function",
                                           code=lambda_.Code.from_asset(
                                               "../bootstrap.zip"),
                                           handler="bootstrap",
                                           architecture=lambda_.Architecture.ARM_64,
                                           runtime=lambda_.Runtime.PROVIDED_AL2,
                                           function_name="nimbusC2Handler"
                                           )

        # define S3 bucket used for storing operator data
        bucket_name = f"red-nimbus-c2-{cdk.Aws.REGION}-{cdk.Aws.ACCOUNT_ID}"
        nimbus_c2_bucket = s3_.Bucket(self, "red-nimbus-c2",
                                      access_control=s3_.BucketAccessControl.BUCKET_OWNER_FULL_CONTROL,
                                      block_public_access=s3_.BlockPublicAccess.BLOCK_ALL,
                                      encryption=s3_.BucketEncryption.S3_MANAGED,
                                      removal_policy=cdk.RemovalPolicy.DESTROY,
                                      auto_delete_objects=True,
                                      bucket_name=bucket_name)
        nimbus_c2_bucket.grant_read_write(nimbusC2Handler)

        # define bucket for unit tests
        test_bucket_name = f"red-nimbus-c2-testing-{cdk.Aws.REGION}-{cdk.Aws.ACCOUNT_ID}"
        test_bucket = s3_.Bucket(self, "red-nimbus-c2-testing",
                                 access_control=s3_.BucketAccessControl.BUCKET_OWNER_FULL_CONTROL,
                                 block_public_access=s3_.BlockPublicAccess.BLOCK_ALL,
                                 encryption=s3_.BucketEncryption.S3_MANAGED,
                                 removal_policy=cdk.RemovalPolicy.DESTROY,
                                 auto_delete_objects=True,
                                 bucket_name=test_bucket_name)
        test_bucket.grant_read_write(nimbusC2Handler)
