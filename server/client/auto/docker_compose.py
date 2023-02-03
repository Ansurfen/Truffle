import yaml

dockerComposeTemplate: str = '''
version: '2'
'''


class DockerCompose:
    _template: any

    def __init__(self, version: str):
        self._template = yaml.load(
            dockerComposeTemplate, Loader=yaml.FullLoader)
        self._template['version'] = version

    def services(self, name: str, image: str, restart: bool, ports: [], networks: []) -> None:
        try:
            if(type(self._template['services']) == type({})):
                pass
        except:
            self._template['services'] = {}
        self._template['services'][name] = {
            "image": image,
            "ports": ports,
            "networks": networks
        }

    def networks(self, exNet: [] = ...) -> None:
        try:
            if(type(self._template['networks']) == type({})):
                pass
        except:
            self._template['networks'] = {}
        if (not exNet is Ellipsis):
            for net in exNet:
                self._template['networks'][net] = {"external": True}

    def gen(self, dir: str) -> None:
        with open('{dir}/docker-compose.yml'.format(dir=dir), 'w+', encoding='utf-8') as fp:
            yaml.dump(self._template, stream=fp, allow_unicode=True)
            fp.close()

    def show(self) -> None:
        print(yaml.dump(self._template))


def TestDockerCompose() -> None:
    dc = DockerCompose('2')
    dc.services('truffle_i18n', 'ansurfen/truffle_i18n:0.0.2',
                True, ['8000-8001:8000'], ['truffle'])
    dc.networks(['truffle'])
    dc.gen('.')


def DockerComposeTemplate(name: str, image: str, ports: [], networks: [] = ...) -> DockerCompose:
    dc = DockerCompose('2')
    dc.services('truffle_{name}'.format(name=name), image,
                True, ports, networks)
    dc.networks(networks)
    return dc


# TestDockerCompose()
dc = DockerComposeTemplate(
    'i18n', 'ansurfen/truffle_i18n:0.0.2', ['8000-8001:8000'], ['truffle'])
dc.show()
