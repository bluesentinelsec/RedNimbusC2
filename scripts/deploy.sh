#!/bin/sh -x

# Deploys Red Nimbus C2 to AWS

# Build and package Operator Lambda
GOOS=linux GOARCH=arm64 go build -o bootstrap -ldflags="-s -w" cmd/lambdaC2/main.go
zip bootstrap.zip bootstrap

# Setup python virtual environment
cd nimbus_cdk
python3 -m venv .venv
source .venv/bin/activate
pip install -r requirements.txt

# Bootstrap CDK for first time use
cdk bootstrap 

# Confirm Red Nimbus C2 builds properly
cdk synth

# Deploy Red Nimbus C2 to AWS
cdk deploy --all --require-approval never --outputs-file ../nimbus_c2_url.json $1

cat ../nimbus_c2_url.json