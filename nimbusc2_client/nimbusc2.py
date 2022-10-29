#!/usr/bin/env python3

# standard library
import argparse
import base64
import json
import logging
import os
import uuid

# external modules
import boto3


def set_aws_profile(profile: str):
    """
    controls which AWS account is used for requests
    see here for info on creating AWS profiles:
    https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-profiles.html    
    """
    os.environ["AWS_PROFILE"] = profile
    logging.debug(f"set aws profile to: {profile}")


def create_task_id() -> str:
    """
    create an ID used to uniquely
    identify each agent task
    """
    logging.debug("creating new task")
    return str(uuid.uuid4())


def parse_task_args(args) -> list[str]:
    """
    convert task args string into list
    """
    arg_list = []
    if args.args:
        arg_list = args.args.split(",")
    else:
        arg_list = [""]
    return arg_list


def new_request(args):
    """
    create a new task request based on CLI args
    """
    logging.debug("initializing new request")

    arg_list = parse_task_args(args)

    if args.task_id == "":
        args.task_id = create_task_id()

    request = {
        "taskID": args.task_id,
        "sessionID": args.session_id,
        "lambdaHandler": "",
        "groupName": args.session_group,
        "execTime": args.time,
        "task": args.cmd,
        "arguments": arg_list
    }
    return request


def invoke_nimbusc2_lambda(request):
    """
    send the task request to Red Nimbus C2
    lambda function
    """
    request_json = json.dumps(request, indent=4)
    print(request_json)

    lambda_client = boto3.client('lambda')

    logging.info("sending request to AWS Lambda: nimbusC2Handler")
    response = lambda_client.invoke(FunctionName='nimbusC2Handler',
                                    InvocationType='RequestResponse',
                                    Payload=request_json)

    logging.info("nimbusC2Handler response:")
    console_output = json.loads(response['Payload'].read())
    console_output = json.dumps(console_output, indent=4, sort_keys=True)
    return console_output


def invoke_set_task(args):
    logging.info("creating set-task request")
    request = new_request(args)
    request["lambdaHandler"] = "HandleSetLambdaTask"
    out = invoke_nimbusc2_lambda(request)
    print(out)


def invoke_get_task(args):
    logging.info("creating get-task request")
    request = new_request(args)
    request["lambdaHandler"] = "HandleGetLambdaTask"
    out = invoke_nimbusc2_lambda(request)
    out = base64.b64decode(out)
    out_pretty = json.loads(out)
    print(json.dumps(out_pretty, indent=4))


def invoke_remove_task(args):
    logging.info("creating remove-task request")
    request = new_request(args)
    request["lambdaHandler"] = "HandleRemoveLambdaTask"
    out = invoke_nimbusc2_lambda(request)
    print(out)


def invoke_get_session(args):
    logging.info("creating get-session request")
    request = new_request(args)
    request["lambdaHandler"] = "HandleGetSession"
    out = invoke_nimbusc2_lambda(request)
    print(out)


def invoke_list_sessions(args):
    logging.info("creating list-sessions request")
    request = new_request(args)
    request["lambdaHandler"] = "HandleListSessions"
    out = invoke_nimbusc2_lambda(request)
    print(out)


def invoke_remove_session(args):
    logging.info("creating remove-sessions request")
    request = new_request(args)
    request["lambdaHandler"] = "HandleRemoveSession"
    out = invoke_nimbusc2_lambda(request)
    print(out)


def main(args):

    # setup console logging
    if args.verbose:
        logging.basicConfig(level=logging.DEBUG,
                            format='%(levelname)s (%(filename)s:%(lineno)s) %(message)s')
        logging.debug("enabled verbose console logging")

    else:
        logging.basicConfig(level=logging.INFO,
                            format='%(levelname)s (%(filename)s:%(lineno)s) %(message)s')

    # handle --aws-profile
    if args.aws_profile:
        set_aws_profile(args.aws_profile)

    # handle agent task commands
    if args.set_task:
        invoke_set_task(args)
        return

    if args.get_task:
        invoke_get_task(args)
        return

    if args.remove_task:
        invoke_remove_task(args)
        return

    # handle agent session commands
    if args.get_session:
        invoke_get_session(args)
        return

    if args.list_sessions:
        invoke_list_sessions(args)
        return

    if args.remove_session:
        invoke_remove_session(args)
        return


if __name__ == "__main__":
    # setup CLI
    parser = argparse.ArgumentParser()

    # define sub-commands
    parser.add_argument("--set-task", action="store_true")
    parser.add_argument("--get-task", action="store_true")
    parser.add_argument("--remove-task", action="store_true")
    parser.add_argument("--get-session", action="store_true")
    parser.add_argument("--list-sessions", action="store_true")
    parser.add_argument("--remove-session", action="store_true")

    # define agent-related arguments
    parser.add_argument("-s", "--session-id", required=False, default="")
    parser.add_argument("-g", "--session-group", required=False, default="")
    parser.add_argument("-i", "--task-id", required=False, default="")
    parser.add_argument("-c", "--cmd", required=False, default="")
    parser.add_argument("-a", "--args", required=False, default="")
    parser.add_argument("-t", "--time", required=False, default="")
    parser.add_argument("-k", "--key-env", required=False, default="")

    # define general arguments
    parser.add_argument("--verbose", action="store_true")
    parser.add_argument("--aws-profile", required=False, default="")

    # parse CLI args
    args = parser.parse_args()

    main(args)
