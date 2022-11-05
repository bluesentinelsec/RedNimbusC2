#!/usr/bin/env python3

import agent_tasks
import s3wrapper

test_session_id = "test-session-123"


def post_test_task():
    bucket = s3wrapper.get_s3_bucket_name()
    dst_file = "tasks/test-task"
    s3wrapper.put_s3_file(bucket, "test_task.json", dst_file)


def test_get_task():
    post_test_task()
    task = agent_tasks.get_task(test_session_id)
    assert (task != "")
