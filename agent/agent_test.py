#!/usr/bin/env python3

import json

import agent


def test_decode_task():
    task = "eyJ0YXNrSUQiOiAiYzViZGE1N2MtN2RlNC00ZDFhLWFlOWItMGIzMjY0MGY1NzQ5IiwgInNlc3Npb25JRCI6ICJkZjY4YjM2NGIwNWFiZTQ5ZDMzOGU2ZWJkMDY1MDdlMCIsICJsYW1iZGFIYW5kbGVyIjogIkhhbmRsZVNldExhbWJkYVRhc2siLCAiZ3JvdXBOYW1lIjogIiIsICJleGVjVGltZSI6ICIiLCAidGFzayI6ICJleGVjLWNtZCIsICJhcmd1bWVudHMiOiBbImFyZzEiLCAiYXJnMiIsICJhcmczIl19"
    decoded_task = agent.decode_task(task)
    assert (decoded_task != "")
    assert(decoded_task["task"] != "")
    assert(decoded_task["arguments"] != [])

def test_exec_cmd():
    
    task_files = ["test_task.json", "test_task_2.json", "test_task_3.json"]
    for file in task_files:
        with open(file, "r") as fp:
            task = json.load(fp)
            print(task)
            out = agent.exec_cmd(task)
            assert(out != "")
            assert(out != None)
