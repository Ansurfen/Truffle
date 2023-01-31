import threading
import os
import yaml
from abc import ABC, abstractmethod
import sys
import time
import logging
from rich.progress import track, Progress
import subprocess
logging.basicConfig(level=logging.DEBUG, filename="debug.log", filemode="w",
                    format="%(asctime)s - %(name)s - %(levelname)-9s - %(filename)-8s : %(lineno)s line - %(message)s", datefmt="%Y-%m-%d %H:%M:%S")
rootPath: str = ''
garbages = []


class OptType:
    COMMON = 1
    MQ = 2
    DOCKER_DESKTOP = 3


class Opt(ABC):

    @abstractmethod
    def Type(self) -> OptType:
        pass


class ServiceOpt:
    rpcPort: int
    httpPort: int
    name: str
    dir: str
    filterList: []

    def __init__(self, name: str, rpcPort: int, httpPort: int, dir: str, filterList: [] = ...):
        self.name = name
        self.rpcPort = rpcPort
        self.httpPort = httpPort
        self.dir = dir
        self.filterList = filterList

    @property
    def Type(self) -> OptType:
        return OptType.COMMON


class MiddlewareOpt:
    optType: OptType

    def __init__(self, type: OptType):
        self.optType = type

    @property
    def Type(self) -> OptType:
        return self.optType


def newConsoleExec(cmd: str) -> None:
    if (sys.platform == 'win32'):
        res = subprocess.run(
            "start powershell.exe cmd /k '{cmd}'".format(cmd=cmd), shell=True)
    else:
        res = subprocess.run(
            "gnome-terminal -e 'bash -c \"{cmd}; exec bash\"'".format(cmd=cmd), shell=True)
    print(res)


def runService(opt: Opt) -> None:
    if (opt.Type == OptType.MQ and isinstance(opt, MiddlewareOpt)):
        os.chdir('D:\\apache-zookeeper-3.8.0-bin\\bin')
        newConsoleExec('zkServer.cmd')
        time.sleep(1)
        os.chdir('D:\kafka_2.12-3.2.0')
        newConsoleExec(
            '.\\bin\windows\kafka-server-start.bat .\config\server.properties')
        time.sleep(3)
    elif (opt.Type == OptType.DOCKER_DESKTOP and isinstance(opt, MiddlewareOpt)):
        os.chdir('C:\Program Files\Docker\Docker')
        newConsoleExec('Docker Desktop.exe')
        with Progress() as progress:
            task = progress.add_task("[green]Launching docker...", total=1000)
            while not progress.finished:
                progress.update(task, advance=10)
                time.sleep(0.02)
        input("Enter any keyword to continue")
    elif (opt.Type == OptType.COMMON and isinstance(opt, ServiceOpt)):
        with open('application.yaml', 'r+', encoding='utf-8') as fp:
            res = yaml.load(fp.read(), Loader=yaml.FullLoader)
            res['server']['base'] = res['server']['port'] = opt.rpcPort
            res['server']['http'] = opt.httpPort
            res['server']['name'] = opt.name
            with open("." + opt.dir + '/application.yaml', 'w+', encoding='utf-8') as target:
                target.seek(0)
                target.truncate()
                if (opt.filterList != Ellipsis):
                    for rule in opt.filterList:
                        res.pop(rule)
                yaml.dump(data=res, stream=target, allow_unicode=True)
                target.close()
            fp.close()
        os.chdir("." + opt.dir)
        newConsoleExec('go run main.go')
        logging.info("{name} service start...".format(name=opt.name))
        garbages.append(rootPath + opt.dir)
        time.sleep(0.5)

    os.chdir(rootPath)


todoList: dict = {
    "desktop": {
        "enable": False,
        "opt": MiddlewareOpt(OptType.DOCKER_DESKTOP)
    },
    "kafka": {
        "enable": True,
        "opt": MiddlewareOpt(OptType.MQ)
    },
    "user": {
        "enable": False,
        "opt": ServiceOpt('user', 8000, 8005, "/user", ['kafka', 'langs', 'timer', 'email', 'breaker'])
    },
    "captcha": {
        "enable": False,
        "opt": ServiceOpt('captcha', 8010, 8015, "/captcha", ['db', 'kafka', 'redis', 'tracer', 'langs'])
    },
    "gateway": {
        "enable": False,
        "opt": ServiceOpt('gateway', 8020, 8025, "/gateway", ['etcd', 'db', 'redis', 'kafka', 'tracer', 'langs', 'timer', 'email', 'breaker'])
    },
    "log": {
        "enable": False,
        "opt": ServiceOpt('log', 8030, 8035, "/log", ['breaker', 'db', 'timer', 'redis', 'email', 'langs'])
    },
    "i18n": {
        "enable": False,
        "opt": ServiceOpt('i18n', 8040, 8045, "/i18n", ['breaker', 'db', 'email', 'kafka', 'redis', 'timer', 'langs'])
    },
    "topic": {
        "enable": False,
        "opt": ServiceOpt('topic', 8060, 8065, "/topic", ['kafka', 'email', 'timer', 'breaker', 'langs'])
    },
    "ws": {
        "enable": False,
        "opt": ServiceOpt('ws', 8070, 8075, "/ws", ['email', 'timer', 'breaker', 'langs'])
    },
}


if __name__ == '__main__':
    path = os.getcwd()
    if (path.rfind("auto") + len("auto") == len(path)):
        os.chdir(os.path.abspath('..'))
    rootPath = os.getcwd()
    for serviceName, serviceOpt in todoList.items():
        if (serviceOpt['enable']):
            runService(serviceOpt['opt'])
    time.sleep(10)
    for garbage in garbages:
        os.remove("{dir}/application.yaml".format(dir=garbage))
