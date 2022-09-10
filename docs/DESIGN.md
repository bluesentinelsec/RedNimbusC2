# Design

## Proposed Workflow

1. Red operator deploys the AWS serverless core (Lambdas, S3, etc.)
     -maybe use CDK for this?

2. Red operator generates an implant
  - do this locally?
  - lambda dynamically generates payload, stores in S3?
  - Why not both?

3. Red operator deploys implant to target
  - how this happens is up to you

4. Implant is executed on target system; implant calls back to implant-handler Lambda
  - effective access control will be key here

5. Red operator issues taskings through a locked down operator lambda

6. storage and state will be placed in S3
  - upload/download files
  - upload/download implant taskings
  - upload/download implant output

7. CloudWatch
  - provides detailed logs
  - red operator accountability

8. Stretch goal - support event based triggers
  - new software installed -> send operator alert
  - low battery -> extend sleep
  - suspicious processes spawned (debuggers, sysinternals, volatility, etc.) -> kill implant, purge artifacts, force reboot

## User Experience

### Installation

```bash
# clone the repo
git clone https://github.com/bluesentinelsec/RedNimbusC2

# install and configure all dependencies
./scripts/install.sh

# boom - you are ready to operate
```

### Generate Implants

```bash
# granular example; most fields will be optional
./create_implant --format exe --arch amd64 --platform windows --c2url https://whatever \
    --callback-frequency 10 --kill-date 365 \
    --implant-group
    --allow-hostname myTarget --allow-user userName --allow-ip 10.10.10.10 \
    --verbose 
```

### Issue taskings
Lambda name: submitTasking

Taskings are sent via HTTP POST requests to API Gateway

API gateway passes to Lambda -> Lambda passes to S3; writes to a bucket - "taskings"

Security here is critical - only red operators should
be able to issue tasks

```bash
# issue singular tasking
./set-task --session-id uuid --cmd supportedCmd --args yourArguments

# task multiple sessions
./set_tasking --session-id uuid1,uuid2,uuid3 --cmd supportedCmd --args yourArguments

# task every session
./set_tasking --session-id all --cmd supportedCmd --args yourArguments

# task implants in a group
./set_tasking --session-group yourGroup --cmd supportedCmd --args yourArguments

# remove task
./remove-task --id <task ID>

# read task
./get-task --id <task ID>

# note: if 'nimbus_secret' env var is set, encrypt command with that key
# implant will have to know about the key
```

### Supported Commands

```
# download a file from target to S3
get-file <src file>,<dst file>

# put file from S3 to target
put-file <src file>,<dst file>

# target system will download a file from a URL
download-file <URL>,<dst file>

# target system will upload a file via HTTP POST to URL
upload-file <URL>

# run command from native shell environment
exec-cmd <command>

# execute a base64 encoded command
exec-b64cmd <base64 encoded command>

# list files in directory
list-files <file or dir>

# list files recursive
list-files-recursive <dir>

# get process listing
get-process

# get implant process ID
get-pid

# get network connections
get-netstat

# get network interface configuration
get-ifconfig

# load dynamic library in process
load-library <library in S3> <pid or proc name>

# load shellcode in process
load-shellcode <shellcode in S3> <pid or proc name>

# exit from the system and delete all artifacts
terminate-session

# change secret key; this will read an env variable, 'nimbus_secret'
set-secret

# set sleep interval
set-sleep <int>

# if implant fails to connect after X times, cleanup and terminate
set-retry-limit <int>
```

### Task structure

Unencrypted:
```
taskID: string
task: string
arguments: []string
sessionID: string
group: string
time: string
```

Encrypted
```
data: <encrypted task object>
```

### Get tasking output

tasks should get written to an S3 bucket - `output`

task output structure:

```
sessionID: string
hostname: string
osName: string
osVersion: string
kernel: string
implantUser: string
implantDir: string
implantProcessID: int
taskID: string
taskIssued: time
taskComplete: time
task: string
error: string
output: string
```

- get psp's should be a module / plugin

- task output can stream into a console window

- task output can be pulled from S3
  - perhaps queried with Athena?

- should we use CloudWatch for task output?

- Send notifications to Slack? Text message? Email?
