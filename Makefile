AWS_PROFILE = ""

all: operator create-agent deployment-package cdk-deploy

# build the nimbusc2 operator client
operator:
	go build -o release/nimbusc2 -ldflags="-s -w" cmd/redOperator/main.go

# build app that creates nimbusc2 agent executables
create-agent:
	go build -o release/nimbusc2-create-agent -ldflags="-s -w" cmd/createAgent/main.go

# build the AWS Lambda deployment package
deployment-package:
	GOOS=linux GOARCH=arm64 go build -o bootstrap -ldflags="-s -w" cmd/lambdaC2/main.go
	zip bootstrap.zip bootstrap

# deploy the AWS CDK app
cdk-deploy: deployment-package
	./scripts/deploy.sh $$AWS_PROFILE

# installs the nimbusc2 operator client
install:
	install release/nimbusc2 /usr/local/bin/
	install release/nimbusc2-create-agent /usr/local/bin/

# removes the nimbusc2 operator client 
# and destroys the CDK app
uninstall:
	rm -f /usr/local/bin/nimbusc2
	rm -f /usr/local/bin/nimbusc2-create-agent
	./scripts/destroy.sh

test:
	go test ./pkg/shellexec/
	go test ./pkg/cli

# intended to be run locally so that
# GitHub Actions does not rack up
# expenses under my AWS account
test-aws:
	go test ./...
