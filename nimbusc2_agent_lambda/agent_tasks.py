#!/usr/bin/env python3

"""
agent_tasks.py defines functions for
getting agent tasks and posting
agent task output to S3
"""

from os.path import normpath, basename

import json
import logging
import os

import s3wrapper


def get_task(session_id):
    """
    read each file in the S3 task folder
    and return the task for the specified
    session id
    """
    logging.info(f"checking tasks for session: {session_id}")

    bucket_name = s3wrapper.get_s3_bucket_name()
    task_files = s3wrapper.list_s3_files(bucket_name, "tasks")

    logging.info(f"task files in s3: {task_files}")

    pending_task = ""

    # read each task file and find a task for this session ID
    for task_file in task_files:

        # download each task file from S3
        dst_file = f"/tmp/{basename(normpath(task_file))}"
        s3wrapper.get_s3_file(bucket_name, task_file, dst_file)

        # read the local task file as dict
        logging.info(f"reading task file {dst_file}")
        with open(dst_file, "r") as fp:
            task_data = json.load(fp)

        # we found a pending task
        if session_id == task_data['sessionID']:
            pending_task = task_data
            logging.info("found a pending task:")
            print(pending_task)

            # delete pending task, otherwise we'll run it repeatedly
            logging.info("removing pending task from S3")
            s3wrapper.remove_s3_file(bucket_name, task_file)

        # delete each task file saved to disk
        logging.info(f"deleting local task file: {dst_file}")
        os.remove(dst_file)

    # return the task for the intended session
    return pending_task