import requests
import json
import websockets
import asyncio


class HttpClient:
    addr: str

    def __init__(self, addr: str):
        self.addr = addr

    def get(self, url: str, **args):
        params = args.get('params')
        headers = args.get('headers')
        try:
            res = requests.get(self.addr + url, params=params, headers=headers)
            return res
        except Exception as e:
            print("get err: %s" % e)

    def postWithResult(self, url: str, **args):
        params = args.get('params')
        headers = args.get('headers')
        data = args.get('data')
        try:
            res = requests.post(self.addr + url, data=data,
                                params=params, headers=headers)
            return json.loads(res.text)
        except Exception as e:
            print("get err: %s" % e)

    def postWithData(self,  url: str, **args) -> dict:
        params = args.get('params')
        headers = args.get('headers')
        data = args.get('data')
        try:
            res = requests.post(self.addr + url, data=data,
                                params=params, headers=headers)
            return json.loads(json.loads(res.text)['data'])
        except Exception as e:
            print("get err: %s" % e)

    def group(self, url: str):
        return HttpClient(self.addr + url)


class WsClient:
    url: str

    def __init__(self, url: str):
        self.url = url

    async def echo(self):
        async with websockets.connect(self.url) as ws:
            await ws.send('hello')
            recv_text = await ws.recv()
            print(recv_text)
