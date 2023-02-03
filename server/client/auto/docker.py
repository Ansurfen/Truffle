import queue


class Dockerfile:
    flows: queue.Queue

    def __init__(self):
        self.flows = queue.Queue()

    def image(self, image: str) -> None:
        self.flows.put('FROM {image}'.format(image=image))

    def copy(self, src: str, dst: str) -> None:
        self.flows.put('COPY {src} {dst}'.format(src=src, dst=dst))

    def run(self, cmd: str) -> None:
        self.flows.put('RUN {cmd}'.format(cmd=cmd))

    def expose(self, ports: []) -> None:
        for port in ports:
            self.flows.put('EXPOSE {port}'.format(port=port))

    def clone(self):
        dockerfile_copy = Dockerfile()
        temp = queue.Queue()
        while not self.flows.empty():
            flow: str = self.flows.get()
            dockerfile_copy.flows.put(flow)
            temp.put(flow)
        self.flows = temp
        return dockerfile_copy

    def cmd(self, cmds: []) -> None:
        self.flows.put('CMD [{cmds}]'.format(
            cmds=','.join('"{0}"'.format(cmd) for cmd in cmds)))

    def volume(self, volumes: []) -> None:
        self.flows.put('VOLUME [{volume}]'.format(
            volume=','.join('"{0}"'.format(volume) for volume in volumes)))

    def workdir(self, dir: str) -> None:
        self.flows.put('WORKDIR {dir}'.format(dir=dir))

    def nop(self) -> None:
        self.flows.put('')

    def show(self) -> None:
        temp = queue.Queue()
        while not self.flows.empty():
            flow: str = self.flows.get()
            print(flow)
            temp.put(flow)
        self.flows = temp

    def gen(self, dir: str) -> None:
        with open('{dir}/dockerfile'.format(dir=dir), 'w+', encoding='utf-8') as fp:
            flows: str = ''
            while not self.flows.empty():
                flows += self.flows.get() + '\n'
            fp.write(flows)
            fp.close()


def DockerfileTemplate(image: str, port: int) -> Dockerfile:
    dockerfile = Dockerfile()
    dockerfile.image(image)
    dockerfile.nop()
    dockerfile.expose([port])
    dockerfile.nop()
    dockerfile.copy('main', '.')
    dockerfile.copy('application.yaml', '.')
    dockerfile.nop()
    dockerfile.run('chmod +x ./main')
    dockerfile.run('mkdir ./log -p')
    return dockerfile


def TestDockerfile() -> None:
    server = Dockerfile()
    server.image('ubuntu:20.04')
    server.expose([8080])
    server.run('apt update -y && apt install -y procps')
    server.copy('main', '.')
    server.copy('application.yaml', '.')
    server.run('chmod +x ./main')
    server.run('mkdir ./log -p')
    server.run('mkdir ./lang -p')
    server.copy('/lang/en_us.yam', '/lang')
    server.copy('/lang/zh_cn.yaml', '/lang')
    server.volume(['/etc/nginx', '/var/log/nginx', '/opt'])
    server.workdir('/opt')
    server.cmd(['./main'])
    server.show()
    server.gen('.')


# TestDockerfile()
doc = DockerfileTemplate('ubuntu:20.04', 8080)
doc.show()
