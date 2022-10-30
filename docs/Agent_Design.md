# Agent Design

## Agent Functions

- **get task**: HTTP POST - check if C2 server has tasking
    - C2 server should automatically register new agent sessions

- **post task output**: HTTP POST - post task output to C2 server
    - C2 server should drop output from unknown sessions

- Agents should have an API key to filter/control access

## Agent Schema

### Get Task

```json
{
    "sessionID": "md5 hash of hostname + user account",
    "apikey": "apikey",
    "hostname": "hostname",
    "osName": "os name",
    "osVersion": "os version",
    "kernelVersion": "kernel version",
    "systemTime": "time unix epoch",
    "agentPID": 1234,
    "agentParentPid": 1,
    "agentProcName": "process name",
    "agentDirectory": "agent location",
    "workingDirectory": "working dir",
    "userName": "userName",
    "internalIP": "internal IP Address",
    "externalIP": "external IP address",
    "sleepInterval": "sleep in seconds + jitter",
    "killDate:": "kill date unix epoch",
}
```

### Post Task

```json
{
    "sessionID": "md5 hash of hostname + user account",
    "apikey": "apikey",
    "taskOutput": "base64 blob?"
}
```

