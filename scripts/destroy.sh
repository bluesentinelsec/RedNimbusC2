#!/bin/sh -x

# Deletes all Red Nimbus C2 components and data from AWS

# setup python virtual environment
cd nimbus_cdk
python3 -m venv .venv
source .venv/bin/activate
pip install -r requirements.txt

# delete all Red Nimbus C2 AWS resources
cdk destroy --all --force $1