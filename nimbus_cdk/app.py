#!/usr/bin/env python3
import os

import aws_cdk as cdk

from nimbus_cdk.nimbus_cdk_stack import NimbusCdkStack


os.system("cd ../ && make deployment-package")
app = cdk.App()
NimbusCdkStack(app, "nimbusC2")

app.synth()
