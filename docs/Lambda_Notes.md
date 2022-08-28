# Lambda Notes

- Lambda lets you run code without provisioning or managing servers
- Lambda removes need for operating system configuration and maintenance
- Scales automatically
- Only pay for compute time you use; no charge while your Lambda is not running
  - also charged for data exchanged between services
- you can mount Elastic File System directory to a local directory
- Execution environment - a secure isolated runtime environment in which Lambda executes your code
- supports code signing; ensures code has not been altered since signing

```
# view current lambda functions
aws lambda list-functions
```

- Infrastructure as code option: `AWS SAM CLI`

## Lambda Terminology
- Function - the code you run on Lambda
- Trigger - configuration that invokes your Lambda function
- Event - JSON formatted request; is an input to your Lambda function
- Deployment package - either a zip archive or a container
- Destination - a resource Lambda can send events to
- Function URLs - offers built in HTTPS endpoint

