import hashlib
import json
import logging
#import boto3


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
    logging.info("stub - handle_get_task")

    # derive session ID from seed values: hostname, username, and agent directory
    session_id = derive_session_id(
        event_body["hostname"], event_body["username"], event_body["agent_dir"], event_body["agent_pid"])

    logging.info(f"agent session ID is: {session_id}")

    # is this an existing session
    is_registered = is_agent_registered(session_id)

    if not is_registered:
        register_agent(session_id, event_body)

    task = get_task(session_id)

    delete_queued_task(session_id)


    return {
        'statusCode': 200,
        'headers': {
            'Content-Type': 'text/plain'
        },
        'body': task
    }

def get_task(session_id):
    logging.info(f"checking tasks for session: {session_id}")
    # read every file in the tasks bucket
    # if the session_id == task_file["session_id"]:
        # get the task and arguments
        # put task and args in a list 
    return "whoami"

def delete_queued_task(session_id):
    logging.info(f"deleting queued tasks for: {session_id}")
    # list every file in S3
    # if file name equals session id
        # delete file
    return


def handle_post_task_output(event_body):
    logging.info("stub - handle_post_task_output")
    return {
        'statusCode': 200,
        'headers': {
            'Content-Type': 'text/plain'
        },
        'body': 'stub - handle_post_task_output'
    }


def register_agent(session_id, event_body):
    logging.info(f"registering agent: {session_id}")
    # write to S3 sessions folder
    return


def remove_agent(session_id):
    logging.info(f"removing agent: {session_id}")
    # delete from S3 sessions folder
    return


def is_agent_registered(session_id):
    logging.info(f"checking if agent is registered: {session_id}")
    # list every file in S3 sessions folder
    # if filename == sessionID
        # return true
    # else return false
    return True


def derive_session_id(hostname: str, user: str, cwd: str, agent_pid) -> str:
    """
    generate a unique session ID
    based on hostname, curent user,
    and current working directory
    """
    seed = hostname + user + cwd + str(agent_pid)
    session_id = hashlib.md5(seed.encode('utf-8')).hexdigest()
    return session_id
