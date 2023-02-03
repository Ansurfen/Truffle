import paramiko
import threading


class SSHOpt(object):
    ip: str
    port: int
    user: str
    pwd: str
    key: str

    def __init__(self, ip: str, port: int, user: str, pwd: str, key: str = ...):
        self.ip = ip
        self.port = port
        self.user = user
        self.pwd = pwd
        self.key = key


class SSHClient(object):
    ssh = None
    ftp = None
    opt: SSHOpt

    def __init__(self, opt: SSHOpt):
        self.opt = opt
        self.ssh = paramiko.SSHClient()
        if (len(opt.key) > 0):
            private_key = paramiko.RSAKey.from_private_key_file(self.opt.key)
            self.ssh.set_missing_host_key_policy(paramiko.AutoAddPolicy())
            self.ssh.connect(opt.ip, opt.port, opt.user, pkey=private_key)
        elif (len(opt.pwd) > 0):
            self.ssh.connect(opt.ip, opt.port, opt.user, opt.pwd)
        else:
            raise Exception("Fail to connect remote server")

    def exec(self, cmd: str) -> None:
        stdin, stdout, stderr = self.ssh.exec_command(cmd)
        res = stdout.read().decode()
        print(res)

    def get(self, src: str, dst: str) -> None:
        if (self.ftp == None):
            self.ftp = self.ssh.open_sftp()
        self.ftp.get(src, dst)

    def put(self, src: str, dst: str) -> None:
        if (self.ftp == None):
            self.ftp = self.ssh.open_sftp()
        self.ftp.put(src, dst)

    def close(self) -> None:
        self.ssh.close()


if __name__ == '__main__':
    def remote_deploy():
        opt = SSHOpt('0.0.0.0', 22, "", '', key='')
        cli = SSHClient(opt)
        cli.exec("ls")
        cli.close()
    threadPool = []
    thread = threading.Thread(target=remote_deploy)
    threadPool.append(thread)
    for td in threadPool:
        td.start()
    for td in threadPool:
        td.join()
