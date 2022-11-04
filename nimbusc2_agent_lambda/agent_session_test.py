#!/usr/bin/env python3

import json

import agent_session

test_session_id = "test-session-id-123"


def test_register_agent():
    event_data = ""
    with open("test_session.json", "r") as fp:
        event_data = json.load(fp)

    agent_session.register_agent(test_session_id, event_data)


def test_is_agent_registered():
    session_exists = agent_session.is_agent_registered(test_session_id)
    assert (session_exists == True)


def test_remove_agent():
    agent_session.remove_agent(test_session_id)


def test_derive_session_id():
    got = agent_session.derive_session_id("hostname", "user", "cwd", 1234)
    want = "66c4b81427185759e15f6d6ddbc37cfd"
    assert (got == want)
