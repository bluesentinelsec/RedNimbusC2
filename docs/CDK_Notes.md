# CDK Notes

- CDK allows you to programatically deploy cloud resources
- output of an AWS CDK program is an AWS CloudFormation template
- use CDK CLI toolkit to work with your CDK apps
- CDK **Constructs** defines concrete AWS resources (S3 bucket, Lambda, etc.)


## Setup Notes

### Pre-reqs

1. Created AWS account, and enabled MFA
2. Created an IAM Admin user + Admin group
3. Created access key and secret key for CLI access under IAM Admin user

### AWS CLI

1. Installed AWS CLI for MacOS

```bash
curl "https://awscli.amazonaws.com/AWSCLIV2.pkg" -o "AWSCLIV2.pkg"
sudo installer -pkg AWSCLIV2.pkg -target /
which aws
```

2. Configure AWS CLI to use your access key and secret key

```
aws configure
```

### AWS CDK

1. Install node

```
brew install node
```

2. Install CDK

```
npm install -g aws-cdk

cdk --version
```

## CDK Tutorial

```
mkdir cdk_workshop && cd cdk_workshop

cdk init sample-app --language python

source .venv/bin/activate

pip install -r requirements.txt
```

- app.py is the entry point
- use `cdk synth` to generate your cloud formation template
- must be in same directory as `cdk.json`
- use `cdk bootstrap` prior to first deployment
- use `cdk deploy` to deploy your app

## Other solutions to look at
- Terraform
- Kubernetes