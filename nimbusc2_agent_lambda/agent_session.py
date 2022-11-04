#!/usr/bin/env python3

"""
agent_session.py contains functions for
registering and removing Nimbus C2 agent
sessions 
"""

import hashlib
import json
import logging

import s3wrapper


def register_agent(session_id, event_body):
    """
    register an agent with the control server
    """

    logging.info(f"registering agent: {session_id}")

    # write session file to /tmp/
    session_file = f"/tmp/{session_id}"

    # convert event_body to json string and write to /tmp/
    event_body_str = json.dumps(event_body)
    with open(session_file, "w") as fp:
        fp.write(event_body_str)

    # get S3 bucket name
    bucket_name = s3wrapper.get_s3_bucket_name()

    # upload session file to S3
    s3_dst_file = f"sessions/{session_id}"
    s3wrapper.put_s3_file(bucket_name, session_file, s3_dst_file)
    return


def remove_agent(session_id):
    """
    delete an agent session from the control server
    """
    logging.info(f"removing agent: {session_id}")

    session_file = f"sessions/{session_id}"
    bucket_name = s3wrapper.get_s3_bucket_name()
    s3wrapper.remove_s3_file(bucket_name, session_file)

    return


def is_agent_registered(session_id):
    """
    check if an agent session is registered
    to the control server
    """
    logging.info(f"checking if agent is registered: {session_id}")

    session_file = f"sessions/{session_id}"
    bucket_name = s3wrapper.get_s3_bucket_name()

    session_files = s3wrapper.list_s3_files(bucket_name, "sessions/")
    for each_session in session_files:
        if each_session == session_file:
            return True
    return False


def derive_session_id(hostname: str, user: str, cwd: str, agent_pid) -> str:
    """
    generate a unique session ID
    based on hostname, curent user,
    and current working directory
    """
    seed = hostname + user + cwd + str(agent_pid)
    session_id = hashlib.md5(seed.encode('utf-8')).hexdigest()
    return session_id
