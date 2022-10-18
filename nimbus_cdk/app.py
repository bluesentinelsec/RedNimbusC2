#!/usr/bin/env python3

import aws_cdk as cdk
from nimbus_cdk.nimbus_cdk_stack import NimbusCdkStack

app = cdk.App()
NimbusCdkStack(app, "nimbusC2")

app.synth()
