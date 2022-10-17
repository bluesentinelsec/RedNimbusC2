# RedNimbusC2
A serverless C2 framework

## Pardon Our Dust

This repository is under active development.

Content is not expected to be stable or even usable at this time.

## Overview

![alt text](images/nimbusC2_architecture.png)

## Prerequisites

### Build Dependencies

You must install the following resources in order to build and operate Red Nimbus C2:

1. [Node.js](https://nodejs.org/en/)
2. [Python](https://www.python.org)
3. [Go](https://go.dev)
4. [AWS CLI](https://docs.aws.amazon.com/cli/latest/userguide/getting-started-install.html)
5. [AWS CDK](https://aws.amazon.com/getting-started/guides/setup-cdk/)
6. [Make Build System](https://www.gnu.org/software/make/)
7. [Git revision control system](https://git-scm.com)

### AWS Account

Red Nimbus C2 makes exclusive use of AWS cloud services.

For this reason, you must have your own AWS account.

Instructions on creating an AWS account are provided [here](https://aws.amazon.com/premiumsupport/knowledge-center/create-and-activate-aws-account/).

## Installation

1. Clone the repository

```bash
git clone https://github.com/bluesentinelsec/RedNimbusC2.git
```

2. Deploy Red Nimbus C2 to AWS using CDK

```bash
# enter the RedNimbusC2 directory
cd RedNimbusC2

# deploy RedNimbusC2 infrastructure
make cdk-deploy
```

3. Validate your installation

```bash
make test
```

## Operator Instructions

TBD

### Generate Agent

TBD

### Deploy Agent to Target System(s)

TBD

### Issue Commands

TBD

