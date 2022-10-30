#!/usr/bin/env python3

class NimbusAgent:
    def __init__(self):
        self.session_id = self.get_session_id()
        self.api_key = self.get_api_key()
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
        self.sleep_interval = self.get_sleep_interval()
        self.kill_date = self.get_kill_date()

    def get_session_id(self):
        return "testing"

    def get_api_key(self):
        return "testing"
    
    def get_hostname(self):
        return "testing"

    def get_os_name(self):
        return "testing"

    def get_os_ver(self):
        return "testing"

    def get_kernel_ver(self):
        return "testing"

    def get_system_time(self):
        return "testing"

    def get_pid(self):
        return "testing"

    def get_parent_pid(self):
        return "testing"

    def get_proc_name(self):
        return "testing"

    def get_agent_dir(self):
        return "testing"

    def get_working_dir(self):
        return "testing"

    def get_username(self):
        return "testing"

    def get_internal_ip(self):
        return "testing"

    def get_external_ip(self):
        return "testing"

    def get_sleep_interval(self):
        return "testing"

    def get_kill_date(self):
        return "testing"


def main():
    agent = NimbusAgent()
    attrs = vars(agent)
    print(attrs)


if __name__ == "__main__":
    main()
