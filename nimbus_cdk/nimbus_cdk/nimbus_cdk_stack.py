import aws_cdk as cdk
from aws_cdk import (
    # Duration,
    Stack,
    aws_lambda as lambda_,
    aws_s3 as s3_,
    aws_apigateway as aws_apigateway_
)
from constructs import Construct


class NimbusCdkStack(Stack):

    def __init__(self, scope: Construct, construct_id: str, **kwargs) -> None:
        super().__init__(scope, construct_id, **kwargs)

        # define Lambda function that handles C2
        operatorHandler = lambda_.Function(self,
                                           "nimbusC2-operator-handler",
                                           code=lambda_.Code.from_asset(
                                               "../bootstrap.zip"),
                                           handler="bootstrap",
                                           architecture=lambda_.Architecture.ARM_64,
                                           runtime=lambda_.Runtime.PROVIDED_AL2,
                                           function_name="nimbusC2Handler",
                                           timeout=cdk.Duration.seconds(300)
                                           )

        # define Lambda function that handles agent requests

        agentHandler = lambda_.Function(self,
                                        "nimbusC2-agent-handler",
                                        code=lambda_.Code.from_asset(
                                            "../nimbusc2_agent_lambda/"),
                                        handler="nimbusc2_agent_lambda.handler",
                                        architecture=lambda_.Architecture.ARM_64,
                                        runtime=lambda_.Runtime.PYTHON_3_9,
                                        function_name="nimbusC2-agent-handler",
                                        timeout=cdk.Duration.seconds(300)
                                        )

        # define API gateway
        # this routes agent task requests
        # to the agent Lambda
        nimbusc2_agent_api = aws_apigateway_.LambdaRestApi(self, "nimbusC2-api-gateway",
                                                           handler=agentHandler)

        # print the agent URL to the console
        # after deployment
        cdk.CfnOutput(self, "NimbusC2-Agent-URL",
                      value=nimbusc2_agent_api.url)

        # define S3 bucket used for storing operator data
        operator_bucket_name = f"red-nimbus-c2-{cdk.Aws.REGION}-{cdk.Aws.ACCOUNT_ID}"
        operator_bucket = s3_.Bucket(self, "red-nimbus-c2",
                                     access_control=s3_.BucketAccessControl.BUCKET_OWNER_FULL_CONTROL,
                                     block_public_access=s3_.BlockPublicAccess.BLOCK_ALL,
                                     encryption=s3_.BucketEncryption.S3_MANAGED,
                                     removal_policy=cdk.RemovalPolicy.DESTROY,
                                     auto_delete_objects=True,
                                     bucket_name=operator_bucket_name)

        # define bucket for unit tests
        test_bucket_name = f"red-nimbus-c2-testing-{cdk.Aws.REGION}-{cdk.Aws.ACCOUNT_ID}"
        test_bucket = s3_.Bucket(self, "red-nimbus-c2-testing",
                                 access_control=s3_.BucketAccessControl.BUCKET_OWNER_FULL_CONTROL,
                                 block_public_access=s3_.BlockPublicAccess.BLOCK_ALL,
                                 encryption=s3_.BucketEncryption.S3_MANAGED,
                                 removal_policy=cdk.RemovalPolicy.DESTROY,
                                 auto_delete_objects=True,
                                 bucket_name=test_bucket_name)

        # set bucket permissions
        # ToDo - try to reduce permissions for more opsec
        operator_bucket.grant_read_write(operatorHandler)
        test_bucket.grant_read_write(operatorHandler)

        operator_bucket.grant_read_write(agentHandler)
        test_bucket.grant_read_write(agentHandler)
