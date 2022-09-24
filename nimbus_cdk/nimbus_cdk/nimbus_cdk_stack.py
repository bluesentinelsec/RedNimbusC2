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

        # define S3 bucket used for storing tasks and output
        # give bucket owner full control
        # block public access
        # let AWS encrypt data for us
        s3Bucket = s3_.Bucket(self, "nimbusC2-s3-bucket",
                              access_control=s3_.BucketAccessControl.BUCKET_OWNER_FULL_CONTROL,
                              block_public_access=s3_.BlockPublicAccess.BLOCK_ALL,
                              encryption=s3_.BucketEncryption.S3_MANAGED)

        # define Lambda function that handles C2
        lambdaFn = lambda_.Function(self,
                                    "nimbusC2-operator-lambda-function",
                                    code=lambda_.Code.from_asset(
                                        "../bootstrap.zip"),
                                    handler="bootstrap",
                                    architecture=lambda_.Architecture.ARM_64,
                                    runtime=lambda_.Runtime.PROVIDED_AL2
                                    )

        # ToDo - create lambda accessed by implants

        # give lambda needed S3 permissions
        # ToDo - roll this back to fewer permissions
        # lambdaFn.role.add_managed_policy(iam_.ManagedPolicy.from_managed_policy_arn(self, "nimbusc2-iam-policy", "arn:aws:iam:aws:policy/AmazonS3FullAccess"))
