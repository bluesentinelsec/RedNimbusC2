#!/usr/bin/env python3

import argparse
import getpass
import hashlib
import json
import logging
import os
import random
import socket
import sys
import time
import urllib.request
import platform


class NimbusAgent:
    def __init__(self, args):

        # members set via CLI
        self.api_key = ""
        self.url = ""
        self.sleep_interval = ""
        self.kill_date = ""

        # members set at runtime
        self.session_id = self.get_session_id()
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

    def get_api_key(self):
        return self.api_key

    def set_api_key(self, api_key: str):
        self.api_key = api_key

    def get_url(self):
        return self.url

    def set_url(self, url: str):
        self.url = url

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
            with urllib.request.urlopen('https://ifconfig.me/') as response:
                html = response.read()
                external_ip = html.decode("utf-8")
        except:
            pass
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
        logging.warn("sorry, this function is not implemented")

# -=-=-=-=-=-=-=-=-=-=-=-=-=
#   Task Commands
# -=-=-=-=-=-=-=-=-=-=-=-=-=

    def get_tasking(self):
        logging.warn("sorry, this function is not implemented")
        return True

    def post_tasking_output(self):
        logging.warn("sorry, this function is not implemented")

    def exec_tasking(self, task, arguments):
        logging.warn("sorry, this function is not implemented")


# -=-=-=-=-=-=-=-=-=-=-=-=-=
#   Agent Commands
# -=-=-=-=-=-=-=-=-=-=-=-=-=

def download_file(url, dst):
    return "sorry this feature is not implemented"


def upload_file(url, src):
    return "sorry this feature is not implemented"


def exec_cmd(cmd):
    return "sorry this feature is not implemented"


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
    agent.set_url(args.url)
    agent.set_api_key(args.api_key)
    agent.set_sleep_interval(args.sleep_interval)
    agent.set_kill_date(args.kill_date)

    # display agent settings for debugging purposes
    agent_settings = json.dumps(vars(agent), indent=4)
    logging.info("agent settings:")
    print(agent_settings)

    while True:
        # check kill date
        logging.debug("checking kill date")
        agent.check_kill_date()

        # sleep
        logging.debug(f"sleeping for {agent.get_sleep_interval()}")
        time.sleep(agent.get_sleep_interval())

        # check if tasking
        logging.debug("getting task from Red Nimbus C2")
        task = agent.get_tasking()

        if not task:
            logging.debug("did not receive a task, restarting C2 loop")
            continue
        else:
            logging.debug("received new task")

        # sleep again
        time.sleep(agent.get_sleep_interval())

        # execute task
        logging.debug("executing task")

        # sleep again
        time.sleep(agent.get_sleep_interval())

        # post task output
        logging.debug("posting task output")


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
