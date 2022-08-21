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

```bash
# issue singular tasking
./set_tasking --session-id uuid --cmd supportedCmd --args yourArguments

# task multiple sessions
./set_tasking --session-id uuid1,uuid2,uuid3 --cmd supportedCmd --args yourArguments

# task every session
./set_tasking --session-id all --cmd supportedCmd --args yourArguments

# task implants in a group
./set_tasking --session-group yourGroup --cmd supportedCmd --args yourArguments
```

### Get tasking output

- task output can stream into a console window

- task output can be pulled from S3
  - perhaps queried with Athena?

- should we use CloudWatch for task output?

- Send notifications to Slack? Text message? Email?
