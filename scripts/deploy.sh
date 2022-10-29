#!/bin/sh -x

# setup python virtual environment
cd nimbus_cdk
python3 -m venv .venv
source .venv/bin/activate
pip install -r requirements.txt

# bootstrap CDK for first time use
cdk bootstrap 

# confirm we can build Red Nimbus C2
cdk synth

# deploy to AWS
cdk deploy --all --require-approval never $1