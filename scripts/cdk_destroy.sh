#!/bin/bash

cd nimbus_cdk
python3 -m venv .venv
source .venv/bin/activate
pip install -r requirements.txt

cdk destroy --all --force